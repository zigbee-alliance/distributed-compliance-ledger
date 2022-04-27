/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { Any } from '../google/protobuf/any'
import { Description } from '../validator/description'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface MsgCreateValidator {
  signer: string
  pubKey: Any | undefined
  description: Description | undefined
}

export interface MsgCreateValidatorResponse {}

export interface MsgProposeDisableValidator {
  creator: string
  address: string
  info: string
  time: number
}

export interface MsgProposeDisableValidatorResponse {}

export interface MsgApproveDisableValidator {
  creator: string
  address: string
  info: string
  time: number
}

export interface MsgApproveDisableValidatorResponse {}

export interface MsgDisableValidator {
  creator: string
}

export interface MsgDisableValidatorResponse {}

export interface MsgEnableValidator {
  creator: string
}

export interface MsgEnableValidatorResponse {}

export interface MsgRejectDisableValidator {
  creator: string
  address: string
  info: string
  time: number
}

export interface MsgRejectDisableValidatorResponse {}

const baseMsgCreateValidator: object = { signer: '' }

export const MsgCreateValidator = {
  encode(message: MsgCreateValidator, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.pubKey !== undefined) {
      Any.encode(message.pubKey, writer.uint32(18).fork()).ldelim()
    }
    if (message.description !== undefined) {
      Description.encode(message.description, writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateValidator } as MsgCreateValidator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.pubKey = Any.decode(reader, reader.uint32())
          break
        case 3:
          message.description = Description.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateValidator {
    const message = { ...baseMsgCreateValidator } as MsgCreateValidator
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = Any.fromJSON(object.pubKey)
    } else {
      message.pubKey = undefined
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = Description.fromJSON(object.description)
    } else {
      message.description = undefined
    }
    return message
  },

  toJSON(message: MsgCreateValidator): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.pubKey !== undefined && (obj.pubKey = message.pubKey ? Any.toJSON(message.pubKey) : undefined)
    message.description !== undefined && (obj.description = message.description ? Description.toJSON(message.description) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateValidator>): MsgCreateValidator {
    const message = { ...baseMsgCreateValidator } as MsgCreateValidator
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = Any.fromPartial(object.pubKey)
    } else {
      message.pubKey = undefined
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = Description.fromPartial(object.description)
    } else {
      message.description = undefined
    }
    return message
  }
}

const baseMsgCreateValidatorResponse: object = {}

export const MsgCreateValidatorResponse = {
  encode(_: MsgCreateValidatorResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateValidatorResponse } as MsgCreateValidatorResponse
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

  fromJSON(_: any): MsgCreateValidatorResponse {
    const message = { ...baseMsgCreateValidatorResponse } as MsgCreateValidatorResponse
    return message
  },

  toJSON(_: MsgCreateValidatorResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateValidatorResponse>): MsgCreateValidatorResponse {
    const message = { ...baseMsgCreateValidatorResponse } as MsgCreateValidatorResponse
    return message
  }
}

const baseMsgProposeDisableValidator: object = { creator: '', address: '', info: '', time: 0 }

export const MsgProposeDisableValidator = {
  encode(message: MsgProposeDisableValidator, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeDisableValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeDisableValidator } as MsgProposeDisableValidator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        case 3:
          message.info = reader.string()
          break
        case 4:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgProposeDisableValidator {
    const message = { ...baseMsgProposeDisableValidator } as MsgProposeDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgProposeDisableValidator): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.address !== undefined && (obj.address = message.address)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgProposeDisableValidator>): MsgProposeDisableValidator {
    const message = { ...baseMsgProposeDisableValidator } as MsgProposeDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgProposeDisableValidatorResponse: object = {}

export const MsgProposeDisableValidatorResponse = {
  encode(_: MsgProposeDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeDisableValidatorResponse } as MsgProposeDisableValidatorResponse
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

  fromJSON(_: any): MsgProposeDisableValidatorResponse {
    const message = { ...baseMsgProposeDisableValidatorResponse } as MsgProposeDisableValidatorResponse
    return message
  },

  toJSON(_: MsgProposeDisableValidatorResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgProposeDisableValidatorResponse>): MsgProposeDisableValidatorResponse {
    const message = { ...baseMsgProposeDisableValidatorResponse } as MsgProposeDisableValidatorResponse
    return message
  }
}

const baseMsgApproveDisableValidator: object = { creator: '', address: '', info: '', time: 0 }

export const MsgApproveDisableValidator = {
  encode(message: MsgApproveDisableValidator, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveDisableValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveDisableValidator } as MsgApproveDisableValidator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        case 3:
          message.info = reader.string()
          break
        case 4:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgApproveDisableValidator {
    const message = { ...baseMsgApproveDisableValidator } as MsgApproveDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgApproveDisableValidator): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.address !== undefined && (obj.address = message.address)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgApproveDisableValidator>): MsgApproveDisableValidator {
    const message = { ...baseMsgApproveDisableValidator } as MsgApproveDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgApproveDisableValidatorResponse: object = {}

export const MsgApproveDisableValidatorResponse = {
  encode(_: MsgApproveDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveDisableValidatorResponse } as MsgApproveDisableValidatorResponse
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

  fromJSON(_: any): MsgApproveDisableValidatorResponse {
    const message = { ...baseMsgApproveDisableValidatorResponse } as MsgApproveDisableValidatorResponse
    return message
  },

  toJSON(_: MsgApproveDisableValidatorResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgApproveDisableValidatorResponse>): MsgApproveDisableValidatorResponse {
    const message = { ...baseMsgApproveDisableValidatorResponse } as MsgApproveDisableValidatorResponse
    return message
  }
}

const baseMsgDisableValidator: object = { creator: '' }

export const MsgDisableValidator = {
  encode(message: MsgDisableValidator, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDisableValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDisableValidator } as MsgDisableValidator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDisableValidator {
    const message = { ...baseMsgDisableValidator } as MsgDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    return message
  },

  toJSON(message: MsgDisableValidator): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDisableValidator>): MsgDisableValidator {
    const message = { ...baseMsgDisableValidator } as MsgDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    return message
  }
}

const baseMsgDisableValidatorResponse: object = {}

export const MsgDisableValidatorResponse = {
  encode(_: MsgDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDisableValidatorResponse } as MsgDisableValidatorResponse
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

  fromJSON(_: any): MsgDisableValidatorResponse {
    const message = { ...baseMsgDisableValidatorResponse } as MsgDisableValidatorResponse
    return message
  },

  toJSON(_: MsgDisableValidatorResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDisableValidatorResponse>): MsgDisableValidatorResponse {
    const message = { ...baseMsgDisableValidatorResponse } as MsgDisableValidatorResponse
    return message
  }
}

const baseMsgEnableValidator: object = { creator: '' }

export const MsgEnableValidator = {
  encode(message: MsgEnableValidator, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEnableValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgEnableValidator } as MsgEnableValidator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgEnableValidator {
    const message = { ...baseMsgEnableValidator } as MsgEnableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    return message
  },

  toJSON(message: MsgEnableValidator): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    return obj
  },

  fromPartial(object: DeepPartial<MsgEnableValidator>): MsgEnableValidator {
    const message = { ...baseMsgEnableValidator } as MsgEnableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    return message
  }
}

const baseMsgEnableValidatorResponse: object = {}

export const MsgEnableValidatorResponse = {
  encode(_: MsgEnableValidatorResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEnableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgEnableValidatorResponse } as MsgEnableValidatorResponse
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

  fromJSON(_: any): MsgEnableValidatorResponse {
    const message = { ...baseMsgEnableValidatorResponse } as MsgEnableValidatorResponse
    return message
  },

  toJSON(_: MsgEnableValidatorResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgEnableValidatorResponse>): MsgEnableValidatorResponse {
    const message = { ...baseMsgEnableValidatorResponse } as MsgEnableValidatorResponse
    return message
  }
}

const baseMsgRejectDisableValidator: object = { creator: '', address: '', info: '', time: 0 }

export const MsgRejectDisableValidator = {
  encode(message: MsgRejectDisableValidator, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.address !== '') {
      writer.uint32(18).string(message.address)
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectDisableValidator {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRejectDisableValidator } as MsgRejectDisableValidator
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.address = reader.string()
          break
        case 3:
          message.info = reader.string()
          break
        case 4:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgRejectDisableValidator {
    const message = { ...baseMsgRejectDisableValidator } as MsgRejectDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address)
    } else {
      message.address = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgRejectDisableValidator): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.address !== undefined && (obj.address = message.address)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgRejectDisableValidator>): MsgRejectDisableValidator {
    const message = { ...baseMsgRejectDisableValidator } as MsgRejectDisableValidator
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address
    } else {
      message.address = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgRejectDisableValidatorResponse: object = {}

export const MsgRejectDisableValidatorResponse = {
  encode(_: MsgRejectDisableValidatorResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectDisableValidatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRejectDisableValidatorResponse } as MsgRejectDisableValidatorResponse
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

  fromJSON(_: any): MsgRejectDisableValidatorResponse {
    const message = { ...baseMsgRejectDisableValidatorResponse } as MsgRejectDisableValidatorResponse
    return message
  },

  toJSON(_: MsgRejectDisableValidatorResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgRejectDisableValidatorResponse>): MsgRejectDisableValidatorResponse {
    const message = { ...baseMsgRejectDisableValidatorResponse } as MsgRejectDisableValidatorResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse>
  ProposeDisableValidator(request: MsgProposeDisableValidator): Promise<MsgProposeDisableValidatorResponse>
  ApproveDisableValidator(request: MsgApproveDisableValidator): Promise<MsgApproveDisableValidatorResponse>
  DisableValidator(request: MsgDisableValidator): Promise<MsgDisableValidatorResponse>
  EnableValidator(request: MsgEnableValidator): Promise<MsgEnableValidatorResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RejectDisableValidator(request: MsgRejectDisableValidator): Promise<MsgRejectDisableValidatorResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse> {
    const data = MsgCreateValidator.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'CreateValidator', data)
    return promise.then((data) => MsgCreateValidatorResponse.decode(new Reader(data)))
  }

  ProposeDisableValidator(request: MsgProposeDisableValidator): Promise<MsgProposeDisableValidatorResponse> {
    const data = MsgProposeDisableValidator.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'ProposeDisableValidator', data)
    return promise.then((data) => MsgProposeDisableValidatorResponse.decode(new Reader(data)))
  }

  ApproveDisableValidator(request: MsgApproveDisableValidator): Promise<MsgApproveDisableValidatorResponse> {
    const data = MsgApproveDisableValidator.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'ApproveDisableValidator', data)
    return promise.then((data) => MsgApproveDisableValidatorResponse.decode(new Reader(data)))
  }

  DisableValidator(request: MsgDisableValidator): Promise<MsgDisableValidatorResponse> {
    const data = MsgDisableValidator.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'DisableValidator', data)
    return promise.then((data) => MsgDisableValidatorResponse.decode(new Reader(data)))
  }

  EnableValidator(request: MsgEnableValidator): Promise<MsgEnableValidatorResponse> {
    const data = MsgEnableValidator.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'EnableValidator', data)
    return promise.then((data) => MsgEnableValidatorResponse.decode(new Reader(data)))
  }

  RejectDisableValidator(request: MsgRejectDisableValidator): Promise<MsgRejectDisableValidatorResponse> {
    const data = MsgRejectDisableValidator.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.validator.Msg', 'RejectDisableValidator', data)
    return promise.then((data) => MsgRejectDisableValidatorResponse.decode(new Reader(data)))
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
