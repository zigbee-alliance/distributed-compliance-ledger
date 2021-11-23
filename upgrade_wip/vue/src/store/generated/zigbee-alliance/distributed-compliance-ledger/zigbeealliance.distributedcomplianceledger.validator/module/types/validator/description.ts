/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface Description {
  /** name. */
  name: string
  /** optional identity signature. */
  identity: string
  /** optional website link. */
  website: string
  /** optional details. */
  details: string
}

const baseDescription: object = { name: '', identity: '', website: '', details: '' }

export const Description = {
  encode(message: Description, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    if (message.identity !== '') {
      writer.uint32(18).string(message.identity)
    }
    if (message.website !== '') {
      writer.uint32(26).string(message.website)
    }
    if (message.details !== '') {
      writer.uint32(34).string(message.details)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Description {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseDescription } as Description
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string()
          break
        case 2:
          message.identity = reader.string()
          break
        case 3:
          message.website = reader.string()
          break
        case 4:
          message.details = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Description {
    const message = { ...baseDescription } as Description
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.identity !== undefined && object.identity !== null) {
      message.identity = String(object.identity)
    } else {
      message.identity = ''
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = String(object.website)
    } else {
      message.website = ''
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = String(object.details)
    } else {
      message.details = ''
    }
    return message
  },

  toJSON(message: Description): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    message.identity !== undefined && (obj.identity = message.identity)
    message.website !== undefined && (obj.website = message.website)
    message.details !== undefined && (obj.details = message.details)
    return obj
  },

  fromPartial(object: DeepPartial<Description>): Description {
    const message = { ...baseDescription } as Description
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.identity !== undefined && object.identity !== null) {
      message.identity = object.identity
    } else {
      message.identity = ''
    }
    if (object.website !== undefined && object.website !== null) {
      message.website = object.website
    } else {
      message.website = ''
    }
    if (object.details !== undefined && object.details !== null) {
      message.details = object.details
    } else {
      message.details = ''
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
