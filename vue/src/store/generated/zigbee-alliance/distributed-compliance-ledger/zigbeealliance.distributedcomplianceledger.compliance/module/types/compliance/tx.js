/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseMsgCertifyModel = {
    signer: '',
    vid: 0,
    pid: 0,
    softwareVersion: 0,
    softwareVersionString: '',
    certificationDate: '',
    certificationType: '',
    reason: ''
};
export const MsgCertifyModel = {
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
        if (message.certificationDate !== '') {
            writer.uint32(50).string(message.certificationDate);
        }
        if (message.certificationType !== '') {
            writer.uint32(58).string(message.certificationType);
        }
        if (message.reason !== '') {
            writer.uint32(66).string(message.reason);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCertifyModel };
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
                    message.certificationDate = reader.string();
                    break;
                case 7:
                    message.certificationType = reader.string();
                    break;
                case 8:
                    message.reason = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCertifyModel };
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
        if (object.certificationDate !== undefined && object.certificationDate !== null) {
            message.certificationDate = String(object.certificationDate);
        }
        else {
            message.certificationDate = '';
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = String(object.reason);
        }
        else {
            message.reason = '';
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
        message.certificationDate !== undefined && (obj.certificationDate = message.certificationDate);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        message.reason !== undefined && (obj.reason = message.reason);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCertifyModel };
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
        if (object.certificationDate !== undefined && object.certificationDate !== null) {
            message.certificationDate = object.certificationDate;
        }
        else {
            message.certificationDate = '';
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = object.reason;
        }
        else {
            message.reason = '';
        }
        return message;
    }
};
const baseMsgCertifyModelResponse = {};
export const MsgCertifyModelResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCertifyModelResponse };
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
        const message = { ...baseMsgCertifyModelResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgCertifyModelResponse };
        return message;
    }
};
const baseMsgRevokeModel = {
    signer: '',
    vid: 0,
    pid: 0,
    softwareVersion: 0,
    softwareVersionString: '',
    revocationDate: '',
    certificationType: '',
    reason: ''
};
export const MsgRevokeModel = {
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
        if (message.revocationDate !== '') {
            writer.uint32(50).string(message.revocationDate);
        }
        if (message.certificationType !== '') {
            writer.uint32(58).string(message.certificationType);
        }
        if (message.reason !== '') {
            writer.uint32(66).string(message.reason);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRevokeModel };
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
                    message.revocationDate = reader.string();
                    break;
                case 7:
                    message.certificationType = reader.string();
                    break;
                case 8:
                    message.reason = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgRevokeModel };
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
        if (object.revocationDate !== undefined && object.revocationDate !== null) {
            message.revocationDate = String(object.revocationDate);
        }
        else {
            message.revocationDate = '';
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = String(object.reason);
        }
        else {
            message.reason = '';
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
        message.revocationDate !== undefined && (obj.revocationDate = message.revocationDate);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        message.reason !== undefined && (obj.reason = message.reason);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRevokeModel };
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
        if (object.revocationDate !== undefined && object.revocationDate !== null) {
            message.revocationDate = object.revocationDate;
        }
        else {
            message.revocationDate = '';
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = object.reason;
        }
        else {
            message.reason = '';
        }
        return message;
    }
};
const baseMsgRevokeModelResponse = {};
export const MsgRevokeModelResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRevokeModelResponse };
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
        const message = { ...baseMsgRevokeModelResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRevokeModelResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CertifyModel(request) {
        const data = MsgCertifyModel.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Msg', 'CertifyModel', data);
        return promise.then((data) => MsgCertifyModelResponse.decode(new Reader(data)));
    }
    RevokeModel(request) {
        const data = MsgRevokeModel.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Msg', 'RevokeModel', data);
        return promise.then((data) => MsgRevokeModelResponse.decode(new Reader(data)));
    }
}
