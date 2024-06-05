/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Grant } from "./grant";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface ProposedCertificate {
  subject: string;
  subjectKeyId: string;
  pemCert: string;
  serialNumber: string;
  owner: string;
  approvals: Grant[];
  subjectAsText: string;
  rejects: Grant[];
  vid: number;
  certSchemaVersion: number;
  schemaVersion: number;
}

function createBaseProposedCertificate(): ProposedCertificate {
  return {
    subject: "",
    subjectKeyId: "",
    pemCert: "",
    serialNumber: "",
    owner: "",
    approvals: [],
    subjectAsText: "",
    rejects: [],
    vid: 0,
    certSchemaVersion: 0,
    schemaVersion: 0,
  };
}

export const ProposedCertificate = {
  encode(message: ProposedCertificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    if (message.pemCert !== "") {
      writer.uint32(26).string(message.pemCert);
    }
    if (message.serialNumber !== "") {
      writer.uint32(34).string(message.serialNumber);
    }
    if (message.owner !== "") {
      writer.uint32(42).string(message.owner);
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.subjectAsText !== "") {
      writer.uint32(58).string(message.subjectAsText);
    }
    for (const v of message.rejects) {
      Grant.encode(v!, writer.uint32(66).fork()).ldelim();
    }
    if (message.vid !== 0) {
      writer.uint32(72).int32(message.vid);
    }
    if (message.certSchemaVersion !== 0) {
      writer.uint32(80).uint32(message.certSchemaVersion);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(88).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProposedCertificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposedCertificate();
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
          message.pemCert = reader.string();
          break;
        case 4:
          message.serialNumber = reader.string();
          break;
        case 5:
          message.owner = reader.string();
          break;
        case 6:
          message.approvals.push(Grant.decode(reader, reader.uint32()));
          break;
        case 7:
          message.subjectAsText = reader.string();
          break;
        case 8:
          message.rejects.push(Grant.decode(reader, reader.uint32()));
          break;
        case 9:
          message.vid = reader.int32();
          break;
        case 10:
          message.certSchemaVersion = reader.uint32();
          break;
        case 11:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ProposedCertificate {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
      pemCert: isSet(object.pemCert) ? String(object.pemCert) : "",
      serialNumber: isSet(object.serialNumber) ? String(object.serialNumber) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      approvals: Array.isArray(object?.approvals) ? object.approvals.map((e: any) => Grant.fromJSON(e)) : [],
      subjectAsText: isSet(object.subjectAsText) ? String(object.subjectAsText) : "",
      rejects: Array.isArray(object?.rejects) ? object.rejects.map((e: any) => Grant.fromJSON(e)) : [],
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      certSchemaVersion: isSet(object.certSchemaVersion) ? Number(object.certSchemaVersion) : 0,
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: ProposedCertificate): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    message.pemCert !== undefined && (obj.pemCert = message.pemCert);
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber);
    message.owner !== undefined && (obj.owner = message.owner);
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => e ? Grant.toJSON(e) : undefined);
    } else {
      obj.approvals = [];
    }
    message.subjectAsText !== undefined && (obj.subjectAsText = message.subjectAsText);
    if (message.rejects) {
      obj.rejects = message.rejects.map((e) => e ? Grant.toJSON(e) : undefined);
    } else {
      obj.rejects = [];
    }
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.certSchemaVersion !== undefined && (obj.certSchemaVersion = Math.round(message.certSchemaVersion));
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ProposedCertificate>, I>>(object: I): ProposedCertificate {
    const message = createBaseProposedCertificate();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    message.pemCert = object.pemCert ?? "";
    message.serialNumber = object.serialNumber ?? "";
    message.owner = object.owner ?? "";
    message.approvals = object.approvals?.map((e) => Grant.fromPartial(e)) || [];
    message.subjectAsText = object.subjectAsText ?? "";
    message.rejects = object.rejects?.map((e) => Grant.fromPartial(e)) || [];
    message.vid = object.vid ?? 0;
    message.certSchemaVersion = object.certSchemaVersion ?? 0;
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
