import { Reader, Writer } from 'protobufjs/minimal';
import { Account } from '../dclauth/account';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { PendingAccount } from '../dclauth/pending_account';
import { PendingAccountRevocation } from '../dclauth/pending_account_revocation';
import { AccountStat } from '../dclauth/account_stat';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface QueryGetAccountRequest {
    address: string;
}
export interface QueryGetAccountResponse {
    account: Account | undefined;
}
export interface QueryAllAccountRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllAccountResponse {
    account: Account[];
    pagination: PageResponse | undefined;
}
export interface QueryGetPendingAccountRequest {
    address: string;
}
export interface QueryGetPendingAccountResponse {
    pendingAccount: PendingAccount | undefined;
}
export interface QueryAllPendingAccountRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllPendingAccountResponse {
    pendingAccount: PendingAccount[];
    pagination: PageResponse | undefined;
}
export interface QueryGetPendingAccountRevocationRequest {
    address: string;
}
export interface QueryGetPendingAccountRevocationResponse {
    pendingAccountRevocation: PendingAccountRevocation | undefined;
}
export interface QueryAllPendingAccountRevocationRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllPendingAccountRevocationResponse {
    pendingAccountRevocation: PendingAccountRevocation[];
    pagination: PageResponse | undefined;
}
export interface QueryGetAccountStatRequest {
}
export interface QueryGetAccountStatResponse {
    AccountStat: AccountStat | undefined;
}
export declare const QueryGetAccountRequest: {
    encode(message: QueryGetAccountRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountRequest;
    fromJSON(object: any): QueryGetAccountRequest;
    toJSON(message: QueryGetAccountRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAccountRequest>): QueryGetAccountRequest;
};
export declare const QueryGetAccountResponse: {
    encode(message: QueryGetAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountResponse;
    fromJSON(object: any): QueryGetAccountResponse;
    toJSON(message: QueryGetAccountResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAccountResponse>): QueryGetAccountResponse;
};
export declare const QueryAllAccountRequest: {
    encode(message: QueryAllAccountRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAccountRequest;
    fromJSON(object: any): QueryAllAccountRequest;
    toJSON(message: QueryAllAccountRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllAccountRequest>): QueryAllAccountRequest;
};
export declare const QueryAllAccountResponse: {
    encode(message: QueryAllAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAccountResponse;
    fromJSON(object: any): QueryAllAccountResponse;
    toJSON(message: QueryAllAccountResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllAccountResponse>): QueryAllAccountResponse;
};
export declare const QueryGetPendingAccountRequest: {
    encode(message: QueryGetPendingAccountRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountRequest;
    fromJSON(object: any): QueryGetPendingAccountRequest;
    toJSON(message: QueryGetPendingAccountRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPendingAccountRequest>): QueryGetPendingAccountRequest;
};
export declare const QueryGetPendingAccountResponse: {
    encode(message: QueryGetPendingAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountResponse;
    fromJSON(object: any): QueryGetPendingAccountResponse;
    toJSON(message: QueryGetPendingAccountResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPendingAccountResponse>): QueryGetPendingAccountResponse;
};
export declare const QueryAllPendingAccountRequest: {
    encode(message: QueryAllPendingAccountRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountRequest;
    fromJSON(object: any): QueryAllPendingAccountRequest;
    toJSON(message: QueryAllPendingAccountRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllPendingAccountRequest>): QueryAllPendingAccountRequest;
};
export declare const QueryAllPendingAccountResponse: {
    encode(message: QueryAllPendingAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountResponse;
    fromJSON(object: any): QueryAllPendingAccountResponse;
    toJSON(message: QueryAllPendingAccountResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllPendingAccountResponse>): QueryAllPendingAccountResponse;
};
export declare const QueryGetPendingAccountRevocationRequest: {
    encode(message: QueryGetPendingAccountRevocationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountRevocationRequest;
    fromJSON(object: any): QueryGetPendingAccountRevocationRequest;
    toJSON(message: QueryGetPendingAccountRevocationRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPendingAccountRevocationRequest>): QueryGetPendingAccountRevocationRequest;
};
export declare const QueryGetPendingAccountRevocationResponse: {
    encode(message: QueryGetPendingAccountRevocationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountRevocationResponse;
    fromJSON(object: any): QueryGetPendingAccountRevocationResponse;
    toJSON(message: QueryGetPendingAccountRevocationResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPendingAccountRevocationResponse>): QueryGetPendingAccountRevocationResponse;
};
export declare const QueryAllPendingAccountRevocationRequest: {
    encode(message: QueryAllPendingAccountRevocationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountRevocationRequest;
    fromJSON(object: any): QueryAllPendingAccountRevocationRequest;
    toJSON(message: QueryAllPendingAccountRevocationRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllPendingAccountRevocationRequest>): QueryAllPendingAccountRevocationRequest;
};
export declare const QueryAllPendingAccountRevocationResponse: {
    encode(message: QueryAllPendingAccountRevocationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountRevocationResponse;
    fromJSON(object: any): QueryAllPendingAccountRevocationResponse;
    toJSON(message: QueryAllPendingAccountRevocationResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllPendingAccountRevocationResponse>): QueryAllPendingAccountRevocationResponse;
};
export declare const QueryGetAccountStatRequest: {
    encode(_: QueryGetAccountStatRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountStatRequest;
    fromJSON(_: any): QueryGetAccountStatRequest;
    toJSON(_: QueryGetAccountStatRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetAccountStatRequest>): QueryGetAccountStatRequest;
};
export declare const QueryGetAccountStatResponse: {
    encode(message: QueryGetAccountStatResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAccountStatResponse;
    fromJSON(object: any): QueryGetAccountStatResponse;
    toJSON(message: QueryGetAccountStatResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAccountStatResponse>): QueryGetAccountStatResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a account by index. */
    Account(request: QueryGetAccountRequest): Promise<QueryGetAccountResponse>;
    /** Queries a list of account items. */
    AccountAll(request: QueryAllAccountRequest): Promise<QueryAllAccountResponse>;
    /** Queries a list of pendingAccount items. */
    PendingAccountAll(request: QueryAllPendingAccountRequest): Promise<QueryAllPendingAccountResponse>;
    /** Queries a list of pendingAccountRevocation items. */
    PendingAccountRevocationAll(request: QueryAllPendingAccountRevocationRequest): Promise<QueryAllPendingAccountRevocationResponse>;
    /** Queries a accountStat by index. */
    AccountStat(request: QueryGetAccountStatRequest): Promise<QueryGetAccountStatResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Account(request: QueryGetAccountRequest): Promise<QueryGetAccountResponse>;
    AccountAll(request: QueryAllAccountRequest): Promise<QueryAllAccountResponse>;
    PendingAccountAll(request: QueryAllPendingAccountRequest): Promise<QueryAllPendingAccountResponse>;
    PendingAccountRevocationAll(request: QueryAllPendingAccountRevocationRequest): Promise<QueryAllPendingAccountRevocationResponse>;
    AccountStat(request: QueryGetAccountStatRequest): Promise<QueryGetAccountStatResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
