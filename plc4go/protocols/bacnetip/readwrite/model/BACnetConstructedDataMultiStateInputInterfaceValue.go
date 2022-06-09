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

// BACnetConstructedDataMultiStateInputInterfaceValue is the data-structure of this message
type BACnetConstructedDataMultiStateInputInterfaceValue struct {
	*BACnetConstructedData
	InterfaceValue *BACnetOptionalBinaryPV

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataMultiStateInputInterfaceValue is the corresponding interface of BACnetConstructedDataMultiStateInputInterfaceValue
type IBACnetConstructedDataMultiStateInputInterfaceValue interface {
	IBACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() *BACnetOptionalBinaryPV
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

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_INPUT
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetInterfaceValue() *BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMultiStateInputInterfaceValue factory function for BACnetConstructedDataMultiStateInputInterfaceValue
func NewBACnetConstructedDataMultiStateInputInterfaceValue(interfaceValue *BACnetOptionalBinaryPV, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataMultiStateInputInterfaceValue {
	_result := &BACnetConstructedDataMultiStateInputInterfaceValue{
		InterfaceValue:        interfaceValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMultiStateInputInterfaceValue(structType interface{}) *BACnetConstructedDataMultiStateInputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataMultiStateInputInterfaceValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateInputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMultiStateInputInterfaceValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMultiStateInputInterfaceValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataMultiStateInputInterfaceValue"
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMultiStateInputInterfaceValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataMultiStateInputInterfaceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateInputInterfaceValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (interfaceValue)
	if pullErr := readBuffer.PullContext("interfaceValue"); pullErr != nil {
		return nil, pullErr
	}
	_interfaceValue, _interfaceValueErr := BACnetOptionalBinaryPVParse(readBuffer)
	if _interfaceValueErr != nil {
		return nil, errors.Wrap(_interfaceValueErr, "Error parsing 'interfaceValue' field")
	}
	interfaceValue := CastBACnetOptionalBinaryPV(_interfaceValue)
	if closeErr := readBuffer.CloseContext("interfaceValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateInputInterfaceValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMultiStateInputInterfaceValue{
		InterfaceValue:        CastBACnetOptionalBinaryPV(interfaceValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateInputInterfaceValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (interfaceValue)
		if pushErr := writeBuffer.PushContext("interfaceValue"); pushErr != nil {
			return pushErr
		}
		_interfaceValueErr := m.InterfaceValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("interfaceValue"); popErr != nil {
			return popErr
		}
		if _interfaceValueErr != nil {
			return errors.Wrap(_interfaceValueErr, "Error serializing 'interfaceValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateInputInterfaceValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMultiStateInputInterfaceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}