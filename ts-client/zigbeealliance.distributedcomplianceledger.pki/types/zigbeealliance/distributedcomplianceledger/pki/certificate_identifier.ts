/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface CertificateIdentifier {
  subject: string;
  subjectKeyId: string;
}

function createBaseCertificateIdentifier(): CertificateIdentifier {
  return { subject: "", subjectKeyId: "" };
}

export const CertificateIdentifier = {
  encode(message: CertificateIdentifier, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CertificateIdentifier {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCertificateIdentifier();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CertificateIdentifier {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: CertificateIdentifier): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CertificateIdentifier>, I>>(object: I): CertificateIdentifier {
    const message = createBaseCertificateIdentifier();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
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
