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
type ApduDataOther struct {
	*ApduData
	ExtendedApdu *ApduDataExt
}

// The corresponding interface
type IApduDataOther interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataOther) ApciType() uint8 {
	return 0xF
}

func (m *ApduDataOther) InitializeParent(parent *ApduData) {}

func NewApduDataOther(extendedApdu *ApduDataExt) *ApduData {
	child := &ApduDataOther{
		ExtendedApdu: extendedApdu,
		ApduData:     NewApduData(),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataOther(structType interface{}) *ApduDataOther {
	castFunc := func(typ interface{}) *ApduDataOther {
		if casted, ok := typ.(ApduDataOther); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataOther); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataOther(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataOther(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataOther) GetTypeName() string {
	return "ApduDataOther"
}

func (m *ApduDataOther) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataOther) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (extendedApdu)
	lengthInBits += m.ExtendedApdu.LengthInBits()

	return lengthInBits
}

func (m *ApduDataOther) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataOtherParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataOther"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (extendedApdu)
	if pullErr := readBuffer.PullContext("extendedApdu"); pullErr != nil {
		return nil, pullErr
	}
	_extendedApdu, _extendedApduErr := ApduDataExtParse(readBuffer, uint8(dataLength))
	if _extendedApduErr != nil {
		return nil, errors.Wrap(_extendedApduErr, "Error parsing 'extendedApdu' field")
	}
	extendedApdu := CastApduDataExt(_extendedApdu)
	if closeErr := readBuffer.CloseContext("extendedApdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataOther"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataOther{
		ExtendedApdu: CastApduDataExt(extendedApdu),
		ApduData:     &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataOther) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataOther"); pushErr != nil {
			return pushErr
		}

		// Simple Field (extendedApdu)
		if pushErr := writeBuffer.PushContext("extendedApdu"); pushErr != nil {
			return pushErr
		}
		_extendedApduErr := m.ExtendedApdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("extendedApdu"); popErr != nil {
			return popErr
		}
		if _extendedApduErr != nil {
			return errors.Wrap(_extendedApduErr, "Error serializing 'extendedApdu' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataOther"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataOther) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
