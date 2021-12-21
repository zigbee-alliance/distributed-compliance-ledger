import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface MsgCreateVendorInfoType {
    creator: string;
    vendorID: number;
    vendorName: string;
    companyLegalName: string;
    companyPrefferedName: string;
    vendorLandingPageURL: string;
}
export interface MsgCreateVendorInfoTypeResponse {
}
export interface MsgUpdateVendorInfoType {
    creator: string;
    vendorID: number;
    vendorName: string;
    companyLegalName: string;
    companyPrefferedName: string;
    vendorLandingPageURL: string;
}
export interface MsgUpdateVendorInfoTypeResponse {
}
export interface MsgDeleteVendorInfoType {
    creator: string;
    vendorID: number;
}
export interface MsgDeleteVendorInfoTypeResponse {
}
export declare const MsgCreateVendorInfoType: {
    encode(message: MsgCreateVendorInfoType, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfoType;
    fromJSON(object: any): MsgCreateVendorInfoType;
    toJSON(message: MsgCreateVendorInfoType): unknown;
    fromPartial(object: DeepPartial<MsgCreateVendorInfoType>): MsgCreateVendorInfoType;
};
export declare const MsgCreateVendorInfoTypeResponse: {
    encode(_: MsgCreateVendorInfoTypeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfoTypeResponse;
    fromJSON(_: any): MsgCreateVendorInfoTypeResponse;
    toJSON(_: MsgCreateVendorInfoTypeResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateVendorInfoTypeResponse>): MsgCreateVendorInfoTypeResponse;
};
export declare const MsgUpdateVendorInfoType: {
    encode(message: MsgUpdateVendorInfoType, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfoType;
    fromJSON(object: any): MsgUpdateVendorInfoType;
    toJSON(message: MsgUpdateVendorInfoType): unknown;
    fromPartial(object: DeepPartial<MsgUpdateVendorInfoType>): MsgUpdateVendorInfoType;
};
export declare const MsgUpdateVendorInfoTypeResponse: {
    encode(_: MsgUpdateVendorInfoTypeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfoTypeResponse;
    fromJSON(_: any): MsgUpdateVendorInfoTypeResponse;
    toJSON(_: MsgUpdateVendorInfoTypeResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateVendorInfoTypeResponse>): MsgUpdateVendorInfoTypeResponse;
};
export declare const MsgDeleteVendorInfoType: {
    encode(message: MsgDeleteVendorInfoType, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfoType;
    fromJSON(object: any): MsgDeleteVendorInfoType;
    toJSON(message: MsgDeleteVendorInfoType): unknown;
    fromPartial(object: DeepPartial<MsgDeleteVendorInfoType>): MsgDeleteVendorInfoType;
};
export declare const MsgDeleteVendorInfoTypeResponse: {
    encode(_: MsgDeleteVendorInfoTypeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfoTypeResponse;
    fromJSON(_: any): MsgDeleteVendorInfoTypeResponse;
    toJSON(_: MsgDeleteVendorInfoTypeResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteVendorInfoTypeResponse>): MsgDeleteVendorInfoTypeResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateVendorInfoType(request: MsgCreateVendorInfoType): Promise<MsgCreateVendorInfoTypeResponse>;
    UpdateVendorInfoType(request: MsgUpdateVendorInfoType): Promise<MsgUpdateVendorInfoTypeResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteVendorInfoType(request: MsgDeleteVendorInfoType): Promise<MsgDeleteVendorInfoTypeResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateVendorInfoType(request: MsgCreateVendorInfoType): Promise<MsgCreateVendorInfoTypeResponse>;
    UpdateVendorInfoType(request: MsgUpdateVendorInfoType): Promise<MsgUpdateVendorInfoTypeResponse>;
    DeleteVendorInfoType(request: MsgDeleteVendorInfoType): Promise<MsgDeleteVendorInfoTypeResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
