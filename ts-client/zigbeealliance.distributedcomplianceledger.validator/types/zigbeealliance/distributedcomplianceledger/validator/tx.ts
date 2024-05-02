/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Any } from "../../../google/protobuf/any";
import { Description } from "./description";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";

export interface MsgCreateValidator {
  signer: string;
  pubKey: Any | undefined;
  description: Description | undefined;
}

export interface MsgCreateValidatorResponse {
}

export interface MsgProposeDisableValidator {
  creator: string;
  address: string;
  info: string;
  time: number;
}

export interface MsgProposeDisableValidatorResponse {
}

export interface MsgApproveDisableValidator {
  creator: string;
  address: string;
  info: string;
  time: number;
}

export interface MsgApproveDisableValidatorResponse {
}

export interface MsgDisableValidator {
  creator: string;
}

export interface MsgDisableValidatorResponse {
}

export interface MsgEnableValidator {
  creator: string;
}

export interface MsgEnableValidatorResponse {
}

export interface MsgRejectDisableValidator {
  creator: string;
  address: string;
  info: string;
  time: number;
}

export interface MsgRejectDisableValidatorResponse {
}

function createBaseMsgCreateValidator(): MsgCreateValidator {
  return { signer: "", pubKey: undefined, description: undefined };
}

export const MsgCreateValidator = {
  encode(message: MsgCreateValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.signer !== "") {
      writer.uint32(10).string(message.signer);
    }
    if (message.pubKey !== undefined) {
      Any.encode(message.pubKey, writer.uint32(18).fork()).ldelim();
    }
    if (message.description !== undefined) {
      Description.encode(message.description, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string();
          break;
        case 2:
          message.pubKey = Any.decode(reader, reader.uint32());
          break;
        case 3:
          message.description = Description.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateValidator {
    return {
      signer: isSet(object.signer) ? String(object.signer) : "",
      pubKey: isSet(object.pubKey) ? Any.fromJSON(object.pubKey) : undefined,
      description: isSet(object.description) ? Description.fromJSON(object.description) : undefined,
    };
  },

  toJSON(message: MsgCreateValidator): unknown {
    const obj: any = {};
    message.signer !== undefined && (obj.signer = message.signer);
    message.pubKey !== undefined && (obj.pubKey = message.pubKey ? Any.toJSON(message.pubKey) : undefined);
    message.description !== undefined
      && (obj.description = message.description ? Description.toJSON(message.description) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateValidator>, I>>(object: I): MsgCreateValidator {
    const message = createBaseMsgCreateValidator();
    message.signer = object.signer ?? "";
    message.pubKey = (object.pubKey !== undefined && object.pubKey !== null)
      ? Any.fromPartial(object.pubKey)
      : undefined;
    message.description = (object.description !== undefined && object.description !== null)
      ? Description.fromPartial(object.description)
      : undefined;
    return message;
  },
};

function createBaseMsgCreateValidatorResponse(): MsgCreateValidatorResponse {
  return {};
}

export const MsgCreateValidatorResponse = {
  encode(_: MsgCreateValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateValidatorResponse();
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

  fromJSON(_: any): MsgCreateValidatorResponse {
    return {};
  },

  toJSON(_: MsgCreateValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateValidatorResponse>, I>>(_: I): MsgCreateValidatorResponse {
    const message = createBaseMsgCreateValidatorResponse();
    return message;
  },
};

function createBaseMsgProposeDisableValidator(): MsgProposeDisableValidator {
  return { creator: "", address: "", info: "", time: 0 };
}

export const MsgProposeDisableValidator = {
  encode(message: MsgProposeDisableValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.info !== "") {
      writer.uint32(26).string(message.info);
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProposeDisableValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProposeDisableValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.info = reader.string();
          break;
        case 4:
          message.time = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgProposeDisableValidator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      info: isSet(object.info) ? String(object.info) : "",
      time: isSet(object.time) ? Number(object.time) : 0,
    };
  },

  toJSON(message: MsgProposeDisableValidator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.info !== undefined && (obj.info = message.info);
    message.time !== undefined && (obj.time = Math.round(message.time));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgProposeDisableValidator>, I>>(object: I): MsgProposeDisableValidator {
    const message = createBaseMsgProposeDisableValidator();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.info = object.info ?? "";
    message.time = object.time ?? 0;
    return message;
  },
};

function createBaseMsgProposeDisableValidatorResponse(): MsgProposeDisableValidatorResponse {
  return {};
}

export const MsgProposeDisableValidatorResponse = {
  encode(_: MsgProposeDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProposeDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProposeDisableValidatorResponse();
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

  fromJSON(_: any): MsgProposeDisableValidatorResponse {
    return {};
  },

  toJSON(_: MsgProposeDisableValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgProposeDisableValidatorResponse>, I>>(
    _: I,
  ): MsgProposeDisableValidatorResponse {
    const message = createBaseMsgProposeDisableValidatorResponse();
    return message;
  },
};

function createBaseMsgApproveDisableValidator(): MsgApproveDisableValidator {
  return { creator: "", address: "", info: "", time: 0 };
}

export const MsgApproveDisableValidator = {
  encode(message: MsgApproveDisableValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.info !== "") {
      writer.uint32(26).string(message.info);
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveDisableValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveDisableValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.info = reader.string();
          break;
        case 4:
          message.time = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApproveDisableValidator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      info: isSet(object.info) ? String(object.info) : "",
      time: isSet(object.time) ? Number(object.time) : 0,
    };
  },

  toJSON(message: MsgApproveDisableValidator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.info !== undefined && (obj.info = message.info);
    message.time !== undefined && (obj.time = Math.round(message.time));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveDisableValidator>, I>>(object: I): MsgApproveDisableValidator {
    const message = createBaseMsgApproveDisableValidator();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.info = object.info ?? "";
    message.time = object.time ?? 0;
    return message;
  },
};

function createBaseMsgApproveDisableValidatorResponse(): MsgApproveDisableValidatorResponse {
  return {};
}

export const MsgApproveDisableValidatorResponse = {
  encode(_: MsgApproveDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveDisableValidatorResponse();
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

  fromJSON(_: any): MsgApproveDisableValidatorResponse {
    return {};
  },

  toJSON(_: MsgApproveDisableValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveDisableValidatorResponse>, I>>(
    _: I,
  ): MsgApproveDisableValidatorResponse {
    const message = createBaseMsgApproveDisableValidatorResponse();
    return message;
  },
};

function createBaseMsgDisableValidator(): MsgDisableValidator {
  return { creator: "" };
}

export const MsgDisableValidator = {
  encode(message: MsgDisableValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDisableValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDisableValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDisableValidator {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgDisableValidator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDisableValidator>, I>>(object: I): MsgDisableValidator {
    const message = createBaseMsgDisableValidator();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgDisableValidatorResponse(): MsgDisableValidatorResponse {
  return {};
}

export const MsgDisableValidatorResponse = {
  encode(_: MsgDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDisableValidatorResponse();
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

  fromJSON(_: any): MsgDisableValidatorResponse {
    return {};
  },

  toJSON(_: MsgDisableValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDisableValidatorResponse>, I>>(_: I): MsgDisableValidatorResponse {
    const message = createBaseMsgDisableValidatorResponse();
    return message;
  },
};

function createBaseMsgEnableValidator(): MsgEnableValidator {
  return { creator: "" };
}

export const MsgEnableValidator = {
  encode(message: MsgEnableValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnableValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnableValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnableValidator {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgEnableValidator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnableValidator>, I>>(object: I): MsgEnableValidator {
    const message = createBaseMsgEnableValidator();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgEnableValidatorResponse(): MsgEnableValidatorResponse {
  return {};
}

export const MsgEnableValidatorResponse = {
  encode(_: MsgEnableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnableValidatorResponse();
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

  fromJSON(_: any): MsgEnableValidatorResponse {
    return {};
  },

  toJSON(_: MsgEnableValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnableValidatorResponse>, I>>(_: I): MsgEnableValidatorResponse {
    const message = createBaseMsgEnableValidatorResponse();
    return message;
  },
};

function createBaseMsgRejectDisableValidator(): MsgRejectDisableValidator {
  return { creator: "", address: "", info: "", time: 0 };
}

export const MsgRejectDisableValidator = {
  encode(message: MsgRejectDisableValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.info !== "") {
      writer.uint32(26).string(message.info);
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectDisableValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRejectDisableValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.info = reader.string();
          break;
        case 4:
          message.time = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRejectDisableValidator {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      info: isSet(object.info) ? String(object.info) : "",
      time: isSet(object.time) ? Number(object.time) : 0,
    };
  },

  toJSON(message: MsgRejectDisableValidator): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.info !== undefined && (obj.info = message.info);
    message.time !== undefined && (obj.time = Math.round(message.time));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRejectDisableValidator>, I>>(object: I): MsgRejectDisableValidator {
    const message = createBaseMsgRejectDisableValidator();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.info = object.info ?? "";
    message.time = object.time ?? 0;
    return message;
  },
};

function createBaseMsgRejectDisableValidatorResponse(): MsgRejectDisableValidatorResponse {
  return {};
}

export const MsgRejectDisableValidatorResponse = {
  encode(_: MsgRejectDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRejectDisableValidatorResponse();
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

  fromJSON(_: any): MsgRejectDisableValidatorResponse {
    return {};
  },

  toJSON(_: MsgRejectDisableValidatorResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRejectDisableValidatorResponse>, I>>(
    _: I,
  ): MsgRejectDisableValidatorResponse {
    const message = createBaseMsgRejectDisableValidatorResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse>;
  ProposeDisableValidator(request: MsgProposeDisableValidator): Promise<MsgProposeDisableValidatorResponse>;
  ApproveDisableValidator(request: MsgApproveDisableValidator): Promise<MsgApproveDisableValidatorResponse>;
  DisableValidator(request: MsgDisableValidator): Promise<MsgDisableValidatorResponse>;
  EnableValidator(request: MsgEnableValidator): Promise<MsgEnableValidatorResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RejectDisableValidator(request: MsgRejectDisableValidator): Promise<MsgRejectDisableValidatorResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateValidator = this.CreateValidator.bind(this);
    this.ProposeDisableValidator = this.ProposeDisableValidator.bind(this);
    this.ApproveDisableValidator = this.ApproveDisableValidator.bind(this);
    this.DisableValidator = this.DisableValidator.bind(this);
    this.EnableValidator = this.EnableValidator.bind(this);
    this.RejectDisableValidator = this.RejectDisableValidator.bind(this);
  }
  CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse> {
    const data = MsgCreateValidator.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Msg",
      "CreateValidator",
      data,
    );
    return promise.then((data) => MsgCreateValidatorResponse.decode(new _m0.Reader(data)));
  }

  ProposeDisableValidator(request: MsgProposeDisableValidator): Promise<MsgProposeDisableValidatorResponse> {
    const data = MsgProposeDisableValidator.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Msg",
      "ProposeDisableValidator",
      data,
    );
    return promise.then((data) => MsgProposeDisableValidatorResponse.decode(new _m0.Reader(data)));
  }

  ApproveDisableValidator(request: MsgApproveDisableValidator): Promise<MsgApproveDisableValidatorResponse> {
    const data = MsgApproveDisableValidator.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Msg",
      "ApproveDisableValidator",
      data,
    );
    return promise.then((data) => MsgApproveDisableValidatorResponse.decode(new _m0.Reader(data)));
  }

  DisableValidator(request: MsgDisableValidator): Promise<MsgDisableValidatorResponse> {
    const data = MsgDisableValidator.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Msg",
      "DisableValidator",
      data,
    );
    return promise.then((data) => MsgDisableValidatorResponse.decode(new _m0.Reader(data)));
  }

  EnableValidator(request: MsgEnableValidator): Promise<MsgEnableValidatorResponse> {
    const data = MsgEnableValidator.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Msg",
      "EnableValidator",
      data,
    );
    return promise.then((data) => MsgEnableValidatorResponse.decode(new _m0.Reader(data)));
  }

  RejectDisableValidator(request: MsgRejectDisableValidator): Promise<MsgRejectDisableValidatorResponse> {
    const data = MsgRejectDisableValidator.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Msg",
      "RejectDisableValidator",
      data,
    );
    return promise.then((data) => MsgRejectDisableValidatorResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
