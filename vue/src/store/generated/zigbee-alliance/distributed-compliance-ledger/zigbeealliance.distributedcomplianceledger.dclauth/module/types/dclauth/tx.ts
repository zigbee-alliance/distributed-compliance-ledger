/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

export interface MsgProposeAddAccount {
  signer: string
  address: string
  pubKey: string
  roles: string[]
  vendorID: number
}

export interface MsgProposeAddAccountResponse {}

export interface MsgApproveAddAccount {
  signer: string
  address: string
}

export interface MsgApproveAddAccountResponse {}

export interface MsgProposeRevokeAccount {
  signer: string
  address: string
}

export interface MsgProposeRevokeAccountResponse {}

export interface MsgApproveRevokeAccount {
  signer: string
  address: string
}

export interface MsgApproveRevokeAccountResponse {}

const baseMsgProposeAddAccount: object = { signer: '', address: '', pubKey: '', roles: '', vendorID: 0 }

export const MsgProposeAddAccount = {
  encode(message: MsgProposeAddAccount, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    if (message.pubKey !== '') {
      writer.uint32(26).string(message.pubKey)
    }
    for (const v of message.roles) {
      writer.uint32(34).string(v!)
    }
    if (message.vendorID !== 0) {
      writer.uint32(40).uint64(message.vendorID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeAddAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeAddAccount } as MsgProposeAddAccount
    message.roles = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        case 3:
          message.pubKey = reader.string()
          break
        case 4:
          message.roles.push(reader.string())
          break
        case 5:
          message.vendorID = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgProposeAddAccount {
    const message = { ...baseMsgProposeAddAccount } as MsgProposeAddAccount
    message.roles = []
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = String(object.pubKey)
    } else {
      message.pubKey = ''
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

  toJSON(message: MsgProposeAddAccount): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.address !== undefined && (obj.address = message.address)
    message.pubKey !== undefined && (obj.pubKey = message.pubKey)
    if (message.roles) {
      obj.roles = message.roles.map((e) => e)
    } else {
      obj.roles = []
    }
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    return obj
  },

  fromPartial(object: DeepPartial<MsgProposeAddAccount>): MsgProposeAddAccount {
    const message = { ...baseMsgProposeAddAccount } as MsgProposeAddAccount
    message.roles = []
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = object.pubKey
    } else {
      message.pubKey = ''
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

const baseMsgProposeAddAccountResponse: object = {}

export const MsgProposeAddAccountResponse = {
  encode(_: MsgProposeAddAccountResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeAddAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeAddAccountResponse } as MsgProposeAddAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgProposeAddAccountResponse {
    const message = { ...baseMsgProposeAddAccountResponse } as MsgProposeAddAccountResponse
    return message
  },

  toJSON(_: MsgProposeAddAccountResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgProposeAddAccountResponse>): MsgProposeAddAccountResponse {
    const message = { ...baseMsgProposeAddAccountResponse } as MsgProposeAddAccountResponse
    return message
  }
}

const baseMsgApproveAddAccount: object = { signer: '', address: '' }

export const MsgApproveAddAccount = {
  encode(message: MsgApproveAddAccount, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveAddAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveAddAccount } as MsgApproveAddAccount
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgApproveAddAccount {
    const message = { ...baseMsgApproveAddAccount } as MsgApproveAddAccount
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgApproveAddAccount): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgApproveAddAccount>): MsgApproveAddAccount {
    const message = { ...baseMsgApproveAddAccount } as MsgApproveAddAccount
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgApproveAddAccountResponse: object = {}

export const MsgApproveAddAccountResponse = {
  encode(_: MsgApproveAddAccountResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveAddAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveAddAccountResponse } as MsgApproveAddAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgApproveAddAccountResponse {
    const message = { ...baseMsgApproveAddAccountResponse } as MsgApproveAddAccountResponse
    return message
  },

  toJSON(_: MsgApproveAddAccountResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgApproveAddAccountResponse>): MsgApproveAddAccountResponse {
    const message = { ...baseMsgApproveAddAccountResponse } as MsgApproveAddAccountResponse
    return message
  }
}

const baseMsgProposeRevokeAccount: object = { signer: '', address: '' }

export const MsgProposeRevokeAccount = {
  encode(message: MsgProposeRevokeAccount, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeRevokeAccount } as MsgProposeRevokeAccount
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgProposeRevokeAccount {
    const message = { ...baseMsgProposeRevokeAccount } as MsgProposeRevokeAccount
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgProposeRevokeAccount): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgProposeRevokeAccount>): MsgProposeRevokeAccount {
    const message = { ...baseMsgProposeRevokeAccount } as MsgProposeRevokeAccount
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgProposeRevokeAccountResponse: object = {}

export const MsgProposeRevokeAccountResponse = {
  encode(_: MsgProposeRevokeAccountResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeRevokeAccountResponse } as MsgProposeRevokeAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgProposeRevokeAccountResponse {
    const message = { ...baseMsgProposeRevokeAccountResponse } as MsgProposeRevokeAccountResponse
    return message
  },

  toJSON(_: MsgProposeRevokeAccountResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgProposeRevokeAccountResponse>): MsgProposeRevokeAccountResponse {
    const message = { ...baseMsgProposeRevokeAccountResponse } as MsgProposeRevokeAccountResponse
    return message
  }
}

const baseMsgApproveRevokeAccount: object = { signer: '', address: '' }

export const MsgApproveRevokeAccount = {
  encode(message: MsgApproveRevokeAccount, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveRevokeAccount } as MsgApproveRevokeAccount
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgApproveRevokeAccount {
    const message = { ...baseMsgApproveRevokeAccount } as MsgApproveRevokeAccount
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: MsgApproveRevokeAccount): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<MsgApproveRevokeAccount>): MsgApproveRevokeAccount {
    const message = { ...baseMsgApproveRevokeAccount } as MsgApproveRevokeAccount
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseMsgApproveRevokeAccountResponse: object = {}

export const MsgApproveRevokeAccountResponse = {
  encode(_: MsgApproveRevokeAccountResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveRevokeAccountResponse } as MsgApproveRevokeAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgApproveRevokeAccountResponse {
    const message = { ...baseMsgApproveRevokeAccountResponse } as MsgApproveRevokeAccountResponse
    return message
  },

  toJSON(_: MsgApproveRevokeAccountResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgApproveRevokeAccountResponse>): MsgApproveRevokeAccountResponse {
    const message = { ...baseMsgApproveRevokeAccountResponse } as MsgApproveRevokeAccountResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  ProposeAddAccount(request: MsgProposeAddAccount): Promise<MsgProposeAddAccountResponse>
  ApproveAddAccount(request: MsgApproveAddAccount): Promise<MsgApproveAddAccountResponse>
  ProposeRevokeAccount(request: MsgProposeRevokeAccount): Promise<MsgProposeRevokeAccountResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ApproveRevokeAccount(request: MsgApproveRevokeAccount): Promise<MsgApproveRevokeAccountResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ProposeAddAccount(request: MsgProposeAddAccount): Promise<MsgProposeAddAccountResponse> {
    const data = MsgProposeAddAccount.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ProposeAddAccount', data)
    return promise.then((data) => MsgProposeAddAccountResponse.decode(new Reader(data)))
  }

  ApproveAddAccount(request: MsgApproveAddAccount): Promise<MsgApproveAddAccountResponse> {
    const data = MsgApproveAddAccount.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ApproveAddAccount', data)
    return promise.then((data) => MsgApproveAddAccountResponse.decode(new Reader(data)))
  }

  ProposeRevokeAccount(request: MsgProposeRevokeAccount): Promise<MsgProposeRevokeAccountResponse> {
    const data = MsgProposeRevokeAccount.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ProposeRevokeAccount', data)
    return promise.then((data) => MsgProposeRevokeAccountResponse.decode(new Reader(data)))
  }

  ApproveRevokeAccount(request: MsgApproveRevokeAccount): Promise<MsgApproveRevokeAccountResponse> {
    const data = MsgApproveRevokeAccount.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Msg', 'ApproveRevokeAccount', data)
    return promise.then((data) => MsgApproveRevokeAccountResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
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
