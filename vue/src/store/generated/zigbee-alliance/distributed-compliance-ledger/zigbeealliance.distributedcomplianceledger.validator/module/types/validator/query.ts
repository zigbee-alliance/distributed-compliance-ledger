/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Validator } from '../validator/validator'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { LastValidatorPower } from '../validator/last_validator_power'
import { ProposedDisableValidator } from '../validator/proposed_disable_validator'
import { DisabledValidator } from '../validator/disabled_validator'
import { RejectedDisableValidator } from '../validator/rejected_validator'

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

export interface QueryGetProposedDisableValidatorRequest {
  address: string
}

export interface QueryGetProposedDisableValidatorResponse {
  proposedDisableValidator: ProposedDisableValidator | undefined
}

export interface QueryAllProposedDisableValidatorRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllProposedDisableValidatorResponse {
  proposedDisableValidator: ProposedDisableValidator[]
  pagination: PageResponse | undefined
}

export interface QueryGetDisabledValidatorRequest {
  address: string
}

export interface QueryGetDisabledValidatorResponse {
  disabledValidator: DisabledValidator | undefined
}

export interface QueryAllDisabledValidatorRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllDisabledValidatorResponse {
  disabledValidator: DisabledValidator[]
  pagination: PageResponse | undefined
}

export interface QueryGetRejectedDisableValidatorRequest {
  owner: string
}

export interface QueryGetRejectedDisableValidatorResponse {
  rejectedValidator: RejectedDisableValidator | undefined
}

export interface QueryAllRejectedDisableValidatorRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllRejectedDisableValidatorResponse {
  rejectedValidator: RejectedDisableValidator[]
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

const baseQueryGetProposedDisableValidatorRequest: object = { address: '' }

export const QueryGetProposedDisableValidatorRequest = {
  encode(message: QueryGetProposedDisableValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedDisableValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedDisableValidatorRequest } as QueryGetProposedDisableValidatorRequest
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

  fromJSON(object: any): QueryGetProposedDisableValidatorRequest {
    const message = { ...baseQueryGetProposedDisableValidatorRequest } as QueryGetProposedDisableValidatorRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetProposedDisableValidatorRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedDisableValidatorRequest>): QueryGetProposedDisableValidatorRequest {
    const message = { ...baseQueryGetProposedDisableValidatorRequest } as QueryGetProposedDisableValidatorRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetProposedDisableValidatorResponse: object = {}

export const QueryGetProposedDisableValidatorResponse = {
  encode(message: QueryGetProposedDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    if (message.proposedDisableValidator !== undefined) {
      ProposedDisableValidator.encode(message.proposedDisableValidator, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedDisableValidatorResponse } as QueryGetProposedDisableValidatorResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedDisableValidator = ProposedDisableValidator.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProposedDisableValidatorResponse {
    const message = { ...baseQueryGetProposedDisableValidatorResponse } as QueryGetProposedDisableValidatorResponse
    if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
      message.proposedDisableValidator = ProposedDisableValidator.fromJSON(object.proposedDisableValidator)
    } else {
      message.proposedDisableValidator = undefined
    }
    return message
  },

  toJSON(message: QueryGetProposedDisableValidatorResponse): unknown {
    const obj: any = {}
    message.proposedDisableValidator !== undefined &&
      (obj.proposedDisableValidator = message.proposedDisableValidator ? ProposedDisableValidator.toJSON(message.proposedDisableValidator) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedDisableValidatorResponse>): QueryGetProposedDisableValidatorResponse {
    const message = { ...baseQueryGetProposedDisableValidatorResponse } as QueryGetProposedDisableValidatorResponse
    if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
      message.proposedDisableValidator = ProposedDisableValidator.fromPartial(object.proposedDisableValidator)
    } else {
      message.proposedDisableValidator = undefined
    }
    return message
  }
}

const baseQueryAllProposedDisableValidatorRequest: object = {}

export const QueryAllProposedDisableValidatorRequest = {
  encode(message: QueryAllProposedDisableValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedDisableValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedDisableValidatorRequest } as QueryAllProposedDisableValidatorRequest
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

  fromJSON(object: any): QueryAllProposedDisableValidatorRequest {
    const message = { ...baseQueryAllProposedDisableValidatorRequest } as QueryAllProposedDisableValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedDisableValidatorRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedDisableValidatorRequest>): QueryAllProposedDisableValidatorRequest {
    const message = { ...baseQueryAllProposedDisableValidatorRequest } as QueryAllProposedDisableValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllProposedDisableValidatorResponse: object = {}

export const QueryAllProposedDisableValidatorResponse = {
  encode(message: QueryAllProposedDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.proposedDisableValidator) {
      ProposedDisableValidator.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedDisableValidatorResponse } as QueryAllProposedDisableValidatorResponse
    message.proposedDisableValidator = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedDisableValidator.push(ProposedDisableValidator.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllProposedDisableValidatorResponse {
    const message = { ...baseQueryAllProposedDisableValidatorResponse } as QueryAllProposedDisableValidatorResponse
    message.proposedDisableValidator = []
    if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
      for (const e of object.proposedDisableValidator) {
        message.proposedDisableValidator.push(ProposedDisableValidator.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedDisableValidatorResponse): unknown {
    const obj: any = {}
    if (message.proposedDisableValidator) {
      obj.proposedDisableValidator = message.proposedDisableValidator.map((e) => (e ? ProposedDisableValidator.toJSON(e) : undefined))
    } else {
      obj.proposedDisableValidator = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedDisableValidatorResponse>): QueryAllProposedDisableValidatorResponse {
    const message = { ...baseQueryAllProposedDisableValidatorResponse } as QueryAllProposedDisableValidatorResponse
    message.proposedDisableValidator = []
    if (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null) {
      for (const e of object.proposedDisableValidator) {
        message.proposedDisableValidator.push(ProposedDisableValidator.fromPartial(e))
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

const baseQueryGetDisabledValidatorRequest: object = { address: '' }

export const QueryGetDisabledValidatorRequest = {
  encode(message: QueryGetDisabledValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDisabledValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetDisabledValidatorRequest } as QueryGetDisabledValidatorRequest
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

  fromJSON(object: any): QueryGetDisabledValidatorRequest {
    const message = { ...baseQueryGetDisabledValidatorRequest } as QueryGetDisabledValidatorRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetDisabledValidatorRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetDisabledValidatorRequest>): QueryGetDisabledValidatorRequest {
    const message = { ...baseQueryGetDisabledValidatorRequest } as QueryGetDisabledValidatorRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetDisabledValidatorResponse: object = {}

export const QueryGetDisabledValidatorResponse = {
  encode(message: QueryGetDisabledValidatorResponse, writer: Writer = Writer.create()): Writer {
    if (message.disabledValidator !== undefined) {
      DisabledValidator.encode(message.disabledValidator, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDisabledValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetDisabledValidatorResponse } as QueryGetDisabledValidatorResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.disabledValidator = DisabledValidator.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetDisabledValidatorResponse {
    const message = { ...baseQueryGetDisabledValidatorResponse } as QueryGetDisabledValidatorResponse
    if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
      message.disabledValidator = DisabledValidator.fromJSON(object.disabledValidator)
    } else {
      message.disabledValidator = undefined
    }
    return message
  },

  toJSON(message: QueryGetDisabledValidatorResponse): unknown {
    const obj: any = {}
    message.disabledValidator !== undefined &&
      (obj.disabledValidator = message.disabledValidator ? DisabledValidator.toJSON(message.disabledValidator) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetDisabledValidatorResponse>): QueryGetDisabledValidatorResponse {
    const message = { ...baseQueryGetDisabledValidatorResponse } as QueryGetDisabledValidatorResponse
    if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
      message.disabledValidator = DisabledValidator.fromPartial(object.disabledValidator)
    } else {
      message.disabledValidator = undefined
    }
    return message
  }
}

const baseQueryAllDisabledValidatorRequest: object = {}

export const QueryAllDisabledValidatorRequest = {
  encode(message: QueryAllDisabledValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDisabledValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllDisabledValidatorRequest } as QueryAllDisabledValidatorRequest
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

  fromJSON(object: any): QueryAllDisabledValidatorRequest {
    const message = { ...baseQueryAllDisabledValidatorRequest } as QueryAllDisabledValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllDisabledValidatorRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllDisabledValidatorRequest>): QueryAllDisabledValidatorRequest {
    const message = { ...baseQueryAllDisabledValidatorRequest } as QueryAllDisabledValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllDisabledValidatorResponse: object = {}

export const QueryAllDisabledValidatorResponse = {
  encode(message: QueryAllDisabledValidatorResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.disabledValidator) {
      DisabledValidator.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDisabledValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllDisabledValidatorResponse } as QueryAllDisabledValidatorResponse
    message.disabledValidator = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.disabledValidator.push(DisabledValidator.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllDisabledValidatorResponse {
    const message = { ...baseQueryAllDisabledValidatorResponse } as QueryAllDisabledValidatorResponse
    message.disabledValidator = []
    if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
      for (const e of object.disabledValidator) {
        message.disabledValidator.push(DisabledValidator.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllDisabledValidatorResponse): unknown {
    const obj: any = {}
    if (message.disabledValidator) {
      obj.disabledValidator = message.disabledValidator.map((e) => (e ? DisabledValidator.toJSON(e) : undefined))
    } else {
      obj.disabledValidator = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllDisabledValidatorResponse>): QueryAllDisabledValidatorResponse {
    const message = { ...baseQueryAllDisabledValidatorResponse } as QueryAllDisabledValidatorResponse
    message.disabledValidator = []
    if (object.disabledValidator !== undefined && object.disabledValidator !== null) {
      for (const e of object.disabledValidator) {
        message.disabledValidator.push(DisabledValidator.fromPartial(e))
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

const baseQueryGetRejectedDisableValidatorRequest: object = { owner: '' }

export const QueryGetRejectedDisableValidatorRequest = {
  encode(message: QueryGetRejectedDisableValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.owner !== '') {
      writer.uint32(10).string(message.owner)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedDisableValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRejectedDisableValidatorRequest } as QueryGetRejectedDisableValidatorRequest
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

  fromJSON(object: any): QueryGetRejectedDisableValidatorRequest {
    const message = { ...baseQueryGetRejectedDisableValidatorRequest } as QueryGetRejectedDisableValidatorRequest
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner)
    } else {
      message.owner = ''
    }
    return message
  },

  toJSON(message: QueryGetRejectedDisableValidatorRequest): unknown {
    const obj: any = {}
    message.owner !== undefined && (obj.owner = message.owner)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRejectedDisableValidatorRequest>): QueryGetRejectedDisableValidatorRequest {
    const message = { ...baseQueryGetRejectedDisableValidatorRequest } as QueryGetRejectedDisableValidatorRequest
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner
    } else {
      message.owner = ''
    }
    return message
  }
}

const baseQueryGetRejectedDisableValidatorResponse: object = {}

export const QueryGetRejectedDisableValidatorResponse = {
  encode(message: QueryGetRejectedDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    if (message.rejectedValidator !== undefined) {
      RejectedDisableValidator.encode(message.rejectedValidator, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRejectedDisableValidatorResponse } as QueryGetRejectedDisableValidatorResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.rejectedValidator = RejectedDisableValidator.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRejectedDisableValidatorResponse {
    const message = { ...baseQueryGetRejectedDisableValidatorResponse } as QueryGetRejectedDisableValidatorResponse
    if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
      message.rejectedValidator = RejectedDisableValidator.fromJSON(object.rejectedValidator)
    } else {
      message.rejectedValidator = undefined
    }
    return message
  },

  toJSON(message: QueryGetRejectedDisableValidatorResponse): unknown {
    const obj: any = {}
    message.rejectedValidator !== undefined &&
      (obj.rejectedValidator = message.rejectedValidator ? RejectedDisableValidator.toJSON(message.rejectedValidator) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRejectedDisableValidatorResponse>): QueryGetRejectedDisableValidatorResponse {
    const message = { ...baseQueryGetRejectedDisableValidatorResponse } as QueryGetRejectedDisableValidatorResponse
    if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
      message.rejectedValidator = RejectedDisableValidator.fromPartial(object.rejectedValidator)
    } else {
      message.rejectedValidator = undefined
    }
    return message
  }
}

const baseQueryAllRejectedDisableValidatorRequest: object = {}

export const QueryAllRejectedDisableValidatorRequest = {
  encode(message: QueryAllRejectedDisableValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedDisableValidatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRejectedDisableValidatorRequest } as QueryAllRejectedDisableValidatorRequest
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

  fromJSON(object: any): QueryAllRejectedDisableValidatorRequest {
    const message = { ...baseQueryAllRejectedDisableValidatorRequest } as QueryAllRejectedDisableValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRejectedDisableValidatorRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRejectedDisableValidatorRequest>): QueryAllRejectedDisableValidatorRequest {
    const message = { ...baseQueryAllRejectedDisableValidatorRequest } as QueryAllRejectedDisableValidatorRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllRejectedDisableValidatorResponse: object = {}

export const QueryAllRejectedDisableValidatorResponse = {
  encode(message: QueryAllRejectedDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.rejectedValidator) {
      RejectedDisableValidator.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRejectedDisableValidatorResponse } as QueryAllRejectedDisableValidatorResponse
    message.rejectedValidator = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.rejectedValidator.push(RejectedDisableValidator.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllRejectedDisableValidatorResponse {
    const message = { ...baseQueryAllRejectedDisableValidatorResponse } as QueryAllRejectedDisableValidatorResponse
    message.rejectedValidator = []
    if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
      for (const e of object.rejectedValidator) {
        message.rejectedValidator.push(RejectedDisableValidator.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRejectedDisableValidatorResponse): unknown {
    const obj: any = {}
    if (message.rejectedValidator) {
      obj.rejectedValidator = message.rejectedValidator.map((e) => (e ? RejectedDisableValidator.toJSON(e) : undefined))
    } else {
      obj.rejectedValidator = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRejectedDisableValidatorResponse>): QueryAllRejectedDisableValidatorResponse {
    const message = { ...baseQueryAllRejectedDisableValidatorResponse } as QueryAllRejectedDisableValidatorResponse
    message.rejectedValidator = []
    if (object.rejectedValidator !== undefined && object.rejectedValidator !== null) {
      for (const e of object.rejectedValidator) {
        message.rejectedValidator.push(RejectedDisableValidator.fromPartial(e))
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
  /** Queries a ProposedDisableValidator by index. */
  ProposedDisableValidator(request: QueryGetProposedDisableValidatorRequest): Promise<QueryGetProposedDisableValidatorResponse>
  /** Queries a list of ProposedDisableValidator items. */
  ProposedDisableValidatorAll(request: QueryAllProposedDisableValidatorRequest): Promise<QueryAllProposedDisableValidatorResponse>
  /** Queries a DisabledValidator by index. */
  DisabledValidator(request: QueryGetDisabledValidatorRequest): Promise<QueryGetDisabledValidatorResponse>
  /** Queries a list of DisabledValidator items. */
  DisabledValidatorAll(request: QueryAllDisabledValidatorRequest): Promise<QueryAllDisabledValidatorResponse>
  /** Queries a RejectedNode by index. */
  RejectedDisableValidator(request: QueryGetRejectedDisableValidatorRequest): Promise<QueryGetRejectedDisableValidatorResponse>
  /** Queries a list of RejectedNode items. */
  RejectedDisableValidatorAll(request: QueryAllRejectedDisableValidatorRequest): Promise<QueryAllRejectedDisableValidatorResponse>
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

  ProposedDisableValidator(request: QueryGetProposedDisableValidatorRequest): Promise<QueryGetProposedDisableValidatorResponse> {
    const data = QueryGetProposedDisableValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ProposedDisableValidator', data)
    return promise.then((data) => QueryGetProposedDisableValidatorResponse.decode(new Reader(data)))
  }

  ProposedDisableValidatorAll(request: QueryAllProposedDisableValidatorRequest): Promise<QueryAllProposedDisableValidatorResponse> {
    const data = QueryAllProposedDisableValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ProposedDisableValidatorAll', data)
    return promise.then((data) => QueryAllProposedDisableValidatorResponse.decode(new Reader(data)))
  }

  DisabledValidator(request: QueryGetDisabledValidatorRequest): Promise<QueryGetDisabledValidatorResponse> {
    const data = QueryGetDisabledValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'DisabledValidator', data)
    return promise.then((data) => QueryGetDisabledValidatorResponse.decode(new Reader(data)))
  }

  DisabledValidatorAll(request: QueryAllDisabledValidatorRequest): Promise<QueryAllDisabledValidatorResponse> {
    const data = QueryAllDisabledValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'DisabledValidatorAll', data)
    return promise.then((data) => QueryAllDisabledValidatorResponse.decode(new Reader(data)))
  }

  RejectedDisableValidator(request: QueryGetRejectedDisableValidatorRequest): Promise<QueryGetRejectedDisableValidatorResponse> {
    const data = QueryGetRejectedDisableValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'RejectedDisableValidator', data)
    return promise.then((data) => QueryGetRejectedDisableValidatorResponse.decode(new Reader(data)))
  }

  RejectedDisableValidatorAll(request: QueryAllRejectedDisableValidatorRequest): Promise<QueryAllRejectedDisableValidatorResponse> {
    const data = QueryAllRejectedDisableValidatorRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'RejectedDisableValidatorAll', data)
    return promise.then((data) => QueryAllRejectedDisableValidatorResponse.decode(new Reader(data)))
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
