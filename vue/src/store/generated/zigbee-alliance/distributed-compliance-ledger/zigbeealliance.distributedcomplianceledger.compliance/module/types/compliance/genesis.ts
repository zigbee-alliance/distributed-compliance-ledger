/* eslint-disable */
import { ComplianceInfo } from '../compliance/compliance_info'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

/** GenesisState defines the compliance module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  complianceInfoList: ComplianceInfo[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.complianceInfoList) {
      ComplianceInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.complianceInfoList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.complianceInfoList.push(ComplianceInfo.decode(reader, reader.uint32()))
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
    message.complianceInfoList = []
    if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
      for (const e of object.complianceInfoList) {
        message.complianceInfoList.push(ComplianceInfo.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.complianceInfoList) {
      obj.complianceInfoList = message.complianceInfoList.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined))
    } else {
      obj.complianceInfoList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.complianceInfoList = []
    if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
      for (const e of object.complianceInfoList) {
        message.complianceInfoList.push(ComplianceInfo.fromPartial(e))
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
