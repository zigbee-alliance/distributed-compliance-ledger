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
