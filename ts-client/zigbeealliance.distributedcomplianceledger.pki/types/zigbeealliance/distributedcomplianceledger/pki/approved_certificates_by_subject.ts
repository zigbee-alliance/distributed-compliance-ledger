/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface ApprovedCertificatesBySubject {
  subject: string;
  subjectKeyIds: string[];
}

function createBaseApprovedCertificatesBySubject(): ApprovedCertificatesBySubject {
  return { subject: "", subjectKeyIds: [] };
}

export const ApprovedCertificatesBySubject = {
  encode(message: ApprovedCertificatesBySubject, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    for (const v of message.subjectKeyIds) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ApprovedCertificatesBySubject {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseApprovedCertificatesBySubject();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyIds.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ApprovedCertificatesBySubject {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyIds: Array.isArray(object?.subjectKeyIds) ? object.subjectKeyIds.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: ApprovedCertificatesBySubject): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    if (message.subjectKeyIds) {
      obj.subjectKeyIds = message.subjectKeyIds.map((e) => e);
    } else {
      obj.subjectKeyIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ApprovedCertificatesBySubject>, I>>(
    object: I,
  ): ApprovedCertificatesBySubject {
    const message = createBaseApprovedCertificatesBySubject();
    message.subject = object.subject ?? "";
    message.subjectKeyIds = object.subjectKeyIds?.map((e) => e) || [];
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