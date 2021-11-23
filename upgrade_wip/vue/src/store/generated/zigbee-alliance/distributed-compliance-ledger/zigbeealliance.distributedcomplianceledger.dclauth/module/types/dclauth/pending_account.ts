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
  address: Account | undefined
  approvals: string[]
}

const basePendingAccount: object = { approvals: '' }

export const PendingAccount = {
  encode(message: PendingAccount, writer: Writer = Writer.create()): Writer {
    if (message.address !== undefined) {
      Account.encode(message.address, writer.uint32(10).fork()).ldelim()
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
          message.address = Account.decode(reader, reader.uint32())
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
      message.address = Account.fromJSON(object.address)
    } else {
      message.address = undefined
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
    message.address !== undefined && (obj.address = message.address ? Account.toJSON(message.address) : undefined)
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
      message.address = Account.fromPartial(object.address)
    } else {
      message.address = undefined
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
