/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Description } from '../validator/description';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator';
const baseMsgCreateValidator = { signer: '', address: '', pubKey: '' };
export const MsgCreateValidator = {
    encode(message, writer = Writer.create()) {
        if (message.signer !== '') {
            writer.uint32(10).string(message.signer);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        if (message.pubKey !== '') {
            writer.uint32(26).string(message.pubKey);
        }
        if (message.description !== undefined) {
            Description.encode(message.description, writer.uint32(34).fork()).ldelim();
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
                    message.address = reader.string();
                    break;
                case 3:
                    message.pubKey = reader.string();
                    break;
                case 4:
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
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = String(object.pubKey);
        }
        else {
            message.pubKey = '';
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
        message.address !== undefined && (obj.address = message.address);
        message.pubKey !== undefined && (obj.pubKey = message.pubKey);
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
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = object.pubKey;
        }
        else {
            message.pubKey = '';
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
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateValidator(request) {
        const data = MsgCreateValidator.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'CreateValidator', data);
        return promise.then((data) => MsgCreateValidatorResponse.decode(new Reader(data)));
    }
}
