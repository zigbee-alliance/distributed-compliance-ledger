/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";

export interface ModelVersion {
  vid: number;
  pid: number;
  softwareVersion: number;
  softwareVersionString: string;
  cdVersionNumber: number;
  firmwareInformation: string;
  softwareVersionValid: boolean;
  otaUrl: string;
  otaFileSize: number;
  otaChecksum: string;
  otaChecksumType: number;
  minApplicableSoftwareVersion: number;
  maxApplicableSoftwareVersion: number;
  releaseNotesUrl: string;
  creator: string;
  schemaVersion: number;
  specificationVersion: number;
}

function createBaseModelVersion(): ModelVersion {
  return {
    vid: 0,
    pid: 0,
    softwareVersion: 0,
    softwareVersionString: "",
    cdVersionNumber: 0,
    firmwareInformation: "",
    softwareVersionValid: false,
    otaUrl: "",
    otaFileSize: 0,
    otaChecksum: "",
    otaChecksumType: 0,
    minApplicableSoftwareVersion: 0,
    maxApplicableSoftwareVersion: 0,
    releaseNotesUrl: "",
    creator: "",
    schemaVersion: 0,
    specificationVersion: 0,
  };
}

export const ModelVersion = {
  encode(message: ModelVersion, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion);
    }
    if (message.softwareVersionString !== "") {
      writer.uint32(34).string(message.softwareVersionString);
    }
    if (message.cdVersionNumber !== 0) {
      writer.uint32(40).int32(message.cdVersionNumber);
    }
    if (message.firmwareInformation !== "") {
      writer.uint32(50).string(message.firmwareInformation);
    }
    if (message.softwareVersionValid === true) {
      writer.uint32(56).bool(message.softwareVersionValid);
    }
    if (message.otaUrl !== "") {
      writer.uint32(66).string(message.otaUrl);
    }
    if (message.otaFileSize !== 0) {
      writer.uint32(72).uint64(message.otaFileSize);
    }
    if (message.otaChecksum !== "") {
      writer.uint32(82).string(message.otaChecksum);
    }
    if (message.otaChecksumType !== 0) {
      writer.uint32(88).int32(message.otaChecksumType);
    }
    if (message.minApplicableSoftwareVersion !== 0) {
      writer.uint32(96).uint32(message.minApplicableSoftwareVersion);
    }
    if (message.maxApplicableSoftwareVersion !== 0) {
      writer.uint32(104).uint32(message.maxApplicableSoftwareVersion);
    }
    if (message.releaseNotesUrl !== "") {
      writer.uint32(114).string(message.releaseNotesUrl);
    }
    if (message.creator !== "") {
      writer.uint32(122).string(message.creator);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(128).uint32(message.schemaVersion);
    }
    if (message.specificationVersion !== 0) {
      writer.uint32(136).uint32(message.specificationVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ModelVersion {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseModelVersion();
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
          message.softwareVersionString = reader.string();
          break;
        case 5:
          message.cdVersionNumber = reader.int32();
          break;
        case 6:
          message.firmwareInformation = reader.string();
          break;
        case 7:
          message.softwareVersionValid = reader.bool();
          break;
        case 8:
          message.otaUrl = reader.string();
          break;
        case 9:
          message.otaFileSize = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.otaChecksum = reader.string();
          break;
        case 11:
          message.otaChecksumType = reader.int32();
          break;
        case 12:
          message.minApplicableSoftwareVersion = reader.uint32();
          break;
        case 13:
          message.maxApplicableSoftwareVersion = reader.uint32();
          break;
        case 14:
          message.releaseNotesUrl = reader.string();
          break;
        case 15:
          message.creator = reader.string();
          break;
        case 16:
          message.schemaVersion = reader.uint32();
          break;
        case 17:
          message.specificationVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ModelVersion {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
      softwareVersionString: isSet(object.softwareVersionString) ? String(object.softwareVersionString) : "",
      cdVersionNumber: isSet(object.cdVersionNumber) ? Number(object.cdVersionNumber) : 0,
      firmwareInformation: isSet(object.firmwareInformation) ? String(object.firmwareInformation) : "",
      softwareVersionValid: isSet(object.softwareVersionValid) ? Boolean(object.softwareVersionValid) : false,
      otaUrl: isSet(object.otaUrl) ? String(object.otaUrl) : "",
      otaFileSize: isSet(object.otaFileSize) ? Number(object.otaFileSize) : 0,
      otaChecksum: isSet(object.otaChecksum) ? String(object.otaChecksum) : "",
      otaChecksumType: isSet(object.otaChecksumType) ? Number(object.otaChecksumType) : 0,
      minApplicableSoftwareVersion: isSet(object.minApplicableSoftwareVersion)
        ? Number(object.minApplicableSoftwareVersion)
        : 0,
      maxApplicableSoftwareVersion: isSet(object.maxApplicableSoftwareVersion)
        ? Number(object.maxApplicableSoftwareVersion)
        : 0,
      releaseNotesUrl: isSet(object.releaseNotesUrl) ? String(object.releaseNotesUrl) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
      specificationVersion: isSet(object.specificationVersion) ? Number(object.specificationVersion) : 0,
    };
  },

  toJSON(message: ModelVersion): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    message.softwareVersionString !== undefined && (obj.softwareVersionString = message.softwareVersionString);
    message.cdVersionNumber !== undefined && (obj.cdVersionNumber = Math.round(message.cdVersionNumber));
    message.firmwareInformation !== undefined && (obj.firmwareInformation = message.firmwareInformation);
    message.softwareVersionValid !== undefined && (obj.softwareVersionValid = message.softwareVersionValid);
    message.otaUrl !== undefined && (obj.otaUrl = message.otaUrl);
    message.otaFileSize !== undefined && (obj.otaFileSize = Math.round(message.otaFileSize));
    message.otaChecksum !== undefined && (obj.otaChecksum = message.otaChecksum);
    message.otaChecksumType !== undefined && (obj.otaChecksumType = Math.round(message.otaChecksumType));
    message.minApplicableSoftwareVersion !== undefined
      && (obj.minApplicableSoftwareVersion = Math.round(message.minApplicableSoftwareVersion));
    message.maxApplicableSoftwareVersion !== undefined
      && (obj.maxApplicableSoftwareVersion = Math.round(message.maxApplicableSoftwareVersion));
    message.releaseNotesUrl !== undefined && (obj.releaseNotesUrl = message.releaseNotesUrl);
    message.creator !== undefined && (obj.creator = message.creator);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    message.specificationVersion !== undefined && (obj.specificationVersion = Math.round(message.specificationVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ModelVersion>, I>>(object: I): ModelVersion {
    const message = createBaseModelVersion();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    message.softwareVersionString = object.softwareVersionString ?? "";
    message.cdVersionNumber = object.cdVersionNumber ?? 0;
    message.firmwareInformation = object.firmwareInformation ?? "";
    message.softwareVersionValid = object.softwareVersionValid ?? false;
    message.otaUrl = object.otaUrl ?? "";
    message.otaFileSize = object.otaFileSize ?? 0;
    message.otaChecksum = object.otaChecksum ?? "";
    message.otaChecksumType = object.otaChecksumType ?? 0;
    message.minApplicableSoftwareVersion = object.minApplicableSoftwareVersion ?? 0;
    message.maxApplicableSoftwareVersion = object.maxApplicableSoftwareVersion ?? 0;
    message.releaseNotesUrl = object.releaseNotesUrl ?? "";
    message.creator = object.creator ?? "";
    message.schemaVersion = object.schemaVersion ?? 0;
    message.specificationVersion = object.specificationVersion ?? 0;
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
