import { Reader, Writer } from 'protobufjs/minimal';
import { Validator } from '../validator/validator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { LastValidatorPower } from '../validator/last_validator_power';
import { ValidatorSigningInfo } from '../validator/validator_signing_info';
import { ValidatorMissedBlockBitArray } from '../validator/validator_missed_block_bit_array';
import { ValidatorOwner } from '../validator/validator_owner';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface QueryGetValidatorRequest {
    address: string;
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
    consensusAddress: string;
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
export interface QueryGetValidatorSigningInfoRequest {
    address: string;
}
export interface QueryGetValidatorSigningInfoResponse {
    validatorSigningInfo: ValidatorSigningInfo | undefined;
}
export interface QueryAllValidatorSigningInfoRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllValidatorSigningInfoResponse {
    validatorSigningInfo: ValidatorSigningInfo[];
    pagination: PageResponse | undefined;
}
export interface QueryGetValidatorMissedBlockBitArrayRequest {
    address: string;
    index: number;
}
export interface QueryGetValidatorMissedBlockBitArrayResponse {
    validatorMissedBlockBitArray: ValidatorMissedBlockBitArray | undefined;
}
export interface QueryAllValidatorMissedBlockBitArrayRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllValidatorMissedBlockBitArrayResponse {
    validatorMissedBlockBitArray: ValidatorMissedBlockBitArray[];
    pagination: PageResponse | undefined;
}
export interface QueryGetValidatorOwnerRequest {
    address: string;
}
export interface QueryGetValidatorOwnerResponse {
    validatorOwner: ValidatorOwner | undefined;
}
export interface QueryAllValidatorOwnerRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllValidatorOwnerResponse {
    validatorOwner: ValidatorOwner[];
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
export declare const QueryGetValidatorSigningInfoRequest: {
    encode(message: QueryGetValidatorSigningInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorSigningInfoRequest;
    fromJSON(object: any): QueryGetValidatorSigningInfoRequest;
    toJSON(message: QueryGetValidatorSigningInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorSigningInfoRequest>): QueryGetValidatorSigningInfoRequest;
};
export declare const QueryGetValidatorSigningInfoResponse: {
    encode(message: QueryGetValidatorSigningInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorSigningInfoResponse;
    fromJSON(object: any): QueryGetValidatorSigningInfoResponse;
    toJSON(message: QueryGetValidatorSigningInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorSigningInfoResponse>): QueryGetValidatorSigningInfoResponse;
};
export declare const QueryAllValidatorSigningInfoRequest: {
    encode(message: QueryAllValidatorSigningInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorSigningInfoRequest;
    fromJSON(object: any): QueryAllValidatorSigningInfoRequest;
    toJSON(message: QueryAllValidatorSigningInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorSigningInfoRequest>): QueryAllValidatorSigningInfoRequest;
};
export declare const QueryAllValidatorSigningInfoResponse: {
    encode(message: QueryAllValidatorSigningInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorSigningInfoResponse;
    fromJSON(object: any): QueryAllValidatorSigningInfoResponse;
    toJSON(message: QueryAllValidatorSigningInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorSigningInfoResponse>): QueryAllValidatorSigningInfoResponse;
};
export declare const QueryGetValidatorMissedBlockBitArrayRequest: {
    encode(message: QueryGetValidatorMissedBlockBitArrayRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorMissedBlockBitArrayRequest;
    fromJSON(object: any): QueryGetValidatorMissedBlockBitArrayRequest;
    toJSON(message: QueryGetValidatorMissedBlockBitArrayRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorMissedBlockBitArrayRequest>): QueryGetValidatorMissedBlockBitArrayRequest;
};
export declare const QueryGetValidatorMissedBlockBitArrayResponse: {
    encode(message: QueryGetValidatorMissedBlockBitArrayResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorMissedBlockBitArrayResponse;
    fromJSON(object: any): QueryGetValidatorMissedBlockBitArrayResponse;
    toJSON(message: QueryGetValidatorMissedBlockBitArrayResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorMissedBlockBitArrayResponse>): QueryGetValidatorMissedBlockBitArrayResponse;
};
export declare const QueryAllValidatorMissedBlockBitArrayRequest: {
    encode(message: QueryAllValidatorMissedBlockBitArrayRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorMissedBlockBitArrayRequest;
    fromJSON(object: any): QueryAllValidatorMissedBlockBitArrayRequest;
    toJSON(message: QueryAllValidatorMissedBlockBitArrayRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorMissedBlockBitArrayRequest>): QueryAllValidatorMissedBlockBitArrayRequest;
};
export declare const QueryAllValidatorMissedBlockBitArrayResponse: {
    encode(message: QueryAllValidatorMissedBlockBitArrayResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorMissedBlockBitArrayResponse;
    fromJSON(object: any): QueryAllValidatorMissedBlockBitArrayResponse;
    toJSON(message: QueryAllValidatorMissedBlockBitArrayResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorMissedBlockBitArrayResponse>): QueryAllValidatorMissedBlockBitArrayResponse;
};
export declare const QueryGetValidatorOwnerRequest: {
    encode(message: QueryGetValidatorOwnerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorOwnerRequest;
    fromJSON(object: any): QueryGetValidatorOwnerRequest;
    toJSON(message: QueryGetValidatorOwnerRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorOwnerRequest>): QueryGetValidatorOwnerRequest;
};
export declare const QueryGetValidatorOwnerResponse: {
    encode(message: QueryGetValidatorOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorOwnerResponse;
    fromJSON(object: any): QueryGetValidatorOwnerResponse;
    toJSON(message: QueryGetValidatorOwnerResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetValidatorOwnerResponse>): QueryGetValidatorOwnerResponse;
};
export declare const QueryAllValidatorOwnerRequest: {
    encode(message: QueryAllValidatorOwnerRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorOwnerRequest;
    fromJSON(object: any): QueryAllValidatorOwnerRequest;
    toJSON(message: QueryAllValidatorOwnerRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorOwnerRequest>): QueryAllValidatorOwnerRequest;
};
export declare const QueryAllValidatorOwnerResponse: {
    encode(message: QueryAllValidatorOwnerResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorOwnerResponse;
    fromJSON(object: any): QueryAllValidatorOwnerResponse;
    toJSON(message: QueryAllValidatorOwnerResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllValidatorOwnerResponse>): QueryAllValidatorOwnerResponse;
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
    /** Queries a validatorSigningInfo by index. */
    ValidatorSigningInfo(request: QueryGetValidatorSigningInfoRequest): Promise<QueryGetValidatorSigningInfoResponse>;
    /** Queries a list of validatorSigningInfo items. */
    ValidatorSigningInfoAll(request: QueryAllValidatorSigningInfoRequest): Promise<QueryAllValidatorSigningInfoResponse>;
    /** Queries a validatorMissedBlockBitArray by index. */
    ValidatorMissedBlockBitArray(request: QueryGetValidatorMissedBlockBitArrayRequest): Promise<QueryGetValidatorMissedBlockBitArrayResponse>;
    /** Queries a list of validatorMissedBlockBitArray items. */
    ValidatorMissedBlockBitArrayAll(request: QueryAllValidatorMissedBlockBitArrayRequest): Promise<QueryAllValidatorMissedBlockBitArrayResponse>;
    /** Queries a validatorOwner by index. */
    ValidatorOwner(request: QueryGetValidatorOwnerRequest): Promise<QueryGetValidatorOwnerResponse>;
    /** Queries a list of validatorOwner items. */
    ValidatorOwnerAll(request: QueryAllValidatorOwnerRequest): Promise<QueryAllValidatorOwnerResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse>;
    ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>;
    LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse>;
    LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse>;
    ValidatorSigningInfo(request: QueryGetValidatorSigningInfoRequest): Promise<QueryGetValidatorSigningInfoResponse>;
    ValidatorSigningInfoAll(request: QueryAllValidatorSigningInfoRequest): Promise<QueryAllValidatorSigningInfoResponse>;
    ValidatorMissedBlockBitArray(request: QueryGetValidatorMissedBlockBitArrayRequest): Promise<QueryGetValidatorMissedBlockBitArrayResponse>;
    ValidatorMissedBlockBitArrayAll(request: QueryAllValidatorMissedBlockBitArrayRequest): Promise<QueryAllValidatorMissedBlockBitArrayResponse>;
    ValidatorOwner(request: QueryGetValidatorOwnerRequest): Promise<QueryGetValidatorOwnerResponse>;
    ValidatorOwnerAll(request: QueryAllValidatorOwnerRequest): Promise<QueryAllValidatorOwnerResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
