import { Grant } from '../validator/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface RejectedDisableNode {
    address: string;
    creator: string;
    approvals: Grant[];
    rejectApprovals: Grant[];
}
export declare const RejectedDisableNode: {
    encode(message: RejectedDisableNode, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RejectedDisableNode;
    fromJSON(object: any): RejectedDisableNode;
    toJSON(message: RejectedDisableNode): unknown;
    fromPartial(object: DeepPartial<RejectedDisableNode>): RejectedDisableNode;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
