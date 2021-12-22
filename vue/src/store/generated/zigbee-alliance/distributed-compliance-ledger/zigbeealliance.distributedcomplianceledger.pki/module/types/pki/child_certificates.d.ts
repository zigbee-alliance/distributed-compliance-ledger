import { CertificateIdentifier } from '../pki/certificate_identifier';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface ChildCertificates {
    issuer: string;
    authorityKeyId: string;
    certIds: CertificateIdentifier[];
}
export declare const ChildCertificates: {
    encode(message: ChildCertificates, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ChildCertificates;
    fromJSON(object: any): ChildCertificates;
    toJSON(message: ChildCertificates): unknown;
    fromPartial(object: DeepPartial<ChildCertificates>): ChildCertificates;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
