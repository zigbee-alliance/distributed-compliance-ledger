/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { Validator } from '../validator/validator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { LastValidatorPower } from '../validator/last_validator_power';
import { ValidatorSigningInfo } from '../validator/validator_signing_info';
import { ValidatorMissedBlockBitArray } from '../validator/validator_missed_block_bit_array';
import { ValidatorOwner } from '../validator/validator_owner';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseQueryGetValidatorRequest = { address: '' };
export const QueryGetValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetValidatorResponse = {};
export const QueryGetValidatorResponse = {
    encode(message, writer = Writer.create()) {
        if (message.validator !== undefined) {
            Validator.encode(message.validator, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validator = Validator.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorResponse };
        if (object.validator !== undefined && object.validator !== null) {
            message.validator = Validator.fromJSON(object.validator);
        }
        else {
            message.validator = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.validator !== undefined && (obj.validator = message.validator ? Validator.toJSON(message.validator) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorResponse };
        if (object.validator !== undefined && object.validator !== null) {
            message.validator = Validator.fromPartial(object.validator);
        }
        else {
            message.validator = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorRequest = {};
export const QueryAllValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorResponse = {};
export const QueryAllValidatorResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.validator) {
            Validator.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorResponse };
        message.validator = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validator.push(Validator.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorResponse };
        message.validator = [];
        if (object.validator !== undefined && object.validator !== null) {
            for (const e of object.validator) {
                message.validator.push(Validator.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.validator) {
            obj.validator = message.validator.map((e) => (e ? Validator.toJSON(e) : undefined));
        }
        else {
            obj.validator = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorResponse };
        message.validator = [];
        if (object.validator !== undefined && object.validator !== null) {
            for (const e of object.validator) {
                message.validator.push(Validator.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetLastValidatorPowerRequest = { consensusAddress: '' };
export const QueryGetLastValidatorPowerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.consensusAddress !== '') {
            writer.uint32(10).string(message.consensusAddress);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetLastValidatorPowerRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.consensusAddress = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetLastValidatorPowerRequest };
        if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
            message.consensusAddress = String(object.consensusAddress);
        }
        else {
            message.consensusAddress = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.consensusAddress !== undefined && (obj.consensusAddress = message.consensusAddress);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetLastValidatorPowerRequest };
        if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
            message.consensusAddress = object.consensusAddress;
        }
        else {
            message.consensusAddress = '';
        }
        return message;
    }
};
const baseQueryGetLastValidatorPowerResponse = {};
export const QueryGetLastValidatorPowerResponse = {
    encode(message, writer = Writer.create()) {
        if (message.lastValidatorPower !== undefined) {
            LastValidatorPower.encode(message.lastValidatorPower, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetLastValidatorPowerResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.lastValidatorPower = LastValidatorPower.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetLastValidatorPowerResponse };
        if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
            message.lastValidatorPower = LastValidatorPower.fromJSON(object.lastValidatorPower);
        }
        else {
            message.lastValidatorPower = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.lastValidatorPower !== undefined &&
            (obj.lastValidatorPower = message.lastValidatorPower ? LastValidatorPower.toJSON(message.lastValidatorPower) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetLastValidatorPowerResponse };
        if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
            message.lastValidatorPower = LastValidatorPower.fromPartial(object.lastValidatorPower);
        }
        else {
            message.lastValidatorPower = undefined;
        }
        return message;
    }
};
const baseQueryAllLastValidatorPowerRequest = {};
export const QueryAllLastValidatorPowerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllLastValidatorPowerRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllLastValidatorPowerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllLastValidatorPowerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllLastValidatorPowerResponse = {};
export const QueryAllLastValidatorPowerResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.lastValidatorPower) {
            LastValidatorPower.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllLastValidatorPowerResponse };
        message.lastValidatorPower = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.lastValidatorPower.push(LastValidatorPower.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllLastValidatorPowerResponse };
        message.lastValidatorPower = [];
        if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
            for (const e of object.lastValidatorPower) {
                message.lastValidatorPower.push(LastValidatorPower.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.lastValidatorPower) {
            obj.lastValidatorPower = message.lastValidatorPower.map((e) => (e ? LastValidatorPower.toJSON(e) : undefined));
        }
        else {
            obj.lastValidatorPower = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllLastValidatorPowerResponse };
        message.lastValidatorPower = [];
        if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
            for (const e of object.lastValidatorPower) {
                message.lastValidatorPower.push(LastValidatorPower.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetValidatorSigningInfoRequest = { address: '' };
export const QueryGetValidatorSigningInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorSigningInfoRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorSigningInfoRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorSigningInfoRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetValidatorSigningInfoResponse = {};
export const QueryGetValidatorSigningInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.validatorSigningInfo !== undefined) {
            ValidatorSigningInfo.encode(message.validatorSigningInfo, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorSigningInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorSigningInfo = ValidatorSigningInfo.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorSigningInfoResponse };
        if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
            message.validatorSigningInfo = ValidatorSigningInfo.fromJSON(object.validatorSigningInfo);
        }
        else {
            message.validatorSigningInfo = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.validatorSigningInfo !== undefined &&
            (obj.validatorSigningInfo = message.validatorSigningInfo ? ValidatorSigningInfo.toJSON(message.validatorSigningInfo) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorSigningInfoResponse };
        if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
            message.validatorSigningInfo = ValidatorSigningInfo.fromPartial(object.validatorSigningInfo);
        }
        else {
            message.validatorSigningInfo = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorSigningInfoRequest = {};
export const QueryAllValidatorSigningInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorSigningInfoRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorSigningInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorSigningInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorSigningInfoResponse = {};
export const QueryAllValidatorSigningInfoResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.validatorSigningInfo) {
            ValidatorSigningInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorSigningInfoResponse };
        message.validatorSigningInfo = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorSigningInfo.push(ValidatorSigningInfo.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorSigningInfoResponse };
        message.validatorSigningInfo = [];
        if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
            for (const e of object.validatorSigningInfo) {
                message.validatorSigningInfo.push(ValidatorSigningInfo.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.validatorSigningInfo) {
            obj.validatorSigningInfo = message.validatorSigningInfo.map((e) => (e ? ValidatorSigningInfo.toJSON(e) : undefined));
        }
        else {
            obj.validatorSigningInfo = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorSigningInfoResponse };
        message.validatorSigningInfo = [];
        if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
            for (const e of object.validatorSigningInfo) {
                message.validatorSigningInfo.push(ValidatorSigningInfo.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetValidatorMissedBlockBitArrayRequest = { address: '', index: 0 };
export const QueryGetValidatorMissedBlockBitArrayRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        if (message.index !== 0) {
            writer.uint32(16).uint64(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorMissedBlockBitArrayRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.index = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorMissedBlockBitArrayRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = Number(object.index);
        }
        else {
            message.index = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorMissedBlockBitArrayRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = 0;
        }
        return message;
    }
};
const baseQueryGetValidatorMissedBlockBitArrayResponse = {};
export const QueryGetValidatorMissedBlockBitArrayResponse = {
    encode(message, writer = Writer.create()) {
        if (message.validatorMissedBlockBitArray !== undefined) {
            ValidatorMissedBlockBitArray.encode(message.validatorMissedBlockBitArray, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorMissedBlockBitArrayResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorMissedBlockBitArray = ValidatorMissedBlockBitArray.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorMissedBlockBitArrayResponse };
        if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
            message.validatorMissedBlockBitArray = ValidatorMissedBlockBitArray.fromJSON(object.validatorMissedBlockBitArray);
        }
        else {
            message.validatorMissedBlockBitArray = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.validatorMissedBlockBitArray !== undefined &&
            (obj.validatorMissedBlockBitArray = message.validatorMissedBlockBitArray
                ? ValidatorMissedBlockBitArray.toJSON(message.validatorMissedBlockBitArray)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorMissedBlockBitArrayResponse };
        if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
            message.validatorMissedBlockBitArray = ValidatorMissedBlockBitArray.fromPartial(object.validatorMissedBlockBitArray);
        }
        else {
            message.validatorMissedBlockBitArray = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorMissedBlockBitArrayRequest = {};
export const QueryAllValidatorMissedBlockBitArrayRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorMissedBlockBitArrayRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorMissedBlockBitArrayRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorMissedBlockBitArrayRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorMissedBlockBitArrayResponse = {};
export const QueryAllValidatorMissedBlockBitArrayResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.validatorMissedBlockBitArray) {
            ValidatorMissedBlockBitArray.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorMissedBlockBitArrayResponse };
        message.validatorMissedBlockBitArray = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorMissedBlockBitArray.push(ValidatorMissedBlockBitArray.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorMissedBlockBitArrayResponse };
        message.validatorMissedBlockBitArray = [];
        if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
            for (const e of object.validatorMissedBlockBitArray) {
                message.validatorMissedBlockBitArray.push(ValidatorMissedBlockBitArray.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.validatorMissedBlockBitArray) {
            obj.validatorMissedBlockBitArray = message.validatorMissedBlockBitArray.map((e) => (e ? ValidatorMissedBlockBitArray.toJSON(e) : undefined));
        }
        else {
            obj.validatorMissedBlockBitArray = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorMissedBlockBitArrayResponse };
        message.validatorMissedBlockBitArray = [];
        if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
            for (const e of object.validatorMissedBlockBitArray) {
                message.validatorMissedBlockBitArray.push(ValidatorMissedBlockBitArray.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetValidatorOwnerRequest = { address: '' };
export const QueryGetValidatorOwnerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorOwnerRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorOwnerRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorOwnerRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetValidatorOwnerResponse = {};
export const QueryGetValidatorOwnerResponse = {
    encode(message, writer = Writer.create()) {
        if (message.validatorOwner !== undefined) {
            ValidatorOwner.encode(message.validatorOwner, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetValidatorOwnerResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorOwner = ValidatorOwner.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetValidatorOwnerResponse };
        if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
            message.validatorOwner = ValidatorOwner.fromJSON(object.validatorOwner);
        }
        else {
            message.validatorOwner = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.validatorOwner !== undefined && (obj.validatorOwner = message.validatorOwner ? ValidatorOwner.toJSON(message.validatorOwner) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorOwnerResponse };
        if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
            message.validatorOwner = ValidatorOwner.fromPartial(object.validatorOwner);
        }
        else {
            message.validatorOwner = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorOwnerRequest = {};
export const QueryAllValidatorOwnerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorOwnerRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorOwnerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorOwnerRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllValidatorOwnerResponse = {};
export const QueryAllValidatorOwnerResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.validatorOwner) {
            ValidatorOwner.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllValidatorOwnerResponse };
        message.validatorOwner = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorOwner.push(ValidatorOwner.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllValidatorOwnerResponse };
        message.validatorOwner = [];
        if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
            for (const e of object.validatorOwner) {
                message.validatorOwner.push(ValidatorOwner.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.validatorOwner) {
            obj.validatorOwner = message.validatorOwner.map((e) => (e ? ValidatorOwner.toJSON(e) : undefined));
        }
        else {
            obj.validatorOwner = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllValidatorOwnerResponse };
        message.validatorOwner = [];
        if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
            for (const e of object.validatorOwner) {
                message.validatorOwner.push(ValidatorOwner.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Validator(request) {
        const data = QueryGetValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'Validator', data);
        return promise.then((data) => QueryGetValidatorResponse.decode(new Reader(data)));
    }
    ValidatorAll(request) {
        const data = QueryAllValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorAll', data);
        return promise.then((data) => QueryAllValidatorResponse.decode(new Reader(data)));
    }
    LastValidatorPower(request) {
        const data = QueryGetLastValidatorPowerRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'LastValidatorPower', data);
        return promise.then((data) => QueryGetLastValidatorPowerResponse.decode(new Reader(data)));
    }
    LastValidatorPowerAll(request) {
        const data = QueryAllLastValidatorPowerRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'LastValidatorPowerAll', data);
        return promise.then((data) => QueryAllLastValidatorPowerResponse.decode(new Reader(data)));
    }
    ValidatorSigningInfo(request) {
        const data = QueryGetValidatorSigningInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorSigningInfo', data);
        return promise.then((data) => QueryGetValidatorSigningInfoResponse.decode(new Reader(data)));
    }
    ValidatorSigningInfoAll(request) {
        const data = QueryAllValidatorSigningInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorSigningInfoAll', data);
        return promise.then((data) => QueryAllValidatorSigningInfoResponse.decode(new Reader(data)));
    }
    ValidatorMissedBlockBitArray(request) {
        const data = QueryGetValidatorMissedBlockBitArrayRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorMissedBlockBitArray', data);
        return promise.then((data) => QueryGetValidatorMissedBlockBitArrayResponse.decode(new Reader(data)));
    }
    ValidatorMissedBlockBitArrayAll(request) {
        const data = QueryAllValidatorMissedBlockBitArrayRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorMissedBlockBitArrayAll', data);
        return promise.then((data) => QueryAllValidatorMissedBlockBitArrayResponse.decode(new Reader(data)));
    }
    ValidatorOwner(request) {
        const data = QueryGetValidatorOwnerRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorOwner', data);
        return promise.then((data) => QueryGetValidatorOwnerResponse.decode(new Reader(data)));
    }
    ValidatorOwnerAll(request) {
        const data = QueryAllValidatorOwnerRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorOwnerAll', data);
        return promise.then((data) => QueryAllValidatorOwnerResponse.decode(new Reader(data)));
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
