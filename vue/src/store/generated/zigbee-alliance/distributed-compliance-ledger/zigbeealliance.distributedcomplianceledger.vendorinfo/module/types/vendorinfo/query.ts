/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { VendorInfo } from '../vendorinfo/vendor_info'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface QueryGetVendorInfoRequest {
  vendorID: number
}

export interface QueryGetVendorInfoResponse {
  vendorInfo: VendorInfo | undefined
}

export interface QueryAllVendorInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllVendorInfoResponse {
  vendorInfo: VendorInfo[]
  pagination: PageResponse | undefined
}

const baseQueryGetVendorInfoRequest: object = { vendorID: 0 }

export const QueryGetVendorInfoRequest = {
  encode(message: QueryGetVendorInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.vendorID !== 0) {
      writer.uint32(8).int32(message.vendorID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetVendorInfoRequest } as QueryGetVendorInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorID = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetVendorInfoRequest {
    const message = { ...baseQueryGetVendorInfoRequest } as QueryGetVendorInfoRequest
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    return message
  },

  toJSON(message: QueryGetVendorInfoRequest): unknown {
    const obj: any = {}
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetVendorInfoRequest>): QueryGetVendorInfoRequest {
    const message = { ...baseQueryGetVendorInfoRequest } as QueryGetVendorInfoRequest
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
    }
    return message
  }
}

const baseQueryGetVendorInfoResponse: object = {}

export const QueryGetVendorInfoResponse = {
  encode(message: QueryGetVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.vendorInfo !== undefined) {
      VendorInfo.encode(message.vendorInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetVendorInfoResponse } as QueryGetVendorInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorInfo = VendorInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetVendorInfoResponse {
    const message = { ...baseQueryGetVendorInfoResponse } as QueryGetVendorInfoResponse
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromJSON(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetVendorInfoResponse): unknown {
    const obj: any = {}
    message.vendorInfo !== undefined && (obj.vendorInfo = message.vendorInfo ? VendorInfo.toJSON(message.vendorInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetVendorInfoResponse>): QueryGetVendorInfoResponse {
    const message = { ...baseQueryGetVendorInfoResponse } as QueryGetVendorInfoResponse
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromPartial(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    return message
  }
}

const baseQueryAllVendorInfoRequest: object = {}

export const QueryAllVendorInfoRequest = {
  encode(message: QueryAllVendorInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllVendorInfoRequest } as QueryAllVendorInfoRequest
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

  fromJSON(object: any): QueryAllVendorInfoRequest {
    const message = { ...baseQueryAllVendorInfoRequest } as QueryAllVendorInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllVendorInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllVendorInfoRequest>): QueryAllVendorInfoRequest {
    const message = { ...baseQueryAllVendorInfoRequest } as QueryAllVendorInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllVendorInfoResponse: object = {}

export const QueryAllVendorInfoResponse = {
  encode(message: QueryAllVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.vendorInfo) {
      VendorInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllVendorInfoResponse } as QueryAllVendorInfoResponse
    message.vendorInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorInfo.push(VendorInfo.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllVendorInfoResponse {
    const message = { ...baseQueryAllVendorInfoResponse } as QueryAllVendorInfoResponse
    message.vendorInfo = []
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      for (const e of object.vendorInfo) {
        message.vendorInfo.push(VendorInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllVendorInfoResponse): unknown {
    const obj: any = {}
    if (message.vendorInfo) {
      obj.vendorInfo = message.vendorInfo.map((e) => (e ? VendorInfo.toJSON(e) : undefined))
    } else {
      obj.vendorInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllVendorInfoResponse>): QueryAllVendorInfoResponse {
    const message = { ...baseQueryAllVendorInfoResponse } as QueryAllVendorInfoResponse
    message.vendorInfo = []
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      for (const e of object.vendorInfo) {
        message.vendorInfo.push(VendorInfo.fromPartial(e))
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
  /** Queries a vendorInfo by index. */
  VendorInfo(request: QueryGetVendorInfoRequest): Promise<QueryGetVendorInfoResponse>
  /** Queries a list of vendorInfo items. */
  VendorInfoAll(request: QueryAllVendorInfoRequest): Promise<QueryAllVendorInfoResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  VendorInfo(request: QueryGetVendorInfoRequest): Promise<QueryGetVendorInfoResponse> {
    const data = QueryGetVendorInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfo', data)
    return promise.then((data) => QueryGetVendorInfoResponse.decode(new Reader(data)))
  }

  VendorInfoAll(request: QueryAllVendorInfoRequest): Promise<QueryAllVendorInfoResponse> {
    const data = QueryAllVendorInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Query', 'VendorInfoAll', data)
    return promise.then((data) => QueryAllVendorInfoResponse.decode(new Reader(data)))
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
