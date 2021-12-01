/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

export interface PendingAccount {
  address: string
  approvals: string[]
}

const basePendingAccount: object = { address: '', approvals: '' }

export const PendingAccount = {
  encode(message: PendingAccount, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    for (const v of message.approvals) {
      writer.uint32(18).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): PendingAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePendingAccount } as PendingAccount
    message.approvals = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
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

  fromJSON(object: any): PendingAccount {
    const message = { ...basePendingAccount } as PendingAccount
    message.approvals = []
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(String(e))
      }
    }
    return message
  },

  toJSON(message: PendingAccount): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => e)
    } else {
      obj.approvals = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<PendingAccount>): PendingAccount {
    const message = { ...basePendingAccount } as PendingAccount
    message.approvals = []
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
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
