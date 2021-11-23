import { Description } from '../validator/description';
import { Any } from '../google/protobuf/any';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface Validator {
    /** the account address of validator owner */
    owner: string;
    /** description of the validator */
    description: Description | undefined;
    /** the consensus public key of the tendermint validator */
    pubKey: Any | undefined;
    /** validator consensus power */
    power: number;
    /** has the validator been removed from validator set */
    jailed: boolean;
    /** the reason of validator jailing */
    jailedReason: string;
}
export declare const Validator: {
    encode(message: Validator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Validator;
    fromJSON(object: any): Validator;
    toJSON(message: Validator): unknown;
    fromPartial(object: DeepPartial<Validator>): Validator;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
