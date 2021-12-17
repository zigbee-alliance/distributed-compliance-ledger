/* eslint-disable */
import * as Long from 'long'
import { util, configure, Writer, Reader } from 'protobufjs/minimal'
import { BaseAccount } from '../cosmos/auth/v1beta1/auth'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

export interface Account {
  baseAccount: BaseAccount | undefined
  /**
   * NOTE. we do not user AccountRoles casting here to preserve repeated form
   *       so protobuf takes care about repeated items in generated code,
   *       (but that might be not the final solution)
   */
  roles: string[]
  vendorID: number
}

const baseAccount: object = { roles: '', vendorID: 0 }

export const Account = {
  encode(message: Account, writer: Writer = Writer.create()): Writer {
    if (message.baseAccount !== undefined) {
      BaseAccount.encode(message.baseAccount, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.roles) {
      writer.uint32(18).string(v!)
    }
    if (message.vendorID !== 0) {
      writer.uint32(24).uint64(message.vendorID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): Account {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseAccount } as Account
    message.roles = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.baseAccount = BaseAccount.decode(reader, reader.uint32())
          break
        case 2:
          message.roles.push(reader.string())
          break
        case 3:
          message.vendorID = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): Account {
    const message = { ...baseAccount } as Account
    message.roles = []
    if (object.baseAccount !== undefined && object.baseAccount !== null) {
      message.baseAccount = BaseAccount.fromJSON(object.baseAccount)
    } else {
      message.baseAccount = undefined
    }
    if (object.roles !== undefined && object.roles !== null) {
      for (const e of object.roles) {
        message.roles.push(String(e))
      }
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    return message
  },

  toJSON(message: Account): unknown {
    const obj: any = {}
    message.baseAccount !== undefined && (obj.baseAccount = message.baseAccount ? BaseAccount.toJSON(message.baseAccount) : undefined)
    if (message.roles) {
      obj.roles = message.roles.map((e) => e)
    } else {
      obj.roles = []
    }
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    return obj
  },

  fromPartial(object: DeepPartial<Account>): Account {
    const message = { ...baseAccount } as Account
    message.roles = []
    if (object.baseAccount !== undefined && object.baseAccount !== null) {
      message.baseAccount = BaseAccount.fromPartial(object.baseAccount)
    } else {
      message.baseAccount = undefined
    }
    if (object.roles !== undefined && object.roles !== null) {
      for (const e of object.roles) {
        message.roles.push(e)
      }
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
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
