/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

import "unsafe"

func (d *Decoder) readLen16() (int, error) {
	var _, err = d.r.Read(d.twoBytes[:])
	return int(d.twoBytes[0])<<8 | int(d.twoBytes[1]), err
}

func (d *Decoder) readLen32() (int, error) {
	var _, err = d.r.Read(d.fourBytes[:])
	return int(d.fourBytes[0])<<24 | int(d.fourBytes[1])<<16 | int(d.fourBytes[2])<<8 | int(d.fourBytes[3]), err
}

func (d *Decoder) readFull(buf unsafe.Pointer) error {
	var (
		err    error
		index  = 0
		n      = 0
		strLen = len((*(*[]byte)(buf)))
		b      = *(*[]byte)(buf)
	)

	for index < strLen {
		if n, err = d.r.Read(b[index:]); err != nil {
			return err
		}
		index += n
	}

	return err
}

//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	var x = uintptr(p)

	return unsafe.Pointer(x ^ 0)
}
