import { Reader, Writer } from 'protobufjs/minimal';
import { VendorInfo } from '../vendorinfo/vendor_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
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
export declare const QueryGetVendorInfoRequest: {
    encode(message: QueryGetVendorInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoRequest;
    fromJSON(object: any): QueryGetVendorInfoRequest;
    toJSON(message: QueryGetVendorInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetVendorInfoRequest>): QueryGetVendorInfoRequest;
};
export declare const QueryGetVendorInfoResponse: {
    encode(message: QueryGetVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoResponse;
    fromJSON(object: any): QueryGetVendorInfoResponse;
    toJSON(message: QueryGetVendorInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetVendorInfoResponse>): QueryGetVendorInfoResponse;
};
export declare const QueryAllVendorInfoRequest: {
    encode(message: QueryAllVendorInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoRequest;
    fromJSON(object: any): QueryAllVendorInfoRequest;
    toJSON(message: QueryAllVendorInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllVendorInfoRequest>): QueryAllVendorInfoRequest;
};
export declare const QueryAllVendorInfoResponse: {
    encode(message: QueryAllVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoResponse;
    fromJSON(object: any): QueryAllVendorInfoResponse;
    toJSON(message: QueryAllVendorInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllVendorInfoResponse>): QueryAllVendorInfoResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a vendorInfo by index. */
    VendorInfo(request: QueryGetVendorInfoRequest): Promise<QueryGetVendorInfoResponse>;
    /** Queries a list of vendorInfo items. */
    VendorInfoAll(request: QueryAllVendorInfoRequest): Promise<QueryAllVendorInfoResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    VendorInfo(request: QueryGetVendorInfoRequest): Promise<QueryGetVendorInfoResponse>;
    VendorInfoAll(request: QueryAllVendorInfoRequest): Promise<QueryAllVendorInfoResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
