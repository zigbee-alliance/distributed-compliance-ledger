/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model'

export interface ModelVersion {
  vid: number
  pid: number
  softwareVersion: number
  softwareVersionString: string
  cdVersionNumber: number
  firmwareInformation: string
  softwareVersionValid: boolean
  otaUrl: string
  otaFileSize: number
  otaChecksum: string
  otaChecksumType: number
  minApplicableSoftwareVersion: number
  maxApplicableSoftwareVersion: number
  releaseNotesUrl: string
  creator: string
}

const baseModelVersion: object = {
  vid: 0,
  pid: 0,
  softwareVersion: 0,
  softwareVersionString: '',
  cdVersionNumber: 0,
  firmwareInformation: '',
  softwareVersionValid: false,
  otaUrl: '',
  otaFileSize: 0,
  otaChecksum: '',
  otaChecksumType: 0,
  minApplicableSoftwareVersion: 0,
  maxApplicableSoftwareVersion: 0,
  releaseNotesUrl: '',
  creator: ''
}

export const ModelVersion = {
  encode(message: ModelVersion, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion)
    }
    if (message.softwareVersionString !== '') {
      writer.uint32(34).string(message.softwareVersionString)
    }
    if (message.cdVersionNumber !== 0) {
      writer.uint32(40).int32(message.cdVersionNumber)
    }
    if (message.firmwareInformation !== '') {
      writer.uint32(50).string(message.firmwareInformation)
    }
    if (message.softwareVersionValid === true) {
      writer.uint32(56).bool(message.softwareVersionValid)
    }
    if (message.otaUrl !== '') {
      writer.uint32(66).string(message.otaUrl)
    }
    if (message.otaFileSize !== 0) {
      writer.uint32(72).uint64(message.otaFileSize)
    }
    if (message.otaChecksum !== '') {
      writer.uint32(82).string(message.otaChecksum)
    }
    if (message.otaChecksumType !== 0) {
      writer.uint32(88).int32(message.otaChecksumType)
    }
    if (message.minApplicableSoftwareVersion !== 0) {
      writer.uint32(96).uint32(message.minApplicableSoftwareVersion)
    }
    if (message.maxApplicableSoftwareVersion !== 0) {
      writer.uint32(104).uint32(message.maxApplicableSoftwareVersion)
    }
    if (message.releaseNotesUrl !== '') {
      writer.uint32(114).string(message.releaseNotesUrl)
    }
    if (message.creator !== '') {
      writer.uint32(122).string(message.creator)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ModelVersion {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseModelVersion } as ModelVersion
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        case 3:
          message.softwareVersion = reader.uint32()
          break
        case 4:
          message.softwareVersionString = reader.string()
          break
        case 5:
          message.cdVersionNumber = reader.int32()
          break
        case 6:
          message.firmwareInformation = reader.string()
          break
        case 7:
          message.softwareVersionValid = reader.bool()
          break
        case 8:
          message.otaUrl = reader.string()
          break
        case 9:
          message.otaFileSize = longToNumber(reader.uint64() as Long)
          break
        case 10:
          message.otaChecksum = reader.string()
          break
        case 11:
          message.otaChecksumType = reader.int32()
          break
        case 12:
          message.minApplicableSoftwareVersion = reader.uint32()
          break
        case 13:
          message.maxApplicableSoftwareVersion = reader.uint32()
          break
        case 14:
          message.releaseNotesUrl = reader.string()
          break
        case 15:
          message.creator = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ModelVersion {
    const message = { ...baseModelVersion } as ModelVersion
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
      message.softwareVersionString = String(object.softwareVersionString)
    } else {
      message.softwareVersionString = ''
    }
    if (object.cdVersionNumber !== undefined && object.cdVersionNumber !== null) {
      message.cdVersionNumber = Number(object.cdVersionNumber)
    } else {
      message.cdVersionNumber = 0
    }
    if (object.firmwareInformation !== undefined && object.firmwareInformation !== null) {
      message.firmwareInformation = String(object.firmwareInformation)
    } else {
      message.firmwareInformation = ''
    }
    if (object.softwareVersionValid !== undefined && object.softwareVersionValid !== null) {
      message.softwareVersionValid = Boolean(object.softwareVersionValid)
    } else {
      message.softwareVersionValid = false
    }
    if (object.otaUrl !== undefined && object.otaUrl !== null) {
      message.otaUrl = String(object.otaUrl)
    } else {
      message.otaUrl = ''
    }
    if (object.otaFileSize !== undefined && object.otaFileSize !== null) {
      message.otaFileSize = Number(object.otaFileSize)
    } else {
      message.otaFileSize = 0
    }
    if (object.otaChecksum !== undefined && object.otaChecksum !== null) {
      message.otaChecksum = String(object.otaChecksum)
    } else {
      message.otaChecksum = ''
    }
    if (object.otaChecksumType !== undefined && object.otaChecksumType !== null) {
      message.otaChecksumType = Number(object.otaChecksumType)
    } else {
      message.otaChecksumType = 0
    }
    if (object.minApplicableSoftwareVersion !== undefined && object.minApplicableSoftwareVersion !== null) {
      message.minApplicableSoftwareVersion = Number(object.minApplicableSoftwareVersion)
    } else {
      message.minApplicableSoftwareVersion = 0
    }
    if (object.maxApplicableSoftwareVersion !== undefined && object.maxApplicableSoftwareVersion !== null) {
      message.maxApplicableSoftwareVersion = Number(object.maxApplicableSoftwareVersion)
    } else {
      message.maxApplicableSoftwareVersion = 0
    }
    if (object.releaseNotesUrl !== undefined && object.releaseNotesUrl !== null) {
      message.releaseNotesUrl = String(object.releaseNotesUrl)
    } else {
      message.releaseNotesUrl = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    return message
  },

  toJSON(message: ModelVersion): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    message.softwareVersionString !== undefined && (obj.softwareVersionString = message.softwareVersionString)
    message.cdVersionNumber !== undefined && (obj.cdVersionNumber = message.cdVersionNumber)
    message.firmwareInformation !== undefined && (obj.firmwareInformation = message.firmwareInformation)
    message.softwareVersionValid !== undefined && (obj.softwareVersionValid = message.softwareVersionValid)
    message.otaUrl !== undefined && (obj.otaUrl = message.otaUrl)
    message.otaFileSize !== undefined && (obj.otaFileSize = message.otaFileSize)
    message.otaChecksum !== undefined && (obj.otaChecksum = message.otaChecksum)
    message.otaChecksumType !== undefined && (obj.otaChecksumType = message.otaChecksumType)
    message.minApplicableSoftwareVersion !== undefined && (obj.minApplicableSoftwareVersion = message.minApplicableSoftwareVersion)
    message.maxApplicableSoftwareVersion !== undefined && (obj.maxApplicableSoftwareVersion = message.maxApplicableSoftwareVersion)
    message.releaseNotesUrl !== undefined && (obj.releaseNotesUrl = message.releaseNotesUrl)
    message.creator !== undefined && (obj.creator = message.creator)
    return obj
  },

  fromPartial(object: DeepPartial<ModelVersion>): ModelVersion {
    const message = { ...baseModelVersion } as ModelVersion
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
      message.softwareVersionString = object.softwareVersionString
    } else {
      message.softwareVersionString = ''
    }
    if (object.cdVersionNumber !== undefined && object.cdVersionNumber !== null) {
      message.cdVersionNumber = object.cdVersionNumber
    } else {
      message.cdVersionNumber = 0
    }
    if (object.firmwareInformation !== undefined && object.firmwareInformation !== null) {
      message.firmwareInformation = object.firmwareInformation
    } else {
      message.firmwareInformation = ''
    }
    if (object.softwareVersionValid !== undefined && object.softwareVersionValid !== null) {
      message.softwareVersionValid = object.softwareVersionValid
    } else {
      message.softwareVersionValid = false
    }
    if (object.otaUrl !== undefined && object.otaUrl !== null) {
      message.otaUrl = object.otaUrl
    } else {
      message.otaUrl = ''
    }
    if (object.otaFileSize !== undefined && object.otaFileSize !== null) {
      message.otaFileSize = object.otaFileSize
    } else {
      message.otaFileSize = 0
    }
    if (object.otaChecksum !== undefined && object.otaChecksum !== null) {
      message.otaChecksum = object.otaChecksum
    } else {
      message.otaChecksum = ''
    }
    if (object.otaChecksumType !== undefined && object.otaChecksumType !== null) {
      message.otaChecksumType = object.otaChecksumType
    } else {
      message.otaChecksumType = 0
    }
    if (object.minApplicableSoftwareVersion !== undefined && object.minApplicableSoftwareVersion !== null) {
      message.minApplicableSoftwareVersion = object.minApplicableSoftwareVersion
    } else {
      message.minApplicableSoftwareVersion = 0
    }
    if (object.maxApplicableSoftwareVersion !== undefined && object.maxApplicableSoftwareVersion !== null) {
      message.maxApplicableSoftwareVersion = object.maxApplicableSoftwareVersion
    } else {
      message.maxApplicableSoftwareVersion = 0
    }
    if (object.releaseNotesUrl !== undefined && object.releaseNotesUrl !== null) {
      message.releaseNotesUrl = object.releaseNotesUrl
    } else {
      message.releaseNotesUrl = ''
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
