import { Product } from '../model/product';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
export interface VendorProducts {
    vid: number;
    products: Product[];
}
export declare const VendorProducts: {
    encode(message: VendorProducts, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VendorProducts;
    fromJSON(object: any): VendorProducts;
    toJSON(message: VendorProducts): unknown;
    fromPartial(object: DeepPartial<VendorProducts>): VendorProducts;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
