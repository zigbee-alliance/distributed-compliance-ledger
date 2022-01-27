/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Params } from '../dclupgrade/params'
import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade'

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined
}

export interface QueryGetProposedUpgradeRequest {
  name: string
}

export interface QueryGetProposedUpgradeResponse {
  proposedUpgrade: ProposedUpgrade | undefined
}

export interface QueryAllProposedUpgradeRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllProposedUpgradeResponse {
  proposedUpgrade: ProposedUpgrade[]
  pagination: PageResponse | undefined
}

const baseQueryParamsRequest: object = {}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest
    return message
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest
    return message
  }
}

const baseQueryParamsResponse: object = {}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params)
    } else {
      message.params = undefined
    }
    return message
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {}
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params)
    } else {
      message.params = undefined
    }
    return message
  }
}

const baseQueryGetProposedUpgradeRequest: object = { name: '' }

export const QueryGetProposedUpgradeRequest = {
  encode(message: QueryGetProposedUpgradeRequest, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedUpgradeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedUpgradeRequest } as QueryGetProposedUpgradeRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProposedUpgradeRequest {
    const message = { ...baseQueryGetProposedUpgradeRequest } as QueryGetProposedUpgradeRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    return message
  },

  toJSON(message: QueryGetProposedUpgradeRequest): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedUpgradeRequest>): QueryGetProposedUpgradeRequest {
    const message = { ...baseQueryGetProposedUpgradeRequest } as QueryGetProposedUpgradeRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    return message
  }
}

const baseQueryGetProposedUpgradeResponse: object = {}

export const QueryGetProposedUpgradeResponse = {
  encode(message: QueryGetProposedUpgradeResponse, writer: Writer = Writer.create()): Writer {
    if (message.proposedUpgrade !== undefined) {
      ProposedUpgrade.encode(message.proposedUpgrade, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedUpgradeResponse } as QueryGetProposedUpgradeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedUpgrade = ProposedUpgrade.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProposedUpgradeResponse {
    const message = { ...baseQueryGetProposedUpgradeResponse } as QueryGetProposedUpgradeResponse
    if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
      message.proposedUpgrade = ProposedUpgrade.fromJSON(object.proposedUpgrade)
    } else {
      message.proposedUpgrade = undefined
    }
    return message
  },

  toJSON(message: QueryGetProposedUpgradeResponse): unknown {
    const obj: any = {}
    message.proposedUpgrade !== undefined && (obj.proposedUpgrade = message.proposedUpgrade ? ProposedUpgrade.toJSON(message.proposedUpgrade) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedUpgradeResponse>): QueryGetProposedUpgradeResponse {
    const message = { ...baseQueryGetProposedUpgradeResponse } as QueryGetProposedUpgradeResponse
    if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
      message.proposedUpgrade = ProposedUpgrade.fromPartial(object.proposedUpgrade)
    } else {
      message.proposedUpgrade = undefined
    }
    return message
  }
}

const baseQueryAllProposedUpgradeRequest: object = {}

export const QueryAllProposedUpgradeRequest = {
  encode(message: QueryAllProposedUpgradeRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedUpgradeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedUpgradeRequest } as QueryAllProposedUpgradeRequest
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

  fromJSON(object: any): QueryAllProposedUpgradeRequest {
    const message = { ...baseQueryAllProposedUpgradeRequest } as QueryAllProposedUpgradeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedUpgradeRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedUpgradeRequest>): QueryAllProposedUpgradeRequest {
    const message = { ...baseQueryAllProposedUpgradeRequest } as QueryAllProposedUpgradeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllProposedUpgradeResponse: object = {}

export const QueryAllProposedUpgradeResponse = {
  encode(message: QueryAllProposedUpgradeResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.proposedUpgrade) {
      ProposedUpgrade.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedUpgradeResponse } as QueryAllProposedUpgradeResponse
    message.proposedUpgrade = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedUpgrade.push(ProposedUpgrade.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllProposedUpgradeResponse {
    const message = { ...baseQueryAllProposedUpgradeResponse } as QueryAllProposedUpgradeResponse
    message.proposedUpgrade = []
    if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
      for (const e of object.proposedUpgrade) {
        message.proposedUpgrade.push(ProposedUpgrade.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedUpgradeResponse): unknown {
    const obj: any = {}
    if (message.proposedUpgrade) {
      obj.proposedUpgrade = message.proposedUpgrade.map((e) => (e ? ProposedUpgrade.toJSON(e) : undefined))
    } else {
      obj.proposedUpgrade = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedUpgradeResponse>): QueryAllProposedUpgradeResponse {
    const message = { ...baseQueryAllProposedUpgradeResponse } as QueryAllProposedUpgradeResponse
    message.proposedUpgrade = []
    if (object.proposedUpgrade !== undefined && object.proposedUpgrade !== null) {
      for (const e of object.proposedUpgrade) {
        message.proposedUpgrade.push(ProposedUpgrade.fromPartial(e))
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
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>
  /** Queries a ProposedUpgrade by index. */
  ProposedUpgrade(request: QueryGetProposedUpgradeRequest): Promise<QueryGetProposedUpgradeResponse>
  /** Queries a list of ProposedUpgrade items. */
  ProposedUpgradeAll(request: QueryAllProposedUpgradeRequest): Promise<QueryAllProposedUpgradeResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'Params', data)
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)))
  }

  ProposedUpgrade(request: QueryGetProposedUpgradeRequest): Promise<QueryGetProposedUpgradeResponse> {
    const data = QueryGetProposedUpgradeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ProposedUpgrade', data)
    return promise.then((data) => QueryGetProposedUpgradeResponse.decode(new Reader(data)))
  }

  ProposedUpgradeAll(request: QueryAllProposedUpgradeRequest): Promise<QueryAllProposedUpgradeResponse> {
    const data = QueryAllProposedUpgradeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ProposedUpgradeAll', data)
    return promise.then((data) => QueryAllProposedUpgradeResponse.decode(new Reader(data)))
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
