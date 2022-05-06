import { Reader, Writer } from 'protobufjs/minimal';
import { Any } from '../google/protobuf/any';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";
export interface MsgProposeAddAccount {
    signer: string;
    address: string;
    pubKey: Any | undefined;
    roles: string[];
    vendorID: number;
    info: string;
    time: number;
}
export interface MsgProposeAddAccountResponse {
}
export interface MsgApproveAddAccount {
    signer: string;
    address: string;
    info: string;
    time: number;
}
export interface MsgApproveAddAccountResponse {
}
export interface MsgProposeRevokeAccount {
    signer: string;
    address: string;
    info: string;
    time: number;
}
export interface MsgProposeRevokeAccountResponse {
}
export interface MsgApproveRevokeAccount {
    signer: string;
    address: string;
    info: string;
    time: number;
}
export interface MsgApproveRevokeAccountResponse {
}
export interface MsgRejectAddAccount {
    signer: string;
    address: string;
    info: string;
    time: number;
}
export interface MsgRejectAddAccountResponse {
}
export declare const MsgProposeAddAccount: {
    encode(message: MsgProposeAddAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeAddAccount;
    fromJSON(object: any): MsgProposeAddAccount;
    toJSON(message: MsgProposeAddAccount): unknown;
    fromPartial(object: DeepPartial<MsgProposeAddAccount>): MsgProposeAddAccount;
};
export declare const MsgProposeAddAccountResponse: {
    encode(_: MsgProposeAddAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeAddAccountResponse;
    fromJSON(_: any): MsgProposeAddAccountResponse;
    toJSON(_: MsgProposeAddAccountResponse): unknown;
    fromPartial(_: DeepPartial<MsgProposeAddAccountResponse>): MsgProposeAddAccountResponse;
};
export declare const MsgApproveAddAccount: {
    encode(message: MsgApproveAddAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveAddAccount;
    fromJSON(object: any): MsgApproveAddAccount;
    toJSON(message: MsgApproveAddAccount): unknown;
    fromPartial(object: DeepPartial<MsgApproveAddAccount>): MsgApproveAddAccount;
};
export declare const MsgApproveAddAccountResponse: {
    encode(_: MsgApproveAddAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveAddAccountResponse;
    fromJSON(_: any): MsgApproveAddAccountResponse;
    toJSON(_: MsgApproveAddAccountResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveAddAccountResponse>): MsgApproveAddAccountResponse;
};
export declare const MsgProposeRevokeAccount: {
    encode(message: MsgProposeRevokeAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeAccount;
    fromJSON(object: any): MsgProposeRevokeAccount;
    toJSON(message: MsgProposeRevokeAccount): unknown;
    fromPartial(object: DeepPartial<MsgProposeRevokeAccount>): MsgProposeRevokeAccount;
};
export declare const MsgProposeRevokeAccountResponse: {
    encode(_: MsgProposeRevokeAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeAccountResponse;
    fromJSON(_: any): MsgProposeRevokeAccountResponse;
    toJSON(_: MsgProposeRevokeAccountResponse): unknown;
    fromPartial(_: DeepPartial<MsgProposeRevokeAccountResponse>): MsgProposeRevokeAccountResponse;
};
export declare const MsgApproveRevokeAccount: {
    encode(message: MsgApproveRevokeAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeAccount;
    fromJSON(object: any): MsgApproveRevokeAccount;
    toJSON(message: MsgApproveRevokeAccount): unknown;
    fromPartial(object: DeepPartial<MsgApproveRevokeAccount>): MsgApproveRevokeAccount;
};
export declare const MsgApproveRevokeAccountResponse: {
    encode(_: MsgApproveRevokeAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeAccountResponse;
    fromJSON(_: any): MsgApproveRevokeAccountResponse;
    toJSON(_: MsgApproveRevokeAccountResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveRevokeAccountResponse>): MsgApproveRevokeAccountResponse;
};
export declare const MsgRejectAddAccount: {
    encode(message: MsgRejectAddAccount, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectAddAccount;
    fromJSON(object: any): MsgRejectAddAccount;
    toJSON(message: MsgRejectAddAccount): unknown;
    fromPartial(object: DeepPartial<MsgRejectAddAccount>): MsgRejectAddAccount;
};
export declare const MsgRejectAddAccountResponse: {
    encode(_: MsgRejectAddAccountResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectAddAccountResponse;
    fromJSON(_: any): MsgRejectAddAccountResponse;
    toJSON(_: MsgRejectAddAccountResponse): unknown;
    fromPartial(_: DeepPartial<MsgRejectAddAccountResponse>): MsgRejectAddAccountResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    ProposeAddAccount(request: MsgProposeAddAccount): Promise<MsgProposeAddAccountResponse>;
    ApproveAddAccount(request: MsgApproveAddAccount): Promise<MsgApproveAddAccountResponse>;
    ProposeRevokeAccount(request: MsgProposeRevokeAccount): Promise<MsgProposeRevokeAccountResponse>;
    ApproveRevokeAccount(request: MsgApproveRevokeAccount): Promise<MsgApproveRevokeAccountResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RejectAddAccount(request: MsgRejectAddAccount): Promise<MsgRejectAddAccountResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    ProposeAddAccount(request: MsgProposeAddAccount): Promise<MsgProposeAddAccountResponse>;
    ApproveAddAccount(request: MsgApproveAddAccount): Promise<MsgApproveAddAccountResponse>;
    ProposeRevokeAccount(request: MsgProposeRevokeAccount): Promise<MsgProposeRevokeAccountResponse>;
    ApproveRevokeAccount(request: MsgApproveRevokeAccount): Promise<MsgApproveRevokeAccountResponse>;
    RejectAddAccount(request: MsgRejectAddAccount): Promise<MsgRejectAddAccountResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
