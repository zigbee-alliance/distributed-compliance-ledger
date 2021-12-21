/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface MsgCreateVendorInfo {
  creator: string
  vendorID: number
  vendorName: string
  companyLegalName: string
  companyPrefferedName: string
  vendorLandingPageURL: string
}

export interface MsgCreateVendorInfoResponse {}

export interface MsgUpdateVendorInfo {
  creator: string
  vendorID: number
  vendorName: string
  companyLegalName: string
  companyPrefferedName: string
  vendorLandingPageURL: string
}

export interface MsgUpdateVendorInfoResponse {}

export interface MsgDeleteVendorInfo {
  creator: string
  vendorID: number
}

export interface MsgDeleteVendorInfoResponse {}

const baseMsgCreateVendorInfo: object = { creator: '', vendorID: 0, vendorName: '', companyLegalName: '', companyPrefferedName: '', vendorLandingPageURL: '' }

export const MsgCreateVendorInfo = {
  encode(message: MsgCreateVendorInfo, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateVendorInfo } as MsgCreateVendorInfo
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

  fromJSON(object: any): MsgCreateVendorInfo {
    const message = { ...baseMsgCreateVendorInfo } as MsgCreateVendorInfo
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

  toJSON(message: MsgCreateVendorInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    message.vendorName !== undefined && (obj.vendorName = message.vendorName)
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName)
    message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName)
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateVendorInfo>): MsgCreateVendorInfo {
    const message = { ...baseMsgCreateVendorInfo } as MsgCreateVendorInfo
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

const baseMsgCreateVendorInfoResponse: object = {}

export const MsgCreateVendorInfoResponse = {
  encode(_: MsgCreateVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateVendorInfoResponse } as MsgCreateVendorInfoResponse
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

  fromJSON(_: any): MsgCreateVendorInfoResponse {
    const message = { ...baseMsgCreateVendorInfoResponse } as MsgCreateVendorInfoResponse
    return message
  },

  toJSON(_: MsgCreateVendorInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateVendorInfoResponse>): MsgCreateVendorInfoResponse {
    const message = { ...baseMsgCreateVendorInfoResponse } as MsgCreateVendorInfoResponse
    return message
  }
}

const baseMsgUpdateVendorInfo: object = { creator: '', vendorID: 0, vendorName: '', companyLegalName: '', companyPrefferedName: '', vendorLandingPageURL: '' }

export const MsgUpdateVendorInfo = {
  encode(message: MsgUpdateVendorInfo, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateVendorInfo } as MsgUpdateVendorInfo
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

  fromJSON(object: any): MsgUpdateVendorInfo {
    const message = { ...baseMsgUpdateVendorInfo } as MsgUpdateVendorInfo
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

  toJSON(message: MsgUpdateVendorInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    message.vendorName !== undefined && (obj.vendorName = message.vendorName)
    message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName)
    message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName)
    message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateVendorInfo>): MsgUpdateVendorInfo {
    const message = { ...baseMsgUpdateVendorInfo } as MsgUpdateVendorInfo
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

const baseMsgUpdateVendorInfoResponse: object = {}

export const MsgUpdateVendorInfoResponse = {
  encode(_: MsgUpdateVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateVendorInfoResponse } as MsgUpdateVendorInfoResponse
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

  fromJSON(_: any): MsgUpdateVendorInfoResponse {
    const message = { ...baseMsgUpdateVendorInfoResponse } as MsgUpdateVendorInfoResponse
    return message
  },

  toJSON(_: MsgUpdateVendorInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateVendorInfoResponse>): MsgUpdateVendorInfoResponse {
    const message = { ...baseMsgUpdateVendorInfoResponse } as MsgUpdateVendorInfoResponse
    return message
  }
}

const baseMsgDeleteVendorInfo: object = { creator: '', vendorID: 0 }

export const MsgDeleteVendorInfo = {
  encode(message: MsgDeleteVendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).uint64(message.vendorID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteVendorInfo } as MsgDeleteVendorInfo
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

  fromJSON(object: any): MsgDeleteVendorInfo {
    const message = { ...baseMsgDeleteVendorInfo } as MsgDeleteVendorInfo
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

  toJSON(message: MsgDeleteVendorInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.vendorID !== undefined && (obj.vendorID = message.vendorID)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteVendorInfo>): MsgDeleteVendorInfo {
    const message = { ...baseMsgDeleteVendorInfo } as MsgDeleteVendorInfo
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

const baseMsgDeleteVendorInfoResponse: object = {}

export const MsgDeleteVendorInfoResponse = {
  encode(_: MsgDeleteVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteVendorInfoResponse } as MsgDeleteVendorInfoResponse
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

  fromJSON(_: any): MsgDeleteVendorInfoResponse {
    const message = { ...baseMsgDeleteVendorInfoResponse } as MsgDeleteVendorInfoResponse
    return message
  },

  toJSON(_: MsgDeleteVendorInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteVendorInfoResponse>): MsgDeleteVendorInfoResponse {
    const message = { ...baseMsgDeleteVendorInfoResponse } as MsgDeleteVendorInfoResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  CreateVendorInfo(request: MsgCreateVendorInfo): Promise<MsgCreateVendorInfoResponse>
  UpdateVendorInfo(request: MsgUpdateVendorInfo): Promise<MsgUpdateVendorInfoResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteVendorInfo(request: MsgDeleteVendorInfo): Promise<MsgDeleteVendorInfoResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateVendorInfo(request: MsgCreateVendorInfo): Promise<MsgCreateVendorInfoResponse> {
    const data = MsgCreateVendorInfo.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'CreateVendorInfo', data)
    return promise.then((data) => MsgCreateVendorInfoResponse.decode(new Reader(data)))
  }

  UpdateVendorInfo(request: MsgUpdateVendorInfo): Promise<MsgUpdateVendorInfoResponse> {
    const data = MsgUpdateVendorInfo.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'UpdateVendorInfo', data)
    return promise.then((data) => MsgUpdateVendorInfoResponse.decode(new Reader(data)))
  }

  DeleteVendorInfo(request: MsgDeleteVendorInfo): Promise<MsgDeleteVendorInfoResponse> {
    const data = MsgDeleteVendorInfo.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'DeleteVendorInfo', data)
    return promise.then((data) => MsgDeleteVendorInfoResponse.decode(new Reader(data)))
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
