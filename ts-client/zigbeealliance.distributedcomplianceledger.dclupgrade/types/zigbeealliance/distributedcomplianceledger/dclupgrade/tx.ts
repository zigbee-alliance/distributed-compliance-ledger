/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Plan } from "../../../cosmos/upgrade/v1beta1/upgrade";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclupgrade";

export interface MsgProposeUpgrade {
  creator: string;
  plan: Plan | undefined;
  info: string;
  time: number;
}

export interface MsgProposeUpgradeResponse {
}

export interface MsgApproveUpgrade {
  creator: string;
  name: string;
  info: string;
  time: number;
}

export interface MsgApproveUpgradeResponse {
}

export interface MsgRejectUpgrade {
  creator: string;
  name: string;
  info: string;
  time: number;
}

export interface MsgRejectUpgradeResponse {
}

function createBaseMsgProposeUpgrade(): MsgProposeUpgrade {
  return { creator: "", plan: undefined, info: "", time: 0 };
}

export const MsgProposeUpgrade = {
  encode(message: MsgProposeUpgrade, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.plan !== undefined) {
      Plan.encode(message.plan, writer.uint32(18).fork()).ldelim();
    }
    if (message.info !== "") {
      writer.uint32(26).string(message.info);
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProposeUpgrade {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProposeUpgrade();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.plan = Plan.decode(reader, reader.uint32());
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

  fromJSON(object: any): MsgProposeUpgrade {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      plan: isSet(object.plan) ? Plan.fromJSON(object.plan) : undefined,
      info: isSet(object.info) ? String(object.info) : "",
      time: isSet(object.time) ? Number(object.time) : 0,
    };
  },

  toJSON(message: MsgProposeUpgrade): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.plan !== undefined && (obj.plan = message.plan ? Plan.toJSON(message.plan) : undefined);
    message.info !== undefined && (obj.info = message.info);
    message.time !== undefined && (obj.time = Math.round(message.time));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgProposeUpgrade>, I>>(object: I): MsgProposeUpgrade {
    const message = createBaseMsgProposeUpgrade();
    message.creator = object.creator ?? "";
    message.plan = (object.plan !== undefined && object.plan !== null) ? Plan.fromPartial(object.plan) : undefined;
    message.info = object.info ?? "";
    message.time = object.time ?? 0;
    return message;
  },
};

function createBaseMsgProposeUpgradeResponse(): MsgProposeUpgradeResponse {
  return {};
}

export const MsgProposeUpgradeResponse = {
  encode(_: MsgProposeUpgradeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProposeUpgradeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProposeUpgradeResponse();
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

  fromJSON(_: any): MsgProposeUpgradeResponse {
    return {};
  },

  toJSON(_: MsgProposeUpgradeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgProposeUpgradeResponse>, I>>(_: I): MsgProposeUpgradeResponse {
    const message = createBaseMsgProposeUpgradeResponse();
    return message;
  },
};

function createBaseMsgApproveUpgrade(): MsgApproveUpgrade {
  return { creator: "", name: "", info: "", time: 0 };
}

export const MsgApproveUpgrade = {
  encode(message: MsgApproveUpgrade, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.info !== "") {
      writer.uint32(26).string(message.info);
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveUpgrade {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveUpgrade();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
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

  fromJSON(object: any): MsgApproveUpgrade {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      info: isSet(object.info) ? String(object.info) : "",
      time: isSet(object.time) ? Number(object.time) : 0,
    };
  },

  toJSON(message: MsgApproveUpgrade): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.info !== undefined && (obj.info = message.info);
    message.time !== undefined && (obj.time = Math.round(message.time));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveUpgrade>, I>>(object: I): MsgApproveUpgrade {
    const message = createBaseMsgApproveUpgrade();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.info = object.info ?? "";
    message.time = object.time ?? 0;
    return message;
  },
};

function createBaseMsgApproveUpgradeResponse(): MsgApproveUpgradeResponse {
  return {};
}

export const MsgApproveUpgradeResponse = {
  encode(_: MsgApproveUpgradeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveUpgradeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveUpgradeResponse();
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

  fromJSON(_: any): MsgApproveUpgradeResponse {
    return {};
  },

  toJSON(_: MsgApproveUpgradeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveUpgradeResponse>, I>>(_: I): MsgApproveUpgradeResponse {
    const message = createBaseMsgApproveUpgradeResponse();
    return message;
  },
};

function createBaseMsgRejectUpgrade(): MsgRejectUpgrade {
  return { creator: "", name: "", info: "", time: 0 };
}

export const MsgRejectUpgrade = {
  encode(message: MsgRejectUpgrade, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.info !== "") {
      writer.uint32(26).string(message.info);
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectUpgrade {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRejectUpgrade();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
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

  fromJSON(object: any): MsgRejectUpgrade {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      info: isSet(object.info) ? String(object.info) : "",
      time: isSet(object.time) ? Number(object.time) : 0,
    };
  },

  toJSON(message: MsgRejectUpgrade): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.info !== undefined && (obj.info = message.info);
    message.time !== undefined && (obj.time = Math.round(message.time));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRejectUpgrade>, I>>(object: I): MsgRejectUpgrade {
    const message = createBaseMsgRejectUpgrade();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.info = object.info ?? "";
    message.time = object.time ?? 0;
    return message;
  },
};

function createBaseMsgRejectUpgradeResponse(): MsgRejectUpgradeResponse {
  return {};
}

export const MsgRejectUpgradeResponse = {
  encode(_: MsgRejectUpgradeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRejectUpgradeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRejectUpgradeResponse();
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

  fromJSON(_: any): MsgRejectUpgradeResponse {
    return {};
  },

  toJSON(_: MsgRejectUpgradeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRejectUpgradeResponse>, I>>(_: I): MsgRejectUpgradeResponse {
    const message = createBaseMsgRejectUpgradeResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse>;
  ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RejectUpgrade(request: MsgRejectUpgrade): Promise<MsgRejectUpgradeResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ProposeUpgrade = this.ProposeUpgrade.bind(this);
    this.ApproveUpgrade = this.ApproveUpgrade.bind(this);
    this.RejectUpgrade = this.RejectUpgrade.bind(this);
  }
  ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse> {
    const data = MsgProposeUpgrade.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.dclupgrade.Msg",
      "ProposeUpgrade",
      data,
    );
    return promise.then((data) => MsgProposeUpgradeResponse.decode(new _m0.Reader(data)));
  }

  ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse> {
    const data = MsgApproveUpgrade.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.dclupgrade.Msg",
      "ApproveUpgrade",
      data,
    );
    return promise.then((data) => MsgApproveUpgradeResponse.decode(new _m0.Reader(data)));
  }

  RejectUpgrade(request: MsgRejectUpgrade): Promise<MsgRejectUpgradeResponse> {
    const data = MsgRejectUpgrade.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.dclupgrade.Msg",
      "RejectUpgrade",
      data,
    );
    return promise.then((data) => MsgRejectUpgradeResponse.decode(new _m0.Reader(data)));
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
