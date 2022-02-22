/* eslint-disable */
import { Grant } from '../dclauth/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
const basePendingAccountRevocation = { address: '' };
export const PendingAccountRevocation = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        for (const v of message.approvals) {
            Grant.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePendingAccountRevocation };
        message.approvals = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                case 2:
                    message.approvals.push(Grant.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...basePendingAccountRevocation };
        message.approvals = [];
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(Grant.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        if (message.approvals) {
            obj.approvals = message.approvals.map((e) => (e ? Grant.toJSON(e) : undefined));
        }
        else {
            obj.approvals = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePendingAccountRevocation };
        message.approvals = [];
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(Grant.fromPartial(e));
            }
        }
        return message;
    }
};
