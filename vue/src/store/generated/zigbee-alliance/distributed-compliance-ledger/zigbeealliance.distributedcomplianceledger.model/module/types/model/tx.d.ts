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
    lsfUrl: string;
    lsfRevision: number;
}
export interface MsgCreateModelResponse {
}
export interface MsgUpdateModel {
    creator: string;
    vid: number;
    pid: number;
    productName: string;
    productLabel: string;
    partNumber: string;
    commissioningCustomFlowUrl: string;
    commissioningModeInitialStepsInstruction: string;
    commissioningModeSecondaryStepsInstruction: string;
    userManualUrl: string;
    supportUrl: string;
    productUrl: string;
    lsfUrl: string;
    lsfRevision: number;
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
    softwareVersionValid: boolean;
    otaUrl: string;
    minApplicableSoftwareVersion: number;
    maxApplicableSoftwareVersion: number;
    releaseNotesUrl: string;
}
export interface MsgUpdateModelVersionResponse {
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
/** Msg defines the Msg service. */
export interface Msg {
    CreateModel(request: MsgCreateModel): Promise<MsgCreateModelResponse>;
    UpdateModel(request: MsgUpdateModel): Promise<MsgUpdateModelResponse>;
    DeleteModel(request: MsgDeleteModel): Promise<MsgDeleteModelResponse>;
    CreateModelVersion(request: MsgCreateModelVersion): Promise<MsgCreateModelVersionResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateModelVersion(request: MsgUpdateModelVersion): Promise<MsgUpdateModelVersionResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateModel(request: MsgCreateModel): Promise<MsgCreateModelResponse>;
    UpdateModel(request: MsgUpdateModel): Promise<MsgUpdateModelResponse>;
    DeleteModel(request: MsgDeleteModel): Promise<MsgDeleteModelResponse>;
    CreateModelVersion(request: MsgCreateModelVersion): Promise<MsgCreateModelVersionResponse>;
    UpdateModelVersion(request: MsgUpdateModelVersion): Promise<MsgUpdateModelVersionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
