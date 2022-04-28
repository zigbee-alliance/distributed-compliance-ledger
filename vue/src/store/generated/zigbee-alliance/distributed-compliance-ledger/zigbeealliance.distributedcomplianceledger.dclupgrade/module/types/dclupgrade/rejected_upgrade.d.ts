import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclupgrade";
export interface RejectedUpgrade {
    name: string;
    plan: string;
    creator: string;
    approvals: string[];
}
export declare const RejectedUpgrade: {
    encode(message: RejectedUpgrade, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RejectedUpgrade;
    fromJSON(object: any): RejectedUpgrade;
    toJSON(message: RejectedUpgrade): unknown;
    fromPartial(object: DeepPartial<RejectedUpgrade>): RejectedUpgrade;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
