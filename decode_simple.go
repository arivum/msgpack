/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

import (
	"encoding/binary"
	"errors"
	"fmt"
	"unsafe"
)

func (d *Decoder) directDecodeFixedString() (string, error) {
	var buf = d.buf.allocateBuffer(int(d.nextByte ^ msgPackFlagFixStr))

	return *(*string)(unsafe.Pointer(&buf)), d.readFull(noescape(unsafe.Pointer(&buf)))
}

func (d *Decoder) directDecodeString8() (string, error) {
	var (
		len, _ = d.r.ReadByte()
		buf    = d.buf.allocateBuffer(int(len))
	)

	return *(*string)(unsafe.Pointer(&buf)), d.readFull(noescape(unsafe.Pointer(&buf)))
}

func (d *Decoder) directDecodeString16() (string, error) {
	var (
		len, _ = d.readLen16()
		buf    = d.buf.allocateBuffer(len)
	)

	return *(*string)(unsafe.Pointer(&buf)), d.readFull(noescape(unsafe.Pointer(&buf)))
}

func (d *Decoder) directDecodeString32() (string, error) {
	var (
		len, err = d.readLen32()
		buf      = d.buf.allocateBuffer(len)
	)

	if err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&buf)), d.readFull(noescape(unsafe.Pointer(&buf)))
}

func (d *Decoder) directDecodeString() (string, error) {
	if (d.nextByte & 0xE0) == msgPackFlagFixStr {
		return d.directDecodeFixedString()
	}
	switch d.nextByte {
	case msgPackFlagStr8:
		return d.directDecodeString8()
	case msgPackFlagStr16:
		return d.directDecodeString16()
	case msgPackFlagStr32:
		return d.directDecodeString32()
	default:
		return "", errors.New("string type unknown")
	}
}

func (d *Decoder) decodeSimple() (interface{}, error) {
	var (
		err error
	)
	switch {
	case d.nextByte == msgPackFlagFloat32:
		var num float32
		if err = binary.Read(d.r, binary.BigEndian, &num); err != nil {
			return nil, err
		}
		return num, nil
	case d.nextByte == msgPackFlagFloat64:
		var num float64
		if err = binary.Read(d.r, binary.BigEndian, &num); err != nil {
			return nil, err
		}
		return num, nil
	case d.nextByte == msgPackFlagUint8:
		return d.r.ReadByte()
	case d.nextByte == msgPackFlagUint16:
		return d.readUint16()
	case d.nextByte == msgPackFlagUint32:
		return d.readUint32()
	case d.nextByte == msgPackFlagUint64:
		return d.readUint64()
	case d.nextByte == msgPackFlagInt8:
		var num byte
		num, err = d.r.ReadByte()
		return int8(num), err
	case d.nextByte == msgPackFlagInt16:
		return d.readInt16()
	case d.nextByte == msgPackFlagInt32:
		return d.readInt32()
	case d.nextByte == msgPackFlagInt64:
		return d.readInt64()
	case (d.nextByte>>7)<<7 == msgPackFlagPosFixInt:
		return int(d.nextByte ^ msgPackFlagPosFixInt), nil
	case (d.nextByte>>5)<<5 == msgPackFlagNegFixInt:
		return int(d.nextByte) | negFixIntMask, nil
	default:
		return nil, fmt.Errorf("could not decode type %x", d.nextByte)
	}
}

func (d *Decoder) readUint16() (uint16, error) {
	var _, err = d.r.Read(d.twoBytes[:])
	return uint16(d.twoBytes[0])<<8 | uint16(d.twoBytes[1]), err
}

func (d *Decoder) readUint32() (uint32, error) {
	var _, err = d.r.Read(d.fourBytes[:])
	return uint32(d.fourBytes[0])<<24 | uint32(d.fourBytes[1])<<16 | uint32(d.fourBytes[2])<<8 | uint32(d.fourBytes[3]), err
}

func (d *Decoder) readUint64() (uint64, error) {
	var _, err = d.r.Read(d.eightBytes[:])
	return uint64(d.eightBytes[0])<<56 | uint64(d.eightBytes[1])<<48 | uint64(d.eightBytes[2])<<40 | uint64(d.eightBytes[3])<<32 | uint64(d.eightBytes[4])<<24 | uint64(d.eightBytes[5])<<16 | uint64(d.eightBytes[6])<<8 | uint64(d.eightBytes[7]), err
}

func (d *Decoder) readInt16() (int16, error) {
	var _, err = d.r.Read(d.twoBytes[:])
	return int16(d.twoBytes[0])<<8 | int16(d.twoBytes[1]), err
}

func (d *Decoder) readInt32() (int32, error) {
	var _, err = d.r.Read(d.fourBytes[:])
	return int32(d.fourBytes[0])<<24 | int32(d.fourBytes[1])<<16 | int32(d.fourBytes[2])<<8 | int32(d.fourBytes[3]), err
}

func (d *Decoder) readInt64() (int64, error) {
	var _, err = d.r.Read(d.eightBytes[:])
	return int64(d.eightBytes[0])<<56 | int64(d.eightBytes[1])<<48 | int64(d.eightBytes[2])<<40 | int64(d.eightBytes[3])<<32 | int64(d.eightBytes[4])<<24 | int64(d.eightBytes[5])<<16 | int64(d.eightBytes[6])<<8 | int64(d.eightBytes[7]), err
}
