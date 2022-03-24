/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

func (d *Decoder) decodeSliceEnrolledSingleEntry() ([]interface{}, error) {
	var (
		err   error
		value interface{}
	)

	// if d.nextByte, err = d.r.ReadByte(); err != nil {
	// 	return nil, err
	// }
	// if key, err = d.directDecodeString(); err != nil {
	// 	return nil, err
	// }
	err = d.handleNext(&value)

	return []interface{}{
		value,
	}, err
}

func (d *Decoder) decodeSliceEnrolledTwoEntries() ([]interface{}, error) {
	var (
		value1, value2 interface{}
		err            error
	)

	if err = d.handleNext(&value1); err != nil {
		return nil, err
	}

	if err = d.handleNext(&value2); err != nil {
		return nil, err
	}

	return []interface{}{
		value1,
		value2,
	}, nil
}

func (d *Decoder) decodeSliceEnrolledThreeEntries() ([]interface{}, error) {
	var (
		value1, value2, value3 interface{}
		err                    error
	)

	if err = d.handleNext(&value1); err != nil {
		return nil, err
	}

	if err = d.handleNext(&value2); err != nil {
		return nil, err
	}

	if err = d.handleNext(&value3); err != nil {
		return nil, err
	}

	return []interface{}{
		value1,
		value2,
		value3,
	}, nil
}

func (d *Decoder) decodeSliceAny(len int) ([]interface{}, error) {
	var (
		s   = make([]interface{}, len)
		err error
		i   = 0
	)

	for ; i < len; i++ {
		//var value interface{}

		if err = d.handleNext(&s[i]); err != nil {
			return nil, err
		}

	}

	return s, nil
}
