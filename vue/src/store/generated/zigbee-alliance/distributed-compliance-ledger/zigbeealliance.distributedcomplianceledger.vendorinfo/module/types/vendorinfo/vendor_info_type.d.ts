import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface VendorInfoType {
    vendorID: number;
    vendorName: string;
    companyLegalName: string;
    companyPrefferedName: string;
    vendorLandingPageURL: string;
    creator: string;
}
export declare const VendorInfoType: {
    encode(message: VendorInfoType, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VendorInfoType;
    fromJSON(object: any): VendorInfoType;
    toJSON(message: VendorInfoType): unknown;
    fromPartial(object: DeepPartial<VendorInfoType>): VendorInfoType;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
