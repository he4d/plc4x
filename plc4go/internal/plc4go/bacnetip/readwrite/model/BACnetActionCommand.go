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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetActionCommand struct {
	DeviceIdentifier   *BACnetContextTagObjectIdentifier
	ObjectIdentifier   *BACnetContextTagObjectIdentifier
	PropertyIdentifier *BACnetContextTagPropertyIdentifier
	ArrayIndex         *BACnetContextTagUnsignedInteger
	PropertyValue      *BACnetConstructedData
	Priority           *BACnetContextTagUnsignedInteger
	PostDelay          *BACnetContextTagBoolean
	QuitOnFailure      *BACnetContextTagBoolean
	WriteSuccessful    *BACnetContextTagBoolean
}

// The corresponding interface
type IBACnetActionCommand interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewBACnetActionCommand(deviceIdentifier *BACnetContextTagObjectIdentifier, objectIdentifier *BACnetContextTagObjectIdentifier, propertyIdentifier *BACnetContextTagPropertyIdentifier, arrayIndex *BACnetContextTagUnsignedInteger, propertyValue *BACnetConstructedData, priority *BACnetContextTagUnsignedInteger, postDelay *BACnetContextTagBoolean, quitOnFailure *BACnetContextTagBoolean, writeSuccessful *BACnetContextTagBoolean) *BACnetActionCommand {
	return &BACnetActionCommand{DeviceIdentifier: deviceIdentifier, ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, Priority: priority, PostDelay: postDelay, QuitOnFailure: quitOnFailure, WriteSuccessful: writeSuccessful}
}

func CastBACnetActionCommand(structType interface{}) *BACnetActionCommand {
	castFunc := func(typ interface{}) *BACnetActionCommand {
		if casted, ok := typ.(BACnetActionCommand); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetActionCommand); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetActionCommand) GetTypeName() string {
	return "BACnetActionCommand"
}

func (m *BACnetActionCommand) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetActionCommand) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += (*m.DeviceIdentifier).LengthInBits()
	}

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.LengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.LengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).LengthInBits()
	}

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += (*m.PropertyValue).LengthInBits()
	}

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += (*m.Priority).LengthInBits()
	}

	// Optional Field (postDelay)
	if m.PostDelay != nil {
		lengthInBits += (*m.PostDelay).LengthInBits()
	}

	// Simple field (quitOnFailure)
	lengthInBits += m.QuitOnFailure.LengthInBits()

	// Simple field (writeSuccessful)
	lengthInBits += m.WriteSuccessful.LengthInBits()

	return lengthInBits
}

func (m *BACnetActionCommand) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetActionCommandParse(readBuffer utils.ReadBuffer) (*BACnetActionCommand, error) {
	if pullErr := readBuffer.PullContext("BACnetActionCommand"); pullErr != nil {
		return nil, pullErr
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if a given expression evaluates to false)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field")
		default:
			deviceIdentifier = CastBACnetContextTagObjectIdentifier(_val)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType_BACNET_PROPERTY_IDENTIFIER)
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetContextTagPropertyIdentifier(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue *BACnetConstructedData = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(4), objectIdentifier.ObjectType, propertyIdentifier)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field")
		default:
			propertyValue = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority *BACnetContextTagUnsignedInteger = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(5), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field")
		default:
			priority = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (postDelay) (Can be skipped, if a given expression evaluates to false)
	var postDelay *BACnetContextTagBoolean = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("postDelay"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(6), BACnetDataType_BOOLEAN)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'postDelay' field")
		default:
			postDelay = CastBACnetContextTagBoolean(_val)
			if closeErr := readBuffer.CloseContext("postDelay"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (quitOnFailure)
	if pullErr := readBuffer.PullContext("quitOnFailure"); pullErr != nil {
		return nil, pullErr
	}
	_quitOnFailure, _quitOnFailureErr := BACnetContextTagParse(readBuffer, uint8(uint8(7)), BACnetDataType_BOOLEAN)
	if _quitOnFailureErr != nil {
		return nil, errors.Wrap(_quitOnFailureErr, "Error parsing 'quitOnFailure' field")
	}
	quitOnFailure := CastBACnetContextTagBoolean(_quitOnFailure)
	if closeErr := readBuffer.CloseContext("quitOnFailure"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (writeSuccessful)
	if pullErr := readBuffer.PullContext("writeSuccessful"); pullErr != nil {
		return nil, pullErr
	}
	_writeSuccessful, _writeSuccessfulErr := BACnetContextTagParse(readBuffer, uint8(uint8(8)), BACnetDataType_BOOLEAN)
	if _writeSuccessfulErr != nil {
		return nil, errors.Wrap(_writeSuccessfulErr, "Error parsing 'writeSuccessful' field")
	}
	writeSuccessful := CastBACnetContextTagBoolean(_writeSuccessful)
	if closeErr := readBuffer.CloseContext("writeSuccessful"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetActionCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetActionCommand(deviceIdentifier, objectIdentifier, propertyIdentifier, arrayIndex, propertyValue, priority, postDelay, quitOnFailure, writeSuccessful), nil
}

func (m *BACnetActionCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetActionCommand"); pushErr != nil {
		return pushErr
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
	var deviceIdentifier *BACnetContextTagObjectIdentifier = nil
	if m.DeviceIdentifier != nil {
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return pushErr
		}
		deviceIdentifier = m.DeviceIdentifier
		_deviceIdentifierErr := deviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return popErr
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return pushErr
	}
	_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return popErr
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return pushErr
	}
	_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return popErr
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	if m.ArrayIndex != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return pushErr
		}
		arrayIndex = m.ArrayIndex
		_arrayIndexErr := arrayIndex.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return popErr
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue *BACnetConstructedData = nil
	if m.PropertyValue != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return pushErr
		}
		propertyValue = m.PropertyValue
		_propertyValueErr := propertyValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return popErr
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (priority) (Can be skipped, if the value is null)
	var priority *BACnetContextTagUnsignedInteger = nil
	if m.Priority != nil {
		if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
			return pushErr
		}
		priority = m.Priority
		_priorityErr := priority.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("priority"); popErr != nil {
			return popErr
		}
		if _priorityErr != nil {
			return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
		}
	}

	// Optional Field (postDelay) (Can be skipped, if the value is null)
	var postDelay *BACnetContextTagBoolean = nil
	if m.PostDelay != nil {
		if pushErr := writeBuffer.PushContext("postDelay"); pushErr != nil {
			return pushErr
		}
		postDelay = m.PostDelay
		_postDelayErr := postDelay.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("postDelay"); popErr != nil {
			return popErr
		}
		if _postDelayErr != nil {
			return errors.Wrap(_postDelayErr, "Error serializing 'postDelay' field")
		}
	}

	// Simple Field (quitOnFailure)
	if pushErr := writeBuffer.PushContext("quitOnFailure"); pushErr != nil {
		return pushErr
	}
	_quitOnFailureErr := m.QuitOnFailure.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("quitOnFailure"); popErr != nil {
		return popErr
	}
	if _quitOnFailureErr != nil {
		return errors.Wrap(_quitOnFailureErr, "Error serializing 'quitOnFailure' field")
	}

	// Simple Field (writeSuccessful)
	if pushErr := writeBuffer.PushContext("writeSuccessful"); pushErr != nil {
		return pushErr
	}
	_writeSuccessfulErr := m.WriteSuccessful.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("writeSuccessful"); popErr != nil {
		return popErr
	}
	if _writeSuccessfulErr != nil {
		return errors.Wrap(_writeSuccessfulErr, "Error serializing 'writeSuccessful' field")
	}

	if popErr := writeBuffer.PopContext("BACnetActionCommand"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetActionCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
