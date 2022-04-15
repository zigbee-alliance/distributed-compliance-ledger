import { Account } from '../dclauth/account';
import { PendingAccount } from '../dclauth/pending_account';
import { PendingAccountRevocation } from '../dclauth/pending_account_revocation';
import { AccountStat } from '../dclauth/account_stat';
import { RevokedAccount } from '../dclauth/revoked_account';
import { RejectedAccount } from '../dclauth/rejected_account';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
/** GenesisState defines the dclauth module's genesis state. */
export interface GenesisState {
    accountList: Account[];
    pendingAccountList: PendingAccount[];
    pendingAccountRevocationList: PendingAccountRevocation[];
    accountStat: AccountStat | undefined;
    revokedAccountList: RevokedAccount[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    rejectedAccountList: RejectedAccount[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
