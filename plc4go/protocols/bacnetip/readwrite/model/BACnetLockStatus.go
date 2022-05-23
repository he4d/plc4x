/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLockStatus is an enum
type BACnetLockStatus uint8

type IBACnetLockStatus interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetLockStatus_LOCKED     BACnetLockStatus = 0
	BACnetLockStatus_UNLOCKED   BACnetLockStatus = 1
	BACnetLockStatus_LOCK_FAULT BACnetLockStatus = 2
	BACnetLockStatus_UNUSED     BACnetLockStatus = 3
	BACnetLockStatus_UNKNOWN    BACnetLockStatus = 4
)

var BACnetLockStatusValues []BACnetLockStatus

func init() {
	_ = errors.New
	BACnetLockStatusValues = []BACnetLockStatus{
		BACnetLockStatus_LOCKED,
		BACnetLockStatus_UNLOCKED,
		BACnetLockStatus_LOCK_FAULT,
		BACnetLockStatus_UNUSED,
		BACnetLockStatus_UNKNOWN,
	}
}

func BACnetLockStatusByValue(value uint8) BACnetLockStatus {
	switch value {
	case 0:
		return BACnetLockStatus_LOCKED
	case 1:
		return BACnetLockStatus_UNLOCKED
	case 2:
		return BACnetLockStatus_LOCK_FAULT
	case 3:
		return BACnetLockStatus_UNUSED
	case 4:
		return BACnetLockStatus_UNKNOWN
	}
	return 0
}

func BACnetLockStatusByName(value string) BACnetLockStatus {
	switch value {
	case "LOCKED":
		return BACnetLockStatus_LOCKED
	case "UNLOCKED":
		return BACnetLockStatus_UNLOCKED
	case "LOCK_FAULT":
		return BACnetLockStatus_LOCK_FAULT
	case "UNUSED":
		return BACnetLockStatus_UNUSED
	case "UNKNOWN":
		return BACnetLockStatus_UNKNOWN
	}
	return 0
}

func BACnetLockStatusKnows(value uint8) bool {
	for _, typeValue := range BACnetLockStatusValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLockStatus(structType interface{}) BACnetLockStatus {
	castFunc := func(typ interface{}) BACnetLockStatus {
		if sBACnetLockStatus, ok := typ.(BACnetLockStatus); ok {
			return sBACnetLockStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLockStatus) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetLockStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLockStatusParse(readBuffer utils.ReadBuffer) (BACnetLockStatus, error) {
	val, err := readBuffer.ReadUint8("BACnetLockStatus", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetLockStatusByValue(val), nil
}

func (e BACnetLockStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetLockStatus", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetLockStatus) name() string {
	switch e {
	case BACnetLockStatus_LOCKED:
		return "LOCKED"
	case BACnetLockStatus_UNLOCKED:
		return "UNLOCKED"
	case BACnetLockStatus_LOCK_FAULT:
		return "LOCK_FAULT"
	case BACnetLockStatus_UNUSED:
		return "UNUSED"
	case BACnetLockStatus_UNKNOWN:
		return "UNKNOWN"
	}
	return ""
}

func (e BACnetLockStatus) String() string {
	return e.name()
}