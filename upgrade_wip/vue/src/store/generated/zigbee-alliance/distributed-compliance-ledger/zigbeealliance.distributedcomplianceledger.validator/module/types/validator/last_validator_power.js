/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseLastValidatorPower = { consensusAddress: '', power: 0 };
export const LastValidatorPower = {
    encode(message, writer = Writer.create()) {
        if (message.consensusAddress !== '') {
            writer.uint32(10).string(message.consensusAddress);
        }
        if (message.power !== 0) {
            writer.uint32(16).int32(message.power);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseLastValidatorPower };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.consensusAddress = reader.string();
                    break;
                case 2:
                    message.power = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseLastValidatorPower };
        if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
            message.consensusAddress = String(object.consensusAddress);
        }
        else {
            message.consensusAddress = '';
        }
        if (object.power !== undefined && object.power !== null) {
            message.power = Number(object.power);
        }
        else {
            message.power = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.consensusAddress !== undefined && (obj.consensusAddress = message.consensusAddress);
        message.power !== undefined && (obj.power = message.power);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseLastValidatorPower };
        if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
            message.consensusAddress = object.consensusAddress;
        }
        else {
            message.consensusAddress = '';
        }
        if (object.power !== undefined && object.power !== null) {
            message.power = object.power;
        }
        else {
            message.power = 0;
        }
        return message;
    }
};
