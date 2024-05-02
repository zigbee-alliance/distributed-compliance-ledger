/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.common";

export interface Uint16Range {
  min: number;
  max: number;
}

function createBaseUint16Range(): Uint16Range {
  return { min: 0, max: 0 };
}

export const Uint16Range = {
  encode(message: Uint16Range, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.min !== 0) {
      writer.uint32(8).int32(message.min);
    }
    if (message.max !== 0) {
      writer.uint32(16).int32(message.max);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Uint16Range {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUint16Range();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.min = reader.int32();
          break;
        case 2:
          message.max = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Uint16Range {
    return { min: isSet(object.min) ? Number(object.min) : 0, max: isSet(object.max) ? Number(object.max) : 0 };
  },

  toJSON(message: Uint16Range): unknown {
    const obj: any = {};
    message.min !== undefined && (obj.min = Math.round(message.min));
    message.max !== undefined && (obj.max = Math.round(message.max));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Uint16Range>, I>>(object: I): Uint16Range {
    const message = createBaseUint16Range();
    message.min = object.min ?? 0;
    message.max = object.max ?? 0;
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
