/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Certificate } from "./certificate";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface NocRootCertificatesByVidAndSkid {
  vid: number;
  subjectKeyId: string;
  certs: Certificate[];
  tq: number;
  schemaVersion: number;
}

function createBaseNocRootCertificatesByVidAndSkid(): NocRootCertificatesByVidAndSkid {
  return { vid: 0, subjectKeyId: "", certs: [], tq: 0, schemaVersion: 0 };
}

export const NocRootCertificatesByVidAndSkid = {
  encode(message: NocRootCertificatesByVidAndSkid, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    for (const v of message.certs) {
      Certificate.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.tq !== 0) {
      writer.uint32(37).float(message.tq);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(40).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NocRootCertificatesByVidAndSkid {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNocRootCertificatesByVidAndSkid();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        case 3:
          message.certs.push(Certificate.decode(reader, reader.uint32()));
          break;
        case 4:
          message.tq = reader.float();
          break;
        case 5:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): NocRootCertificatesByVidAndSkid {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
      certs: Array.isArray(object?.certs) ? object.certs.map((e: any) => Certificate.fromJSON(e)) : [],
      tq: isSet(object.tq) ? Number(object.tq) : 0,
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: NocRootCertificatesByVidAndSkid): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    if (message.certs) {
      obj.certs = message.certs.map((e) => e ? Certificate.toJSON(e) : undefined);
    } else {
      obj.certs = [];
    }
    message.tq !== undefined && (obj.tq = message.tq);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NocRootCertificatesByVidAndSkid>, I>>(
    object: I,
  ): NocRootCertificatesByVidAndSkid {
    const message = createBaseNocRootCertificatesByVidAndSkid();
    message.vid = object.vid ?? 0;
    message.subjectKeyId = object.subjectKeyId ?? "";
    message.certs = object.certs?.map((e) => Certificate.fromPartial(e)) || [];
    message.tq = object.tq ?? 0;
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
