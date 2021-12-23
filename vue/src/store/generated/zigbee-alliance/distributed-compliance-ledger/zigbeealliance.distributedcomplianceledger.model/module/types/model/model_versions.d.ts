import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
export interface ModelVersions {
    vid: number;
    pid: number;
    softwareVersions: number[];
}
export declare const ModelVersions: {
    encode(message: ModelVersions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ModelVersions;
    fromJSON(object: any): ModelVersions;
    toJSON(message: ModelVersions): unknown;
    fromPartial(object: DeepPartial<ModelVersions>): ModelVersions;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
