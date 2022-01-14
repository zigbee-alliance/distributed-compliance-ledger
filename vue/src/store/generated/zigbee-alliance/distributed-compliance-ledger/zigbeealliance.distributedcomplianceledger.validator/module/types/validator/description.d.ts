import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface Description {
    /** a human-readable name for the validator. */
    moniker: string;
    /** optional identity signature. */
    identity: string;
    /** optional website link. */
    website: string;
    /** optional details. */
    details: string;
}
export declare const Description: {
    encode(message: Description, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Description;
    fromJSON(object: any): Description;
    toJSON(message: Description): unknown;
    fromPartial(object: DeepPartial<Description>): Description;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
