import { Reader, Writer } from 'protobufjs/minimal';
import { NewVendorInfo } from '../vendorinfo/new_vendor_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface QueryGetNewVendorInfoRequest {
    index: string;
}
export interface QueryGetNewVendorInfoResponse {
    newVendorInfo: NewVendorInfo | undefined;
}
export interface QueryAllNewVendorInfoRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllNewVendorInfoResponse {
    newVendorInfo: NewVendorInfo[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetNewVendorInfoRequest: {
    encode(message: QueryGetNewVendorInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNewVendorInfoRequest;
    fromJSON(object: any): QueryGetNewVendorInfoRequest;
    toJSON(message: QueryGetNewVendorInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetNewVendorInfoRequest>): QueryGetNewVendorInfoRequest;
};
export declare const QueryGetNewVendorInfoResponse: {
    encode(message: QueryGetNewVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNewVendorInfoResponse;
    fromJSON(object: any): QueryGetNewVendorInfoResponse;
    toJSON(message: QueryGetNewVendorInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetNewVendorInfoResponse>): QueryGetNewVendorInfoResponse;
};
export declare const QueryAllNewVendorInfoRequest: {
    encode(message: QueryAllNewVendorInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNewVendorInfoRequest;
    fromJSON(object: any): QueryAllNewVendorInfoRequest;
    toJSON(message: QueryAllNewVendorInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllNewVendorInfoRequest>): QueryAllNewVendorInfoRequest;
};
export declare const QueryAllNewVendorInfoResponse: {
    encode(message: QueryAllNewVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNewVendorInfoResponse;
    fromJSON(object: any): QueryAllNewVendorInfoResponse;
    toJSON(message: QueryAllNewVendorInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllNewVendorInfoResponse>): QueryAllNewVendorInfoResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a newVendorInfo by index. */
    NewVendorInfo(request: QueryGetNewVendorInfoRequest): Promise<QueryGetNewVendorInfoResponse>;
    /** Queries a list of newVendorInfo items. */
    NewVendorInfoAll(request: QueryAllNewVendorInfoRequest): Promise<QueryAllNewVendorInfoResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    NewVendorInfo(request: QueryGetNewVendorInfoRequest): Promise<QueryGetNewVendorInfoResponse>;
    NewVendorInfoAll(request: QueryAllNewVendorInfoRequest): Promise<QueryAllNewVendorInfoResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
