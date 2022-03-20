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
	"sync"
)

type Decoder struct {
	buf                  *preallocBuf
	underlayingReader    io.Reader
	r                    *bufio.Reader
	s                    chan interface{}
	mapDecoderFuncsByLen []func() (map[string]interface{}, error)
	lastError            error
	twoBytes             [2]byte
	fourBytes            [4]byte
	eightBytes           [8]byte
	nextByte             byte
	closeOnce            sync.Once
	closeChan            chan struct{}
}

type preallocBuf struct {
	buf   []byte
	index int
}
