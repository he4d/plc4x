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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetTimeStampDateTime struct {
	*BACnetTimeStamp
	DateTimeValue *BACnetDateTime
}

// The corresponding interface
type IBACnetTimeStampDateTime interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTimeStampDateTime) PeekedTagNumber() uint8 {
	return uint8(2)
}

func (m *BACnetTimeStampDateTime) InitializeParent(parent *BACnetTimeStamp, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) {
	m.BACnetTimeStamp.OpeningTag = openingTag
	m.BACnetTimeStamp.PeekedTagHeader = peekedTagHeader
	m.BACnetTimeStamp.ClosingTag = closingTag
	m.BACnetTimeStamp.PeekedTagNumber = peekedTagNumber
}

func NewBACnetTimeStampDateTime(dateTimeValue *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) *BACnetTimeStamp {
	child := &BACnetTimeStampDateTime{
		DateTimeValue:   dateTimeValue,
		BACnetTimeStamp: NewBACnetTimeStamp(openingTag, peekedTagHeader, closingTag, peekedTagNumber),
	}
	child.Child = child
	return child.BACnetTimeStamp
}

func CastBACnetTimeStampDateTime(structType interface{}) *BACnetTimeStampDateTime {
	castFunc := func(typ interface{}) *BACnetTimeStampDateTime {
		if casted, ok := typ.(BACnetTimeStampDateTime); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTimeStampDateTime); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTimeStamp); ok {
			return CastBACnetTimeStampDateTime(casted.Child)
		}
		if casted, ok := typ.(*BACnetTimeStamp); ok {
			return CastBACnetTimeStampDateTime(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTimeStampDateTime) GetTypeName() string {
	return "BACnetTimeStampDateTime"
}

func (m *BACnetTimeStampDateTime) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTimeStampDateTime) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.LengthInBits()

	return lengthInBits
}

func (m *BACnetTimeStampDateTime) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTimeStampDateTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetTimeStamp, error) {
	if pullErr := readBuffer.PullContext("BACnetTimeStampDateTime"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, pullErr
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeParse(readBuffer, uint8(uint8(2)))
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field")
	}
	dateTimeValue := CastBACnetDateTime(_dateTimeValue)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStampDateTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTimeStampDateTime{
		DateTimeValue:   CastBACnetDateTime(dateTimeValue),
		BACnetTimeStamp: &BACnetTimeStamp{},
	}
	_child.BACnetTimeStamp.Child = _child
	return _child.BACnetTimeStamp, nil
}

func (m *BACnetTimeStampDateTime) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimeStampDateTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dateTimeValue)
		if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
			return pushErr
		}
		_dateTimeValueErr := m.DateTimeValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
			return popErr
		}
		if _dateTimeValueErr != nil {
			return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimeStampDateTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTimeStampDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
