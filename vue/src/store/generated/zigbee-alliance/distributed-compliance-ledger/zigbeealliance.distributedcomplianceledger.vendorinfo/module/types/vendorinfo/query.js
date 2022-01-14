/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { VendorInfo } from '../vendorinfo/vendor_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseQueryGetVendorInfoRequest = { vendorID: 0 };
export const QueryGetVendorInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vendorID !== 0) {
            writer.uint32(8).int32(message.vendorID);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetVendorInfoRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorID = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetVendorInfoRequest };
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
        const message = { ...baseQueryGetVendorInfoRequest };
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = object.vendorID;
        }
        else {
            message.vendorID = 0;
        }
        return message;
    }
};
const baseQueryGetVendorInfoResponse = {};
export const QueryGetVendorInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.vendorInfo !== undefined) {
            VendorInfo.encode(message.vendorInfo, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetVendorInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorInfo = VendorInfo.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetVendorInfoResponse };
        if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
            message.vendorInfo = VendorInfo.fromJSON(object.vendorInfo);
        }
        else {
            message.vendorInfo = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vendorInfo !== undefined && (obj.vendorInfo = message.vendorInfo ? VendorInfo.toJSON(message.vendorInfo) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetVendorInfoResponse };
        if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
            message.vendorInfo = VendorInfo.fromPartial(object.vendorInfo);
        }
        else {
            message.vendorInfo = undefined;
        }
        return message;
    }
};
const baseQueryAllVendorInfoRequest = {};
export const QueryAllVendorInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllVendorInfoRequest };
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
        const message = { ...baseQueryAllVendorInfoRequest };
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
        const message = { ...baseQueryAllVendorInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllVendorInfoResponse = {};
export const QueryAllVendorInfoResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vendorInfo) {
            VendorInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllVendorInfoResponse };
        message.vendorInfo = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorInfo.push(VendorInfo.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllVendorInfoResponse };
        message.vendorInfo = [];
        if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
            for (const e of object.vendorInfo) {
                message.vendorInfo.push(VendorInfo.fromJSON(e));
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
        if (message.vendorInfo) {
            obj.vendorInfo = message.vendorInfo.map((e) => (e ? VendorInfo.toJSON(e) : undefined));
        }
        else {
            obj.vendorInfo = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllVendorInfoResponse };
        message.vendorInfo = [];
        if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
            for (const e of object.vendorInfo) {
                message.vendorInfo.push(VendorInfo.fromPartial(e));
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
    VendorInfo(request) {
        const data = QueryGetVendorInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfo', data);
        return promise.then((data) => QueryGetVendorInfoResponse.decode(new Reader(data)));
    }
    VendorInfoAll(request) {
        const data = QueryAllVendorInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfoAll', data);
        return promise.then((data) => QueryAllVendorInfoResponse.decode(new Reader(data)));
    }
}
