/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Grant } from "./grant";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";

export interface ProposedDisableValidator {
  address: string;
  creator: string;
  approvals: Grant[];
  rejects: Grant[];
  schemaVersion: number;
}

function createBaseProposedDisableValidator(): ProposedDisableValidator {
  return { address: "", creator: "", approvals: [], rejects: [], schemaVersion: 0 };
}

export const ProposedDisableValidator = {
  encode(message: ProposedDisableValidator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.rejects) {
      Grant.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(40).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ProposedDisableValidator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseProposedDisableValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.approvals.push(Grant.decode(reader, reader.uint32()));
          break;
        case 4:
          message.rejects.push(Grant.decode(reader, reader.uint32()));
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

  fromJSON(object: any): ProposedDisableValidator {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      approvals: Array.isArray(object?.approvals) ? object.approvals.map((e: any) => Grant.fromJSON(e)) : [],
      rejects: Array.isArray(object?.rejects) ? object.rejects.map((e: any) => Grant.fromJSON(e)) : [],
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: ProposedDisableValidator): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.creator !== undefined && (obj.creator = message.creator);
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => e ? Grant.toJSON(e) : undefined);
    } else {
      obj.approvals = [];
    }
    if (message.rejects) {
      obj.rejects = message.rejects.map((e) => e ? Grant.toJSON(e) : undefined);
    } else {
      obj.rejects = [];
    }
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ProposedDisableValidator>, I>>(object: I): ProposedDisableValidator {
    const message = createBaseProposedDisableValidator();
    message.address = object.address ?? "";
    message.creator = object.creator ?? "";
    message.approvals = object.approvals?.map((e) => Grant.fromPartial(e)) || [];
    message.rejects = object.rejects?.map((e) => Grant.fromPartial(e)) || [];
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
