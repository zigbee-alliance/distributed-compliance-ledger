/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model'

export interface Product {
  pid: number
  name: string
  partNumber: string
}

const baseProduct: object = { pid: 0, name: '', partNumber: '' }

export const Product = {
  encode(message: Product, writer: Writer = Writer.create()): Writer {
    if (message.pid !== 0) {
      writer.uint32(8).int32(message.pid)
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name)
    }
    if (message.partNumber !== '') {
      writer.uint32(26).string(message.partNumber)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Product {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseProduct } as Product
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pid = reader.int32()
          break
        case 2:
          message.name = reader.string()
          break
        case 3:
          message.partNumber = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Product {
    const message = { ...baseProduct } as Product
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.partNumber !== undefined && object.partNumber !== null) {
      message.partNumber = String(object.partNumber)
    } else {
      message.partNumber = ''
    }
    return message
  },

  toJSON(message: Product): unknown {
    const obj: any = {}
    message.pid !== undefined && (obj.pid = message.pid)
    message.name !== undefined && (obj.name = message.name)
    message.partNumber !== undefined && (obj.partNumber = message.partNumber)
    return obj
  },

  fromPartial(object: DeepPartial<Product>): Product {
    const message = { ...baseProduct } as Product
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.partNumber !== undefined && object.partNumber !== null) {
      message.partNumber = object.partNumber
    } else {
      message.partNumber = ''
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
