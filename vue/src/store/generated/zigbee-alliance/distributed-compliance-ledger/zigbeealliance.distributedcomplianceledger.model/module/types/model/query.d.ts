import { Reader, Writer } from 'protobufjs/minimal';
import { VendorProducts } from '../model/vendor_products';
import { Model } from '../model/model';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { ModelVersion } from '../model/model_version';
import { ModelVersions } from '../model/model_versions';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
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
export declare const QueryGetVendorProductsRequest: {
    encode(message: QueryGetVendorProductsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetVendorProductsRequest;
    fromJSON(object: any): QueryGetVendorProductsRequest;
    toJSON(message: QueryGetVendorProductsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetVendorProductsRequest>): QueryGetVendorProductsRequest;
};
export declare const QueryGetVendorProductsResponse: {
    encode(message: QueryGetVendorProductsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetVendorProductsResponse;
    fromJSON(object: any): QueryGetVendorProductsResponse;
    toJSON(message: QueryGetVendorProductsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetVendorProductsResponse>): QueryGetVendorProductsResponse;
};
export declare const QueryGetModelRequest: {
    encode(message: QueryGetModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetModelRequest;
    fromJSON(object: any): QueryGetModelRequest;
    toJSON(message: QueryGetModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetModelRequest>): QueryGetModelRequest;
};
export declare const QueryGetModelResponse: {
    encode(message: QueryGetModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetModelResponse;
    fromJSON(object: any): QueryGetModelResponse;
    toJSON(message: QueryGetModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetModelResponse>): QueryGetModelResponse;
};
export declare const QueryAllModelRequest: {
    encode(message: QueryAllModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllModelRequest;
    fromJSON(object: any): QueryAllModelRequest;
    toJSON(message: QueryAllModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllModelRequest>): QueryAllModelRequest;
};
export declare const QueryAllModelResponse: {
    encode(message: QueryAllModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllModelResponse;
    fromJSON(object: any): QueryAllModelResponse;
    toJSON(message: QueryAllModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllModelResponse>): QueryAllModelResponse;
};
export declare const QueryGetModelVersionRequest: {
    encode(message: QueryGetModelVersionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionRequest;
    fromJSON(object: any): QueryGetModelVersionRequest;
    toJSON(message: QueryGetModelVersionRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetModelVersionRequest>): QueryGetModelVersionRequest;
};
export declare const QueryGetModelVersionResponse: {
    encode(message: QueryGetModelVersionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionResponse;
    fromJSON(object: any): QueryGetModelVersionResponse;
    toJSON(message: QueryGetModelVersionResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetModelVersionResponse>): QueryGetModelVersionResponse;
};
export declare const QueryGetModelVersionsRequest: {
    encode(message: QueryGetModelVersionsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionsRequest;
    fromJSON(object: any): QueryGetModelVersionsRequest;
    toJSON(message: QueryGetModelVersionsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetModelVersionsRequest>): QueryGetModelVersionsRequest;
};
export declare const QueryGetModelVersionsResponse: {
    encode(message: QueryGetModelVersionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionsResponse;
    fromJSON(object: any): QueryGetModelVersionsResponse;
    toJSON(message: QueryGetModelVersionsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetModelVersionsResponse>): QueryGetModelVersionsResponse;
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
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    VendorProducts(request: QueryGetVendorProductsRequest): Promise<QueryGetVendorProductsResponse>;
    Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse>;
    ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse>;
    ModelVersion(request: QueryGetModelVersionRequest): Promise<QueryGetModelVersionResponse>;
    ModelVersions(request: QueryGetModelVersionsRequest): Promise<QueryGetModelVersionsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
