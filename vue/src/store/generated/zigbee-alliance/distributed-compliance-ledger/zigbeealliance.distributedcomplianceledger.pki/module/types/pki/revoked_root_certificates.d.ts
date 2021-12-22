import { CertificateIdentifier } from '../pki/certificate_identifier';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface RevokedRootCertificates {
    certs: CertificateIdentifier[];
}
export declare const RevokedRootCertificates: {
    encode(message: RevokedRootCertificates, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RevokedRootCertificates;
    fromJSON(object: any): RevokedRootCertificates;
    toJSON(message: RevokedRootCertificates): unknown;
    fromPartial(object: DeepPartial<RevokedRootCertificates>): RevokedRootCertificates;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
