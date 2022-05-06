/* eslint-disable */
import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade'
import { ApprovedUpgrade } from '../dclupgrade/approved_upgrade'
import { RejectedUpgrade } from '../dclupgrade/rejected_upgrade'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade'

/** GenesisState defines the dclupgrade module's genesis state. */
export interface GenesisState {
  proposedUpgradeList: ProposedUpgrade[]
  approvedUpgradeList: ApprovedUpgrade[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  rejectedUpgradeList: RejectedUpgrade[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.proposedUpgradeList) {
      ProposedUpgrade.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.approvedUpgradeList) {
      ApprovedUpgrade.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    for (const v of message.rejectedUpgradeList) {
      RejectedUpgrade.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.proposedUpgradeList = []
    message.approvedUpgradeList = []
    message.rejectedUpgradeList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedUpgradeList.push(ProposedUpgrade.decode(reader, reader.uint32()))
          break
        case 2:
          message.approvedUpgradeList.push(ApprovedUpgrade.decode(reader, reader.uint32()))
          break
        case 3:
          message.rejectedUpgradeList.push(RejectedUpgrade.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.proposedUpgradeList = []
    message.approvedUpgradeList = []
    message.rejectedUpgradeList = []
    if (object.proposedUpgradeList !== undefined && object.proposedUpgradeList !== null) {
      for (const e of object.proposedUpgradeList) {
        message.proposedUpgradeList.push(ProposedUpgrade.fromJSON(e))
      }
    }
    if (object.approvedUpgradeList !== undefined && object.approvedUpgradeList !== null) {
      for (const e of object.approvedUpgradeList) {
        message.approvedUpgradeList.push(ApprovedUpgrade.fromJSON(e))
      }
    }
    if (object.rejectedUpgradeList !== undefined && object.rejectedUpgradeList !== null) {
      for (const e of object.rejectedUpgradeList) {
        message.rejectedUpgradeList.push(RejectedUpgrade.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.proposedUpgradeList) {
      obj.proposedUpgradeList = message.proposedUpgradeList.map((e) => (e ? ProposedUpgrade.toJSON(e) : undefined))
    } else {
      obj.proposedUpgradeList = []
    }
    if (message.approvedUpgradeList) {
      obj.approvedUpgradeList = message.approvedUpgradeList.map((e) => (e ? ApprovedUpgrade.toJSON(e) : undefined))
    } else {
      obj.approvedUpgradeList = []
    }
    if (message.rejectedUpgradeList) {
      obj.rejectedUpgradeList = message.rejectedUpgradeList.map((e) => (e ? RejectedUpgrade.toJSON(e) : undefined))
    } else {
      obj.rejectedUpgradeList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.proposedUpgradeList = []
    message.approvedUpgradeList = []
    message.rejectedUpgradeList = []
    if (object.proposedUpgradeList !== undefined && object.proposedUpgradeList !== null) {
      for (const e of object.proposedUpgradeList) {
        message.proposedUpgradeList.push(ProposedUpgrade.fromPartial(e))
      }
    }
    if (object.approvedUpgradeList !== undefined && object.approvedUpgradeList !== null) {
      for (const e of object.approvedUpgradeList) {
        message.approvedUpgradeList.push(ApprovedUpgrade.fromPartial(e))
      }
    }
    if (object.rejectedUpgradeList !== undefined && object.rejectedUpgradeList !== null) {
      for (const e of object.rejectedUpgradeList) {
        message.rejectedUpgradeList.push(RejectedUpgrade.fromPartial(e))
      }
    }
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
