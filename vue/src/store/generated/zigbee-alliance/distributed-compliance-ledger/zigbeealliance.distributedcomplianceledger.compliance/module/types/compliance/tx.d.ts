import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface MsgCertifyModel {
    signer: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    certificationDate: string;
    certificationType: string;
    reason: string;
}
export interface MsgCertifyModelResponse {
}
export interface MsgRevokeModel {
    signer: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    revocationDate: string;
    certificationType: string;
    reason: string;
}
export interface MsgRevokeModelResponse {
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
/** Msg defines the Msg service. */
export interface Msg {
    CertifyModel(request: MsgCertifyModel): Promise<MsgCertifyModelResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RevokeModel(request: MsgRevokeModel): Promise<MsgRevokeModelResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CertifyModel(request: MsgCertifyModel): Promise<MsgCertifyModelResponse>;
    RevokeModel(request: MsgRevokeModel): Promise<MsgRevokeModelResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
