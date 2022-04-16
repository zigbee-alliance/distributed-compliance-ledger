/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface RejectedNode {
  owner: string
  approvals: string[]
}

const baseRejectedNode: object = { owner: '', approvals: '' }

export const RejectedNode = {
  encode(message: RejectedNode, writer: Writer = Writer.create()): Writer {
    if (message.owner !== '') {
      writer.uint32(10).string(message.owner)
    }
    for (const v of message.approvals) {
      writer.uint32(18).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RejectedNode {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRejectedNode } as RejectedNode
    message.approvals = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string()
          break
        case 2:
          message.approvals.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RejectedNode {
    const message = { ...baseRejectedNode } as RejectedNode
    message.approvals = []
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(String(e))
      }
    }
    return message
  },

  toJSON(message: RejectedNode): unknown {
    const obj: any = {}
    message.owner !== undefined && (obj.owner = message.owner)
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => e)
    } else {
      obj.approvals = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<RejectedNode>): RejectedNode {
    const message = { ...baseRejectedNode } as RejectedNode
    message.approvals = []
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(e)
      }
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
