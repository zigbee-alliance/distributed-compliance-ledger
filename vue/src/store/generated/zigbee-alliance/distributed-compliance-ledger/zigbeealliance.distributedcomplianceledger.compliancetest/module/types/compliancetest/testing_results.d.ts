import { TestingResult } from '../compliancetest/testing_result';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliancetest";
export interface TestingResults {
    vid: number;
    pid: number;
    softwareVersion: number;
    results: TestingResult[];
    softwareVersionString: string;
}
export declare const TestingResults: {
    encode(message: TestingResults, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): TestingResults;
    fromJSON(object: any): TestingResults;
    toJSON(message: TestingResults): unknown;
    fromPartial(object: DeepPartial<TestingResults>): TestingResults;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
