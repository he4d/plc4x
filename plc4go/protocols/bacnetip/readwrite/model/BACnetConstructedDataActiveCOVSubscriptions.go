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

// BACnetConstructedDataActiveCOVSubscriptions is the data-structure of this message
type BACnetConstructedDataActiveCOVSubscriptions struct {
	*BACnetConstructedData
	ActiveCOVSubscriptions []*BACnetCOVSubscription

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataActiveCOVSubscriptions is the corresponding interface of BACnetConstructedDataActiveCOVSubscriptions
type IBACnetConstructedDataActiveCOVSubscriptions interface {
	IBACnetConstructedData
	// GetActiveCOVSubscriptions returns ActiveCOVSubscriptions (property field)
	GetActiveCOVSubscriptions() []*BACnetCOVSubscription
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

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_COV_SUBSCRIPTIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataActiveCOVSubscriptions) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetActiveCOVSubscriptions() []*BACnetCOVSubscription {
	return m.ActiveCOVSubscriptions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActiveCOVSubscriptions factory function for BACnetConstructedDataActiveCOVSubscriptions
func NewBACnetConstructedDataActiveCOVSubscriptions(activeCOVSubscriptions []*BACnetCOVSubscription, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataActiveCOVSubscriptions {
	_result := &BACnetConstructedDataActiveCOVSubscriptions{
		ActiveCOVSubscriptions: activeCOVSubscriptions,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataActiveCOVSubscriptions(structType interface{}) *BACnetConstructedDataActiveCOVSubscriptions {
	if casted, ok := structType.(BACnetConstructedDataActiveCOVSubscriptions); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveCOVSubscriptions); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataActiveCOVSubscriptions(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataActiveCOVSubscriptions(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetTypeName() string {
	return "BACnetConstructedDataActiveCOVSubscriptions"
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.ActiveCOVSubscriptions) > 0 {
		for _, element := range m.ActiveCOVSubscriptions {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataActiveCOVSubscriptionsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataActiveCOVSubscriptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveCOVSubscriptions"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (activeCOVSubscriptions)
	if pullErr := readBuffer.PullContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	activeCOVSubscriptions := make([]*BACnetCOVSubscription, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetCOVSubscriptionParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'activeCOVSubscriptions' field")
			}
			activeCOVSubscriptions = append(activeCOVSubscriptions, CastBACnetCOVSubscription(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveCOVSubscriptions"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataActiveCOVSubscriptions{
		ActiveCOVSubscriptions: activeCOVSubscriptions,
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveCOVSubscriptions"); pushErr != nil {
			return pushErr
		}

		// Array Field (activeCOVSubscriptions)
		if m.ActiveCOVSubscriptions != nil {
			if pushErr := writeBuffer.PushContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.ActiveCOVSubscriptions {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'activeCOVSubscriptions' field")
				}
			}
			if popErr := writeBuffer.PopContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveCOVSubscriptions"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataActiveCOVSubscriptions) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}