/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface MsgCreateVendorInfoType {
  creator: string
  vendorID: number
  vendorName: string
  companyLegalName: string
  companyPrefferedName: string
  vendorLandingPageURL: string
}

export interface MsgCreateVendorInfoTypeResponse {}

export interface MsgUpdateVendorInfoType {
  creator: string
  vendorID: number
  vendorName: string
  companyLegalName: string
  companyPrefferedName: string
  vendorLandingPageURL: string
}

export interface MsgUpdateVendorInfoTypeResponse {}

export interface MsgDeleteVendorInfoType {
  creator: string
  vendorID: number
}

export interface MsgDeleteVendorInfoTypeResponse {}

const baseMsgCreateVendorInfoType: object = {
  creator: '',
  vendorID: 0,
  vendorName: '',
  companyLegalName: '',
  companyPrefferedName: '',
  vendorLandingPageURL: ''
}

export const MsgCreateVendorInfoType = {
  encode(message: MsgCreateVendorInfoType, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).uint64(message.vendorID)
    }
    if (message.vendorName !== '') {
      writer.uint32(26).string(message.vendorName)
    }
    if (message.companyLegalName !== '') {
      writer.uint32(34).string(message.companyLegalName)
    }
    if (message.companyPrefferedName !== '') {
      writer.uint32(42).string(message.companyPrefferedName)
    }
    if (message.vendorLandingPageURL !== '') {
      writer.uint32(50).string(message.vendorLandingPageURL)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfoType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateVendorInfoType } as MsgCreateVendorInfoType
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.vendorID = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.vendorName = reader.string()
          break
        case 4:
          message.companyLegalName = reader.string()
          break
        case 5:
          message.companyPrefferedName = reader.string()
          break
        case 6:
          message.vendorLandingPageURL = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateVendorInfoType {
    const message = { ...baseMsgCreateVendorInfoType } as MsgCreateVendorInfoType
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    if (object.vendorName !== undefined && object.vendorName !== null) {
      message.vendorName = String(object.vendorName)
    } else {
      message.vendorName = ''
    }
    if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
      message.companyLegalName = String(object.companyLegalName)
    } else {
      message.companyLegalName = ''
    }
    if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
      message.companyPrefferedName = String(object.companyPrefferedName)
    } else {
      message.companyPrefferedName = ''
    }
    if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
      message.vendorLandingPageURL = String(object.vendorLandingPageURL)
    } else {
      message.vendorLandingPageURL = ''
    }
    return message
  },

  toJSON(message: MsgCreateVendorInfoType): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    message.vendorName !== undefined && (obj.vendorName = message.vendorName)
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName)
    message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName)
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateVendorInfoType>): MsgCreateVendorInfoType {
    const message = { ...baseMsgCreateVendorInfoType } as MsgCreateVendorInfoType
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
    }
    if (object.vendorName !== undefined && object.vendorName !== null) {
      message.vendorName = object.vendorName
    } else {
      message.vendorName = ''
    }
    if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
      message.companyLegalName = object.companyLegalName
    } else {
      message.companyLegalName = ''
    }
    if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
      message.companyPrefferedName = object.companyPrefferedName
    } else {
      message.companyPrefferedName = ''
    }
    if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
      message.vendorLandingPageURL = object.vendorLandingPageURL
    } else {
      message.vendorLandingPageURL = ''
    }
    return message
  }
}

const baseMsgCreateVendorInfoTypeResponse: object = {}

export const MsgCreateVendorInfoTypeResponse = {
  encode(_: MsgCreateVendorInfoTypeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfoTypeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateVendorInfoTypeResponse } as MsgCreateVendorInfoTypeResponse
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

  fromJSON(_: any): MsgCreateVendorInfoTypeResponse {
    const message = { ...baseMsgCreateVendorInfoTypeResponse } as MsgCreateVendorInfoTypeResponse
    return message
  },

  toJSON(_: MsgCreateVendorInfoTypeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateVendorInfoTypeResponse>): MsgCreateVendorInfoTypeResponse {
    const message = { ...baseMsgCreateVendorInfoTypeResponse } as MsgCreateVendorInfoTypeResponse
    return message
  }
}

const baseMsgUpdateVendorInfoType: object = {
  creator: '',
  vendorID: 0,
  vendorName: '',
  companyLegalName: '',
  companyPrefferedName: '',
  vendorLandingPageURL: ''
}

export const MsgUpdateVendorInfoType = {
  encode(message: MsgUpdateVendorInfoType, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).uint64(message.vendorID)
    }
    if (message.vendorName !== '') {
      writer.uint32(26).string(message.vendorName)
    }
    if (message.companyLegalName !== '') {
      writer.uint32(34).string(message.companyLegalName)
    }
    if (message.companyPrefferedName !== '') {
      writer.uint32(42).string(message.companyPrefferedName)
    }
    if (message.vendorLandingPageURL !== '') {
      writer.uint32(50).string(message.vendorLandingPageURL)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfoType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateVendorInfoType } as MsgUpdateVendorInfoType
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.vendorID = longToNumber(reader.uint64() as Long)
          break
        case 3:
          message.vendorName = reader.string()
          break
        case 4:
          message.companyLegalName = reader.string()
          break
        case 5:
          message.companyPrefferedName = reader.string()
          break
        case 6:
          message.vendorLandingPageURL = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateVendorInfoType {
    const message = { ...baseMsgUpdateVendorInfoType } as MsgUpdateVendorInfoType
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    if (object.vendorName !== undefined && object.vendorName !== null) {
      message.vendorName = String(object.vendorName)
    } else {
      message.vendorName = ''
    }
    if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
      message.companyLegalName = String(object.companyLegalName)
    } else {
      message.companyLegalName = ''
    }
    if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
      message.companyPrefferedName = String(object.companyPrefferedName)
    } else {
      message.companyPrefferedName = ''
    }
    if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
      message.vendorLandingPageURL = String(object.vendorLandingPageURL)
    } else {
      message.vendorLandingPageURL = ''
    }
    return message
  },

  toJSON(message: MsgUpdateVendorInfoType): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    message.vendorName !== undefined && (obj.vendorName = message.vendorName)
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName)
    message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName)
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateVendorInfoType>): MsgUpdateVendorInfoType {
    const message = { ...baseMsgUpdateVendorInfoType } as MsgUpdateVendorInfoType
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
    }
    if (object.vendorName !== undefined && object.vendorName !== null) {
      message.vendorName = object.vendorName
    } else {
      message.vendorName = ''
    }
    if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
      message.companyLegalName = object.companyLegalName
    } else {
      message.companyLegalName = ''
    }
    if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
      message.companyPrefferedName = object.companyPrefferedName
    } else {
      message.companyPrefferedName = ''
    }
    if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
      message.vendorLandingPageURL = object.vendorLandingPageURL
    } else {
      message.vendorLandingPageURL = ''
    }
    return message
  }
}

const baseMsgUpdateVendorInfoTypeResponse: object = {}

export const MsgUpdateVendorInfoTypeResponse = {
  encode(_: MsgUpdateVendorInfoTypeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfoTypeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateVendorInfoTypeResponse } as MsgUpdateVendorInfoTypeResponse
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

  fromJSON(_: any): MsgUpdateVendorInfoTypeResponse {
    const message = { ...baseMsgUpdateVendorInfoTypeResponse } as MsgUpdateVendorInfoTypeResponse
    return message
  },

  toJSON(_: MsgUpdateVendorInfoTypeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateVendorInfoTypeResponse>): MsgUpdateVendorInfoTypeResponse {
    const message = { ...baseMsgUpdateVendorInfoTypeResponse } as MsgUpdateVendorInfoTypeResponse
    return message
  }
}

const baseMsgDeleteVendorInfoType: object = { creator: '', vendorID: 0 }

export const MsgDeleteVendorInfoType = {
  encode(message: MsgDeleteVendorInfoType, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).uint64(message.vendorID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfoType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteVendorInfoType } as MsgDeleteVendorInfoType
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.vendorID = longToNumber(reader.uint64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeleteVendorInfoType {
    const message = { ...baseMsgDeleteVendorInfoType } as MsgDeleteVendorInfoType
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = Number(object.vendorID)
    } else {
      message.vendorID = 0
    }
    return message
  },

  toJSON(message: MsgDeleteVendorInfoType): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteVendorInfoType>): MsgDeleteVendorInfoType {
    const message = { ...baseMsgDeleteVendorInfoType } as MsgDeleteVendorInfoType
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.vendorID !== undefined && object.vendorID !== null) {
      message.vendorID = object.vendorID
    } else {
      message.vendorID = 0
    }
    return message
  }
}

const baseMsgDeleteVendorInfoTypeResponse: object = {}

export const MsgDeleteVendorInfoTypeResponse = {
  encode(_: MsgDeleteVendorInfoTypeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfoTypeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteVendorInfoTypeResponse } as MsgDeleteVendorInfoTypeResponse
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

  fromJSON(_: any): MsgDeleteVendorInfoTypeResponse {
    const message = { ...baseMsgDeleteVendorInfoTypeResponse } as MsgDeleteVendorInfoTypeResponse
    return message
  },

  toJSON(_: MsgDeleteVendorInfoTypeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteVendorInfoTypeResponse>): MsgDeleteVendorInfoTypeResponse {
    const message = { ...baseMsgDeleteVendorInfoTypeResponse } as MsgDeleteVendorInfoTypeResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  CreateVendorInfoType(request: MsgCreateVendorInfoType): Promise<MsgCreateVendorInfoTypeResponse>
  UpdateVendorInfoType(request: MsgUpdateVendorInfoType): Promise<MsgUpdateVendorInfoTypeResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteVendorInfoType(request: MsgDeleteVendorInfoType): Promise<MsgDeleteVendorInfoTypeResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateVendorInfoType(request: MsgCreateVendorInfoType): Promise<MsgCreateVendorInfoTypeResponse> {
    const data = MsgCreateVendorInfoType.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'CreateVendorInfoType', data)
    return promise.then((data) => MsgCreateVendorInfoTypeResponse.decode(new Reader(data)))
  }

  UpdateVendorInfoType(request: MsgUpdateVendorInfoType): Promise<MsgUpdateVendorInfoTypeResponse> {
    const data = MsgUpdateVendorInfoType.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'UpdateVendorInfoType', data)
    return promise.then((data) => MsgUpdateVendorInfoTypeResponse.decode(new Reader(data)))
  }

  DeleteVendorInfoType(request: MsgDeleteVendorInfoType): Promise<MsgDeleteVendorInfoTypeResponse> {
    const data = MsgDeleteVendorInfoType.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'DeleteVendorInfoType', data)
    return promise.then((data) => MsgDeleteVendorInfoTypeResponse.decode(new Reader(data)))
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
