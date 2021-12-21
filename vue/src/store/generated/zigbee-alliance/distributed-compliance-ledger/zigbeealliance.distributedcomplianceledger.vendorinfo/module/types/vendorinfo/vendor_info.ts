/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface VendorInfo {
  vendorID: number
  vendorName: string
  companyLegalName: string
  companyPrefferedName: string
  vendorLandingPageURL: string
  creator: string
}

const baseVendorInfo: object = { vendorID: 0, vendorName: '', companyLegalName: '', companyPrefferedName: '', vendorLandingPageURL: '', creator: '' }

export const VendorInfo = {
  encode(message: VendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.vendorID !== 0) {
      writer.uint32(8).uint64(message.vendorID)
    }
    if (message.vendorName !== '') {
      writer.uint32(18).string(message.vendorName)
    }
    if (message.companyLegalName !== '') {
      writer.uint32(26).string(message.companyLegalName)
    }
    if (message.companyPrefferedName !== '') {
      writer.uint32(34).string(message.companyPrefferedName)
    }
    if (message.vendorLandingPageURL !== '') {
      writer.uint32(42).string(message.vendorLandingPageURL)
    }
    if (message.creator !== '') {
      writer.uint32(50).string(message.creator)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): VendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseVendorInfo } as VendorInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorID = longToNumber(reader.uint64() as Long)
          break
        case 2:
          message.vendorName = reader.string()
          break
        case 3:
          message.companyLegalName = reader.string()
          break
        case 4:
          message.companyPrefferedName = reader.string()
          break
        case 5:
          message.vendorLandingPageURL = reader.string()
          break
        case 6:
          message.creator = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): VendorInfo {
    const message = { ...baseVendorInfo } as VendorInfo
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    if (object.vendorName !== undefined && object.vendorName !== null) {
      message.vendorName = String(object.vendorName)
    } else {
      message.vendorName = ''
    }
    if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
      message.companyLegalName = String(object.companyLegalName)
    } else {
      message.companyLegalName = ''
    }
    if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
      message.companyPrefferedName = String(object.companyPrefferedName)
    } else {
      message.companyPrefferedName = ''
    }
    if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
      message.vendorLandingPageURL = String(object.vendorLandingPageURL)
    } else {
      message.vendorLandingPageURL = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    return message
  },

  toJSON(message: VendorInfo): unknown {
    const obj: any = {}
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    message.vendorName !== undefined && (obj.vendorName = message.vendorName)
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName)
    message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName)
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL)
    message.creator !== undefined && (obj.creator = message.creator)
    return obj
  },

  fromPartial(object: DeepPartial<VendorInfo>): VendorInfo {
    const message = { ...baseVendorInfo } as VendorInfo
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
    }
    if (object.vendorName !== undefined && object.vendorName !== null) {
      message.vendorName = object.vendorName
    } else {
      message.vendorName = ''
    }
    if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
      message.companyLegalName = object.companyLegalName
    } else {
      message.companyLegalName = ''
    }
    if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
      message.companyPrefferedName = object.companyPrefferedName
    } else {
      message.companyPrefferedName = ''
    }
    if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
      message.vendorLandingPageURL = object.vendorLandingPageURL
    } else {
      message.vendorLandingPageURL = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    return message
  }
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
