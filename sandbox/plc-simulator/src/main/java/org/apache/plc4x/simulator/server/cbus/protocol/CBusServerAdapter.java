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
package org.apache.plc4x.simulator.server.cbus.protocol;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import org.apache.plc4x.java.cbus.readwrite.*;
import org.apache.plc4x.simulator.model.Context;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class CBusServerAdapter extends ChannelInboundHandlerAdapter {

    private static final Logger LOGGER = LoggerFactory.getLogger(CBusServerAdapter.class);

    private Context context;

    private static final RequestContext requestContext = new RequestContext(false, false, false);
    private static final CBusOptions cBusOptions = new CBusOptions(false, false, false, false, false, false, false, false, true);

    private Lock writeLock = new ReentrantLock();

    private ScheduledFuture<?> sf;

    public CBusServerAdapter(Context context) {
        this.context = context;
    }

    @Override
    public void channelActive(ChannelHandlerContext ctx) throws Exception {
        super.channelActive(ctx);
        sf = ctx.executor().scheduleAtFixedRate(() -> {
            try {
                writeLock.lock();
                MonitoredSAL monitoredSAL;
                if (cBusOptions.getExstat()) {
                    // TODO implement long from
                    throw new IllegalStateException("Not implemented");
                } else {
                    LightingData lightingData;
                    double random = Math.random();
                    if (random < 0.25) {
                        lightingData = new LightingDataOn(LightingCommandTypeContainer.LightingCommandOn, (byte) 0xAF);
                    } else if (random > 0.25 && random < 0.5) {
                        lightingData = new LightingDataOff(LightingCommandTypeContainer.LightingCommandOff, (byte) 0xAF);
                    } else if (random > 0.5 && random < 0.75) {
                        lightingData = new LightingDataRampToLevel(LightingCommandTypeContainer.LightingCommandRampToLevel_20Second, (byte) 0xAF, (byte) 0xE0);
                    } else {
                        lightingData = new LightingDataTerminateRamp(LightingCommandTypeContainer.LightingCommandTerminateRamp, (byte) 0xAF);
                    }
                    SALData salData = new SALDataLighting(null, lightingData);
                    monitoredSAL = new MonitoredSALShortFormBasicMode((byte) 0x0, (byte) 0x0, (short) 0x0, (short) 0x0, (byte) 0x0, ApplicationIdContainer.LIGHTING_38, salData, cBusOptions);
                }
                EncodedReply encodedReply = new MonitoredSALReply((byte) 0x0, monitoredSAL, cBusOptions, requestContext);
                Reply reply = new ReplyEncodedReply((byte) 0x0, encodedReply, null, cBusOptions, requestContext);
                ReplyOrConfirmation replyOrConfirmation = new ReplyOrConfirmationReply((byte) 0x00, reply, new ResponseTermination(), cBusOptions, requestContext);
                CBusMessage message = new CBusMessageToClient(replyOrConfirmation, requestContext, cBusOptions);
                LOGGER.info("[Monitor] Sending out\n{}\n{}", message, encodedReply);
                ctx.writeAndFlush(message);
            } finally {
                writeLock.unlock();
            }
        }, 5, 5, TimeUnit.SECONDS);
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) throws Exception {
        sf.cancel(false);
        super.channelInactive(ctx);
    }

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        if (!(msg instanceof CBusMessage)) {
            return;
        }
        try {
            writeLock.lock();
            CBusMessage packet = (CBusMessage) msg;
            if (packet instanceof CBusMessageToClient) {
                LOGGER.info("Message to client not supported\n{}", packet);
                return;
            }
            CBusMessageToServer cBusMessageToServer = (CBusMessageToServer) packet;
            Request request = cBusMessageToServer.getRequest();
            if (request instanceof RequestEmpty || request instanceof RequestNull) {
                LOGGER.debug("Ignoring\n{}", request);
                return;
            }
            if (request instanceof RequestDirectCommandAccess) {
                RequestDirectCommandAccess requestDirectCommandAccess = (RequestDirectCommandAccess) request;
                LOGGER.info("Handling RequestDirectCommandAccess\n{}", requestDirectCommandAccess);
                // TODO: handle this
                return;
            }
            if (request instanceof RequestCommand) {
                RequestCommand requestCommand = (RequestCommand) request;
                LOGGER.info("Handling RequestCommand\n{}", requestCommand);
                CBusCommand cbusCommand = requestCommand.getCbusCommand();
                LOGGER.info("Handling CBusCommand\n{}", cbusCommand);
                if (cbusCommand instanceof CBusCommandPointToPoint) {
                    CBusCommandPointToPoint cBusCommandPointToPoint = (CBusCommandPointToPoint) cbusCommand;
                    LOGGER.info("Handling CBusCommandPointToPoint\n{}", cBusCommandPointToPoint);
                    // TODO: handle this
                    return;
                }
                if (cbusCommand instanceof CBusCommandPointToMultiPoint) {
                    CBusCommandPointToMultiPoint cBusCommandPointToMultiPoint = (CBusCommandPointToMultiPoint) cbusCommand;
                    CBusPointToMultiPointCommand command = cBusCommandPointToMultiPoint.getCommand();
                    if (command instanceof CBusPointToMultiPointCommandStatus) {
                        CBusPointToMultiPointCommandStatus cBusPointToMultiPointCommandStatus = (CBusPointToMultiPointCommandStatus) command;
                        StatusRequest statusRequest = cBusPointToMultiPointCommandStatus.getStatusRequest();
                        if (statusRequest instanceof StatusRequestBinaryState) {
                            StatusRequestBinaryState statusRequestBinaryState = (StatusRequestBinaryState) statusRequest;
                            StatusHeader statusHeader = new StatusHeader((short) (2 + 1)); // 2 we have always + 1 as we got one status byte
                            // TODO: map actuall values from simulator
                            byte blockStart = 0x0;
                            List<StatusByte> statusBytes = List.of(new StatusByte(GAVState.ON, GAVState.ERROR, GAVState.OFF, GAVState.DOES_NOT_EXIST));
                            // TODO: this might be extended or standard depeding on exstat
                            StandardFormatStatusReply standardFormatStatusReply = new StandardFormatStatusReply(statusHeader, statusRequestBinaryState.getApplication(), blockStart, statusBytes);
                            EncodedReply encodedReply = new EncodedReplyStandardFormatStatusReply((byte) 0xC0, standardFormatStatusReply, cBusOptions, requestContext);
                            ReplyEncodedReply replyEncodedReply = new ReplyEncodedReply((byte) 0xC0, encodedReply, null, cBusOptions, requestContext);
                            ReplyOrConfirmation replyOrConfirmation = new ReplyOrConfirmationReply((byte) 0xFF, replyEncodedReply, new ResponseTermination(), cBusOptions, requestContext);
                            Alpha alpha = requestCommand.getAlpha();
                            if (alpha != null) {
                                Confirmation confirmation = new Confirmation(alpha, null, ConfirmationType.CONFIRMATION_SUCCESSFUL);
                                replyOrConfirmation = new ReplyOrConfirmationConfirmation(alpha.getCharacter(), confirmation, replyOrConfirmation, cBusOptions, requestContext);
                            }
                            CBusMessage response = new CBusMessageToClient(replyOrConfirmation, requestContext, cBusOptions);
                            LOGGER.info("Send binary status response\n{}", response);
                            ctx.writeAndFlush(response);
                            return;
                        }
                        if (statusRequest instanceof StatusRequestBinaryStateDeprecated) {
                            StatusRequestBinaryStateDeprecated statusRequestBinaryStateDeprecated = (StatusRequestBinaryStateDeprecated) statusRequest;
                            LOGGER.info("Handling StatusRequestBinaryStateDeprecated\n{}", statusRequestBinaryStateDeprecated);
                            // TODO: handle this
                            return;
                        }
                        if (statusRequest instanceof StatusRequestLevel) {
                            StatusRequestLevel statusRequestLevel = (StatusRequestLevel) statusRequest;
                            ExtendedStatusHeader statusHeader = new ExtendedStatusHeader((short) (3 + 2)); // 3 we have always (coding is extra opposed to the standard) + 2 as we got one level information
                            StatusCoding coding = StatusCoding.LEVEL_BY_THIS_SERIAL_INTERFACE;
                            // TODO: map actuall values from simulator
                            byte blockStart = statusRequestLevel.getStartingGroupAddressLabel();
                            List<LevelInformation> levelInformations = List.of(new LevelInformationNormal(0x5555, LevelInformationNibblePair.Value_F, LevelInformationNibblePair.Value_F));
                            ExtendedFormatStatusReply extendedFormatStatusReply = new ExtendedFormatStatusReply(statusHeader, coding, statusRequestLevel.getApplication(), blockStart, null, levelInformations);
                            EncodedReply encodedReply = new EncodedReplyExtendedFormatStatusReply((byte) 0xC0, extendedFormatStatusReply, cBusOptions, requestContext);
                            ReplyEncodedReply replyEncodedReply = new ReplyEncodedReply((byte) 0xC0, encodedReply, null, cBusOptions, requestContext);
                            ReplyOrConfirmation replyOrConfirmation = new ReplyOrConfirmationReply((byte) 0xFF, replyEncodedReply, new ResponseTermination(), cBusOptions, requestContext);
                            Alpha alpha = requestCommand.getAlpha();
                            if (alpha != null) {
                                Confirmation confirmation = new Confirmation(alpha, null, ConfirmationType.CONFIRMATION_SUCCESSFUL);
                                replyOrConfirmation = new ReplyOrConfirmationConfirmation(alpha.getCharacter(), confirmation, replyOrConfirmation, cBusOptions, requestContext);
                            }
                            CBusMessage response = new CBusMessageToClient(replyOrConfirmation, requestContext, cBusOptions);
                            LOGGER.info("Send level status response\n{}", response);
                            ctx.writeAndFlush(response);
                            return;
                        }
                        // TODO: handle this
                        return;
                    }
                    if (command instanceof CBusPointToMultiPointCommandNormal) {
                        CBusPointToMultiPointCommandNormal cBusPointToMultiPointCommandNormal = (CBusPointToMultiPointCommandNormal) command;
                        LOGGER.info("Handling CBusPointToMultiPointCommandNormal\n{}", cBusPointToMultiPointCommandNormal);
                        return;
                    }
                    // TODO: handle this
                    return;
                }
                if (cbusCommand instanceof CBusCommandPointToPointToMultiPoint) {
                    CBusCommandPointToPointToMultiPoint cBusCommandPointToPointToMultiPoint = (CBusCommandPointToPointToMultiPoint) cbusCommand;
                    LOGGER.info("Handling CBusCommandPointToPointToMultiPoint\n{}", cBusCommandPointToPointToMultiPoint);
                    // TODO: handle this
                    return;
                }
                if (cbusCommand instanceof CBusCommandDeviceManagement) {
                    CBusCommandDeviceManagement cBusCommandDeviceManagement = (CBusCommandDeviceManagement) cbusCommand;
                    LOGGER.info("Handling CBusCommandDeviceManagement\n{}", cBusCommandDeviceManagement);
                    // TODO: handle this
                    return;
                }

                Alpha alpha = requestCommand.getAlpha();
                if (alpha != null) {
                    Confirmation confirmation = new Confirmation(alpha, null, ConfirmationType.NOT_TRANSMITTED_CORRUPTION);
                    ReplyOrConfirmationConfirmation replyOrConfirmationConfirmation = new ReplyOrConfirmationConfirmation(alpha.getCharacter(), confirmation, null, cBusOptions, requestContext);
                    CBusMessage response = new CBusMessageToClient(replyOrConfirmationConfirmation, requestContext, cBusOptions);
                    LOGGER.info("Send response\n{}", response);
                    ctx.writeAndFlush(response);
                }
                return;
            }
            if (request instanceof RequestObsolete) {
                RequestObsolete requestObsolete = (RequestObsolete) request;
                LOGGER.info("Handling RequestObsolete\n{}", requestObsolete);
                // TODO: handle this
                return;
            }
            if (request instanceof RequestReset) {
                RequestReset requestReset = (RequestReset) request;
                LOGGER.info("Handling RequestReset\n{}", requestReset);
                // TODO: handle this
                return;
            }
            if (request instanceof RequestSmartConnectShortcut) {
                RequestSmartConnectShortcut requestSmartConnectShortcut = (RequestSmartConnectShortcut) request;
                LOGGER.info("Handling RequestSmartConnectShortcut\n{}", requestSmartConnectShortcut);
                // TODO: handle this
                return;
            }
        } finally {
            writeLock.unlock();
        }
    }

}