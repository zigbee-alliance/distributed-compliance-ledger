/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade';
const baseApprovedUpgrade = { name: '', approvals: '' };
export const ApprovedUpgrade = {
    encode(message, writer = Writer.create()) {
        if (message.name !== '') {
            writer.uint32(10).string(message.name);
        }
        for (const v of message.approvals) {
            writer.uint32(18).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseApprovedUpgrade };
        message.approvals = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.approvals.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseApprovedUpgrade };
        message.approvals = [];
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
        }
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        if (message.approvals) {
            obj.approvals = message.approvals.map((e) => e);
        }
        else {
            obj.approvals = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseApprovedUpgrade };
        message.approvals = [];
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(e);
            }
        }
        return message;
    }
};
