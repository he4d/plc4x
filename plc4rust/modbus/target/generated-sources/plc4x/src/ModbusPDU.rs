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
// package org.apache.plc4x.rust.modbus.readwrite;

// Code generated by code-generation. DO NOT EDIT.
use std::io::{Error, ErrorKind, Read, Write};
use plc4rust::{Message, NoOption};
use plc4rust::read_buffer::ReadBuffer;
use plc4rust::write_buffer::WriteBuffer;

use crate::ModbusPDUError::ModbusPDUError;
use crate::ModbusPDUError::ModbusPDUErrorOptions;
use crate::ModbusPDUReadDiscreteInputsRequest::ModbusPDUReadDiscreteInputsRequest;
use crate::ModbusPDUReadDiscreteInputsRequest::ModbusPDUReadDiscreteInputsRequestOptions;
use crate::ModbusPDUReadDiscreteInputsResponse::ModbusPDUReadDiscreteInputsResponse;
use crate::ModbusPDUReadDiscreteInputsResponse::ModbusPDUReadDiscreteInputsResponseOptions;
use crate::ModbusPDUReadCoilsRequest::ModbusPDUReadCoilsRequest;
use crate::ModbusPDUReadCoilsRequest::ModbusPDUReadCoilsRequestOptions;
use crate::ModbusPDUReadCoilsResponse::ModbusPDUReadCoilsResponse;
use crate::ModbusPDUReadCoilsResponse::ModbusPDUReadCoilsResponseOptions;
use crate::ModbusPDUWriteSingleCoilRequest::ModbusPDUWriteSingleCoilRequest;
use crate::ModbusPDUWriteSingleCoilRequest::ModbusPDUWriteSingleCoilRequestOptions;
use crate::ModbusPDUWriteSingleCoilResponse::ModbusPDUWriteSingleCoilResponse;
use crate::ModbusPDUWriteSingleCoilResponse::ModbusPDUWriteSingleCoilResponseOptions;
use crate::ModbusPDUWriteMultipleCoilsRequest::ModbusPDUWriteMultipleCoilsRequest;
use crate::ModbusPDUWriteMultipleCoilsRequest::ModbusPDUWriteMultipleCoilsRequestOptions;
use crate::ModbusPDUWriteMultipleCoilsResponse::ModbusPDUWriteMultipleCoilsResponse;
use crate::ModbusPDUWriteMultipleCoilsResponse::ModbusPDUWriteMultipleCoilsResponseOptions;
use crate::ModbusPDUReadInputRegistersRequest::ModbusPDUReadInputRegistersRequest;
use crate::ModbusPDUReadInputRegistersRequest::ModbusPDUReadInputRegistersRequestOptions;
use crate::ModbusPDUReadInputRegistersResponse::ModbusPDUReadInputRegistersResponse;
use crate::ModbusPDUReadInputRegistersResponse::ModbusPDUReadInputRegistersResponseOptions;
use crate::ModbusPDUReadHoldingRegistersRequest::ModbusPDUReadHoldingRegistersRequest;
use crate::ModbusPDUReadHoldingRegistersRequest::ModbusPDUReadHoldingRegistersRequestOptions;
use crate::ModbusPDUReadHoldingRegistersResponse::ModbusPDUReadHoldingRegistersResponse;
use crate::ModbusPDUReadHoldingRegistersResponse::ModbusPDUReadHoldingRegistersResponseOptions;
use crate::ModbusPDUWriteSingleRegisterRequest::ModbusPDUWriteSingleRegisterRequest;
use crate::ModbusPDUWriteSingleRegisterRequest::ModbusPDUWriteSingleRegisterRequestOptions;
use crate::ModbusPDUWriteSingleRegisterResponse::ModbusPDUWriteSingleRegisterResponse;
use crate::ModbusPDUWriteSingleRegisterResponse::ModbusPDUWriteSingleRegisterResponseOptions;
use crate::ModbusPDUWriteMultipleHoldingRegistersRequest::ModbusPDUWriteMultipleHoldingRegistersRequest;
use crate::ModbusPDUWriteMultipleHoldingRegistersRequest::ModbusPDUWriteMultipleHoldingRegistersRequestOptions;
use crate::ModbusPDUWriteMultipleHoldingRegistersResponse::ModbusPDUWriteMultipleHoldingRegistersResponse;
use crate::ModbusPDUWriteMultipleHoldingRegistersResponse::ModbusPDUWriteMultipleHoldingRegistersResponseOptions;
use crate::ModbusPDUReadWriteMultipleHoldingRegistersRequest::ModbusPDUReadWriteMultipleHoldingRegistersRequest;
use crate::ModbusPDUReadWriteMultipleHoldingRegistersRequest::ModbusPDUReadWriteMultipleHoldingRegistersRequestOptions;
use crate::ModbusPDUReadWriteMultipleHoldingRegistersResponse::ModbusPDUReadWriteMultipleHoldingRegistersResponse;
use crate::ModbusPDUReadWriteMultipleHoldingRegistersResponse::ModbusPDUReadWriteMultipleHoldingRegistersResponseOptions;
use crate::ModbusPDUMaskWriteHoldingRegisterRequest::ModbusPDUMaskWriteHoldingRegisterRequest;
use crate::ModbusPDUMaskWriteHoldingRegisterRequest::ModbusPDUMaskWriteHoldingRegisterRequestOptions;
use crate::ModbusPDUMaskWriteHoldingRegisterResponse::ModbusPDUMaskWriteHoldingRegisterResponse;
use crate::ModbusPDUMaskWriteHoldingRegisterResponse::ModbusPDUMaskWriteHoldingRegisterResponseOptions;
use crate::ModbusPDUReadFifoQueueRequest::ModbusPDUReadFifoQueueRequest;
use crate::ModbusPDUReadFifoQueueRequest::ModbusPDUReadFifoQueueRequestOptions;
use crate::ModbusPDUReadFifoQueueResponse::ModbusPDUReadFifoQueueResponse;
use crate::ModbusPDUReadFifoQueueResponse::ModbusPDUReadFifoQueueResponseOptions;
use crate::ModbusPDUReadFileRecordRequest::ModbusPDUReadFileRecordRequest;
use crate::ModbusPDUReadFileRecordRequest::ModbusPDUReadFileRecordRequestOptions;
use crate::ModbusPDUReadFileRecordResponse::ModbusPDUReadFileRecordResponse;
use crate::ModbusPDUReadFileRecordResponse::ModbusPDUReadFileRecordResponseOptions;
use crate::ModbusPDUWriteFileRecordRequest::ModbusPDUWriteFileRecordRequest;
use crate::ModbusPDUWriteFileRecordRequest::ModbusPDUWriteFileRecordRequestOptions;
use crate::ModbusPDUWriteFileRecordResponse::ModbusPDUWriteFileRecordResponse;
use crate::ModbusPDUWriteFileRecordResponse::ModbusPDUWriteFileRecordResponseOptions;
use crate::ModbusPDUReadExceptionStatusRequest::ModbusPDUReadExceptionStatusRequest;
use crate::ModbusPDUReadExceptionStatusRequest::ModbusPDUReadExceptionStatusRequestOptions;
use crate::ModbusPDUReadExceptionStatusResponse::ModbusPDUReadExceptionStatusResponse;
use crate::ModbusPDUReadExceptionStatusResponse::ModbusPDUReadExceptionStatusResponseOptions;
use crate::ModbusPDUDiagnosticRequest::ModbusPDUDiagnosticRequest;
use crate::ModbusPDUDiagnosticRequest::ModbusPDUDiagnosticRequestOptions;
use crate::ModbusPDUDiagnosticResponse::ModbusPDUDiagnosticResponse;
use crate::ModbusPDUDiagnosticResponse::ModbusPDUDiagnosticResponseOptions;
use crate::ModbusPDUGetComEventCounterRequest::ModbusPDUGetComEventCounterRequest;
use crate::ModbusPDUGetComEventCounterRequest::ModbusPDUGetComEventCounterRequestOptions;
use crate::ModbusPDUGetComEventCounterResponse::ModbusPDUGetComEventCounterResponse;
use crate::ModbusPDUGetComEventCounterResponse::ModbusPDUGetComEventCounterResponseOptions;
use crate::ModbusPDUGetComEventLogRequest::ModbusPDUGetComEventLogRequest;
use crate::ModbusPDUGetComEventLogRequest::ModbusPDUGetComEventLogRequestOptions;
use crate::ModbusPDUGetComEventLogResponse::ModbusPDUGetComEventLogResponse;
use crate::ModbusPDUGetComEventLogResponse::ModbusPDUGetComEventLogResponseOptions;
use crate::ModbusPDUReportServerIdRequest::ModbusPDUReportServerIdRequest;
use crate::ModbusPDUReportServerIdRequest::ModbusPDUReportServerIdRequestOptions;
use crate::ModbusPDUReportServerIdResponse::ModbusPDUReportServerIdResponse;
use crate::ModbusPDUReportServerIdResponse::ModbusPDUReportServerIdResponseOptions;
use crate::ModbusPDUReadDeviceIdentificationRequest::ModbusPDUReadDeviceIdentificationRequest;
use crate::ModbusPDUReadDeviceIdentificationRequest::ModbusPDUReadDeviceIdentificationRequestOptions;
use crate::ModbusPDUReadDeviceIdentificationResponse::ModbusPDUReadDeviceIdentificationResponse;
use crate::ModbusPDUReadDeviceIdentificationResponse::ModbusPDUReadDeviceIdentificationResponseOptions;

#[derive(PartialEq, Debug, Clone)]
pub struct ModbusPDUOptions {
    pub response: bool
}

#[derive(PartialEq, Debug, Clone)]
pub enum ModbusPDU {
    ModbusPDUError(ModbusPDUError),
    ModbusPDUReadDiscreteInputsRequest(ModbusPDUReadDiscreteInputsRequest),
    ModbusPDUReadDiscreteInputsResponse(ModbusPDUReadDiscreteInputsResponse),
    ModbusPDUReadCoilsRequest(ModbusPDUReadCoilsRequest),
    ModbusPDUReadCoilsResponse(ModbusPDUReadCoilsResponse),
    ModbusPDUWriteSingleCoilRequest(ModbusPDUWriteSingleCoilRequest),
    ModbusPDUWriteSingleCoilResponse(ModbusPDUWriteSingleCoilResponse),
    ModbusPDUWriteMultipleCoilsRequest(ModbusPDUWriteMultipleCoilsRequest),
    ModbusPDUWriteMultipleCoilsResponse(ModbusPDUWriteMultipleCoilsResponse),
    ModbusPDUReadInputRegistersRequest(ModbusPDUReadInputRegistersRequest),
    ModbusPDUReadInputRegistersResponse(ModbusPDUReadInputRegistersResponse),
    ModbusPDUReadHoldingRegistersRequest(ModbusPDUReadHoldingRegistersRequest),
    ModbusPDUReadHoldingRegistersResponse(ModbusPDUReadHoldingRegistersResponse),
    ModbusPDUWriteSingleRegisterRequest(ModbusPDUWriteSingleRegisterRequest),
    ModbusPDUWriteSingleRegisterResponse(ModbusPDUWriteSingleRegisterResponse),
    ModbusPDUWriteMultipleHoldingRegistersRequest(ModbusPDUWriteMultipleHoldingRegistersRequest),
    ModbusPDUWriteMultipleHoldingRegistersResponse(ModbusPDUWriteMultipleHoldingRegistersResponse),
    ModbusPDUReadWriteMultipleHoldingRegistersRequest(ModbusPDUReadWriteMultipleHoldingRegistersRequest),
    ModbusPDUReadWriteMultipleHoldingRegistersResponse(ModbusPDUReadWriteMultipleHoldingRegistersResponse),
    ModbusPDUMaskWriteHoldingRegisterRequest(ModbusPDUMaskWriteHoldingRegisterRequest),
    ModbusPDUMaskWriteHoldingRegisterResponse(ModbusPDUMaskWriteHoldingRegisterResponse),
    ModbusPDUReadFifoQueueRequest(ModbusPDUReadFifoQueueRequest),
    ModbusPDUReadFifoQueueResponse(ModbusPDUReadFifoQueueResponse),
    ModbusPDUReadFileRecordRequest(ModbusPDUReadFileRecordRequest),
    ModbusPDUReadFileRecordResponse(ModbusPDUReadFileRecordResponse),
    ModbusPDUWriteFileRecordRequest(ModbusPDUWriteFileRecordRequest),
    ModbusPDUWriteFileRecordResponse(ModbusPDUWriteFileRecordResponse),
    ModbusPDUReadExceptionStatusRequest(ModbusPDUReadExceptionStatusRequest),
    ModbusPDUReadExceptionStatusResponse(ModbusPDUReadExceptionStatusResponse),
    ModbusPDUDiagnosticRequest(ModbusPDUDiagnosticRequest),
    ModbusPDUDiagnosticResponse(ModbusPDUDiagnosticResponse),
    ModbusPDUGetComEventCounterRequest(ModbusPDUGetComEventCounterRequest),
    ModbusPDUGetComEventCounterResponse(ModbusPDUGetComEventCounterResponse),
    ModbusPDUGetComEventLogRequest(ModbusPDUGetComEventLogRequest),
    ModbusPDUGetComEventLogResponse(ModbusPDUGetComEventLogResponse),
    ModbusPDUReportServerIdRequest(ModbusPDUReportServerIdRequest),
    ModbusPDUReportServerIdResponse(ModbusPDUReportServerIdResponse),
    ModbusPDUReadDeviceIdentificationRequest(ModbusPDUReadDeviceIdentificationRequest),
    ModbusPDUReadDeviceIdentificationResponse(ModbusPDUReadDeviceIdentificationResponse)
}

impl Message for ModbusPDU {
    type M = ModbusPDU;
    type P = ModbusPDUOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        match self {
            ModbusPDU::ModbusPDUError(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadDiscreteInputsRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadDiscreteInputsResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadCoilsRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadCoilsResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteSingleCoilRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteSingleCoilResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteMultipleCoilsRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteMultipleCoilsResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadInputRegistersRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadInputRegistersResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadHoldingRegistersRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadHoldingRegistersResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteSingleRegisterRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteSingleRegisterResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteMultipleHoldingRegistersRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteMultipleHoldingRegistersResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadWriteMultipleHoldingRegistersRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadWriteMultipleHoldingRegistersResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUMaskWriteHoldingRegisterRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUMaskWriteHoldingRegisterResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadFifoQueueRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadFifoQueueResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadFileRecordRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadFileRecordResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteFileRecordRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUWriteFileRecordResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadExceptionStatusRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadExceptionStatusResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUDiagnosticRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUDiagnosticResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUGetComEventCounterRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUGetComEventCounterResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUGetComEventLogRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUGetComEventLogResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReportServerIdRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReportServerIdResponse(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadDeviceIdentificationRequest(msg) => {
                msg.serialize(writer)
            }
            ModbusPDU::ModbusPDUReadDeviceIdentificationResponse(msg) => {
                msg.serialize(writer)
            }
        }
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let response = parameter.response;
        let errorFlag = reader.read_bit()?;
        let functionFlag = reader.read_u_n(7)? as u8;
        match (errorFlag, functionFlag, response) {
            (true, _, _) => {
                Ok(ModbusPDU::ModbusPDUError(ModbusPDUError::parse::<T>(reader, Some(ModbusPDUErrorOptions {
                    response
                }))?))
            }
            (false, 0x02, false) => {
                Ok(ModbusPDU::ModbusPDUReadDiscreteInputsRequest(ModbusPDUReadDiscreteInputsRequest::parse::<T>(reader, Some(ModbusPDUReadDiscreteInputsRequestOptions {
                    response
                }))?))
            }
            (false, 0x02, true) => {
                Ok(ModbusPDU::ModbusPDUReadDiscreteInputsResponse(ModbusPDUReadDiscreteInputsResponse::parse::<T>(reader, Some(ModbusPDUReadDiscreteInputsResponseOptions {
                    response
                }))?))
            }
            (false, 0x01, false) => {
                Ok(ModbusPDU::ModbusPDUReadCoilsRequest(ModbusPDUReadCoilsRequest::parse::<T>(reader, Some(ModbusPDUReadCoilsRequestOptions {
                    response
                }))?))
            }
            (false, 0x01, true) => {
                Ok(ModbusPDU::ModbusPDUReadCoilsResponse(ModbusPDUReadCoilsResponse::parse::<T>(reader, Some(ModbusPDUReadCoilsResponseOptions {
                    response
                }))?))
            }
            (false, 0x05, false) => {
                Ok(ModbusPDU::ModbusPDUWriteSingleCoilRequest(ModbusPDUWriteSingleCoilRequest::parse::<T>(reader, Some(ModbusPDUWriteSingleCoilRequestOptions {
                    response
                }))?))
            }
            (false, 0x05, true) => {
                Ok(ModbusPDU::ModbusPDUWriteSingleCoilResponse(ModbusPDUWriteSingleCoilResponse::parse::<T>(reader, Some(ModbusPDUWriteSingleCoilResponseOptions {
                    response
                }))?))
            }
            (false, 0x0F, false) => {
                Ok(ModbusPDU::ModbusPDUWriteMultipleCoilsRequest(ModbusPDUWriteMultipleCoilsRequest::parse::<T>(reader, Some(ModbusPDUWriteMultipleCoilsRequestOptions {
                    response
                }))?))
            }
            (false, 0x0F, true) => {
                Ok(ModbusPDU::ModbusPDUWriteMultipleCoilsResponse(ModbusPDUWriteMultipleCoilsResponse::parse::<T>(reader, Some(ModbusPDUWriteMultipleCoilsResponseOptions {
                    response
                }))?))
            }
            (false, 0x04, false) => {
                Ok(ModbusPDU::ModbusPDUReadInputRegistersRequest(ModbusPDUReadInputRegistersRequest::parse::<T>(reader, Some(ModbusPDUReadInputRegistersRequestOptions {
                    response
                }))?))
            }
            (false, 0x04, true) => {
                Ok(ModbusPDU::ModbusPDUReadInputRegistersResponse(ModbusPDUReadInputRegistersResponse::parse::<T>(reader, Some(ModbusPDUReadInputRegistersResponseOptions {
                    response
                }))?))
            }
            (false, 0x03, false) => {
                Ok(ModbusPDU::ModbusPDUReadHoldingRegistersRequest(ModbusPDUReadHoldingRegistersRequest::parse::<T>(reader, Some(ModbusPDUReadHoldingRegistersRequestOptions {
                    response
                }))?))
            }
            (false, 0x03, true) => {
                Ok(ModbusPDU::ModbusPDUReadHoldingRegistersResponse(ModbusPDUReadHoldingRegistersResponse::parse::<T>(reader, Some(ModbusPDUReadHoldingRegistersResponseOptions {
                    response
                }))?))
            }
            (false, 0x06, false) => {
                Ok(ModbusPDU::ModbusPDUWriteSingleRegisterRequest(ModbusPDUWriteSingleRegisterRequest::parse::<T>(reader, Some(ModbusPDUWriteSingleRegisterRequestOptions {
                    response
                }))?))
            }
            (false, 0x06, true) => {
                Ok(ModbusPDU::ModbusPDUWriteSingleRegisterResponse(ModbusPDUWriteSingleRegisterResponse::parse::<T>(reader, Some(ModbusPDUWriteSingleRegisterResponseOptions {
                    response
                }))?))
            }
            (false, 0x10, false) => {
                Ok(ModbusPDU::ModbusPDUWriteMultipleHoldingRegistersRequest(ModbusPDUWriteMultipleHoldingRegistersRequest::parse::<T>(reader, Some(ModbusPDUWriteMultipleHoldingRegistersRequestOptions {
                    response
                }))?))
            }
            (false, 0x10, true) => {
                Ok(ModbusPDU::ModbusPDUWriteMultipleHoldingRegistersResponse(ModbusPDUWriteMultipleHoldingRegistersResponse::parse::<T>(reader, Some(ModbusPDUWriteMultipleHoldingRegistersResponseOptions {
                    response
                }))?))
            }
            (false, 0x17, false) => {
                Ok(ModbusPDU::ModbusPDUReadWriteMultipleHoldingRegistersRequest(ModbusPDUReadWriteMultipleHoldingRegistersRequest::parse::<T>(reader, Some(ModbusPDUReadWriteMultipleHoldingRegistersRequestOptions {
                    response
                }))?))
            }
            (false, 0x17, true) => {
                Ok(ModbusPDU::ModbusPDUReadWriteMultipleHoldingRegistersResponse(ModbusPDUReadWriteMultipleHoldingRegistersResponse::parse::<T>(reader, Some(ModbusPDUReadWriteMultipleHoldingRegistersResponseOptions {
                    response
                }))?))
            }
            (false, 0x16, false) => {
                Ok(ModbusPDU::ModbusPDUMaskWriteHoldingRegisterRequest(ModbusPDUMaskWriteHoldingRegisterRequest::parse::<T>(reader, Some(ModbusPDUMaskWriteHoldingRegisterRequestOptions {
                    response
                }))?))
            }
            (false, 0x16, true) => {
                Ok(ModbusPDU::ModbusPDUMaskWriteHoldingRegisterResponse(ModbusPDUMaskWriteHoldingRegisterResponse::parse::<T>(reader, Some(ModbusPDUMaskWriteHoldingRegisterResponseOptions {
                    response
                }))?))
            }
            (false, 0x18, false) => {
                Ok(ModbusPDU::ModbusPDUReadFifoQueueRequest(ModbusPDUReadFifoQueueRequest::parse::<T>(reader, Some(ModbusPDUReadFifoQueueRequestOptions {
                    response
                }))?))
            }
            (false, 0x18, true) => {
                Ok(ModbusPDU::ModbusPDUReadFifoQueueResponse(ModbusPDUReadFifoQueueResponse::parse::<T>(reader, Some(ModbusPDUReadFifoQueueResponseOptions {
                    response
                }))?))
            }
            (false, 0x14, false) => {
                Ok(ModbusPDU::ModbusPDUReadFileRecordRequest(ModbusPDUReadFileRecordRequest::parse::<T>(reader, Some(ModbusPDUReadFileRecordRequestOptions {
                    response
                }))?))
            }
            (false, 0x14, true) => {
                Ok(ModbusPDU::ModbusPDUReadFileRecordResponse(ModbusPDUReadFileRecordResponse::parse::<T>(reader, Some(ModbusPDUReadFileRecordResponseOptions {
                    response
                }))?))
            }
            (false, 0x15, false) => {
                Ok(ModbusPDU::ModbusPDUWriteFileRecordRequest(ModbusPDUWriteFileRecordRequest::parse::<T>(reader, Some(ModbusPDUWriteFileRecordRequestOptions {
                    response
                }))?))
            }
            (false, 0x15, true) => {
                Ok(ModbusPDU::ModbusPDUWriteFileRecordResponse(ModbusPDUWriteFileRecordResponse::parse::<T>(reader, Some(ModbusPDUWriteFileRecordResponseOptions {
                    response
                }))?))
            }
            (false, 0x07, false) => {
                Ok(ModbusPDU::ModbusPDUReadExceptionStatusRequest(ModbusPDUReadExceptionStatusRequest::parse::<T>(reader, Some(ModbusPDUReadExceptionStatusRequestOptions {
                    response
                }))?))
            }
            (false, 0x07, true) => {
                Ok(ModbusPDU::ModbusPDUReadExceptionStatusResponse(ModbusPDUReadExceptionStatusResponse::parse::<T>(reader, Some(ModbusPDUReadExceptionStatusResponseOptions {
                    response
                }))?))
            }
            (false, 0x08, false) => {
                Ok(ModbusPDU::ModbusPDUDiagnosticRequest(ModbusPDUDiagnosticRequest::parse::<T>(reader, Some(ModbusPDUDiagnosticRequestOptions {
                    response
                }))?))
            }
            (false, 0x08, true) => {
                Ok(ModbusPDU::ModbusPDUDiagnosticResponse(ModbusPDUDiagnosticResponse::parse::<T>(reader, Some(ModbusPDUDiagnosticResponseOptions {
                    response
                }))?))
            }
            (false, 0x0B, false) => {
                Ok(ModbusPDU::ModbusPDUGetComEventCounterRequest(ModbusPDUGetComEventCounterRequest::parse::<T>(reader, Some(ModbusPDUGetComEventCounterRequestOptions {
                    response
                }))?))
            }
            (false, 0x0B, true) => {
                Ok(ModbusPDU::ModbusPDUGetComEventCounterResponse(ModbusPDUGetComEventCounterResponse::parse::<T>(reader, Some(ModbusPDUGetComEventCounterResponseOptions {
                    response
                }))?))
            }
            (false, 0x0C, false) => {
                Ok(ModbusPDU::ModbusPDUGetComEventLogRequest(ModbusPDUGetComEventLogRequest::parse::<T>(reader, Some(ModbusPDUGetComEventLogRequestOptions {
                    response
                }))?))
            }
            (false, 0x0C, true) => {
                Ok(ModbusPDU::ModbusPDUGetComEventLogResponse(ModbusPDUGetComEventLogResponse::parse::<T>(reader, Some(ModbusPDUGetComEventLogResponseOptions {
                    response
                }))?))
            }
            (false, 0x11, false) => {
                Ok(ModbusPDU::ModbusPDUReportServerIdRequest(ModbusPDUReportServerIdRequest::parse::<T>(reader, Some(ModbusPDUReportServerIdRequestOptions {
                    response
                }))?))
            }
            (false, 0x11, true) => {
                Ok(ModbusPDU::ModbusPDUReportServerIdResponse(ModbusPDUReportServerIdResponse::parse::<T>(reader, Some(ModbusPDUReportServerIdResponseOptions {
                    response
                }))?))
            }
            (false, 0x2B, false) => {
                Ok(ModbusPDU::ModbusPDUReadDeviceIdentificationRequest(ModbusPDUReadDeviceIdentificationRequest::parse::<T>(reader, Some(ModbusPDUReadDeviceIdentificationRequestOptions {
                    response
                }))?))
            }
            (false, 0x2B, true) => {
                Ok(ModbusPDU::ModbusPDUReadDeviceIdentificationResponse(ModbusPDUReadDeviceIdentificationResponse::parse::<T>(reader, Some(ModbusPDUReadDeviceIdentificationResponseOptions {
                    response
                }))?))
            }
            _ => {
                panic!("Unable to parse!");
            }
        }
    }
}


