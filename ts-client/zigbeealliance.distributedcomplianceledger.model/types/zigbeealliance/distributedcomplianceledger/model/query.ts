/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { Model } from "./model";
import { ModelVersion } from "./model_version";
import { ModelVersions } from "./model_versions";
import { VendorProducts } from "./vendor_products";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";

export interface QueryGetVendorProductsRequest {
  vid: number;
}

export interface QueryGetVendorProductsResponse {
  vendorProducts: VendorProducts | undefined;
}

export interface QueryGetModelRequest {
  vid: number;
  pid: number;
}

export interface QueryGetModelResponse {
  model: Model | undefined;
}

export interface QueryAllModelRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllModelResponse {
  model: Model[];
  pagination: PageResponse | undefined;
}

export interface QueryGetModelVersionRequest {
  vid: number;
  pid: number;
  softwareVersion: number;
}

export interface QueryGetModelVersionResponse {
  modelVersion: ModelVersion | undefined;
}

export interface QueryGetModelVersionsRequest {
  vid: number;
  pid: number;
}

export interface QueryGetModelVersionsResponse {
  modelVersions: ModelVersions | undefined;
}

function createBaseQueryGetVendorProductsRequest(): QueryGetVendorProductsRequest {
  return { vid: 0 };
}

export const QueryGetVendorProductsRequest = {
  encode(message: QueryGetVendorProductsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVendorProductsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVendorProductsRequest();
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

  fromJSON(object: any): QueryGetVendorProductsRequest {
    return { vid: isSet(object.vid) ? Number(object.vid) : 0 };
  },

  toJSON(message: QueryGetVendorProductsRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVendorProductsRequest>, I>>(
    object: I,
  ): QueryGetVendorProductsRequest {
    const message = createBaseQueryGetVendorProductsRequest();
    message.vid = object.vid ?? 0;
    return message;
  },
};

function createBaseQueryGetVendorProductsResponse(): QueryGetVendorProductsResponse {
  return { vendorProducts: undefined };
}

export const QueryGetVendorProductsResponse = {
  encode(message: QueryGetVendorProductsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vendorProducts !== undefined) {
      VendorProducts.encode(message.vendorProducts, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVendorProductsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVendorProductsResponse();
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

  fromJSON(object: any): QueryGetVendorProductsResponse {
    return {
      vendorProducts: isSet(object.vendorProducts) ? VendorProducts.fromJSON(object.vendorProducts) : undefined,
    };
  },

  toJSON(message: QueryGetVendorProductsResponse): unknown {
    const obj: any = {};
    message.vendorProducts !== undefined
      && (obj.vendorProducts = message.vendorProducts ? VendorProducts.toJSON(message.vendorProducts) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVendorProductsResponse>, I>>(
    object: I,
  ): QueryGetVendorProductsResponse {
    const message = createBaseQueryGetVendorProductsResponse();
    message.vendorProducts = (object.vendorProducts !== undefined && object.vendorProducts !== null)
      ? VendorProducts.fromPartial(object.vendorProducts)
      : undefined;
    return message;
  },
};

function createBaseQueryGetModelRequest(): QueryGetModelRequest {
  return { vid: 0, pid: 0 };
}

export const QueryGetModelRequest = {
  encode(message: QueryGetModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelRequest();
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

  fromJSON(object: any): QueryGetModelRequest {
    return { vid: isSet(object.vid) ? Number(object.vid) : 0, pid: isSet(object.pid) ? Number(object.pid) : 0 };
  },

  toJSON(message: QueryGetModelRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelRequest>, I>>(object: I): QueryGetModelRequest {
    const message = createBaseQueryGetModelRequest();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    return message;
  },
};

function createBaseQueryGetModelResponse(): QueryGetModelResponse {
  return { model: undefined };
}

export const QueryGetModelResponse = {
  encode(message: QueryGetModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.model !== undefined) {
      Model.encode(message.model, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelResponse();
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

  fromJSON(object: any): QueryGetModelResponse {
    return { model: isSet(object.model) ? Model.fromJSON(object.model) : undefined };
  },

  toJSON(message: QueryGetModelResponse): unknown {
    const obj: any = {};
    message.model !== undefined && (obj.model = message.model ? Model.toJSON(message.model) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelResponse>, I>>(object: I): QueryGetModelResponse {
    const message = createBaseQueryGetModelResponse();
    message.model = (object.model !== undefined && object.model !== null) ? Model.fromPartial(object.model) : undefined;
    return message;
  },
};

function createBaseQueryAllModelRequest(): QueryAllModelRequest {
  return { pagination: undefined };
}

export const QueryAllModelRequest = {
  encode(message: QueryAllModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllModelRequest();
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

  fromJSON(object: any): QueryAllModelRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllModelRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllModelRequest>, I>>(object: I): QueryAllModelRequest {
    const message = createBaseQueryAllModelRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllModelResponse(): QueryAllModelResponse {
  return { model: [], pagination: undefined };
}

export const QueryAllModelResponse = {
  encode(message: QueryAllModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.model) {
      Model.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllModelResponse();
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

  fromJSON(object: any): QueryAllModelResponse {
    return {
      model: Array.isArray(object?.model) ? object.model.map((e: any) => Model.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllModelResponse): unknown {
    const obj: any = {};
    if (message.model) {
      obj.model = message.model.map((e) => e ? Model.toJSON(e) : undefined);
    } else {
      obj.model = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllModelResponse>, I>>(object: I): QueryAllModelResponse {
    const message = createBaseQueryAllModelResponse();
    message.model = object.model?.map((e) => Model.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetModelVersionRequest(): QueryGetModelVersionRequest {
  return { vid: 0, pid: 0, softwareVersion: 0 };
}

export const QueryGetModelVersionRequest = {
  encode(message: QueryGetModelVersionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelVersionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelVersionRequest();
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
          message.softwareVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetModelVersionRequest {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
    };
  },

  toJSON(message: QueryGetModelVersionRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelVersionRequest>, I>>(object: I): QueryGetModelVersionRequest {
    const message = createBaseQueryGetModelVersionRequest();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    return message;
  },
};

function createBaseQueryGetModelVersionResponse(): QueryGetModelVersionResponse {
  return { modelVersion: undefined };
}

export const QueryGetModelVersionResponse = {
  encode(message: QueryGetModelVersionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.modelVersion !== undefined) {
      ModelVersion.encode(message.modelVersion, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelVersionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelVersionResponse();
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

  fromJSON(object: any): QueryGetModelVersionResponse {
    return { modelVersion: isSet(object.modelVersion) ? ModelVersion.fromJSON(object.modelVersion) : undefined };
  },

  toJSON(message: QueryGetModelVersionResponse): unknown {
    const obj: any = {};
    message.modelVersion !== undefined
      && (obj.modelVersion = message.modelVersion ? ModelVersion.toJSON(message.modelVersion) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelVersionResponse>, I>>(object: I): QueryGetModelVersionResponse {
    const message = createBaseQueryGetModelVersionResponse();
    message.modelVersion = (object.modelVersion !== undefined && object.modelVersion !== null)
      ? ModelVersion.fromPartial(object.modelVersion)
      : undefined;
    return message;
  },
};

function createBaseQueryGetModelVersionsRequest(): QueryGetModelVersionsRequest {
  return { vid: 0, pid: 0 };
}

export const QueryGetModelVersionsRequest = {
  encode(message: QueryGetModelVersionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelVersionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelVersionsRequest();
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

  fromJSON(object: any): QueryGetModelVersionsRequest {
    return { vid: isSet(object.vid) ? Number(object.vid) : 0, pid: isSet(object.pid) ? Number(object.pid) : 0 };
  },

  toJSON(message: QueryGetModelVersionsRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelVersionsRequest>, I>>(object: I): QueryGetModelVersionsRequest {
    const message = createBaseQueryGetModelVersionsRequest();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    return message;
  },
};

function createBaseQueryGetModelVersionsResponse(): QueryGetModelVersionsResponse {
  return { modelVersions: undefined };
}

export const QueryGetModelVersionsResponse = {
  encode(message: QueryGetModelVersionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.modelVersions !== undefined) {
      ModelVersions.encode(message.modelVersions, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetModelVersionsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetModelVersionsResponse();
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

  fromJSON(object: any): QueryGetModelVersionsResponse {
    return { modelVersions: isSet(object.modelVersions) ? ModelVersions.fromJSON(object.modelVersions) : undefined };
  },

  toJSON(message: QueryGetModelVersionsResponse): unknown {
    const obj: any = {};
    message.modelVersions !== undefined
      && (obj.modelVersions = message.modelVersions ? ModelVersions.toJSON(message.modelVersions) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetModelVersionsResponse>, I>>(
    object: I,
  ): QueryGetModelVersionsResponse {
    const message = createBaseQueryGetModelVersionsResponse();
    message.modelVersions = (object.modelVersions !== undefined && object.modelVersions !== null)
      ? ModelVersions.fromPartial(object.modelVersions)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries VendorProducts by index. */
  VendorProducts(request: QueryGetVendorProductsRequest): Promise<QueryGetVendorProductsResponse>;
  /** Queries a Model by index. */
  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse>;
  /** Queries a list of all Model items. */
  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse>;
  /** Queries a ModelVersion by index. */
  ModelVersion(request: QueryGetModelVersionRequest): Promise<QueryGetModelVersionResponse>;
  /** Queries ModelVersions by index. */
  ModelVersions(request: QueryGetModelVersionsRequest): Promise<QueryGetModelVersionsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.VendorProducts = this.VendorProducts.bind(this);
    this.Model = this.Model.bind(this);
    this.ModelAll = this.ModelAll.bind(this);
    this.ModelVersion = this.ModelVersion.bind(this);
    this.ModelVersions = this.ModelVersions.bind(this);
  }
  VendorProducts(request: QueryGetVendorProductsRequest): Promise<QueryGetVendorProductsResponse> {
    const data = QueryGetVendorProductsRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.model.Query", "VendorProducts", data);
    return promise.then((data) => QueryGetVendorProductsResponse.decode(new _m0.Reader(data)));
  }

  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse> {
    const data = QueryGetModelRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.model.Query", "Model", data);
    return promise.then((data) => QueryGetModelResponse.decode(new _m0.Reader(data)));
  }

  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse> {
    const data = QueryAllModelRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.model.Query", "ModelAll", data);
    return promise.then((data) => QueryAllModelResponse.decode(new _m0.Reader(data)));
  }

  ModelVersion(request: QueryGetModelVersionRequest): Promise<QueryGetModelVersionResponse> {
    const data = QueryGetModelVersionRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.model.Query", "ModelVersion", data);
    return promise.then((data) => QueryGetModelVersionResponse.decode(new _m0.Reader(data)));
  }

  ModelVersions(request: QueryGetModelVersionsRequest): Promise<QueryGetModelVersionsResponse> {
    const data = QueryGetModelVersionsRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.model.Query", "ModelVersions", data);
    return promise.then((data) => QueryGetModelVersionsResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
