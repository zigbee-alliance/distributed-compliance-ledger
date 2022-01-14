import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface CertificateIdentifier {
    subject: string;
    subjectKeyId: string;
}
export declare const CertificateIdentifier: {
    encode(message: CertificateIdentifier, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CertificateIdentifier;
    fromJSON(object: any): CertificateIdentifier;
    toJSON(message: CertificateIdentifier): unknown;
    fromPartial(object: DeepPartial<CertificateIdentifier>): CertificateIdentifier;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
