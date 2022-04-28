import { Plan } from '../cosmos/upgrade/v1beta1/upgrade';
import { Grant } from '../dclupgrade/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclupgrade";
export interface ProposedUpgrade {
    plan: Plan | undefined;
    creator: string;
    approvals: Grant[];
    rejects: Grant[];
}
export declare const ProposedUpgrade: {
    encode(message: ProposedUpgrade, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ProposedUpgrade;
    fromJSON(object: any): ProposedUpgrade;
    toJSON(message: ProposedUpgrade): unknown;
    fromPartial(object: DeepPartial<ProposedUpgrade>): ProposedUpgrade;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
