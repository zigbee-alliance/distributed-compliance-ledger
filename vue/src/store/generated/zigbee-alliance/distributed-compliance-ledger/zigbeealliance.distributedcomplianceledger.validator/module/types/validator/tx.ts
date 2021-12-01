/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { Any } from '../google/protobuf/any'
import { Description } from '../validator/description'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface MsgCreateValidator {
  signer: string
  pubKey: Any | undefined
  description: Description | undefined
}

export interface MsgCreateValidatorResponse {}

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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateValidator(request: MsgCreateValidator): Promise<MsgCreateValidatorResponse>
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
