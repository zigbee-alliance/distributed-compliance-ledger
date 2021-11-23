/* eslint-disable */
import { Validator } from '../validator/validator';
import { LastValidatorPower } from '../validator/last_validator_power';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.validatorList) {
            Validator.encode(v, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.lastValidatorPowerList) {
            LastValidatorPower.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.validatorList = [];
        message.lastValidatorPowerList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.validatorList.push(Validator.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.lastValidatorPowerList.push(LastValidatorPower.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.validatorList = [];
        message.lastValidatorPowerList = [];
        if (object.validatorList !== undefined && object.validatorList !== null) {
            for (const e of object.validatorList) {
                message.validatorList.push(Validator.fromJSON(e));
            }
        }
        if (object.lastValidatorPowerList !== undefined && object.lastValidatorPowerList !== null) {
            for (const e of object.lastValidatorPowerList) {
                message.lastValidatorPowerList.push(LastValidatorPower.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.validatorList) {
            obj.validatorList = message.validatorList.map((e) => (e ? Validator.toJSON(e) : undefined));
        }
        else {
            obj.validatorList = [];
        }
        if (message.lastValidatorPowerList) {
            obj.lastValidatorPowerList = message.lastValidatorPowerList.map((e) => (e ? LastValidatorPower.toJSON(e) : undefined));
        }
        else {
            obj.lastValidatorPowerList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.validatorList = [];
        message.lastValidatorPowerList = [];
        if (object.validatorList !== undefined && object.validatorList !== null) {
            for (const e of object.validatorList) {
                message.validatorList.push(Validator.fromPartial(e));
            }
        }
        if (object.lastValidatorPowerList !== undefined && object.lastValidatorPowerList !== null) {
            for (const e of object.lastValidatorPowerList) {
                message.lastValidatorPowerList.push(LastValidatorPower.fromPartial(e));
            }
        }
        return message;
    }
};
