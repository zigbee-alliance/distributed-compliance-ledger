import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface ValidatorSigningInfo {
    address: string;
    startHeight: number;
    indexOffset: number;
    missedBlocksCounter: number;
}
export declare const ValidatorSigningInfo: {
    encode(message: ValidatorSigningInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ValidatorSigningInfo;
    fromJSON(object: any): ValidatorSigningInfo;
    toJSON(message: ValidatorSigningInfo): unknown;
    fromPartial(object: DeepPartial<ValidatorSigningInfo>): ValidatorSigningInfo;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
