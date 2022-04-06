/* eslint-disable */
import { Account } from '../dclauth/account'
import { Grant } from '../dclauth/grant'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

export interface RevokedAccount {
  account: Account | undefined
  revokeApprovals: Grant[]
  revokedReason: string
}

const baseRevokedAccount: object = { revokedReason: '' }

export const RevokedAccount = {
  encode(message: RevokedAccount, writer: Writer = Writer.create()): Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.revokeApprovals) {
      Grant.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    if (message.revokedReason !== '') {
      writer.uint32(26).string(message.revokedReason)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RevokedAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRevokedAccount } as RevokedAccount
    message.revokeApprovals = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32())
          break
        case 2:
          message.revokeApprovals.push(Grant.decode(reader, reader.uint32()))
          break
        case 3:
          message.revokedReason = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RevokedAccount {
    const message = { ...baseRevokedAccount } as RevokedAccount
    message.revokeApprovals = []
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account)
    } else {
      message.account = undefined
    }
    if (object.revokeApprovals !== undefined && object.revokeApprovals !== null) {
      for (const e of object.revokeApprovals) {
        message.revokeApprovals.push(Grant.fromJSON(e))
      }
    }
    if (object.revokedReason !== undefined && object.revokedReason !== null) {
      message.revokedReason = String(object.revokedReason)
    } else {
      message.revokedReason = ''
    }
    return message
  },

  toJSON(message: RevokedAccount): unknown {
    const obj: any = {}
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined)
    if (message.revokeApprovals) {
      obj.revokeApprovals = message.revokeApprovals.map((e) => (e ? Grant.toJSON(e) : undefined))
    } else {
      obj.revokeApprovals = []
    }
    message.revokedReason !== undefined && (obj.revokedReason = message.revokedReason)
    return obj
  },

  fromPartial(object: DeepPartial<RevokedAccount>): RevokedAccount {
    const message = { ...baseRevokedAccount } as RevokedAccount
    message.revokeApprovals = []
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromPartial(object.account)
    } else {
      message.account = undefined
    }
    if (object.revokeApprovals !== undefined && object.revokeApprovals !== null) {
      for (const e of object.revokeApprovals) {
        message.revokeApprovals.push(Grant.fromPartial(e))
      }
    }
    if (object.revokedReason !== undefined && object.revokedReason !== null) {
      message.revokedReason = object.revokedReason
    } else {
      message.revokedReason = ''
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
