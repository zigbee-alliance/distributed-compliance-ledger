import { Reader, Writer } from 'protobufjs/minimal';
import { Validator } from '../validator/validator';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { LastValidatorPower } from '../validator/last_validator_power';
import { ProposedDisableValidator } from '../validator/proposed_disable_validator';
import { DisabledValidator } from '../validator/disabled_validator';
import { RejectedDisableValidator } from '../validator/rejected_validator';
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
export interface QueryGetProposedDisableValidatorRequest {
    address: string;
}
export interface QueryGetProposedDisableValidatorResponse {
    proposedDisableValidator: ProposedDisableValidator | undefined;
}
export interface QueryAllProposedDisableValidatorRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllProposedDisableValidatorResponse {
    proposedDisableValidator: ProposedDisableValidator[];
    pagination: PageResponse | undefined;
}
export interface QueryGetDisabledValidatorRequest {
    address: string;
}
export interface QueryGetDisabledValidatorResponse {
    disabledValidator: DisabledValidator | undefined;
}
export interface QueryAllDisabledValidatorRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllDisabledValidatorResponse {
    disabledValidator: DisabledValidator[];
    pagination: PageResponse | undefined;
}
export interface QueryGetRejectedDisableValidatorRequest {
    owner: string;
}
export interface QueryGetRejectedDisableValidatorResponse {
    rejectedValidator: RejectedDisableValidator | undefined;
}
export interface QueryAllRejectedDisableValidatorRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllRejectedDisableValidatorResponse {
    rejectedValidator: RejectedDisableValidator[];
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
export declare const QueryGetProposedDisableValidatorRequest: {
    encode(message: QueryGetProposedDisableValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedDisableValidatorRequest;
    fromJSON(object: any): QueryGetProposedDisableValidatorRequest;
    toJSON(message: QueryGetProposedDisableValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedDisableValidatorRequest>): QueryGetProposedDisableValidatorRequest;
};
export declare const QueryGetProposedDisableValidatorResponse: {
    encode(message: QueryGetProposedDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedDisableValidatorResponse;
    fromJSON(object: any): QueryGetProposedDisableValidatorResponse;
    toJSON(message: QueryGetProposedDisableValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedDisableValidatorResponse>): QueryGetProposedDisableValidatorResponse;
};
export declare const QueryAllProposedDisableValidatorRequest: {
    encode(message: QueryAllProposedDisableValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedDisableValidatorRequest;
    fromJSON(object: any): QueryAllProposedDisableValidatorRequest;
    toJSON(message: QueryAllProposedDisableValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedDisableValidatorRequest>): QueryAllProposedDisableValidatorRequest;
};
export declare const QueryAllProposedDisableValidatorResponse: {
    encode(message: QueryAllProposedDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedDisableValidatorResponse;
    fromJSON(object: any): QueryAllProposedDisableValidatorResponse;
    toJSON(message: QueryAllProposedDisableValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedDisableValidatorResponse>): QueryAllProposedDisableValidatorResponse;
};
export declare const QueryGetDisabledValidatorRequest: {
    encode(message: QueryGetDisabledValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDisabledValidatorRequest;
    fromJSON(object: any): QueryGetDisabledValidatorRequest;
    toJSON(message: QueryGetDisabledValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetDisabledValidatorRequest>): QueryGetDisabledValidatorRequest;
};
export declare const QueryGetDisabledValidatorResponse: {
    encode(message: QueryGetDisabledValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDisabledValidatorResponse;
    fromJSON(object: any): QueryGetDisabledValidatorResponse;
    toJSON(message: QueryGetDisabledValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetDisabledValidatorResponse>): QueryGetDisabledValidatorResponse;
};
export declare const QueryAllDisabledValidatorRequest: {
    encode(message: QueryAllDisabledValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDisabledValidatorRequest;
    fromJSON(object: any): QueryAllDisabledValidatorRequest;
    toJSON(message: QueryAllDisabledValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllDisabledValidatorRequest>): QueryAllDisabledValidatorRequest;
};
export declare const QueryAllDisabledValidatorResponse: {
    encode(message: QueryAllDisabledValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDisabledValidatorResponse;
    fromJSON(object: any): QueryAllDisabledValidatorResponse;
    toJSON(message: QueryAllDisabledValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllDisabledValidatorResponse>): QueryAllDisabledValidatorResponse;
};
export declare const QueryGetRejectedDisableValidatorRequest: {
    encode(message: QueryGetRejectedDisableValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedDisableValidatorRequest;
    fromJSON(object: any): QueryGetRejectedDisableValidatorRequest;
    toJSON(message: QueryGetRejectedDisableValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRejectedDisableValidatorRequest>): QueryGetRejectedDisableValidatorRequest;
};
export declare const QueryGetRejectedDisableValidatorResponse: {
    encode(message: QueryGetRejectedDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedDisableValidatorResponse;
    fromJSON(object: any): QueryGetRejectedDisableValidatorResponse;
    toJSON(message: QueryGetRejectedDisableValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRejectedDisableValidatorResponse>): QueryGetRejectedDisableValidatorResponse;
};
export declare const QueryAllRejectedDisableValidatorRequest: {
    encode(message: QueryAllRejectedDisableValidatorRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedDisableValidatorRequest;
    fromJSON(object: any): QueryAllRejectedDisableValidatorRequest;
    toJSON(message: QueryAllRejectedDisableValidatorRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRejectedDisableValidatorRequest>): QueryAllRejectedDisableValidatorRequest;
};
export declare const QueryAllRejectedDisableValidatorResponse: {
    encode(message: QueryAllRejectedDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedDisableValidatorResponse;
    fromJSON(object: any): QueryAllRejectedDisableValidatorResponse;
    toJSON(message: QueryAllRejectedDisableValidatorResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRejectedDisableValidatorResponse>): QueryAllRejectedDisableValidatorResponse;
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
    /** Queries a ProposedDisableValidator by index. */
    ProposedDisableValidator(request: QueryGetProposedDisableValidatorRequest): Promise<QueryGetProposedDisableValidatorResponse>;
    /** Queries a list of ProposedDisableValidator items. */
    ProposedDisableValidatorAll(request: QueryAllProposedDisableValidatorRequest): Promise<QueryAllProposedDisableValidatorResponse>;
    /** Queries a DisabledValidator by index. */
    DisabledValidator(request: QueryGetDisabledValidatorRequest): Promise<QueryGetDisabledValidatorResponse>;
    /** Queries a list of DisabledValidator items. */
    DisabledValidatorAll(request: QueryAllDisabledValidatorRequest): Promise<QueryAllDisabledValidatorResponse>;
    /** Queries a RejectedNode by index. */
    RejectedDisableValidator(request: QueryGetRejectedDisableValidatorRequest): Promise<QueryGetRejectedDisableValidatorResponse>;
    /** Queries a list of RejectedNode items. */
    RejectedDisableValidatorAll(request: QueryAllRejectedDisableValidatorRequest): Promise<QueryAllRejectedDisableValidatorResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse>;
    ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>;
    LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse>;
    LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse>;
    ProposedDisableValidator(request: QueryGetProposedDisableValidatorRequest): Promise<QueryGetProposedDisableValidatorResponse>;
    ProposedDisableValidatorAll(request: QueryAllProposedDisableValidatorRequest): Promise<QueryAllProposedDisableValidatorResponse>;
    DisabledValidator(request: QueryGetDisabledValidatorRequest): Promise<QueryGetDisabledValidatorResponse>;
    DisabledValidatorAll(request: QueryAllDisabledValidatorRequest): Promise<QueryAllDisabledValidatorResponse>;
    RejectedDisableValidator(request: QueryGetRejectedDisableValidatorRequest): Promise<QueryGetRejectedDisableValidatorResponse>;
    RejectedDisableValidatorAll(request: QueryAllRejectedDisableValidatorRequest): Promise<QueryAllRejectedDisableValidatorResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
