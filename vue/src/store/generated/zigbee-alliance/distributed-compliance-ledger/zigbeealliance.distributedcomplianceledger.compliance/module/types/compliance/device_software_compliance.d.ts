import { ComplianceInfo } from '../compliance/compliance_info';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface DeviceSoftwareCompliance {
    cDCertificateId: string;
    complianceInfo: ComplianceInfo[];
}
export declare const DeviceSoftwareCompliance: {
    encode(message: DeviceSoftwareCompliance, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DeviceSoftwareCompliance;
    fromJSON(object: any): DeviceSoftwareCompliance;
    toJSON(message: DeviceSoftwareCompliance): unknown;
    fromPartial(object: DeepPartial<DeviceSoftwareCompliance>): DeviceSoftwareCompliance;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
