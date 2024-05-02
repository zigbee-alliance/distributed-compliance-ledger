/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { CertificateIdentifier } from "./certificate_identifier";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface RevokedRootCertificates {
  certs: CertificateIdentifier[];
}

function createBaseRevokedRootCertificates(): RevokedRootCertificates {
  return { certs: [] };
}

export const RevokedRootCertificates = {
  encode(message: RevokedRootCertificates, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.certs) {
      CertificateIdentifier.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RevokedRootCertificates {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRevokedRootCertificates();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.certs.push(CertificateIdentifier.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RevokedRootCertificates {
    return {
      certs: Array.isArray(object?.certs) ? object.certs.map((e: any) => CertificateIdentifier.fromJSON(e)) : [],
    };
  },

  toJSON(message: RevokedRootCertificates): unknown {
    const obj: any = {};
    if (message.certs) {
      obj.certs = message.certs.map((e) => e ? CertificateIdentifier.toJSON(e) : undefined);
    } else {
      obj.certs = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RevokedRootCertificates>, I>>(object: I): RevokedRootCertificates {
    const message = createBaseRevokedRootCertificates();
    message.certs = object.certs?.map((e) => CertificateIdentifier.fromPartial(e)) || [];
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
