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

// BACnetEscalatorOperationDirection is an enum
type BACnetEscalatorOperationDirection uint16

type IBACnetEscalatorOperationDirection interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetEscalatorOperationDirection_UNKNOWN                  BACnetEscalatorOperationDirection = 0
	BACnetEscalatorOperationDirection_STOPPED                  BACnetEscalatorOperationDirection = 1
	BACnetEscalatorOperationDirection_UP_RATED_SPEED           BACnetEscalatorOperationDirection = 2
	BACnetEscalatorOperationDirection_UP_REDUCED_SPEED         BACnetEscalatorOperationDirection = 3
	BACnetEscalatorOperationDirection_DOWN_RATED_SPEED         BACnetEscalatorOperationDirection = 4
	BACnetEscalatorOperationDirection_DOWN_REDUCED_SPEED       BACnetEscalatorOperationDirection = 5
	BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE BACnetEscalatorOperationDirection = 0xFFFF
)

var BACnetEscalatorOperationDirectionValues []BACnetEscalatorOperationDirection

func init() {
	_ = errors.New
	BACnetEscalatorOperationDirectionValues = []BACnetEscalatorOperationDirection{
		BACnetEscalatorOperationDirection_UNKNOWN,
		BACnetEscalatorOperationDirection_STOPPED,
		BACnetEscalatorOperationDirection_UP_RATED_SPEED,
		BACnetEscalatorOperationDirection_UP_REDUCED_SPEED,
		BACnetEscalatorOperationDirection_DOWN_RATED_SPEED,
		BACnetEscalatorOperationDirection_DOWN_REDUCED_SPEED,
		BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetEscalatorOperationDirectionByValue(value uint16) (enum BACnetEscalatorOperationDirection, ok bool) {
	switch value {
	case 0:
		return BACnetEscalatorOperationDirection_UNKNOWN, true
	case 0xFFFF:
		return BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetEscalatorOperationDirection_STOPPED, true
	case 2:
		return BACnetEscalatorOperationDirection_UP_RATED_SPEED, true
	case 3:
		return BACnetEscalatorOperationDirection_UP_REDUCED_SPEED, true
	case 4:
		return BACnetEscalatorOperationDirection_DOWN_RATED_SPEED, true
	case 5:
		return BACnetEscalatorOperationDirection_DOWN_REDUCED_SPEED, true
	}
	return 0, false
}

func BACnetEscalatorOperationDirectionByName(value string) (enum BACnetEscalatorOperationDirection, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetEscalatorOperationDirection_UNKNOWN, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE, true
	case "STOPPED":
		return BACnetEscalatorOperationDirection_STOPPED, true
	case "UP_RATED_SPEED":
		return BACnetEscalatorOperationDirection_UP_RATED_SPEED, true
	case "UP_REDUCED_SPEED":
		return BACnetEscalatorOperationDirection_UP_REDUCED_SPEED, true
	case "DOWN_RATED_SPEED":
		return BACnetEscalatorOperationDirection_DOWN_RATED_SPEED, true
	case "DOWN_REDUCED_SPEED":
		return BACnetEscalatorOperationDirection_DOWN_REDUCED_SPEED, true
	}
	return 0, false
}

func BACnetEscalatorOperationDirectionKnows(value uint16) bool {
	for _, typeValue := range BACnetEscalatorOperationDirectionValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetEscalatorOperationDirection(structType interface{}) BACnetEscalatorOperationDirection {
	castFunc := func(typ interface{}) BACnetEscalatorOperationDirection {
		if sBACnetEscalatorOperationDirection, ok := typ.(BACnetEscalatorOperationDirection); ok {
			return sBACnetEscalatorOperationDirection
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetEscalatorOperationDirection) GetLengthInBits() uint16 {
	return 16
}

func (m BACnetEscalatorOperationDirection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEscalatorOperationDirectionParse(readBuffer utils.ReadBuffer) (BACnetEscalatorOperationDirection, error) {
	val, err := readBuffer.ReadUint16("BACnetEscalatorOperationDirection", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetEscalatorOperationDirection")
	}
	if enum, ok := BACnetEscalatorOperationDirectionByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetEscalatorOperationDirection(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetEscalatorOperationDirection) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("BACnetEscalatorOperationDirection", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetEscalatorOperationDirection) PLC4XEnumName() string {
	switch e {
	case BACnetEscalatorOperationDirection_UNKNOWN:
		return "UNKNOWN"
	case BACnetEscalatorOperationDirection_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetEscalatorOperationDirection_STOPPED:
		return "STOPPED"
	case BACnetEscalatorOperationDirection_UP_RATED_SPEED:
		return "UP_RATED_SPEED"
	case BACnetEscalatorOperationDirection_UP_REDUCED_SPEED:
		return "UP_REDUCED_SPEED"
	case BACnetEscalatorOperationDirection_DOWN_RATED_SPEED:
		return "DOWN_RATED_SPEED"
	case BACnetEscalatorOperationDirection_DOWN_REDUCED_SPEED:
		return "DOWN_REDUCED_SPEED"
	}
	return ""
}

func (e BACnetEscalatorOperationDirection) String() string {
	return e.PLC4XEnumName()
}
