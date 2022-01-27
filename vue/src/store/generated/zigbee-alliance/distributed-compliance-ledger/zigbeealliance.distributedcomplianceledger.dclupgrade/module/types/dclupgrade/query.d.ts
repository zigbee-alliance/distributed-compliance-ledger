import { Reader, Writer } from 'protobufjs/minimal';
import { Params } from '../dclupgrade/params';
import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclupgrade";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetProposedUpgradeRequest {
    name: string;
}
export interface QueryGetProposedUpgradeResponse {
    proposedUpgrade: ProposedUpgrade | undefined;
}
export interface QueryAllProposedUpgradeRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllProposedUpgradeResponse {
    proposedUpgrade: ProposedUpgrade[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetProposedUpgradeRequest: {
    encode(message: QueryGetProposedUpgradeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedUpgradeRequest;
    fromJSON(object: any): QueryGetProposedUpgradeRequest;
    toJSON(message: QueryGetProposedUpgradeRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedUpgradeRequest>): QueryGetProposedUpgradeRequest;
};
export declare const QueryGetProposedUpgradeResponse: {
    encode(message: QueryGetProposedUpgradeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedUpgradeResponse;
    fromJSON(object: any): QueryGetProposedUpgradeResponse;
    toJSON(message: QueryGetProposedUpgradeResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedUpgradeResponse>): QueryGetProposedUpgradeResponse;
};
export declare const QueryAllProposedUpgradeRequest: {
    encode(message: QueryAllProposedUpgradeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedUpgradeRequest;
    fromJSON(object: any): QueryAllProposedUpgradeRequest;
    toJSON(message: QueryAllProposedUpgradeRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedUpgradeRequest>): QueryAllProposedUpgradeRequest;
};
export declare const QueryAllProposedUpgradeResponse: {
    encode(message: QueryAllProposedUpgradeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedUpgradeResponse;
    fromJSON(object: any): QueryAllProposedUpgradeResponse;
    toJSON(message: QueryAllProposedUpgradeResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedUpgradeResponse>): QueryAllProposedUpgradeResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a ProposedUpgrade by index. */
    ProposedUpgrade(request: QueryGetProposedUpgradeRequest): Promise<QueryGetProposedUpgradeResponse>;
    /** Queries a list of ProposedUpgrade items. */
    ProposedUpgradeAll(request: QueryAllProposedUpgradeRequest): Promise<QueryAllProposedUpgradeResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    ProposedUpgrade(request: QueryGetProposedUpgradeRequest): Promise<QueryGetProposedUpgradeResponse>;
    ProposedUpgradeAll(request: QueryAllProposedUpgradeRequest): Promise<QueryAllProposedUpgradeResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
