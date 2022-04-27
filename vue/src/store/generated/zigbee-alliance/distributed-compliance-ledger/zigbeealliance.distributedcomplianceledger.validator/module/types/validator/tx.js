/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { Any } from '../google/protobuf/any';
import { Description } from '../validator/description';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseMsgCreateValidator = { signer: '' };
export const MsgCreateValidator = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.pubKey !== undefined) {
            Any.encode(message.pubKey, writer.uint32(18).fork()).ldelim();
        }
        if (message.description !== undefined) {
            Description.encode(message.description, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateValidator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.signer = reader.string();
                    break;
                case 2:
                    message.pubKey = Any.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.description = Description.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateValidator };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = String(object.signer);
        }
        else {
            message.signer = '';
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = Any.fromJSON(object.pubKey);
        }
        else {
            message.pubKey = undefined;
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = Description.fromJSON(object.description);
        }
        else {
            message.description = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.signer !== undefined && (obj.signer = message.signer);
        message.pubKey !== undefined && (obj.pubKey = message.pubKey ? Any.toJSON(message.pubKey) : undefined);
        message.description !== undefined && (obj.description = message.description ? Description.toJSON(message.description) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateValidator };
        if (object.signer !== undefined && object.signer !== null) {
            message.signer = object.signer;
        }
        else {
            message.signer = '';
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = Any.fromPartial(object.pubKey);
        }
        else {
            message.pubKey = undefined;
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = Description.fromPartial(object.description);
        }
        else {
            message.description = undefined;
        }
        return message;
    }
};
const baseMsgCreateValidatorResponse = {};
export const MsgCreateValidatorResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateValidatorResponse };
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
        const message = { ...baseMsgCreateValidatorResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgCreateValidatorResponse };
        return message;
    }
};
const baseMsgProposeDisableValidator = { creator: '', address: '', info: '', time: 0 };
export const MsgProposeDisableValidator = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
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
        const message = { ...baseMsgProposeDisableValidator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
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
        const message = { ...baseMsgProposeDisableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.address !== undefined && (obj.address = message.address);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgProposeDisableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
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
const baseMsgProposeDisableValidatorResponse = {};
export const MsgProposeDisableValidatorResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeDisableValidatorResponse };
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
        const message = { ...baseMsgProposeDisableValidatorResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgProposeDisableValidatorResponse };
        return message;
    }
};
const baseMsgApproveDisableValidator = { creator: '', address: '', info: '', time: 0 };
export const MsgApproveDisableValidator = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
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
        const message = { ...baseMsgApproveDisableValidator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
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
        const message = { ...baseMsgApproveDisableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.address !== undefined && (obj.address = message.address);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgApproveDisableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
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
const baseMsgApproveDisableValidatorResponse = {};
export const MsgApproveDisableValidatorResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveDisableValidatorResponse };
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
        const message = { ...baseMsgApproveDisableValidatorResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgApproveDisableValidatorResponse };
        return message;
    }
};
const baseMsgDisableValidator = { creator: '' };
export const MsgDisableValidator = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDisableValidator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseMsgDisableValidator };
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
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDisableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        return message;
    }
};
const baseMsgDisableValidatorResponse = {};
export const MsgDisableValidatorResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDisableValidatorResponse };
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
        const message = { ...baseMsgDisableValidatorResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDisableValidatorResponse };
        return message;
    }
};
const baseMsgEnableValidator = { creator: '' };
export const MsgEnableValidator = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgEnableValidator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseMsgEnableValidator };
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
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgEnableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        return message;
    }
};
const baseMsgEnableValidatorResponse = {};
export const MsgEnableValidatorResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgEnableValidatorResponse };
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
        const message = { ...baseMsgEnableValidatorResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgEnableValidatorResponse };
        return message;
    }
};
const baseMsgRejectDisableValidator = { creator: '', address: '', info: '', time: 0 };
export const MsgRejectDisableValidator = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
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
        const message = { ...baseMsgRejectDisableValidator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
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
        const message = { ...baseMsgRejectDisableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
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
        message.creator !== undefined && (obj.creator = message.creator);
        message.address !== undefined && (obj.address = message.address);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRejectDisableValidator };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
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
const baseMsgRejectDisableValidatorResponse = {};
export const MsgRejectDisableValidatorResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRejectDisableValidatorResponse };
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
        const message = { ...baseMsgRejectDisableValidatorResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRejectDisableValidatorResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateValidator(request) {
        const data = MsgCreateValidator.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'CreateValidator', data);
        return promise.then((data) => MsgCreateValidatorResponse.decode(new Reader(data)));
    }
    ProposeDisableValidator(request) {
        const data = MsgProposeDisableValidator.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'ProposeDisableValidator', data);
        return promise.then((data) => MsgProposeDisableValidatorResponse.decode(new Reader(data)));
    }
    ApproveDisableValidator(request) {
        const data = MsgApproveDisableValidator.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'ApproveDisableValidator', data);
        return promise.then((data) => MsgApproveDisableValidatorResponse.decode(new Reader(data)));
    }
    DisableValidator(request) {
        const data = MsgDisableValidator.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'DisableValidator', data);
        return promise.then((data) => MsgDisableValidatorResponse.decode(new Reader(data)));
    }
    EnableValidator(request) {
        const data = MsgEnableValidator.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'EnableValidator', data);
        return promise.then((data) => MsgEnableValidatorResponse.decode(new Reader(data)));
    }
    RejectDisableValidator(request) {
        const data = MsgRejectDisableValidator.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'RejectDisableValidator', data);
        return promise.then((data) => MsgRejectDisableValidatorResponse.decode(new Reader(data)));
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
