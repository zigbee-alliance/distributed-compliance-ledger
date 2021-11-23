import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface AccountStat {
    number: number;
}
export declare const AccountStat: {
    encode(message: AccountStat, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): AccountStat;
    fromJSON(object: any): AccountStat;
    toJSON(message: AccountStat): unknown;
    fromPartial(object: DeepPartial<AccountStat>): AccountStat;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
