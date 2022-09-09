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

// Code generated by "plc4xgenerator -type=DefaultPlcBrowseQueryResult"; DO NOT EDIT.

package model

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *DefaultPlcBrowseQueryResult) Serialize(writeBuffer utils.WriteBuffer) error {
	if err := writeBuffer.PushContext("PlcBrowseQueryResult"); err != nil {
		return err
	}

	if d.Field != nil {
		if serializableField, ok := d.Field.(utils.Serializable); ok {
			if err := writeBuffer.PushContext("field"); err != nil {
				return err
			}
			if err := serializableField.Serialize(writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("field"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.Field)
			if err := writeBuffer.WriteString("field", uint32(len(stringValue)*8), "UTF-8", stringValue); err != nil {
				return err
			}
		}
	}

	if err := writeBuffer.WriteString("name", uint32(len(d.Name)*8), "UTF-8", d.Name); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("readable", d.Readable); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("writable", d.Writable); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("subscribable", d.Subscribable); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("possibleDataTypes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.PossibleDataTypes {
		if err := writeBuffer.WriteString("", uint32(len(elem)*8), "UTF-8", elem); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("possibleDataTypes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("attributes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for name, elemValue := range d.Attributes {
		if serializable, ok := elemValue.(utils.Serializable); ok {
			if err := writeBuffer.PushContext(name); err != nil {
				return err
			}
			if err := serializable.Serialize(writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(name); err != nil {
				return err
			}
		} else {
			elemValueAsString := fmt.Sprintf("%v", elemValue)
			if err := writeBuffer.WriteString(name, uint32(len(elemValueAsString)*8), "UTF-8", elemValueAsString); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("attributes", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("PlcBrowseQueryResult"); err != nil {
		return err
	}
	return nil
}

func (d *DefaultPlcBrowseQueryResult) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
