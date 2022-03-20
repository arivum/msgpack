/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package main

import (
	"io"
	"os"
)

func (j *JSONFileInputAdapter) GetReader() (io.Reader, error) {
	var (
		file, err = os.OpenFile(j.Filename, os.O_RDONLY, 0600)
		r, w      = io.Pipe()
		fInfo     os.FileInfo
	)

	if fInfo, err = os.Stat(j.Filename); err != nil {
		return nil, err
	}

	j.filesize = fInfo.Size()

	go j.readLoop(w, file)
	//go j.printStats()

	return r, err
}
