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

// BACnetPropertyValues is the corresponding interface of BACnetPropertyValues
type BACnetPropertyValues interface {
	utils.LengthAware
	utils.Serializable
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetData returns Data (property field)
	GetData() []BACnetPropertyValue
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetPropertyValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyValues.
// This is useful for switch cases.
type BACnetPropertyValuesExactly interface {
	BACnetPropertyValues
	isBACnetPropertyValues() bool
}

// _BACnetPropertyValues is the data-structure of this message
type _BACnetPropertyValues struct {
	InnerOpeningTag BACnetOpeningTag
	Data            []BACnetPropertyValue
	InnerClosingTag BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyValues) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetPropertyValues) GetData() []BACnetPropertyValue {
	return m.Data
}

func (m *_BACnetPropertyValues) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyValues factory function for _BACnetPropertyValues
func NewBACnetPropertyValues(innerOpeningTag BACnetOpeningTag, data []BACnetPropertyValue, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetPropertyValues {
	return &_BACnetPropertyValues{InnerOpeningTag: innerOpeningTag, Data: data, InnerClosingTag: innerClosingTag, TagNumber: tagNumber, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyValues(structType interface{}) BACnetPropertyValues {
	if casted, ok := structType.(BACnetPropertyValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyValues) GetTypeName() string {
	return "BACnetPropertyValues"
}

func (m *_BACnetPropertyValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Array field
	if len(m.Data) > 0 {
		for _, element := range m.Data {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyValuesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetPropertyValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetPropertyValues")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Terminated array
	var data []BACnetPropertyValue
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetPropertyValueParse(readBuffer, objectTypeArgument)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field of BACnetPropertyValues")
			}
			data = append(data, _item.(BACnetPropertyValue))

		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetPropertyValues")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyValues")
	}

	// Create the instance
	return &_BACnetPropertyValues{
		TagNumber:          tagNumber,
		ObjectTypeArgument: objectTypeArgument,
		InnerOpeningTag:    innerOpeningTag,
		Data:               data,
		InnerClosingTag:    innerClosingTag,
	}, nil
}

func (m *_BACnetPropertyValues) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyValues")
	}

	// Simple Field (innerOpeningTag)
	if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
	}
	_innerOpeningTagErr := writeBuffer.WriteSerializable(m.GetInnerOpeningTag())
	if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for innerOpeningTag")
	}
	if _innerOpeningTagErr != nil {
		return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
	}

	// Array Field (data)
	if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for data")
	}
	for _, _element := range m.GetData() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'data' field")
		}
	}
	if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for data")
	}

	// Simple Field (innerClosingTag)
	if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
	}
	_innerClosingTagErr := writeBuffer.WriteSerializable(m.GetInnerClosingTag())
	if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for innerClosingTag")
	}
	if _innerClosingTagErr != nil {
		return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetPropertyValues) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetPropertyValues) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}

//
////

func (m *_BACnetPropertyValues) isBACnetPropertyValues() bool {
	return true
}

func (m *_BACnetPropertyValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
