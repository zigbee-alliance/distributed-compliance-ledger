/* eslint-disable */
import { VendorInfo } from '../vendorinfo/vendor_info'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface NewVendorInfo {
  index: string
  vendorInfo: VendorInfo | undefined
  creator: string
}

const baseNewVendorInfo: object = { index: '', creator: '' }

export const NewVendorInfo = {
  encode(message: NewVendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    if (message.vendorInfo !== undefined) {
      VendorInfo.encode(message.vendorInfo, writer.uint32(18).fork()).ldelim()
    }
    if (message.creator !== '') {
      writer.uint32(26).string(message.creator)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): NewVendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseNewVendorInfo } as NewVendorInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string()
          break
        case 2:
          message.vendorInfo = VendorInfo.decode(reader, reader.uint32())
          break
        case 3:
          message.creator = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): NewVendorInfo {
    const message = { ...baseNewVendorInfo } as NewVendorInfo
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromJSON(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    return message
  },

  toJSON(message: NewVendorInfo): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    message.vendorInfo !== undefined && (obj.vendorInfo = message.vendorInfo ? VendorInfo.toJSON(message.vendorInfo) : undefined)
    message.creator !== undefined && (obj.creator = message.creator)
    return obj
  },

  fromPartial(object: DeepPartial<NewVendorInfo>): NewVendorInfo {
    const message = { ...baseNewVendorInfo } as NewVendorInfo
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromPartial(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
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
