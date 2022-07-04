import { ComplianceHistoryItem } from '../compliance/compliance_history_item';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface ComplianceInfo {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
    softwareVersionString: string;
    cDVersionNumber: number;
    softwareVersionCertificationStatus: number;
    date: string;
    reason: string;
    owner: string;
    history: ComplianceHistoryItem[];
    cDCertificateId: string;
    certificationRoute: string;
    programType: string;
    programTypeVersion: string;
    compliantPlatformUsed: string;
    compliantPlatformVersion: string;
    transport: string;
    familyId: string;
    supportedClusters: string;
    OSVersion: string;
    parentChild: string;
    certificationIdOfSoftwareComponent: string;
}
export declare const ComplianceInfo: {
    encode(message: ComplianceInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ComplianceInfo;
    fromJSON(object: any): ComplianceInfo;
    toJSON(message: ComplianceInfo): unknown;
    fromPartial(object: DeepPartial<ComplianceInfo>): ComplianceInfo;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
