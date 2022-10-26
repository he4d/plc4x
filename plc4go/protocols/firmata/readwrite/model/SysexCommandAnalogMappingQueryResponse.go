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

// SysexCommandAnalogMappingQueryResponse is the corresponding interface of SysexCommandAnalogMappingQueryResponse
type SysexCommandAnalogMappingQueryResponse interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
}

// SysexCommandAnalogMappingQueryResponseExactly can be used when we want exactly this type and not a type which fulfills SysexCommandAnalogMappingQueryResponse.
// This is useful for switch cases.
type SysexCommandAnalogMappingQueryResponseExactly interface {
	SysexCommandAnalogMappingQueryResponse
	isSysexCommandAnalogMappingQueryResponse() bool
}

// _SysexCommandAnalogMappingQueryResponse is the data-structure of this message
type _SysexCommandAnalogMappingQueryResponse struct {
	*_SysexCommand
	Pin uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandAnalogMappingQueryResponse) GetCommandType() uint8 {
	return 0x69
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandAnalogMappingQueryResponse) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandAnalogMappingQueryResponse) GetParent() SysexCommand {
	return m._SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandAnalogMappingQueryResponse) GetPin() uint8 {
	return m.Pin
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandAnalogMappingQueryResponse factory function for _SysexCommandAnalogMappingQueryResponse
func NewSysexCommandAnalogMappingQueryResponse(pin uint8) *_SysexCommandAnalogMappingQueryResponse {
	_result := &_SysexCommandAnalogMappingQueryResponse{
		Pin:           pin,
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandAnalogMappingQueryResponse(structType interface{}) SysexCommandAnalogMappingQueryResponse {
	if casted, ok := structType.(SysexCommandAnalogMappingQueryResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandAnalogMappingQueryResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetTypeName() string {
	return "SysexCommandAnalogMappingQueryResponse"
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SysexCommandAnalogMappingQueryResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandAnalogMappingQueryResponseParse(readBuffer utils.ReadBuffer, response bool) (SysexCommandAnalogMappingQueryResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandAnalogMappingQueryResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandAnalogMappingQueryResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field of SysexCommandAnalogMappingQueryResponse")
	}
	pin := _pin

	if closeErr := readBuffer.CloseContext("SysexCommandAnalogMappingQueryResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandAnalogMappingQueryResponse")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandAnalogMappingQueryResponse{
		_SysexCommand: &_SysexCommand{},
		Pin:           pin,
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandAnalogMappingQueryResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandAnalogMappingQueryResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandAnalogMappingQueryResponse")
		}

		// Simple Field (pin)
		pin := uint8(m.GetPin())
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandAnalogMappingQueryResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandAnalogMappingQueryResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SysexCommandAnalogMappingQueryResponse) isSysexCommandAnalogMappingQueryResponse() bool {
	return true
}

func (m *_SysexCommandAnalogMappingQueryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
