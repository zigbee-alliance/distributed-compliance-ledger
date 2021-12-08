/* eslint-disable */
import { Product } from '../model/product'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model'

export interface VendorProducts {
  vid: number
  products: Product | undefined
}

const baseVendorProducts: object = { vid: 0 }

export const VendorProducts = {
  encode(message: VendorProducts, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.products !== undefined) {
      Product.encode(message.products, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): VendorProducts {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseVendorProducts } as VendorProducts
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.products = Product.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): VendorProducts {
    const message = { ...baseVendorProducts } as VendorProducts
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.products !== undefined && object.products !== null) {
      message.products = Product.fromJSON(object.products)
    } else {
      message.products = undefined
    }
    return message
  },

  toJSON(message: VendorProducts): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.products !== undefined && (obj.products = message.products ? Product.toJSON(message.products) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<VendorProducts>): VendorProducts {
    const message = { ...baseVendorProducts } as VendorProducts
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.products !== undefined && object.products !== null) {
      message.products = Product.fromPartial(object.products)
    } else {
      message.products = undefined
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
