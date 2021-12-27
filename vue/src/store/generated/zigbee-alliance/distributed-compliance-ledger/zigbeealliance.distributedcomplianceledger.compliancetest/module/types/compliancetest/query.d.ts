import { Reader, Writer } from 'protobufjs/minimal';
import { TestingResults } from '../compliancetest/testing_results';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliancetest";
export interface QueryGetTestingResultsRequest {
    vid: number;
    pid: number;
    softwareVersion: number;
}
export interface QueryGetTestingResultsResponse {
    testingResults: TestingResults | undefined;
}
export interface QueryAllTestingResultsRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllTestingResultsResponse {
    testingResults: TestingResults[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetTestingResultsRequest: {
    encode(message: QueryGetTestingResultsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTestingResultsRequest;
    fromJSON(object: any): QueryGetTestingResultsRequest;
    toJSON(message: QueryGetTestingResultsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTestingResultsRequest>): QueryGetTestingResultsRequest;
};
export declare const QueryGetTestingResultsResponse: {
    encode(message: QueryGetTestingResultsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTestingResultsResponse;
    fromJSON(object: any): QueryGetTestingResultsResponse;
    toJSON(message: QueryGetTestingResultsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTestingResultsResponse>): QueryGetTestingResultsResponse;
};
export declare const QueryAllTestingResultsRequest: {
    encode(message: QueryAllTestingResultsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTestingResultsRequest;
    fromJSON(object: any): QueryAllTestingResultsRequest;
    toJSON(message: QueryAllTestingResultsRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllTestingResultsRequest>): QueryAllTestingResultsRequest;
};
export declare const QueryAllTestingResultsResponse: {
    encode(message: QueryAllTestingResultsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTestingResultsResponse;
    fromJSON(object: any): QueryAllTestingResultsResponse;
    toJSON(message: QueryAllTestingResultsResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllTestingResultsResponse>): QueryAllTestingResultsResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a TestingResults by index. */
    TestingResults(request: QueryGetTestingResultsRequest): Promise<QueryGetTestingResultsResponse>;
    /** Queries a list of TestingResults items. */
    TestingResultsAll(request: QueryAllTestingResultsRequest): Promise<QueryAllTestingResultsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    TestingResults(request: QueryGetTestingResultsRequest): Promise<QueryGetTestingResultsResponse>;
    TestingResultsAll(request: QueryAllTestingResultsRequest): Promise<QueryAllTestingResultsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
