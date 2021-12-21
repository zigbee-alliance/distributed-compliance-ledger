import { Reader, Writer } from 'protobufjs/minimal';
import { VendorInfoType } from '../vendorinfo/vendor_info_type';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface QueryGetVendorInfoTypeRequest {
    vendorID: number;
}
export interface QueryGetVendorInfoTypeResponse {
    vendorInfoType: VendorInfoType | undefined;
}
export interface QueryAllVendorInfoTypeRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllVendorInfoTypeResponse {
    vendorInfoType: VendorInfoType[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetVendorInfoTypeRequest: {
    encode(message: QueryGetVendorInfoTypeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoTypeRequest;
    fromJSON(object: any): QueryGetVendorInfoTypeRequest;
    toJSON(message: QueryGetVendorInfoTypeRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetVendorInfoTypeRequest>): QueryGetVendorInfoTypeRequest;
};
export declare const QueryGetVendorInfoTypeResponse: {
    encode(message: QueryGetVendorInfoTypeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoTypeResponse;
    fromJSON(object: any): QueryGetVendorInfoTypeResponse;
    toJSON(message: QueryGetVendorInfoTypeResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetVendorInfoTypeResponse>): QueryGetVendorInfoTypeResponse;
};
export declare const QueryAllVendorInfoTypeRequest: {
    encode(message: QueryAllVendorInfoTypeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoTypeRequest;
    fromJSON(object: any): QueryAllVendorInfoTypeRequest;
    toJSON(message: QueryAllVendorInfoTypeRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllVendorInfoTypeRequest>): QueryAllVendorInfoTypeRequest;
};
export declare const QueryAllVendorInfoTypeResponse: {
    encode(message: QueryAllVendorInfoTypeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoTypeResponse;
    fromJSON(object: any): QueryAllVendorInfoTypeResponse;
    toJSON(message: QueryAllVendorInfoTypeResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllVendorInfoTypeResponse>): QueryAllVendorInfoTypeResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a vendorInfoType by index. */
    VendorInfoType(request: QueryGetVendorInfoTypeRequest): Promise<QueryGetVendorInfoTypeResponse>;
    /** Queries a list of vendorInfoType items. */
    VendorInfoTypeAll(request: QueryAllVendorInfoTypeRequest): Promise<QueryAllVendorInfoTypeResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    VendorInfoType(request: QueryGetVendorInfoTypeRequest): Promise<QueryGetVendorInfoTypeResponse>;
    VendorInfoTypeAll(request: QueryAllVendorInfoTypeRequest): Promise<QueryAllVendorInfoTypeResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
