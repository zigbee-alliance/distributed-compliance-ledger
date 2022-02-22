import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface Grant {
    address: string;
    /** number of nanoseconds elapsed since January 1, 1970 UTC */
    time: number;
    info: string;
}
export declare const Grant: {
    encode(message: Grant, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Grant;
    fromJSON(object: any): Grant;
    toJSON(message: Grant): unknown;
    fromPartial(object: DeepPartial<Grant>): Grant;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
