/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Any } from "../../../google/protobuf/any";
import { Description } from "./description";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";

export interface Validator {
  /** the account address of validator owner */
  owner: string;
  /** description of the validator */
  description:
    | Description
    | undefined;
  /** the consensus public key of the tendermint validator */
  pubKey:
    | Any
    | undefined;
  /** validator consensus power */
  power: number;
  /** has the validator been removed from validator set */
  jailed: boolean;
  /** the reason of validator jailing */
  jailedReason: string;
  schemaVersion: number;
}

function createBaseValidator(): Validator {
  return {
    owner: "",
    description: undefined,
    pubKey: undefined,
    power: 0,
    jailed: false,
    jailedReason: "",
    schemaVersion: 0,
  };
}

export const Validator = {
  encode(message: Validator, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.description !== undefined) {
      Description.encode(message.description, writer.uint32(18).fork()).ldelim();
    }
    if (message.pubKey !== undefined) {
      Any.encode(message.pubKey, writer.uint32(26).fork()).ldelim();
    }
    if (message.power !== 0) {
      writer.uint32(32).int32(message.power);
    }
    if (message.jailed === true) {
      writer.uint32(40).bool(message.jailed);
    }
    if (message.jailedReason !== "") {
      writer.uint32(50).string(message.jailedReason);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(56).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Validator {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseValidator();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.description = Description.decode(reader, reader.uint32());
          break;
        case 3:
          message.pubKey = Any.decode(reader, reader.uint32());
          break;
        case 4:
          message.power = reader.int32();
          break;
        case 5:
          message.jailed = reader.bool();
          break;
        case 6:
          message.jailedReason = reader.string();
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

  fromJSON(object: any): Validator {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      description: isSet(object.description) ? Description.fromJSON(object.description) : undefined,
      pubKey: isSet(object.pubKey) ? Any.fromJSON(object.pubKey) : undefined,
      power: isSet(object.power) ? Number(object.power) : 0,
      jailed: isSet(object.jailed) ? Boolean(object.jailed) : false,
      jailedReason: isSet(object.jailedReason) ? String(object.jailedReason) : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: Validator): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.description !== undefined
      && (obj.description = message.description ? Description.toJSON(message.description) : undefined);
    message.pubKey !== undefined && (obj.pubKey = message.pubKey ? Any.toJSON(message.pubKey) : undefined);
    message.power !== undefined && (obj.power = Math.round(message.power));
    message.jailed !== undefined && (obj.jailed = message.jailed);
    message.jailedReason !== undefined && (obj.jailedReason = message.jailedReason);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Validator>, I>>(object: I): Validator {
    const message = createBaseValidator();
    message.owner = object.owner ?? "";
    message.description = (object.description !== undefined && object.description !== null)
      ? Description.fromPartial(object.description)
      : undefined;
    message.pubKey = (object.pubKey !== undefined && object.pubKey !== null)
      ? Any.fromPartial(object.pubKey)
      : undefined;
    message.power = object.power ?? 0;
    message.jailed = object.jailed ?? false;
    message.jailedReason = object.jailedReason ?? "";
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
