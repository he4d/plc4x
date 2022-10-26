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

// BACnetConstructedDataDoorPulseTime is the corresponding interface of BACnetConstructedDataDoorPulseTime
type BACnetConstructedDataDoorPulseTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoorPulseTime returns DoorPulseTime (property field)
	GetDoorPulseTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDoorPulseTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDoorPulseTime.
// This is useful for switch cases.
type BACnetConstructedDataDoorPulseTimeExactly interface {
	BACnetConstructedDataDoorPulseTime
	isBACnetConstructedDataDoorPulseTime() bool
}

// _BACnetConstructedDataDoorPulseTime is the data-structure of this message
type _BACnetConstructedDataDoorPulseTime struct {
	*_BACnetConstructedData
	DoorPulseTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorPulseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorPulseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_PULSE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorPulseTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDoorPulseTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorPulseTime) GetDoorPulseTime() BACnetApplicationTagUnsignedInteger {
	return m.DoorPulseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorPulseTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetDoorPulseTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorPulseTime factory function for _BACnetConstructedDataDoorPulseTime
func NewBACnetConstructedDataDoorPulseTime(doorPulseTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorPulseTime {
	_result := &_BACnetConstructedDataDoorPulseTime{
		DoorPulseTime:          doorPulseTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorPulseTime(structType interface{}) BACnetConstructedDataDoorPulseTime {
	if casted, ok := structType.(BACnetConstructedDataDoorPulseTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorPulseTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorPulseTime) GetTypeName() string {
	return "BACnetConstructedDataDoorPulseTime"
}

func (m *_BACnetConstructedDataDoorPulseTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDoorPulseTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (doorPulseTime)
	lengthInBits += m.DoorPulseTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorPulseTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDoorPulseTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorPulseTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorPulseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorPulseTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorPulseTime)
	if pullErr := readBuffer.PullContext("doorPulseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorPulseTime")
	}
	_doorPulseTime, _doorPulseTimeErr := BACnetApplicationTagParse(readBuffer)
	if _doorPulseTimeErr != nil {
		return nil, errors.Wrap(_doorPulseTimeErr, "Error parsing 'doorPulseTime' field of BACnetConstructedDataDoorPulseTime")
	}
	doorPulseTime := _doorPulseTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("doorPulseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorPulseTime")
	}

	// Virtual field
	_actualValue := doorPulseTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorPulseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorPulseTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDoorPulseTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DoorPulseTime: doorPulseTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDoorPulseTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorPulseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorPulseTime")
		}

		// Simple Field (doorPulseTime)
		if pushErr := writeBuffer.PushContext("doorPulseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorPulseTime")
		}
		_doorPulseTimeErr := writeBuffer.WriteSerializable(m.GetDoorPulseTime())
		if popErr := writeBuffer.PopContext("doorPulseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorPulseTime")
		}
		if _doorPulseTimeErr != nil {
			return errors.Wrap(_doorPulseTimeErr, "Error serializing 'doorPulseTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorPulseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorPulseTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorPulseTime) isBACnetConstructedDataDoorPulseTime() bool {
	return true
}

func (m *_BACnetConstructedDataDoorPulseTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
