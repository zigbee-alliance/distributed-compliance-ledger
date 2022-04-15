/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface MsgCreateVendorInfo {
  creator: string
  vendorID: number
  vendorName: string
  companyLegalName: string
  companyPreferredName: string
  vendorLandingPageURL: string
}

export interface MsgCreateVendorInfoResponse {}

export interface MsgUpdateVendorInfo {
  creator: string
  vendorID: number
  vendorName: string
  companyLegalName: string
  companyPreferredName: string
  vendorLandingPageURL: string
}

export interface MsgUpdateVendorInfoResponse {}

const baseMsgCreateVendorInfo: object = { creator: '', vendorID: 0, vendorName: '', companyLegalName: '', companyPreferredName: '', vendorLandingPageURL: '' }

export const MsgCreateVendorInfo = {
  encode(message: MsgCreateVendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).int32(message.vendorID)
    }
    if (message.vendorName !== '') {
      writer.uint32(26).string(message.vendorName)
    }
    if (message.companyLegalName !== '') {
      writer.uint32(34).string(message.companyLegalName)
    }
    if (message.companyPreferredName !== '') {
      writer.uint32(42).string(message.companyPreferredName)
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
          message.vendorID = reader.int32()
          break
        case 3:
          message.vendorName = reader.string()
          break
        case 4:
          message.companyLegalName = reader.string()
          break
        case 5:
          message.companyPreferredName = reader.string()
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
    if (object.companyPreferredName !== undefined && object.companyPreferredName !== null) {
      message.companyPreferredName = String(object.companyPreferredName)
    } else {
      message.companyPreferredName = ''
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
    message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName)
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
    if (object.companyPreferredName !== undefined && object.companyPreferredName !== null) {
      message.companyPreferredName = object.companyPreferredName
    } else {
      message.companyPreferredName = ''
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

const baseMsgUpdateVendorInfo: object = { creator: '', vendorID: 0, vendorName: '', companyLegalName: '', companyPreferredName: '', vendorLandingPageURL: '' }

export const MsgUpdateVendorInfo = {
  encode(message: MsgUpdateVendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.vendorID !== 0) {
      writer.uint32(16).int32(message.vendorID)
    }
    if (message.vendorName !== '') {
      writer.uint32(26).string(message.vendorName)
    }
    if (message.companyLegalName !== '') {
      writer.uint32(34).string(message.companyLegalName)
    }
    if (message.companyPreferredName !== '') {
      writer.uint32(42).string(message.companyPreferredName)
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
          message.vendorID = reader.int32()
          break
        case 3:
          message.vendorName = reader.string()
          break
        case 4:
          message.companyLegalName = reader.string()
          break
        case 5:
          message.companyPreferredName = reader.string()
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
    if (object.companyPreferredName !== undefined && object.companyPreferredName !== null) {
      message.companyPreferredName = String(object.companyPreferredName)
    } else {
      message.companyPreferredName = ''
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
    message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName)
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
    if (object.companyPreferredName !== undefined && object.companyPreferredName !== null) {
      message.companyPreferredName = object.companyPreferredName
    } else {
      message.companyPreferredName = ''
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

/** Msg defines the Msg service. */
export interface Msg {
  CreateVendorInfo(request: MsgCreateVendorInfo): Promise<MsgCreateVendorInfoResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateVendorInfo(request: MsgUpdateVendorInfo): Promise<MsgUpdateVendorInfoResponse>
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
