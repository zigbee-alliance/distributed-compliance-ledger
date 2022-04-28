/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { ApprovedUpgrade } from '../dclupgrade/approved_upgrade';
import { RejectedUpgrade } from '../dclupgrade/rejected_upgrade';
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
const baseQueryGetApprovedUpgradeRequest = { name: '' };
export const QueryGetApprovedUpgradeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.name !== '') {
            writer.uint32(10).string(message.name);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetApprovedUpgradeRequest };
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
        const message = { ...baseQueryGetApprovedUpgradeRequest };
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
        const message = { ...baseQueryGetApprovedUpgradeRequest };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        return message;
    }
};
const baseQueryGetApprovedUpgradeResponse = {};
export const QueryGetApprovedUpgradeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.approvedUpgrade !== undefined) {
            ApprovedUpgrade.encode(message.approvedUpgrade, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetApprovedUpgradeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.approvedUpgrade = ApprovedUpgrade.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetApprovedUpgradeResponse };
        if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
            message.approvedUpgrade = ApprovedUpgrade.fromJSON(object.approvedUpgrade);
        }
        else {
            message.approvedUpgrade = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.approvedUpgrade !== undefined && (obj.approvedUpgrade = message.approvedUpgrade ? ApprovedUpgrade.toJSON(message.approvedUpgrade) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetApprovedUpgradeResponse };
        if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
            message.approvedUpgrade = ApprovedUpgrade.fromPartial(object.approvedUpgrade);
        }
        else {
            message.approvedUpgrade = undefined;
        }
        return message;
    }
};
const baseQueryAllApprovedUpgradeRequest = {};
export const QueryAllApprovedUpgradeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllApprovedUpgradeRequest };
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
        const message = { ...baseQueryAllApprovedUpgradeRequest };
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
        const message = { ...baseQueryAllApprovedUpgradeRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllApprovedUpgradeResponse = {};
export const QueryAllApprovedUpgradeResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.approvedUpgrade) {
            ApprovedUpgrade.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllApprovedUpgradeResponse };
        message.approvedUpgrade = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.approvedUpgrade.push(ApprovedUpgrade.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllApprovedUpgradeResponse };
        message.approvedUpgrade = [];
        if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
            for (const e of object.approvedUpgrade) {
                message.approvedUpgrade.push(ApprovedUpgrade.fromJSON(e));
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
        if (message.approvedUpgrade) {
            obj.approvedUpgrade = message.approvedUpgrade.map((e) => (e ? ApprovedUpgrade.toJSON(e) : undefined));
        }
        else {
            obj.approvedUpgrade = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllApprovedUpgradeResponse };
        message.approvedUpgrade = [];
        if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
            for (const e of object.approvedUpgrade) {
                message.approvedUpgrade.push(ApprovedUpgrade.fromPartial(e));
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
const baseQueryGetRejectedUpgradeRequest = { name: '' };
export const QueryGetRejectedUpgradeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.name !== '') {
            writer.uint32(10).string(message.name);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRejectedUpgradeRequest };
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
        const message = { ...baseQueryGetRejectedUpgradeRequest };
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
        const message = { ...baseQueryGetRejectedUpgradeRequest };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        return message;
    }
};
const baseQueryGetRejectedUpgradeResponse = {};
export const QueryGetRejectedUpgradeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.rejectedUpgrade !== undefined) {
            RejectedUpgrade.encode(message.rejectedUpgrade, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRejectedUpgradeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.rejectedUpgrade = RejectedUpgrade.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetRejectedUpgradeResponse };
        if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
            message.rejectedUpgrade = RejectedUpgrade.fromJSON(object.rejectedUpgrade);
        }
        else {
            message.rejectedUpgrade = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.rejectedUpgrade !== undefined && (obj.rejectedUpgrade = message.rejectedUpgrade ? RejectedUpgrade.toJSON(message.rejectedUpgrade) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRejectedUpgradeResponse };
        if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
            message.rejectedUpgrade = RejectedUpgrade.fromPartial(object.rejectedUpgrade);
        }
        else {
            message.rejectedUpgrade = undefined;
        }
        return message;
    }
};
const baseQueryAllRejectedUpgradeRequest = {};
export const QueryAllRejectedUpgradeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRejectedUpgradeRequest };
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
        const message = { ...baseQueryAllRejectedUpgradeRequest };
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
        const message = { ...baseQueryAllRejectedUpgradeRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllRejectedUpgradeResponse = {};
export const QueryAllRejectedUpgradeResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.rejectedUpgrade) {
            RejectedUpgrade.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRejectedUpgradeResponse };
        message.rejectedUpgrade = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.rejectedUpgrade.push(RejectedUpgrade.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllRejectedUpgradeResponse };
        message.rejectedUpgrade = [];
        if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
            for (const e of object.rejectedUpgrade) {
                message.rejectedUpgrade.push(RejectedUpgrade.fromJSON(e));
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
        if (message.rejectedUpgrade) {
            obj.rejectedUpgrade = message.rejectedUpgrade.map((e) => (e ? RejectedUpgrade.toJSON(e) : undefined));
        }
        else {
            obj.rejectedUpgrade = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRejectedUpgradeResponse };
        message.rejectedUpgrade = [];
        if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
            for (const e of object.rejectedUpgrade) {
                message.rejectedUpgrade.push(RejectedUpgrade.fromPartial(e));
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
    ApprovedUpgrade(request) {
        const data = QueryGetApprovedUpgradeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ApprovedUpgrade', data);
        return promise.then((data) => QueryGetApprovedUpgradeResponse.decode(new Reader(data)));
    }
    ApprovedUpgradeAll(request) {
        const data = QueryAllApprovedUpgradeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ApprovedUpgradeAll', data);
        return promise.then((data) => QueryAllApprovedUpgradeResponse.decode(new Reader(data)));
    }
    RejectedUpgrade(request) {
        const data = QueryGetRejectedUpgradeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'RejectedUpgrade', data);
        return promise.then((data) => QueryGetRejectedUpgradeResponse.decode(new Reader(data)));
    }
    RejectedUpgradeAll(request) {
        const data = QueryAllRejectedUpgradeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'RejectedUpgradeAll', data);
        return promise.then((data) => QueryAllRejectedUpgradeResponse.decode(new Reader(data)));
    }
}
