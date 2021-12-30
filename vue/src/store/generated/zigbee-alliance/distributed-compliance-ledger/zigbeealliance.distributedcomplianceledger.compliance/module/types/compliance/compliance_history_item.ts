/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

export interface ComplianceHistoryItem {
  softwareVersionCertificationStatus: number
  date: string
  reason: string
  cDVersionNumber: number
}

const baseComplianceHistoryItem: object = { softwareVersionCertificationStatus: 0, date: '', reason: '', cDVersionNumber: 0 }

export const ComplianceHistoryItem = {
  encode(message: ComplianceHistoryItem, writer: Writer = Writer.create()): Writer {
    if (message.softwareVersionCertificationStatus !== 0) {
      writer.uint32(8).uint32(message.softwareVersionCertificationStatus)
    }
    if (message.date !== '') {
      writer.uint32(18).string(message.date)
    }
    if (message.reason !== '') {
      writer.uint32(26).string(message.reason)
    }
    if (message.cDVersionNumber !== 0) {
      writer.uint32(32).uint32(message.cDVersionNumber)
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
          message.softwareVersionCertificationStatus = reader.uint32()
          break
        case 2:
          message.date = reader.string()
          break
        case 3:
          message.reason = reader.string()
          break
        case 4:
          message.cDVersionNumber = reader.uint32()
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
    if (object.cDVersionNumber !== undefined && object.cDVersionNumber !== null) {
      message.cDVersionNumber = Number(object.cDVersionNumber)
    } else {
      message.cDVersionNumber = 0
    }
    return message
  },

  toJSON(message: ComplianceHistoryItem): unknown {
    const obj: any = {}
    message.softwareVersionCertificationStatus !== undefined && (obj.softwareVersionCertificationStatus = message.softwareVersionCertificationStatus)
    message.date !== undefined && (obj.date = message.date)
    message.reason !== undefined && (obj.reason = message.reason)
    message.cDVersionNumber !== undefined && (obj.cDVersionNumber = message.cDVersionNumber)
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
    if (object.cDVersionNumber !== undefined && object.cDVersionNumber !== null) {
      message.cDVersionNumber = object.cDVersionNumber
    } else {
      message.cDVersionNumber = 0
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
