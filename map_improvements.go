/*
 * Copyright (c) 2022, arivum.
 * All rights reserved.
 * SPDX-License-Identifier: MIT
 * For full license text, see the LICENSE file in the repo root or https://opensource.org/licenses/MIT
 */

package msgpack

func (d *Decoder) decodeMapEnrolledSingleEntry() (map[string]interface{}, error) {
	var (
		err   error
		key   string
		value interface{}
	)

	if d.nextByte, err = d.r.ReadByte(); err != nil {
		return nil, err
	}
	if key, err = d.directDecodeString(); err != nil {
		return nil, err
	}
	err = d.handleNext(&value)

	return map[string]interface{}{
		key: value,
	}, err
}

func (d *Decoder) decodeMapEnrolledTwoEntries() (map[string]interface{}, error) {
	var (
		value1, value2 interface{}
		key1, key2     string
		err            error
	)

	if d.nextByte, err = d.r.ReadByte(); err != nil {
		return nil, err
	}
	if key1, err = d.directDecodeString(); err != nil {
		return nil, err
	}
	if err = d.handleNext(&value1); err != nil {
		return nil, err
	}

	if d.nextByte, err = d.r.ReadByte(); err != nil {
		return nil, err
	}
	if key2, err = d.directDecodeString(); err != nil {
		return nil, err
	}
	if err = d.handleNext(&value2); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		key1: value1,
		key2: value2,
	}, nil
}

func (d *Decoder) decodeMapEnrolledThreeEntries() (map[string]interface{}, error) {
	var (
		value1, value2, value3 interface{}
		key1, key2, key3       string
		err                    error
	)

	if d.nextByte, err = d.r.ReadByte(); err != nil {
		return nil, err
	}
	if key1, err = d.directDecodeString(); err != nil {
		return nil, err
	}
	if err = d.handleNext(&value1); err != nil {
		return nil, err
	}

	if d.nextByte, err = d.r.ReadByte(); err != nil {
		return nil, err
	}
	if key2, err = d.directDecodeString(); err != nil {
		return nil, err
	}
	if err = d.handleNext(&value2); err != nil {
		return nil, err
	}

	if d.nextByte, err = d.r.ReadByte(); err != nil {
		return nil, err
	}
	if key3, err = d.directDecodeString(); err != nil {
		return nil, err
	}
	if err = d.handleNext(&value3); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		key1: value1,
		key2: value2,
		key3: value3,
	}, nil
}

func (d *Decoder) decodeMapAny(len int) (map[string]interface{}, error) {
	var (
		m   = make(map[string]interface{}, len)
		err error
		key string
		i   = 0
	)

	for ; i < len; i++ {
		var value interface{}

		if d.nextByte, err = d.r.ReadByte(); err != nil {
			return nil, err
		}
		if key, err = d.directDecodeString(); err != nil {
			return nil, err
		}
		if err = d.handleNext(&value); err != nil {
			return nil, err
		}
		m[key] = value

	}

	return m, nil
}
