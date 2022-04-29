/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { Plan } from '../cosmos/upgrade/v1beta1/upgrade';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade';
const baseMsgProposeUpgrade = { creator: '', info: '', time: 0 };
export const MsgProposeUpgrade = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.plan !== undefined) {
            Plan.encode(message.plan, writer.uint32(18).fork()).ldelim();
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
        const message = { ...baseMsgProposeUpgrade };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.plan = Plan.decode(reader, reader.uint32());
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
        const message = { ...baseMsgProposeUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.plan !== undefined && object.plan !== null) {
            message.plan = Plan.fromJSON(object.plan);
        }
        else {
            message.plan = undefined;
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
        message.plan !== undefined && (obj.plan = message.plan ? Plan.toJSON(message.plan) : undefined);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgProposeUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.plan !== undefined && object.plan !== null) {
            message.plan = Plan.fromPartial(object.plan);
        }
        else {
            message.plan = undefined;
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
const baseMsgProposeUpgradeResponse = {};
export const MsgProposeUpgradeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeUpgradeResponse };
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
        const message = { ...baseMsgProposeUpgradeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgProposeUpgradeResponse };
        return message;
    }
};
const baseMsgApproveUpgrade = { creator: '', name: '', info: '', time: 0 };
export const MsgApproveUpgrade = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.name !== '') {
            writer.uint32(18).string(message.name);
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
        const message = { ...baseMsgApproveUpgrade };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
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
        const message = { ...baseMsgApproveUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
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
        message.name !== undefined && (obj.name = message.name);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgApproveUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
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
const baseMsgApproveUpgradeResponse = {};
export const MsgApproveUpgradeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgApproveUpgradeResponse };
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
        const message = { ...baseMsgApproveUpgradeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgApproveUpgradeResponse };
        return message;
    }
};
const baseMsgRejectUpgrade = { creator: '', name: '', info: '', time: 0 };
export const MsgRejectUpgrade = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.name !== '') {
            writer.uint32(18).string(message.name);
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
        const message = { ...baseMsgRejectUpgrade };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.name = reader.string();
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
        const message = { ...baseMsgRejectUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
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
        message.name !== undefined && (obj.name = message.name);
        message.info !== undefined && (obj.info = message.info);
        message.time !== undefined && (obj.time = message.time);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgRejectUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
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
const baseMsgRejectUpgradeResponse = {};
export const MsgRejectUpgradeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgRejectUpgradeResponse };
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
        const message = { ...baseMsgRejectUpgradeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgRejectUpgradeResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ProposeUpgrade(request) {
        const data = MsgProposeUpgrade.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'ProposeUpgrade', data);
        return promise.then((data) => MsgProposeUpgradeResponse.decode(new Reader(data)));
    }
    ApproveUpgrade(request) {
        const data = MsgApproveUpgrade.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'ApproveUpgrade', data);
        return promise.then((data) => MsgApproveUpgradeResponse.decode(new Reader(data)));
    }
    RejectUpgrade(request) {
        const data = MsgRejectUpgrade.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'RejectUpgrade', data);
        return promise.then((data) => MsgRejectUpgradeResponse.decode(new Reader(data)));
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
