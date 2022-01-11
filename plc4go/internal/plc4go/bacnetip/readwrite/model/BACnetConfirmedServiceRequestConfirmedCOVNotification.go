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
type BACnetConfirmedServiceRequestConfirmedCOVNotification struct {
	*BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier *BACnetContextTagUnsignedInteger
	MonitoredObjectIdentifier   *BACnetContextTagObjectIdentifier
	IssueConfirmed              *BACnetContextTagBoolean
	LifetimeInSeconds           *BACnetContextTagUnsignedInteger
	ListOfValues                *BACnetPropertyValues
}

// The corresponding interface
type IBACnetConfirmedServiceRequestConfirmedCOVNotification interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) ServiceChoice() uint8 {
	return 0x01
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func NewBACnetConfirmedServiceRequestConfirmedCOVNotification(subscriberProcessIdentifier *BACnetContextTagUnsignedInteger, monitoredObjectIdentifier *BACnetContextTagObjectIdentifier, issueConfirmed *BACnetContextTagBoolean, lifetimeInSeconds *BACnetContextTagUnsignedInteger, listOfValues *BACnetPropertyValues) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestConfirmedCOVNotification{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		MonitoredObjectIdentifier:     monitoredObjectIdentifier,
		IssueConfirmed:                issueConfirmed,
		LifetimeInSeconds:             lifetimeInSeconds,
		ListOfValues:                  listOfValues,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestConfirmedCOVNotification(structType interface{}) *BACnetConfirmedServiceRequestConfirmedCOVNotification {
	castFunc := func(typ interface{}) *BACnetConfirmedServiceRequestConfirmedCOVNotification {
		if casted, ok := typ.(BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
			return casted
		}
		if casted, ok := typ.(BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestConfirmedCOVNotification(casted.Child)
		}
		if casted, ok := typ.(*BACnetConfirmedServiceRequest); ok {
			return CastBACnetConfirmedServiceRequestConfirmedCOVNotification(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotification"
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.LengthInBits()

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.LengthInBits()

	// Simple field (issueConfirmed)
	lengthInBits += m.IssueConfirmed.LengthInBits()

	// Simple field (lifetimeInSeconds)
	lengthInBits += m.LifetimeInSeconds.LengthInBits()

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.LengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_UNSIGNED_INTEGER)
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}
	subscriberProcessIdentifier := CastBACnetContextTagUnsignedInteger(_subscriberProcessIdentifier)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field")
	}
	monitoredObjectIdentifier := CastBACnetContextTagObjectIdentifier(_monitoredObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (issueConfirmed)
	if pullErr := readBuffer.PullContext("issueConfirmed"); pullErr != nil {
		return nil, pullErr
	}
	_issueConfirmed, _issueConfirmedErr := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_BOOLEAN)
	if _issueConfirmedErr != nil {
		return nil, errors.Wrap(_issueConfirmedErr, "Error parsing 'issueConfirmed' field")
	}
	issueConfirmed := CastBACnetContextTagBoolean(_issueConfirmed)
	if closeErr := readBuffer.CloseContext("issueConfirmed"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lifetimeInSeconds)
	if pullErr := readBuffer.PullContext("lifetimeInSeconds"); pullErr != nil {
		return nil, pullErr
	}
	_lifetimeInSeconds, _lifetimeInSecondsErr := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_UNSIGNED_INTEGER)
	if _lifetimeInSecondsErr != nil {
		return nil, errors.Wrap(_lifetimeInSecondsErr, "Error parsing 'lifetimeInSeconds' field")
	}
	lifetimeInSeconds := CastBACnetContextTagUnsignedInteger(_lifetimeInSeconds)
	if closeErr := readBuffer.CloseContext("lifetimeInSeconds"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfValues, _listOfValuesErr := BACnetPropertyValuesParse(readBuffer, uint8(4))
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field")
	}
	listOfValues := CastBACnetPropertyValues(_listOfValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedCOVNotification{
		SubscriberProcessIdentifier:   CastBACnetContextTagUnsignedInteger(subscriberProcessIdentifier),
		MonitoredObjectIdentifier:     CastBACnetContextTagObjectIdentifier(monitoredObjectIdentifier),
		IssueConfirmed:                CastBACnetContextTagBoolean(issueConfirmed),
		LifetimeInSeconds:             CastBACnetContextTagUnsignedInteger(lifetimeInSeconds),
		ListOfValues:                  CastBACnetPropertyValues(listOfValues),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return pushErr
		}
		_subscriberProcessIdentifierErr := m.SubscriberProcessIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return popErr
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (monitoredObjectIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
			return pushErr
		}
		_monitoredObjectIdentifierErr := m.MonitoredObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
			return popErr
		}
		if _monitoredObjectIdentifierErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
		}

		// Simple Field (issueConfirmed)
		if pushErr := writeBuffer.PushContext("issueConfirmed"); pushErr != nil {
			return pushErr
		}
		_issueConfirmedErr := m.IssueConfirmed.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("issueConfirmed"); popErr != nil {
			return popErr
		}
		if _issueConfirmedErr != nil {
			return errors.Wrap(_issueConfirmedErr, "Error serializing 'issueConfirmed' field")
		}

		// Simple Field (lifetimeInSeconds)
		if pushErr := writeBuffer.PushContext("lifetimeInSeconds"); pushErr != nil {
			return pushErr
		}
		_lifetimeInSecondsErr := m.LifetimeInSeconds.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lifetimeInSeconds"); popErr != nil {
			return popErr
		}
		if _lifetimeInSecondsErr != nil {
			return errors.Wrap(_lifetimeInSecondsErr, "Error serializing 'lifetimeInSeconds' field")
		}

		// Simple Field (listOfValues)
		if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
			return pushErr
		}
		_listOfValuesErr := m.ListOfValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
			return popErr
		}
		if _listOfValuesErr != nil {
			return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
