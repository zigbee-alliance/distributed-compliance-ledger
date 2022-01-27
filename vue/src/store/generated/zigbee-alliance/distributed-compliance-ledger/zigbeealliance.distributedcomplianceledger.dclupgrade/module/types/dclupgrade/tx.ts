/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade'

export interface MsgProposeUpgrade {
  creator: string
  plan: string
}

export interface MsgProposeUpgradeResponse {}

const baseMsgProposeUpgrade: object = { creator: '', plan: '' }

export const MsgProposeUpgrade = {
  encode(message: MsgProposeUpgrade, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.plan !== '') {
      writer.uint32(18).string(message.plan)
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
          message.plan = reader.string()
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
      message.plan = String(object.plan)
    } else {
      message.plan = ''
    }
    return message
  },

  toJSON(message: MsgProposeUpgrade): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.plan !== undefined && (obj.plan = message.plan)
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
      message.plan = object.plan
    } else {
      message.plan = ''
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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ProposeUpgrade(request: MsgProposeUpgrade): Promise<MsgProposeUpgradeResponse>
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
