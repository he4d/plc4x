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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// LabelType is an enum
type LabelType uint8

type ILabelType interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	LabelType_TEXT_LABEL             LabelType = 0
	LabelType_PREDEFINED_ICON        LabelType = 1
	LabelType_LOAD_DYNAMIC_ICON      LabelType = 2
	LabelType_SET_PREFERRED_LANGUAGE LabelType = 3
)

var LabelTypeValues []LabelType

func init() {
	_ = errors.New
	LabelTypeValues = []LabelType{
		LabelType_TEXT_LABEL,
		LabelType_PREDEFINED_ICON,
		LabelType_LOAD_DYNAMIC_ICON,
		LabelType_SET_PREFERRED_LANGUAGE,
	}
}

func LabelTypeByValue(value uint8) (enum LabelType, ok bool) {
	switch value {
	case 0:
		return LabelType_TEXT_LABEL, true
	case 1:
		return LabelType_PREDEFINED_ICON, true
	case 2:
		return LabelType_LOAD_DYNAMIC_ICON, true
	case 3:
		return LabelType_SET_PREFERRED_LANGUAGE, true
	}
	return 0, false
}

func LabelTypeByName(value string) (enum LabelType, ok bool) {
	switch value {
	case "TEXT_LABEL":
		return LabelType_TEXT_LABEL, true
	case "PREDEFINED_ICON":
		return LabelType_PREDEFINED_ICON, true
	case "LOAD_DYNAMIC_ICON":
		return LabelType_LOAD_DYNAMIC_ICON, true
	case "SET_PREFERRED_LANGUAGE":
		return LabelType_SET_PREFERRED_LANGUAGE, true
	}
	return 0, false
}

func LabelTypeKnows(value uint8) bool {
	for _, typeValue := range LabelTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastLabelType(structType interface{}) LabelType {
	castFunc := func(typ interface{}) LabelType {
		if sLabelType, ok := typ.(LabelType); ok {
			return sLabelType
		}
		return 0
	}
	return castFunc(structType)
}

func (m LabelType) GetLengthInBits() uint16 {
	return 2
}

func (m LabelType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LabelTypeParse(readBuffer utils.ReadBuffer) (LabelType, error) {
	val, err := readBuffer.ReadUint8("LabelType", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LabelType")
	}
	if enum, ok := LabelTypeByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return LabelType(val), nil
	} else {
		return enum, nil
	}
}

func (e LabelType) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("LabelType", 2, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LabelType) PLC4XEnumName() string {
	switch e {
	case LabelType_TEXT_LABEL:
		return "TEXT_LABEL"
	case LabelType_PREDEFINED_ICON:
		return "PREDEFINED_ICON"
	case LabelType_LOAD_DYNAMIC_ICON:
		return "LOAD_DYNAMIC_ICON"
	case LabelType_SET_PREFERRED_LANGUAGE:
		return "SET_PREFERRED_LANGUAGE"
	}
	return ""
}

func (e LabelType) String() string {
	return e.PLC4XEnumName()
}