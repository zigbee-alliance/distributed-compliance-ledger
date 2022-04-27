import { Grant } from '../validator/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface RejectedDisableValidator {
    address: string;
    creator: string;
    approvals: Grant[];
    rejects: Grant[];
}
export declare const RejectedDisableValidator: {
    encode(message: RejectedDisableValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RejectedDisableValidator;
    fromJSON(object: any): RejectedDisableValidator;
    toJSON(message: RejectedDisableValidator): unknown;
    fromPartial(object: DeepPartial<RejectedDisableValidator>): RejectedDisableValidator;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
