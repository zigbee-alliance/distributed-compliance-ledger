/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

export interface ComplianceHistoryItem {
  softwareVersionCertificationStatus: number
  date: string
  reason: string
}

const baseComplianceHistoryItem: object = { softwareVersionCertificationStatus: 0, date: '', reason: '' }

export const ComplianceHistoryItem = {
  encode(message: ComplianceHistoryItem, writer: Writer = Writer.create()): Writer {
    if (message.softwareVersionCertificationStatus !== 0) {
      writer.uint32(8).uint64(message.softwareVersionCertificationStatus)
    }
    if (message.date !== '') {
      writer.uint32(18).string(message.date)
    }
    if (message.reason !== '') {
      writer.uint32(26).string(message.reason)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ComplianceHistoryItem {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseComplianceHistoryItem } as ComplianceHistoryItem
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.softwareVersionCertificationStatus = longToNumber(reader.uint64() as Long)
          break
        case 2:
          message.date = reader.string()
          break
        case 3:
          message.reason = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ComplianceHistoryItem {
    const message = { ...baseComplianceHistoryItem } as ComplianceHistoryItem
    if (object.softwareVersionCertificationStatus !== undefined && object.softwareVersionCertificationStatus !== null) {
      message.softwareVersionCertificationStatus = Number(object.softwareVersionCertificationStatus)
    } else {
      message.softwareVersionCertificationStatus = 0
    }
    if (object.date !== undefined && object.date !== null) {
      message.date = String(object.date)
    } else {
      message.date = ''
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = String(object.reason)
    } else {
      message.reason = ''
    }
    return message
  },

  toJSON(message: ComplianceHistoryItem): unknown {
    const obj: any = {}
    message.softwareVersionCertificationStatus !== undefined && (obj.softwareVersionCertificationStatus = message.softwareVersionCertificationStatus)
    message.date !== undefined && (obj.date = message.date)
    message.reason !== undefined && (obj.reason = message.reason)
    return obj
  },

  fromPartial(object: DeepPartial<ComplianceHistoryItem>): ComplianceHistoryItem {
    const message = { ...baseComplianceHistoryItem } as ComplianceHistoryItem
    if (object.softwareVersionCertificationStatus !== undefined && object.softwareVersionCertificationStatus !== null) {
      message.softwareVersionCertificationStatus = object.softwareVersionCertificationStatus
    } else {
      message.softwareVersionCertificationStatus = 0
    }
    if (object.date !== undefined && object.date !== null) {
      message.date = object.date
    } else {
      message.date = ''
    }
    if (object.reason !== undefined && object.reason !== null) {
      message.reason = object.reason
    } else {
      message.reason = ''
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
