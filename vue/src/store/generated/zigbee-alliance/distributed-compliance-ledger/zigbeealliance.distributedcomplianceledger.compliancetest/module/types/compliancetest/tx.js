/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliancetest';
const baseMsgAddTestingResult = { signer: '', vid: 0, pid: 0, softwareVersion: 0, softwareVersionString: '', testResult: '', testDate: '' };
export const MsgAddTestingResult = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.vid !== 0) {
            writer.uint32(16).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(24).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(32).uint32(message.softwareVersion);
        }
        if (message.softwareVersionString !== '') {
            writer.uint32(42).string(message.softwareVersionString);
        }
        if (message.testResult !== '') {
            writer.uint32(50).string(message.testResult);
        }
        if (message.testDate !== '') {
            writer.uint32(58).string(message.testDate);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddTestingResult };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.vid = reader.int32();
                    break;
                case 3:
                    message.pid = reader.int32();
                    break;
                case 4:
                    message.softwareVersion = reader.uint32();
                    break;
                case 5:
                    message.softwareVersionString = reader.string();
                    break;
                case 6:
                    message.testResult = reader.string();
                    break;
                case 7:
                    message.testDate = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAddTestingResult };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
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
        if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
            message.softwareVersionString = String(object.softwareVersionString);
        }
        else {
            message.softwareVersionString = '';
        }
        if (object.testResult !== undefined && object.testResult !== null) {
            message.testResult = String(object.testResult);
        }
        else {
            message.testResult = '';
        }
        if (object.testDate !== undefined && object.testDate !== null) {
            message.testDate = String(object.testDate);
        }
        else {
            message.testDate = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.softwareVersionString !== undefined && (obj.softwareVersionString = message.softwareVersionString);
        message.testResult !== undefined && (obj.testResult = message.testResult);
        message.testDate !== undefined && (obj.testDate = message.testDate);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddTestingResult };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
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
        if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
            message.softwareVersionString = object.softwareVersionString;
        }
        else {
            message.softwareVersionString = '';
        }
        if (object.testResult !== undefined && object.testResult !== null) {
            message.testResult = object.testResult;
        }
        else {
            message.testResult = '';
        }
        if (object.testDate !== undefined && object.testDate !== null) {
            message.testDate = object.testDate;
        }
        else {
            message.testDate = '';
        }
        return message;
    }
};
const baseMsgAddTestingResultResponse = {};
export const MsgAddTestingResultResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddTestingResultResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgAddTestingResultResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddTestingResultResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    AddTestingResult(request) {
        const data = MsgAddTestingResult.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliancetest.Msg', 'AddTestingResult', data);
        return promise.then((data) => MsgAddTestingResultResponse.decode(new Reader(data)));
    }
}
