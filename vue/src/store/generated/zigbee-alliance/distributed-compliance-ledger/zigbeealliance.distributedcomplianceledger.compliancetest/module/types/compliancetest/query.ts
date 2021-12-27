/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { TestingResults } from '../compliancetest/testing_results'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliancetest'

export interface QueryGetTestingResultsRequest {
  vid: number
  pid: number
  softwareVersion: number
}

export interface QueryGetTestingResultsResponse {
  testingResults: TestingResults | undefined
}

export interface QueryAllTestingResultsRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllTestingResultsResponse {
  testingResults: TestingResults[]
  pagination: PageResponse | undefined
}

const baseQueryGetTestingResultsRequest: object = { vid: 0, pid: 0, softwareVersion: 0 }

export const QueryGetTestingResultsRequest = {
  encode(message: QueryGetTestingResultsRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTestingResultsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTestingResultsRequest } as QueryGetTestingResultsRequest
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
          message.softwareVersion = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetTestingResultsRequest {
    const message = { ...baseQueryGetTestingResultsRequest } as QueryGetTestingResultsRequest
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
    return message
  },

  toJSON(message: QueryGetTestingResultsRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetTestingResultsRequest>): QueryGetTestingResultsRequest {
    const message = { ...baseQueryGetTestingResultsRequest } as QueryGetTestingResultsRequest
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
    return message
  }
}

const baseQueryGetTestingResultsResponse: object = {}

export const QueryGetTestingResultsResponse = {
  encode(message: QueryGetTestingResultsResponse, writer: Writer = Writer.create()): Writer {
    if (message.testingResults !== undefined) {
      TestingResults.encode(message.testingResults, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetTestingResultsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetTestingResultsResponse } as QueryGetTestingResultsResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.testingResults = TestingResults.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetTestingResultsResponse {
    const message = { ...baseQueryGetTestingResultsResponse } as QueryGetTestingResultsResponse
    if (object.testingResults !== undefined && object.testingResults !== null) {
      message.testingResults = TestingResults.fromJSON(object.testingResults)
    } else {
      message.testingResults = undefined
    }
    return message
  },

  toJSON(message: QueryGetTestingResultsResponse): unknown {
    const obj: any = {}
    message.testingResults !== undefined && (obj.testingResults = message.testingResults ? TestingResults.toJSON(message.testingResults) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetTestingResultsResponse>): QueryGetTestingResultsResponse {
    const message = { ...baseQueryGetTestingResultsResponse } as QueryGetTestingResultsResponse
    if (object.testingResults !== undefined && object.testingResults !== null) {
      message.testingResults = TestingResults.fromPartial(object.testingResults)
    } else {
      message.testingResults = undefined
    }
    return message
  }
}

const baseQueryAllTestingResultsRequest: object = {}

export const QueryAllTestingResultsRequest = {
  encode(message: QueryAllTestingResultsRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTestingResultsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllTestingResultsRequest } as QueryAllTestingResultsRequest
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

  fromJSON(object: any): QueryAllTestingResultsRequest {
    const message = { ...baseQueryAllTestingResultsRequest } as QueryAllTestingResultsRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllTestingResultsRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllTestingResultsRequest>): QueryAllTestingResultsRequest {
    const message = { ...baseQueryAllTestingResultsRequest } as QueryAllTestingResultsRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllTestingResultsResponse: object = {}

export const QueryAllTestingResultsResponse = {
  encode(message: QueryAllTestingResultsResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.testingResults) {
      TestingResults.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllTestingResultsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllTestingResultsResponse } as QueryAllTestingResultsResponse
    message.testingResults = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.testingResults.push(TestingResults.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllTestingResultsResponse {
    const message = { ...baseQueryAllTestingResultsResponse } as QueryAllTestingResultsResponse
    message.testingResults = []
    if (object.testingResults !== undefined && object.testingResults !== null) {
      for (const e of object.testingResults) {
        message.testingResults.push(TestingResults.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllTestingResultsResponse): unknown {
    const obj: any = {}
    if (message.testingResults) {
      obj.testingResults = message.testingResults.map((e) => (e ? TestingResults.toJSON(e) : undefined))
    } else {
      obj.testingResults = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllTestingResultsResponse>): QueryAllTestingResultsResponse {
    const message = { ...baseQueryAllTestingResultsResponse } as QueryAllTestingResultsResponse
    message.testingResults = []
    if (object.testingResults !== undefined && object.testingResults !== null) {
      for (const e of object.testingResults) {
        message.testingResults.push(TestingResults.fromPartial(e))
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
  /** Queries a TestingResults by index. */
  TestingResults(request: QueryGetTestingResultsRequest): Promise<QueryGetTestingResultsResponse>
  /** Queries a list of TestingResults items. */
  TestingResultsAll(request: QueryAllTestingResultsRequest): Promise<QueryAllTestingResultsResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  TestingResults(request: QueryGetTestingResultsRequest): Promise<QueryGetTestingResultsResponse> {
    const data = QueryGetTestingResultsRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliancetest.Query', 'TestingResults', data)
    return promise.then((data) => QueryGetTestingResultsResponse.decode(new Reader(data)))
  }

  TestingResultsAll(request: QueryAllTestingResultsRequest): Promise<QueryAllTestingResultsResponse> {
    const data = QueryAllTestingResultsRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliancetest.Query', 'TestingResultsAll', data)
    return promise.then((data) => QueryAllTestingResultsResponse.decode(new Reader(data)))
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
