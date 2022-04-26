import { Certificate } from '../pki/certificate';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface RejectedCertificate {
    subject: string;
    subjectKeyId: string;
    certs: Certificate[];
}
export declare const RejectedCertificate: {
    encode(message: RejectedCertificate, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RejectedCertificate;
    fromJSON(object: any): RejectedCertificate;
    toJSON(message: RejectedCertificate): unknown;
    fromPartial(object: DeepPartial<RejectedCertificate>): RejectedCertificate;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
