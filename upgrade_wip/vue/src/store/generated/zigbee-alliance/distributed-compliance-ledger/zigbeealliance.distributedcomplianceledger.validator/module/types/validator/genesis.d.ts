import { Validator } from '../validator/validator';
import { LastValidatorPower } from '../validator/last_validator_power';
import { ValidatorSigningInfo } from '../validator/validator_signing_info';
import { ValidatorMissedBlockBitArray } from '../validator/validator_missed_block_bit_array';
import { ValidatorOwner } from '../validator/validator_owner';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
/** GenesisState defines the validator module's genesis state. */
export interface GenesisState {
    validatorList: Validator[];
    lastValidatorPowerList: LastValidatorPower[];
    validatorSigningInfoList: ValidatorSigningInfo[];
    validatorMissedBlockBitArrayList: ValidatorMissedBlockBitArray[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    validatorOwnerList: ValidatorOwner[];
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
