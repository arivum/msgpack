/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

import "unsafe"

func (d *Decoder) readLen16() (int, error) {
	var err error

	if d.twoBytes[0], err = d.r.ReadByte(); err != nil {
		return 0, err
	}
	d.twoBytes[1], err = d.r.ReadByte()

	return int(d.twoBytes[0])<<8 | int(d.twoBytes[1]), err
}

func (d *Decoder) readLen32() (int, error) {
	var err error

	if _, err = d.r.Read(d.fourBytes[:]); err != nil {
		return 0, err
	}

	return int(d.fourBytes[0])<<24 | int(d.fourBytes[1])<<16 | int(d.fourBytes[2])<<8 | int(d.fourBytes[3]), nil
}

func (d *Decoder) readFull(buf unsafe.Pointer) error {
	var (
		index = 0
		n     = 0
		err   error

		strLen = len((*(*[]byte)(buf)))
	)

	for index < strLen {
		if n, err = d.r.Read((*(*[]byte)(buf))[index:]); err != nil {
			return err
		}
		index += n
	}

	return nil
}

//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	var x = uintptr(p)

	return unsafe.Pointer(x ^ 0)
}
