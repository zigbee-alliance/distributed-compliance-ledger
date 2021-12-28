import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliancetest";
export interface MsgAddTestingResult {
    signer: string;
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    testResult: string;
    testDate: string;
}
export interface MsgAddTestingResultResponse {
}
export declare const MsgAddTestingResult: {
    encode(message: MsgAddTestingResult, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddTestingResult;
    fromJSON(object: any): MsgAddTestingResult;
    toJSON(message: MsgAddTestingResult): unknown;
    fromPartial(object: DeepPartial<MsgAddTestingResult>): MsgAddTestingResult;
};
export declare const MsgAddTestingResultResponse: {
    encode(_: MsgAddTestingResultResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddTestingResultResponse;
    fromJSON(_: any): MsgAddTestingResultResponse;
    toJSON(_: MsgAddTestingResultResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddTestingResultResponse>): MsgAddTestingResultResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    AddTestingResult(request: MsgAddTestingResult): Promise<MsgAddTestingResultResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    AddTestingResult(request: MsgAddTestingResult): Promise<MsgAddTestingResultResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
