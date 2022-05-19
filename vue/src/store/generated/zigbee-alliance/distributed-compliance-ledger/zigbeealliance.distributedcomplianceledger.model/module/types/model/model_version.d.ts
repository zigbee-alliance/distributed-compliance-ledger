import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
export interface ModelVersion {
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    cdVersionNumber: number;
    firmwareInformation: string;
    softwareVersionValid: boolean;
    otaUrl: string;
    otaFileSize: number;
    otaChecksum: string;
    otaChecksumType: number;
    minApplicableSoftwareVersion: number;
    maxApplicableSoftwareVersion: number;
    releaseNotesUrl: string;
    creator: string;
}
export declare const ModelVersion: {
    encode(message: ModelVersion, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ModelVersion;
    fromJSON(object: any): ModelVersion;
    toJSON(message: ModelVersion): unknown;
    fromPartial(object: DeepPartial<ModelVersion>): ModelVersion;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
