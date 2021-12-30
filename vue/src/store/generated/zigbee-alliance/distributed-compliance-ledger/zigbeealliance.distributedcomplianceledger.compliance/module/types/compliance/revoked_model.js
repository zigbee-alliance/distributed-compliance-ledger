/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseRevokedModel = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '', value: false };
export const RevokedModel = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint32(message.softwareVersion);
        }
        if (message.certificationType !== '') {
            writer.uint32(34).string(message.certificationType);
        }
        if (message.value === true) {
            writer.uint32(40).bool(message.value);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRevokedModel };
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
                    message.softwareVersion = reader.uint32();
                    break;
                case 4:
                    message.certificationType = reader.string();
                    break;
                case 5:
                    message.value = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseRevokedModel };
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
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = Boolean(object.value);
        }
        else {
            message.value = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        message.value !== undefined && (obj.value = message.value);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRevokedModel };
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
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        if (object.value !== undefined && object.value !== null) {
            message.value = object.value;
        }
        else {
            message.value = false;
        }
        return message;
    }
};
