/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliancetest'

export interface MsgAddTestingResult {
  signer: string
  vid: number
  pid: number
  softwareVersion: number
  softwareVersionString: string
  testResult: string
  testDate: string
}

export interface MsgAddTestingResultResponse {}

const baseMsgAddTestingResult: object = { signer: '', vid: 0, pid: 0, softwareVersion: 0, softwareVersionString: '', testResult: '', testDate: '' }

export const MsgAddTestingResult = {
  encode(message: MsgAddTestingResult, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.vid !== 0) {
      writer.uint32(16).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(24).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(32).uint32(message.softwareVersion)
    }
    if (message.softwareVersionString !== '') {
      writer.uint32(42).string(message.softwareVersionString)
    }
    if (message.testResult !== '') {
      writer.uint32(50).string(message.testResult)
    }
    if (message.testDate !== '') {
      writer.uint32(58).string(message.testDate)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddTestingResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddTestingResult } as MsgAddTestingResult
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.vid = reader.int32()
          break
        case 3:
          message.pid = reader.int32()
          break
        case 4:
          message.softwareVersion = reader.uint32()
          break
        case 5:
          message.softwareVersionString = reader.string()
          break
        case 6:
          message.testResult = reader.string()
          break
        case 7:
          message.testDate = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgAddTestingResult {
    const message = { ...baseMsgAddTestingResult } as MsgAddTestingResult
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
      message.softwareVersionString = String(object.softwareVersionString)
    } else {
      message.softwareVersionString = ''
    }
    if (object.testResult !== undefined && object.testResult !== null) {
      message.testResult = String(object.testResult)
    } else {
      message.testResult = ''
    }
    if (object.testDate !== undefined && object.testDate !== null) {
      message.testDate = String(object.testDate)
    } else {
      message.testDate = ''
    }
    return message
  },

  toJSON(message: MsgAddTestingResult): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    message.softwareVersionString !== undefined && (obj.softwareVersionString = message.softwareVersionString)
    message.testResult !== undefined && (obj.testResult = message.testResult)
    message.testDate !== undefined && (obj.testDate = message.testDate)
    return obj
  },

  fromPartial(object: DeepPartial<MsgAddTestingResult>): MsgAddTestingResult {
    const message = { ...baseMsgAddTestingResult } as MsgAddTestingResult
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
      message.softwareVersionString = object.softwareVersionString
    } else {
      message.softwareVersionString = ''
    }
    if (object.testResult !== undefined && object.testResult !== null) {
      message.testResult = object.testResult
    } else {
      message.testResult = ''
    }
    if (object.testDate !== undefined && object.testDate !== null) {
      message.testDate = object.testDate
    } else {
      message.testDate = ''
    }
    return message
  }
}

const baseMsgAddTestingResultResponse: object = {}

export const MsgAddTestingResultResponse = {
  encode(_: MsgAddTestingResultResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddTestingResultResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddTestingResultResponse } as MsgAddTestingResultResponse
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

  fromJSON(_: any): MsgAddTestingResultResponse {
    const message = { ...baseMsgAddTestingResultResponse } as MsgAddTestingResultResponse
    return message
  },

  toJSON(_: MsgAddTestingResultResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgAddTestingResultResponse>): MsgAddTestingResultResponse {
    const message = { ...baseMsgAddTestingResultResponse } as MsgAddTestingResultResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AddTestingResult(request: MsgAddTestingResult): Promise<MsgAddTestingResultResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  AddTestingResult(request: MsgAddTestingResult): Promise<MsgAddTestingResultResponse> {
    const data = MsgAddTestingResult.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliancetest.Msg', 'AddTestingResult', data)
    return promise.then((data) => MsgAddTestingResultResponse.decode(new Reader(data)))
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
