import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface ValidatorMissedBlockBitArray {
    address: string;
    index: number;
}
export declare const ValidatorMissedBlockBitArray: {
    encode(message: ValidatorMissedBlockBitArray, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ValidatorMissedBlockBitArray;
    fromJSON(object: any): ValidatorMissedBlockBitArray;
    toJSON(message: ValidatorMissedBlockBitArray): unknown;
    fromPartial(object: DeepPartial<ValidatorMissedBlockBitArray>): ValidatorMissedBlockBitArray;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
