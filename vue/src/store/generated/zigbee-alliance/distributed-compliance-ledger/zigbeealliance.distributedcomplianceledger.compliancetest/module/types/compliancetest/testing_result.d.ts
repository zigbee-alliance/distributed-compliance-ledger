import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliancetest";
export interface TestingResult {
    vid: number;
    pid: number;
    softwareVersion: number;
    softwareVersionString: string;
    owner: string;
    testResult: string;
    testDate: string;
}
export declare const TestingResult: {
    encode(message: TestingResult, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): TestingResult;
    fromJSON(object: any): TestingResult;
    toJSON(message: TestingResult): unknown;
    fromPartial(object: DeepPartial<TestingResult>): TestingResult;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
