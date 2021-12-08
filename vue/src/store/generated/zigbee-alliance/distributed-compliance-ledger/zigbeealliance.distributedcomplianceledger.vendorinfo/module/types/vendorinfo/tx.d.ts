import { Reader, Writer } from 'protobufjs/minimal';
import { VendorInfo } from '../vendorinfo/vendor_info';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface MsgCreateNewVendorInfo {
    creator: string;
    index: string;
    vendorInfo: VendorInfo | undefined;
}
export interface MsgCreateNewVendorInfoResponse {
}
export interface MsgUpdateNewVendorInfo {
    creator: string;
    index: string;
    vendorInfo: VendorInfo | undefined;
}
export interface MsgUpdateNewVendorInfoResponse {
}
export interface MsgDeleteNewVendorInfo {
    creator: string;
    index: string;
}
export interface MsgDeleteNewVendorInfoResponse {
}
export declare const MsgCreateNewVendorInfo: {
    encode(message: MsgCreateNewVendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateNewVendorInfo;
    fromJSON(object: any): MsgCreateNewVendorInfo;
    toJSON(message: MsgCreateNewVendorInfo): unknown;
    fromPartial(object: DeepPartial<MsgCreateNewVendorInfo>): MsgCreateNewVendorInfo;
};
export declare const MsgCreateNewVendorInfoResponse: {
    encode(_: MsgCreateNewVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateNewVendorInfoResponse;
    fromJSON(_: any): MsgCreateNewVendorInfoResponse;
    toJSON(_: MsgCreateNewVendorInfoResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateNewVendorInfoResponse>): MsgCreateNewVendorInfoResponse;
};
export declare const MsgUpdateNewVendorInfo: {
    encode(message: MsgUpdateNewVendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateNewVendorInfo;
    fromJSON(object: any): MsgUpdateNewVendorInfo;
    toJSON(message: MsgUpdateNewVendorInfo): unknown;
    fromPartial(object: DeepPartial<MsgUpdateNewVendorInfo>): MsgUpdateNewVendorInfo;
};
export declare const MsgUpdateNewVendorInfoResponse: {
    encode(_: MsgUpdateNewVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateNewVendorInfoResponse;
    fromJSON(_: any): MsgUpdateNewVendorInfoResponse;
    toJSON(_: MsgUpdateNewVendorInfoResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateNewVendorInfoResponse>): MsgUpdateNewVendorInfoResponse;
};
export declare const MsgDeleteNewVendorInfo: {
    encode(message: MsgDeleteNewVendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteNewVendorInfo;
    fromJSON(object: any): MsgDeleteNewVendorInfo;
    toJSON(message: MsgDeleteNewVendorInfo): unknown;
    fromPartial(object: DeepPartial<MsgDeleteNewVendorInfo>): MsgDeleteNewVendorInfo;
};
export declare const MsgDeleteNewVendorInfoResponse: {
    encode(_: MsgDeleteNewVendorInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteNewVendorInfoResponse;
    fromJSON(_: any): MsgDeleteNewVendorInfoResponse;
    toJSON(_: MsgDeleteNewVendorInfoResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteNewVendorInfoResponse>): MsgDeleteNewVendorInfoResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateNewVendorInfo(request: MsgCreateNewVendorInfo): Promise<MsgCreateNewVendorInfoResponse>;
    UpdateNewVendorInfo(request: MsgUpdateNewVendorInfo): Promise<MsgUpdateNewVendorInfoResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteNewVendorInfo(request: MsgDeleteNewVendorInfo): Promise<MsgDeleteNewVendorInfoResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateNewVendorInfo(request: MsgCreateNewVendorInfo): Promise<MsgCreateNewVendorInfoResponse>;
    UpdateNewVendorInfo(request: MsgUpdateNewVendorInfo): Promise<MsgUpdateNewVendorInfoResponse>;
    DeleteNewVendorInfo(request: MsgDeleteNewVendorInfo): Promise<MsgDeleteNewVendorInfoResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
