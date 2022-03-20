/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

func (d *Decoder) decodeFixMap() (map[string]interface{}, error) {
	var len = d.nextByte ^ msgPackFlagFixMap

	if len < 4 {
		return d.mapDecoderFuncsByLen[len]()
	}

	return d.decodeMapAny(int(len))
}

func (d *Decoder) decodeMap16() (map[string]interface{}, error) {
	var len, err = d.readLen16()

	if err != nil {
		return nil, err
	}

	return d.decodeMapAny(len)
}

func (d *Decoder) decodeMap32() (map[string]interface{}, error) {
	var len, err = d.readLen32()

	if err != nil {
		return nil, err
	}

	return d.decodeMapAny(len)
}

func (d *Decoder) decodeFixSlice() ([]interface{}, error) {
	return d.decodeSlice(int(d.nextByte ^ msgPackFlagFixArray))
}

func (d *Decoder) decodeSlice16() ([]interface{}, error) {
	var len, err = d.readLen16()

	if err != nil {
		return nil, err
	}

	return d.decodeSlice(len)
}

func (d *Decoder) decodeSlice32() ([]interface{}, error) {
	var len, err = d.readLen32()

	if err != nil {
		return nil, err
	}

	return d.decodeSlice(len)
}

func (d *Decoder) decodeSlice(len int) ([]interface{}, error) {
	var (
		err error
		s   = make([]interface{}, len)
		i   = 0
	)

	for ; i < len; i++ {
		var value interface{}
		if err = d.handleNext(&value); err != nil {
			return nil, err
		}
		s[i] = value
	}

	return s, nil
}
