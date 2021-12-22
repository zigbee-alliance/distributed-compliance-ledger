import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface ApprovedCertificatesBySubject {
    subject: string;
    subjectKeyIds: string[];
}
export declare const ApprovedCertificatesBySubject: {
    encode(message: ApprovedCertificatesBySubject, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ApprovedCertificatesBySubject;
    fromJSON(object: any): ApprovedCertificatesBySubject;
    toJSON(message: ApprovedCertificatesBySubject): unknown;
    fromPartial(object: DeepPartial<ApprovedCertificatesBySubject>): ApprovedCertificatesBySubject;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
