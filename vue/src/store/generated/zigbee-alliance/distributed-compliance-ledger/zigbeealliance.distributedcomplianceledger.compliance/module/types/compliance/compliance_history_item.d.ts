import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface ComplianceHistoryItem {
    softwareVersionCertificationStatus: number;
    date: string;
    reason: string;
    cDVersionNumber: number;
}
export declare const ComplianceHistoryItem: {
    encode(message: ComplianceHistoryItem, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ComplianceHistoryItem;
    fromJSON(object: any): ComplianceHistoryItem;
    toJSON(message: ComplianceHistoryItem): unknown;
    fromPartial(object: DeepPartial<ComplianceHistoryItem>): ComplianceHistoryItem;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
