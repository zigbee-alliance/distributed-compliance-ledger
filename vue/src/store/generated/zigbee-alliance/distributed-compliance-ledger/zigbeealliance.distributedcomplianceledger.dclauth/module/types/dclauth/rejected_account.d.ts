import { Account } from '../dclauth/account';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface RejectedAccount {
    account: Account | undefined;
}
export declare const RejectedAccount: {
    encode(message: RejectedAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RejectedAccount;
    fromJSON(object: any): RejectedAccount;
    toJSON(message: RejectedAccount): unknown;
    fromPartial(object: DeepPartial<RejectedAccount>): RejectedAccount;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
