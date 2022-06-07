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

use crate::DriverType::DriverType;
use crate::ModbusPDU::ModbusPDU;
use crate::ModbusPDU::ModbusPDUOptions;

#[derive(PartialEq, Debug, Clone)]
pub struct ModbusTcpADUOptions {
    pub driverType: DriverType, 
    pub response: bool
}

#[derive(PartialEq, Debug, Clone)]
pub struct ModbusTcpADU {
    pub transactionIdentifier: u16,
    pub unitIdentifier: u8,
    pub pdu: ModbusPDU
}

impl ModbusTcpADU {
    pub fn length(&self) -> u16 {
        (self.pdu.get_length_in_bytes() + 1) as u16
    }
}

impl Message for ModbusTcpADU {
    type M = ModbusTcpADU;
    type P = ModbusTcpADUOptions;

    fn get_length_in_bits(&self) -> u32 {
        todo!()
    }

    fn serialize<T: Write>(&self, writer: &mut WriteBuffer<T>) -> Result<usize, Error> {
        writer.write_u16(self.transactionIdentifier)?;
        writer.write_u16(0x0000)?;
        writer.write_u16(self.length())?;
        writer.write_u8(self.unitIdentifier)?;
        self.pdu.serialize(writer)?;
        Ok(0)
    }

    fn parse<T: Read>(reader: &mut ReadBuffer<T>, parameter: Option<Self::P>) -> Result<Self::M, Error> {
        // (Re-)define the options
        let parameter = parameter.unwrap();
        let driverType = parameter.driverType;
        let response = parameter.response;
        let transactionIdentifier = reader.read_u16()?;
        let protocolIdentifier = reader.read_u16()?;
        // assert value of constant
        assert_eq!(0x0000, protocolIdentifier);
        let length = reader.read_u16()?;
        let unitIdentifier = reader.read_u8()?;
        let pdu = ModbusPDU::parse(reader, Some(ModbusPDUOptions { response }))?;
        Ok(Self::M {
            transactionIdentifier,
            unitIdentifier,
            pdu
        })
    }
}


