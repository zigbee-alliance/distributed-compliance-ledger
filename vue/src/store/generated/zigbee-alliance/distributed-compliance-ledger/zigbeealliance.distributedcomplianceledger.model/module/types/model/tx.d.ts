import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
export interface MsgCreateModel {
    creator: string;
    vid: number;
    pid: number;
    deviceTypeId: number;
    productName: string;
    productLabel: string;
    partNumber: string;
    commissioningCustomFlow: number;
    commissioningCustomFlowUrl: string;
    commissioningModeInitialStepsHint: number;
    commissioningModeInitialStepsInstruction: string;
    commissioningModeSecondaryStepsHint: number;
    commissioningModeSecondaryStepsInstruction: string;
    userManualUrl: string;
    supportUrl: string;
    productUrl: string;
}
export interface MsgCreateModelResponse {
}
export interface MsgUpdateModel {
    creator: string;
    vid: number;
    pid: number;
    deviceTypeId: number;
    productName: string;
    productLabel: string;
    partNumber: string;
    commissioningCustomFlow: number;
    commissioningCustomFlowUrl: string;
    commissioningModeInitialStepsHint: number;
    commissioningModeInitialStepsInstruction: string;
    commissioningModeSecondaryStepsHint: number;
    commissioningModeSecondaryStepsInstruction: string;
    userManualUrl: string;
    supportUrl: string;
    productUrl: string;
}
export interface MsgUpdateModelResponse {
}
export interface MsgDeleteModel {
    creator: string;
    vid: number;
    pid: number;
}
export interface MsgDeleteModelResponse {
}
export interface MsgCreateModelVersion {
    creator: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    cdVersionNumber: number;
    firmwareDigests: string;
    softwareVersionValid: boolean;
    otaUrl: string;
    otaFileSize: number;
    otaChecksum: string;
    otaChecksumType: number;
    minApplicableSoftwareVersion: number;
    maxApplicableSoftwareVersion: number;
    releaseNotesUrl: string;
}
export interface MsgCreateModelVersionResponse {
}
export interface MsgUpdateModelVersion {
    creator: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    cdVersionNumber: number;
    firmwareDigests: string;
    softwareVersionValid: boolean;
    otaUrl: string;
    otaFileSize: number;
    otaChecksum: string;
    otaChecksumType: number;
    minApplicableSoftwareVersion: number;
    maxApplicableSoftwareVersion: number;
    releaseNotesUrl: string;
}
export interface MsgUpdateModelVersionResponse {
}
export interface MsgDeleteModelVersion {
    creator: string;
    vid: number;
    pid: number;
    softwareVersion: number;
}
export interface MsgDeleteModelVersionResponse {
}
export interface MsgCreateModelVersions {
    creator: string;
    vid: number;
    pid: number;
    softwareVersions: number[];
}
export interface MsgCreateModelVersionsResponse {
}
export interface MsgUpdateModelVersions {
    creator: string;
    vid: number;
    pid: number;
    softwareVersions: number[];
}
export interface MsgUpdateModelVersionsResponse {
}
export interface MsgDeleteModelVersions {
    creator: string;
    vid: number;
    pid: number;
}
export interface MsgDeleteModelVersionsResponse {
}
export declare const MsgCreateModel: {
    encode(message: MsgCreateModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateModel;
    fromJSON(object: any): MsgCreateModel;
    toJSON(message: MsgCreateModel): unknown;
    fromPartial(object: DeepPartial<MsgCreateModel>): MsgCreateModel;
};
export declare const MsgCreateModelResponse: {
    encode(_: MsgCreateModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateModelResponse;
    fromJSON(_: any): MsgCreateModelResponse;
    toJSON(_: MsgCreateModelResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateModelResponse>): MsgCreateModelResponse;
};
export declare const MsgUpdateModel: {
    encode(message: MsgUpdateModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateModel;
    fromJSON(object: any): MsgUpdateModel;
    toJSON(message: MsgUpdateModel): unknown;
    fromPartial(object: DeepPartial<MsgUpdateModel>): MsgUpdateModel;
};
export declare const MsgUpdateModelResponse: {
    encode(_: MsgUpdateModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateModelResponse;
    fromJSON(_: any): MsgUpdateModelResponse;
    toJSON(_: MsgUpdateModelResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateModelResponse>): MsgUpdateModelResponse;
};
export declare const MsgDeleteModel: {
    encode(message: MsgDeleteModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteModel;
    fromJSON(object: any): MsgDeleteModel;
    toJSON(message: MsgDeleteModel): unknown;
    fromPartial(object: DeepPartial<MsgDeleteModel>): MsgDeleteModel;
};
export declare const MsgDeleteModelResponse: {
    encode(_: MsgDeleteModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteModelResponse;
    fromJSON(_: any): MsgDeleteModelResponse;
    toJSON(_: MsgDeleteModelResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteModelResponse>): MsgDeleteModelResponse;
};
export declare const MsgCreateModelVersion: {
    encode(message: MsgCreateModelVersion, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateModelVersion;
    fromJSON(object: any): MsgCreateModelVersion;
    toJSON(message: MsgCreateModelVersion): unknown;
    fromPartial(object: DeepPartial<MsgCreateModelVersion>): MsgCreateModelVersion;
};
export declare const MsgCreateModelVersionResponse: {
    encode(_: MsgCreateModelVersionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateModelVersionResponse;
    fromJSON(_: any): MsgCreateModelVersionResponse;
    toJSON(_: MsgCreateModelVersionResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateModelVersionResponse>): MsgCreateModelVersionResponse;
};
export declare const MsgUpdateModelVersion: {
    encode(message: MsgUpdateModelVersion, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateModelVersion;
    fromJSON(object: any): MsgUpdateModelVersion;
    toJSON(message: MsgUpdateModelVersion): unknown;
    fromPartial(object: DeepPartial<MsgUpdateModelVersion>): MsgUpdateModelVersion;
};
export declare const MsgUpdateModelVersionResponse: {
    encode(_: MsgUpdateModelVersionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateModelVersionResponse;
    fromJSON(_: any): MsgUpdateModelVersionResponse;
    toJSON(_: MsgUpdateModelVersionResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateModelVersionResponse>): MsgUpdateModelVersionResponse;
};
export declare const MsgDeleteModelVersion: {
    encode(message: MsgDeleteModelVersion, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteModelVersion;
    fromJSON(object: any): MsgDeleteModelVersion;
    toJSON(message: MsgDeleteModelVersion): unknown;
    fromPartial(object: DeepPartial<MsgDeleteModelVersion>): MsgDeleteModelVersion;
};
export declare const MsgDeleteModelVersionResponse: {
    encode(_: MsgDeleteModelVersionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteModelVersionResponse;
    fromJSON(_: any): MsgDeleteModelVersionResponse;
    toJSON(_: MsgDeleteModelVersionResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteModelVersionResponse>): MsgDeleteModelVersionResponse;
};
export declare const MsgCreateModelVersions: {
    encode(message: MsgCreateModelVersions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateModelVersions;
    fromJSON(object: any): MsgCreateModelVersions;
    toJSON(message: MsgCreateModelVersions): unknown;
    fromPartial(object: DeepPartial<MsgCreateModelVersions>): MsgCreateModelVersions;
};
export declare const MsgCreateModelVersionsResponse: {
    encode(_: MsgCreateModelVersionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateModelVersionsResponse;
    fromJSON(_: any): MsgCreateModelVersionsResponse;
    toJSON(_: MsgCreateModelVersionsResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateModelVersionsResponse>): MsgCreateModelVersionsResponse;
};
export declare const MsgUpdateModelVersions: {
    encode(message: MsgUpdateModelVersions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateModelVersions;
    fromJSON(object: any): MsgUpdateModelVersions;
    toJSON(message: MsgUpdateModelVersions): unknown;
    fromPartial(object: DeepPartial<MsgUpdateModelVersions>): MsgUpdateModelVersions;
};
export declare const MsgUpdateModelVersionsResponse: {
    encode(_: MsgUpdateModelVersionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateModelVersionsResponse;
    fromJSON(_: any): MsgUpdateModelVersionsResponse;
    toJSON(_: MsgUpdateModelVersionsResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateModelVersionsResponse>): MsgUpdateModelVersionsResponse;
};
export declare const MsgDeleteModelVersions: {
    encode(message: MsgDeleteModelVersions, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteModelVersions;
    fromJSON(object: any): MsgDeleteModelVersions;
    toJSON(message: MsgDeleteModelVersions): unknown;
    fromPartial(object: DeepPartial<MsgDeleteModelVersions>): MsgDeleteModelVersions;
};
export declare const MsgDeleteModelVersionsResponse: {
    encode(_: MsgDeleteModelVersionsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteModelVersionsResponse;
    fromJSON(_: any): MsgDeleteModelVersionsResponse;
    toJSON(_: MsgDeleteModelVersionsResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteModelVersionsResponse>): MsgDeleteModelVersionsResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateModel(request: MsgCreateModel): Promise<MsgCreateModelResponse>;
    UpdateModel(request: MsgUpdateModel): Promise<MsgUpdateModelResponse>;
    DeleteModel(request: MsgDeleteModel): Promise<MsgDeleteModelResponse>;
    CreateModelVersion(request: MsgCreateModelVersion): Promise<MsgCreateModelVersionResponse>;
    UpdateModelVersion(request: MsgUpdateModelVersion): Promise<MsgUpdateModelVersionResponse>;
    DeleteModelVersion(request: MsgDeleteModelVersion): Promise<MsgDeleteModelVersionResponse>;
    CreateModelVersions(request: MsgCreateModelVersions): Promise<MsgCreateModelVersionsResponse>;
    UpdateModelVersions(request: MsgUpdateModelVersions): Promise<MsgUpdateModelVersionsResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteModelVersions(request: MsgDeleteModelVersions): Promise<MsgDeleteModelVersionsResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateModel(request: MsgCreateModel): Promise<MsgCreateModelResponse>;
    UpdateModel(request: MsgUpdateModel): Promise<MsgUpdateModelResponse>;
    DeleteModel(request: MsgDeleteModel): Promise<MsgDeleteModelResponse>;
    CreateModelVersion(request: MsgCreateModelVersion): Promise<MsgCreateModelVersionResponse>;
    UpdateModelVersion(request: MsgUpdateModelVersion): Promise<MsgUpdateModelVersionResponse>;
    DeleteModelVersion(request: MsgDeleteModelVersion): Promise<MsgDeleteModelVersionResponse>;
    CreateModelVersions(request: MsgCreateModelVersions): Promise<MsgCreateModelVersionsResponse>;
    UpdateModelVersions(request: MsgUpdateModelVersions): Promise<MsgUpdateModelVersionsResponse>;
    DeleteModelVersions(request: MsgDeleteModelVersions): Promise<MsgDeleteModelVersionsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
