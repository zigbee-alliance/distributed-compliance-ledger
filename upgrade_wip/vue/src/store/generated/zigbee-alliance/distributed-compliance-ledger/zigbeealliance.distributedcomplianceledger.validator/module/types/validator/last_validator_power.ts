/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

export interface LastValidatorPower {
  consensusAddress: string
  power: number
}

const baseLastValidatorPower: object = { consensusAddress: '', power: 0 }

export const LastValidatorPower = {
  encode(message: LastValidatorPower, writer: Writer = Writer.create()): Writer {
    if (message.consensusAddress !== '') {
      writer.uint32(10).string(message.consensusAddress)
    }
    if (message.power !== 0) {
      writer.uint32(16).int32(message.power)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): LastValidatorPower {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseLastValidatorPower } as LastValidatorPower
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.consensusAddress = reader.string()
          break
        case 2:
          message.power = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): LastValidatorPower {
    const message = { ...baseLastValidatorPower } as LastValidatorPower
    if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
      message.consensusAddress = String(object.consensusAddress)
    } else {
      message.consensusAddress = ''
    }
    if (object.power !== undefined && object.power !== null) {
      message.power = Number(object.power)
    } else {
      message.power = 0
    }
    return message
  },

  toJSON(message: LastValidatorPower): unknown {
    const obj: any = {}
    message.consensusAddress !== undefined && (obj.consensusAddress = message.consensusAddress)
    message.power !== undefined && (obj.power = message.power)
    return obj
  },

  fromPartial(object: DeepPartial<LastValidatorPower>): LastValidatorPower {
    const message = { ...baseLastValidatorPower } as LastValidatorPower
    if (object.consensusAddress !== undefined && object.consensusAddress !== null) {
      message.consensusAddress = object.consensusAddress
    } else {
      message.consensusAddress = ''
    }
    if (object.power !== undefined && object.power !== null) {
      message.power = object.power
    } else {
      message.power = 0
    }
    return message
  }
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
