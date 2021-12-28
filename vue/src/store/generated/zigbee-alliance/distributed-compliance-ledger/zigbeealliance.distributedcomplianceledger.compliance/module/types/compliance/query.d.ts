import { Reader, Writer } from 'protobufjs/minimal';
import { ComplianceInfo } from '../compliance/compliance_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface QueryGetComplianceInfoRequest {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
}
export interface QueryGetComplianceInfoResponse {
    complianceInfo: ComplianceInfo | undefined;
}
export interface QueryAllComplianceInfoRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllComplianceInfoResponse {
    complianceInfo: ComplianceInfo[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetComplianceInfoRequest: {
    encode(message: QueryGetComplianceInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoRequest;
    fromJSON(object: any): QueryGetComplianceInfoRequest;
    toJSON(message: QueryGetComplianceInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetComplianceInfoRequest>): QueryGetComplianceInfoRequest;
};
export declare const QueryGetComplianceInfoResponse: {
    encode(message: QueryGetComplianceInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoResponse;
    fromJSON(object: any): QueryGetComplianceInfoResponse;
    toJSON(message: QueryGetComplianceInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetComplianceInfoResponse>): QueryGetComplianceInfoResponse;
};
export declare const QueryAllComplianceInfoRequest: {
    encode(message: QueryAllComplianceInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoRequest;
    fromJSON(object: any): QueryAllComplianceInfoRequest;
    toJSON(message: QueryAllComplianceInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllComplianceInfoRequest>): QueryAllComplianceInfoRequest;
};
export declare const QueryAllComplianceInfoResponse: {
    encode(message: QueryAllComplianceInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoResponse;
    fromJSON(object: any): QueryAllComplianceInfoResponse;
    toJSON(message: QueryAllComplianceInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllComplianceInfoResponse>): QueryAllComplianceInfoResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a ComplianceInfo by index. */
    ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse>;
    /** Queries a list of ComplianceInfo items. */
    ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse>;
    ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
