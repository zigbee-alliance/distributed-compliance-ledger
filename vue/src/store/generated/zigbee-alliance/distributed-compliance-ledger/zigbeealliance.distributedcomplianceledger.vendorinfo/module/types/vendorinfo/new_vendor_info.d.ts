import { VendorInfo } from '../vendorinfo/vendor_info';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";
export interface NewVendorInfo {
    index: string;
    vendorInfo: VendorInfo | undefined;
    creator: string;
}
export declare const NewVendorInfo: {
    encode(message: NewVendorInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NewVendorInfo;
    fromJSON(object: any): NewVendorInfo;
    toJSON(message: NewVendorInfo): unknown;
    fromPartial(object: DeepPartial<NewVendorInfo>): NewVendorInfo;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
