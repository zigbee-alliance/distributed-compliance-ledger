/* eslint-disable */
import { Account } from '../dclauth/account'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

/**
 * TODO issue 99: do we need that ???
 * option (gogoproto.goproto_getters)  = false;
 * option (gogoproto.goproto_stringer) = false;
 */
export interface PendingAccount {
  account: Account | undefined
}

const basePendingAccount: object = {}

export const PendingAccount = {
  encode(message: PendingAccount, writer: Writer = Writer.create()): Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): PendingAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePendingAccount } as PendingAccount
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

  fromJSON(object: any): PendingAccount {
    const message = { ...basePendingAccount } as PendingAccount
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account)
    } else {
      message.account = undefined
    }
    return message
  },

  toJSON(message: PendingAccount): unknown {
    const obj: any = {}
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<PendingAccount>): PendingAccount {
    const message = { ...basePendingAccount } as PendingAccount
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
