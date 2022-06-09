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

// BACnetConstructedDataLastCredentialRemovedTime is the data-structure of this message
type BACnetConstructedDataLastCredentialRemovedTime struct {
	*BACnetConstructedData
	LastCredentialRemovedTime *BACnetDateTime

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataLastCredentialRemovedTime is the corresponding interface of BACnetConstructedDataLastCredentialRemovedTime
type IBACnetConstructedDataLastCredentialRemovedTime interface {
	IBACnetConstructedData
	// GetLastCredentialRemovedTime returns LastCredentialRemovedTime (property field)
	GetLastCredentialRemovedTime() *BACnetDateTime
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_CREDENTIAL_REMOVED_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLastCredentialRemovedTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetLastCredentialRemovedTime() *BACnetDateTime {
	return m.LastCredentialRemovedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastCredentialRemovedTime factory function for BACnetConstructedDataLastCredentialRemovedTime
func NewBACnetConstructedDataLastCredentialRemovedTime(lastCredentialRemovedTime *BACnetDateTime, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataLastCredentialRemovedTime {
	_result := &BACnetConstructedDataLastCredentialRemovedTime{
		LastCredentialRemovedTime: lastCredentialRemovedTime,
		BACnetConstructedData:     NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLastCredentialRemovedTime(structType interface{}) *BACnetConstructedDataLastCredentialRemovedTime {
	if casted, ok := structType.(BACnetConstructedDataLastCredentialRemovedTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialRemovedTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastCredentialRemovedTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastCredentialRemovedTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialRemovedTime"
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastCredentialRemovedTime)
	lengthInBits += m.LastCredentialRemovedTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastCredentialRemovedTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataLastCredentialRemovedTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialRemovedTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastCredentialRemovedTime)
	if pullErr := readBuffer.PullContext("lastCredentialRemovedTime"); pullErr != nil {
		return nil, pullErr
	}
	_lastCredentialRemovedTime, _lastCredentialRemovedTimeErr := BACnetDateTimeParse(readBuffer)
	if _lastCredentialRemovedTimeErr != nil {
		return nil, errors.Wrap(_lastCredentialRemovedTimeErr, "Error parsing 'lastCredentialRemovedTime' field")
	}
	lastCredentialRemovedTime := CastBACnetDateTime(_lastCredentialRemovedTime)
	if closeErr := readBuffer.CloseContext("lastCredentialRemovedTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialRemovedTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLastCredentialRemovedTime{
		LastCredentialRemovedTime: CastBACnetDateTime(lastCredentialRemovedTime),
		BACnetConstructedData:     &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialRemovedTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (lastCredentialRemovedTime)
		if pushErr := writeBuffer.PushContext("lastCredentialRemovedTime"); pushErr != nil {
			return pushErr
		}
		_lastCredentialRemovedTimeErr := m.LastCredentialRemovedTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lastCredentialRemovedTime"); popErr != nil {
			return popErr
		}
		if _lastCredentialRemovedTimeErr != nil {
			return errors.Wrap(_lastCredentialRemovedTimeErr, "Error serializing 'lastCredentialRemovedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialRemovedTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLastCredentialRemovedTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}