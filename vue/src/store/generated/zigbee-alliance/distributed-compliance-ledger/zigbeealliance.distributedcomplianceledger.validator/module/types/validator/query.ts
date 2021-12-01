/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Validator } from '../validator/validator'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { LastValidatorPower } from '../validator/last_validator_power'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface QueryGetValidatorRequest {
  owner: string
}

export interface QueryGetValidatorResponse {
  validator: Validator | undefined
}

export interface QueryAllValidatorRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllValidatorResponse {
  validator: Validator[]
  pagination: PageResponse | undefined
}

export interface QueryGetLastValidatorPowerRequest {
  owner: string
}

export interface QueryGetLastValidatorPowerResponse {
  lastValidatorPower: LastValidatorPower | undefined
}

export interface QueryAllLastValidatorPowerRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllLastValidatorPowerResponse {
  lastValidatorPower: LastValidatorPower[]
  pagination: PageResponse | undefined
}

const baseQueryGetValidatorRequest: object = { owner: '' }

export const QueryGetValidatorRequest = {
  encode(message: QueryGetValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.owner !== '') {
      writer.uint32(10).string(message.owner)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorRequest } as QueryGetValidatorRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetValidatorRequest {
    const message = { ...baseQueryGetValidatorRequest } as QueryGetValidatorRequest
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    return message
  },

  toJSON(message: QueryGetValidatorRequest): unknown {
    const obj: any = {}
    message.owner !== undefined && (obj.owner = message.owner)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorRequest>): QueryGetValidatorRequest {
    const message = { ...baseQueryGetValidatorRequest } as QueryGetValidatorRequest
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    return message
  }
}

const baseQueryGetValidatorResponse: object = {}

export const QueryGetValidatorResponse = {
  encode(message: QueryGetValidatorResponse, writer: Writer = Writer.create()): Writer {
    if (message.validator !== undefined) {
      Validator.encode(message.validator, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorResponse } as QueryGetValidatorResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validator = Validator.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetValidatorResponse {
    const message = { ...baseQueryGetValidatorResponse } as QueryGetValidatorResponse
    if (object.validator !== undefined && object.validator !== null) {
      message.validator = Validator.fromJSON(object.validator)
    } else {
      message.validator = undefined
    }
    return message
  },

  toJSON(message: QueryGetValidatorResponse): unknown {
    const obj: any = {}
    message.validator !== undefined && (obj.validator = message.validator ? Validator.toJSON(message.validator) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorResponse>): QueryGetValidatorResponse {
    const message = { ...baseQueryGetValidatorResponse } as QueryGetValidatorResponse
    if (object.validator !== undefined && object.validator !== null) {
      message.validator = Validator.fromPartial(object.validator)
    } else {
      message.validator = undefined
    }
    return message
  }
}

const baseQueryAllValidatorRequest: object = {}

export const QueryAllValidatorRequest = {
  encode(message: QueryAllValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorRequest } as QueryAllValidatorRequest
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

  fromJSON(object: any): QueryAllValidatorRequest {
    const message = { ...baseQueryAllValidatorRequest } as QueryAllValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorRequest>): QueryAllValidatorRequest {
    const message = { ...baseQueryAllValidatorRequest } as QueryAllValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllValidatorResponse: object = {}

export const QueryAllValidatorResponse = {
  encode(message: QueryAllValidatorResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.validator) {
      Validator.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorResponse } as QueryAllValidatorResponse
    message.validator = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validator.push(Validator.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllValidatorResponse {
    const message = { ...baseQueryAllValidatorResponse } as QueryAllValidatorResponse
    message.validator = []
    if (object.validator !== undefined && object.validator !== null) {
      for (const e of object.validator) {
        message.validator.push(Validator.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorResponse): unknown {
    const obj: any = {}
    if (message.validator) {
      obj.validator = message.validator.map((e) => (e ? Validator.toJSON(e) : undefined))
    } else {
      obj.validator = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorResponse>): QueryAllValidatorResponse {
    const message = { ...baseQueryAllValidatorResponse } as QueryAllValidatorResponse
    message.validator = []
    if (object.validator !== undefined && object.validator !== null) {
      for (const e of object.validator) {
        message.validator.push(Validator.fromPartial(e))
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

const baseQueryGetLastValidatorPowerRequest: object = { owner: '' }

export const QueryGetLastValidatorPowerRequest = {
  encode(message: QueryGetLastValidatorPowerRequest, writer: Writer = Writer.create()): Writer {
    if (message.owner !== '') {
      writer.uint32(10).string(message.owner)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLastValidatorPowerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetLastValidatorPowerRequest } as QueryGetLastValidatorPowerRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetLastValidatorPowerRequest {
    const message = { ...baseQueryGetLastValidatorPowerRequest } as QueryGetLastValidatorPowerRequest
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    return message
  },

  toJSON(message: QueryGetLastValidatorPowerRequest): unknown {
    const obj: any = {}
    message.owner !== undefined && (obj.owner = message.owner)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetLastValidatorPowerRequest>): QueryGetLastValidatorPowerRequest {
    const message = { ...baseQueryGetLastValidatorPowerRequest } as QueryGetLastValidatorPowerRequest
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    return message
  }
}

const baseQueryGetLastValidatorPowerResponse: object = {}

export const QueryGetLastValidatorPowerResponse = {
  encode(message: QueryGetLastValidatorPowerResponse, writer: Writer = Writer.create()): Writer {
    if (message.lastValidatorPower !== undefined) {
      LastValidatorPower.encode(message.lastValidatorPower, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLastValidatorPowerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetLastValidatorPowerResponse } as QueryGetLastValidatorPowerResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.lastValidatorPower = LastValidatorPower.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetLastValidatorPowerResponse {
    const message = { ...baseQueryGetLastValidatorPowerResponse } as QueryGetLastValidatorPowerResponse
    if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
      message.lastValidatorPower = LastValidatorPower.fromJSON(object.lastValidatorPower)
    } else {
      message.lastValidatorPower = undefined
    }
    return message
  },

  toJSON(message: QueryGetLastValidatorPowerResponse): unknown {
    const obj: any = {}
    message.lastValidatorPower !== undefined &&
      (obj.lastValidatorPower = message.lastValidatorPower ? LastValidatorPower.toJSON(message.lastValidatorPower) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetLastValidatorPowerResponse>): QueryGetLastValidatorPowerResponse {
    const message = { ...baseQueryGetLastValidatorPowerResponse } as QueryGetLastValidatorPowerResponse
    if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
      message.lastValidatorPower = LastValidatorPower.fromPartial(object.lastValidatorPower)
    } else {
      message.lastValidatorPower = undefined
    }
    return message
  }
}

const baseQueryAllLastValidatorPowerRequest: object = {}

export const QueryAllLastValidatorPowerRequest = {
  encode(message: QueryAllLastValidatorPowerRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllLastValidatorPowerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllLastValidatorPowerRequest } as QueryAllLastValidatorPowerRequest
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

  fromJSON(object: any): QueryAllLastValidatorPowerRequest {
    const message = { ...baseQueryAllLastValidatorPowerRequest } as QueryAllLastValidatorPowerRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllLastValidatorPowerRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllLastValidatorPowerRequest>): QueryAllLastValidatorPowerRequest {
    const message = { ...baseQueryAllLastValidatorPowerRequest } as QueryAllLastValidatorPowerRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllLastValidatorPowerResponse: object = {}

export const QueryAllLastValidatorPowerResponse = {
  encode(message: QueryAllLastValidatorPowerResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.lastValidatorPower) {
      LastValidatorPower.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllLastValidatorPowerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllLastValidatorPowerResponse } as QueryAllLastValidatorPowerResponse
    message.lastValidatorPower = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.lastValidatorPower.push(LastValidatorPower.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllLastValidatorPowerResponse {
    const message = { ...baseQueryAllLastValidatorPowerResponse } as QueryAllLastValidatorPowerResponse
    message.lastValidatorPower = []
    if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
      for (const e of object.lastValidatorPower) {
        message.lastValidatorPower.push(LastValidatorPower.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllLastValidatorPowerResponse): unknown {
    const obj: any = {}
    if (message.lastValidatorPower) {
      obj.lastValidatorPower = message.lastValidatorPower.map((e) => (e ? LastValidatorPower.toJSON(e) : undefined))
    } else {
      obj.lastValidatorPower = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllLastValidatorPowerResponse>): QueryAllLastValidatorPowerResponse {
    const message = { ...baseQueryAllLastValidatorPowerResponse } as QueryAllLastValidatorPowerResponse
    message.lastValidatorPower = []
    if (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null) {
      for (const e of object.lastValidatorPower) {
        message.lastValidatorPower.push(LastValidatorPower.fromPartial(e))
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

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a validator by index. */
  Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse>
  /** Queries a list of validator items. */
  ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>
  /** Queries a lastValidatorPower by index. */
  LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse>
  /** Queries a list of lastValidatorPower items. */
  LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse> {
    const data = QueryGetValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'Validator', data)
    return promise.then((data) => QueryGetValidatorResponse.decode(new Reader(data)))
  }

  ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse> {
    const data = QueryAllValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorAll', data)
    return promise.then((data) => QueryAllValidatorResponse.decode(new Reader(data)))
  }

  LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse> {
    const data = QueryGetLastValidatorPowerRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'LastValidatorPower', data)
    return promise.then((data) => QueryGetLastValidatorPowerResponse.decode(new Reader(data)))
  }

  LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse> {
    const data = QueryAllLastValidatorPowerRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'LastValidatorPowerAll', data)
    return promise.then((data) => QueryAllLastValidatorPowerResponse.decode(new Reader(data)))
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
