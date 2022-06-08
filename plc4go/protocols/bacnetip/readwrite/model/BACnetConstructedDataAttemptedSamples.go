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

// BACnetConstructedDataAttemptedSamples is the data-structure of this message
type BACnetConstructedDataAttemptedSamples struct {
	*BACnetConstructedData
	AttemptedSamples *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAttemptedSamples is the corresponding interface of BACnetConstructedDataAttemptedSamples
type IBACnetConstructedDataAttemptedSamples interface {
	IBACnetConstructedData
	// GetAttemptedSamples returns AttemptedSamples (property field)
	GetAttemptedSamples() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataAttemptedSamples) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAttemptedSamples) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ATTEMPTED_SAMPLES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAttemptedSamples) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAttemptedSamples) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAttemptedSamples) GetAttemptedSamples() *BACnetApplicationTagUnsignedInteger {
	return m.AttemptedSamples
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAttemptedSamples factory function for BACnetConstructedDataAttemptedSamples
func NewBACnetConstructedDataAttemptedSamples(attemptedSamples *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAttemptedSamples {
	_result := &BACnetConstructedDataAttemptedSamples{
		AttemptedSamples:      attemptedSamples,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAttemptedSamples(structType interface{}) *BACnetConstructedDataAttemptedSamples {
	if casted, ok := structType.(BACnetConstructedDataAttemptedSamples); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAttemptedSamples); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAttemptedSamples(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAttemptedSamples(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAttemptedSamples) GetTypeName() string {
	return "BACnetConstructedDataAttemptedSamples"
}

func (m *BACnetConstructedDataAttemptedSamples) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAttemptedSamples) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (attemptedSamples)
	lengthInBits += m.AttemptedSamples.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAttemptedSamples) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAttemptedSamplesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAttemptedSamples, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAttemptedSamples"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (attemptedSamples)
	if pullErr := readBuffer.PullContext("attemptedSamples"); pullErr != nil {
		return nil, pullErr
	}
	_attemptedSamples, _attemptedSamplesErr := BACnetApplicationTagParse(readBuffer)
	if _attemptedSamplesErr != nil {
		return nil, errors.Wrap(_attemptedSamplesErr, "Error parsing 'attemptedSamples' field")
	}
	attemptedSamples := CastBACnetApplicationTagUnsignedInteger(_attemptedSamples)
	if closeErr := readBuffer.CloseContext("attemptedSamples"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAttemptedSamples"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAttemptedSamples{
		AttemptedSamples:      CastBACnetApplicationTagUnsignedInteger(attemptedSamples),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAttemptedSamples) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAttemptedSamples"); pushErr != nil {
			return pushErr
		}

		// Simple Field (attemptedSamples)
		if pushErr := writeBuffer.PushContext("attemptedSamples"); pushErr != nil {
			return pushErr
		}
		_attemptedSamplesErr := m.AttemptedSamples.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("attemptedSamples"); popErr != nil {
			return popErr
		}
		if _attemptedSamplesErr != nil {
			return errors.Wrap(_attemptedSamplesErr, "Error serializing 'attemptedSamples' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAttemptedSamples"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAttemptedSamples) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
