/* eslint-disable */
import { NewVendorInfo } from '../vendorinfo/new_vendor_info'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

/** GenesisState defines the vendorinfo module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  newVendorInfoList: NewVendorInfo[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.newVendorInfoList) {
      NewVendorInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.newVendorInfoList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.newVendorInfoList.push(NewVendorInfo.decode(reader, reader.uint32()))
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
    message.newVendorInfoList = []
    if (object.newVendorInfoList !== undefined && object.newVendorInfoList !== null) {
      for (const e of object.newVendorInfoList) {
        message.newVendorInfoList.push(NewVendorInfo.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.newVendorInfoList) {
      obj.newVendorInfoList = message.newVendorInfoList.map((e) => (e ? NewVendorInfo.toJSON(e) : undefined))
    } else {
      obj.newVendorInfoList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.newVendorInfoList = []
    if (object.newVendorInfoList !== undefined && object.newVendorInfoList !== null) {
      for (const e of object.newVendorInfoList) {
        message.newVendorInfoList.push(NewVendorInfo.fromPartial(e))
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
