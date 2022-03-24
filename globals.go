/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

const (
	msgPackFlagFixStr    = byte(0xA0)
	msgPackFlagStr8      = byte(0xd9)
	msgPackFlagStr16     = byte(0xda)
	msgPackFlagStr32     = byte(0xdb)
	msgPackFlagFixArray  = byte(0x90)
	msgPackFlagArray16   = byte(0xdc)
	msgPackFlagArray32   = byte(0xdd)
	msgPackFlagFixMap    = byte(0x80)
	msgPackFlagMap16     = byte(0xde)
	msgPackFlagMap32     = byte(0xdf)
	msgPackFlagFloat32   = byte(0xca)
	msgPackFlagFloat64   = byte(0xcb)
	msgPackFlagPosFixInt = byte(0x00)
	msgPackFlagNegFixInt = byte(0xe0)
	msgPackFlagInt8      = byte(0xd0)
	msgPackFlagInt16     = byte(0xd1)
	msgPackFlagInt32     = byte(0xd2)
	msgPackFlagInt64     = byte(0xd3)
	msgPackFlagUint8     = byte(0xcc)
	msgPackFlagUint16    = byte(0xcd)
	msgPackFlagUint32    = byte(0xce)
	msgPackFlagUint64    = byte(0xcf)
	msgPackFlagNil       = byte(0xc0)
	msgPackFlagBoolFalse = byte(0xc2)
	msgPackFlagBoolTrue  = byte(0xc3)
	max8BitPlusOne       = (1 << 8)
	max16BitPlusOne      = (1 << 16)
	max32BitPlusOne      = (1 << 32)
	max5BitPlusOne       = (1 << 5)
	max4BitPlusOne       = (1 << 4)
	negFixIntMask        = (-1 << 5)
	accBufSize           = 4 << 10
	accBufMax            = 1 << 10
)
