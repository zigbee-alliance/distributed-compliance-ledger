/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { ComplianceInfo } from '../compliance/compliance_info'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

export interface QueryGetComplianceInfoRequest {
  vid: number
  pid: number
  softwareVersion: number
  certificationType: string
}

export interface QueryGetComplianceInfoResponse {
  complianceInfo: ComplianceInfo | undefined
}

export interface QueryAllComplianceInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllComplianceInfoResponse {
  complianceInfo: ComplianceInfo[]
  pagination: PageResponse | undefined
}

const baseQueryGetComplianceInfoRequest: object = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' }

export const QueryGetComplianceInfoRequest = {
  encode(message: QueryGetComplianceInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint64(message.softwareVersion)
    }
    if (message.certificationType !== '') {
      writer.uint32(34).string(message.certificationType)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetComplianceInfoRequest } as QueryGetComplianceInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        case 3:
          message.softwareVersion = longToNumber(reader.uint64() as Long)
          break
        case 4:
          message.certificationType = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetComplianceInfoRequest {
    const message = { ...baseQueryGetComplianceInfoRequest } as QueryGetComplianceInfoRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = String(object.certificationType)
    } else {
      message.certificationType = ''
    }
    return message
  },

  toJSON(message: QueryGetComplianceInfoRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    message.certificationType !== undefined && (obj.certificationType = message.certificationType)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetComplianceInfoRequest>): QueryGetComplianceInfoRequest {
    const message = { ...baseQueryGetComplianceInfoRequest } as QueryGetComplianceInfoRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = object.certificationType
    } else {
      message.certificationType = ''
    }
    return message
  }
}

const baseQueryGetComplianceInfoResponse: object = {}

export const QueryGetComplianceInfoResponse = {
  encode(message: QueryGetComplianceInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.complianceInfo !== undefined) {
      ComplianceInfo.encode(message.complianceInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetComplianceInfoResponse } as QueryGetComplianceInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.complianceInfo = ComplianceInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetComplianceInfoResponse {
    const message = { ...baseQueryGetComplianceInfoResponse } as QueryGetComplianceInfoResponse
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      message.complianceInfo = ComplianceInfo.fromJSON(object.complianceInfo)
    } else {
      message.complianceInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetComplianceInfoResponse): unknown {
    const obj: any = {}
    message.complianceInfo !== undefined && (obj.complianceInfo = message.complianceInfo ? ComplianceInfo.toJSON(message.complianceInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetComplianceInfoResponse>): QueryGetComplianceInfoResponse {
    const message = { ...baseQueryGetComplianceInfoResponse } as QueryGetComplianceInfoResponse
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      message.complianceInfo = ComplianceInfo.fromPartial(object.complianceInfo)
    } else {
      message.complianceInfo = undefined
    }
    return message
  }
}

const baseQueryAllComplianceInfoRequest: object = {}

export const QueryAllComplianceInfoRequest = {
  encode(message: QueryAllComplianceInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllComplianceInfoRequest } as QueryAllComplianceInfoRequest
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

  fromJSON(object: any): QueryAllComplianceInfoRequest {
    const message = { ...baseQueryAllComplianceInfoRequest } as QueryAllComplianceInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllComplianceInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllComplianceInfoRequest>): QueryAllComplianceInfoRequest {
    const message = { ...baseQueryAllComplianceInfoRequest } as QueryAllComplianceInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllComplianceInfoResponse: object = {}

export const QueryAllComplianceInfoResponse = {
  encode(message: QueryAllComplianceInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.complianceInfo) {
      ComplianceInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllComplianceInfoResponse } as QueryAllComplianceInfoResponse
    message.complianceInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.complianceInfo.push(ComplianceInfo.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllComplianceInfoResponse {
    const message = { ...baseQueryAllComplianceInfoResponse } as QueryAllComplianceInfoResponse
    message.complianceInfo = []
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(ComplianceInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllComplianceInfoResponse): unknown {
    const obj: any = {}
    if (message.complianceInfo) {
      obj.complianceInfo = message.complianceInfo.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined))
    } else {
      obj.complianceInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllComplianceInfoResponse>): QueryAllComplianceInfoResponse {
    const message = { ...baseQueryAllComplianceInfoResponse } as QueryAllComplianceInfoResponse
    message.complianceInfo = []
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(ComplianceInfo.fromPartial(e))
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
  /** Queries a ComplianceInfo by index. */
  ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse>
  /** Queries a list of ComplianceInfo items. */
  ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse> {
    const data = QueryGetComplianceInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfo', data)
    return promise.then((data) => QueryGetComplianceInfoResponse.decode(new Reader(data)))
  }

  ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse> {
    const data = QueryAllComplianceInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfoAll', data)
    return promise.then((data) => QueryAllComplianceInfoResponse.decode(new Reader(data)))
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
