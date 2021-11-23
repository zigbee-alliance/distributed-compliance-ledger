import { Reader, Writer } from 'protobufjs/minimal';
import { Validator } from '../validator/validator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { LastValidatorPower } from '../validator/last_validator_power';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface QueryGetValidatorRequest {
    owner: string;
}
export interface QueryGetValidatorResponse {
    validator: Validator | undefined;
}
export interface QueryAllValidatorRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllValidatorResponse {
    validator: Validator[];
    pagination: PageResponse | undefined;
}
export interface QueryGetLastValidatorPowerRequest {
    owner: string;
}
export interface QueryGetLastValidatorPowerResponse {
    lastValidatorPower: LastValidatorPower | undefined;
}
export interface QueryAllLastValidatorPowerRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllLastValidatorPowerResponse {
    lastValidatorPower: LastValidatorPower[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetValidatorRequest: {
    encode(message: QueryGetValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorRequest;
    fromJSON(object: any): QueryGetValidatorRequest;
    toJSON(message: QueryGetValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorRequest>): QueryGetValidatorRequest;
};
export declare const QueryGetValidatorResponse: {
    encode(message: QueryGetValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorResponse;
    fromJSON(object: any): QueryGetValidatorResponse;
    toJSON(message: QueryGetValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorResponse>): QueryGetValidatorResponse;
};
export declare const QueryAllValidatorRequest: {
    encode(message: QueryAllValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorRequest;
    fromJSON(object: any): QueryAllValidatorRequest;
    toJSON(message: QueryAllValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorRequest>): QueryAllValidatorRequest;
};
export declare const QueryAllValidatorResponse: {
    encode(message: QueryAllValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorResponse;
    fromJSON(object: any): QueryAllValidatorResponse;
    toJSON(message: QueryAllValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorResponse>): QueryAllValidatorResponse;
};
export declare const QueryGetLastValidatorPowerRequest: {
    encode(message: QueryGetLastValidatorPowerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLastValidatorPowerRequest;
    fromJSON(object: any): QueryGetLastValidatorPowerRequest;
    toJSON(message: QueryGetLastValidatorPowerRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetLastValidatorPowerRequest>): QueryGetLastValidatorPowerRequest;
};
export declare const QueryGetLastValidatorPowerResponse: {
    encode(message: QueryGetLastValidatorPowerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetLastValidatorPowerResponse;
    fromJSON(object: any): QueryGetLastValidatorPowerResponse;
    toJSON(message: QueryGetLastValidatorPowerResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetLastValidatorPowerResponse>): QueryGetLastValidatorPowerResponse;
};
export declare const QueryAllLastValidatorPowerRequest: {
    encode(message: QueryAllLastValidatorPowerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllLastValidatorPowerRequest;
    fromJSON(object: any): QueryAllLastValidatorPowerRequest;
    toJSON(message: QueryAllLastValidatorPowerRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllLastValidatorPowerRequest>): QueryAllLastValidatorPowerRequest;
};
export declare const QueryAllLastValidatorPowerResponse: {
    encode(message: QueryAllLastValidatorPowerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllLastValidatorPowerResponse;
    fromJSON(object: any): QueryAllLastValidatorPowerResponse;
    toJSON(message: QueryAllLastValidatorPowerResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllLastValidatorPowerResponse>): QueryAllLastValidatorPowerResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a validator by index. */
    Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse>;
    /** Queries a list of validator items. */
    ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>;
    /** Queries a lastValidatorPower by index. */
    LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse>;
    /** Queries a list of lastValidatorPower items. */
    LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse>;
    ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>;
    LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse>;
    LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
