import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface UniqueCertificate {
    issuer: string;
    serialNumber: string;
    present: boolean;
}
export declare const UniqueCertificate: {
    encode(message: UniqueCertificate, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): UniqueCertificate;
    fromJSON(object: any): UniqueCertificate;
    toJSON(message: UniqueCertificate): unknown;
    fromPartial(object: DeepPartial<UniqueCertificate>): UniqueCertificate;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
