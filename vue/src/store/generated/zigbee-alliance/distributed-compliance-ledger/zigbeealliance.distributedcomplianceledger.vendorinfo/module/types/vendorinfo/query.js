/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { VendorInfoType } from '../vendorinfo/vendor_info_type';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseQueryGetVendorInfoTypeRequest = { vendorID: 0 };
export const QueryGetVendorInfoTypeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vendorID !== 0) {
            writer.uint32(8).uint64(message.vendorID);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetVendorInfoTypeRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorID = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetVendorInfoTypeRequest };
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = Number(object.vendorID);
        }
        else {
            message.vendorID = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vendorID !== undefined && (obj.vendorID = message.vendorID);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetVendorInfoTypeRequest };
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = object.vendorID;
        }
        else {
            message.vendorID = 0;
        }
        return message;
    }
};
const baseQueryGetVendorInfoTypeResponse = {};
export const QueryGetVendorInfoTypeResponse = {
    encode(message, writer = Writer.create()) {
        if (message.vendorInfoType !== undefined) {
            VendorInfoType.encode(message.vendorInfoType, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetVendorInfoTypeResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorInfoType = VendorInfoType.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetVendorInfoTypeResponse };
        if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
            message.vendorInfoType = VendorInfoType.fromJSON(object.vendorInfoType);
        }
        else {
            message.vendorInfoType = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vendorInfoType !== undefined && (obj.vendorInfoType = message.vendorInfoType ? VendorInfoType.toJSON(message.vendorInfoType) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetVendorInfoTypeResponse };
        if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
            message.vendorInfoType = VendorInfoType.fromPartial(object.vendorInfoType);
        }
        else {
            message.vendorInfoType = undefined;
        }
        return message;
    }
};
const baseQueryAllVendorInfoTypeRequest = {};
export const QueryAllVendorInfoTypeRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllVendorInfoTypeRequest };
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
        const message = { ...baseQueryAllVendorInfoTypeRequest };
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
        const message = { ...baseQueryAllVendorInfoTypeRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllVendorInfoTypeResponse = {};
export const QueryAllVendorInfoTypeResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vendorInfoType) {
            VendorInfoType.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllVendorInfoTypeResponse };
        message.vendorInfoType = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorInfoType.push(VendorInfoType.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllVendorInfoTypeResponse };
        message.vendorInfoType = [];
        if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
            for (const e of object.vendorInfoType) {
                message.vendorInfoType.push(VendorInfoType.fromJSON(e));
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
        if (message.vendorInfoType) {
            obj.vendorInfoType = message.vendorInfoType.map((e) => (e ? VendorInfoType.toJSON(e) : undefined));
        }
        else {
            obj.vendorInfoType = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllVendorInfoTypeResponse };
        message.vendorInfoType = [];
        if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
            for (const e of object.vendorInfoType) {
                message.vendorInfoType.push(VendorInfoType.fromPartial(e));
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
    VendorInfoType(request) {
        const data = QueryGetVendorInfoTypeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfoType', data);
        return promise.then((data) => QueryGetVendorInfoTypeResponse.decode(new Reader(data)));
    }
    VendorInfoTypeAll(request) {
        const data = QueryAllVendorInfoTypeRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfoTypeAll', data);
        return promise.then((data) => QueryAllVendorInfoTypeResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
