import { Grant } from '../validator/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface DisabledValidator {
    address: string;
    creator: string;
    approvals: Grant[];
    disabledByNodeAdmin: boolean;
    rejects: Grant[];
}
export declare const DisabledValidator: {
    encode(message: DisabledValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DisabledValidator;
    fromJSON(object: any): DisabledValidator;
    toJSON(message: DisabledValidator): unknown;
    fromPartial(object: DeepPartial<DisabledValidator>): DisabledValidator;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
