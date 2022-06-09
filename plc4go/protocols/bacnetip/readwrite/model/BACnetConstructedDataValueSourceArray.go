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

// BACnetConstructedDataValueSourceArray is the data-structure of this message
type BACnetConstructedDataValueSourceArray struct {
	*BACnetConstructedData
	VtClassesSupported []*BACnetValueSource

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataValueSourceArray is the corresponding interface of BACnetConstructedDataValueSourceArray
type IBACnetConstructedDataValueSourceArray interface {
	IBACnetConstructedData
	// GetVtClassesSupported returns VtClassesSupported (property field)
	GetVtClassesSupported() []*BACnetValueSource
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

func (m *BACnetConstructedDataValueSourceArray) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataValueSourceArray) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_SOURCE_ARRAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataValueSourceArray) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataValueSourceArray) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataValueSourceArray) GetVtClassesSupported() []*BACnetValueSource {
	return m.VtClassesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataValueSourceArray factory function for BACnetConstructedDataValueSourceArray
func NewBACnetConstructedDataValueSourceArray(vtClassesSupported []*BACnetValueSource, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataValueSourceArray {
	_result := &BACnetConstructedDataValueSourceArray{
		VtClassesSupported:    vtClassesSupported,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataValueSourceArray(structType interface{}) *BACnetConstructedDataValueSourceArray {
	if casted, ok := structType.(BACnetConstructedDataValueSourceArray); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueSourceArray); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataValueSourceArray(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataValueSourceArray(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataValueSourceArray) GetTypeName() string {
	return "BACnetConstructedDataValueSourceArray"
}

func (m *BACnetConstructedDataValueSourceArray) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataValueSourceArray) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.VtClassesSupported) > 0 {
		for _, element := range m.VtClassesSupported {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataValueSourceArray) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataValueSourceArrayParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataValueSourceArray, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueSourceArray"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (vtClassesSupported)
	if pullErr := readBuffer.PullContext("vtClassesSupported", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	vtClassesSupported := make([]*BACnetValueSource, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetValueSourceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'vtClassesSupported' field")
			}
			vtClassesSupported = append(vtClassesSupported, CastBACnetValueSource(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("vtClassesSupported", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueSourceArray"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataValueSourceArray{
		VtClassesSupported:    vtClassesSupported,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataValueSourceArray) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueSourceArray"); pushErr != nil {
			return pushErr
		}

		// Array Field (vtClassesSupported)
		if m.VtClassesSupported != nil {
			if pushErr := writeBuffer.PushContext("vtClassesSupported", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.VtClassesSupported {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'vtClassesSupported' field")
				}
			}
			if popErr := writeBuffer.PopContext("vtClassesSupported", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueSourceArray"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataValueSourceArray) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}