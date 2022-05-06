/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Validator } from '../validator/validator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { LastValidatorPower } from '../validator/last_validator_power';
import { ProposedDisableValidator } from '../validator/proposed_disable_validator';
import { DisabledValidator } from '../validator/disabled_validator';
import { RejectedDisableValidator } from '../validator/rejected_validator';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseQueryGetValidatorRequest = { owner: '' };
export const QueryGetValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== '') {
            writer.uint32(10).string(message.owner);
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
        const message = { ...baseQueryGetValidatorRequest };
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
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetValidatorRequest };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = '';
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
const baseQueryGetLastValidatorPowerRequest = { owner: '' };
export const QueryGetLastValidatorPowerRequest = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== '') {
            writer.uint32(10).string(message.owner);
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
        const message = { ...baseQueryGetLastValidatorPowerRequest };
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
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetLastValidatorPowerRequest };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = '';
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
const baseQueryGetProposedDisableValidatorRequest = { address: '' };
export const QueryGetProposedDisableValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedDisableValidatorRequest };
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
        const message = { ...baseQueryGetProposedDisableValidatorRequest };
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
        const message = { ...baseQueryGetProposedDisableValidatorRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetProposedDisableValidatorResponse = {};
export const QueryGetProposedDisableValidatorResponse = {
    encode(message, writer = Writer.create()) {
        if (message.proposedDisableValidator !== undefined) {
            ProposedDisableValidator.encode(message.proposedDisableValidator, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedDisableValidatorResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposedDisableValidator = ProposedDisableValidator.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetProposedDisableValidatorResponse };
        if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
            message.proposedDisableValidator = ProposedDisableValidator.fromJSON(object.proposedDisableValidator);
        }
        else {
            message.proposedDisableValidator = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposedDisableValidator !== undefined &&
            (obj.proposedDisableValidator = message.proposedDisableValidator ? ProposedDisableValidator.toJSON(message.proposedDisableValidator) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProposedDisableValidatorResponse };
        if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
            message.proposedDisableValidator = ProposedDisableValidator.fromPartial(object.proposedDisableValidator);
        }
        else {
            message.proposedDisableValidator = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedDisableValidatorRequest = {};
export const QueryAllProposedDisableValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedDisableValidatorRequest };
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
        const message = { ...baseQueryAllProposedDisableValidatorRequest };
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
        const message = { ...baseQueryAllProposedDisableValidatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedDisableValidatorResponse = {};
export const QueryAllProposedDisableValidatorResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.proposedDisableValidator) {
            ProposedDisableValidator.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedDisableValidatorResponse };
        message.proposedDisableValidator = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposedDisableValidator.push(ProposedDisableValidator.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllProposedDisableValidatorResponse };
        message.proposedDisableValidator = [];
        if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
            for (const e of object.proposedDisableValidator) {
                message.proposedDisableValidator.push(ProposedDisableValidator.fromJSON(e));
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
        if (message.proposedDisableValidator) {
            obj.proposedDisableValidator = message.proposedDisableValidator.map((e) => (e ? ProposedDisableValidator.toJSON(e) : undefined));
        }
        else {
            obj.proposedDisableValidator = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllProposedDisableValidatorResponse };
        message.proposedDisableValidator = [];
        if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
            for (const e of object.proposedDisableValidator) {
                message.proposedDisableValidator.push(ProposedDisableValidator.fromPartial(e));
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
const baseQueryGetDisabledValidatorRequest = { address: '' };
export const QueryGetDisabledValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetDisabledValidatorRequest };
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
        const message = { ...baseQueryGetDisabledValidatorRequest };
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
        const message = { ...baseQueryGetDisabledValidatorRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetDisabledValidatorResponse = {};
export const QueryGetDisabledValidatorResponse = {
    encode(message, writer = Writer.create()) {
        if (message.disabledValidator !== undefined) {
            DisabledValidator.encode(message.disabledValidator, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetDisabledValidatorResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.disabledValidator = DisabledValidator.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetDisabledValidatorResponse };
        if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
            message.disabledValidator = DisabledValidator.fromJSON(object.disabledValidator);
        }
        else {
            message.disabledValidator = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.disabledValidator !== undefined &&
            (obj.disabledValidator = message.disabledValidator ? DisabledValidator.toJSON(message.disabledValidator) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetDisabledValidatorResponse };
        if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
            message.disabledValidator = DisabledValidator.fromPartial(object.disabledValidator);
        }
        else {
            message.disabledValidator = undefined;
        }
        return message;
    }
};
const baseQueryAllDisabledValidatorRequest = {};
export const QueryAllDisabledValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllDisabledValidatorRequest };
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
        const message = { ...baseQueryAllDisabledValidatorRequest };
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
        const message = { ...baseQueryAllDisabledValidatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllDisabledValidatorResponse = {};
export const QueryAllDisabledValidatorResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.disabledValidator) {
            DisabledValidator.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllDisabledValidatorResponse };
        message.disabledValidator = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.disabledValidator.push(DisabledValidator.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllDisabledValidatorResponse };
        message.disabledValidator = [];
        if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
            for (const e of object.disabledValidator) {
                message.disabledValidator.push(DisabledValidator.fromJSON(e));
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
        if (message.disabledValidator) {
            obj.disabledValidator = message.disabledValidator.map((e) => (e ? DisabledValidator.toJSON(e) : undefined));
        }
        else {
            obj.disabledValidator = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllDisabledValidatorResponse };
        message.disabledValidator = [];
        if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
            for (const e of object.disabledValidator) {
                message.disabledValidator.push(DisabledValidator.fromPartial(e));
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
const baseQueryGetRejectedDisableValidatorRequest = { owner: '' };
export const QueryGetRejectedDisableValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.owner !== '') {
            writer.uint32(10).string(message.owner);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRejectedDisableValidatorRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
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
        const message = { ...baseQueryGetRejectedDisableValidatorRequest };
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
        message.owner !== undefined && (obj.owner = message.owner);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRejectedDisableValidatorRequest };
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = '';
        }
        return message;
    }
};
const baseQueryGetRejectedDisableValidatorResponse = {};
export const QueryGetRejectedDisableValidatorResponse = {
    encode(message, writer = Writer.create()) {
        if (message.rejectedValidator !== undefined) {
            RejectedDisableValidator.encode(message.rejectedValidator, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRejectedDisableValidatorResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.rejectedValidator = RejectedDisableValidator.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetRejectedDisableValidatorResponse };
        if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
            message.rejectedValidator = RejectedDisableValidator.fromJSON(object.rejectedValidator);
        }
        else {
            message.rejectedValidator = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.rejectedValidator !== undefined &&
            (obj.rejectedValidator = message.rejectedValidator ? RejectedDisableValidator.toJSON(message.rejectedValidator) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRejectedDisableValidatorResponse };
        if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
            message.rejectedValidator = RejectedDisableValidator.fromPartial(object.rejectedValidator);
        }
        else {
            message.rejectedValidator = undefined;
        }
        return message;
    }
};
const baseQueryAllRejectedDisableValidatorRequest = {};
export const QueryAllRejectedDisableValidatorRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRejectedDisableValidatorRequest };
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
        const message = { ...baseQueryAllRejectedDisableValidatorRequest };
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
        const message = { ...baseQueryAllRejectedDisableValidatorRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllRejectedDisableValidatorResponse = {};
export const QueryAllRejectedDisableValidatorResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.rejectedValidator) {
            RejectedDisableValidator.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRejectedDisableValidatorResponse };
        message.rejectedValidator = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.rejectedValidator.push(RejectedDisableValidator.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllRejectedDisableValidatorResponse };
        message.rejectedValidator = [];
        if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
            for (const e of object.rejectedValidator) {
                message.rejectedValidator.push(RejectedDisableValidator.fromJSON(e));
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
        if (message.rejectedValidator) {
            obj.rejectedValidator = message.rejectedValidator.map((e) => (e ? RejectedDisableValidator.toJSON(e) : undefined));
        }
        else {
            obj.rejectedValidator = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRejectedDisableValidatorResponse };
        message.rejectedValidator = [];
        if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
            for (const e of object.rejectedValidator) {
                message.rejectedValidator.push(RejectedDisableValidator.fromPartial(e));
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
    ProposedDisableValidator(request) {
        const data = QueryGetProposedDisableValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ProposedDisableValidator', data);
        return promise.then((data) => QueryGetProposedDisableValidatorResponse.decode(new Reader(data)));
    }
    ProposedDisableValidatorAll(request) {
        const data = QueryAllProposedDisableValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ProposedDisableValidatorAll', data);
        return promise.then((data) => QueryAllProposedDisableValidatorResponse.decode(new Reader(data)));
    }
    DisabledValidator(request) {
        const data = QueryGetDisabledValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'DisabledValidator', data);
        return promise.then((data) => QueryGetDisabledValidatorResponse.decode(new Reader(data)));
    }
    DisabledValidatorAll(request) {
        const data = QueryAllDisabledValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'DisabledValidatorAll', data);
        return promise.then((data) => QueryAllDisabledValidatorResponse.decode(new Reader(data)));
    }
    RejectedDisableValidator(request) {
        const data = QueryGetRejectedDisableValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'RejectedDisableValidator', data);
        return promise.then((data) => QueryGetRejectedDisableValidatorResponse.decode(new Reader(data)));
    }
    RejectedDisableValidatorAll(request) {
        const data = QueryAllRejectedDisableValidatorRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'RejectedDisableValidatorAll', data);
        return promise.then((data) => QueryAllRejectedDisableValidatorResponse.decode(new Reader(data)));
    }
}
