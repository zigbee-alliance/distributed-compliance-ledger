import { Writer, Reader } from 'protobufjs/minimal';
import { BaseAccount } from '../cosmos/auth/v1beta1/auth';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface Account {
    baseAccount: BaseAccount | undefined;
    /**
     * NOTE. we do not user AccountRoles casting here to preserve repeated form
     *       so protobuf takes care about repeated items in generated code,
     *       (but that might be not the final solution)
     */
    roles: string[];
    vendorID: number;
}
export declare const Account: {
    encode(message: Account, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Account;
    fromJSON(object: any): Account;
    toJSON(message: Account): unknown;
    fromPartial(object: DeepPartial<Account>): Account;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
