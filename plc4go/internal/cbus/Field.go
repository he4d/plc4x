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

package cbus

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
)

type PlcField interface {
	GetApplicationId() readWriteModel.ApplicationId
}

type plcField struct {
	NumElements uint16
}

func NewField(numElements uint16) PlcField {
	return &plcField{
		NumElements: numElements,
	}
}

func (m plcField) GetAddressString() string {
	// TODO: implement me
	return fmt.Sprintf("TODO[%d]", m.NumElements)
}

func (m plcField) GetTypeName() string {
	return "TODO"
}

func (m plcField) GetApplicationId() readWriteModel.ApplicationId {
	//TODO implement me
	panic("implement me")
}

func (m plcField) GetQuantity() uint16 {
	return m.NumElements
}

func (m plcField) Serialize(writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("TODO"); err != nil {
		return err
	}

	if err := writeBuffer.PopContext("TODO"); err != nil {
		return err
	}
	return nil
}