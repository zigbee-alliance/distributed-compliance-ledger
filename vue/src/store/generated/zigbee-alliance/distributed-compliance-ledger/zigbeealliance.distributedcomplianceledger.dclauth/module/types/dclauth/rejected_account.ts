/* eslint-disable */
import { Account } from '../dclauth/account'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

export interface RejectedAccount {
  account: Account | undefined
}

const baseRejectedAccount: object = {}

export const RejectedAccount = {
  encode(message: RejectedAccount, writer: Writer = Writer.create()): Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RejectedAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRejectedAccount } as RejectedAccount
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RejectedAccount {
    const message = { ...baseRejectedAccount } as RejectedAccount
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account)
    } else {
      message.account = undefined
    }
    return message
  },

  toJSON(message: RejectedAccount): unknown {
    const obj: any = {}
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<RejectedAccount>): RejectedAccount {
    const message = { ...baseRejectedAccount } as RejectedAccount
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromPartial(object.account)
    } else {
      message.account = undefined
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
