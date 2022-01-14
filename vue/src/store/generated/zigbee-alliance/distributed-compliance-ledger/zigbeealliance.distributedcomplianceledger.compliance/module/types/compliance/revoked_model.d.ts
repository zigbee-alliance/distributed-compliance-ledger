import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface RevokedModel {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
    value: boolean;
}
export declare const RevokedModel: {
    encode(message: RevokedModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RevokedModel;
    fromJSON(object: any): RevokedModel;
    toJSON(message: RevokedModel): unknown;
    fromPartial(object: DeepPartial<RevokedModel>): RevokedModel;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
