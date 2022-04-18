import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface VendorInfo {
    vendorID: number;
    vendorName: string;
    companyLegalName: string;
    companyPreferredName: string;
    vendorLandingPageURL: string;
    creator: string;
}
export declare const VendorInfo: {
    encode(message: VendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VendorInfo;
    fromJSON(object: any): VendorInfo;
    toJSON(message: VendorInfo): unknown;
    fromPartial(object: DeepPartial<VendorInfo>): VendorInfo;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
