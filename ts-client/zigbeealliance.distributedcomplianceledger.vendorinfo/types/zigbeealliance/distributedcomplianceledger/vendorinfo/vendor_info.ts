/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";

export interface VendorInfo {
  vendorID: number;
  vendorName: string;
  companyLegalName: string;
  companyPreferredName: string;
  vendorLandingPageURL: string;
  creator: string;
  schemaVersion: number;
}

function createBaseVendorInfo(): VendorInfo {
  return {
    vendorID: 0,
    vendorName: "",
    companyLegalName: "",
    companyPreferredName: "",
    vendorLandingPageURL: "",
    creator: "",
    schemaVersion: 0,
  };
}

export const VendorInfo = {
  encode(message: VendorInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vendorID !== 0) {
      writer.uint32(8).int32(message.vendorID);
    }
    if (message.vendorName !== "") {
      writer.uint32(18).string(message.vendorName);
    }
    if (message.companyLegalName !== "") {
      writer.uint32(26).string(message.companyLegalName);
    }
    if (message.companyPreferredName !== "") {
      writer.uint32(34).string(message.companyPreferredName);
    }
    if (message.vendorLandingPageURL !== "") {
      writer.uint32(42).string(message.vendorLandingPageURL);
    }
    if (message.creator !== "") {
      writer.uint32(50).string(message.creator);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(56).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VendorInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVendorInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vendorID = reader.int32();
          break;
        case 2:
          message.vendorName = reader.string();
          break;
        case 3:
          message.companyLegalName = reader.string();
          break;
        case 4:
          message.companyPreferredName = reader.string();
          break;
        case 5:
          message.vendorLandingPageURL = reader.string();
          break;
        case 6:
          message.creator = reader.string();
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

  fromJSON(object: any): VendorInfo {
    return {
      vendorID: isSet(object.vendorID) ? Number(object.vendorID) : 0,
      vendorName: isSet(object.vendorName) ? String(object.vendorName) : "",
      companyLegalName: isSet(object.companyLegalName) ? String(object.companyLegalName) : "",
      companyPreferredName: isSet(object.companyPreferredName) ? String(object.companyPreferredName) : "",
      vendorLandingPageURL: isSet(object.vendorLandingPageURL) ? String(object.vendorLandingPageURL) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: VendorInfo): unknown {
    const obj: any = {};
    message.vendorID !== undefined && (obj.vendorID = Math.round(message.vendorID));
    message.vendorName !== undefined && (obj.vendorName = message.vendorName);
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName);
    message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName);
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
    message.creator !== undefined && (obj.creator = message.creator);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VendorInfo>, I>>(object: I): VendorInfo {
    const message = createBaseVendorInfo();
    message.vendorID = object.vendorID ?? 0;
    message.vendorName = object.vendorName ?? "";
    message.companyLegalName = object.companyLegalName ?? "";
    message.companyPreferredName = object.companyPreferredName ?? "";
    message.vendorLandingPageURL = object.vendorLandingPageURL ?? "";
    message.creator = object.creator ?? "";
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
