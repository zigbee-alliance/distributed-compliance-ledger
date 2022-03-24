/* eslint-disable */
import { Account } from '../dclauth/account';
import { Grant } from '../dclauth/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
const baseRevokedAccount = {};
export const RevokedAccount = {
    encode(message, writer = Writer.create()) {
        if (message.account !== undefined) {
            Account.encode(message.account, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.revokeApprovals) {
            Grant.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRevokedAccount };
        message.revokeApprovals = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.account = Account.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.revokeApprovals.push(Grant.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseRevokedAccount };
        message.revokeApprovals = [];
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromJSON(object.account);
        }
        else {
            message.account = undefined;
        }
        if (object.revokeApprovals !== undefined && object.revokeApprovals !== null) {
            for (const e of object.revokeApprovals) {
                message.revokeApprovals.push(Grant.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
        if (message.revokeApprovals) {
            obj.revokeApprovals = message.revokeApprovals.map((e) => (e ? Grant.toJSON(e) : undefined));
        }
        else {
            obj.revokeApprovals = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRevokedAccount };
        message.revokeApprovals = [];
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromPartial(object.account);
        }
        else {
            message.account = undefined;
        }
        if (object.revokeApprovals !== undefined && object.revokeApprovals !== null) {
            for (const e of object.revokeApprovals) {
                message.revokeApprovals.push(Grant.fromPartial(e));
            }
        }
        return message;
    }
};
