/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Account } from '../dclauth/account'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { PendingAccount } from '../dclauth/pending_account'
import { PendingAccountRevocation } from '../dclauth/pending_account_revocation'
import { AccountStat } from '../dclauth/account_stat'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth'

export interface QueryGetAccountRequest {
  address: string
}

export interface QueryGetAccountResponse {
  account: Account | undefined
}

export interface QueryAllAccountRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllAccountResponse {
  account: Account[]
  pagination: PageResponse | undefined
}

export interface QueryGetPendingAccountRequest {
  address: string
}

export interface QueryGetPendingAccountResponse {
  pendingAccount: PendingAccount | undefined
}

export interface QueryAllPendingAccountRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllPendingAccountResponse {
  pendingAccount: PendingAccount[]
  pagination: PageResponse | undefined
}

export interface QueryGetPendingAccountRevocationRequest {
  address: string
}

export interface QueryGetPendingAccountRevocationResponse {
  pendingAccountRevocation: PendingAccountRevocation | undefined
}

export interface QueryAllPendingAccountRevocationRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllPendingAccountRevocationResponse {
  pendingAccountRevocation: PendingAccountRevocation[]
  pagination: PageResponse | undefined
}

export interface QueryGetAccountStatRequest {}

export interface QueryGetAccountStatResponse {
  AccountStat: AccountStat | undefined
}

const baseQueryGetAccountRequest: object = { address: '' }

export const QueryGetAccountRequest = {
  encode(message: QueryGetAccountRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAccountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAccountRequest } as QueryGetAccountRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetAccountRequest {
    const message = { ...baseQueryGetAccountRequest } as QueryGetAccountRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetAccountRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetAccountRequest>): QueryGetAccountRequest {
    const message = { ...baseQueryGetAccountRequest } as QueryGetAccountRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetAccountResponse: object = {}

export const QueryGetAccountResponse = {
  encode(message: QueryGetAccountResponse, writer: Writer = Writer.create()): Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAccountResponse } as QueryGetAccountResponse
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

  fromJSON(object: any): QueryGetAccountResponse {
    const message = { ...baseQueryGetAccountResponse } as QueryGetAccountResponse
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromJSON(object.account)
    } else {
      message.account = undefined
    }
    return message
  },

  toJSON(message: QueryGetAccountResponse): unknown {
    const obj: any = {}
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetAccountResponse>): QueryGetAccountResponse {
    const message = { ...baseQueryGetAccountResponse } as QueryGetAccountResponse
    if (object.account !== undefined && object.account !== null) {
      message.account = Account.fromPartial(object.account)
    } else {
      message.account = undefined
    }
    return message
  }
}

const baseQueryAllAccountRequest: object = {}

export const QueryAllAccountRequest = {
  encode(message: QueryAllAccountRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllAccountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllAccountRequest } as QueryAllAccountRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllAccountRequest {
    const message = { ...baseQueryAllAccountRequest } as QueryAllAccountRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllAccountRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllAccountRequest>): QueryAllAccountRequest {
    const message = { ...baseQueryAllAccountRequest } as QueryAllAccountRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllAccountResponse: object = {}

export const QueryAllAccountResponse = {
  encode(message: QueryAllAccountResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.account) {
      Account.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllAccountResponse } as QueryAllAccountResponse
    message.account = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.account.push(Account.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllAccountResponse {
    const message = { ...baseQueryAllAccountResponse } as QueryAllAccountResponse
    message.account = []
    if (object.account !== undefined && object.account !== null) {
      for (const e of object.account) {
        message.account.push(Account.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllAccountResponse): unknown {
    const obj: any = {}
    if (message.account) {
      obj.account = message.account.map((e) => (e ? Account.toJSON(e) : undefined))
    } else {
      obj.account = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllAccountResponse>): QueryAllAccountResponse {
    const message = { ...baseQueryAllAccountResponse } as QueryAllAccountResponse
    message.account = []
    if (object.account !== undefined && object.account !== null) {
      for (const e of object.account) {
        message.account.push(Account.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetPendingAccountRequest: object = { address: '' }

export const QueryGetPendingAccountRequest = {
  encode(message: QueryGetPendingAccountRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPendingAccountRequest } as QueryGetPendingAccountRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPendingAccountRequest {
    const message = { ...baseQueryGetPendingAccountRequest } as QueryGetPendingAccountRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetPendingAccountRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPendingAccountRequest>): QueryGetPendingAccountRequest {
    const message = { ...baseQueryGetPendingAccountRequest } as QueryGetPendingAccountRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetPendingAccountResponse: object = {}

export const QueryGetPendingAccountResponse = {
  encode(message: QueryGetPendingAccountResponse, writer: Writer = Writer.create()): Writer {
    if (message.pendingAccount !== undefined) {
      PendingAccount.encode(message.pendingAccount, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPendingAccountResponse } as QueryGetPendingAccountResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pendingAccount = PendingAccount.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPendingAccountResponse {
    const message = { ...baseQueryGetPendingAccountResponse } as QueryGetPendingAccountResponse
    if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
      message.pendingAccount = PendingAccount.fromJSON(object.pendingAccount)
    } else {
      message.pendingAccount = undefined
    }
    return message
  },

  toJSON(message: QueryGetPendingAccountResponse): unknown {
    const obj: any = {}
    message.pendingAccount !== undefined && (obj.pendingAccount = message.pendingAccount ? PendingAccount.toJSON(message.pendingAccount) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPendingAccountResponse>): QueryGetPendingAccountResponse {
    const message = { ...baseQueryGetPendingAccountResponse } as QueryGetPendingAccountResponse
    if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
      message.pendingAccount = PendingAccount.fromPartial(object.pendingAccount)
    } else {
      message.pendingAccount = undefined
    }
    return message
  }
}

const baseQueryAllPendingAccountRequest: object = {}

export const QueryAllPendingAccountRequest = {
  encode(message: QueryAllPendingAccountRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPendingAccountRequest } as QueryAllPendingAccountRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPendingAccountRequest {
    const message = { ...baseQueryAllPendingAccountRequest } as QueryAllPendingAccountRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPendingAccountRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPendingAccountRequest>): QueryAllPendingAccountRequest {
    const message = { ...baseQueryAllPendingAccountRequest } as QueryAllPendingAccountRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllPendingAccountResponse: object = {}

export const QueryAllPendingAccountResponse = {
  encode(message: QueryAllPendingAccountResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.pendingAccount) {
      PendingAccount.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPendingAccountResponse } as QueryAllPendingAccountResponse
    message.pendingAccount = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pendingAccount.push(PendingAccount.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPendingAccountResponse {
    const message = { ...baseQueryAllPendingAccountResponse } as QueryAllPendingAccountResponse
    message.pendingAccount = []
    if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
      for (const e of object.pendingAccount) {
        message.pendingAccount.push(PendingAccount.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPendingAccountResponse): unknown {
    const obj: any = {}
    if (message.pendingAccount) {
      obj.pendingAccount = message.pendingAccount.map((e) => (e ? PendingAccount.toJSON(e) : undefined))
    } else {
      obj.pendingAccount = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPendingAccountResponse>): QueryAllPendingAccountResponse {
    const message = { ...baseQueryAllPendingAccountResponse } as QueryAllPendingAccountResponse
    message.pendingAccount = []
    if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
      for (const e of object.pendingAccount) {
        message.pendingAccount.push(PendingAccount.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetPendingAccountRevocationRequest: object = { address: '' }

export const QueryGetPendingAccountRevocationRequest = {
  encode(message: QueryGetPendingAccountRevocationRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountRevocationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPendingAccountRevocationRequest } as QueryGetPendingAccountRevocationRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPendingAccountRevocationRequest {
    const message = { ...baseQueryGetPendingAccountRevocationRequest } as QueryGetPendingAccountRevocationRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetPendingAccountRevocationRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPendingAccountRevocationRequest>): QueryGetPendingAccountRevocationRequest {
    const message = { ...baseQueryGetPendingAccountRevocationRequest } as QueryGetPendingAccountRevocationRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetPendingAccountRevocationResponse: object = {}

export const QueryGetPendingAccountRevocationResponse = {
  encode(message: QueryGetPendingAccountRevocationResponse, writer: Writer = Writer.create()): Writer {
    if (message.pendingAccountRevocation !== undefined) {
      PendingAccountRevocation.encode(message.pendingAccountRevocation, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPendingAccountRevocationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPendingAccountRevocationResponse } as QueryGetPendingAccountRevocationResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pendingAccountRevocation = PendingAccountRevocation.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPendingAccountRevocationResponse {
    const message = { ...baseQueryGetPendingAccountRevocationResponse } as QueryGetPendingAccountRevocationResponse
    if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
      message.pendingAccountRevocation = PendingAccountRevocation.fromJSON(object.pendingAccountRevocation)
    } else {
      message.pendingAccountRevocation = undefined
    }
    return message
  },

  toJSON(message: QueryGetPendingAccountRevocationResponse): unknown {
    const obj: any = {}
    message.pendingAccountRevocation !== undefined &&
      (obj.pendingAccountRevocation = message.pendingAccountRevocation ? PendingAccountRevocation.toJSON(message.pendingAccountRevocation) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPendingAccountRevocationResponse>): QueryGetPendingAccountRevocationResponse {
    const message = { ...baseQueryGetPendingAccountRevocationResponse } as QueryGetPendingAccountRevocationResponse
    if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
      message.pendingAccountRevocation = PendingAccountRevocation.fromPartial(object.pendingAccountRevocation)
    } else {
      message.pendingAccountRevocation = undefined
    }
    return message
  }
}

const baseQueryAllPendingAccountRevocationRequest: object = {}

export const QueryAllPendingAccountRevocationRequest = {
  encode(message: QueryAllPendingAccountRevocationRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountRevocationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPendingAccountRevocationRequest } as QueryAllPendingAccountRevocationRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPendingAccountRevocationRequest {
    const message = { ...baseQueryAllPendingAccountRevocationRequest } as QueryAllPendingAccountRevocationRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPendingAccountRevocationRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPendingAccountRevocationRequest>): QueryAllPendingAccountRevocationRequest {
    const message = { ...baseQueryAllPendingAccountRevocationRequest } as QueryAllPendingAccountRevocationRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllPendingAccountRevocationResponse: object = {}

export const QueryAllPendingAccountRevocationResponse = {
  encode(message: QueryAllPendingAccountRevocationResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.pendingAccountRevocation) {
      PendingAccountRevocation.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPendingAccountRevocationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPendingAccountRevocationResponse } as QueryAllPendingAccountRevocationResponse
    message.pendingAccountRevocation = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pendingAccountRevocation.push(PendingAccountRevocation.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllPendingAccountRevocationResponse {
    const message = { ...baseQueryAllPendingAccountRevocationResponse } as QueryAllPendingAccountRevocationResponse
    message.pendingAccountRevocation = []
    if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
      for (const e of object.pendingAccountRevocation) {
        message.pendingAccountRevocation.push(PendingAccountRevocation.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPendingAccountRevocationResponse): unknown {
    const obj: any = {}
    if (message.pendingAccountRevocation) {
      obj.pendingAccountRevocation = message.pendingAccountRevocation.map((e) => (e ? PendingAccountRevocation.toJSON(e) : undefined))
    } else {
      obj.pendingAccountRevocation = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPendingAccountRevocationResponse>): QueryAllPendingAccountRevocationResponse {
    const message = { ...baseQueryAllPendingAccountRevocationResponse } as QueryAllPendingAccountRevocationResponse
    message.pendingAccountRevocation = []
    if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
      for (const e of object.pendingAccountRevocation) {
        message.pendingAccountRevocation.push(PendingAccountRevocation.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetAccountStatRequest: object = {}

export const QueryGetAccountStatRequest = {
  encode(_: QueryGetAccountStatRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAccountStatRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAccountStatRequest } as QueryGetAccountStatRequest
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

  fromJSON(_: any): QueryGetAccountStatRequest {
    const message = { ...baseQueryGetAccountStatRequest } as QueryGetAccountStatRequest
    return message
  },

  toJSON(_: QueryGetAccountStatRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetAccountStatRequest>): QueryGetAccountStatRequest {
    const message = { ...baseQueryGetAccountStatRequest } as QueryGetAccountStatRequest
    return message
  }
}

const baseQueryGetAccountStatResponse: object = {}

export const QueryGetAccountStatResponse = {
  encode(message: QueryGetAccountStatResponse, writer: Writer = Writer.create()): Writer {
    if (message.AccountStat !== undefined) {
      AccountStat.encode(message.AccountStat, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetAccountStatResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetAccountStatResponse } as QueryGetAccountStatResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.AccountStat = AccountStat.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetAccountStatResponse {
    const message = { ...baseQueryGetAccountStatResponse } as QueryGetAccountStatResponse
    if (object.AccountStat !== undefined && object.AccountStat !== null) {
      message.AccountStat = AccountStat.fromJSON(object.AccountStat)
    } else {
      message.AccountStat = undefined
    }
    return message
  },

  toJSON(message: QueryGetAccountStatResponse): unknown {
    const obj: any = {}
    message.AccountStat !== undefined && (obj.AccountStat = message.AccountStat ? AccountStat.toJSON(message.AccountStat) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetAccountStatResponse>): QueryGetAccountStatResponse {
    const message = { ...baseQueryGetAccountStatResponse } as QueryGetAccountStatResponse
    if (object.AccountStat !== undefined && object.AccountStat !== null) {
      message.AccountStat = AccountStat.fromPartial(object.AccountStat)
    } else {
      message.AccountStat = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a account by index. */
  Account(request: QueryGetAccountRequest): Promise<QueryGetAccountResponse>
  /** Queries a list of account items. */
  AccountAll(request: QueryAllAccountRequest): Promise<QueryAllAccountResponse>
  /** Queries a pendingAccount by index. */
  PendingAccount(request: QueryGetPendingAccountRequest): Promise<QueryGetPendingAccountResponse>
  /** Queries a list of pendingAccount items. */
  PendingAccountAll(request: QueryAllPendingAccountRequest): Promise<QueryAllPendingAccountResponse>
  /** Queries a pendingAccountRevocation by index. */
  PendingAccountRevocation(request: QueryGetPendingAccountRevocationRequest): Promise<QueryGetPendingAccountRevocationResponse>
  /** Queries a list of pendingAccountRevocation items. */
  PendingAccountRevocationAll(request: QueryAllPendingAccountRevocationRequest): Promise<QueryAllPendingAccountRevocationResponse>
  /** Queries a accountStat by index. */
  AccountStat(request: QueryGetAccountStatRequest): Promise<QueryGetAccountStatResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Account(request: QueryGetAccountRequest): Promise<QueryGetAccountResponse> {
    const data = QueryGetAccountRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'Account', data)
    return promise.then((data) => QueryGetAccountResponse.decode(new Reader(data)))
  }

  AccountAll(request: QueryAllAccountRequest): Promise<QueryAllAccountResponse> {
    const data = QueryAllAccountRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'AccountAll', data)
    return promise.then((data) => QueryAllAccountResponse.decode(new Reader(data)))
  }

  PendingAccount(request: QueryGetPendingAccountRequest): Promise<QueryGetPendingAccountResponse> {
    const data = QueryGetPendingAccountRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccount', data)
    return promise.then((data) => QueryGetPendingAccountResponse.decode(new Reader(data)))
  }

  PendingAccountAll(request: QueryAllPendingAccountRequest): Promise<QueryAllPendingAccountResponse> {
    const data = QueryAllPendingAccountRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccountAll', data)
    return promise.then((data) => QueryAllPendingAccountResponse.decode(new Reader(data)))
  }

  PendingAccountRevocation(request: QueryGetPendingAccountRevocationRequest): Promise<QueryGetPendingAccountRevocationResponse> {
    const data = QueryGetPendingAccountRevocationRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccountRevocation', data)
    return promise.then((data) => QueryGetPendingAccountRevocationResponse.decode(new Reader(data)))
  }

  PendingAccountRevocationAll(request: QueryAllPendingAccountRevocationRequest): Promise<QueryAllPendingAccountRevocationResponse> {
    const data = QueryAllPendingAccountRevocationRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccountRevocationAll', data)
    return promise.then((data) => QueryAllPendingAccountRevocationResponse.decode(new Reader(data)))
  }

  AccountStat(request: QueryGetAccountStatRequest): Promise<QueryGetAccountStatResponse> {
    const data = QueryGetAccountStatRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'AccountStat', data)
    return promise.then((data) => QueryGetAccountStatResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
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
