/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */
package org.apache.plc4x.java.ads.api.commands;

import org.apache.plc4x.java.ads.api.commands.types.ADSState;
import org.apache.plc4x.java.ads.api.commands.types.Data;
import org.apache.plc4x.java.ads.api.commands.types.DeviceState;
import org.apache.plc4x.java.ads.api.commands.types.Length;
import org.apache.plc4x.java.ads.api.generic.ADSData;
import org.apache.plc4x.java.ads.api.generic.AMSHeader;
import org.apache.plc4x.java.ads.api.generic.AMSTCPHeader;
import org.apache.plc4x.java.ads.api.generic.types.AMSNetId;
import org.apache.plc4x.java.ads.api.generic.types.AMSPort;
import org.apache.plc4x.java.ads.api.generic.types.Command;
import org.apache.plc4x.java.ads.api.generic.types.Invoke;
import org.apache.plc4x.java.ads.api.util.LengthSupplier;

import static java.util.Objects.requireNonNull;

/**
 * Changes the ADS status and the device status of an ADS device.
 * Additionally it is possible to send data to the ADS device to transfer further information.
 * These data were not analysed from the current ADS devices (PLC, NC, ...)
 */
@ADSCommandType(Command.ADS_WRITE_CONTROL)
public class ADSWriteControlRequest extends ADSAbstractRequest {

    /**
     * 2 bytes	New ADS status (see data type ADSSTATE of the ADS-DLL).
     */
    private final ADSState adsState;
    /**
     * 2 bytes	New device status.
     */
    private final DeviceState deviceState;
    /**
     * 4 bytes	Length of data in byte.
     */
    private final Length length;
    /**
     * n bytes	Additional data which are sent to the ADS device
     */
    private final Data data;

    ////
    // Used when fields should be calculated. TODO: check if we better work with a subclass.
    private final LengthSupplier lengthSupplier;
    private final boolean calculated;
    //
    ///

    private ADSWriteControlRequest(AMSTCPHeader amstcpHeader, AMSHeader amsHeader, ADSState adsState, DeviceState deviceState, Length length, Data data) {
        super(amstcpHeader, amsHeader);
        this.adsState = requireNonNull(adsState);
        this.deviceState = requireNonNull(deviceState);
        this.length = requireNonNull(length);
        this.data = requireNonNull(data);
        this.lengthSupplier = null;
        this.calculated = false;
    }

    private ADSWriteControlRequest(AMSHeader amsHeader, ADSState adsState, DeviceState deviceState, Length length, Data data) {
        super(amsHeader);
        this.adsState = requireNonNull(adsState);
        this.deviceState = requireNonNull(deviceState);
        this.length = requireNonNull(length);
        this.data = requireNonNull(data);
        this.lengthSupplier = null;
        this.calculated = false;
    }

    private ADSWriteControlRequest(AMSNetId targetAmsNetId, AMSPort targetAmsPort, AMSNetId sourceAmsNetId, AMSPort sourceAmsPort, Invoke invokeId, ADSState adsState, DeviceState deviceState, Data data) {
        super(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, invokeId);
        this.adsState = requireNonNull(adsState);
        this.deviceState = requireNonNull(deviceState);
        this.length = null;
        this.data = requireNonNull(data);
        this.lengthSupplier = data;
        this.calculated = true;
    }

    public static ADSWriteControlRequest of(AMSTCPHeader amstcpHeader, AMSHeader amsHeader, ADSState adsState, DeviceState deviceState, Length length, Data data) {
        return new ADSWriteControlRequest(amstcpHeader, amsHeader, adsState, deviceState, length, data);
    }

    public static ADSWriteControlRequest of(AMSHeader amsHeader, ADSState adsState, DeviceState deviceState, Length length, Data data) {
        return new ADSWriteControlRequest(amsHeader, adsState, deviceState, length, data);
    }

    public static ADSWriteControlRequest of(AMSNetId targetAmsNetId, AMSPort targetAmsPort, AMSNetId sourceAmsNetId, AMSPort sourceAmsPort, Invoke invokeId, ADSState adsState, DeviceState deviceState, Data data) {
        return new ADSWriteControlRequest(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, invokeId, adsState, deviceState, data);
    }

    @Override
    public ADSData getAdsData() {
        return buildADSData(adsState, deviceState, calculated ? Length.of(lengthSupplier.getCalculatedLength()) : length, data);
    }

    @Override
    public String toString() {
        return "ADSWriteControlRequest{" +
            "adsState=" + adsState +
            ", deviceState=" + deviceState +
            ", length=" + (calculated ? Length.of(lengthSupplier.getCalculatedLength()) : length) +
            ", data=" + data +
            "} " + super.toString();
    }
}
