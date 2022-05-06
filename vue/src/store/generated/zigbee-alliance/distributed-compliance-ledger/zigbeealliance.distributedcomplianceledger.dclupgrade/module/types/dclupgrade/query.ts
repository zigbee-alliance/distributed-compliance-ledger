/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { ApprovedUpgrade } from '../dclupgrade/approved_upgrade'
import { RejectedUpgrade } from '../dclupgrade/rejected_upgrade'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade'

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

export interface QueryGetApprovedUpgradeRequest {
  name: string
}

export interface QueryGetApprovedUpgradeResponse {
  approvedUpgrade: ApprovedUpgrade | undefined
}

export interface QueryAllApprovedUpgradeRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllApprovedUpgradeResponse {
  approvedUpgrade: ApprovedUpgrade[]
  pagination: PageResponse | undefined
}

export interface QueryGetRejectedUpgradeRequest {
  name: string
}

export interface QueryGetRejectedUpgradeResponse {
  rejectedUpgrade: RejectedUpgrade | undefined
}

export interface QueryAllRejectedUpgradeRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllRejectedUpgradeResponse {
  rejectedUpgrade: RejectedUpgrade[]
  pagination: PageResponse | undefined
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

const baseQueryGetApprovedUpgradeRequest: object = { name: '' }

export const QueryGetApprovedUpgradeRequest = {
  encode(message: QueryGetApprovedUpgradeRequest, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedUpgradeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedUpgradeRequest } as QueryGetApprovedUpgradeRequest
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

  fromJSON(object: any): QueryGetApprovedUpgradeRequest {
    const message = { ...baseQueryGetApprovedUpgradeRequest } as QueryGetApprovedUpgradeRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    return message
  },

  toJSON(message: QueryGetApprovedUpgradeRequest): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetApprovedUpgradeRequest>): QueryGetApprovedUpgradeRequest {
    const message = { ...baseQueryGetApprovedUpgradeRequest } as QueryGetApprovedUpgradeRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    return message
  }
}

const baseQueryGetApprovedUpgradeResponse: object = {}

export const QueryGetApprovedUpgradeResponse = {
  encode(message: QueryGetApprovedUpgradeResponse, writer: Writer = Writer.create()): Writer {
    if (message.approvedUpgrade !== undefined) {
      ApprovedUpgrade.encode(message.approvedUpgrade, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedUpgradeResponse } as QueryGetApprovedUpgradeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.approvedUpgrade = ApprovedUpgrade.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetApprovedUpgradeResponse {
    const message = { ...baseQueryGetApprovedUpgradeResponse } as QueryGetApprovedUpgradeResponse
    if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
      message.approvedUpgrade = ApprovedUpgrade.fromJSON(object.approvedUpgrade)
    } else {
      message.approvedUpgrade = undefined
    }
    return message
  },

  toJSON(message: QueryGetApprovedUpgradeResponse): unknown {
    const obj: any = {}
    message.approvedUpgrade !== undefined && (obj.approvedUpgrade = message.approvedUpgrade ? ApprovedUpgrade.toJSON(message.approvedUpgrade) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetApprovedUpgradeResponse>): QueryGetApprovedUpgradeResponse {
    const message = { ...baseQueryGetApprovedUpgradeResponse } as QueryGetApprovedUpgradeResponse
    if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
      message.approvedUpgrade = ApprovedUpgrade.fromPartial(object.approvedUpgrade)
    } else {
      message.approvedUpgrade = undefined
    }
    return message
  }
}

const baseQueryAllApprovedUpgradeRequest: object = {}

export const QueryAllApprovedUpgradeRequest = {
  encode(message: QueryAllApprovedUpgradeRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllApprovedUpgradeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllApprovedUpgradeRequest } as QueryAllApprovedUpgradeRequest
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

  fromJSON(object: any): QueryAllApprovedUpgradeRequest {
    const message = { ...baseQueryAllApprovedUpgradeRequest } as QueryAllApprovedUpgradeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllApprovedUpgradeRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllApprovedUpgradeRequest>): QueryAllApprovedUpgradeRequest {
    const message = { ...baseQueryAllApprovedUpgradeRequest } as QueryAllApprovedUpgradeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllApprovedUpgradeResponse: object = {}

export const QueryAllApprovedUpgradeResponse = {
  encode(message: QueryAllApprovedUpgradeResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.approvedUpgrade) {
      ApprovedUpgrade.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllApprovedUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllApprovedUpgradeResponse } as QueryAllApprovedUpgradeResponse
    message.approvedUpgrade = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.approvedUpgrade.push(ApprovedUpgrade.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllApprovedUpgradeResponse {
    const message = { ...baseQueryAllApprovedUpgradeResponse } as QueryAllApprovedUpgradeResponse
    message.approvedUpgrade = []
    if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
      for (const e of object.approvedUpgrade) {
        message.approvedUpgrade.push(ApprovedUpgrade.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllApprovedUpgradeResponse): unknown {
    const obj: any = {}
    if (message.approvedUpgrade) {
      obj.approvedUpgrade = message.approvedUpgrade.map((e) => (e ? ApprovedUpgrade.toJSON(e) : undefined))
    } else {
      obj.approvedUpgrade = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllApprovedUpgradeResponse>): QueryAllApprovedUpgradeResponse {
    const message = { ...baseQueryAllApprovedUpgradeResponse } as QueryAllApprovedUpgradeResponse
    message.approvedUpgrade = []
    if (object.approvedUpgrade !== undefined && object.approvedUpgrade !== null) {
      for (const e of object.approvedUpgrade) {
        message.approvedUpgrade.push(ApprovedUpgrade.fromPartial(e))
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

const baseQueryGetRejectedUpgradeRequest: object = { name: '' }

export const QueryGetRejectedUpgradeRequest = {
  encode(message: QueryGetRejectedUpgradeRequest, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedUpgradeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRejectedUpgradeRequest } as QueryGetRejectedUpgradeRequest
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

  fromJSON(object: any): QueryGetRejectedUpgradeRequest {
    const message = { ...baseQueryGetRejectedUpgradeRequest } as QueryGetRejectedUpgradeRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    return message
  },

  toJSON(message: QueryGetRejectedUpgradeRequest): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRejectedUpgradeRequest>): QueryGetRejectedUpgradeRequest {
    const message = { ...baseQueryGetRejectedUpgradeRequest } as QueryGetRejectedUpgradeRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    return message
  }
}

const baseQueryGetRejectedUpgradeResponse: object = {}

export const QueryGetRejectedUpgradeResponse = {
  encode(message: QueryGetRejectedUpgradeResponse, writer: Writer = Writer.create()): Writer {
    if (message.rejectedUpgrade !== undefined) {
      RejectedUpgrade.encode(message.rejectedUpgrade, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRejectedUpgradeResponse } as QueryGetRejectedUpgradeResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.rejectedUpgrade = RejectedUpgrade.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRejectedUpgradeResponse {
    const message = { ...baseQueryGetRejectedUpgradeResponse } as QueryGetRejectedUpgradeResponse
    if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
      message.rejectedUpgrade = RejectedUpgrade.fromJSON(object.rejectedUpgrade)
    } else {
      message.rejectedUpgrade = undefined
    }
    return message
  },

  toJSON(message: QueryGetRejectedUpgradeResponse): unknown {
    const obj: any = {}
    message.rejectedUpgrade !== undefined && (obj.rejectedUpgrade = message.rejectedUpgrade ? RejectedUpgrade.toJSON(message.rejectedUpgrade) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRejectedUpgradeResponse>): QueryGetRejectedUpgradeResponse {
    const message = { ...baseQueryGetRejectedUpgradeResponse } as QueryGetRejectedUpgradeResponse
    if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
      message.rejectedUpgrade = RejectedUpgrade.fromPartial(object.rejectedUpgrade)
    } else {
      message.rejectedUpgrade = undefined
    }
    return message
  }
}

const baseQueryAllRejectedUpgradeRequest: object = {}

export const QueryAllRejectedUpgradeRequest = {
  encode(message: QueryAllRejectedUpgradeRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedUpgradeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRejectedUpgradeRequest } as QueryAllRejectedUpgradeRequest
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

  fromJSON(object: any): QueryAllRejectedUpgradeRequest {
    const message = { ...baseQueryAllRejectedUpgradeRequest } as QueryAllRejectedUpgradeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRejectedUpgradeRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRejectedUpgradeRequest>): QueryAllRejectedUpgradeRequest {
    const message = { ...baseQueryAllRejectedUpgradeRequest } as QueryAllRejectedUpgradeRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllRejectedUpgradeResponse: object = {}

export const QueryAllRejectedUpgradeResponse = {
  encode(message: QueryAllRejectedUpgradeResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.rejectedUpgrade) {
      RejectedUpgrade.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRejectedUpgradeResponse } as QueryAllRejectedUpgradeResponse
    message.rejectedUpgrade = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.rejectedUpgrade.push(RejectedUpgrade.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllRejectedUpgradeResponse {
    const message = { ...baseQueryAllRejectedUpgradeResponse } as QueryAllRejectedUpgradeResponse
    message.rejectedUpgrade = []
    if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
      for (const e of object.rejectedUpgrade) {
        message.rejectedUpgrade.push(RejectedUpgrade.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRejectedUpgradeResponse): unknown {
    const obj: any = {}
    if (message.rejectedUpgrade) {
      obj.rejectedUpgrade = message.rejectedUpgrade.map((e) => (e ? RejectedUpgrade.toJSON(e) : undefined))
    } else {
      obj.rejectedUpgrade = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRejectedUpgradeResponse>): QueryAllRejectedUpgradeResponse {
    const message = { ...baseQueryAllRejectedUpgradeResponse } as QueryAllRejectedUpgradeResponse
    message.rejectedUpgrade = []
    if (object.rejectedUpgrade !== undefined && object.rejectedUpgrade !== null) {
      for (const e of object.rejectedUpgrade) {
        message.rejectedUpgrade.push(RejectedUpgrade.fromPartial(e))
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
  /** Queries a ProposedUpgrade by index. */
  ProposedUpgrade(request: QueryGetProposedUpgradeRequest): Promise<QueryGetProposedUpgradeResponse>
  /** Queries a list of ProposedUpgrade items. */
  ProposedUpgradeAll(request: QueryAllProposedUpgradeRequest): Promise<QueryAllProposedUpgradeResponse>
  /** Queries a ApprovedUpgrade by index. */
  ApprovedUpgrade(request: QueryGetApprovedUpgradeRequest): Promise<QueryGetApprovedUpgradeResponse>
  /** Queries a list of ApprovedUpgrade items. */
  ApprovedUpgradeAll(request: QueryAllApprovedUpgradeRequest): Promise<QueryAllApprovedUpgradeResponse>
  /** Queries a RejectedUpgrade by index. */
  RejectedUpgrade(request: QueryGetRejectedUpgradeRequest): Promise<QueryGetRejectedUpgradeResponse>
  /** Queries a list of RejectedUpgrade items. */
  RejectedUpgradeAll(request: QueryAllRejectedUpgradeRequest): Promise<QueryAllRejectedUpgradeResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
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

  ApprovedUpgrade(request: QueryGetApprovedUpgradeRequest): Promise<QueryGetApprovedUpgradeResponse> {
    const data = QueryGetApprovedUpgradeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ApprovedUpgrade', data)
    return promise.then((data) => QueryGetApprovedUpgradeResponse.decode(new Reader(data)))
  }

  ApprovedUpgradeAll(request: QueryAllApprovedUpgradeRequest): Promise<QueryAllApprovedUpgradeResponse> {
    const data = QueryAllApprovedUpgradeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'ApprovedUpgradeAll', data)
    return promise.then((data) => QueryAllApprovedUpgradeResponse.decode(new Reader(data)))
  }

  RejectedUpgrade(request: QueryGetRejectedUpgradeRequest): Promise<QueryGetRejectedUpgradeResponse> {
    const data = QueryGetRejectedUpgradeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'RejectedUpgrade', data)
    return promise.then((data) => QueryGetRejectedUpgradeResponse.decode(new Reader(data)))
  }

  RejectedUpgradeAll(request: QueryAllRejectedUpgradeRequest): Promise<QueryAllRejectedUpgradeResponse> {
    const data = QueryAllRejectedUpgradeRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Query', 'RejectedUpgradeAll', data)
    return promise.then((data) => QueryAllRejectedUpgradeResponse.decode(new Reader(data)))
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
