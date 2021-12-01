import { Reader, Writer } from 'protobufjs/minimal';
import { Description } from '../validator/description';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface MsgCreateValidator {
    signer: string;
    pubKey: string;
    description: Description | undefined;
}
export interface MsgCreateValidatorResponse {
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
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
