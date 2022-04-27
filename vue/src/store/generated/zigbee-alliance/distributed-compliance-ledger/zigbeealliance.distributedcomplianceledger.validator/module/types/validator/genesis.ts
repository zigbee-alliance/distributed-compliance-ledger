/* eslint-disable */
import { Validator } from '../validator/validator'
import { LastValidatorPower } from '../validator/last_validator_power'
import { ProposedDisableValidator } from '../validator/proposed_disable_validator'
import { DisabledValidator } from '../validator/disabled_validator'
import { RejectedDisableValidator } from '../validator/rejected_validator'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

/** GenesisState defines the validator module's genesis state. */
export interface GenesisState {
  validatorList: Validator[]
  lastValidatorPowerList: LastValidatorPower[]
  proposedDisableValidatorList: ProposedDisableValidator[]
  disabledValidatorList: DisabledValidator[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  rejectedValidatorList: RejectedDisableValidator[]
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
    for (const v of message.proposedDisableValidatorList) {
      ProposedDisableValidator.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.disabledValidatorList) {
      DisabledValidator.encode(v!, writer.uint32(34).fork()).ldelim()
    }
    for (const v of message.rejectedValidatorList) {
      RejectedDisableValidator.encode(v!, writer.uint32(42).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.validatorList = []
    message.lastValidatorPowerList = []
    message.proposedDisableValidatorList = []
    message.disabledValidatorList = []
    message.rejectedValidatorList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.validatorList.push(Validator.decode(reader, reader.uint32()))
          break
        case 2:
          message.lastValidatorPowerList.push(LastValidatorPower.decode(reader, reader.uint32()))
          break
        case 3:
          message.proposedDisableValidatorList.push(ProposedDisableValidator.decode(reader, reader.uint32()))
          break
        case 4:
          message.disabledValidatorList.push(DisabledValidator.decode(reader, reader.uint32()))
          break
        case 5:
          message.rejectedValidatorList.push(RejectedDisableValidator.decode(reader, reader.uint32()))
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
    message.proposedDisableValidatorList = []
    message.disabledValidatorList = []
    message.rejectedValidatorList = []
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
    if (object.proposedDisableValidatorList !== undefined && object.proposedDisableValidatorList !== null) {
      for (const e of object.proposedDisableValidatorList) {
        message.proposedDisableValidatorList.push(ProposedDisableValidator.fromJSON(e))
      }
    }
    if (object.disabledValidatorList !== undefined && object.disabledValidatorList !== null) {
      for (const e of object.disabledValidatorList) {
        message.disabledValidatorList.push(DisabledValidator.fromJSON(e))
      }
    }
    if (object.rejectedValidatorList !== undefined && object.rejectedValidatorList !== null) {
      for (const e of object.rejectedValidatorList) {
        message.rejectedValidatorList.push(RejectedDisableValidator.fromJSON(e))
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
    if (message.proposedDisableValidatorList) {
      obj.proposedDisableValidatorList = message.proposedDisableValidatorList.map((e) => (e ? ProposedDisableValidator.toJSON(e) : undefined))
    } else {
      obj.proposedDisableValidatorList = []
    }
    if (message.disabledValidatorList) {
      obj.disabledValidatorList = message.disabledValidatorList.map((e) => (e ? DisabledValidator.toJSON(e) : undefined))
    } else {
      obj.disabledValidatorList = []
    }
    if (message.rejectedValidatorList) {
      obj.rejectedValidatorList = message.rejectedValidatorList.map((e) => (e ? RejectedDisableValidator.toJSON(e) : undefined))
    } else {
      obj.rejectedValidatorList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.validatorList = []
    message.lastValidatorPowerList = []
    message.proposedDisableValidatorList = []
    message.disabledValidatorList = []
    message.rejectedValidatorList = []
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
    if (object.proposedDisableValidatorList !== undefined && object.proposedDisableValidatorList !== null) {
      for (const e of object.proposedDisableValidatorList) {
        message.proposedDisableValidatorList.push(ProposedDisableValidator.fromPartial(e))
      }
    }
    if (object.disabledValidatorList !== undefined && object.disabledValidatorList !== null) {
      for (const e of object.disabledValidatorList) {
        message.disabledValidatorList.push(DisabledValidator.fromPartial(e))
      }
    }
    if (object.rejectedValidatorList !== undefined && object.rejectedValidatorList !== null) {
      for (const e of object.rejectedValidatorList) {
        message.rejectedValidatorList.push(RejectedDisableValidator.fromPartial(e))
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
