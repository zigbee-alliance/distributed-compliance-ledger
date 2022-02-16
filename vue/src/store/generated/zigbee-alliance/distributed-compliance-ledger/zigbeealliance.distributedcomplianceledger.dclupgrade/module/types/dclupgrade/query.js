/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade';
const baseQueryGetProposedUpgradeRequest = { name: '' };
export const QueryGetProposedUpgradeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.name !== '') {
            writer.uint32(10).string(message.name);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedUpgradeRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetProposedUpgradeRequest };
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProposedUpgradeRequest };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        return message;
    }
};
const baseQueryGetProposedUpgradeResponse = {};
export const QueryGetProposedUpgradeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.proposedUpgrade !== undefined) {
            ProposedUpgrade.encode(message.proposedUpgrade, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedUpgradeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposedUpgrade = ProposedUpgrade.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetProposedUpgradeResponse };
        if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
            message.proposedUpgrade = ProposedUpgrade.fromJSON(object.proposedUpgrade);
        }
        else {
            message.proposedUpgrade = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposedUpgrade !== undefined && (obj.proposedUpgrade = message.proposedUpgrade ? ProposedUpgrade.toJSON(message.proposedUpgrade) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProposedUpgradeResponse };
        if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
            message.proposedUpgrade = ProposedUpgrade.fromPartial(object.proposedUpgrade);
        }
        else {
            message.proposedUpgrade = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedUpgradeRequest = {};
export const QueryAllProposedUpgradeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedUpgradeRequest };
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
        const message = { ...baseQueryAllProposedUpgradeRequest };
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
        const message = { ...baseQueryAllProposedUpgradeRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedUpgradeResponse = {};
export const QueryAllProposedUpgradeResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.proposedUpgrade) {
            ProposedUpgrade.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedUpgradeResponse };
        message.proposedUpgrade = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposedUpgrade.push(ProposedUpgrade.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllProposedUpgradeResponse };
        message.proposedUpgrade = [];
        if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
            for (const e of object.proposedUpgrade) {
                message.proposedUpgrade.push(ProposedUpgrade.fromJSON(e));
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
        if (message.proposedUpgrade) {
            obj.proposedUpgrade = message.proposedUpgrade.map((e) => (e ? ProposedUpgrade.toJSON(e) : undefined));
        }
        else {
            obj.proposedUpgrade = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllProposedUpgradeResponse };
        message.proposedUpgrade = [];
        if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
            for (const e of object.proposedUpgrade) {
                message.proposedUpgrade.push(ProposedUpgrade.fromPartial(e));
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
    ProposedUpgrade(request) {
        const data = QueryGetProposedUpgradeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ProposedUpgrade', data);
        return promise.then((data) => QueryGetProposedUpgradeResponse.decode(new Reader(data)));
    }
    ProposedUpgradeAll(request) {
        const data = QueryAllProposedUpgradeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ProposedUpgradeAll', data);
        return promise.then((data) => QueryAllProposedUpgradeResponse.decode(new Reader(data)));
    }
}
