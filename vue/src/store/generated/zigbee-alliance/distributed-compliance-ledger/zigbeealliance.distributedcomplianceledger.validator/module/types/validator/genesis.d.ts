import { Validator } from '../validator/validator';
import { LastValidatorPower } from '../validator/last_validator_power';
import { ProposedDisableValidator } from '../validator/proposed_disable_validator';
import { DisabledValidator } from '../validator/disabled_validator';
import { RejectedDisableValidator } from '../validator/rejected_validator';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
/** GenesisState defines the validator module's genesis state. */
export interface GenesisState {
    validatorList: Validator[];
    lastValidatorPowerList: LastValidatorPower[];
    proposedDisableValidatorList: ProposedDisableValidator[];
    disabledValidatorList: DisabledValidator[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    rejectedValidatorList: RejectedDisableValidator[];
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
