import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface ValidatorOwner {
    address: string;
}
export declare const ValidatorOwner: {
    encode(message: ValidatorOwner, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ValidatorOwner;
    fromJSON(object: any): ValidatorOwner;
    toJSON(message: ValidatorOwner): unknown;
    fromPartial(object: DeepPartial<ValidatorOwner>): ValidatorOwner;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
