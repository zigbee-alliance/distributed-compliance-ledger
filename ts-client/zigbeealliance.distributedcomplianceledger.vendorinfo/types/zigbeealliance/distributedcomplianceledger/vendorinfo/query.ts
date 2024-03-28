/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { VendorInfo } from "./vendor_info";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";

export interface QueryGetVendorInfoRequest {
  vendorID: number;
}

export interface QueryGetVendorInfoResponse {
  vendorInfo: VendorInfo | undefined;
}

export interface QueryAllVendorInfoRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllVendorInfoResponse {
  vendorInfo: VendorInfo[];
  pagination: PageResponse | undefined;
}

function createBaseQueryGetVendorInfoRequest(): QueryGetVendorInfoRequest {
  return { vendorID: 0 };
}

export const QueryGetVendorInfoRequest = {
  encode(message: QueryGetVendorInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vendorID !== 0) {
      writer.uint32(8).int32(message.vendorID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVendorInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVendorInfoRequest();
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

  fromJSON(object: any): QueryGetVendorInfoRequest {
    return { vendorID: isSet(object.vendorID) ? Number(object.vendorID) : 0 };
  },

  toJSON(message: QueryGetVendorInfoRequest): unknown {
    const obj: any = {};
    message.vendorID !== undefined && (obj.vendorID = Math.round(message.vendorID));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVendorInfoRequest>, I>>(object: I): QueryGetVendorInfoRequest {
    const message = createBaseQueryGetVendorInfoRequest();
    message.vendorID = object.vendorID ?? 0;
    return message;
  },
};

function createBaseQueryGetVendorInfoResponse(): QueryGetVendorInfoResponse {
  return { vendorInfo: undefined };
}

export const QueryGetVendorInfoResponse = {
  encode(message: QueryGetVendorInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vendorInfo !== undefined) {
      VendorInfo.encode(message.vendorInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetVendorInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetVendorInfoResponse();
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

  fromJSON(object: any): QueryGetVendorInfoResponse {
    return { vendorInfo: isSet(object.vendorInfo) ? VendorInfo.fromJSON(object.vendorInfo) : undefined };
  },

  toJSON(message: QueryGetVendorInfoResponse): unknown {
    const obj: any = {};
    message.vendorInfo !== undefined
      && (obj.vendorInfo = message.vendorInfo ? VendorInfo.toJSON(message.vendorInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetVendorInfoResponse>, I>>(object: I): QueryGetVendorInfoResponse {
    const message = createBaseQueryGetVendorInfoResponse();
    message.vendorInfo = (object.vendorInfo !== undefined && object.vendorInfo !== null)
      ? VendorInfo.fromPartial(object.vendorInfo)
      : undefined;
    return message;
  },
};

function createBaseQueryAllVendorInfoRequest(): QueryAllVendorInfoRequest {
  return { pagination: undefined };
}

export const QueryAllVendorInfoRequest = {
  encode(message: QueryAllVendorInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllVendorInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllVendorInfoRequest();
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

  fromJSON(object: any): QueryAllVendorInfoRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllVendorInfoRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllVendorInfoRequest>, I>>(object: I): QueryAllVendorInfoRequest {
    const message = createBaseQueryAllVendorInfoRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllVendorInfoResponse(): QueryAllVendorInfoResponse {
  return { vendorInfo: [], pagination: undefined };
}

export const QueryAllVendorInfoResponse = {
  encode(message: QueryAllVendorInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.vendorInfo) {
      VendorInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllVendorInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllVendorInfoResponse();
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

  fromJSON(object: any): QueryAllVendorInfoResponse {
    return {
      vendorInfo: Array.isArray(object?.vendorInfo) ? object.vendorInfo.map((e: any) => VendorInfo.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllVendorInfoResponse): unknown {
    const obj: any = {};
    if (message.vendorInfo) {
      obj.vendorInfo = message.vendorInfo.map((e) => e ? VendorInfo.toJSON(e) : undefined);
    } else {
      obj.vendorInfo = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllVendorInfoResponse>, I>>(object: I): QueryAllVendorInfoResponse {
    const message = createBaseQueryAllVendorInfoResponse();
    message.vendorInfo = object.vendorInfo?.map((e) => VendorInfo.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a vendorInfo by index. */
  VendorInfo(request: QueryGetVendorInfoRequest): Promise<QueryGetVendorInfoResponse>;
  /** Queries a list of vendorInfo items. */
  VendorInfoAll(request: QueryAllVendorInfoRequest): Promise<QueryAllVendorInfoResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.VendorInfo = this.VendorInfo.bind(this);
    this.VendorInfoAll = this.VendorInfoAll.bind(this);
  }
  VendorInfo(request: QueryGetVendorInfoRequest): Promise<QueryGetVendorInfoResponse> {
    const data = QueryGetVendorInfoRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.vendorinfo.Query", "VendorInfo", data);
    return promise.then((data) => QueryGetVendorInfoResponse.decode(new _m0.Reader(data)));
  }

  VendorInfoAll(request: QueryAllVendorInfoRequest): Promise<QueryAllVendorInfoResponse> {
    const data = QueryAllVendorInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.vendorinfo.Query",
      "VendorInfoAll",
      data,
    );
    return promise.then((data) => QueryAllVendorInfoResponse.decode(new _m0.Reader(data)));
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
