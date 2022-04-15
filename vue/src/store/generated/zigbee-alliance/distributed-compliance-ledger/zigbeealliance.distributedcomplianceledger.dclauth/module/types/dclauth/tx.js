/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { Any } from '../google/protobuf/any';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
const baseMsgProposeAddAccount = { signer: '', address: '', roles: '', vendorID: 0, info: '', time: 0 };
export const MsgProposeAddAccount = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        if (message.pubKey !== undefined) {
            Any.encode(message.pubKey, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.roles) {
            writer.uint32(34).string(v);
        }
        if (message.vendorID !== 0) {
            writer.uint32(40).int32(message.vendorID);
        }
        if (message.info !== '') {
            writer.uint32(50).string(message.info);
        }
        if (message.time !== 0) {
            writer.uint32(56).int64(message.time);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeAddAccount };
        message.roles = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                case 3:
                    message.pubKey = Any.decode(reader, reader.uint32());
                    break;
                case 4:
                    message.roles.push(reader.string());
                    break;
                case 5:
                    message.vendorID = reader.int32();
                    break;
                case 6:
                    message.info = reader.string();
                    break;
                case 7:
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
        const message = { ...baseMsgProposeAddAccount };
        message.roles = [];
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = Any.fromJSON(object.pubKey);
        }
        else {
            message.pubKey = undefined;
        }
        if (object.roles !== undefined && object.roles !== null) {
            for (const e of object.roles) {
                message.roles.push(String(e));
            }
        }
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = Number(object.vendorID);
        }
        else {
            message.vendorID = 0;
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
        message.address !== undefined && (obj.address = message.address);
        message.pubKey !== undefined && (obj.pubKey = message.pubKey ? Any.toJSON(message.pubKey) : undefined);
        if (message.roles) {
            obj.roles = message.roles.map((e) => e);
        }
        else {
            obj.roles = [];
        }
        message.vendorID !== undefined && (obj.vendorID = message.vendorID);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgProposeAddAccount };
        message.roles = [];
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = Any.fromPartial(object.pubKey);
        }
        else {
            message.pubKey = undefined;
        }
        if (object.roles !== undefined && object.roles !== null) {
            for (const e of object.roles) {
                message.roles.push(e);
            }
        }
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = object.vendorID;
        }
        else {
            message.vendorID = 0;
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
const baseMsgProposeAddAccountResponse = {};
export const MsgProposeAddAccountResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeAddAccountResponse };
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
        const message = { ...baseMsgProposeAddAccountResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgProposeAddAccountResponse };
        return message;
    }
};
const baseMsgApproveAddAccount = { signer: '', address: '', info: '', time: 0 };
export const MsgApproveAddAccount = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
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
        const message = { ...baseMsgApproveAddAccount };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
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
        const message = { ...baseMsgApproveAddAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
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
        message.address !== undefined && (obj.address = message.address);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgApproveAddAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
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
const baseMsgApproveAddAccountResponse = {};
export const MsgApproveAddAccountResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveAddAccountResponse };
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
        const message = { ...baseMsgApproveAddAccountResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgApproveAddAccountResponse };
        return message;
    }
};
const baseMsgProposeRevokeAccount = { signer: '', address: '', info: '', time: 0 };
export const MsgProposeRevokeAccount = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
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
        const message = { ...baseMsgProposeRevokeAccount };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
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
        const message = { ...baseMsgProposeRevokeAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
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
        message.address !== undefined && (obj.address = message.address);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgProposeRevokeAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
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
const baseMsgProposeRevokeAccountResponse = {};
export const MsgProposeRevokeAccountResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeRevokeAccountResponse };
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
        const message = { ...baseMsgProposeRevokeAccountResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgProposeRevokeAccountResponse };
        return message;
    }
};
const baseMsgApproveRevokeAccount = { signer: '', address: '', info: '', time: 0 };
export const MsgApproveRevokeAccount = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
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
        const message = { ...baseMsgApproveRevokeAccount };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
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
        const message = { ...baseMsgApproveRevokeAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
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
        message.address !== undefined && (obj.address = message.address);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgApproveRevokeAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
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
const baseMsgApproveRevokeAccountResponse = {};
export const MsgApproveRevokeAccountResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveRevokeAccountResponse };
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
        const message = { ...baseMsgApproveRevokeAccountResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgApproveRevokeAccountResponse };
        return message;
    }
};
const baseMsgRejectAddAccount = { signer: '', address: '', info: '', time: 0 };
export const MsgRejectAddAccount = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
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
        const message = { ...baseMsgRejectAddAccount };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.address = reader.string();
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
        const message = { ...baseMsgRejectAddAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
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
        message.address !== undefined && (obj.address = message.address);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRejectAddAccount };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
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
const baseMsgRejectAddAccountResponse = {};
export const MsgRejectAddAccountResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRejectAddAccountResponse };
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
        const message = { ...baseMsgRejectAddAccountResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRejectAddAccountResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ProposeAddAccount(request) {
        const data = MsgProposeAddAccount.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ProposeAddAccount', data);
        return promise.then((data) => MsgProposeAddAccountResponse.decode(new Reader(data)));
    }
    ApproveAddAccount(request) {
        const data = MsgApproveAddAccount.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ApproveAddAccount', data);
        return promise.then((data) => MsgApproveAddAccountResponse.decode(new Reader(data)));
    }
    ProposeRevokeAccount(request) {
        const data = MsgProposeRevokeAccount.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ProposeRevokeAccount', data);
        return promise.then((data) => MsgProposeRevokeAccountResponse.decode(new Reader(data)));
    }
    ApproveRevokeAccount(request) {
        const data = MsgApproveRevokeAccount.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ApproveRevokeAccount', data);
        return promise.then((data) => MsgApproveRevokeAccountResponse.decode(new Reader(data)));
    }
    RejectAddAccount(request) {
        const data = MsgRejectAddAccount.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'RejectAddAccount', data);
        return promise.then((data) => MsgRejectAddAccountResponse.decode(new Reader(data)));
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
