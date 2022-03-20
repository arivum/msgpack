/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/arivum/json2msgpackStreamer"
	arivumMsgPack "github.com/arivum/msgpack"
	"github.com/vmihailenco/msgpack/v5"
)

var (
	filename = ""
	d        = 1 * time.Second
)

type Stat struct {
	MBPerSec float64
	Time     time.Time
}

type Test struct {
	Name    string
	Stats   []Stat
	RunFunc func(io.Reader) `json:"-"`
}

func (t *Test) Run() error {
	var (
		err   error
		avg   float64
		input = JSONFileInputAdapter{
			Filename: filename,
		}
		r io.Reader
	)

	if r, err = input.GetReader(); err != nil {
		log.Panic(err)
	}

	go t.RunFunc(r)

	for {
		time.Sleep(d)
		if avg = input.getMBperSec(d); avg == 0 {
			break
		}
		t.Stats = append(t.Stats, Stat{
			Time:     time.Now(),
			MBPerSec: avg,
		})
		fmt.Println(avg)
	}
	return nil
}

var tests = []*Test{
	{
		Name:  "converter+arivumMsgPack",
		Stats: make([]Stat, 0),
		RunFunc: func(r io.Reader) {
			json2msgpack := json2msgpackStreamer.NewJSON2MsgPackStreamer(r)
			msgpackDec := arivumMsgPack.NewDecoder(json2msgpack)

			for range msgpackDec.Stream() {
				if msgpackDec.LastError() != nil {
					break
				}
			}
		},
	},
	{
		Name:  "converter+vmihailenco",
		Stats: make([]Stat, 0),
		RunFunc: func(r io.Reader) {
			json2msgpack := json2msgpackStreamer.NewJSON2MsgPackStreamer(r)
			msgpackDec := msgpack.NewDecoder(json2msgpack)

			for {
				var entry interface{}

				if err := msgpackDec.Decode(&entry); err != nil {
					break
				}
			}
		},
	},
	{
		Name:  "converter-only",
		Stats: make([]Stat, 0),
		RunFunc: func(r io.Reader) {
			json2msgpack := json2msgpackStreamer.NewJSON2MsgPackStreamer(r)
			bufR := bufio.NewReader(json2msgpack)
			for {
				if _, err := bufR.ReadByte(); err != nil {
					break
				}
			}
		},
	},
	{
		Name:  "encoding/json",
		Stats: make([]Stat, 0),
		RunFunc: func(r io.Reader) {
			var (
				jsonDec = json.NewDecoder(r)
				err     error
			)

			for {
				var v interface{}
				if err = jsonDec.Decode(&v); err != nil {
					fmt.Println(err)
					break
				}
			}
		},
	},
}

func parse() {
	flag.StringVar(&filename, "jsonfile", "", "(required) JSON file to run performance tests on")
	flag.Parse()
	if len(filename) == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	parse()
	for _, test := range tests {
		fmt.Printf("Running performance test for %s\n", test.Name)
		test.Run()
	}

	if values, err := json.Marshal(tests); err != nil {
		panic(err)
	} else {
		fmt.Println(tests)
		if err = ioutil.WriteFile("performance.json", values, 0777); err != nil {
			log.Panic(err)
		}
	}
}
