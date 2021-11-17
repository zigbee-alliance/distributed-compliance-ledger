/* eslint-disable */
import { Description } from '../validator/description';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseValidator = { address: '', pubKey: '', power: 0, jailed: false, jailedReason: '', owner: '' };
export const Validator = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.description !== undefined) {
            Description.encode(message.description, writer.uint32(18).fork()).ldelim();
        }
        if (message.pubKey !== '') {
            writer.uint32(26).string(message.pubKey);
        }
        if (message.power !== 0) {
            writer.uint32(32).int32(message.power);
        }
        if (message.jailed === true) {
            writer.uint32(40).bool(message.jailed);
        }
        if (message.jailedReason !== '') {
            writer.uint32(50).string(message.jailedReason);
        }
        if (message.owner !== '') {
            writer.uint32(58).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseValidator };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.description = Description.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.pubKey = reader.string();
                    break;
                case 4:
                    message.power = reader.int32();
                    break;
                case 5:
                    message.jailed = reader.bool();
                    break;
                case 6:
                    message.jailedReason = reader.string();
                    break;
                case 7:
                    message.owner = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseValidator };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = Description.fromJSON(object.description);
        }
        else {
            message.description = undefined;
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = String(object.pubKey);
        }
        else {
            message.pubKey = '';
        }
        if (object.power !== undefined && object.power !== null) {
            message.power = Number(object.power);
        }
        else {
            message.power = 0;
        }
        if (object.jailed !== undefined && object.jailed !== null) {
            message.jailed = Boolean(object.jailed);
        }
        else {
            message.jailed = false;
        }
        if (object.jailedReason !== undefined && object.jailedReason !== null) {
            message.jailedReason = String(object.jailedReason);
        }
        else {
            message.jailedReason = '';
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.description !== undefined && (obj.description = message.description ? Description.toJSON(message.description) : undefined);
        message.pubKey !== undefined && (obj.pubKey = message.pubKey);
        message.power !== undefined && (obj.power = message.power);
        message.jailed !== undefined && (obj.jailed = message.jailed);
        message.jailedReason !== undefined && (obj.jailedReason = message.jailedReason);
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseValidator };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.description !== undefined && object.description !== null) {
            message.description = Description.fromPartial(object.description);
        }
        else {
            message.description = undefined;
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = object.pubKey;
        }
        else {
            message.pubKey = '';
        }
        if (object.power !== undefined && object.power !== null) {
            message.power = object.power;
        }
        else {
            message.power = 0;
        }
        if (object.jailed !== undefined && object.jailed !== null) {
            message.jailed = object.jailed;
        }
        else {
            message.jailed = false;
        }
        if (object.jailedReason !== undefined && object.jailedReason !== null) {
            message.jailedReason = object.jailedReason;
        }
        else {
            message.jailedReason = '';
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = '';
        }
        return message;
    }
};
