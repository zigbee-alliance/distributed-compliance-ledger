import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
export interface Model {
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
    creator: string;
}
export declare const Model: {
    encode(message: Model, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Model;
    fromJSON(object: any): Model;
    toJSON(message: Model): unknown;
    fromPartial(object: DeepPartial<Model>): Model;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
