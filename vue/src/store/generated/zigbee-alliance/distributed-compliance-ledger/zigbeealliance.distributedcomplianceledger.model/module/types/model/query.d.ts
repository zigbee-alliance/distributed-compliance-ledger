import { Reader, Writer } from 'protobufjs/minimal';
import { VendorProducts } from '../model/vendor_products';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { Model } from '../model/model';
import { ModelVersion } from '../model/model_version';
import { ModelVersions } from '../model/model_versions';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
export interface QueryGetVendorProductsRequest {
    vid: number;
}
export interface QueryGetVendorProductsResponse {
    vendorProducts: VendorProducts | undefined;
}
export interface QueryAllVendorProductsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllVendorProductsResponse {
    vendorProducts: VendorProducts[];
    pagination: PageResponse | undefined;
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
export interface QueryAllModelVersionRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllModelVersionResponse {
    modelVersion: ModelVersion[];
    pagination: PageResponse | undefined;
}
export interface QueryGetModelVersionsRequest {
    vid: number;
    pid: number;
}
export interface QueryGetModelVersionsResponse {
    modelVersions: ModelVersions | undefined;
}
export interface QueryAllModelVersionsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllModelVersionsResponse {
    modelVersions: ModelVersions[];
    pagination: PageResponse | undefined;
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
export declare const QueryAllVendorProductsRequest: {
    encode(message: QueryAllVendorProductsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllVendorProductsRequest;
    fromJSON(object: any): QueryAllVendorProductsRequest;
    toJSON(message: QueryAllVendorProductsRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllVendorProductsRequest>): QueryAllVendorProductsRequest;
};
export declare const QueryAllVendorProductsResponse: {
    encode(message: QueryAllVendorProductsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllVendorProductsResponse;
    fromJSON(object: any): QueryAllVendorProductsResponse;
    toJSON(message: QueryAllVendorProductsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllVendorProductsResponse>): QueryAllVendorProductsResponse;
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
export declare const QueryAllModelVersionRequest: {
    encode(message: QueryAllModelVersionRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionRequest;
    fromJSON(object: any): QueryAllModelVersionRequest;
    toJSON(message: QueryAllModelVersionRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllModelVersionRequest>): QueryAllModelVersionRequest;
};
export declare const QueryAllModelVersionResponse: {
    encode(message: QueryAllModelVersionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionResponse;
    fromJSON(object: any): QueryAllModelVersionResponse;
    toJSON(message: QueryAllModelVersionResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllModelVersionResponse>): QueryAllModelVersionResponse;
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
export declare const QueryAllModelVersionsRequest: {
    encode(message: QueryAllModelVersionsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionsRequest;
    fromJSON(object: any): QueryAllModelVersionsRequest;
    toJSON(message: QueryAllModelVersionsRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllModelVersionsRequest>): QueryAllModelVersionsRequest;
};
export declare const QueryAllModelVersionsResponse: {
    encode(message: QueryAllModelVersionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionsResponse;
    fromJSON(object: any): QueryAllModelVersionsResponse;
    toJSON(message: QueryAllModelVersionsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllModelVersionsResponse>): QueryAllModelVersionsResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a VendorProducts by index. */
    VendorProducts(request: QueryGetVendorProductsRequest): Promise<QueryGetVendorProductsResponse>;
    /** Queries a list of VendorProducts items. */
    VendorProductsAll(request: QueryAllVendorProductsRequest): Promise<QueryAllVendorProductsResponse>;
    /** Queries a Model by index. */
    Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse>;
    /** Queries a list of Model items. */
    ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse>;
    /** Queries a ModelVersion by index. */
    ModelVersion(request: QueryGetModelVersionRequest): Promise<QueryGetModelVersionResponse>;
    /** Queries a list of ModelVersion items. */
    ModelVersionAll(request: QueryAllModelVersionRequest): Promise<QueryAllModelVersionResponse>;
    /** Queries a ModelVersions by index. */
    ModelVersions(request: QueryGetModelVersionsRequest): Promise<QueryGetModelVersionsResponse>;
    /** Queries a list of ModelVersions items. */
    ModelVersionsAll(request: QueryAllModelVersionsRequest): Promise<QueryAllModelVersionsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    VendorProducts(request: QueryGetVendorProductsRequest): Promise<QueryGetVendorProductsResponse>;
    VendorProductsAll(request: QueryAllVendorProductsRequest): Promise<QueryAllVendorProductsResponse>;
    Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse>;
    ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse>;
    ModelVersion(request: QueryGetModelVersionRequest): Promise<QueryGetModelVersionResponse>;
    ModelVersionAll(request: QueryAllModelVersionRequest): Promise<QueryAllModelVersionResponse>;
    ModelVersions(request: QueryGetModelVersionsRequest): Promise<QueryGetModelVersionsResponse>;
    ModelVersionsAll(request: QueryAllModelVersionsRequest): Promise<QueryAllModelVersionsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
