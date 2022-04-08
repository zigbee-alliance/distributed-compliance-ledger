/* eslint-disable */
import { Account } from '../dclauth/account';
import { Grant } from '../dclauth/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
const baseRejectedAccount = {};
export const RejectedAccount = {
    encode(message, writer = Writer.create()) {
        if (message.account !== undefined) {
            Account.encode(message.account, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.rejectApprovals) {
            Grant.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRejectedAccount };
        message.rejectApprovals = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.account = Account.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.rejectApprovals.push(Grant.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseRejectedAccount };
        message.rejectApprovals = [];
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromJSON(object.account);
        }
        else {
            message.account = undefined;
        }
        if (object.rejectApprovals !== undefined && object.rejectApprovals !== null) {
            for (const e of object.rejectApprovals) {
                message.rejectApprovals.push(Grant.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
        if (message.rejectApprovals) {
            obj.rejectApprovals = message.rejectApprovals.map((e) => (e ? Grant.toJSON(e) : undefined));
        }
        else {
            obj.rejectApprovals = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRejectedAccount };
        message.rejectApprovals = [];
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromPartial(object.account);
        }
        else {
            message.account = undefined;
        }
        if (object.rejectApprovals !== undefined && object.rejectApprovals !== null) {
            for (const e of object.rejectApprovals) {
                message.rejectApprovals.push(Grant.fromPartial(e));
            }
        }
        return message;
    }
};
