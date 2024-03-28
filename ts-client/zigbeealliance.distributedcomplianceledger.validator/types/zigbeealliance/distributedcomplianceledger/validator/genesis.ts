/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DisabledValidator } from "./disabled_validator";
import { LastValidatorPower } from "./last_validator_power";
import { ProposedDisableValidator } from "./proposed_disable_validator";
import { RejectedDisableValidator } from "./rejected_validator";
import { Validator } from "./validator";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";

/** GenesisState defines the validator module's genesis state. */
export interface GenesisState {
  validatorList: Validator[];
  lastValidatorPowerList: LastValidatorPower[];
  proposedDisableValidatorList: ProposedDisableValidator[];
  disabledValidatorList: DisabledValidator[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  rejectedValidatorList: RejectedDisableValidator[];
}

function createBaseGenesisState(): GenesisState {
  return {
    validatorList: [],
    lastValidatorPowerList: [],
    proposedDisableValidatorList: [],
    disabledValidatorList: [],
    rejectedValidatorList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.validatorList) {
      Validator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.lastValidatorPowerList) {
      LastValidatorPower.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.proposedDisableValidatorList) {
      ProposedDisableValidator.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.disabledValidatorList) {
      DisabledValidator.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.rejectedValidatorList) {
      RejectedDisableValidator.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validatorList.push(Validator.decode(reader, reader.uint32()));
          break;
        case 2:
          message.lastValidatorPowerList.push(LastValidatorPower.decode(reader, reader.uint32()));
          break;
        case 3:
          message.proposedDisableValidatorList.push(ProposedDisableValidator.decode(reader, reader.uint32()));
          break;
        case 4:
          message.disabledValidatorList.push(DisabledValidator.decode(reader, reader.uint32()));
          break;
        case 5:
          message.rejectedValidatorList.push(RejectedDisableValidator.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      validatorList: Array.isArray(object?.validatorList)
        ? object.validatorList.map((e: any) => Validator.fromJSON(e))
        : [],
      lastValidatorPowerList: Array.isArray(object?.lastValidatorPowerList)
        ? object.lastValidatorPowerList.map((e: any) => LastValidatorPower.fromJSON(e))
        : [],
      proposedDisableValidatorList: Array.isArray(object?.proposedDisableValidatorList)
        ? object.proposedDisableValidatorList.map((e: any) => ProposedDisableValidator.fromJSON(e))
        : [],
      disabledValidatorList: Array.isArray(object?.disabledValidatorList)
        ? object.disabledValidatorList.map((e: any) => DisabledValidator.fromJSON(e))
        : [],
      rejectedValidatorList: Array.isArray(object?.rejectedValidatorList)
        ? object.rejectedValidatorList.map((e: any) => RejectedDisableValidator.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.validatorList) {
      obj.validatorList = message.validatorList.map((e) => e ? Validator.toJSON(e) : undefined);
    } else {
      obj.validatorList = [];
    }
    if (message.lastValidatorPowerList) {
      obj.lastValidatorPowerList = message.lastValidatorPowerList.map((e) =>
        e ? LastValidatorPower.toJSON(e) : undefined
      );
    } else {
      obj.lastValidatorPowerList = [];
    }
    if (message.proposedDisableValidatorList) {
      obj.proposedDisableValidatorList = message.proposedDisableValidatorList.map((e) =>
        e ? ProposedDisableValidator.toJSON(e) : undefined
      );
    } else {
      obj.proposedDisableValidatorList = [];
    }
    if (message.disabledValidatorList) {
      obj.disabledValidatorList = message.disabledValidatorList.map((e) => e ? DisabledValidator.toJSON(e) : undefined);
    } else {
      obj.disabledValidatorList = [];
    }
    if (message.rejectedValidatorList) {
      obj.rejectedValidatorList = message.rejectedValidatorList.map((e) =>
        e ? RejectedDisableValidator.toJSON(e) : undefined
      );
    } else {
      obj.rejectedValidatorList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.validatorList = object.validatorList?.map((e) => Validator.fromPartial(e)) || [];
    message.lastValidatorPowerList = object.lastValidatorPowerList?.map((e) => LastValidatorPower.fromPartial(e)) || [];
    message.proposedDisableValidatorList =
      object.proposedDisableValidatorList?.map((e) => ProposedDisableValidator.fromPartial(e)) || [];
    message.disabledValidatorList = object.disabledValidatorList?.map((e) => DisabledValidator.fromPartial(e)) || [];
    message.rejectedValidatorList = object.rejectedValidatorList?.map((e) => RejectedDisableValidator.fromPartial(e))
      || [];
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
