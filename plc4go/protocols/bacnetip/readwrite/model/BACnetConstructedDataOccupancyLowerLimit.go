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

// BACnetConstructedDataOccupancyLowerLimit is the data-structure of this message
type BACnetConstructedDataOccupancyLowerLimit struct {
	*BACnetConstructedData
	OccupancyLowerLimit *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataOccupancyLowerLimit is the corresponding interface of BACnetConstructedDataOccupancyLowerLimit
type IBACnetConstructedDataOccupancyLowerLimit interface {
	IBACnetConstructedData
	// GetOccupancyLowerLimit returns OccupancyLowerLimit (property field)
	GetOccupancyLowerLimit() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataOccupancyLowerLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataOccupancyLowerLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_LOWER_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataOccupancyLowerLimit) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataOccupancyLowerLimit) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataOccupancyLowerLimit) GetOccupancyLowerLimit() *BACnetApplicationTagUnsignedInteger {
	return m.OccupancyLowerLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyLowerLimit factory function for BACnetConstructedDataOccupancyLowerLimit
func NewBACnetConstructedDataOccupancyLowerLimit(occupancyLowerLimit *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataOccupancyLowerLimit {
	_result := &BACnetConstructedDataOccupancyLowerLimit{
		OccupancyLowerLimit:   occupancyLowerLimit,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataOccupancyLowerLimit(structType interface{}) *BACnetConstructedDataOccupancyLowerLimit {
	if casted, ok := structType.(BACnetConstructedDataOccupancyLowerLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyLowerLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyLowerLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyLowerLimit(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataOccupancyLowerLimit) GetTypeName() string {
	return "BACnetConstructedDataOccupancyLowerLimit"
}

func (m *BACnetConstructedDataOccupancyLowerLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataOccupancyLowerLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (occupancyLowerLimit)
	lengthInBits += m.OccupancyLowerLimit.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataOccupancyLowerLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOccupancyLowerLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataOccupancyLowerLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyLowerLimit"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyLowerLimit)
	if pullErr := readBuffer.PullContext("occupancyLowerLimit"); pullErr != nil {
		return nil, pullErr
	}
	_occupancyLowerLimit, _occupancyLowerLimitErr := BACnetApplicationTagParse(readBuffer)
	if _occupancyLowerLimitErr != nil {
		return nil, errors.Wrap(_occupancyLowerLimitErr, "Error parsing 'occupancyLowerLimit' field")
	}
	occupancyLowerLimit := CastBACnetApplicationTagUnsignedInteger(_occupancyLowerLimit)
	if closeErr := readBuffer.CloseContext("occupancyLowerLimit"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyLowerLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataOccupancyLowerLimit{
		OccupancyLowerLimit:   CastBACnetApplicationTagUnsignedInteger(occupancyLowerLimit),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataOccupancyLowerLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyLowerLimit"); pushErr != nil {
			return pushErr
		}

		// Simple Field (occupancyLowerLimit)
		if pushErr := writeBuffer.PushContext("occupancyLowerLimit"); pushErr != nil {
			return pushErr
		}
		_occupancyLowerLimitErr := m.OccupancyLowerLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("occupancyLowerLimit"); popErr != nil {
			return popErr
		}
		if _occupancyLowerLimitErr != nil {
			return errors.Wrap(_occupancyLowerLimitErr, "Error serializing 'occupancyLowerLimit' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyLowerLimit"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataOccupancyLowerLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}