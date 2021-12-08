/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { NewVendorInfo } from '../vendorinfo/new_vendor_info'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface QueryGetNewVendorInfoRequest {
  index: string
}

export interface QueryGetNewVendorInfoResponse {
  newVendorInfo: NewVendorInfo | undefined
}

export interface QueryAllNewVendorInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllNewVendorInfoResponse {
  newVendorInfo: NewVendorInfo[]
  pagination: PageResponse | undefined
}

const baseQueryGetNewVendorInfoRequest: object = { index: '' }

export const QueryGetNewVendorInfoRequest = {
  encode(message: QueryGetNewVendorInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.index !== '') {
      writer.uint32(10).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNewVendorInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNewVendorInfoRequest } as QueryGetNewVendorInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNewVendorInfoRequest {
    const message = { ...baseQueryGetNewVendorInfoRequest } as QueryGetNewVendorInfoRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: QueryGetNewVendorInfoRequest): unknown {
    const obj: any = {}
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNewVendorInfoRequest>): QueryGetNewVendorInfoRequest {
    const message = { ...baseQueryGetNewVendorInfoRequest } as QueryGetNewVendorInfoRequest
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseQueryGetNewVendorInfoResponse: object = {}

export const QueryGetNewVendorInfoResponse = {
  encode(message: QueryGetNewVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.newVendorInfo !== undefined) {
      NewVendorInfo.encode(message.newVendorInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNewVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNewVendorInfoResponse } as QueryGetNewVendorInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.newVendorInfo = NewVendorInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNewVendorInfoResponse {
    const message = { ...baseQueryGetNewVendorInfoResponse } as QueryGetNewVendorInfoResponse
    if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
      message.newVendorInfo = NewVendorInfo.fromJSON(object.newVendorInfo)
    } else {
      message.newVendorInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetNewVendorInfoResponse): unknown {
    const obj: any = {}
    message.newVendorInfo !== undefined && (obj.newVendorInfo = message.newVendorInfo ? NewVendorInfo.toJSON(message.newVendorInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNewVendorInfoResponse>): QueryGetNewVendorInfoResponse {
    const message = { ...baseQueryGetNewVendorInfoResponse } as QueryGetNewVendorInfoResponse
    if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
      message.newVendorInfo = NewVendorInfo.fromPartial(object.newVendorInfo)
    } else {
      message.newVendorInfo = undefined
    }
    return message
  }
}

const baseQueryAllNewVendorInfoRequest: object = {}

export const QueryAllNewVendorInfoRequest = {
  encode(message: QueryAllNewVendorInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNewVendorInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNewVendorInfoRequest } as QueryAllNewVendorInfoRequest
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

  fromJSON(object: any): QueryAllNewVendorInfoRequest {
    const message = { ...baseQueryAllNewVendorInfoRequest } as QueryAllNewVendorInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNewVendorInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNewVendorInfoRequest>): QueryAllNewVendorInfoRequest {
    const message = { ...baseQueryAllNewVendorInfoRequest } as QueryAllNewVendorInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllNewVendorInfoResponse: object = {}

export const QueryAllNewVendorInfoResponse = {
  encode(message: QueryAllNewVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.newVendorInfo) {
      NewVendorInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNewVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNewVendorInfoResponse } as QueryAllNewVendorInfoResponse
    message.newVendorInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.newVendorInfo.push(NewVendorInfo.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllNewVendorInfoResponse {
    const message = { ...baseQueryAllNewVendorInfoResponse } as QueryAllNewVendorInfoResponse
    message.newVendorInfo = []
    if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
      for (const e of object.newVendorInfo) {
        message.newVendorInfo.push(NewVendorInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNewVendorInfoResponse): unknown {
    const obj: any = {}
    if (message.newVendorInfo) {
      obj.newVendorInfo = message.newVendorInfo.map((e) => (e ? NewVendorInfo.toJSON(e) : undefined))
    } else {
      obj.newVendorInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNewVendorInfoResponse>): QueryAllNewVendorInfoResponse {
    const message = { ...baseQueryAllNewVendorInfoResponse } as QueryAllNewVendorInfoResponse
    message.newVendorInfo = []
    if (object.newVendorInfo !== undefined && object.newVendorInfo !== null) {
      for (const e of object.newVendorInfo) {
        message.newVendorInfo.push(NewVendorInfo.fromPartial(e))
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
  /** Queries a newVendorInfo by index. */
  NewVendorInfo(request: QueryGetNewVendorInfoRequest): Promise<QueryGetNewVendorInfoResponse>
  /** Queries a list of newVendorInfo items. */
  NewVendorInfoAll(request: QueryAllNewVendorInfoRequest): Promise<QueryAllNewVendorInfoResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  NewVendorInfo(request: QueryGetNewVendorInfoRequest): Promise<QueryGetNewVendorInfoResponse> {
    const data = QueryGetNewVendorInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'NewVendorInfo', data)
    return promise.then((data) => QueryGetNewVendorInfoResponse.decode(new Reader(data)))
  }

  NewVendorInfoAll(request: QueryAllNewVendorInfoRequest): Promise<QueryAllNewVendorInfoResponse> {
    const data = QueryAllNewVendorInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'NewVendorInfoAll', data)
    return promise.then((data) => QueryAllNewVendorInfoResponse.decode(new Reader(data)))
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
