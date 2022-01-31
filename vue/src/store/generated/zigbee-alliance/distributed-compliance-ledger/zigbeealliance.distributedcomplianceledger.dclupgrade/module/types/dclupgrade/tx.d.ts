import { Reader, Writer } from 'protobufjs/minimal';
import { Plan } from '../cosmos/upgrade/v1beta1/upgrade';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclupgrade";
export interface MsgProposeUpgrade {
    creator: string;
    plan: Plan | undefined;
}
export interface MsgProposeUpgradeResponse {
}
export interface MsgApproveUpgrade {
    creator: string;
    name: string;
}
export interface MsgApproveUpgradeResponse {
}
export declare const MsgProposeUpgrade: {
    encode(message: MsgProposeUpgrade, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeUpgrade;
    fromJSON(object: any): MsgProposeUpgrade;
    toJSON(message: MsgProposeUpgrade): unknown;
    fromPartial(object: DeepPartial<MsgProposeUpgrade>): MsgProposeUpgrade;
};
export declare const MsgProposeUpgradeResponse: {
    encode(_: MsgProposeUpgradeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeUpgradeResponse;
    fromJSON(_: any): MsgProposeUpgradeResponse;
    toJSON(_: MsgProposeUpgradeResponse): unknown;
    fromPartial(_: DeepPartial<MsgProposeUpgradeResponse>): MsgProposeUpgradeResponse;
};
export declare const MsgApproveUpgrade: {
    encode(message: MsgApproveUpgrade, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveUpgrade;
    fromJSON(object: any): MsgApproveUpgrade;
    toJSON(message: MsgApproveUpgrade): unknown;
    fromPartial(object: DeepPartial<MsgApproveUpgrade>): MsgApproveUpgrade;
};
export declare const MsgApproveUpgradeResponse: {
    encode(_: MsgApproveUpgradeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveUpgradeResponse;
    fromJSON(_: any): MsgApproveUpgradeResponse;
    toJSON(_: MsgApproveUpgradeResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveUpgradeResponse>): MsgApproveUpgradeResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse>;
    ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
