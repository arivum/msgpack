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
	"testing"
	"unsafe"
)

func BenchmarkReadFull(b *testing.B) {
	rd, wt := io.Pipe()
	go func() {
		for i := 0; i < b.N; i++ {
			wt.Write([]byte{byte(b.N >> 24), byte(b.N >> 16), byte(b.N >> 8), byte(b.N)})
		}
	}()
	r := bufio.NewReader(rd)
	d := NewDecoder(r)
	bb := make([]byte, 4)
	buf := noescape(unsafe.Pointer(&bb))
	for i := 0; i < b.N; i++ {
		d.readFull(buf)
	}
}

func BenchmarkReadLen32(b *testing.B) {
	rd, wt := io.Pipe()
	go func() {
		for i := 0; i < b.N; i++ {
			wt.Write([]byte{byte(b.N >> 24), byte(b.N >> 16), byte(b.N >> 8), byte(b.N)})
		}
	}()
	r := bufio.NewReader(rd)
	d := NewDecoder(r)
	for i := 0; i < b.N; i++ {
		d.readLen16()
	}
}

func BenchmarkReadLen16(b *testing.B) {
	rd, wt := io.Pipe()
	go func() {
		for i := 0; i < b.N; i++ {
			wt.Write([]byte{byte(b.N >> 24), byte(b.N >> 16), byte(b.N >> 8), byte(b.N)})
		}
	}()
	r := bufio.NewReader(rd)
	d := NewDecoder(r)
	for i := 0; i < b.N; i++ {
		d.readLen16()
	}
}
