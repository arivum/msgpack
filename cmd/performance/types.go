/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package main

type JSONFileInputAdapter struct {
	Filename                string `yaml:"filename" jsonschema:"required"`
	readBytes, oldReadBytes float64
	filesize                int64
}

type readerFunc func(p []byte) (n int, err error)

func (read readerFunc) Read(p []byte) (n int, err error) {
	return read(p)
}
