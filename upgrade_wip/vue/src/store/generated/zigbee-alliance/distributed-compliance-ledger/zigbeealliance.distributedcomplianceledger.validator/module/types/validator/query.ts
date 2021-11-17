/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { Validator } from '../validator/validator'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { LastValidatorPower } from '../validator/last_validator_power'
import { ValidatorSigningInfo } from '../validator/validator_signing_info'
import { ValidatorMissedBlockBitArray } from '../validator/validator_missed_block_bit_array'
import { ValidatorOwner } from '../validator/validator_owner'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface QueryGetValidatorRequest {
  address: string
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
  consensusAddress: string
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

export interface QueryGetValidatorSigningInfoRequest {
  address: string
}

export interface QueryGetValidatorSigningInfoResponse {
  validatorSigningInfo: ValidatorSigningInfo | undefined
}

export interface QueryAllValidatorSigningInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllValidatorSigningInfoResponse {
  validatorSigningInfo: ValidatorSigningInfo[]
  pagination: PageResponse | undefined
}

export interface QueryGetValidatorMissedBlockBitArrayRequest {
  address: string
  index: number
}

export interface QueryGetValidatorMissedBlockBitArrayResponse {
  validatorMissedBlockBitArray: ValidatorMissedBlockBitArray | undefined
}

export interface QueryAllValidatorMissedBlockBitArrayRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllValidatorMissedBlockBitArrayResponse {
  validatorMissedBlockBitArray: ValidatorMissedBlockBitArray[]
  pagination: PageResponse | undefined
}

export interface QueryGetValidatorOwnerRequest {
  address: string
}

export interface QueryGetValidatorOwnerResponse {
  validatorOwner: ValidatorOwner | undefined
}

export interface QueryAllValidatorOwnerRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllValidatorOwnerResponse {
  validatorOwner: ValidatorOwner[]
  pagination: PageResponse | undefined
}

const baseQueryGetValidatorRequest: object = { address: '' }

export const QueryGetValidatorRequest = {
  encode(message: QueryGetValidatorRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
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
          message.address = reader.string()
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
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetValidatorRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorRequest>): QueryGetValidatorRequest {
    const message = { ...baseQueryGetValidatorRequest } as QueryGetValidatorRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
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

const baseQueryGetLastValidatorPowerRequest: object = { consensusAddress: '' }

export const QueryGetLastValidatorPowerRequest = {
  encode(message: QueryGetLastValidatorPowerRequest, writer: Writer = Writer.create()): Writer {
    if (message.consensusAddress !== '') {
      writer.uint32(10).string(message.consensusAddress)
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
          message.consensusAddress = reader.string()
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
    if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
      message.consensusAddress = String(object.consensusAddress)
    } else {
      message.consensusAddress = ''
    }
    return message
  },

  toJSON(message: QueryGetLastValidatorPowerRequest): unknown {
    const obj: any = {}
    message.consensusAddress !== undefined && (obj.consensusAddress = message.consensusAddress)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetLastValidatorPowerRequest>): QueryGetLastValidatorPowerRequest {
    const message = { ...baseQueryGetLastValidatorPowerRequest } as QueryGetLastValidatorPowerRequest
    if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
      message.consensusAddress = object.consensusAddress
    } else {
      message.consensusAddress = ''
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

const baseQueryGetValidatorSigningInfoRequest: object = { address: '' }

export const QueryGetValidatorSigningInfoRequest = {
  encode(message: QueryGetValidatorSigningInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorSigningInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorSigningInfoRequest } as QueryGetValidatorSigningInfoRequest
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

  fromJSON(object: any): QueryGetValidatorSigningInfoRequest {
    const message = { ...baseQueryGetValidatorSigningInfoRequest } as QueryGetValidatorSigningInfoRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetValidatorSigningInfoRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorSigningInfoRequest>): QueryGetValidatorSigningInfoRequest {
    const message = { ...baseQueryGetValidatorSigningInfoRequest } as QueryGetValidatorSigningInfoRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetValidatorSigningInfoResponse: object = {}

export const QueryGetValidatorSigningInfoResponse = {
  encode(message: QueryGetValidatorSigningInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.validatorSigningInfo !== undefined) {
      ValidatorSigningInfo.encode(message.validatorSigningInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorSigningInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorSigningInfoResponse } as QueryGetValidatorSigningInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorSigningInfo = ValidatorSigningInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetValidatorSigningInfoResponse {
    const message = { ...baseQueryGetValidatorSigningInfoResponse } as QueryGetValidatorSigningInfoResponse
    if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
      message.validatorSigningInfo = ValidatorSigningInfo.fromJSON(object.validatorSigningInfo)
    } else {
      message.validatorSigningInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetValidatorSigningInfoResponse): unknown {
    const obj: any = {}
    message.validatorSigningInfo !== undefined &&
      (obj.validatorSigningInfo = message.validatorSigningInfo ? ValidatorSigningInfo.toJSON(message.validatorSigningInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorSigningInfoResponse>): QueryGetValidatorSigningInfoResponse {
    const message = { ...baseQueryGetValidatorSigningInfoResponse } as QueryGetValidatorSigningInfoResponse
    if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
      message.validatorSigningInfo = ValidatorSigningInfo.fromPartial(object.validatorSigningInfo)
    } else {
      message.validatorSigningInfo = undefined
    }
    return message
  }
}

const baseQueryAllValidatorSigningInfoRequest: object = {}

export const QueryAllValidatorSigningInfoRequest = {
  encode(message: QueryAllValidatorSigningInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorSigningInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorSigningInfoRequest } as QueryAllValidatorSigningInfoRequest
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

  fromJSON(object: any): QueryAllValidatorSigningInfoRequest {
    const message = { ...baseQueryAllValidatorSigningInfoRequest } as QueryAllValidatorSigningInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorSigningInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorSigningInfoRequest>): QueryAllValidatorSigningInfoRequest {
    const message = { ...baseQueryAllValidatorSigningInfoRequest } as QueryAllValidatorSigningInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllValidatorSigningInfoResponse: object = {}

export const QueryAllValidatorSigningInfoResponse = {
  encode(message: QueryAllValidatorSigningInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.validatorSigningInfo) {
      ValidatorSigningInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorSigningInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorSigningInfoResponse } as QueryAllValidatorSigningInfoResponse
    message.validatorSigningInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorSigningInfo.push(ValidatorSigningInfo.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllValidatorSigningInfoResponse {
    const message = { ...baseQueryAllValidatorSigningInfoResponse } as QueryAllValidatorSigningInfoResponse
    message.validatorSigningInfo = []
    if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
      for (const e of object.validatorSigningInfo) {
        message.validatorSigningInfo.push(ValidatorSigningInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorSigningInfoResponse): unknown {
    const obj: any = {}
    if (message.validatorSigningInfo) {
      obj.validatorSigningInfo = message.validatorSigningInfo.map((e) => (e ? ValidatorSigningInfo.toJSON(e) : undefined))
    } else {
      obj.validatorSigningInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorSigningInfoResponse>): QueryAllValidatorSigningInfoResponse {
    const message = { ...baseQueryAllValidatorSigningInfoResponse } as QueryAllValidatorSigningInfoResponse
    message.validatorSigningInfo = []
    if (object.validatorSigningInfo !== undefined && object.validatorSigningInfo !== null) {
      for (const e of object.validatorSigningInfo) {
        message.validatorSigningInfo.push(ValidatorSigningInfo.fromPartial(e))
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

const baseQueryGetValidatorMissedBlockBitArrayRequest: object = { address: '', index: 0 }

export const QueryGetValidatorMissedBlockBitArrayRequest = {
  encode(message: QueryGetValidatorMissedBlockBitArrayRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorMissedBlockBitArrayRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorMissedBlockBitArrayRequest } as QueryGetValidatorMissedBlockBitArrayRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string()
          break
        case 2:
          message.index = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetValidatorMissedBlockBitArrayRequest {
    const message = { ...baseQueryGetValidatorMissedBlockBitArrayRequest } as QueryGetValidatorMissedBlockBitArrayRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = Number(object.index)
    } else {
      message.index = 0
    }
    return message
  },

  toJSON(message: QueryGetValidatorMissedBlockBitArrayRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorMissedBlockBitArrayRequest>): QueryGetValidatorMissedBlockBitArrayRequest {
    const message = { ...baseQueryGetValidatorMissedBlockBitArrayRequest } as QueryGetValidatorMissedBlockBitArrayRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = 0
    }
    return message
  }
}

const baseQueryGetValidatorMissedBlockBitArrayResponse: object = {}

export const QueryGetValidatorMissedBlockBitArrayResponse = {
  encode(message: QueryGetValidatorMissedBlockBitArrayResponse, writer: Writer = Writer.create()): Writer {
    if (message.validatorMissedBlockBitArray !== undefined) {
      ValidatorMissedBlockBitArray.encode(message.validatorMissedBlockBitArray, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorMissedBlockBitArrayResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorMissedBlockBitArrayResponse } as QueryGetValidatorMissedBlockBitArrayResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorMissedBlockBitArray = ValidatorMissedBlockBitArray.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetValidatorMissedBlockBitArrayResponse {
    const message = { ...baseQueryGetValidatorMissedBlockBitArrayResponse } as QueryGetValidatorMissedBlockBitArrayResponse
    if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
      message.validatorMissedBlockBitArray = ValidatorMissedBlockBitArray.fromJSON(object.validatorMissedBlockBitArray)
    } else {
      message.validatorMissedBlockBitArray = undefined
    }
    return message
  },

  toJSON(message: QueryGetValidatorMissedBlockBitArrayResponse): unknown {
    const obj: any = {}
    message.validatorMissedBlockBitArray !== undefined &&
      (obj.validatorMissedBlockBitArray = message.validatorMissedBlockBitArray
        ? ValidatorMissedBlockBitArray.toJSON(message.validatorMissedBlockBitArray)
        : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorMissedBlockBitArrayResponse>): QueryGetValidatorMissedBlockBitArrayResponse {
    const message = { ...baseQueryGetValidatorMissedBlockBitArrayResponse } as QueryGetValidatorMissedBlockBitArrayResponse
    if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
      message.validatorMissedBlockBitArray = ValidatorMissedBlockBitArray.fromPartial(object.validatorMissedBlockBitArray)
    } else {
      message.validatorMissedBlockBitArray = undefined
    }
    return message
  }
}

const baseQueryAllValidatorMissedBlockBitArrayRequest: object = {}

export const QueryAllValidatorMissedBlockBitArrayRequest = {
  encode(message: QueryAllValidatorMissedBlockBitArrayRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorMissedBlockBitArrayRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorMissedBlockBitArrayRequest } as QueryAllValidatorMissedBlockBitArrayRequest
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

  fromJSON(object: any): QueryAllValidatorMissedBlockBitArrayRequest {
    const message = { ...baseQueryAllValidatorMissedBlockBitArrayRequest } as QueryAllValidatorMissedBlockBitArrayRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorMissedBlockBitArrayRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorMissedBlockBitArrayRequest>): QueryAllValidatorMissedBlockBitArrayRequest {
    const message = { ...baseQueryAllValidatorMissedBlockBitArrayRequest } as QueryAllValidatorMissedBlockBitArrayRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllValidatorMissedBlockBitArrayResponse: object = {}

export const QueryAllValidatorMissedBlockBitArrayResponse = {
  encode(message: QueryAllValidatorMissedBlockBitArrayResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.validatorMissedBlockBitArray) {
      ValidatorMissedBlockBitArray.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorMissedBlockBitArrayResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorMissedBlockBitArrayResponse } as QueryAllValidatorMissedBlockBitArrayResponse
    message.validatorMissedBlockBitArray = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorMissedBlockBitArray.push(ValidatorMissedBlockBitArray.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllValidatorMissedBlockBitArrayResponse {
    const message = { ...baseQueryAllValidatorMissedBlockBitArrayResponse } as QueryAllValidatorMissedBlockBitArrayResponse
    message.validatorMissedBlockBitArray = []
    if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
      for (const e of object.validatorMissedBlockBitArray) {
        message.validatorMissedBlockBitArray.push(ValidatorMissedBlockBitArray.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorMissedBlockBitArrayResponse): unknown {
    const obj: any = {}
    if (message.validatorMissedBlockBitArray) {
      obj.validatorMissedBlockBitArray = message.validatorMissedBlockBitArray.map((e) => (e ? ValidatorMissedBlockBitArray.toJSON(e) : undefined))
    } else {
      obj.validatorMissedBlockBitArray = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorMissedBlockBitArrayResponse>): QueryAllValidatorMissedBlockBitArrayResponse {
    const message = { ...baseQueryAllValidatorMissedBlockBitArrayResponse } as QueryAllValidatorMissedBlockBitArrayResponse
    message.validatorMissedBlockBitArray = []
    if (object.validatorMissedBlockBitArray !== undefined && object.validatorMissedBlockBitArray !== null) {
      for (const e of object.validatorMissedBlockBitArray) {
        message.validatorMissedBlockBitArray.push(ValidatorMissedBlockBitArray.fromPartial(e))
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

const baseQueryGetValidatorOwnerRequest: object = { address: '' }

export const QueryGetValidatorOwnerRequest = {
  encode(message: QueryGetValidatorOwnerRequest, writer: Writer = Writer.create()): Writer {
    if (message.address !== '') {
      writer.uint32(10).string(message.address)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorOwnerRequest } as QueryGetValidatorOwnerRequest
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

  fromJSON(object: any): QueryGetValidatorOwnerRequest {
    const message = { ...baseQueryGetValidatorOwnerRequest } as QueryGetValidatorOwnerRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    return message
  },

  toJSON(message: QueryGetValidatorOwnerRequest): unknown {
    const obj: any = {}
    message.address !== undefined && (obj.address = message.address)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorOwnerRequest>): QueryGetValidatorOwnerRequest {
    const message = { ...baseQueryGetValidatorOwnerRequest } as QueryGetValidatorOwnerRequest
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    return message
  }
}

const baseQueryGetValidatorOwnerResponse: object = {}

export const QueryGetValidatorOwnerResponse = {
  encode(message: QueryGetValidatorOwnerResponse, writer: Writer = Writer.create()): Writer {
    if (message.validatorOwner !== undefined) {
      ValidatorOwner.encode(message.validatorOwner, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetValidatorOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetValidatorOwnerResponse } as QueryGetValidatorOwnerResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorOwner = ValidatorOwner.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetValidatorOwnerResponse {
    const message = { ...baseQueryGetValidatorOwnerResponse } as QueryGetValidatorOwnerResponse
    if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
      message.validatorOwner = ValidatorOwner.fromJSON(object.validatorOwner)
    } else {
      message.validatorOwner = undefined
    }
    return message
  },

  toJSON(message: QueryGetValidatorOwnerResponse): unknown {
    const obj: any = {}
    message.validatorOwner !== undefined && (obj.validatorOwner = message.validatorOwner ? ValidatorOwner.toJSON(message.validatorOwner) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetValidatorOwnerResponse>): QueryGetValidatorOwnerResponse {
    const message = { ...baseQueryGetValidatorOwnerResponse } as QueryGetValidatorOwnerResponse
    if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
      message.validatorOwner = ValidatorOwner.fromPartial(object.validatorOwner)
    } else {
      message.validatorOwner = undefined
    }
    return message
  }
}

const baseQueryAllValidatorOwnerRequest: object = {}

export const QueryAllValidatorOwnerRequest = {
  encode(message: QueryAllValidatorOwnerRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorOwnerRequest } as QueryAllValidatorOwnerRequest
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

  fromJSON(object: any): QueryAllValidatorOwnerRequest {
    const message = { ...baseQueryAllValidatorOwnerRequest } as QueryAllValidatorOwnerRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorOwnerRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorOwnerRequest>): QueryAllValidatorOwnerRequest {
    const message = { ...baseQueryAllValidatorOwnerRequest } as QueryAllValidatorOwnerRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllValidatorOwnerResponse: object = {}

export const QueryAllValidatorOwnerResponse = {
  encode(message: QueryAllValidatorOwnerResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.validatorOwner) {
      ValidatorOwner.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllValidatorOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllValidatorOwnerResponse } as QueryAllValidatorOwnerResponse
    message.validatorOwner = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorOwner.push(ValidatorOwner.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllValidatorOwnerResponse {
    const message = { ...baseQueryAllValidatorOwnerResponse } as QueryAllValidatorOwnerResponse
    message.validatorOwner = []
    if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
      for (const e of object.validatorOwner) {
        message.validatorOwner.push(ValidatorOwner.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllValidatorOwnerResponse): unknown {
    const obj: any = {}
    if (message.validatorOwner) {
      obj.validatorOwner = message.validatorOwner.map((e) => (e ? ValidatorOwner.toJSON(e) : undefined))
    } else {
      obj.validatorOwner = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllValidatorOwnerResponse>): QueryAllValidatorOwnerResponse {
    const message = { ...baseQueryAllValidatorOwnerResponse } as QueryAllValidatorOwnerResponse
    message.validatorOwner = []
    if (object.validatorOwner !== undefined && object.validatorOwner !== null) {
      for (const e of object.validatorOwner) {
        message.validatorOwner.push(ValidatorOwner.fromPartial(e))
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
  /** Queries a validatorSigningInfo by index. */
  ValidatorSigningInfo(request: QueryGetValidatorSigningInfoRequest): Promise<QueryGetValidatorSigningInfoResponse>
  /** Queries a list of validatorSigningInfo items. */
  ValidatorSigningInfoAll(request: QueryAllValidatorSigningInfoRequest): Promise<QueryAllValidatorSigningInfoResponse>
  /** Queries a validatorMissedBlockBitArray by index. */
  ValidatorMissedBlockBitArray(request: QueryGetValidatorMissedBlockBitArrayRequest): Promise<QueryGetValidatorMissedBlockBitArrayResponse>
  /** Queries a list of validatorMissedBlockBitArray items. */
  ValidatorMissedBlockBitArrayAll(request: QueryAllValidatorMissedBlockBitArrayRequest): Promise<QueryAllValidatorMissedBlockBitArrayResponse>
  /** Queries a validatorOwner by index. */
  ValidatorOwner(request: QueryGetValidatorOwnerRequest): Promise<QueryGetValidatorOwnerResponse>
  /** Queries a list of validatorOwner items. */
  ValidatorOwnerAll(request: QueryAllValidatorOwnerRequest): Promise<QueryAllValidatorOwnerResponse>
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

  ValidatorSigningInfo(request: QueryGetValidatorSigningInfoRequest): Promise<QueryGetValidatorSigningInfoResponse> {
    const data = QueryGetValidatorSigningInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorSigningInfo', data)
    return promise.then((data) => QueryGetValidatorSigningInfoResponse.decode(new Reader(data)))
  }

  ValidatorSigningInfoAll(request: QueryAllValidatorSigningInfoRequest): Promise<QueryAllValidatorSigningInfoResponse> {
    const data = QueryAllValidatorSigningInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorSigningInfoAll', data)
    return promise.then((data) => QueryAllValidatorSigningInfoResponse.decode(new Reader(data)))
  }

  ValidatorMissedBlockBitArray(request: QueryGetValidatorMissedBlockBitArrayRequest): Promise<QueryGetValidatorMissedBlockBitArrayResponse> {
    const data = QueryGetValidatorMissedBlockBitArrayRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorMissedBlockBitArray', data)
    return promise.then((data) => QueryGetValidatorMissedBlockBitArrayResponse.decode(new Reader(data)))
  }

  ValidatorMissedBlockBitArrayAll(request: QueryAllValidatorMissedBlockBitArrayRequest): Promise<QueryAllValidatorMissedBlockBitArrayResponse> {
    const data = QueryAllValidatorMissedBlockBitArrayRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorMissedBlockBitArrayAll', data)
    return promise.then((data) => QueryAllValidatorMissedBlockBitArrayResponse.decode(new Reader(data)))
  }

  ValidatorOwner(request: QueryGetValidatorOwnerRequest): Promise<QueryGetValidatorOwnerResponse> {
    const data = QueryGetValidatorOwnerRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorOwner', data)
    return promise.then((data) => QueryGetValidatorOwnerResponse.decode(new Reader(data)))
  }

  ValidatorOwnerAll(request: QueryAllValidatorOwnerRequest): Promise<QueryAllValidatorOwnerResponse> {
    const data = QueryAllValidatorOwnerRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Query', 'ValidatorOwnerAll', data)
    return promise.then((data) => QueryAllValidatorOwnerResponse.decode(new Reader(data)))
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
