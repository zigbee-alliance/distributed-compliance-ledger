/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseValidatorSigningInfo = { address: '', startHeight: 0, indexOffset: 0, missedBlocksCounter: 0 };
export const ValidatorSigningInfo = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.startHeight !== 0) {
            writer.uint32(16).uint64(message.startHeight);
        }
        if (message.indexOffset !== 0) {
            writer.uint32(24).uint64(message.indexOffset);
        }
        if (message.missedBlocksCounter !== 0) {
            writer.uint32(32).uint64(message.missedBlocksCounter);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseValidatorSigningInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.startHeight = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.indexOffset = longToNumber(reader.uint64());
                    break;
                case 4:
                    message.missedBlocksCounter = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseValidatorSigningInfo };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.startHeight !== undefined && object.startHeight !== null) {
            message.startHeight = Number(object.startHeight);
        }
        else {
            message.startHeight = 0;
        }
        if (object.indexOffset !== undefined && object.indexOffset !== null) {
            message.indexOffset = Number(object.indexOffset);
        }
        else {
            message.indexOffset = 0;
        }
        if (object.missedBlocksCounter !== undefined && object.missedBlocksCounter !== null) {
            message.missedBlocksCounter = Number(object.missedBlocksCounter);
        }
        else {
            message.missedBlocksCounter = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.startHeight !== undefined && (obj.startHeight = message.startHeight);
        message.indexOffset !== undefined && (obj.indexOffset = message.indexOffset);
        message.missedBlocksCounter !== undefined && (obj.missedBlocksCounter = message.missedBlocksCounter);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseValidatorSigningInfo };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.startHeight !== undefined && object.startHeight !== null) {
            message.startHeight = object.startHeight;
        }
        else {
            message.startHeight = 0;
        }
        if (object.indexOffset !== undefined && object.indexOffset !== null) {
            message.indexOffset = object.indexOffset;
        }
        else {
            message.indexOffset = 0;
        }
        if (object.missedBlocksCounter !== undefined && object.missedBlocksCounter !== null) {
            message.missedBlocksCounter = object.missedBlocksCounter;
        }
        else {
            message.missedBlocksCounter = 0;
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
