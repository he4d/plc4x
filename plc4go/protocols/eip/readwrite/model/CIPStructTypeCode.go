/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// CIPStructTypeCode is an enum
type CIPStructTypeCode uint16

type ICIPStructTypeCode interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	CIPStructTypeCode_STRING CIPStructTypeCode = 0x0FCE
)

var CIPStructTypeCodeValues []CIPStructTypeCode

func init() {
	_ = errors.New
	CIPStructTypeCodeValues = []CIPStructTypeCode{
		CIPStructTypeCode_STRING,
	}
}

func CIPStructTypeCodeByValue(value uint16) (enum CIPStructTypeCode, ok bool) {
	switch value {
	case 0x0FCE:
		return CIPStructTypeCode_STRING, true
	}
	return 0, false
}

func CIPStructTypeCodeByName(value string) (enum CIPStructTypeCode, ok bool) {
	switch value {
	case "STRING":
		return CIPStructTypeCode_STRING, true
	}
	return 0, false
}

func CIPStructTypeCodeKnows(value uint16) bool {
	for _, typeValue := range CIPStructTypeCodeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCIPStructTypeCode(structType interface{}) CIPStructTypeCode {
	castFunc := func(typ interface{}) CIPStructTypeCode {
		if sCIPStructTypeCode, ok := typ.(CIPStructTypeCode); ok {
			return sCIPStructTypeCode
		}
		return 0
	}
	return castFunc(structType)
}

func (m CIPStructTypeCode) GetLengthInBits() uint16 {
	return 16
}

func (m CIPStructTypeCode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CIPStructTypeCodeParse(readBuffer utils.ReadBuffer) (CIPStructTypeCode, error) {
	val, err := readBuffer.ReadUint16("CIPStructTypeCode", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CIPStructTypeCode")
	}
	if enum, ok := CIPStructTypeCodeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return CIPStructTypeCode(val), nil
	} else {
		return enum, nil
	}
}

func (e CIPStructTypeCode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("CIPStructTypeCode", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CIPStructTypeCode) PLC4XEnumName() string {
	switch e {
	case CIPStructTypeCode_STRING:
		return "STRING"
	}
	return ""
}

func (e CIPStructTypeCode) String() string {
	return e.PLC4XEnumName()
}
