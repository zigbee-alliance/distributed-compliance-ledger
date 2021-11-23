/* eslint-disable */
import { Validator } from '../validator/validator'
import { LastValidatorPower } from '../validator/last_validator_power'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

/** GenesisState defines the validator module's genesis state. */
export interface GenesisState {
  validatorList: Validator[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  lastValidatorPowerList: LastValidatorPower[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.validatorList) {
      Validator.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.lastValidatorPowerList) {
      LastValidatorPower.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.validatorList = []
    message.lastValidatorPowerList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorList.push(Validator.decode(reader, reader.uint32()))
          break
        case 2:
          message.lastValidatorPowerList.push(LastValidatorPower.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.validatorList = []
    message.lastValidatorPowerList = []
    if (object.validatorList !== undefined && object.validatorList !== null) {
      for (const e of object.validatorList) {
        message.validatorList.push(Validator.fromJSON(e))
      }
    }
    if (object.lastValidatorPowerList !== undefined && object.lastValidatorPowerList !== null) {
      for (const e of object.lastValidatorPowerList) {
        message.lastValidatorPowerList.push(LastValidatorPower.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.validatorList) {
      obj.validatorList = message.validatorList.map((e) => (e ? Validator.toJSON(e) : undefined))
    } else {
      obj.validatorList = []
    }
    if (message.lastValidatorPowerList) {
      obj.lastValidatorPowerList = message.lastValidatorPowerList.map((e) => (e ? LastValidatorPower.toJSON(e) : undefined))
    } else {
      obj.lastValidatorPowerList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.validatorList = []
    message.lastValidatorPowerList = []
    if (object.validatorList !== undefined && object.validatorList !== null) {
      for (const e of object.validatorList) {
        message.validatorList.push(Validator.fromPartial(e))
      }
    }
    if (object.lastValidatorPowerList !== undefined && object.lastValidatorPowerList !== null) {
      for (const e of object.lastValidatorPowerList) {
        message.lastValidatorPowerList.push(LastValidatorPower.fromPartial(e))
      }
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
