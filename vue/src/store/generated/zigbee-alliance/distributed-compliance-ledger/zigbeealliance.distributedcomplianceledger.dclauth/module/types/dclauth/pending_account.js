/* eslint-disable */
import { Account } from '../dclauth/account';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
const basePendingAccount = {};
export const PendingAccount = {
    encode(message, writer = Writer.create()) {
        if (message.account !== undefined) {
            Account.encode(message.account, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePendingAccount };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.account = Account.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...basePendingAccount };
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromJSON(object.account);
        }
        else {
            message.account = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePendingAccount };
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromPartial(object.account);
        }
        else {
            message.account = undefined;
        }
        return message;
    }
};
