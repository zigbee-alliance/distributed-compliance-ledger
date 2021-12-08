/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model'

export interface ModelVersions {
  vid: number
  pid: number
  softwareVersions: number[]
  creator: string
}

const baseModelVersions: object = { vid: 0, pid: 0, softwareVersions: 0, creator: '' }

export const ModelVersions = {
  encode(message: ModelVersions, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    writer.uint32(26).fork()
    for (const v of message.softwareVersions) {
      writer.uint64(v)
    }
    writer.ldelim()
    if (message.creator !== '') {
      writer.uint32(34).string(message.creator)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ModelVersions {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseModelVersions } as ModelVersions
    message.softwareVersions = []
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
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos
            while (reader.pos < end2) {
              message.softwareVersions.push(longToNumber(reader.uint64() as Long))
            }
          } else {
            message.softwareVersions.push(longToNumber(reader.uint64() as Long))
          }
          break
        case 4:
          message.creator = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ModelVersions {
    const message = { ...baseModelVersions } as ModelVersions
    message.softwareVersions = []
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
    if (object.softwareVersions !== undefined && object.softwareVersions !== null) {
      for (const e of object.softwareVersions) {
        message.softwareVersions.push(Number(e))
      }
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    return message
  },

  toJSON(message: ModelVersions): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    if (message.softwareVersions) {
      obj.softwareVersions = message.softwareVersions.map((e) => e)
    } else {
      obj.softwareVersions = []
    }
    message.creator !== undefined && (obj.creator = message.creator)
    return obj
  },

  fromPartial(object: DeepPartial<ModelVersions>): ModelVersions {
    const message = { ...baseModelVersions } as ModelVersions
    message.softwareVersions = []
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
    if (object.softwareVersions !== undefined && object.softwareVersions !== null) {
      for (const e of object.softwareVersions) {
        message.softwareVersions.push(e)
      }
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
