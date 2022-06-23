import { ComplianceInfo } from '../compliance/compliance_info';
import { CertifiedModel } from '../compliance/certified_model';
import { RevokedModel } from '../compliance/revoked_model';
import { ProvisionalModel } from '../compliance/provisional_model';
import { DeviceSoftwareCompliance } from '../compliance/device_software_compliance';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
/** GenesisState defines the compliance module's genesis state. */
export interface GenesisState {
    complianceInfoList: ComplianceInfo[];
    certifiedModelList: CertifiedModel[];
    revokedModelList: RevokedModel[];
    provisionalModelList: ProvisionalModel[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    deviceSoftwareComplianceList: DeviceSoftwareCompliance[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
