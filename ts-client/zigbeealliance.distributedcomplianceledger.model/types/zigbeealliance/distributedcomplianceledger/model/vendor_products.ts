/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Product } from "./product";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";

export interface VendorProducts {
  vid: number;
  products: Product[];
  schemaVersion: number;
}

function createBaseVendorProducts(): VendorProducts {
  return { vid: 0, products: [], schemaVersion: 0 };
}

export const VendorProducts = {
  encode(message: VendorProducts, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    for (const v of message.products) {
      Product.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(24).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VendorProducts {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVendorProducts();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.products.push(Product.decode(reader, reader.uint32()));
          break;
        case 3:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VendorProducts {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      products: Array.isArray(object?.products) ? object.products.map((e: any) => Product.fromJSON(e)) : [],
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: VendorProducts): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    if (message.products) {
      obj.products = message.products.map((e) => e ? Product.toJSON(e) : undefined);
    } else {
      obj.products = [];
    }
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VendorProducts>, I>>(object: I): VendorProducts {
    const message = createBaseVendorProducts();
    message.vid = object.vid ?? 0;
    message.products = object.products?.map((e) => Product.fromPartial(e)) || [];
    message.schemaVersion = object.schemaVersion ?? 0;
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
