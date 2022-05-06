/* eslint-disable */
import { Account } from '../dclauth/account'
import { PendingAccount } from '../dclauth/pending_account'
import { PendingAccountRevocation } from '../dclauth/pending_account_revocation'
import { AccountStat } from '../dclauth/account_stat'
import { RevokedAccount } from '../dclauth/revoked_account'
import { RejectedAccount } from '../dclauth/rejected_account'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

/** GenesisState defines the dclauth module's genesis state. */
export interface GenesisState {
  accountList: Account[]
  pendingAccountList: PendingAccount[]
  pendingAccountRevocationList: PendingAccountRevocation[]
  accountStat: AccountStat | undefined
  revokedAccountList: RevokedAccount[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  rejectedAccountList: RejectedAccount[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.accountList) {
      Account.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.pendingAccountList) {
      PendingAccount.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    for (const v of message.pendingAccountRevocationList) {
      PendingAccountRevocation.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    if (message.accountStat !== undefined) {
      AccountStat.encode(message.accountStat, writer.uint32(34).fork()).ldelim()
    }
    for (const v of message.revokedAccountList) {
      RevokedAccount.encode(v!, writer.uint32(42).fork()).ldelim()
    }
    for (const v of message.rejectedAccountList) {
      RejectedAccount.encode(v!, writer.uint32(50).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.accountList = []
    message.pendingAccountList = []
    message.pendingAccountRevocationList = []
    message.revokedAccountList = []
    message.rejectedAccountList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.accountList.push(Account.decode(reader, reader.uint32()))
          break
        case 2:
          message.pendingAccountList.push(PendingAccount.decode(reader, reader.uint32()))
          break
        case 3:
          message.pendingAccountRevocationList.push(PendingAccountRevocation.decode(reader, reader.uint32()))
          break
        case 4:
          message.accountStat = AccountStat.decode(reader, reader.uint32())
          break
        case 5:
          message.revokedAccountList.push(RevokedAccount.decode(reader, reader.uint32()))
          break
        case 6:
          message.rejectedAccountList.push(RejectedAccount.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.accountList = []
    message.pendingAccountList = []
    message.pendingAccountRevocationList = []
    message.revokedAccountList = []
    message.rejectedAccountList = []
    if (object.accountList !== undefined && object.accountList !== null) {
      for (const e of object.accountList) {
        message.accountList.push(Account.fromJSON(e))
      }
    }
    if (object.pendingAccountList !== undefined && object.pendingAccountList !== null) {
      for (const e of object.pendingAccountList) {
        message.pendingAccountList.push(PendingAccount.fromJSON(e))
      }
    }
    if (object.pendingAccountRevocationList !== undefined && object.pendingAccountRevocationList !== null) {
      for (const e of object.pendingAccountRevocationList) {
        message.pendingAccountRevocationList.push(PendingAccountRevocation.fromJSON(e))
      }
    }
    if (object.accountStat !== undefined && object.accountStat !== null) {
      message.accountStat = AccountStat.fromJSON(object.accountStat)
    } else {
      message.accountStat = undefined
    }
    if (object.revokedAccountList !== undefined && object.revokedAccountList !== null) {
      for (const e of object.revokedAccountList) {
        message.revokedAccountList.push(RevokedAccount.fromJSON(e))
      }
    }
    if (object.rejectedAccountList !== undefined && object.rejectedAccountList !== null) {
      for (const e of object.rejectedAccountList) {
        message.rejectedAccountList.push(RejectedAccount.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.accountList) {
      obj.accountList = message.accountList.map((e) => (e ? Account.toJSON(e) : undefined))
    } else {
      obj.accountList = []
    }
    if (message.pendingAccountList) {
      obj.pendingAccountList = message.pendingAccountList.map((e) => (e ? PendingAccount.toJSON(e) : undefined))
    } else {
      obj.pendingAccountList = []
    }
    if (message.pendingAccountRevocationList) {
      obj.pendingAccountRevocationList = message.pendingAccountRevocationList.map((e) => (e ? PendingAccountRevocation.toJSON(e) : undefined))
    } else {
      obj.pendingAccountRevocationList = []
    }
    message.accountStat !== undefined && (obj.accountStat = message.accountStat ? AccountStat.toJSON(message.accountStat) : undefined)
    if (message.revokedAccountList) {
      obj.revokedAccountList = message.revokedAccountList.map((e) => (e ? RevokedAccount.toJSON(e) : undefined))
    } else {
      obj.revokedAccountList = []
    }
    if (message.rejectedAccountList) {
      obj.rejectedAccountList = message.rejectedAccountList.map((e) => (e ? RejectedAccount.toJSON(e) : undefined))
    } else {
      obj.rejectedAccountList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.accountList = []
    message.pendingAccountList = []
    message.pendingAccountRevocationList = []
    message.revokedAccountList = []
    message.rejectedAccountList = []
    if (object.accountList !== undefined && object.accountList !== null) {
      for (const e of object.accountList) {
        message.accountList.push(Account.fromPartial(e))
      }
    }
    if (object.pendingAccountList !== undefined && object.pendingAccountList !== null) {
      for (const e of object.pendingAccountList) {
        message.pendingAccountList.push(PendingAccount.fromPartial(e))
      }
    }
    if (object.pendingAccountRevocationList !== undefined && object.pendingAccountRevocationList !== null) {
      for (const e of object.pendingAccountRevocationList) {
        message.pendingAccountRevocationList.push(PendingAccountRevocation.fromPartial(e))
      }
    }
    if (object.accountStat !== undefined && object.accountStat !== null) {
      message.accountStat = AccountStat.fromPartial(object.accountStat)
    } else {
      message.accountStat = undefined
    }
    if (object.revokedAccountList !== undefined && object.revokedAccountList !== null) {
      for (const e of object.revokedAccountList) {
        message.revokedAccountList.push(RevokedAccount.fromPartial(e))
      }
    }
    if (object.rejectedAccountList !== undefined && object.rejectedAccountList !== null) {
      for (const e of object.rejectedAccountList) {
        message.rejectedAccountList.push(RejectedAccount.fromPartial(e))
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
