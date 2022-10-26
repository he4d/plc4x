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

// ReplyOrConfirmationReply is the corresponding interface of ReplyOrConfirmationReply
type ReplyOrConfirmationReply interface {
	utils.LengthAware
	utils.Serializable
	ReplyOrConfirmation
	// GetReply returns Reply (property field)
	GetReply() Reply
	// GetTermination returns Termination (property field)
	GetTermination() ResponseTermination
}

// ReplyOrConfirmationReplyExactly can be used when we want exactly this type and not a type which fulfills ReplyOrConfirmationReply.
// This is useful for switch cases.
type ReplyOrConfirmationReplyExactly interface {
	ReplyOrConfirmationReply
	isReplyOrConfirmationReply() bool
}

// _ReplyOrConfirmationReply is the data-structure of this message
type _ReplyOrConfirmationReply struct {
	*_ReplyOrConfirmation
	Reply       Reply
	Termination ResponseTermination
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReplyOrConfirmationReply) InitializeParent(parent ReplyOrConfirmation, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ReplyOrConfirmationReply) GetParent() ReplyOrConfirmation {
	return m._ReplyOrConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReplyOrConfirmationReply) GetReply() Reply {
	return m.Reply
}

func (m *_ReplyOrConfirmationReply) GetTermination() ResponseTermination {
	return m.Termination
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReplyOrConfirmationReply factory function for _ReplyOrConfirmationReply
func NewReplyOrConfirmationReply(reply Reply, termination ResponseTermination, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ReplyOrConfirmationReply {
	_result := &_ReplyOrConfirmationReply{
		Reply:                reply,
		Termination:          termination,
		_ReplyOrConfirmation: NewReplyOrConfirmation(peekedByte, cBusOptions, requestContext),
	}
	_result._ReplyOrConfirmation._ReplyOrConfirmationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReplyOrConfirmationReply(structType interface{}) ReplyOrConfirmationReply {
	if casted, ok := structType.(ReplyOrConfirmationReply); ok {
		return casted
	}
	if casted, ok := structType.(*ReplyOrConfirmationReply); ok {
		return *casted
	}
	return nil
}

func (m *_ReplyOrConfirmationReply) GetTypeName() string {
	return "ReplyOrConfirmationReply"
}

func (m *_ReplyOrConfirmationReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ReplyOrConfirmationReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reply)
	lengthInBits += m.Reply.GetLengthInBits()

	// Simple field (termination)
	lengthInBits += m.Termination.GetLengthInBits()

	return lengthInBits
}

func (m *_ReplyOrConfirmationReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyOrConfirmationReplyParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (ReplyOrConfirmationReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReplyOrConfirmationReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReplyOrConfirmationReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reply)
	if pullErr := readBuffer.PullContext("reply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reply")
	}
	_reply, _replyErr := ReplyParse(readBuffer, cBusOptions, requestContext)
	if _replyErr != nil {
		return nil, errors.Wrap(_replyErr, "Error parsing 'reply' field of ReplyOrConfirmationReply")
	}
	reply := _reply.(Reply)
	if closeErr := readBuffer.CloseContext("reply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reply")
	}

	// Simple Field (termination)
	if pullErr := readBuffer.PullContext("termination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for termination")
	}
	_termination, _terminationErr := ResponseTerminationParse(readBuffer)
	if _terminationErr != nil {
		return nil, errors.Wrap(_terminationErr, "Error parsing 'termination' field of ReplyOrConfirmationReply")
	}
	termination := _termination.(ResponseTermination)
	if closeErr := readBuffer.CloseContext("termination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for termination")
	}

	if closeErr := readBuffer.CloseContext("ReplyOrConfirmationReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReplyOrConfirmationReply")
	}

	// Create a partially initialized instance
	_child := &_ReplyOrConfirmationReply{
		_ReplyOrConfirmation: &_ReplyOrConfirmation{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
		Reply:       reply,
		Termination: termination,
	}
	_child._ReplyOrConfirmation._ReplyOrConfirmationChildRequirements = _child
	return _child, nil
}

func (m *_ReplyOrConfirmationReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReplyOrConfirmationReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReplyOrConfirmationReply")
		}

		// Simple Field (reply)
		if pushErr := writeBuffer.PushContext("reply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reply")
		}
		_replyErr := writeBuffer.WriteSerializable(m.GetReply())
		if popErr := writeBuffer.PopContext("reply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reply")
		}
		if _replyErr != nil {
			return errors.Wrap(_replyErr, "Error serializing 'reply' field")
		}

		// Simple Field (termination)
		if pushErr := writeBuffer.PushContext("termination"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for termination")
		}
		_terminationErr := writeBuffer.WriteSerializable(m.GetTermination())
		if popErr := writeBuffer.PopContext("termination"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for termination")
		}
		if _terminationErr != nil {
			return errors.Wrap(_terminationErr, "Error serializing 'termination' field")
		}

		if popErr := writeBuffer.PopContext("ReplyOrConfirmationReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReplyOrConfirmationReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ReplyOrConfirmationReply) isReplyOrConfirmationReply() bool {
	return true
}

func (m *_ReplyOrConfirmationReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
