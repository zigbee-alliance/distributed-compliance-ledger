import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface CertifiedModel {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
    value: boolean;
}
export declare const CertifiedModel: {
    encode(message: CertifiedModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CertifiedModel;
    fromJSON(object: any): CertifiedModel;
    toJSON(message: CertifiedModel): unknown;
    fromPartial(object: DeepPartial<CertifiedModel>): CertifiedModel;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
