/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface ValidatorSigningInfo {
  address: string
  startHeight: number
  indexOffset: number
  missedBlocksCounter: number
}

const baseValidatorSigningInfo: object = { address: '', startHeight: 0, indexOffset: 0, missedBlocksCounter: 0 }

export const ValidatorSigningInfo = {
  encode(message: ValidatorSigningInfo, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.startHeight !== 0) {
      writer.uint32(16).uint64(message.startHeight)
    }
    if (message.indexOffset !== 0) {
      writer.uint32(24).uint64(message.indexOffset)
    }
    if (message.missedBlocksCounter !== 0) {
      writer.uint32(32).uint64(message.missedBlocksCounter)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ValidatorSigningInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseValidatorSigningInfo } as ValidatorSigningInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.startHeight = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.indexOffset = longToNumber(reader.uint64() as Long)
          break
        case 4:
          message.missedBlocksCounter = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ValidatorSigningInfo {
    const message = { ...baseValidatorSigningInfo } as ValidatorSigningInfo
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.startHeight !== undefined && object.startHeight !== null) {
      message.startHeight = Number(object.startHeight)
    } else {
      message.startHeight = 0
    }
    if (object.indexOffset !== undefined && object.indexOffset !== null) {
      message.indexOffset = Number(object.indexOffset)
    } else {
      message.indexOffset = 0
    }
    if (object.missedBlocksCounter !== undefined && object.missedBlocksCounter !== null) {
      message.missedBlocksCounter = Number(object.missedBlocksCounter)
    } else {
      message.missedBlocksCounter = 0
    }
    return message
  },

  toJSON(message: ValidatorSigningInfo): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.startHeight !== undefined && (obj.startHeight = message.startHeight)
    message.indexOffset !== undefined && (obj.indexOffset = message.indexOffset)
    message.missedBlocksCounter !== undefined && (obj.missedBlocksCounter = message.missedBlocksCounter)
    return obj
  },

  fromPartial(object: DeepPartial<ValidatorSigningInfo>): ValidatorSigningInfo {
    const message = { ...baseValidatorSigningInfo } as ValidatorSigningInfo
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.startHeight !== undefined && object.startHeight !== null) {
      message.startHeight = object.startHeight
    } else {
      message.startHeight = 0
    }
    if (object.indexOffset !== undefined && object.indexOffset !== null) {
      message.indexOffset = object.indexOffset
    } else {
      message.indexOffset = 0
    }
    if (object.missedBlocksCounter !== undefined && object.missedBlocksCounter !== null) {
      message.missedBlocksCounter = object.missedBlocksCounter
    } else {
      message.missedBlocksCounter = 0
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
