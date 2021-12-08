/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { NewVendorInfo } from '../vendorinfo/new_vendor_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseQueryGetNewVendorInfoRequest = { index: '' };
export const QueryGetNewVendorInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNewVendorInfoRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.index = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetNewVendorInfoRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetNewVendorInfoRequest };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        return message;
    }
};
const baseQueryGetNewVendorInfoResponse = {};
export const QueryGetNewVendorInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.newVendorInfo !== undefined) {
            NewVendorInfo.encode(message.newVendorInfo, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNewVendorInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.newVendorInfo = NewVendorInfo.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetNewVendorInfoResponse };
        if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
            message.newVendorInfo = NewVendorInfo.fromJSON(object.newVendorInfo);
        }
        else {
            message.newVendorInfo = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.newVendorInfo !== undefined && (obj.newVendorInfo = message.newVendorInfo ? NewVendorInfo.toJSON(message.newVendorInfo) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetNewVendorInfoResponse };
        if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
            message.newVendorInfo = NewVendorInfo.fromPartial(object.newVendorInfo);
        }
        else {
            message.newVendorInfo = undefined;
        }
        return message;
    }
};
const baseQueryAllNewVendorInfoRequest = {};
export const QueryAllNewVendorInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllNewVendorInfoRequest };
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
        const message = { ...baseQueryAllNewVendorInfoRequest };
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
        const message = { ...baseQueryAllNewVendorInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllNewVendorInfoResponse = {};
export const QueryAllNewVendorInfoResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.newVendorInfo) {
            NewVendorInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllNewVendorInfoResponse };
        message.newVendorInfo = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.newVendorInfo.push(NewVendorInfo.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllNewVendorInfoResponse };
        message.newVendorInfo = [];
        if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
            for (const e of object.newVendorInfo) {
                message.newVendorInfo.push(NewVendorInfo.fromJSON(e));
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
        if (message.newVendorInfo) {
            obj.newVendorInfo = message.newVendorInfo.map((e) => (e ? NewVendorInfo.toJSON(e) : undefined));
        }
        else {
            obj.newVendorInfo = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllNewVendorInfoResponse };
        message.newVendorInfo = [];
        if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
            for (const e of object.newVendorInfo) {
                message.newVendorInfo.push(NewVendorInfo.fromPartial(e));
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
    NewVendorInfo(request) {
        const data = QueryGetNewVendorInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'NewVendorInfo', data);
        return promise.then((data) => QueryGetNewVendorInfoResponse.decode(new Reader(data)));
    }
    NewVendorInfoAll(request) {
        const data = QueryAllNewVendorInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'NewVendorInfoAll', data);
        return promise.then((data) => QueryAllNewVendorInfoResponse.decode(new Reader(data)));
    }
}
