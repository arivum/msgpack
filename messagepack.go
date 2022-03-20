/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

import (
	"bufio"
	"io"
	"time"
)

func NewDecoder(r io.Reader) *Decoder {
	var (
		buf = bufio.NewReader(r)
		d   = &Decoder{
			underlayingReader: r,
			r:                 buf,
			s:                 make(chan interface{}, 100),
			buf:               newPreallocBuf(),
		}
	)
	d.mapDecoderFuncsByLen = []func() (map[string]interface{}, error){
		func() (map[string]interface{}, error) {
			return emptyMap, nil
		},
		d.decodeMapEnrolledSingleEntry,
		d.decodeMapEnrolledTwoEntries,
		d.decodeMapEnrolledThreeEntries,
	}
	return d
}

func (d *Decoder) Stream() chan interface{} {
	go func() {
		for {
			select {
			case <-d.closeChan:
				d.Close()
			default:
				var v interface{}
				if d.lastError = d.handleNext(&v); d.lastError != nil {
					if v != nil {
						d.s <- v
					}
					d.Close()
					return
				}
				d.s <- v
			}
		}
	}()

	return d.s
}

func (d *Decoder) Close() {
	d.closeOnce.Do(func() {
		for len(d.s) > 0 {
			time.Sleep(10 * time.Microsecond)
		}
		close(d.s)
	})
}

func (d *Decoder) LastError() error {
	if d.lastError == io.EOF {
		d.lastError = nil
	}
	return d.lastError
}

func (d *Decoder) handleNext(v *interface{}) error {
	var err error

	if d.nextByte, err = d.r.ReadByte(); err != nil {
		return err
	}

	if (d.nextByte & 0xE0) == msgPackFlagFixStr {
		*v, err = d.directDecodeFixedString()
		return err
	}

	switch d.nextByte & 0xF0 {
	case msgPackFlagFixMap:
		*v, err = d.decodeFixMap()
		return err
	case msgPackFlagFixArray:
		*v, err = d.decodeFixSlice()
		return err
	}

	switch d.nextByte {
	case msgPackFlagMap16:
		*v, err = d.decodeMap16()
		return err
	case msgPackFlagMap32:
		*v, err = d.decodeMap32()
		return err
	case msgPackFlagArray16:
		*v, err = d.decodeSlice16()
		return err
	case msgPackFlagArray32:
		*v, err = d.decodeSlice32()
		return err
	case msgPackFlagStr8:
		*v, err = d.directDecodeString8()
		return err
	case msgPackFlagStr16:
		*v, err = d.directDecodeString16()
		return err
	case msgPackFlagStr32:
		*v, err = d.directDecodeString32()
		return err
	case msgPackFlagNil:
		*v = nil
		return nil
	case msgPackFlagBoolFalse:
		*v = false
		return nil
	case msgPackFlagBoolTrue:
		*v = true
		return nil
	default:
		*v, err = d.decodeSimple()
		return err
	}
}
