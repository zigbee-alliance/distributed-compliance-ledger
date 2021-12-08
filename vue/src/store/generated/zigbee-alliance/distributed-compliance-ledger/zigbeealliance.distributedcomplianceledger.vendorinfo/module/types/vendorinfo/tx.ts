/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { VendorInfo } from '../vendorinfo/vendor_info'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo'

export interface MsgCreateNewVendorInfo {
  creator: string
  index: string
  vendorInfo: VendorInfo | undefined
}

export interface MsgCreateNewVendorInfoResponse {}

export interface MsgUpdateNewVendorInfo {
  creator: string
  index: string
  vendorInfo: VendorInfo | undefined
}

export interface MsgUpdateNewVendorInfoResponse {}

export interface MsgDeleteNewVendorInfo {
  creator: string
  index: string
}

export interface MsgDeleteNewVendorInfoResponse {}

const baseMsgCreateNewVendorInfo: object = { creator: '', index: '' }

export const MsgCreateNewVendorInfo = {
  encode(message: MsgCreateNewVendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.vendorInfo !== undefined) {
      VendorInfo.encode(message.vendorInfo, writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateNewVendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateNewVendorInfo } as MsgCreateNewVendorInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.index = reader.string()
          break
        case 3:
          message.vendorInfo = VendorInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgCreateNewVendorInfo {
    const message = { ...baseMsgCreateNewVendorInfo } as MsgCreateNewVendorInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromJSON(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    return message
  },

  toJSON(message: MsgCreateNewVendorInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.vendorInfo !== undefined && (obj.vendorInfo = message.vendorInfo ? VendorInfo.toJSON(message.vendorInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgCreateNewVendorInfo>): MsgCreateNewVendorInfo {
    const message = { ...baseMsgCreateNewVendorInfo } as MsgCreateNewVendorInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromPartial(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    return message
  }
}

const baseMsgCreateNewVendorInfoResponse: object = {}

export const MsgCreateNewVendorInfoResponse = {
  encode(_: MsgCreateNewVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateNewVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgCreateNewVendorInfoResponse } as MsgCreateNewVendorInfoResponse
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

  fromJSON(_: any): MsgCreateNewVendorInfoResponse {
    const message = { ...baseMsgCreateNewVendorInfoResponse } as MsgCreateNewVendorInfoResponse
    return message
  },

  toJSON(_: MsgCreateNewVendorInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgCreateNewVendorInfoResponse>): MsgCreateNewVendorInfoResponse {
    const message = { ...baseMsgCreateNewVendorInfoResponse } as MsgCreateNewVendorInfoResponse
    return message
  }
}

const baseMsgUpdateNewVendorInfo: object = { creator: '', index: '' }

export const MsgUpdateNewVendorInfo = {
  encode(message: MsgUpdateNewVendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    if (message.vendorInfo !== undefined) {
      VendorInfo.encode(message.vendorInfo, writer.uint32(26).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNewVendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNewVendorInfo } as MsgUpdateNewVendorInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.index = reader.string()
          break
        case 3:
          message.vendorInfo = VendorInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdateNewVendorInfo {
    const message = { ...baseMsgUpdateNewVendorInfo } as MsgUpdateNewVendorInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromJSON(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    return message
  },

  toJSON(message: MsgUpdateNewVendorInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    message.vendorInfo !== undefined && (obj.vendorInfo = message.vendorInfo ? VendorInfo.toJSON(message.vendorInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdateNewVendorInfo>): MsgUpdateNewVendorInfo {
    const message = { ...baseMsgUpdateNewVendorInfo } as MsgUpdateNewVendorInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
      message.vendorInfo = VendorInfo.fromPartial(object.vendorInfo)
    } else {
      message.vendorInfo = undefined
    }
    return message
  }
}

const baseMsgUpdateNewVendorInfoResponse: object = {}

export const MsgUpdateNewVendorInfoResponse = {
  encode(_: MsgUpdateNewVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateNewVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdateNewVendorInfoResponse } as MsgUpdateNewVendorInfoResponse
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

  fromJSON(_: any): MsgUpdateNewVendorInfoResponse {
    const message = { ...baseMsgUpdateNewVendorInfoResponse } as MsgUpdateNewVendorInfoResponse
    return message
  },

  toJSON(_: MsgUpdateNewVendorInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdateNewVendorInfoResponse>): MsgUpdateNewVendorInfoResponse {
    const message = { ...baseMsgUpdateNewVendorInfoResponse } as MsgUpdateNewVendorInfoResponse
    return message
  }
}

const baseMsgDeleteNewVendorInfo: object = { creator: '', index: '' }

export const MsgDeleteNewVendorInfo = {
  encode(message: MsgDeleteNewVendorInfo, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.index !== '') {
      writer.uint32(18).string(message.index)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteNewVendorInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteNewVendorInfo } as MsgDeleteNewVendorInfo
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.index = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeleteNewVendorInfo {
    const message = { ...baseMsgDeleteNewVendorInfo } as MsgDeleteNewVendorInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index)
    } else {
      message.index = ''
    }
    return message
  },

  toJSON(message: MsgDeleteNewVendorInfo): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.index !== undefined && (obj.index = message.index)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeleteNewVendorInfo>): MsgDeleteNewVendorInfo {
    const message = { ...baseMsgDeleteNewVendorInfo } as MsgDeleteNewVendorInfo
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index
    } else {
      message.index = ''
    }
    return message
  }
}

const baseMsgDeleteNewVendorInfoResponse: object = {}

export const MsgDeleteNewVendorInfoResponse = {
  encode(_: MsgDeleteNewVendorInfoResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteNewVendorInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeleteNewVendorInfoResponse } as MsgDeleteNewVendorInfoResponse
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

  fromJSON(_: any): MsgDeleteNewVendorInfoResponse {
    const message = { ...baseMsgDeleteNewVendorInfoResponse } as MsgDeleteNewVendorInfoResponse
    return message
  },

  toJSON(_: MsgDeleteNewVendorInfoResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeleteNewVendorInfoResponse>): MsgDeleteNewVendorInfoResponse {
    const message = { ...baseMsgDeleteNewVendorInfoResponse } as MsgDeleteNewVendorInfoResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  CreateNewVendorInfo(request: MsgCreateNewVendorInfo): Promise<MsgCreateNewVendorInfoResponse>
  UpdateNewVendorInfo(request: MsgUpdateNewVendorInfo): Promise<MsgUpdateNewVendorInfoResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteNewVendorInfo(request: MsgDeleteNewVendorInfo): Promise<MsgDeleteNewVendorInfoResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  CreateNewVendorInfo(request: MsgCreateNewVendorInfo): Promise<MsgCreateNewVendorInfoResponse> {
    const data = MsgCreateNewVendorInfo.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'CreateNewVendorInfo', data)
    return promise.then((data) => MsgCreateNewVendorInfoResponse.decode(new Reader(data)))
  }

  UpdateNewVendorInfo(request: MsgUpdateNewVendorInfo): Promise<MsgUpdateNewVendorInfoResponse> {
    const data = MsgUpdateNewVendorInfo.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'UpdateNewVendorInfo', data)
    return promise.then((data) => MsgUpdateNewVendorInfoResponse.decode(new Reader(data)))
  }

  DeleteNewVendorInfo(request: MsgDeleteNewVendorInfo): Promise<MsgDeleteNewVendorInfoResponse> {
    const data = MsgDeleteNewVendorInfo.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'DeleteNewVendorInfo', data)
    return promise.then((data) => MsgDeleteNewVendorInfoResponse.decode(new Reader(data)))
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
