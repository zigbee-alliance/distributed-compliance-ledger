/* eslint-disable */
import { VendorInfo } from '../vendorinfo/vendor_info'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

/** GenesisState defines the vendorinfo module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  vendorInfoList: VendorInfo[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.vendorInfoList) {
      VendorInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.vendorInfoList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorInfoList.push(VendorInfo.decode(reader, reader.uint32()))
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
    message.vendorInfoList = []
    if (object.vendorInfoList !== undefined && object.vendorInfoList !== null) {
      for (const e of object.vendorInfoList) {
        message.vendorInfoList.push(VendorInfo.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.vendorInfoList) {
      obj.vendorInfoList = message.vendorInfoList.map((e) => (e ? VendorInfo.toJSON(e) : undefined))
    } else {
      obj.vendorInfoList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.vendorInfoList = []
    if (object.vendorInfoList !== undefined && object.vendorInfoList !== null) {
      for (const e of object.vendorInfoList) {
        message.vendorInfoList.push(VendorInfo.fromPartial(e))
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
