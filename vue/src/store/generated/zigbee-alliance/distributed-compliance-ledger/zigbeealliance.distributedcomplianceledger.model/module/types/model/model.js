/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model';
const baseModel = {
    vid: 0,
    pid: 0,
    deviceTypeId: 0,
    productName: '',
    productLabel: '',
    partNumber: '',
    commissioningCustomFlow: 0,
    commissioningCustomFlowUrl: '',
    commissioningModeInitialStepsHint: 0,
    commissioningModeInitialStepsInstruction: '',
    commissioningModeSecondaryStepsHint: 0,
    commissioningModeSecondaryStepsInstruction: '',
    userManualUrl: '',
    supportUrl: '',
    productUrl: '',
    lsfUrl: '',
    lsfRevision: 0,
    creator: ''
};
export const Model = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.deviceTypeId !== 0) {
            writer.uint32(24).int32(message.deviceTypeId);
        }
        if (message.productName !== '') {
            writer.uint32(34).string(message.productName);
        }
        if (message.productLabel !== '') {
            writer.uint32(42).string(message.productLabel);
        }
        if (message.partNumber !== '') {
            writer.uint32(50).string(message.partNumber);
        }
        if (message.commissioningCustomFlow !== 0) {
            writer.uint32(56).int32(message.commissioningCustomFlow);
        }
        if (message.commissioningCustomFlowUrl !== '') {
            writer.uint32(66).string(message.commissioningCustomFlowUrl);
        }
        if (message.commissioningModeInitialStepsHint !== 0) {
            writer.uint32(72).uint32(message.commissioningModeInitialStepsHint);
        }
        if (message.commissioningModeInitialStepsInstruction !== '') {
            writer.uint32(82).string(message.commissioningModeInitialStepsInstruction);
        }
        if (message.commissioningModeSecondaryStepsHint !== 0) {
            writer.uint32(88).uint32(message.commissioningModeSecondaryStepsHint);
        }
        if (message.commissioningModeSecondaryStepsInstruction !== '') {
            writer.uint32(98).string(message.commissioningModeSecondaryStepsInstruction);
        }
        if (message.userManualUrl !== '') {
            writer.uint32(106).string(message.userManualUrl);
        }
        if (message.supportUrl !== '') {
            writer.uint32(114).string(message.supportUrl);
        }
        if (message.productUrl !== '') {
            writer.uint32(122).string(message.productUrl);
        }
        if (message.lsfUrl !== '') {
            writer.uint32(130).string(message.lsfUrl);
        }
        if (message.lsfRevision !== 0) {
            writer.uint32(136).int32(message.lsfRevision);
        }
        if (message.creator !== '') {
            writer.uint32(146).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseModel };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.pid = reader.int32();
                    break;
                case 3:
                    message.deviceTypeId = reader.int32();
                    break;
                case 4:
                    message.productName = reader.string();
                    break;
                case 5:
                    message.productLabel = reader.string();
                    break;
                case 6:
                    message.partNumber = reader.string();
                    break;
                case 7:
                    message.commissioningCustomFlow = reader.int32();
                    break;
                case 8:
                    message.commissioningCustomFlowUrl = reader.string();
                    break;
                case 9:
                    message.commissioningModeInitialStepsHint = reader.uint32();
                    break;
                case 10:
                    message.commissioningModeInitialStepsInstruction = reader.string();
                    break;
                case 11:
                    message.commissioningModeSecondaryStepsHint = reader.uint32();
                    break;
                case 12:
                    message.commissioningModeSecondaryStepsInstruction = reader.string();
                    break;
                case 13:
                    message.userManualUrl = reader.string();
                    break;
                case 14:
                    message.supportUrl = reader.string();
                    break;
                case 15:
                    message.productUrl = reader.string();
                    break;
                case 16:
                    message.lsfUrl = reader.string();
                    break;
                case 17:
                    message.lsfRevision = reader.int32();
                    break;
                case 18:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseModel };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.deviceTypeId !== undefined && object.deviceTypeId !== null) {
            message.deviceTypeId = Number(object.deviceTypeId);
        }
        else {
            message.deviceTypeId = 0;
        }
        if (object.productName !== undefined && object.productName !== null) {
            message.productName = String(object.productName);
        }
        else {
            message.productName = '';
        }
        if (object.productLabel !== undefined && object.productLabel !== null) {
            message.productLabel = String(object.productLabel);
        }
        else {
            message.productLabel = '';
        }
        if (object.partNumber !== undefined && object.partNumber !== null) {
            message.partNumber = String(object.partNumber);
        }
        else {
            message.partNumber = '';
        }
        if (object.commissioningCustomFlow !== undefined && object.commissioningCustomFlow !== null) {
            message.commissioningCustomFlow = Number(object.commissioningCustomFlow);
        }
        else {
            message.commissioningCustomFlow = 0;
        }
        if (object.commissioningCustomFlowUrl !== undefined && object.commissioningCustomFlowUrl !== null) {
            message.commissioningCustomFlowUrl = String(object.commissioningCustomFlowUrl);
        }
        else {
            message.commissioningCustomFlowUrl = '';
        }
        if (object.commissioningModeInitialStepsHint !== undefined && object.commissioningModeInitialStepsHint !== null) {
            message.commissioningModeInitialStepsHint = Number(object.commissioningModeInitialStepsHint);
        }
        else {
            message.commissioningModeInitialStepsHint = 0;
        }
        if (object.commissioningModeInitialStepsInstruction !== undefined && object.commissioningModeInitialStepsInstruction !== null) {
            message.commissioningModeInitialStepsInstruction = String(object.commissioningModeInitialStepsInstruction);
        }
        else {
            message.commissioningModeInitialStepsInstruction = '';
        }
        if (object.commissioningModeSecondaryStepsHint !== undefined && object.commissioningModeSecondaryStepsHint !== null) {
            message.commissioningModeSecondaryStepsHint = Number(object.commissioningModeSecondaryStepsHint);
        }
        else {
            message.commissioningModeSecondaryStepsHint = 0;
        }
        if (object.commissioningModeSecondaryStepsInstruction !== undefined && object.commissioningModeSecondaryStepsInstruction !== null) {
            message.commissioningModeSecondaryStepsInstruction = String(object.commissioningModeSecondaryStepsInstruction);
        }
        else {
            message.commissioningModeSecondaryStepsInstruction = '';
        }
        if (object.userManualUrl !== undefined && object.userManualUrl !== null) {
            message.userManualUrl = String(object.userManualUrl);
        }
        else {
            message.userManualUrl = '';
        }
        if (object.supportUrl !== undefined && object.supportUrl !== null) {
            message.supportUrl = String(object.supportUrl);
        }
        else {
            message.supportUrl = '';
        }
        if (object.productUrl !== undefined && object.productUrl !== null) {
            message.productUrl = String(object.productUrl);
        }
        else {
            message.productUrl = '';
        }
        if (object.lsfUrl !== undefined && object.lsfUrl !== null) {
            message.lsfUrl = String(object.lsfUrl);
        }
        else {
            message.lsfUrl = '';
        }
        if (object.lsfRevision !== undefined && object.lsfRevision !== null) {
            message.lsfRevision = Number(object.lsfRevision);
        }
        else {
            message.lsfRevision = 0;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.deviceTypeId !== undefined && (obj.deviceTypeId = message.deviceTypeId);
        message.productName !== undefined && (obj.productName = message.productName);
        message.productLabel !== undefined && (obj.productLabel = message.productLabel);
        message.partNumber !== undefined && (obj.partNumber = message.partNumber);
        message.commissioningCustomFlow !== undefined && (obj.commissioningCustomFlow = message.commissioningCustomFlow);
        message.commissioningCustomFlowUrl !== undefined && (obj.commissioningCustomFlowUrl = message.commissioningCustomFlowUrl);
        message.commissioningModeInitialStepsHint !== undefined && (obj.commissioningModeInitialStepsHint = message.commissioningModeInitialStepsHint);
        message.commissioningModeInitialStepsInstruction !== undefined &&
            (obj.commissioningModeInitialStepsInstruction = message.commissioningModeInitialStepsInstruction);
        message.commissioningModeSecondaryStepsHint !== undefined && (obj.commissioningModeSecondaryStepsHint = message.commissioningModeSecondaryStepsHint);
        message.commissioningModeSecondaryStepsInstruction !== undefined &&
            (obj.commissioningModeSecondaryStepsInstruction = message.commissioningModeSecondaryStepsInstruction);
        message.userManualUrl !== undefined && (obj.userManualUrl = message.userManualUrl);
        message.supportUrl !== undefined && (obj.supportUrl = message.supportUrl);
        message.productUrl !== undefined && (obj.productUrl = message.productUrl);
        message.lsfUrl !== undefined && (obj.lsfUrl = message.lsfUrl);
        message.lsfRevision !== undefined && (obj.lsfRevision = message.lsfRevision);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseModel };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.deviceTypeId !== undefined && object.deviceTypeId !== null) {
            message.deviceTypeId = object.deviceTypeId;
        }
        else {
            message.deviceTypeId = 0;
        }
        if (object.productName !== undefined && object.productName !== null) {
            message.productName = object.productName;
        }
        else {
            message.productName = '';
        }
        if (object.productLabel !== undefined && object.productLabel !== null) {
            message.productLabel = object.productLabel;
        }
        else {
            message.productLabel = '';
        }
        if (object.partNumber !== undefined && object.partNumber !== null) {
            message.partNumber = object.partNumber;
        }
        else {
            message.partNumber = '';
        }
        if (object.commissioningCustomFlow !== undefined && object.commissioningCustomFlow !== null) {
            message.commissioningCustomFlow = object.commissioningCustomFlow;
        }
        else {
            message.commissioningCustomFlow = 0;
        }
        if (object.commissioningCustomFlowUrl !== undefined && object.commissioningCustomFlowUrl !== null) {
            message.commissioningCustomFlowUrl = object.commissioningCustomFlowUrl;
        }
        else {
            message.commissioningCustomFlowUrl = '';
        }
        if (object.commissioningModeInitialStepsHint !== undefined && object.commissioningModeInitialStepsHint !== null) {
            message.commissioningModeInitialStepsHint = object.commissioningModeInitialStepsHint;
        }
        else {
            message.commissioningModeInitialStepsHint = 0;
        }
        if (object.commissioningModeInitialStepsInstruction !== undefined && object.commissioningModeInitialStepsInstruction !== null) {
            message.commissioningModeInitialStepsInstruction = object.commissioningModeInitialStepsInstruction;
        }
        else {
            message.commissioningModeInitialStepsInstruction = '';
        }
        if (object.commissioningModeSecondaryStepsHint !== undefined && object.commissioningModeSecondaryStepsHint !== null) {
            message.commissioningModeSecondaryStepsHint = object.commissioningModeSecondaryStepsHint;
        }
        else {
            message.commissioningModeSecondaryStepsHint = 0;
        }
        if (object.commissioningModeSecondaryStepsInstruction !== undefined && object.commissioningModeSecondaryStepsInstruction !== null) {
            message.commissioningModeSecondaryStepsInstruction = object.commissioningModeSecondaryStepsInstruction;
        }
        else {
            message.commissioningModeSecondaryStepsInstruction = '';
        }
        if (object.userManualUrl !== undefined && object.userManualUrl !== null) {
            message.userManualUrl = object.userManualUrl;
        }
        else {
            message.userManualUrl = '';
        }
        if (object.supportUrl !== undefined && object.supportUrl !== null) {
            message.supportUrl = object.supportUrl;
        }
        else {
            message.supportUrl = '';
        }
        if (object.productUrl !== undefined && object.productUrl !== null) {
            message.productUrl = object.productUrl;
        }
        else {
            message.productUrl = '';
        }
        if (object.lsfUrl !== undefined && object.lsfUrl !== null) {
            message.lsfUrl = object.lsfUrl;
        }
        else {
            message.lsfUrl = '';
        }
        if (object.lsfRevision !== undefined && object.lsfRevision !== null) {
            message.lsfRevision = object.lsfRevision;
        }
        else {
            message.lsfRevision = 0;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        return message;
    }
};
