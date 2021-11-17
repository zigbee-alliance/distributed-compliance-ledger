/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface ValidatorMissedBlockBitArray {
  address: string
  index: number
}

const baseValidatorMissedBlockBitArray: object = { address: '', index: 0 }

export const ValidatorMissedBlockBitArray = {
  encode(message: ValidatorMissedBlockBitArray, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ValidatorMissedBlockBitArray {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseValidatorMissedBlockBitArray } as ValidatorMissedBlockBitArray
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.index = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ValidatorMissedBlockBitArray {
    const message = { ...baseValidatorMissedBlockBitArray } as ValidatorMissedBlockBitArray
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = Number(object.index)
    } else {
      message.index = 0
    }
    return message
  },

  toJSON(message: ValidatorMissedBlockBitArray): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<ValidatorMissedBlockBitArray>): ValidatorMissedBlockBitArray {
    const message = { ...baseValidatorMissedBlockBitArray } as ValidatorMissedBlockBitArray
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = 0
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
