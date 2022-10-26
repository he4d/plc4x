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

// ApduDataMemoryResponse is the corresponding interface of ApduDataMemoryResponse
type ApduDataMemoryResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduData
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetData returns Data (property field)
	GetData() []byte
}

// ApduDataMemoryResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataMemoryResponse.
// This is useful for switch cases.
type ApduDataMemoryResponseExactly interface {
	ApduDataMemoryResponse
	isApduDataMemoryResponse() bool
}

// _ApduDataMemoryResponse is the data-structure of this message
type _ApduDataMemoryResponse struct {
	*_ApduData
	Address uint16
	Data    []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataMemoryResponse) GetApciType() uint8 {
	return 0x9
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataMemoryResponse) InitializeParent(parent ApduData) {}

func (m *_ApduDataMemoryResponse) GetParent() ApduData {
	return m._ApduData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataMemoryResponse) GetAddress() uint16 {
	return m.Address
}

func (m *_ApduDataMemoryResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataMemoryResponse factory function for _ApduDataMemoryResponse
func NewApduDataMemoryResponse(address uint16, data []byte, dataLength uint8) *_ApduDataMemoryResponse {
	_result := &_ApduDataMemoryResponse{
		Address:   address,
		Data:      data,
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataMemoryResponse(structType interface{}) ApduDataMemoryResponse {
	if casted, ok := structType.(ApduDataMemoryResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataMemoryResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataMemoryResponse) GetTypeName() string {
	return "ApduDataMemoryResponse"
}

func (m *_ApduDataMemoryResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataMemoryResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (numBytes)
	lengthInBits += 6

	// Simple field (address)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataMemoryResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataMemoryResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataMemoryResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataMemoryResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataMemoryResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (numBytes) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numBytes, _numBytesErr := readBuffer.ReadUint8("numBytes", 6)
	_ = numBytes
	if _numBytesErr != nil {
		return nil, errors.Wrap(_numBytesErr, "Error parsing 'numBytes' field of ApduDataMemoryResponse")
	}

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ApduDataMemoryResponse")
	}
	address := _address
	// Byte Array field (data)
	numberOfBytesdata := int(numBytes)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ApduDataMemoryResponse")
	}

	if closeErr := readBuffer.CloseContext("ApduDataMemoryResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataMemoryResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataMemoryResponse{
		_ApduData: &_ApduData{
			DataLength: dataLength,
		},
		Address: address,
		Data:    data,
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataMemoryResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataMemoryResponse")
		}

		// Implicit Field (numBytes) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numBytes := uint8(uint8(len(m.GetData())))
		_numBytesErr := writeBuffer.WriteUint8("numBytes", 6, (numBytes))
		if _numBytesErr != nil {
			return errors.Wrap(_numBytesErr, "Error serializing 'numBytes' field")
		}

		// Simple Field (address)
		address := uint16(m.GetAddress())
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataMemoryResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataMemoryResponse) isApduDataMemoryResponse() bool {
	return true
}

func (m *_ApduDataMemoryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
