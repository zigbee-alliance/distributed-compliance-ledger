/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { CertificateIdentifier } from "./certificate_identifier";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface ChildCertificates {
  issuer: string;
  authorityKeyId: string;
  certIds: CertificateIdentifier[];
  schemaVersion: number;
}

function createBaseChildCertificates(): ChildCertificates {
  return { issuer: "", authorityKeyId: "", certIds: [], schemaVersion: 0 };
}

export const ChildCertificates = {
  encode(message: ChildCertificates, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    if (message.authorityKeyId !== "") {
      writer.uint32(18).string(message.authorityKeyId);
    }
    for (const v of message.certIds) {
      CertificateIdentifier.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(32).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChildCertificates {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChildCertificates();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string();
          break;
        case 2:
          message.authorityKeyId = reader.string();
          break;
        case 3:
          message.certIds.push(CertificateIdentifier.decode(reader, reader.uint32()));
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

  fromJSON(object: any): ChildCertificates {
    return {
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      authorityKeyId: isSet(object.authorityKeyId) ? String(object.authorityKeyId) : "",
      certIds: Array.isArray(object?.certIds) ? object.certIds.map((e: any) => CertificateIdentifier.fromJSON(e)) : [],
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: ChildCertificates): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId);
    if (message.certIds) {
      obj.certIds = message.certIds.map((e) => e ? CertificateIdentifier.toJSON(e) : undefined);
    } else {
      obj.certIds = [];
    }
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChildCertificates>, I>>(object: I): ChildCertificates {
    const message = createBaseChildCertificates();
    message.issuer = object.issuer ?? "";
    message.authorityKeyId = object.authorityKeyId ?? "";
    message.certIds = object.certIds?.map((e) => CertificateIdentifier.fromPartial(e)) || [];
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
