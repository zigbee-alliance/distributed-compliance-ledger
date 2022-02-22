import { Grant } from '../dclauth/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface PendingAccountRevocation {
    address: string;
    approvals: Grant[];
}
export declare const PendingAccountRevocation: {
    encode(message: PendingAccountRevocation, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PendingAccountRevocation;
    fromJSON(object: any): PendingAccountRevocation;
    toJSON(message: PendingAccountRevocation): unknown;
    fromPartial(object: DeepPartial<PendingAccountRevocation>): PendingAccountRevocation;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
