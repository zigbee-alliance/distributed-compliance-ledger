/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseGrant = { address: '', time: 0, info: '' };
export const Grant = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.time !== 0) {
            writer.uint32(16).int64(message.time);
        }
        if (message.info !== '') {
            writer.uint32(26).string(message.info);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGrant };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.time = longToNumber(reader.int64());
                    break;
                case 3:
                    message.info = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGrant };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = Number(object.time);
        }
        else {
            message.time = 0;
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = String(object.info);
        }
        else {
            message.info = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.time !== undefined && (obj.time = message.time);
        message.info !== undefined && (obj.info = message.info);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGrant };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.time !== undefined && object.time !== null) {
            message.time = object.time;
        }
        else {
            message.time = 0;
        }
        if (object.info !== undefined && object.info !== null) {
            message.info = object.info;
        }
        else {
            message.info = '';
        }
        return message;
    }
};
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
