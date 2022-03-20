/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

func newPreallocBuf() *preallocBuf {
	return &preallocBuf{
		buf:   make([]byte, 1024),
		index: 0,
	}
}

func (g *preallocBuf) allocateBuffer(len int) []byte {
	if g.index+len < 1024 {
		g.index += len
		return g.buf[g.index-len : g.index]
	}
	if len >= 1024 {
		return make([]byte, len)
	}
	g.buf = make([]byte, 1024)
	g.index = len
	return g.buf[0:len]
}
