/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";

export interface Product {
  pid: number;
  name: string;
  partNumber: string;
  schemaVersion: number;
}

function createBaseProduct(): Product {
  return { pid: 0, name: "", partNumber: "", schemaVersion: 0 };
}

export const Product = {
  encode(message: Product, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pid !== 0) {
      writer.uint32(8).int32(message.pid);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.partNumber !== "") {
      writer.uint32(26).string(message.partNumber);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(32).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Product {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProduct();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pid = reader.int32();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.partNumber = reader.string();
          break;
        case 4:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Product {
    return {
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      partNumber: isSet(object.partNumber) ? String(object.partNumber) : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: Product): unknown {
    const obj: any = {};
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.name !== undefined && (obj.name = message.name);
    message.partNumber !== undefined && (obj.partNumber = message.partNumber);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Product>, I>>(object: I): Product {
    const message = createBaseProduct();
    message.pid = object.pid ?? 0;
    message.name = object.name ?? "";
    message.partNumber = object.partNumber ?? "";
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
