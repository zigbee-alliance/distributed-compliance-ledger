import { Account } from '../dclauth/account';
import { Grant } from '../dclauth/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface RevokedAccount {
    account: Account | undefined;
    revokeApprovals: Grant[];
}
export declare const RevokedAccount: {
    encode(message: RevokedAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RevokedAccount;
    fromJSON(object: any): RevokedAccount;
    toJSON(message: RevokedAccount): unknown;
    fromPartial(object: DeepPartial<RevokedAccount>): RevokedAccount;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
