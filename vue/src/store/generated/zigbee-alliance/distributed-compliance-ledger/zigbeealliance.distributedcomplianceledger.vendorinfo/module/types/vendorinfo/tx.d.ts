import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface MsgCreateVendorInfo {
    creator: string;
    vendorID: number;
    vendorName: string;
    companyLegalName: string;
    companyPrefferedName: string;
    vendorLandingPageURL: string;
}
export interface MsgCreateVendorInfoResponse {
}
export interface MsgUpdateVendorInfo {
    creator: string;
    vendorID: number;
    vendorName: string;
    companyLegalName: string;
    companyPrefferedName: string;
    vendorLandingPageURL: string;
}
export interface MsgUpdateVendorInfoResponse {
}
export interface MsgDeleteVendorInfo {
    creator: string;
    vendorID: number;
}
export interface MsgDeleteVendorInfoResponse {
}
export declare const MsgCreateVendorInfo: {
    encode(message: MsgCreateVendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfo;
    fromJSON(object: any): MsgCreateVendorInfo;
    toJSON(message: MsgCreateVendorInfo): unknown;
    fromPartial(object: DeepPartial<MsgCreateVendorInfo>): MsgCreateVendorInfo;
};
export declare const MsgCreateVendorInfoResponse: {
    encode(_: MsgCreateVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfoResponse;
    fromJSON(_: any): MsgCreateVendorInfoResponse;
    toJSON(_: MsgCreateVendorInfoResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateVendorInfoResponse>): MsgCreateVendorInfoResponse;
};
export declare const MsgUpdateVendorInfo: {
    encode(message: MsgUpdateVendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfo;
    fromJSON(object: any): MsgUpdateVendorInfo;
    toJSON(message: MsgUpdateVendorInfo): unknown;
    fromPartial(object: DeepPartial<MsgUpdateVendorInfo>): MsgUpdateVendorInfo;
};
export declare const MsgUpdateVendorInfoResponse: {
    encode(_: MsgUpdateVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfoResponse;
    fromJSON(_: any): MsgUpdateVendorInfoResponse;
    toJSON(_: MsgUpdateVendorInfoResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateVendorInfoResponse>): MsgUpdateVendorInfoResponse;
};
export declare const MsgDeleteVendorInfo: {
    encode(message: MsgDeleteVendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfo;
    fromJSON(object: any): MsgDeleteVendorInfo;
    toJSON(message: MsgDeleteVendorInfo): unknown;
    fromPartial(object: DeepPartial<MsgDeleteVendorInfo>): MsgDeleteVendorInfo;
};
export declare const MsgDeleteVendorInfoResponse: {
    encode(_: MsgDeleteVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfoResponse;
    fromJSON(_: any): MsgDeleteVendorInfoResponse;
    toJSON(_: MsgDeleteVendorInfoResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteVendorInfoResponse>): MsgDeleteVendorInfoResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateVendorInfo(request: MsgCreateVendorInfo): Promise<MsgCreateVendorInfoResponse>;
    UpdateVendorInfo(request: MsgUpdateVendorInfo): Promise<MsgUpdateVendorInfoResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteVendorInfo(request: MsgDeleteVendorInfo): Promise<MsgDeleteVendorInfoResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateVendorInfo(request: MsgCreateVendorInfo): Promise<MsgCreateVendorInfoResponse>;
    UpdateVendorInfo(request: MsgUpdateVendorInfo): Promise<MsgUpdateVendorInfoResponse>;
    DeleteVendorInfo(request: MsgDeleteVendorInfo): Promise<MsgDeleteVendorInfoResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
