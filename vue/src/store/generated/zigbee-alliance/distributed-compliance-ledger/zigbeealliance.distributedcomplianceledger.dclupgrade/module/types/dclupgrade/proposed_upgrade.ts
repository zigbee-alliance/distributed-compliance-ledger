/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade'

export interface ProposedUpgrade {
  name: string
  plan: string
  creator: string
  approvals: string[]
}

const baseProposedUpgrade: object = { name: '', plan: '', creator: '', approvals: '' }

export const ProposedUpgrade = {
  encode(message: ProposedUpgrade, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    if (message.plan !== '') {
      writer.uint32(18).string(message.plan)
    }
    if (message.creator !== '') {
      writer.uint32(26).string(message.creator)
    }
    for (const v of message.approvals) {
      writer.uint32(34).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ProposedUpgrade {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseProposedUpgrade } as ProposedUpgrade
    message.approvals = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string()
          break
        case 2:
          message.plan = reader.string()
          break
        case 3:
          message.creator = reader.string()
          break
        case 4:
          message.approvals.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ProposedUpgrade {
    const message = { ...baseProposedUpgrade } as ProposedUpgrade
    message.approvals = []
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    if (object.plan !== undefined && object.plan !== null) {
      message.plan = String(object.plan)
    } else {
      message.plan = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(String(e))
      }
    }
    return message
  },

  toJSON(message: ProposedUpgrade): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    message.plan !== undefined && (obj.plan = message.plan)
    message.creator !== undefined && (obj.creator = message.creator)
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => e)
    } else {
      obj.approvals = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<ProposedUpgrade>): ProposedUpgrade {
    const message = { ...baseProposedUpgrade } as ProposedUpgrade
    message.approvals = []
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    if (object.plan !== undefined && object.plan !== null) {
      message.plan = object.plan
    } else {
      message.plan = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
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
