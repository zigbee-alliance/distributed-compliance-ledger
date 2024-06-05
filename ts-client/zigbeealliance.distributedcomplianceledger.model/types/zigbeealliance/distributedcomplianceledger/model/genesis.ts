/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Model } from "./model";
import { ModelVersion } from "./model_version";
import { ModelVersions } from "./model_versions";
import { VendorProducts } from "./vendor_products";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";

/** GenesisState defines the model module's genesis state. */
export interface GenesisState {
  vendorProductsList: VendorProducts[];
  modelList: Model[];
  modelVersionList: ModelVersion[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  modelVersionsList: ModelVersions[];
}

function createBaseGenesisState(): GenesisState {
  return { vendorProductsList: [], modelList: [], modelVersionList: [], modelVersionsList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.vendorProductsList) {
      VendorProducts.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.modelList) {
      Model.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.modelVersionList) {
      ModelVersion.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.modelVersionsList) {
      ModelVersions.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vendorProductsList.push(VendorProducts.decode(reader, reader.uint32()));
          break;
        case 2:
          message.modelList.push(Model.decode(reader, reader.uint32()));
          break;
        case 3:
          message.modelVersionList.push(ModelVersion.decode(reader, reader.uint32()));
          break;
        case 4:
          message.modelVersionsList.push(ModelVersions.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      vendorProductsList: Array.isArray(object?.vendorProductsList)
        ? object.vendorProductsList.map((e: any) => VendorProducts.fromJSON(e))
        : [],
      modelList: Array.isArray(object?.modelList) ? object.modelList.map((e: any) => Model.fromJSON(e)) : [],
      modelVersionList: Array.isArray(object?.modelVersionList)
        ? object.modelVersionList.map((e: any) => ModelVersion.fromJSON(e))
        : [],
      modelVersionsList: Array.isArray(object?.modelVersionsList)
        ? object.modelVersionsList.map((e: any) => ModelVersions.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.vendorProductsList) {
      obj.vendorProductsList = message.vendorProductsList.map((e) => e ? VendorProducts.toJSON(e) : undefined);
    } else {
      obj.vendorProductsList = [];
    }
    if (message.modelList) {
      obj.modelList = message.modelList.map((e) => e ? Model.toJSON(e) : undefined);
    } else {
      obj.modelList = [];
    }
    if (message.modelVersionList) {
      obj.modelVersionList = message.modelVersionList.map((e) => e ? ModelVersion.toJSON(e) : undefined);
    } else {
      obj.modelVersionList = [];
    }
    if (message.modelVersionsList) {
      obj.modelVersionsList = message.modelVersionsList.map((e) => e ? ModelVersions.toJSON(e) : undefined);
    } else {
      obj.modelVersionsList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.vendorProductsList = object.vendorProductsList?.map((e) => VendorProducts.fromPartial(e)) || [];
    message.modelList = object.modelList?.map((e) => Model.fromPartial(e)) || [];
    message.modelVersionList = object.modelVersionList?.map((e) => ModelVersion.fromPartial(e)) || [];
    message.modelVersionsList = object.modelVersionsList?.map((e) => ModelVersions.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };
