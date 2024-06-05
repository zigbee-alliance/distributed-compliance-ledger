/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";

export interface RevokedModel {
  vid: number;
  pid: number;
  softwareVersion: number;
  certificationType: string;
  value: boolean;
  schemaVersion: number;
}

function createBaseRevokedModel(): RevokedModel {
  return { vid: 0, pid: 0, softwareVersion: 0, certificationType: "", value: false, schemaVersion: 0 };
}

export const RevokedModel = {
  encode(message: RevokedModel, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion);
    }
    if (message.certificationType !== "") {
      writer.uint32(34).string(message.certificationType);
    }
    if (message.value === true) {
      writer.uint32(40).bool(message.value);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(48).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RevokedModel {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRevokedModel();
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
          message.softwareVersion = reader.uint32();
          break;
        case 4:
          message.certificationType = reader.string();
          break;
        case 5:
          message.value = reader.bool();
          break;
        case 6:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RevokedModel {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
      certificationType: isSet(object.certificationType) ? String(object.certificationType) : "",
      value: isSet(object.value) ? Boolean(object.value) : false,
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: RevokedModel): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    message.certificationType !== undefined && (obj.certificationType = message.certificationType);
    message.value !== undefined && (obj.value = message.value);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RevokedModel>, I>>(object: I): RevokedModel {
    const message = createBaseRevokedModel();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    message.certificationType = object.certificationType ?? "";
    message.value = object.value ?? false;
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
