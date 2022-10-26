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

// SecurityCommandType is an enum
type SecurityCommandType uint8

type ISecurityCommandType interface {
	NumberOfArguments() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	SecurityCommandType_OFF   SecurityCommandType = 0x00
	SecurityCommandType_ON    SecurityCommandType = 0x01
	SecurityCommandType_EVENT SecurityCommandType = 0x02
)

var SecurityCommandTypeValues []SecurityCommandType

func init() {
	_ = errors.New
	SecurityCommandTypeValues = []SecurityCommandType{
		SecurityCommandType_OFF,
		SecurityCommandType_ON,
		SecurityCommandType_EVENT,
	}
}

func (e SecurityCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 0
		}
	case 0x01:
		{ /* '0x01' */
			return 1
		}
	case 0x02:
		{ /* '0x02' */
			return 0xFF
		}
	default:
		{
			return 0
		}
	}
}

func SecurityCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (SecurityCommandType, error) {
	for _, sizeValue := range SecurityCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func SecurityCommandTypeByValue(value uint8) (enum SecurityCommandType, ok bool) {
	switch value {
	case 0x00:
		return SecurityCommandType_OFF, true
	case 0x01:
		return SecurityCommandType_ON, true
	case 0x02:
		return SecurityCommandType_EVENT, true
	}
	return 0, false
}

func SecurityCommandTypeByName(value string) (enum SecurityCommandType, ok bool) {
	switch value {
	case "OFF":
		return SecurityCommandType_OFF, true
	case "ON":
		return SecurityCommandType_ON, true
	case "EVENT":
		return SecurityCommandType_EVENT, true
	}
	return 0, false
}

func SecurityCommandTypeKnows(value uint8) bool {
	for _, typeValue := range SecurityCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastSecurityCommandType(structType interface{}) SecurityCommandType {
	castFunc := func(typ interface{}) SecurityCommandType {
		if sSecurityCommandType, ok := typ.(SecurityCommandType); ok {
			return sSecurityCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m SecurityCommandType) GetLengthInBits() uint16 {
	return 4
}

func (m SecurityCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityCommandTypeParse(readBuffer utils.ReadBuffer) (SecurityCommandType, error) {
	val, err := readBuffer.ReadUint8("SecurityCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading SecurityCommandType")
	}
	if enum, ok := SecurityCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return SecurityCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e SecurityCommandType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("SecurityCommandType", 4, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e SecurityCommandType) PLC4XEnumName() string {
	switch e {
	case SecurityCommandType_OFF:
		return "OFF"
	case SecurityCommandType_ON:
		return "ON"
	case SecurityCommandType_EVENT:
		return "EVENT"
	}
	return ""
}

func (e SecurityCommandType) String() string {
	return e.PLC4XEnumName()
}
