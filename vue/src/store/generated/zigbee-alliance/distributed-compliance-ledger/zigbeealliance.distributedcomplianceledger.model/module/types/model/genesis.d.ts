import { VendorProducts } from '../model/vendor_products';
import { Model } from '../model/model';
import { ModelVersion } from '../model/model_version';
import { ModelVersions } from '../model/model_versions';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";
/** GenesisState defines the model module's genesis state. */
export interface GenesisState {
    vendorProductsList: VendorProducts[];
    modelList: Model[];
    modelVersionList: ModelVersion[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    modelVersionsList: ModelVersions[];
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
