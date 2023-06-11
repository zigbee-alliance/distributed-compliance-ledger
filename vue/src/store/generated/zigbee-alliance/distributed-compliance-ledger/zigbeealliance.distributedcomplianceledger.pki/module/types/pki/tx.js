/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseMsgProposeAddX509RootCert = { signer: '', cert: '', info: '', time: 0 };
export const MsgProposeAddX509RootCert = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.cert !== '') {
            writer.uint32(18).string(message.cert);
        }
        if (message.info !== '') {
            writer.uint32(26).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(32).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeAddX509RootCert };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.cert = reader.string();
                    break;
                case 3:
                    message.info = reader.string();
                    break;
                case 4:
                    message.time = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgProposeAddX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.cert !== undefined && object.cert !== null) {
            message.cert = String(object.cert);
        }
        else {
            message.cert = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.cert !== undefined && (obj.cert = message.cert);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgProposeAddX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.cert !== undefined && object.cert !== null) {
            message.cert = object.cert;
        }
        else {
            message.cert = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        return message;
    }
};
const baseMsgProposeAddX509RootCertResponse = {};
export const MsgProposeAddX509RootCertResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeAddX509RootCertResponse };
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
        const message = { ...baseMsgProposeAddX509RootCertResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgProposeAddX509RootCertResponse };
        return message;
    }
};
const baseMsgApproveAddX509RootCert = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 };
export const MsgApproveAddX509RootCert = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.subject !== '') {
            writer.uint32(18).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(26).string(message.subjectKeyId);
        }
        if (message.info !== '') {
            writer.uint32(34).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(40).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveAddX509RootCert };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.subject = reader.string();
                    break;
                case 3:
                    message.subjectKeyId = reader.string();
                    break;
                case 4:
                    message.info = reader.string();
                    break;
                case 5:
                    message.time = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgApproveAddX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgApproveAddX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        return message;
    }
};
const baseMsgApproveAddX509RootCertResponse = {};
export const MsgApproveAddX509RootCertResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveAddX509RootCertResponse };
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
        const message = { ...baseMsgApproveAddX509RootCertResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgApproveAddX509RootCertResponse };
        return message;
    }
};
const baseMsgAddX509Cert = { signer: '', cert: '', info: '', time: 0 };
export const MsgAddX509Cert = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.cert !== '') {
            writer.uint32(18).string(message.cert);
        }
        if (message.info !== '') {
            writer.uint32(26).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(32).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddX509Cert };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.cert = reader.string();
                    break;
                case 3:
                    message.info = reader.string();
                    break;
                case 4:
                    message.time = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAddX509Cert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.cert !== undefined && object.cert !== null) {
            message.cert = String(object.cert);
        }
        else {
            message.cert = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.cert !== undefined && (obj.cert = message.cert);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddX509Cert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.cert !== undefined && object.cert !== null) {
            message.cert = object.cert;
        }
        else {
            message.cert = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        return message;
    }
};
const baseMsgAddX509CertResponse = {};
export const MsgAddX509CertResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddX509CertResponse };
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
        const message = { ...baseMsgAddX509CertResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddX509CertResponse };
        return message;
    }
};
const baseMsgProposeRevokeX509RootCert = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 };
export const MsgProposeRevokeX509RootCert = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.subject !== '') {
            writer.uint32(18).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(26).string(message.subjectKeyId);
        }
        if (message.info !== '') {
            writer.uint32(34).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(40).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeRevokeX509RootCert };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.subject = reader.string();
                    break;
                case 3:
                    message.subjectKeyId = reader.string();
                    break;
                case 4:
                    message.info = reader.string();
                    break;
                case 5:
                    message.time = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgProposeRevokeX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgProposeRevokeX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        return message;
    }
};
const baseMsgProposeRevokeX509RootCertResponse = {};
export const MsgProposeRevokeX509RootCertResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeRevokeX509RootCertResponse };
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
        const message = { ...baseMsgProposeRevokeX509RootCertResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgProposeRevokeX509RootCertResponse };
        return message;
    }
};
const baseMsgApproveRevokeX509RootCert = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 };
export const MsgApproveRevokeX509RootCert = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.subject !== '') {
            writer.uint32(18).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(26).string(message.subjectKeyId);
        }
        if (message.info !== '') {
            writer.uint32(42).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(48).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveRevokeX509RootCert };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.subject = reader.string();
                    break;
                case 3:
                    message.subjectKeyId = reader.string();
                    break;
                case 5:
                    message.info = reader.string();
                    break;
                case 6:
                    message.time = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgApproveRevokeX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgApproveRevokeX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        return message;
    }
};
const baseMsgApproveRevokeX509RootCertResponse = {};
export const MsgApproveRevokeX509RootCertResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveRevokeX509RootCertResponse };
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
        const message = { ...baseMsgApproveRevokeX509RootCertResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgApproveRevokeX509RootCertResponse };
        return message;
    }
};
const baseMsgRevokeX509Cert = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 };
export const MsgRevokeX509Cert = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.subject !== '') {
            writer.uint32(18).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(26).string(message.subjectKeyId);
        }
        if (message.info !== '') {
            writer.uint32(34).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(40).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRevokeX509Cert };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.subject = reader.string();
                    break;
                case 3:
                    message.subjectKeyId = reader.string();
                    break;
                case 4:
                    message.info = reader.string();
                    break;
                case 5:
                    message.time = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgRevokeX509Cert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRevokeX509Cert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        return message;
    }
};
const baseMsgRevokeX509CertResponse = {};
export const MsgRevokeX509CertResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRevokeX509CertResponse };
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
        const message = { ...baseMsgRevokeX509CertResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRevokeX509CertResponse };
        return message;
    }
};
const baseMsgRejectAddX509RootCert = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 };
export const MsgRejectAddX509RootCert = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.subject !== '') {
            writer.uint32(18).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(26).string(message.subjectKeyId);
        }
        if (message.info !== '') {
            writer.uint32(34).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(40).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRejectAddX509RootCert };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.subject = reader.string();
                    break;
                case 3:
                    message.subjectKeyId = reader.string();
                    break;
                case 4:
                    message.info = reader.string();
                    break;
                case 5:
                    message.time = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgRejectAddX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRejectAddX509RootCert };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        return message;
    }
};
const baseMsgRejectAddX509RootCertResponse = {};
export const MsgRejectAddX509RootCertResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRejectAddX509RootCertResponse };
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
        const message = { ...baseMsgRejectAddX509RootCertResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRejectAddX509RootCertResponse };
        return message;
    }
};
const baseMsgAddPkiRevocationDistributionPoint = {
    signer: '',
    vid: 0,
    pid: 0,
    isPAA: false,
    label: '',
    crlSignerCertificate: '',
    issuerSubjectKeyID: '',
    dataURL: '',
    dataFileSize: 0,
    dataDigest: '',
    dataDigestType: 0,
    revocationType: 0
};
export const MsgAddPkiRevocationDistributionPoint = {
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
        if (message.isPAA === true) {
            writer.uint32(32).bool(message.isPAA);
        }
        if (message.label !== '') {
            writer.uint32(42).string(message.label);
        }
        if (message.crlSignerCertificate !== '') {
            writer.uint32(50).string(message.crlSignerCertificate);
        }
        if (message.issuerSubjectKeyID !== '') {
            writer.uint32(58).string(message.issuerSubjectKeyID);
        }
        if (message.dataURL !== '') {
            writer.uint32(66).string(message.dataURL);
        }
        if (message.dataFileSize !== 0) {
            writer.uint32(72).uint64(message.dataFileSize);
        }
        if (message.dataDigest !== '') {
            writer.uint32(82).string(message.dataDigest);
        }
        if (message.dataDigestType !== 0) {
            writer.uint32(88).uint32(message.dataDigestType);
        }
        if (message.revocationType !== 0) {
            writer.uint32(96).uint64(message.revocationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddPkiRevocationDistributionPoint };
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
                    message.isPAA = reader.bool();
                    break;
                case 5:
                    message.label = reader.string();
                    break;
                case 6:
                    message.crlSignerCertificate = reader.string();
                    break;
                case 7:
                    message.issuerSubjectKeyID = reader.string();
                    break;
                case 8:
                    message.dataURL = reader.string();
                    break;
                case 9:
                    message.dataFileSize = longToNumber(reader.uint64());
                    break;
                case 10:
                    message.dataDigest = reader.string();
                    break;
                case 11:
                    message.dataDigestType = reader.uint32();
                    break;
                case 12:
                    message.revocationType = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgAddPkiRevocationDistributionPoint };
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
        if (object.isPAA !== undefined && object.isPAA !== null) {
            message.isPAA = Boolean(object.isPAA);
        }
        else {
            message.isPAA = false;
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = '';
        }
        if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
            message.crlSignerCertificate = String(object.crlSignerCertificate);
        }
        else {
            message.crlSignerCertificate = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = String(object.issuerSubjectKeyID);
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        if (object.dataURL !== undefined && object.dataURL !== null) {
            message.dataURL = String(object.dataURL);
        }
        else {
            message.dataURL = '';
        }
        if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
            message.dataFileSize = Number(object.dataFileSize);
        }
        else {
            message.dataFileSize = 0;
        }
        if (object.dataDigest !== undefined && object.dataDigest !== null) {
            message.dataDigest = String(object.dataDigest);
        }
        else {
            message.dataDigest = '';
        }
        if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
            message.dataDigestType = Number(object.dataDigestType);
        }
        else {
            message.dataDigestType = 0;
        }
        if (object.revocationType !== undefined && object.revocationType !== null) {
            message.revocationType = Number(object.revocationType);
        }
        else {
            message.revocationType = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.isPAA !== undefined && (obj.isPAA = message.isPAA);
        message.label !== undefined && (obj.label = message.label);
        message.crlSignerCertificate !== undefined && (obj.crlSignerCertificate = message.crlSignerCertificate);
        message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID);
        message.dataURL !== undefined && (obj.dataURL = message.dataURL);
        message.dataFileSize !== undefined && (obj.dataFileSize = message.dataFileSize);
        message.dataDigest !== undefined && (obj.dataDigest = message.dataDigest);
        message.dataDigestType !== undefined && (obj.dataDigestType = message.dataDigestType);
        message.revocationType !== undefined && (obj.revocationType = message.revocationType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgAddPkiRevocationDistributionPoint };
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
        if (object.isPAA !== undefined && object.isPAA !== null) {
            message.isPAA = object.isPAA;
        }
        else {
            message.isPAA = false;
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = '';
        }
        if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
            message.crlSignerCertificate = object.crlSignerCertificate;
        }
        else {
            message.crlSignerCertificate = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = object.issuerSubjectKeyID;
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        if (object.dataURL !== undefined && object.dataURL !== null) {
            message.dataURL = object.dataURL;
        }
        else {
            message.dataURL = '';
        }
        if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
            message.dataFileSize = object.dataFileSize;
        }
        else {
            message.dataFileSize = 0;
        }
        if (object.dataDigest !== undefined && object.dataDigest !== null) {
            message.dataDigest = object.dataDigest;
        }
        else {
            message.dataDigest = '';
        }
        if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
            message.dataDigestType = object.dataDigestType;
        }
        else {
            message.dataDigestType = 0;
        }
        if (object.revocationType !== undefined && object.revocationType !== null) {
            message.revocationType = object.revocationType;
        }
        else {
            message.revocationType = 0;
        }
        return message;
    }
};
const baseMsgAddPkiRevocationDistributionPointResponse = {};
export const MsgAddPkiRevocationDistributionPointResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgAddPkiRevocationDistributionPointResponse };
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
        const message = { ...baseMsgAddPkiRevocationDistributionPointResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgAddPkiRevocationDistributionPointResponse };
        return message;
    }
};
const baseMsgUpdatePkiRevocationDistributionPoint = {
    signer: '',
    vid: 0,
    label: '',
    crlSignerCertificate: '',
    issuerSubjectKeyID: '',
    dataURL: '',
    dataFileSize: 0,
    dataDigest: '',
    dataDigestType: 0
};
export const MsgUpdatePkiRevocationDistributionPoint = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.vid !== 0) {
            writer.uint32(16).int32(message.vid);
        }
        if (message.label !== '') {
            writer.uint32(26).string(message.label);
        }
        if (message.crlSignerCertificate !== '') {
            writer.uint32(34).string(message.crlSignerCertificate);
        }
        if (message.issuerSubjectKeyID !== '') {
            writer.uint32(42).string(message.issuerSubjectKeyID);
        }
        if (message.dataURL !== '') {
            writer.uint32(50).string(message.dataURL);
        }
        if (message.dataFileSize !== 0) {
            writer.uint32(56).uint64(message.dataFileSize);
        }
        if (message.dataDigest !== '') {
            writer.uint32(66).string(message.dataDigest);
        }
        if (message.dataDigestType !== 0) {
            writer.uint32(72).uint32(message.dataDigestType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdatePkiRevocationDistributionPoint };
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
                    message.label = reader.string();
                    break;
                case 4:
                    message.crlSignerCertificate = reader.string();
                    break;
                case 5:
                    message.issuerSubjectKeyID = reader.string();
                    break;
                case 6:
                    message.dataURL = reader.string();
                    break;
                case 7:
                    message.dataFileSize = longToNumber(reader.uint64());
                    break;
                case 8:
                    message.dataDigest = reader.string();
                    break;
                case 9:
                    message.dataDigestType = reader.uint32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdatePkiRevocationDistributionPoint };
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
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = '';
        }
        if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
            message.crlSignerCertificate = String(object.crlSignerCertificate);
        }
        else {
            message.crlSignerCertificate = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = String(object.issuerSubjectKeyID);
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        if (object.dataURL !== undefined && object.dataURL !== null) {
            message.dataURL = String(object.dataURL);
        }
        else {
            message.dataURL = '';
        }
        if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
            message.dataFileSize = Number(object.dataFileSize);
        }
        else {
            message.dataFileSize = 0;
        }
        if (object.dataDigest !== undefined && object.dataDigest !== null) {
            message.dataDigest = String(object.dataDigest);
        }
        else {
            message.dataDigest = '';
        }
        if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
            message.dataDigestType = Number(object.dataDigestType);
        }
        else {
            message.dataDigestType = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.vid !== undefined && (obj.vid = message.vid);
        message.label !== undefined && (obj.label = message.label);
        message.crlSignerCertificate !== undefined && (obj.crlSignerCertificate = message.crlSignerCertificate);
        message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID);
        message.dataURL !== undefined && (obj.dataURL = message.dataURL);
        message.dataFileSize !== undefined && (obj.dataFileSize = message.dataFileSize);
        message.dataDigest !== undefined && (obj.dataDigest = message.dataDigest);
        message.dataDigestType !== undefined && (obj.dataDigestType = message.dataDigestType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdatePkiRevocationDistributionPoint };
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
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = '';
        }
        if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
            message.crlSignerCertificate = object.crlSignerCertificate;
        }
        else {
            message.crlSignerCertificate = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = object.issuerSubjectKeyID;
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        if (object.dataURL !== undefined && object.dataURL !== null) {
            message.dataURL = object.dataURL;
        }
        else {
            message.dataURL = '';
        }
        if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
            message.dataFileSize = object.dataFileSize;
        }
        else {
            message.dataFileSize = 0;
        }
        if (object.dataDigest !== undefined && object.dataDigest !== null) {
            message.dataDigest = object.dataDigest;
        }
        else {
            message.dataDigest = '';
        }
        if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
            message.dataDigestType = object.dataDigestType;
        }
        else {
            message.dataDigestType = 0;
        }
        return message;
    }
};
const baseMsgUpdatePkiRevocationDistributionPointResponse = {};
export const MsgUpdatePkiRevocationDistributionPointResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdatePkiRevocationDistributionPointResponse };
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
        const message = { ...baseMsgUpdatePkiRevocationDistributionPointResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUpdatePkiRevocationDistributionPointResponse };
        return message;
    }
};
const baseMsgDeletePkiRevocationDistributionPoint = { signer: '', vid: 0, label: '', issuerSubjectKeyID: '' };
export const MsgDeletePkiRevocationDistributionPoint = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.vid !== 0) {
            writer.uint32(16).int32(message.vid);
        }
        if (message.label !== '') {
            writer.uint32(26).string(message.label);
        }
        if (message.issuerSubjectKeyID !== '') {
            writer.uint32(34).string(message.issuerSubjectKeyID);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeletePkiRevocationDistributionPoint };
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
                    message.label = reader.string();
                    break;
                case 4:
                    message.issuerSubjectKeyID = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgDeletePkiRevocationDistributionPoint };
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
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = String(object.issuerSubjectKeyID);
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.vid !== undefined && (obj.vid = message.vid);
        message.label !== undefined && (obj.label = message.label);
        message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeletePkiRevocationDistributionPoint };
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
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = object.issuerSubjectKeyID;
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        return message;
    }
};
const baseMsgDeletePkiRevocationDistributionPointResponse = {};
export const MsgDeletePkiRevocationDistributionPointResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeletePkiRevocationDistributionPointResponse };
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
        const message = { ...baseMsgDeletePkiRevocationDistributionPointResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDeletePkiRevocationDistributionPointResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ProposeAddX509RootCert(request) {
        const data = MsgProposeAddX509RootCert.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ProposeAddX509RootCert', data);
        return promise.then((data) => MsgProposeAddX509RootCertResponse.decode(new Reader(data)));
    }
    ApproveAddX509RootCert(request) {
        const data = MsgApproveAddX509RootCert.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ApproveAddX509RootCert', data);
        return promise.then((data) => MsgApproveAddX509RootCertResponse.decode(new Reader(data)));
    }
    AddX509Cert(request) {
        const data = MsgAddX509Cert.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'AddX509Cert', data);
        return promise.then((data) => MsgAddX509CertResponse.decode(new Reader(data)));
    }
    ProposeRevokeX509RootCert(request) {
        const data = MsgProposeRevokeX509RootCert.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ProposeRevokeX509RootCert', data);
        return promise.then((data) => MsgProposeRevokeX509RootCertResponse.decode(new Reader(data)));
    }
    ApproveRevokeX509RootCert(request) {
        const data = MsgApproveRevokeX509RootCert.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ApproveRevokeX509RootCert', data);
        return promise.then((data) => MsgApproveRevokeX509RootCertResponse.decode(new Reader(data)));
    }
    RevokeX509Cert(request) {
        const data = MsgRevokeX509Cert.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'RevokeX509Cert', data);
        return promise.then((data) => MsgRevokeX509CertResponse.decode(new Reader(data)));
    }
    RejectAddX509RootCert(request) {
        const data = MsgRejectAddX509RootCert.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'RejectAddX509RootCert', data);
        return promise.then((data) => MsgRejectAddX509RootCertResponse.decode(new Reader(data)));
    }
    AddPkiRevocationDistributionPoint(request) {
        const data = MsgAddPkiRevocationDistributionPoint.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'AddPkiRevocationDistributionPoint', data);
        return promise.then((data) => MsgAddPkiRevocationDistributionPointResponse.decode(new Reader(data)));
    }
    UpdatePkiRevocationDistributionPoint(request) {
        const data = MsgUpdatePkiRevocationDistributionPoint.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'UpdatePkiRevocationDistributionPoint', data);
        return promise.then((data) => MsgUpdatePkiRevocationDistributionPointResponse.decode(new Reader(data)));
    }
    DeletePkiRevocationDistributionPoint(request) {
        const data = MsgDeletePkiRevocationDistributionPoint.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'DeletePkiRevocationDistributionPoint', data);
        return promise.then((data) => MsgDeletePkiRevocationDistributionPointResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
