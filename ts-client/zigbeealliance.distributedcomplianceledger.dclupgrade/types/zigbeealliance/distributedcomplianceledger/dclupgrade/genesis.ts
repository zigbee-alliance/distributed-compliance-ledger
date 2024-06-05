/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { ApprovedUpgrade } from "./approved_upgrade";
import { ProposedUpgrade } from "./proposed_upgrade";
import { RejectedUpgrade } from "./rejected_upgrade";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclupgrade";

/** GenesisState defines the dclupgrade module's genesis state. */
export interface GenesisState {
  proposedUpgradeList: ProposedUpgrade[];
  approvedUpgradeList: ApprovedUpgrade[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  rejectedUpgradeList: RejectedUpgrade[];
}

function createBaseGenesisState(): GenesisState {
  return { proposedUpgradeList: [], approvedUpgradeList: [], rejectedUpgradeList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.proposedUpgradeList) {
      ProposedUpgrade.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.approvedUpgradeList) {
      ApprovedUpgrade.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.rejectedUpgradeList) {
      RejectedUpgrade.encode(v!, writer.uint32(26).fork()).ldelim();
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
          message.proposedUpgradeList.push(ProposedUpgrade.decode(reader, reader.uint32()));
          break;
        case 2:
          message.approvedUpgradeList.push(ApprovedUpgrade.decode(reader, reader.uint32()));
          break;
        case 3:
          message.rejectedUpgradeList.push(RejectedUpgrade.decode(reader, reader.uint32()));
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
      proposedUpgradeList: Array.isArray(object?.proposedUpgradeList)
        ? object.proposedUpgradeList.map((e: any) => ProposedUpgrade.fromJSON(e))
        : [],
      approvedUpgradeList: Array.isArray(object?.approvedUpgradeList)
        ? object.approvedUpgradeList.map((e: any) => ApprovedUpgrade.fromJSON(e))
        : [],
      rejectedUpgradeList: Array.isArray(object?.rejectedUpgradeList)
        ? object.rejectedUpgradeList.map((e: any) => RejectedUpgrade.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.proposedUpgradeList) {
      obj.proposedUpgradeList = message.proposedUpgradeList.map((e) => e ? ProposedUpgrade.toJSON(e) : undefined);
    } else {
      obj.proposedUpgradeList = [];
    }
    if (message.approvedUpgradeList) {
      obj.approvedUpgradeList = message.approvedUpgradeList.map((e) => e ? ApprovedUpgrade.toJSON(e) : undefined);
    } else {
      obj.approvedUpgradeList = [];
    }
    if (message.rejectedUpgradeList) {
      obj.rejectedUpgradeList = message.rejectedUpgradeList.map((e) => e ? RejectedUpgrade.toJSON(e) : undefined);
    } else {
      obj.rejectedUpgradeList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.proposedUpgradeList = object.proposedUpgradeList?.map((e) => ProposedUpgrade.fromPartial(e)) || [];
    message.approvedUpgradeList = object.approvedUpgradeList?.map((e) => ApprovedUpgrade.fromPartial(e)) || [];
    message.rejectedUpgradeList = object.rejectedUpgradeList?.map((e) => RejectedUpgrade.fromPartial(e)) || [];
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
