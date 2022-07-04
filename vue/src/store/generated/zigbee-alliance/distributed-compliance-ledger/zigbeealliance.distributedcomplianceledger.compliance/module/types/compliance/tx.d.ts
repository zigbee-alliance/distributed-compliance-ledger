import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface MsgCertifyModel {
    signer: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    cDVersionNumber: number;
    certificationDate: string;
    certificationType: string;
    reason: string;
    programTypeVersion: string;
    cDCertificateId: string;
    familyId: string;
    supportedClusters: string;
    compliantPlatformUsed: string;
    compliantPlatformVersion: string;
    OSVersion: string;
    certificationRoute: string;
    programType: string;
    transport: string;
    parentChild: string;
    certificationIdOfSoftwareComponent: string;
}
export interface MsgCertifyModelResponse {
}
export interface MsgRevokeModel {
    signer: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    cDVersionNumber: number;
    revocationDate: string;
    certificationType: string;
    reason: string;
}
export interface MsgRevokeModelResponse {
}
export interface MsgProvisionModel {
    signer: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    cDVersionNumber: number;
    provisionalDate: string;
    certificationType: string;
    reason: string;
    programTypeVersion: string;
    cDCertificateId: string;
    familyId: string;
    supportedClusters: string;
    compliantPlatformUsed: string;
    compliantPlatformVersion: string;
    OSVersion: string;
    certificationRoute: string;
    programType: string;
    transport: string;
    parentChild: string;
    certificationIdOfSoftwareComponent: string;
}
export interface MsgProvisionModelResponse {
}
export declare const MsgCertifyModel: {
    encode(message: MsgCertifyModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCertifyModel;
    fromJSON(object: any): MsgCertifyModel;
    toJSON(message: MsgCertifyModel): unknown;
    fromPartial(object: DeepPartial<MsgCertifyModel>): MsgCertifyModel;
};
export declare const MsgCertifyModelResponse: {
    encode(_: MsgCertifyModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCertifyModelResponse;
    fromJSON(_: any): MsgCertifyModelResponse;
    toJSON(_: MsgCertifyModelResponse): unknown;
    fromPartial(_: DeepPartial<MsgCertifyModelResponse>): MsgCertifyModelResponse;
};
export declare const MsgRevokeModel: {
    encode(message: MsgRevokeModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRevokeModel;
    fromJSON(object: any): MsgRevokeModel;
    toJSON(message: MsgRevokeModel): unknown;
    fromPartial(object: DeepPartial<MsgRevokeModel>): MsgRevokeModel;
};
export declare const MsgRevokeModelResponse: {
    encode(_: MsgRevokeModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRevokeModelResponse;
    fromJSON(_: any): MsgRevokeModelResponse;
    toJSON(_: MsgRevokeModelResponse): unknown;
    fromPartial(_: DeepPartial<MsgRevokeModelResponse>): MsgRevokeModelResponse;
};
export declare const MsgProvisionModel: {
    encode(message: MsgProvisionModel, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProvisionModel;
    fromJSON(object: any): MsgProvisionModel;
    toJSON(message: MsgProvisionModel): unknown;
    fromPartial(object: DeepPartial<MsgProvisionModel>): MsgProvisionModel;
};
export declare const MsgProvisionModelResponse: {
    encode(_: MsgProvisionModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProvisionModelResponse;
    fromJSON(_: any): MsgProvisionModelResponse;
    toJSON(_: MsgProvisionModelResponse): unknown;
    fromPartial(_: DeepPartial<MsgProvisionModelResponse>): MsgProvisionModelResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CertifyModel(request: MsgCertifyModel): Promise<MsgCertifyModelResponse>;
    RevokeModel(request: MsgRevokeModel): Promise<MsgRevokeModelResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    ProvisionModel(request: MsgProvisionModel): Promise<MsgProvisionModelResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CertifyModel(request: MsgCertifyModel): Promise<MsgCertifyModelResponse>;
    RevokeModel(request: MsgRevokeModel): Promise<MsgRevokeModelResponse>;
    ProvisionModel(request: MsgProvisionModel): Promise<MsgProvisionModelResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
