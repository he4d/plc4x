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

// BACnetConstructedDataBBMDBroadcastDistributionTable is the corresponding interface of BACnetConstructedDataBBMDBroadcastDistributionTable
type BACnetConstructedDataBBMDBroadcastDistributionTable interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBbmdBroadcastDistributionTable returns BbmdBroadcastDistributionTable (property field)
	GetBbmdBroadcastDistributionTable() []BACnetBDTEntry
}

// BACnetConstructedDataBBMDBroadcastDistributionTableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBBMDBroadcastDistributionTable.
// This is useful for switch cases.
type BACnetConstructedDataBBMDBroadcastDistributionTableExactly interface {
	BACnetConstructedDataBBMDBroadcastDistributionTable
	isBACnetConstructedDataBBMDBroadcastDistributionTable() bool
}

// _BACnetConstructedDataBBMDBroadcastDistributionTable is the data-structure of this message
type _BACnetConstructedDataBBMDBroadcastDistributionTable struct {
	*_BACnetConstructedData
	BbmdBroadcastDistributionTable []BACnetBDTEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BBMD_BROADCAST_DISTRIBUTION_TABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetBbmdBroadcastDistributionTable() []BACnetBDTEntry {
	return m.BbmdBroadcastDistributionTable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBBMDBroadcastDistributionTable factory function for _BACnetConstructedDataBBMDBroadcastDistributionTable
func NewBACnetConstructedDataBBMDBroadcastDistributionTable(bbmdBroadcastDistributionTable []BACnetBDTEntry, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBBMDBroadcastDistributionTable {
	_result := &_BACnetConstructedDataBBMDBroadcastDistributionTable{
		BbmdBroadcastDistributionTable: bbmdBroadcastDistributionTable,
		_BACnetConstructedData:         NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBBMDBroadcastDistributionTable(structType interface{}) BACnetConstructedDataBBMDBroadcastDistributionTable {
	if casted, ok := structType.(BACnetConstructedDataBBMDBroadcastDistributionTable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDBroadcastDistributionTable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetTypeName() string {
	return "BACnetConstructedDataBBMDBroadcastDistributionTable"
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.BbmdBroadcastDistributionTable) > 0 {
		for _, element := range m.BbmdBroadcastDistributionTable {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBBMDBroadcastDistributionTableParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBBMDBroadcastDistributionTable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDBroadcastDistributionTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (bbmdBroadcastDistributionTable)
	if pullErr := readBuffer.PullContext("bbmdBroadcastDistributionTable", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bbmdBroadcastDistributionTable")
	}
	// Terminated array
	var bbmdBroadcastDistributionTable []BACnetBDTEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetBDTEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'bbmdBroadcastDistributionTable' field of BACnetConstructedDataBBMDBroadcastDistributionTable")
			}
			bbmdBroadcastDistributionTable = append(bbmdBroadcastDistributionTable, _item.(BACnetBDTEntry))

		}
	}
	if closeErr := readBuffer.CloseContext("bbmdBroadcastDistributionTable", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bbmdBroadcastDistributionTable")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDBroadcastDistributionTable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBBMDBroadcastDistributionTable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BbmdBroadcastDistributionTable: bbmdBroadcastDistributionTable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDBroadcastDistributionTable")
		}

		// Array Field (bbmdBroadcastDistributionTable)
		if pushErr := writeBuffer.PushContext("bbmdBroadcastDistributionTable", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bbmdBroadcastDistributionTable")
		}
		for _, _element := range m.GetBbmdBroadcastDistributionTable() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'bbmdBroadcastDistributionTable' field")
			}
		}
		if popErr := writeBuffer.PopContext("bbmdBroadcastDistributionTable", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bbmdBroadcastDistributionTable")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDBroadcastDistributionTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDBroadcastDistributionTable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) isBACnetConstructedDataBBMDBroadcastDistributionTable() bool {
	return true
}

func (m *_BACnetConstructedDataBBMDBroadcastDistributionTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
