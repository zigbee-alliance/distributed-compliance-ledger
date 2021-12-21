/* eslint-disable */
import { VendorInfoType } from '../vendorinfo/vendor_info_type'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

/** GenesisState defines the vendorinfo module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  vendorInfoTypeList: VendorInfoType[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.vendorInfoTypeList) {
      VendorInfoType.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.vendorInfoTypeList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorInfoTypeList.push(VendorInfoType.decode(reader, reader.uint32()))
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
    message.vendorInfoTypeList = []
    if (object.vendorInfoTypeList !== undefined && object.vendorInfoTypeList !== null) {
      for (const e of object.vendorInfoTypeList) {
        message.vendorInfoTypeList.push(VendorInfoType.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.vendorInfoTypeList) {
      obj.vendorInfoTypeList = message.vendorInfoTypeList.map((e) => (e ? VendorInfoType.toJSON(e) : undefined))
    } else {
      obj.vendorInfoTypeList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.vendorInfoTypeList = []
    if (object.vendorInfoTypeList !== undefined && object.vendorInfoTypeList !== null) {
      for (const e of object.vendorInfoTypeList) {
        message.vendorInfoTypeList.push(VendorInfoType.fromPartial(e))
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
