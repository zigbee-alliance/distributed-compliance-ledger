/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.vendorinfo";

export interface MsgCreateVendorInfo {
  creator: string;
  vendorID: number;
  vendorName: string;
  companyLegalName: string;
  companyPreferredName: string;
  vendorLandingPageURL: string;
  schemaVersion: number;
}

export interface MsgCreateVendorInfoResponse {
}

export interface MsgUpdateVendorInfo {
  creator: string;
  vendorID: number;
  vendorName: string;
  companyLegalName: string;
  companyPreferredName: string;
  vendorLandingPageURL: string;
  schemaVersion: number;
}

export interface MsgUpdateVendorInfoResponse {
}

function createBaseMsgCreateVendorInfo(): MsgCreateVendorInfo {
  return {
    creator: "",
    vendorID: 0,
    vendorName: "",
    companyLegalName: "",
    companyPreferredName: "",
    vendorLandingPageURL: "",
    schemaVersion: 0,
  };
}

export const MsgCreateVendorInfo = {
  encode(message: MsgCreateVendorInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).int32(message.vendorID);
    }
    if (message.vendorName !== "") {
      writer.uint32(26).string(message.vendorName);
    }
    if (message.companyLegalName !== "") {
      writer.uint32(34).string(message.companyLegalName);
    }
    if (message.companyPreferredName !== "") {
      writer.uint32(42).string(message.companyPreferredName);
    }
    if (message.vendorLandingPageURL !== "") {
      writer.uint32(50).string(message.vendorLandingPageURL);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(56).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVendorInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVendorInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.vendorID = reader.int32();
          break;
        case 3:
          message.vendorName = reader.string();
          break;
        case 4:
          message.companyLegalName = reader.string();
          break;
        case 5:
          message.companyPreferredName = reader.string();
          break;
        case 6:
          message.vendorLandingPageURL = reader.string();
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

  fromJSON(object: any): MsgCreateVendorInfo {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      vendorID: isSet(object.vendorID) ? Number(object.vendorID) : 0,
      vendorName: isSet(object.vendorName) ? String(object.vendorName) : "",
      companyLegalName: isSet(object.companyLegalName) ? String(object.companyLegalName) : "",
      companyPreferredName: isSet(object.companyPreferredName) ? String(object.companyPreferredName) : "",
      vendorLandingPageURL: isSet(object.vendorLandingPageURL) ? String(object.vendorLandingPageURL) : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: MsgCreateVendorInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.vendorID !== undefined && (obj.vendorID = Math.round(message.vendorID));
    message.vendorName !== undefined && (obj.vendorName = message.vendorName);
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName);
    message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName);
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVendorInfo>, I>>(object: I): MsgCreateVendorInfo {
    const message = createBaseMsgCreateVendorInfo();
    message.creator = object.creator ?? "";
    message.vendorID = object.vendorID ?? 0;
    message.vendorName = object.vendorName ?? "";
    message.companyLegalName = object.companyLegalName ?? "";
    message.companyPreferredName = object.companyPreferredName ?? "";
    message.vendorLandingPageURL = object.vendorLandingPageURL ?? "";
    message.schemaVersion = object.schemaVersion ?? 0;
    return message;
  },
};

function createBaseMsgCreateVendorInfoResponse(): MsgCreateVendorInfoResponse {
  return {};
}

export const MsgCreateVendorInfoResponse = {
  encode(_: MsgCreateVendorInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVendorInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVendorInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateVendorInfoResponse {
    return {};
  },

  toJSON(_: MsgCreateVendorInfoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVendorInfoResponse>, I>>(_: I): MsgCreateVendorInfoResponse {
    const message = createBaseMsgCreateVendorInfoResponse();
    return message;
  },
};

function createBaseMsgUpdateVendorInfo(): MsgUpdateVendorInfo {
  return {
    creator: "",
    vendorID: 0,
    vendorName: "",
    companyLegalName: "",
    companyPreferredName: "",
    vendorLandingPageURL: "",
    schemaVersion: 0,
  };
}

export const MsgUpdateVendorInfo = {
  encode(message: MsgUpdateVendorInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).int32(message.vendorID);
    }
    if (message.vendorName !== "") {
      writer.uint32(26).string(message.vendorName);
    }
    if (message.companyLegalName !== "") {
      writer.uint32(34).string(message.companyLegalName);
    }
    if (message.companyPreferredName !== "") {
      writer.uint32(42).string(message.companyPreferredName);
    }
    if (message.vendorLandingPageURL !== "") {
      writer.uint32(50).string(message.vendorLandingPageURL);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(56).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateVendorInfo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateVendorInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.vendorID = reader.int32();
          break;
        case 3:
          message.vendorName = reader.string();
          break;
        case 4:
          message.companyLegalName = reader.string();
          break;
        case 5:
          message.companyPreferredName = reader.string();
          break;
        case 6:
          message.vendorLandingPageURL = reader.string();
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

  fromJSON(object: any): MsgUpdateVendorInfo {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      vendorID: isSet(object.vendorID) ? Number(object.vendorID) : 0,
      vendorName: isSet(object.vendorName) ? String(object.vendorName) : "",
      companyLegalName: isSet(object.companyLegalName) ? String(object.companyLegalName) : "",
      companyPreferredName: isSet(object.companyPreferredName) ? String(object.companyPreferredName) : "",
      vendorLandingPageURL: isSet(object.vendorLandingPageURL) ? String(object.vendorLandingPageURL) : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: MsgUpdateVendorInfo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.vendorID !== undefined && (obj.vendorID = Math.round(message.vendorID));
    message.vendorName !== undefined && (obj.vendorName = message.vendorName);
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName);
    message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName);
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateVendorInfo>, I>>(object: I): MsgUpdateVendorInfo {
    const message = createBaseMsgUpdateVendorInfo();
    message.creator = object.creator ?? "";
    message.vendorID = object.vendorID ?? 0;
    message.vendorName = object.vendorName ?? "";
    message.companyLegalName = object.companyLegalName ?? "";
    message.companyPreferredName = object.companyPreferredName ?? "";
    message.vendorLandingPageURL = object.vendorLandingPageURL ?? "";
    message.schemaVersion = object.schemaVersion ?? 0;
    return message;
  },
};

function createBaseMsgUpdateVendorInfoResponse(): MsgUpdateVendorInfoResponse {
  return {};
}

export const MsgUpdateVendorInfoResponse = {
  encode(_: MsgUpdateVendorInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateVendorInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateVendorInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateVendorInfoResponse {
    return {};
  },

  toJSON(_: MsgUpdateVendorInfoResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateVendorInfoResponse>, I>>(_: I): MsgUpdateVendorInfoResponse {
    const message = createBaseMsgUpdateVendorInfoResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVendorInfo(request: MsgCreateVendorInfo): Promise<MsgCreateVendorInfoResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateVendorInfo(request: MsgUpdateVendorInfo): Promise<MsgUpdateVendorInfoResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateVendorInfo = this.CreateVendorInfo.bind(this);
    this.UpdateVendorInfo = this.UpdateVendorInfo.bind(this);
  }
  CreateVendorInfo(request: MsgCreateVendorInfo): Promise<MsgCreateVendorInfoResponse> {
    const data = MsgCreateVendorInfo.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.vendorinfo.Msg",
      "CreateVendorInfo",
      data,
    );
    return promise.then((data) => MsgCreateVendorInfoResponse.decode(new _m0.Reader(data)));
  }

  UpdateVendorInfo(request: MsgUpdateVendorInfo): Promise<MsgUpdateVendorInfoResponse> {
    const data = MsgUpdateVendorInfo.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.vendorinfo.Msg",
      "UpdateVendorInfo",
      data,
    );
    return promise.then((data) => MsgUpdateVendorInfoResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
