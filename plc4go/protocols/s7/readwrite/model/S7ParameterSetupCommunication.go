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

// S7ParameterSetupCommunication is the corresponding interface of S7ParameterSetupCommunication
type S7ParameterSetupCommunication interface {
	utils.LengthAware
	utils.Serializable
	S7Parameter
	// GetMaxAmqCaller returns MaxAmqCaller (property field)
	GetMaxAmqCaller() uint16
	// GetMaxAmqCallee returns MaxAmqCallee (property field)
	GetMaxAmqCallee() uint16
	// GetPduLength returns PduLength (property field)
	GetPduLength() uint16
}

// S7ParameterSetupCommunicationExactly can be used when we want exactly this type and not a type which fulfills S7ParameterSetupCommunication.
// This is useful for switch cases.
type S7ParameterSetupCommunicationExactly interface {
	S7ParameterSetupCommunication
	isS7ParameterSetupCommunication() bool
}

// _S7ParameterSetupCommunication is the data-structure of this message
type _S7ParameterSetupCommunication struct {
	*_S7Parameter
	MaxAmqCaller uint16
	MaxAmqCallee uint16
	PduLength    uint16
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterSetupCommunication) GetParameterType() uint8 {
	return 0xF0
}

func (m *_S7ParameterSetupCommunication) GetMessageType() uint8 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterSetupCommunication) InitializeParent(parent S7Parameter) {}

func (m *_S7ParameterSetupCommunication) GetParent() S7Parameter {
	return m._S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterSetupCommunication) GetMaxAmqCaller() uint16 {
	return m.MaxAmqCaller
}

func (m *_S7ParameterSetupCommunication) GetMaxAmqCallee() uint16 {
	return m.MaxAmqCallee
}

func (m *_S7ParameterSetupCommunication) GetPduLength() uint16 {
	return m.PduLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterSetupCommunication factory function for _S7ParameterSetupCommunication
func NewS7ParameterSetupCommunication(maxAmqCaller uint16, maxAmqCallee uint16, pduLength uint16) *_S7ParameterSetupCommunication {
	_result := &_S7ParameterSetupCommunication{
		MaxAmqCaller: maxAmqCaller,
		MaxAmqCallee: maxAmqCallee,
		PduLength:    pduLength,
		_S7Parameter: NewS7Parameter(),
	}
	_result._S7Parameter._S7ParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7ParameterSetupCommunication(structType interface{}) S7ParameterSetupCommunication {
	if casted, ok := structType.(S7ParameterSetupCommunication); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterSetupCommunication); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterSetupCommunication) GetTypeName() string {
	return "S7ParameterSetupCommunication"
}

func (m *_S7ParameterSetupCommunication) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7ParameterSetupCommunication) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (maxAmqCaller)
	lengthInBits += 16

	// Simple field (maxAmqCallee)
	lengthInBits += 16

	// Simple field (pduLength)
	lengthInBits += 16

	return lengthInBits
}

func (m *_S7ParameterSetupCommunication) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterSetupCommunicationParse(readBuffer utils.ReadBuffer, messageType uint8) (S7ParameterSetupCommunication, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterSetupCommunication"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterSetupCommunication")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of S7ParameterSetupCommunication")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (maxAmqCaller)
	_maxAmqCaller, _maxAmqCallerErr := readBuffer.ReadUint16("maxAmqCaller", 16)
	if _maxAmqCallerErr != nil {
		return nil, errors.Wrap(_maxAmqCallerErr, "Error parsing 'maxAmqCaller' field of S7ParameterSetupCommunication")
	}
	maxAmqCaller := _maxAmqCaller

	// Simple Field (maxAmqCallee)
	_maxAmqCallee, _maxAmqCalleeErr := readBuffer.ReadUint16("maxAmqCallee", 16)
	if _maxAmqCalleeErr != nil {
		return nil, errors.Wrap(_maxAmqCalleeErr, "Error parsing 'maxAmqCallee' field of S7ParameterSetupCommunication")
	}
	maxAmqCallee := _maxAmqCallee

	// Simple Field (pduLength)
	_pduLength, _pduLengthErr := readBuffer.ReadUint16("pduLength", 16)
	if _pduLengthErr != nil {
		return nil, errors.Wrap(_pduLengthErr, "Error parsing 'pduLength' field of S7ParameterSetupCommunication")
	}
	pduLength := _pduLength

	if closeErr := readBuffer.CloseContext("S7ParameterSetupCommunication"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterSetupCommunication")
	}

	// Create a partially initialized instance
	_child := &_S7ParameterSetupCommunication{
		_S7Parameter:   &_S7Parameter{},
		MaxAmqCaller:   maxAmqCaller,
		MaxAmqCallee:   maxAmqCallee,
		PduLength:      pduLength,
		reservedField0: reservedField0,
	}
	_child._S7Parameter._S7ParameterChildRequirements = _child
	return _child, nil
}

func (m *_S7ParameterSetupCommunication) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterSetupCommunication"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterSetupCommunication")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (maxAmqCaller)
		maxAmqCaller := uint16(m.GetMaxAmqCaller())
		_maxAmqCallerErr := writeBuffer.WriteUint16("maxAmqCaller", 16, (maxAmqCaller))
		if _maxAmqCallerErr != nil {
			return errors.Wrap(_maxAmqCallerErr, "Error serializing 'maxAmqCaller' field")
		}

		// Simple Field (maxAmqCallee)
		maxAmqCallee := uint16(m.GetMaxAmqCallee())
		_maxAmqCalleeErr := writeBuffer.WriteUint16("maxAmqCallee", 16, (maxAmqCallee))
		if _maxAmqCalleeErr != nil {
			return errors.Wrap(_maxAmqCalleeErr, "Error serializing 'maxAmqCallee' field")
		}

		// Simple Field (pduLength)
		pduLength := uint16(m.GetPduLength())
		_pduLengthErr := writeBuffer.WriteUint16("pduLength", 16, (pduLength))
		if _pduLengthErr != nil {
			return errors.Wrap(_pduLengthErr, "Error serializing 'pduLength' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterSetupCommunication"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterSetupCommunication")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7ParameterSetupCommunication) isS7ParameterSetupCommunication() bool {
	return true
}

func (m *_S7ParameterSetupCommunication) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
