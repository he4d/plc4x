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
type BACnetPropertyStatesAction struct {
	*BACnetPropertyStates
	Action *BACnetAction
}

// The corresponding interface
type IBACnetPropertyStatesAction interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////

func (m *BACnetPropertyStatesAction) InitializeParent(parent *BACnetPropertyStates, openingTag *BACnetOpeningTag, peekedTagNumber uint8, closingTag *BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagNumber = peekedTagNumber
	m.ClosingTag = closingTag
}

func NewBACnetPropertyStatesAction(action *BACnetAction, openingTag *BACnetOpeningTag, peekedTagNumber uint8, closingTag *BACnetClosingTag) *BACnetPropertyStates {
	child := &BACnetPropertyStatesAction{
		Action:               action,
		BACnetPropertyStates: NewBACnetPropertyStates(openingTag, peekedTagNumber, closingTag),
	}
	child.Child = child
	return child.BACnetPropertyStates
}

func CastBACnetPropertyStatesAction(structType interface{}) *BACnetPropertyStatesAction {
	castFunc := func(typ interface{}) *BACnetPropertyStatesAction {
		if casted, ok := typ.(BACnetPropertyStatesAction); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetPropertyStatesAction); ok {
			return casted
		}
		if casted, ok := typ.(BACnetPropertyStates); ok {
			return CastBACnetPropertyStatesAction(casted.Child)
		}
		if casted, ok := typ.(*BACnetPropertyStates); ok {
			return CastBACnetPropertyStatesAction(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetPropertyStatesAction) GetTypeName() string {
	return "BACnetPropertyStatesAction"
}

func (m *BACnetPropertyStatesAction) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetPropertyStatesAction) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Optional Field (action)
	if m.Action != nil {
		lengthInBits += (*m.Action).LengthInBits()
	}

	return lengthInBits
}

func (m *BACnetPropertyStatesAction) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetPropertyStatesActionParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (*BACnetPropertyStates, error) {
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAction"); pullErr != nil {
		return nil, pullErr
	}

	// Optional Field (action) (Can be skipped, if a given expression evaluates to false)
	var action *BACnetAction = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("action"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetActionParse(readBuffer, peekedTagNumber)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'action' field")
		case _err == utils.ParseAssertError:
			readBuffer.Reset(currentPos)
		default:
			action = CastBACnetAction(_val)
			if closeErr := readBuffer.CloseContext("action"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAction"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetPropertyStatesAction{
		Action:               CastBACnetAction(action),
		BACnetPropertyStates: &BACnetPropertyStates{},
	}
	_child.BACnetPropertyStates.Child = _child
	return _child.BACnetPropertyStates, nil
}

func (m *BACnetPropertyStatesAction) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAction"); pushErr != nil {
			return pushErr
		}

		// Optional Field (action) (Can be skipped, if the value is null)
		var action *BACnetAction = nil
		if m.Action != nil {
			if pushErr := writeBuffer.PushContext("action"); pushErr != nil {
				return pushErr
			}
			action = m.Action
			_actionErr := action.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("action"); popErr != nil {
				return popErr
			}
			if _actionErr != nil {
				return errors.Wrap(_actionErr, "Error serializing 'action' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAction"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetPropertyStatesAction) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}