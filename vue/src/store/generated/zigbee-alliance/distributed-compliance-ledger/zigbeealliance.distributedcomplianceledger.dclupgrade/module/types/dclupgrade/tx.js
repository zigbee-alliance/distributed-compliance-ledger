/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade';
const baseMsgProposeUpgrade = { creator: '', plan: '' };
export const MsgProposeUpgrade = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.plan !== '') {
            writer.uint32(18).string(message.plan);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeUpgrade };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.plan = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgProposeUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        if (object.plan !== undefined && object.plan !== null) {
            message.plan = String(object.plan);
        }
        else {
            message.plan = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.plan !== undefined && (obj.plan = message.plan);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgProposeUpgrade };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        if (object.plan !== undefined && object.plan !== null) {
            message.plan = object.plan;
        }
        else {
            message.plan = '';
        }
        return message;
    }
};
const baseMsgProposeUpgradeResponse = {};
export const MsgProposeUpgradeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgProposeUpgradeResponse };
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
        const message = { ...baseMsgProposeUpgradeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgProposeUpgradeResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ProposeUpgrade(request) {
        const data = MsgProposeUpgrade.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'ProposeUpgrade', data);
        return promise.then((data) => MsgProposeUpgradeResponse.decode(new Reader(data)));
    }
}
