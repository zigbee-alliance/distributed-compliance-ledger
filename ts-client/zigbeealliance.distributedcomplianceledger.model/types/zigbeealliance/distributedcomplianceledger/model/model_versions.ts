/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";

export interface ModelVersions {
  vid: number;
  pid: number;
  softwareVersions: number[];
  schemaVersion: number;
}

function createBaseModelVersions(): ModelVersions {
  return { vid: 0, pid: 0, softwareVersions: [], schemaVersion: 0 };
}

export const ModelVersions = {
  encode(message: ModelVersions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    writer.uint32(26).fork();
    for (const v of message.softwareVersions) {
      writer.uint32(v);
    }
    writer.ldelim();
    if (message.schemaVersion !== 0) {
      writer.uint32(32).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ModelVersions {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseModelVersions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.pid = reader.int32();
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.softwareVersions.push(reader.uint32());
            }
          } else {
            message.softwareVersions.push(reader.uint32());
          }
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

  fromJSON(object: any): ModelVersions {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersions: Array.isArray(object?.softwareVersions)
        ? object.softwareVersions.map((e: any) => Number(e))
        : [],
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: ModelVersions): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    if (message.softwareVersions) {
      obj.softwareVersions = message.softwareVersions.map((e) => Math.round(e));
    } else {
      obj.softwareVersions = [];
    }
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ModelVersions>, I>>(object: I): ModelVersions {
    const message = createBaseModelVersions();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersions = object.softwareVersions?.map((e) => e) || [];
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
