import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface ProposedCertificateRevocation {
    subject: string;
    subjectKeyId: string;
    approvals: string[];
}
export declare const ProposedCertificateRevocation: {
    encode(message: ProposedCertificateRevocation, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ProposedCertificateRevocation;
    fromJSON(object: any): ProposedCertificateRevocation;
    toJSON(message: ProposedCertificateRevocation): unknown;
    fromPartial(object: DeepPartial<ProposedCertificateRevocation>): ProposedCertificateRevocation;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
