/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
import { VendorProducts } from '../model/vendor_products';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { Model } from '../model/model';
import { ModelVersion } from '../model/model_version';
import { ModelVersions } from '../model/model_versions';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model';
const baseQueryGetVendorProductsRequest = { vid: 0 };
export const QueryGetVendorProductsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetVendorProductsRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetVendorProductsRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetVendorProductsRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        return message;
    }
};
const baseQueryGetVendorProductsResponse = {};
export const QueryGetVendorProductsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.vendorProducts !== undefined) {
            VendorProducts.encode(message.vendorProducts, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetVendorProductsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorProducts = VendorProducts.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetVendorProductsResponse };
        if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
            message.vendorProducts = VendorProducts.fromJSON(object.vendorProducts);
        }
        else {
            message.vendorProducts = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vendorProducts !== undefined && (obj.vendorProducts = message.vendorProducts ? VendorProducts.toJSON(message.vendorProducts) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetVendorProductsResponse };
        if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
            message.vendorProducts = VendorProducts.fromPartial(object.vendorProducts);
        }
        else {
            message.vendorProducts = undefined;
        }
        return message;
    }
};
const baseQueryAllVendorProductsRequest = {};
export const QueryAllVendorProductsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllVendorProductsRequest };
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
        const message = { ...baseQueryAllVendorProductsRequest };
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
        const message = { ...baseQueryAllVendorProductsRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllVendorProductsResponse = {};
export const QueryAllVendorProductsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vendorProducts) {
            VendorProducts.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllVendorProductsResponse };
        message.vendorProducts = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorProducts.push(VendorProducts.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllVendorProductsResponse };
        message.vendorProducts = [];
        if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
            for (const e of object.vendorProducts) {
                message.vendorProducts.push(VendorProducts.fromJSON(e));
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
        if (message.vendorProducts) {
            obj.vendorProducts = message.vendorProducts.map((e) => (e ? VendorProducts.toJSON(e) : undefined));
        }
        else {
            obj.vendorProducts = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllVendorProductsResponse };
        message.vendorProducts = [];
        if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
            for (const e of object.vendorProducts) {
                message.vendorProducts.push(VendorProducts.fromPartial(e));
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
const baseQueryGetModelRequest = { vid: 0, pid: 0 };
export const QueryGetModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetModelRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.pid = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        return message;
    }
};
const baseQueryGetModelResponse = {};
export const QueryGetModelResponse = {
    encode(message, writer = Writer.create()) {
        if (message.model !== undefined) {
            Model.encode(message.model, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetModelResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.model = Model.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetModelResponse };
        if (object.model !== undefined && object.model !== null) {
            message.model = Model.fromJSON(object.model);
        }
        else {
            message.model = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.model !== undefined && (obj.model = message.model ? Model.toJSON(message.model) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetModelResponse };
        if (object.model !== undefined && object.model !== null) {
            message.model = Model.fromPartial(object.model);
        }
        else {
            message.model = undefined;
        }
        return message;
    }
};
const baseQueryAllModelRequest = {};
export const QueryAllModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllModelRequest };
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
        const message = { ...baseQueryAllModelRequest };
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
        const message = { ...baseQueryAllModelRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllModelResponse = {};
export const QueryAllModelResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.model) {
            Model.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllModelResponse };
        message.model = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.model.push(Model.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllModelResponse };
        message.model = [];
        if (object.model !== undefined && object.model !== null) {
            for (const e of object.model) {
                message.model.push(Model.fromJSON(e));
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
        if (message.model) {
            obj.model = message.model.map((e) => (e ? Model.toJSON(e) : undefined));
        }
        else {
            obj.model = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllModelResponse };
        message.model = [];
        if (object.model !== undefined && object.model !== null) {
            for (const e of object.model) {
                message.model.push(Model.fromPartial(e));
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
const baseQueryGetModelVersionRequest = { vid: 0, pid: 0, softwareVersion: 0 };
export const QueryGetModelVersionRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint64(message.softwareVersion);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetModelVersionRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.pid = reader.int32();
                    break;
                case 3:
                    message.softwareVersion = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetModelVersionRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetModelVersionRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        return message;
    }
};
const baseQueryGetModelVersionResponse = {};
export const QueryGetModelVersionResponse = {
    encode(message, writer = Writer.create()) {
        if (message.modelVersion !== undefined) {
            ModelVersion.encode(message.modelVersion, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetModelVersionResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.modelVersion = ModelVersion.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetModelVersionResponse };
        if (object.modelVersion !== undefined && object.modelVersion !== null) {
            message.modelVersion = ModelVersion.fromJSON(object.modelVersion);
        }
        else {
            message.modelVersion = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.modelVersion !== undefined && (obj.modelVersion = message.modelVersion ? ModelVersion.toJSON(message.modelVersion) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetModelVersionResponse };
        if (object.modelVersion !== undefined && object.modelVersion !== null) {
            message.modelVersion = ModelVersion.fromPartial(object.modelVersion);
        }
        else {
            message.modelVersion = undefined;
        }
        return message;
    }
};
const baseQueryAllModelVersionRequest = {};
export const QueryAllModelVersionRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllModelVersionRequest };
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
        const message = { ...baseQueryAllModelVersionRequest };
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
        const message = { ...baseQueryAllModelVersionRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllModelVersionResponse = {};
export const QueryAllModelVersionResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.modelVersion) {
            ModelVersion.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllModelVersionResponse };
        message.modelVersion = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.modelVersion.push(ModelVersion.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllModelVersionResponse };
        message.modelVersion = [];
        if (object.modelVersion !== undefined && object.modelVersion !== null) {
            for (const e of object.modelVersion) {
                message.modelVersion.push(ModelVersion.fromJSON(e));
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
        if (message.modelVersion) {
            obj.modelVersion = message.modelVersion.map((e) => (e ? ModelVersion.toJSON(e) : undefined));
        }
        else {
            obj.modelVersion = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllModelVersionResponse };
        message.modelVersion = [];
        if (object.modelVersion !== undefined && object.modelVersion !== null) {
            for (const e of object.modelVersion) {
                message.modelVersion.push(ModelVersion.fromPartial(e));
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
const baseQueryGetModelVersionsRequest = { vid: 0, pid: 0 };
export const QueryGetModelVersionsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetModelVersionsRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.pid = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetModelVersionsRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetModelVersionsRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        return message;
    }
};
const baseQueryGetModelVersionsResponse = {};
export const QueryGetModelVersionsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.modelVersions !== undefined) {
            ModelVersions.encode(message.modelVersions, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetModelVersionsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.modelVersions = ModelVersions.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetModelVersionsResponse };
        if (object.modelVersions !== undefined && object.modelVersions !== null) {
            message.modelVersions = ModelVersions.fromJSON(object.modelVersions);
        }
        else {
            message.modelVersions = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.modelVersions !== undefined && (obj.modelVersions = message.modelVersions ? ModelVersions.toJSON(message.modelVersions) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetModelVersionsResponse };
        if (object.modelVersions !== undefined && object.modelVersions !== null) {
            message.modelVersions = ModelVersions.fromPartial(object.modelVersions);
        }
        else {
            message.modelVersions = undefined;
        }
        return message;
    }
};
const baseQueryAllModelVersionsRequest = {};
export const QueryAllModelVersionsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllModelVersionsRequest };
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
        const message = { ...baseQueryAllModelVersionsRequest };
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
        const message = { ...baseQueryAllModelVersionsRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllModelVersionsResponse = {};
export const QueryAllModelVersionsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.modelVersions) {
            ModelVersions.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllModelVersionsResponse };
        message.modelVersions = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.modelVersions.push(ModelVersions.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllModelVersionsResponse };
        message.modelVersions = [];
        if (object.modelVersions !== undefined && object.modelVersions !== null) {
            for (const e of object.modelVersions) {
                message.modelVersions.push(ModelVersions.fromJSON(e));
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
        if (message.modelVersions) {
            obj.modelVersions = message.modelVersions.map((e) => (e ? ModelVersions.toJSON(e) : undefined));
        }
        else {
            obj.modelVersions = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllModelVersionsResponse };
        message.modelVersions = [];
        if (object.modelVersions !== undefined && object.modelVersions !== null) {
            for (const e of object.modelVersions) {
                message.modelVersions.push(ModelVersions.fromPartial(e));
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
    VendorProducts(request) {
        const data = QueryGetVendorProductsRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'VendorProducts', data);
        return promise.then((data) => QueryGetVendorProductsResponse.decode(new Reader(data)));
    }
    VendorProductsAll(request) {
        const data = QueryAllVendorProductsRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'VendorProductsAll', data);
        return promise.then((data) => QueryAllVendorProductsResponse.decode(new Reader(data)));
    }
    Model(request) {
        const data = QueryGetModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'Model', data);
        return promise.then((data) => QueryGetModelResponse.decode(new Reader(data)));
    }
    ModelAll(request) {
        const data = QueryAllModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelAll', data);
        return promise.then((data) => QueryAllModelResponse.decode(new Reader(data)));
    }
    ModelVersion(request) {
        const data = QueryGetModelVersionRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersion', data);
        return promise.then((data) => QueryGetModelVersionResponse.decode(new Reader(data)));
    }
    ModelVersionAll(request) {
        const data = QueryAllModelVersionRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersionAll', data);
        return promise.then((data) => QueryAllModelVersionResponse.decode(new Reader(data)));
    }
    ModelVersions(request) {
        const data = QueryGetModelVersionsRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersions', data);
        return promise.then((data) => QueryGetModelVersionsResponse.decode(new Reader(data)));
    }
    ModelVersionsAll(request) {
        const data = QueryAllModelVersionsRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersionsAll', data);
        return promise.then((data) => QueryAllModelVersionsResponse.decode(new Reader(data)));
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
