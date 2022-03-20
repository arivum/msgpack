/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package main

import (
	"fmt"
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

func (j *JSONFileInputAdapter) readLoop(w io.WriteCloser, r io.Reader) {
	var (
		err   error
		start time.Time
	)
	defer func() {
		_ = w.Close()
		//close(j.finished)
	}()

	start = time.Now()
	if _, err = io.Copy(w, readerFunc(func(p []byte) (int, error) {
		var (
			err error
			n   int
		)

		if n, err = r.Read(p); err == nil {
			j.readBytes += float64(n)
		}
		return n, err
	})); err != nil {
		logrus.Error(err)
	}
	logrus.Infof("processed file %s with an avg of %s", j.Filename, j.avgThroughput(time.Since(start)))
}

func (j *JSONFileInputAdapter) avgThroughput(duration time.Duration) string {
	var (
		throughput = float64(j.filesize) / float64(duration.Seconds())
		multiplier = 0
		units      = []string{
			"B/s",
			"kB/s",
			"MB/s",
			"GB/s",
		}
	)

	for throughput/1024 > 1 {
		throughput /= 1024
		multiplier++
	}

	return fmt.Sprintf("%.2f %s", throughput, units[multiplier])
}

func (j *JSONFileInputAdapter) getMBperSec(resolution time.Duration) float64 {

	var (
		avg = (j.readBytes - j.oldReadBytes) / (1 << 20) / resolution.Seconds()
	)

	j.oldReadBytes = j.readBytes
	return avg
}

// func (j *JSONFileInputAdapter) printStats() {
// 	var (
// 		readBytes  = 0
// 		resolution = 5 * time.Second
// 	)

// 	for {
// 		select {
// 		case <-j.finished:
// 			return
// 		default:
// 			time.Sleep(resolution)
// 			logrus.Infof("file %s processed: %.2f%%, avg: %.2fMB/s", j.Filename, j.readBytes*100/float64(j.filesize), (j.readBytes-float64(readBytes))/(1<<20)/resolution.Seconds())
// 			readBytes = int(j.readBytes)
// 		}
// 	}
// }
