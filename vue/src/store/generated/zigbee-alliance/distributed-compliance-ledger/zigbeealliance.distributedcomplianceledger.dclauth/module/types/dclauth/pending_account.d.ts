import { Account } from '../dclauth/account';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
/**
 * TODO issue 99: do we need that ???
 * option (gogoproto.goproto_getters)  = false;
 * option (gogoproto.goproto_stringer) = false;
 */
export interface PendingAccount {
    account: Account | undefined;
}
export declare const PendingAccount: {
    encode(message: PendingAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PendingAccount;
    fromJSON(object: any): PendingAccount;
    toJSON(message: PendingAccount): unknown;
    fromPartial(object: DeepPartial<PendingAccount>): PendingAccount;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
