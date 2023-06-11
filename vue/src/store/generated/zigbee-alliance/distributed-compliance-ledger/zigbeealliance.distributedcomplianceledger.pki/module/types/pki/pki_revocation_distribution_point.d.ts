import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface PkiRevocationDistributionPoint {
    vid: number;
    label: string;
    issuerSubjectKeyID: string;
    pid: number;
    isPAA: boolean;
    crlSignerCertificate: string;
    dataURL: string;
    dataFileSize: number;
    dataDigest: string;
    dataDigestType: number;
    revocationType: number;
}
export declare const PkiRevocationDistributionPoint: {
    encode(message: PkiRevocationDistributionPoint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PkiRevocationDistributionPoint;
    fromJSON(object: any): PkiRevocationDistributionPoint;
    toJSON(message: PkiRevocationDistributionPoint): unknown;
    fromPartial(object: DeepPartial<PkiRevocationDistributionPoint>): PkiRevocationDistributionPoint;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
