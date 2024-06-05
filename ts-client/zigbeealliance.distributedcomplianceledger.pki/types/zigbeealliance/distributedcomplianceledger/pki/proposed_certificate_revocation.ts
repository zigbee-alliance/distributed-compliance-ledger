/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Grant } from "./grant";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface ProposedCertificateRevocation {
  subject: string;
  subjectKeyId: string;
  approvals: Grant[];
  subjectAsText: string;
  serialNumber: string;
  revokeChild: boolean;
  schemaVersion: number;
}

function createBaseProposedCertificateRevocation(): ProposedCertificateRevocation {
  return {
    subject: "",
    subjectKeyId: "",
    approvals: [],
    subjectAsText: "",
    serialNumber: "",
    revokeChild: false,
    schemaVersion: 0,
  };
}

export const ProposedCertificateRevocation = {
  encode(message: ProposedCertificateRevocation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.subjectAsText !== "") {
      writer.uint32(34).string(message.subjectAsText);
    }
    if (message.serialNumber !== "") {
      writer.uint32(42).string(message.serialNumber);
    }
    if (message.revokeChild === true) {
      writer.uint32(48).bool(message.revokeChild);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(56).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProposedCertificateRevocation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposedCertificateRevocation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        case 3:
          message.approvals.push(Grant.decode(reader, reader.uint32()));
          break;
        case 4:
          message.subjectAsText = reader.string();
          break;
        case 5:
          message.serialNumber = reader.string();
          break;
        case 6:
          message.revokeChild = reader.bool();
          break;
        case 7:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ProposedCertificateRevocation {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
      approvals: Array.isArray(object?.approvals) ? object.approvals.map((e: any) => Grant.fromJSON(e)) : [],
      subjectAsText: isSet(object.subjectAsText) ? String(object.subjectAsText) : "",
      serialNumber: isSet(object.serialNumber) ? String(object.serialNumber) : "",
      revokeChild: isSet(object.revokeChild) ? Boolean(object.revokeChild) : false,
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: ProposedCertificateRevocation): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => e ? Grant.toJSON(e) : undefined);
    } else {
      obj.approvals = [];
    }
    message.subjectAsText !== undefined && (obj.subjectAsText = message.subjectAsText);
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber);
    message.revokeChild !== undefined && (obj.revokeChild = message.revokeChild);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ProposedCertificateRevocation>, I>>(
    object: I,
  ): ProposedCertificateRevocation {
    const message = createBaseProposedCertificateRevocation();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    message.approvals = object.approvals?.map((e) => Grant.fromPartial(e)) || [];
    message.subjectAsText = object.subjectAsText ?? "";
    message.serialNumber = object.serialNumber ?? "";
    message.revokeChild = object.revokeChild ?? false;
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
