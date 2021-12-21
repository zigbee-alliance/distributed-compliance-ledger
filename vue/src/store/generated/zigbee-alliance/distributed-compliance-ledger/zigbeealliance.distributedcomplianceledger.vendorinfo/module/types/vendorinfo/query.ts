/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { VendorInfoType } from '../vendorinfo/vendor_info_type'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface QueryGetVendorInfoTypeRequest {
  vendorID: number
}

export interface QueryGetVendorInfoTypeResponse {
  vendorInfoType: VendorInfoType | undefined
}

export interface QueryAllVendorInfoTypeRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllVendorInfoTypeResponse {
  vendorInfoType: VendorInfoType[]
  pagination: PageResponse | undefined
}

const baseQueryGetVendorInfoTypeRequest: object = { vendorID: 0 }

export const QueryGetVendorInfoTypeRequest = {
  encode(message: QueryGetVendorInfoTypeRequest, writer: Writer = Writer.create()): Writer {
    if (message.vendorID !== 0) {
      writer.uint32(8).uint64(message.vendorID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoTypeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetVendorInfoTypeRequest } as QueryGetVendorInfoTypeRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorID = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetVendorInfoTypeRequest {
    const message = { ...baseQueryGetVendorInfoTypeRequest } as QueryGetVendorInfoTypeRequest
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    return message
  },

  toJSON(message: QueryGetVendorInfoTypeRequest): unknown {
    const obj: any = {}
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetVendorInfoTypeRequest>): QueryGetVendorInfoTypeRequest {
    const message = { ...baseQueryGetVendorInfoTypeRequest } as QueryGetVendorInfoTypeRequest
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
    }
    return message
  }
}

const baseQueryGetVendorInfoTypeResponse: object = {}

export const QueryGetVendorInfoTypeResponse = {
  encode(message: QueryGetVendorInfoTypeResponse, writer: Writer = Writer.create()): Writer {
    if (message.vendorInfoType !== undefined) {
      VendorInfoType.encode(message.vendorInfoType, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoTypeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetVendorInfoTypeResponse } as QueryGetVendorInfoTypeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorInfoType = VendorInfoType.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetVendorInfoTypeResponse {
    const message = { ...baseQueryGetVendorInfoTypeResponse } as QueryGetVendorInfoTypeResponse
    if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
      message.vendorInfoType = VendorInfoType.fromJSON(object.vendorInfoType)
    } else {
      message.vendorInfoType = undefined
    }
    return message
  },

  toJSON(message: QueryGetVendorInfoTypeResponse): unknown {
    const obj: any = {}
    message.vendorInfoType !== undefined && (obj.vendorInfoType = message.vendorInfoType ? VendorInfoType.toJSON(message.vendorInfoType) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetVendorInfoTypeResponse>): QueryGetVendorInfoTypeResponse {
    const message = { ...baseQueryGetVendorInfoTypeResponse } as QueryGetVendorInfoTypeResponse
    if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
      message.vendorInfoType = VendorInfoType.fromPartial(object.vendorInfoType)
    } else {
      message.vendorInfoType = undefined
    }
    return message
  }
}

const baseQueryAllVendorInfoTypeRequest: object = {}

export const QueryAllVendorInfoTypeRequest = {
  encode(message: QueryAllVendorInfoTypeRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoTypeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllVendorInfoTypeRequest } as QueryAllVendorInfoTypeRequest
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

  fromJSON(object: any): QueryAllVendorInfoTypeRequest {
    const message = { ...baseQueryAllVendorInfoTypeRequest } as QueryAllVendorInfoTypeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllVendorInfoTypeRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllVendorInfoTypeRequest>): QueryAllVendorInfoTypeRequest {
    const message = { ...baseQueryAllVendorInfoTypeRequest } as QueryAllVendorInfoTypeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllVendorInfoTypeResponse: object = {}

export const QueryAllVendorInfoTypeResponse = {
  encode(message: QueryAllVendorInfoTypeResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.vendorInfoType) {
      VendorInfoType.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoTypeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllVendorInfoTypeResponse } as QueryAllVendorInfoTypeResponse
    message.vendorInfoType = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorInfoType.push(VendorInfoType.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllVendorInfoTypeResponse {
    const message = { ...baseQueryAllVendorInfoTypeResponse } as QueryAllVendorInfoTypeResponse
    message.vendorInfoType = []
    if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
      for (const e of object.vendorInfoType) {
        message.vendorInfoType.push(VendorInfoType.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllVendorInfoTypeResponse): unknown {
    const obj: any = {}
    if (message.vendorInfoType) {
      obj.vendorInfoType = message.vendorInfoType.map((e) => (e ? VendorInfoType.toJSON(e) : undefined))
    } else {
      obj.vendorInfoType = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllVendorInfoTypeResponse>): QueryAllVendorInfoTypeResponse {
    const message = { ...baseQueryAllVendorInfoTypeResponse } as QueryAllVendorInfoTypeResponse
    message.vendorInfoType = []
    if (object.vendorInfoType !== undefined && object.vendorInfoType !== null) {
      for (const e of object.vendorInfoType) {
        message.vendorInfoType.push(VendorInfoType.fromPartial(e))
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
  /** Queries a vendorInfoType by index. */
  VendorInfoType(request: QueryGetVendorInfoTypeRequest): Promise<QueryGetVendorInfoTypeResponse>
  /** Queries a list of vendorInfoType items. */
  VendorInfoTypeAll(request: QueryAllVendorInfoTypeRequest): Promise<QueryAllVendorInfoTypeResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  VendorInfoType(request: QueryGetVendorInfoTypeRequest): Promise<QueryGetVendorInfoTypeResponse> {
    const data = QueryGetVendorInfoTypeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfoType', data)
    return promise.then((data) => QueryGetVendorInfoTypeResponse.decode(new Reader(data)))
  }

  VendorInfoTypeAll(request: QueryAllVendorInfoTypeRequest): Promise<QueryAllVendorInfoTypeResponse> {
    const data = QueryAllVendorInfoTypeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfoTypeAll', data)
    return promise.then((data) => QueryAllVendorInfoTypeResponse.decode(new Reader(data)))
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
