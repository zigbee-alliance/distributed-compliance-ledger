/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { ComplianceHistoryItem } from "./compliance_history_item";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";

export interface ComplianceInfo {
  vid: number;
  pid: number;
  softwareVersion: number;
  certificationType: string;
  softwareVersionString: string;
  cDVersionNumber: number;
  softwareVersionCertificationStatus: number;
  date: string;
  reason: string;
  owner: string;
  history: ComplianceHistoryItem[];
  cDCertificateId: string;
  certificationRoute: string;
  programType: string;
  programTypeVersion: string;
  compliantPlatformUsed: string;
  compliantPlatformVersion: string;
  transport: string;
  familyId: string;
  supportedClusters: string;
  OSVersion: string;
  parentChild: string;
  certificationIdOfSoftwareComponent: string;
  schemaVersion: number;
}

function createBaseComplianceInfo(): ComplianceInfo {
  return {
    vid: 0,
    pid: 0,
    softwareVersion: 0,
    certificationType: "",
    softwareVersionString: "",
    cDVersionNumber: 0,
    softwareVersionCertificationStatus: 0,
    date: "",
    reason: "",
    owner: "",
    history: [],
    cDCertificateId: "",
    certificationRoute: "",
    programType: "",
    programTypeVersion: "",
    compliantPlatformUsed: "",
    compliantPlatformVersion: "",
    transport: "",
    familyId: "",
    supportedClusters: "",
    OSVersion: "",
    parentChild: "",
    certificationIdOfSoftwareComponent: "",
    schemaVersion: 0,
  };
}

export const ComplianceInfo = {
  encode(message: ComplianceInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.softwareVersionString !== "") {
      writer.uint32(42).string(message.softwareVersionString);
    }
    if (message.cDVersionNumber !== 0) {
      writer.uint32(48).uint32(message.cDVersionNumber);
    }
    if (message.softwareVersionCertificationStatus !== 0) {
      writer.uint32(56).uint32(message.softwareVersionCertificationStatus);
    }
    if (message.date !== "") {
      writer.uint32(66).string(message.date);
    }
    if (message.reason !== "") {
      writer.uint32(74).string(message.reason);
    }
    if (message.owner !== "") {
      writer.uint32(82).string(message.owner);
    }
    for (const v of message.history) {
      ComplianceHistoryItem.encode(v!, writer.uint32(90).fork()).ldelim();
    }
    if (message.cDCertificateId !== "") {
      writer.uint32(98).string(message.cDCertificateId);
    }
    if (message.certificationRoute !== "") {
      writer.uint32(106).string(message.certificationRoute);
    }
    if (message.programType !== "") {
      writer.uint32(114).string(message.programType);
    }
    if (message.programTypeVersion !== "") {
      writer.uint32(122).string(message.programTypeVersion);
    }
    if (message.compliantPlatformUsed !== "") {
      writer.uint32(130).string(message.compliantPlatformUsed);
    }
    if (message.compliantPlatformVersion !== "") {
      writer.uint32(138).string(message.compliantPlatformVersion);
    }
    if (message.transport !== "") {
      writer.uint32(146).string(message.transport);
    }
    if (message.familyId !== "") {
      writer.uint32(154).string(message.familyId);
    }
    if (message.supportedClusters !== "") {
      writer.uint32(162).string(message.supportedClusters);
    }
    if (message.OSVersion !== "") {
      writer.uint32(170).string(message.OSVersion);
    }
    if (message.parentChild !== "") {
      writer.uint32(178).string(message.parentChild);
    }
    if (message.certificationIdOfSoftwareComponent !== "") {
      writer.uint32(186).string(message.certificationIdOfSoftwareComponent);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(192).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ComplianceInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseComplianceInfo();
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
          message.softwareVersionString = reader.string();
          break;
        case 6:
          message.cDVersionNumber = reader.uint32();
          break;
        case 7:
          message.softwareVersionCertificationStatus = reader.uint32();
          break;
        case 8:
          message.date = reader.string();
          break;
        case 9:
          message.reason = reader.string();
          break;
        case 10:
          message.owner = reader.string();
          break;
        case 11:
          message.history.push(ComplianceHistoryItem.decode(reader, reader.uint32()));
          break;
        case 12:
          message.cDCertificateId = reader.string();
          break;
        case 13:
          message.certificationRoute = reader.string();
          break;
        case 14:
          message.programType = reader.string();
          break;
        case 15:
          message.programTypeVersion = reader.string();
          break;
        case 16:
          message.compliantPlatformUsed = reader.string();
          break;
        case 17:
          message.compliantPlatformVersion = reader.string();
          break;
        case 18:
          message.transport = reader.string();
          break;
        case 19:
          message.familyId = reader.string();
          break;
        case 20:
          message.supportedClusters = reader.string();
          break;
        case 21:
          message.OSVersion = reader.string();
          break;
        case 22:
          message.parentChild = reader.string();
          break;
        case 23:
          message.certificationIdOfSoftwareComponent = reader.string();
          break;
        case 24:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ComplianceInfo {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
      certificationType: isSet(object.certificationType) ? String(object.certificationType) : "",
      softwareVersionString: isSet(object.softwareVersionString) ? String(object.softwareVersionString) : "",
      cDVersionNumber: isSet(object.cDVersionNumber) ? Number(object.cDVersionNumber) : 0,
      softwareVersionCertificationStatus: isSet(object.softwareVersionCertificationStatus)
        ? Number(object.softwareVersionCertificationStatus)
        : 0,
      date: isSet(object.date) ? String(object.date) : "",
      reason: isSet(object.reason) ? String(object.reason) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      history: Array.isArray(object?.history) ? object.history.map((e: any) => ComplianceHistoryItem.fromJSON(e)) : [],
      cDCertificateId: isSet(object.cDCertificateId) ? String(object.cDCertificateId) : "",
      certificationRoute: isSet(object.certificationRoute) ? String(object.certificationRoute) : "",
      programType: isSet(object.programType) ? String(object.programType) : "",
      programTypeVersion: isSet(object.programTypeVersion) ? String(object.programTypeVersion) : "",
      compliantPlatformUsed: isSet(object.compliantPlatformUsed) ? String(object.compliantPlatformUsed) : "",
      compliantPlatformVersion: isSet(object.compliantPlatformVersion) ? String(object.compliantPlatformVersion) : "",
      transport: isSet(object.transport) ? String(object.transport) : "",
      familyId: isSet(object.familyId) ? String(object.familyId) : "",
      supportedClusters: isSet(object.supportedClusters) ? String(object.supportedClusters) : "",
      OSVersion: isSet(object.OSVersion) ? String(object.OSVersion) : "",
      parentChild: isSet(object.parentChild) ? String(object.parentChild) : "",
      certificationIdOfSoftwareComponent: isSet(object.certificationIdOfSoftwareComponent)
        ? String(object.certificationIdOfSoftwareComponent)
        : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: ComplianceInfo): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    message.certificationType !== undefined && (obj.certificationType = message.certificationType);
    message.softwareVersionString !== undefined && (obj.softwareVersionString = message.softwareVersionString);
    message.cDVersionNumber !== undefined && (obj.cDVersionNumber = Math.round(message.cDVersionNumber));
    message.softwareVersionCertificationStatus !== undefined
      && (obj.softwareVersionCertificationStatus = Math.round(message.softwareVersionCertificationStatus));
    message.date !== undefined && (obj.date = message.date);
    message.reason !== undefined && (obj.reason = message.reason);
    message.owner !== undefined && (obj.owner = message.owner);
    if (message.history) {
      obj.history = message.history.map((e) => e ? ComplianceHistoryItem.toJSON(e) : undefined);
    } else {
      obj.history = [];
    }
    message.cDCertificateId !== undefined && (obj.cDCertificateId = message.cDCertificateId);
    message.certificationRoute !== undefined && (obj.certificationRoute = message.certificationRoute);
    message.programType !== undefined && (obj.programType = message.programType);
    message.programTypeVersion !== undefined && (obj.programTypeVersion = message.programTypeVersion);
    message.compliantPlatformUsed !== undefined && (obj.compliantPlatformUsed = message.compliantPlatformUsed);
    message.compliantPlatformVersion !== undefined && (obj.compliantPlatformVersion = message.compliantPlatformVersion);
    message.transport !== undefined && (obj.transport = message.transport);
    message.familyId !== undefined && (obj.familyId = message.familyId);
    message.supportedClusters !== undefined && (obj.supportedClusters = message.supportedClusters);
    message.OSVersion !== undefined && (obj.OSVersion = message.OSVersion);
    message.parentChild !== undefined && (obj.parentChild = message.parentChild);
    message.certificationIdOfSoftwareComponent !== undefined
      && (obj.certificationIdOfSoftwareComponent = message.certificationIdOfSoftwareComponent);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ComplianceInfo>, I>>(object: I): ComplianceInfo {
    const message = createBaseComplianceInfo();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    message.certificationType = object.certificationType ?? "";
    message.softwareVersionString = object.softwareVersionString ?? "";
    message.cDVersionNumber = object.cDVersionNumber ?? 0;
    message.softwareVersionCertificationStatus = object.softwareVersionCertificationStatus ?? 0;
    message.date = object.date ?? "";
    message.reason = object.reason ?? "";
    message.owner = object.owner ?? "";
    message.history = object.history?.map((e) => ComplianceHistoryItem.fromPartial(e)) || [];
    message.cDCertificateId = object.cDCertificateId ?? "";
    message.certificationRoute = object.certificationRoute ?? "";
    message.programType = object.programType ?? "";
    message.programTypeVersion = object.programTypeVersion ?? "";
    message.compliantPlatformUsed = object.compliantPlatformUsed ?? "";
    message.compliantPlatformVersion = object.compliantPlatformVersion ?? "";
    message.transport = object.transport ?? "";
    message.familyId = object.familyId ?? "";
    message.supportedClusters = object.supportedClusters ?? "";
    message.OSVersion = object.OSVersion ?? "";
    message.parentChild = object.parentChild ?? "";
    message.certificationIdOfSoftwareComponent = object.certificationIdOfSoftwareComponent ?? "";
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
