import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";
export interface RejectedNode {
    owner: string;
    approvals: string[];
}
export declare const RejectedNode: {
    encode(message: RejectedNode, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RejectedNode;
    fromJSON(object: any): RejectedNode;
    toJSON(message: RejectedNode): unknown;
    fromPartial(object: DeepPartial<RejectedNode>): RejectedNode;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
