import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface ProposedCertificate {
    subject: string;
    subjectKeyId: string;
    pemCert: string;
    serialNumber: string;
    owner: string;
    approvals: string[];
}
export declare const ProposedCertificate: {
    encode(message: ProposedCertificate, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ProposedCertificate;
    fromJSON(object: any): ProposedCertificate;
    toJSON(message: ProposedCertificate): unknown;
    fromPartial(object: DeepPartial<ProposedCertificate>): ProposedCertificate;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
