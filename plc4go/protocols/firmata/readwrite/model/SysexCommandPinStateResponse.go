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

// SysexCommandPinStateResponse is the corresponding interface of SysexCommandPinStateResponse
type SysexCommandPinStateResponse interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetPinMode returns PinMode (property field)
	GetPinMode() uint8
	// GetPinState returns PinState (property field)
	GetPinState() uint8
}

// SysexCommandPinStateResponseExactly can be used when we want exactly this type and not a type which fulfills SysexCommandPinStateResponse.
// This is useful for switch cases.
type SysexCommandPinStateResponseExactly interface {
	SysexCommandPinStateResponse
	isSysexCommandPinStateResponse() bool
}

// _SysexCommandPinStateResponse is the data-structure of this message
type _SysexCommandPinStateResponse struct {
	*_SysexCommand
	Pin      uint8
	PinMode  uint8
	PinState uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandPinStateResponse) GetCommandType() uint8 {
	return 0x6E
}

func (m *_SysexCommandPinStateResponse) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandPinStateResponse) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandPinStateResponse) GetParent() SysexCommand {
	return m._SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandPinStateResponse) GetPin() uint8 {
	return m.Pin
}

func (m *_SysexCommandPinStateResponse) GetPinMode() uint8 {
	return m.PinMode
}

func (m *_SysexCommandPinStateResponse) GetPinState() uint8 {
	return m.PinState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSysexCommandPinStateResponse factory function for _SysexCommandPinStateResponse
func NewSysexCommandPinStateResponse(pin uint8, pinMode uint8, pinState uint8) *_SysexCommandPinStateResponse {
	_result := &_SysexCommandPinStateResponse{
		Pin:           pin,
		PinMode:       pinMode,
		PinState:      pinState,
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandPinStateResponse(structType interface{}) SysexCommandPinStateResponse {
	if casted, ok := structType.(SysexCommandPinStateResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandPinStateResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandPinStateResponse) GetTypeName() string {
	return "SysexCommandPinStateResponse"
}

func (m *_SysexCommandPinStateResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandPinStateResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	// Simple field (pinMode)
	lengthInBits += 8

	// Simple field (pinState)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SysexCommandPinStateResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandPinStateResponseParse(readBuffer utils.ReadBuffer, response bool) (SysexCommandPinStateResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandPinStateResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandPinStateResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
	_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field of SysexCommandPinStateResponse")
	}
	pin := _pin

	// Simple Field (pinMode)
	_pinMode, _pinModeErr := readBuffer.ReadUint8("pinMode", 8)
	if _pinModeErr != nil {
		return nil, errors.Wrap(_pinModeErr, "Error parsing 'pinMode' field of SysexCommandPinStateResponse")
	}
	pinMode := _pinMode

	// Simple Field (pinState)
	_pinState, _pinStateErr := readBuffer.ReadUint8("pinState", 8)
	if _pinStateErr != nil {
		return nil, errors.Wrap(_pinStateErr, "Error parsing 'pinState' field of SysexCommandPinStateResponse")
	}
	pinState := _pinState

	if closeErr := readBuffer.CloseContext("SysexCommandPinStateResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandPinStateResponse")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandPinStateResponse{
		_SysexCommand: &_SysexCommand{},
		Pin:           pin,
		PinMode:       pinMode,
		PinState:      pinState,
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandPinStateResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandPinStateResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandPinStateResponse")
		}

		// Simple Field (pin)
		pin := uint8(m.GetPin())
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		// Simple Field (pinMode)
		pinMode := uint8(m.GetPinMode())
		_pinModeErr := writeBuffer.WriteUint8("pinMode", 8, (pinMode))
		if _pinModeErr != nil {
			return errors.Wrap(_pinModeErr, "Error serializing 'pinMode' field")
		}

		// Simple Field (pinState)
		pinState := uint8(m.GetPinState())
		_pinStateErr := writeBuffer.WriteUint8("pinState", 8, (pinState))
		if _pinStateErr != nil {
			return errors.Wrap(_pinStateErr, "Error serializing 'pinState' field")
		}

		if popErr := writeBuffer.PopContext("SysexCommandPinStateResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandPinStateResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SysexCommandPinStateResponse) isSysexCommandPinStateResponse() bool {
	return true
}

func (m *_SysexCommandPinStateResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
