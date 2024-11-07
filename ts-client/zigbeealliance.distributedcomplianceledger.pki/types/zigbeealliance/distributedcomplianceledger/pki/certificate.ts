/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Grant } from "./grant";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export enum CertificateType {
  DeviceAttestationPKI = 0,
  OperationalPKI = 1,
  VIDSignerPKI = 2,
  UNRECOGNIZED = -1,
}

export function certificateTypeFromJSON(object: any): CertificateType {
  switch (object) {
    case 0:
    case "DeviceAttestationPKI":
      return CertificateType.DeviceAttestationPKI;
    case 1:
    case "OperationalPKI":
      return CertificateType.OperationalPKI;
    case 2:
    case "VIDSignerPKI":
      return CertificateType.VIDSignerPKI;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CertificateType.UNRECOGNIZED;
  }
}

export function certificateTypeToJSON(object: CertificateType): string {
  switch (object) {
    case CertificateType.DeviceAttestationPKI:
      return "DeviceAttestationPKI";
    case CertificateType.OperationalPKI:
      return "OperationalPKI";
    case CertificateType.VIDSignerPKI:
      return "VIDSignerPKI";
    case CertificateType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Certificate {
  pemCert: string;
  serialNumber: string;
  issuer: string;
  authorityKeyId: string;
  rootSubject: string;
  rootSubjectKeyId: string;
  isRoot: boolean;
  owner: string;
  subject: string;
  subjectKeyId: string;
  approvals: Grant[];
  subjectAsText: string;
  rejects: Grant[];
  vid: number;
  certificateType: CertificateType;
  schemaVersion: number;
}

function createBaseCertificate(): Certificate {
  return {
    pemCert: "",
    serialNumber: "",
    issuer: "",
    authorityKeyId: "",
    rootSubject: "",
    rootSubjectKeyId: "",
    isRoot: false,
    owner: "",
    subject: "",
    subjectKeyId: "",
    approvals: [],
    subjectAsText: "",
    rejects: [],
    vid: 0,
    certificateType: 0,
    schemaVersion: 0,
  };
}

export const Certificate = {
  encode(message: Certificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pemCert !== "") {
      writer.uint32(10).string(message.pemCert);
    }
    if (message.serialNumber !== "") {
      writer.uint32(18).string(message.serialNumber);
    }
    if (message.issuer !== "") {
      writer.uint32(26).string(message.issuer);
    }
    if (message.authorityKeyId !== "") {
      writer.uint32(34).string(message.authorityKeyId);
    }
    if (message.rootSubject !== "") {
      writer.uint32(42).string(message.rootSubject);
    }
    if (message.rootSubjectKeyId !== "") {
      writer.uint32(50).string(message.rootSubjectKeyId);
    }
    if (message.isRoot === true) {
      writer.uint32(56).bool(message.isRoot);
    }
    if (message.owner !== "") {
      writer.uint32(66).string(message.owner);
    }
    if (message.subject !== "") {
      writer.uint32(74).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(82).string(message.subjectKeyId);
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(90).fork()).ldelim();
    }
    if (message.subjectAsText !== "") {
      writer.uint32(98).string(message.subjectAsText);
    }
    for (const v of message.rejects) {
      Grant.encode(v!, writer.uint32(106).fork()).ldelim();
    }
    if (message.vid !== 0) {
      writer.uint32(112).int32(message.vid);
    }
    if (message.certificateType !== 0) {
      writer.uint32(120).int32(message.certificateType);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(128).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Certificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pemCert = reader.string();
          break;
        case 2:
          message.serialNumber = reader.string();
          break;
        case 3:
          message.issuer = reader.string();
          break;
        case 4:
          message.authorityKeyId = reader.string();
          break;
        case 5:
          message.rootSubject = reader.string();
          break;
        case 6:
          message.rootSubjectKeyId = reader.string();
          break;
        case 7:
          message.isRoot = reader.bool();
          break;
        case 8:
          message.owner = reader.string();
          break;
        case 9:
          message.subject = reader.string();
          break;
        case 10:
          message.subjectKeyId = reader.string();
          break;
        case 11:
          message.approvals.push(Grant.decode(reader, reader.uint32()));
          break;
        case 12:
          message.subjectAsText = reader.string();
          break;
        case 13:
          message.rejects.push(Grant.decode(reader, reader.uint32()));
          break;
        case 14:
          message.vid = reader.int32();
          break;
        case 15:
          message.certificateType = reader.int32() as any;
          break;
        case 16:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Certificate {
    return {
      pemCert: isSet(object.pemCert) ? String(object.pemCert) : "",
      serialNumber: isSet(object.serialNumber) ? String(object.serialNumber) : "",
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      authorityKeyId: isSet(object.authorityKeyId) ? String(object.authorityKeyId) : "",
      rootSubject: isSet(object.rootSubject) ? String(object.rootSubject) : "",
      rootSubjectKeyId: isSet(object.rootSubjectKeyId) ? String(object.rootSubjectKeyId) : "",
      isRoot: isSet(object.isRoot) ? Boolean(object.isRoot) : false,
      owner: isSet(object.owner) ? String(object.owner) : "",
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
      approvals: Array.isArray(object?.approvals) ? object.approvals.map((e: any) => Grant.fromJSON(e)) : [],
      subjectAsText: isSet(object.subjectAsText) ? String(object.subjectAsText) : "",
      rejects: Array.isArray(object?.rejects) ? object.rejects.map((e: any) => Grant.fromJSON(e)) : [],
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      certificateType: isSet(object.certificateType) ? certificateTypeFromJSON(object.certificateType) : 0,
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: Certificate): unknown {
    const obj: any = {};
    message.pemCert !== undefined && (obj.pemCert = message.pemCert);
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber);
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId);
    message.rootSubject !== undefined && (obj.rootSubject = message.rootSubject);
    message.rootSubjectKeyId !== undefined && (obj.rootSubjectKeyId = message.rootSubjectKeyId);
    message.isRoot !== undefined && (obj.isRoot = message.isRoot);
    message.owner !== undefined && (obj.owner = message.owner);
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
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
    message.certificateType !== undefined && (obj.certificateType = certificateTypeToJSON(message.certificateType));
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Certificate>, I>>(object: I): Certificate {
    const message = createBaseCertificate();
    message.pemCert = object.pemCert ?? "";
    message.serialNumber = object.serialNumber ?? "";
    message.issuer = object.issuer ?? "";
    message.authorityKeyId = object.authorityKeyId ?? "";
    message.rootSubject = object.rootSubject ?? "";
    message.rootSubjectKeyId = object.rootSubjectKeyId ?? "";
    message.isRoot = object.isRoot ?? false;
    message.owner = object.owner ?? "";
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    message.approvals = object.approvals?.map((e) => Grant.fromPartial(e)) || [];
    message.subjectAsText = object.subjectAsText ?? "";
    message.rejects = object.rejects?.map((e) => Grant.fromPartial(e)) || [];
    message.vid = object.vid ?? 0;
    message.certificateType = object.certificateType ?? 0;
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
