/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'
import { Plan } from '../cosmos/upgrade/v1beta1/upgrade'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade'

export interface MsgProposeUpgrade {
  creator: string
  plan: Plan | undefined
  info: string
  time: number
}

export interface MsgProposeUpgradeResponse {}

export interface MsgApproveUpgrade {
  creator: string
  name: string
  info: string
  time: number
}

export interface MsgApproveUpgradeResponse {}

export interface MsgRejectUpgrade {
  creator: string
  name: string
  info: string
  time: number
}

export interface MsgRejectUpgradeResponse {}

const baseMsgProposeUpgrade: object = { creator: '', info: '', time: 0 }

export const MsgProposeUpgrade = {
  encode(message: MsgProposeUpgrade, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.plan !== undefined) {
      Plan.encode(message.plan, writer.uint32(18).fork()).ldelim()
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeUpgrade {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeUpgrade } as MsgProposeUpgrade
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.plan = Plan.decode(reader, reader.uint32())
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

  fromJSON(object: any): MsgProposeUpgrade {
    const message = { ...baseMsgProposeUpgrade } as MsgProposeUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.plan !== undefined && object.plan !== null) {
      message.plan = Plan.fromJSON(object.plan)
    } else {
      message.plan = undefined
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

  toJSON(message: MsgProposeUpgrade): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.plan !== undefined && (obj.plan = message.plan ? Plan.toJSON(message.plan) : undefined)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgProposeUpgrade>): MsgProposeUpgrade {
    const message = { ...baseMsgProposeUpgrade } as MsgProposeUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.plan !== undefined && object.plan !== null) {
      message.plan = Plan.fromPartial(object.plan)
    } else {
      message.plan = undefined
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

const baseMsgProposeUpgradeResponse: object = {}

export const MsgProposeUpgradeResponse = {
  encode(_: MsgProposeUpgradeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeUpgradeResponse } as MsgProposeUpgradeResponse
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

  fromJSON(_: any): MsgProposeUpgradeResponse {
    const message = { ...baseMsgProposeUpgradeResponse } as MsgProposeUpgradeResponse
    return message
  },

  toJSON(_: MsgProposeUpgradeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgProposeUpgradeResponse>): MsgProposeUpgradeResponse {
    const message = { ...baseMsgProposeUpgradeResponse } as MsgProposeUpgradeResponse
    return message
  }
}

const baseMsgApproveUpgrade: object = { creator: '', name: '', info: '', time: 0 }

export const MsgApproveUpgrade = {
  encode(message: MsgApproveUpgrade, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name)
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveUpgrade {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveUpgrade } as MsgApproveUpgrade
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.name = reader.string()
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

  fromJSON(object: any): MsgApproveUpgrade {
    const message = { ...baseMsgApproveUpgrade } as MsgApproveUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
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

  toJSON(message: MsgApproveUpgrade): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.name !== undefined && (obj.name = message.name)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgApproveUpgrade>): MsgApproveUpgrade {
    const message = { ...baseMsgApproveUpgrade } as MsgApproveUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
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

const baseMsgApproveUpgradeResponse: object = {}

export const MsgApproveUpgradeResponse = {
  encode(_: MsgApproveUpgradeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveUpgradeResponse } as MsgApproveUpgradeResponse
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

  fromJSON(_: any): MsgApproveUpgradeResponse {
    const message = { ...baseMsgApproveUpgradeResponse } as MsgApproveUpgradeResponse
    return message
  },

  toJSON(_: MsgApproveUpgradeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgApproveUpgradeResponse>): MsgApproveUpgradeResponse {
    const message = { ...baseMsgApproveUpgradeResponse } as MsgApproveUpgradeResponse
    return message
  }
}

const baseMsgRejectUpgrade: object = { creator: '', name: '', info: '', time: 0 }

export const MsgRejectUpgrade = {
  encode(message: MsgRejectUpgrade, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.name !== '') {
      writer.uint32(18).string(message.name)
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectUpgrade {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRejectUpgrade } as MsgRejectUpgrade
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.name = reader.string()
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

  fromJSON(object: any): MsgRejectUpgrade {
    const message = { ...baseMsgRejectUpgrade } as MsgRejectUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
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

  toJSON(message: MsgRejectUpgrade): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.name !== undefined && (obj.name = message.name)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgRejectUpgrade>): MsgRejectUpgrade {
    const message = { ...baseMsgRejectUpgrade } as MsgRejectUpgrade
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
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

const baseMsgRejectUpgradeResponse: object = {}

export const MsgRejectUpgradeResponse = {
  encode(_: MsgRejectUpgradeResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectUpgradeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRejectUpgradeResponse } as MsgRejectUpgradeResponse
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

  fromJSON(_: any): MsgRejectUpgradeResponse {
    const message = { ...baseMsgRejectUpgradeResponse } as MsgRejectUpgradeResponse
    return message
  },

  toJSON(_: MsgRejectUpgradeResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgRejectUpgradeResponse>): MsgRejectUpgradeResponse {
    const message = { ...baseMsgRejectUpgradeResponse } as MsgRejectUpgradeResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse>
  ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RejectUpgrade(request: MsgRejectUpgrade): Promise<MsgRejectUpgradeResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse> {
    const data = MsgProposeUpgrade.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'ProposeUpgrade', data)
    return promise.then((data) => MsgProposeUpgradeResponse.decode(new Reader(data)))
  }

  ApproveUpgrade(request: MsgApproveUpgrade): Promise<MsgApproveUpgradeResponse> {
    const data = MsgApproveUpgrade.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'ApproveUpgrade', data)
    return promise.then((data) => MsgApproveUpgradeResponse.decode(new Reader(data)))
  }

  RejectUpgrade(request: MsgRejectUpgrade): Promise<MsgRejectUpgradeResponse> {
    const data = MsgRejectUpgrade.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclupgrade.Msg', 'RejectUpgrade', data)
    return promise.then((data) => MsgRejectUpgradeResponse.decode(new Reader(data)))
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
