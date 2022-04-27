import { Reader, Writer } from 'protobufjs/minimal';
import { Any } from '../google/protobuf/any';
import { Description } from '../validator/description';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface MsgCreateValidator {
    signer: string;
    pubKey: Any | undefined;
    description: Description | undefined;
}
export interface MsgCreateValidatorResponse {
}
export interface MsgProposeDisableValidator {
    creator: string;
    address: string;
    info: string;
    time: number;
}
export interface MsgProposeDisableValidatorResponse {
}
export interface MsgApproveDisableValidator {
    creator: string;
    address: string;
    info: string;
    time: number;
}
export interface MsgApproveDisableValidatorResponse {
}
export interface MsgDisableValidator {
    creator: string;
}
export interface MsgDisableValidatorResponse {
}
export interface MsgEnableValidator {
    creator: string;
}
export interface MsgEnableValidatorResponse {
}
export interface MsgRejectDisableValidator {
    creator: string;
    address: string;
    info: string;
    time: number;
}
export interface MsgRejectDisableValidatorResponse {
}
export declare const MsgCreateValidator: {
    encode(message: MsgCreateValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateValidator;
    fromJSON(object: any): MsgCreateValidator;
    toJSON(message: MsgCreateValidator): unknown;
    fromPartial(object: DeepPartial<MsgCreateValidator>): MsgCreateValidator;
};
export declare const MsgCreateValidatorResponse: {
    encode(_: MsgCreateValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateValidatorResponse;
    fromJSON(_: any): MsgCreateValidatorResponse;
    toJSON(_: MsgCreateValidatorResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateValidatorResponse>): MsgCreateValidatorResponse;
};
export declare const MsgProposeDisableValidator: {
    encode(message: MsgProposeDisableValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeDisableValidator;
    fromJSON(object: any): MsgProposeDisableValidator;
    toJSON(message: MsgProposeDisableValidator): unknown;
    fromPartial(object: DeepPartial<MsgProposeDisableValidator>): MsgProposeDisableValidator;
};
export declare const MsgProposeDisableValidatorResponse: {
    encode(_: MsgProposeDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeDisableValidatorResponse;
    fromJSON(_: any): MsgProposeDisableValidatorResponse;
    toJSON(_: MsgProposeDisableValidatorResponse): unknown;
    fromPartial(_: DeepPartial<MsgProposeDisableValidatorResponse>): MsgProposeDisableValidatorResponse;
};
export declare const MsgApproveDisableValidator: {
    encode(message: MsgApproveDisableValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveDisableValidator;
    fromJSON(object: any): MsgApproveDisableValidator;
    toJSON(message: MsgApproveDisableValidator): unknown;
    fromPartial(object: DeepPartial<MsgApproveDisableValidator>): MsgApproveDisableValidator;
};
export declare const MsgApproveDisableValidatorResponse: {
    encode(_: MsgApproveDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveDisableValidatorResponse;
    fromJSON(_: any): MsgApproveDisableValidatorResponse;
    toJSON(_: MsgApproveDisableValidatorResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveDisableValidatorResponse>): MsgApproveDisableValidatorResponse;
};
export declare const MsgDisableValidator: {
    encode(message: MsgDisableValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDisableValidator;
    fromJSON(object: any): MsgDisableValidator;
    toJSON(message: MsgDisableValidator): unknown;
    fromPartial(object: DeepPartial<MsgDisableValidator>): MsgDisableValidator;
};
export declare const MsgDisableValidatorResponse: {
    encode(_: MsgDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDisableValidatorResponse;
    fromJSON(_: any): MsgDisableValidatorResponse;
    toJSON(_: MsgDisableValidatorResponse): unknown;
    fromPartial(_: DeepPartial<MsgDisableValidatorResponse>): MsgDisableValidatorResponse;
};
export declare const MsgEnableValidator: {
    encode(message: MsgEnableValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgEnableValidator;
    fromJSON(object: any): MsgEnableValidator;
    toJSON(message: MsgEnableValidator): unknown;
    fromPartial(object: DeepPartial<MsgEnableValidator>): MsgEnableValidator;
};
export declare const MsgEnableValidatorResponse: {
    encode(_: MsgEnableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgEnableValidatorResponse;
    fromJSON(_: any): MsgEnableValidatorResponse;
    toJSON(_: MsgEnableValidatorResponse): unknown;
    fromPartial(_: DeepPartial<MsgEnableValidatorResponse>): MsgEnableValidatorResponse;
};
export declare const MsgRejectDisableValidator: {
    encode(message: MsgRejectDisableValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectDisableValidator;
    fromJSON(object: any): MsgRejectDisableValidator;
    toJSON(message: MsgRejectDisableValidator): unknown;
    fromPartial(object: DeepPartial<MsgRejectDisableValidator>): MsgRejectDisableValidator;
};
export declare const MsgRejectDisableValidatorResponse: {
    encode(_: MsgRejectDisableValidatorResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectDisableValidatorResponse;
    fromJSON(_: any): MsgRejectDisableValidatorResponse;
    toJSON(_: MsgRejectDisableValidatorResponse): unknown;
    fromPartial(_: DeepPartial<MsgRejectDisableValidatorResponse>): MsgRejectDisableValidatorResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse>;
    ProposeDisableValidator(request: MsgProposeDisableValidator): Promise<MsgProposeDisableValidatorResponse>;
    ApproveDisableValidator(request: MsgApproveDisableValidator): Promise<MsgApproveDisableValidatorResponse>;
    DisableValidator(request: MsgDisableValidator): Promise<MsgDisableValidatorResponse>;
    EnableValidator(request: MsgEnableValidator): Promise<MsgEnableValidatorResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RejectDisableValidator(request: MsgRejectDisableValidator): Promise<MsgRejectDisableValidatorResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse>;
    ProposeDisableValidator(request: MsgProposeDisableValidator): Promise<MsgProposeDisableValidatorResponse>;
    ApproveDisableValidator(request: MsgApproveDisableValidator): Promise<MsgApproveDisableValidatorResponse>;
    DisableValidator(request: MsgDisableValidator): Promise<MsgDisableValidatorResponse>;
    EnableValidator(request: MsgEnableValidator): Promise<MsgEnableValidatorResponse>;
    RejectDisableValidator(request: MsgRejectDisableValidator): Promise<MsgRejectDisableValidatorResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
