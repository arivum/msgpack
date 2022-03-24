/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

func (d *Decoder) decodeFixMap() (map[string]interface{}, error) {
	var len = int(d.nextByte ^ msgPackFlagFixMap)

	if len < 4 {
		return d.mapDecoderFuncsByLen[len]()
	}

	return d.decodeMapAny(len)
}

func (d *Decoder) decodeMap16() (map[string]interface{}, error) {
	var len, err = d.readLen16()

	if err != nil {
		return map[string]interface{}{}, err
	}

	return d.decodeMapAny(len)
}

func (d *Decoder) decodeMap32() (map[string]interface{}, error) {
	var len, err = d.readLen32()

	if err != nil {
		return map[string]interface{}{}, err
	}

	return d.decodeMapAny(len)
}

func (d *Decoder) decodeFixSlice() ([]interface{}, error) {
	var len = int(d.nextByte ^ msgPackFlagFixArray)

	if len < 4 {
		return d.sliceDecoderFuncsByLen[len]()
	}

	return d.decodeSliceAny(len)
}

func (d *Decoder) decodeSlice16() ([]interface{}, error) {
	var len, err = d.readLen16()

	if err != nil {
		return []interface{}{}, err
	}

	return d.decodeSliceAny(len)
}

func (d *Decoder) decodeSlice32() ([]interface{}, error) {
	var len, err = d.readLen32()

	if err != nil {
		return []interface{}{}, err
	}

	return d.decodeSliceAny(len)
}
