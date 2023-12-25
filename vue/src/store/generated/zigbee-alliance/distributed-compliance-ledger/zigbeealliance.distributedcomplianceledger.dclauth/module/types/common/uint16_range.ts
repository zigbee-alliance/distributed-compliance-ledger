/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.common'

export interface Uint16Range {
  min: number
  max: number
}

const baseUint16Range: object = { min: 0, max: 0 }

export const Uint16Range = {
  encode(message: Uint16Range, writer: Writer = Writer.create()): Writer {
    if (message.min !== 0) {
      writer.uint32(8).int32(message.min)
    }
    if (message.max !== 0) {
      writer.uint32(16).int32(message.max)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Uint16Range {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseUint16Range } as Uint16Range
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.min = reader.int32()
          break
        case 2:
          message.max = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Uint16Range {
    const message = { ...baseUint16Range } as Uint16Range
    if (object.min !== undefined && object.min !== null) {
      message.min = Number(object.min)
    } else {
      message.min = 0
    }
    if (object.max !== undefined && object.max !== null) {
      message.max = Number(object.max)
    } else {
      message.max = 0
    }
    return message
  },

  toJSON(message: Uint16Range): unknown {
    const obj: any = {}
    message.min !== undefined && (obj.min = message.min)
    message.max !== undefined && (obj.max = message.max)
    return obj
  },

  fromPartial(object: DeepPartial<Uint16Range>): Uint16Range {
    const message = { ...baseUint16Range } as Uint16Range
    if (object.min !== undefined && object.min !== null) {
      message.min = object.min
    } else {
      message.min = 0
    }
    if (object.max !== undefined && object.max !== null) {
      message.max = object.max
    } else {
      message.max = 0
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
