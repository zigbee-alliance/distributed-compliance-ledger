import { Grant } from '../validator/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface ProposedDisableValidator {
    address: string;
    creator: string;
    approvals: Grant[];
    rejects: Grant[];
}
export declare const ProposedDisableValidator: {
    encode(message: ProposedDisableValidator, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ProposedDisableValidator;
    fromJSON(object: any): ProposedDisableValidator;
    toJSON(message: ProposedDisableValidator): unknown;
    fromPartial(object: DeepPartial<ProposedDisableValidator>): ProposedDisableValidator;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
