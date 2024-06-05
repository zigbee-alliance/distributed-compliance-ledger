/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";

export interface ComplianceHistoryItem {
  softwareVersionCertificationStatus: number;
  date: string;
  reason: string;
  cDVersionNumber: number;
  schemaVersion: number;
}

function createBaseComplianceHistoryItem(): ComplianceHistoryItem {
  return { softwareVersionCertificationStatus: 0, date: "", reason: "", cDVersionNumber: 0, schemaVersion: 0 };
}

export const ComplianceHistoryItem = {
  encode(message: ComplianceHistoryItem, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.softwareVersionCertificationStatus !== 0) {
      writer.uint32(8).uint32(message.softwareVersionCertificationStatus);
    }
    if (message.date !== "") {
      writer.uint32(18).string(message.date);
    }
    if (message.reason !== "") {
      writer.uint32(26).string(message.reason);
    }
    if (message.cDVersionNumber !== 0) {
      writer.uint32(32).uint32(message.cDVersionNumber);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(40).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ComplianceHistoryItem {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseComplianceHistoryItem();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.softwareVersionCertificationStatus = reader.uint32();
          break;
        case 2:
          message.date = reader.string();
          break;
        case 3:
          message.reason = reader.string();
          break;
        case 4:
          message.cDVersionNumber = reader.uint32();
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

  fromJSON(object: any): ComplianceHistoryItem {
    return {
      softwareVersionCertificationStatus: isSet(object.softwareVersionCertificationStatus)
        ? Number(object.softwareVersionCertificationStatus)
        : 0,
      date: isSet(object.date) ? String(object.date) : "",
      reason: isSet(object.reason) ? String(object.reason) : "",
      cDVersionNumber: isSet(object.cDVersionNumber) ? Number(object.cDVersionNumber) : 0,
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: ComplianceHistoryItem): unknown {
    const obj: any = {};
    message.softwareVersionCertificationStatus !== undefined
      && (obj.softwareVersionCertificationStatus = Math.round(message.softwareVersionCertificationStatus));
    message.date !== undefined && (obj.date = message.date);
    message.reason !== undefined && (obj.reason = message.reason);
    message.cDVersionNumber !== undefined && (obj.cDVersionNumber = Math.round(message.cDVersionNumber));
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ComplianceHistoryItem>, I>>(object: I): ComplianceHistoryItem {
    const message = createBaseComplianceHistoryItem();
    message.softwareVersionCertificationStatus = object.softwareVersionCertificationStatus ?? 0;
    message.date = object.date ?? "";
    message.reason = object.reason ?? "";
    message.cDVersionNumber = object.cDVersionNumber ?? 0;
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
