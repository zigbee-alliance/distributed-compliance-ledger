import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade';
import { ApprovedUpgrade } from '../dclupgrade/approved_upgrade';
import { RejectedUpgrade } from '../dclupgrade/rejected_upgrade';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclupgrade";
/** GenesisState defines the dclupgrade module's genesis state. */
export interface GenesisState {
    proposedUpgradeList: ProposedUpgrade[];
    approvedUpgradeList: ApprovedUpgrade[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    rejectedUpgradeList: RejectedUpgrade[];
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
