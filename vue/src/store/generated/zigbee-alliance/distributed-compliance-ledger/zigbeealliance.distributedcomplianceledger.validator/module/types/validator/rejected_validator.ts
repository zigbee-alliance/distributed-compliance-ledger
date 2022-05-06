/* eslint-disable */
import { Grant } from '../validator/grant'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface RejectedDisableValidator {
  address: string
  creator: string
  approvals: Grant[]
  rejects: Grant[]
}

const baseRejectedDisableValidator: object = { address: '', creator: '' }

export const RejectedDisableValidator = {
  encode(message: RejectedDisableValidator, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.creator !== '') {
      writer.uint32(18).string(message.creator)
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.rejects) {
      Grant.encode(v!, writer.uint32(34).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RejectedDisableValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRejectedDisableValidator } as RejectedDisableValidator
    message.approvals = []
    message.rejects = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.creator = reader.string()
          break
        case 3:
          message.approvals.push(Grant.decode(reader, reader.uint32()))
          break
        case 4:
          message.rejects.push(Grant.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RejectedDisableValidator {
    const message = { ...baseRejectedDisableValidator } as RejectedDisableValidator
    message.approvals = []
    message.rejects = []
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromJSON(e))
      }
    }
    if (object.rejects !== undefined && object.rejects !== null) {
      for (const e of object.rejects) {
        message.rejects.push(Grant.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: RejectedDisableValidator): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.creator !== undefined && (obj.creator = message.creator)
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => (e ? Grant.toJSON(e) : undefined))
    } else {
      obj.approvals = []
    }
    if (message.rejects) {
      obj.rejects = message.rejects.map((e) => (e ? Grant.toJSON(e) : undefined))
    } else {
      obj.rejects = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<RejectedDisableValidator>): RejectedDisableValidator {
    const message = { ...baseRejectedDisableValidator } as RejectedDisableValidator
    message.approvals = []
    message.rejects = []
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromPartial(e))
      }
    }
    if (object.rejects !== undefined && object.rejects !== null) {
      for (const e of object.rejects) {
        message.rejects.push(Grant.fromPartial(e))
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
