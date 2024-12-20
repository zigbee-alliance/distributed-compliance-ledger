/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface NocCertificatesBySubject {
  subject: string;
  subjectKeyIds: string[];
  schemaVersion: number;
}

function createBaseNocCertificatesBySubject(): NocCertificatesBySubject {
  return { subject: "", subjectKeyIds: [], schemaVersion: 0 };
}

export const NocCertificatesBySubject = {
  encode(message: NocCertificatesBySubject, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    for (const v of message.subjectKeyIds) {
      writer.uint32(18).string(v!);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(24).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NocCertificatesBySubject {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNocCertificatesBySubject();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyIds.push(reader.string());
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

  fromJSON(object: any): NocCertificatesBySubject {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyIds: Array.isArray(object?.subjectKeyIds) ? object.subjectKeyIds.map((e: any) => String(e)) : [],
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: NocCertificatesBySubject): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    if (message.subjectKeyIds) {
      obj.subjectKeyIds = message.subjectKeyIds.map((e) => e);
    } else {
      obj.subjectKeyIds = [];
    }
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NocCertificatesBySubject>, I>>(object: I): NocCertificatesBySubject {
    const message = createBaseNocCertificatesBySubject();
    message.subject = object.subject ?? "";
    message.subjectKeyIds = object.subjectKeyIds?.map((e) => e) || [];
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
