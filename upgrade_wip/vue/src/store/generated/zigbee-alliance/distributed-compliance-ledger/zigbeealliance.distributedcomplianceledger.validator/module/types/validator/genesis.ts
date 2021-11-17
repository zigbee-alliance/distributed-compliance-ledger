/* eslint-disable */
import { Validator } from '../validator/validator'
import { LastValidatorPower } from '../validator/last_validator_power'
import { ValidatorSigningInfo } from '../validator/validator_signing_info'
import { ValidatorMissedBlockBitArray } from '../validator/validator_missed_block_bit_array'
import { ValidatorOwner } from '../validator/validator_owner'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.validator'

/** GenesisState defines the validator module's genesis state. */
export interface GenesisState {
  validatorList: Validator[]
  lastValidatorPowerList: LastValidatorPower[]
  validatorSigningInfoList: ValidatorSigningInfo[]
  validatorMissedBlockBitArrayList: ValidatorMissedBlockBitArray[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  validatorOwnerList: ValidatorOwner[]
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
    for (const v of message.validatorSigningInfoList) {
      ValidatorSigningInfo.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.validatorMissedBlockBitArrayList) {
      ValidatorMissedBlockBitArray.encode(v!, writer.uint32(34).fork()).ldelim()
    }
    for (const v of message.validatorOwnerList) {
      ValidatorOwner.encode(v!, writer.uint32(42).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.validatorList = []
    message.lastValidatorPowerList = []
    message.validatorSigningInfoList = []
    message.validatorMissedBlockBitArrayList = []
    message.validatorOwnerList = []
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
          message.validatorSigningInfoList.push(ValidatorSigningInfo.decode(reader, reader.uint32()))
          break
        case 4:
          message.validatorMissedBlockBitArrayList.push(ValidatorMissedBlockBitArray.decode(reader, reader.uint32()))
          break
        case 5:
          message.validatorOwnerList.push(ValidatorOwner.decode(reader, reader.uint32()))
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
    message.validatorSigningInfoList = []
    message.validatorMissedBlockBitArrayList = []
    message.validatorOwnerList = []
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
    if (object.validatorSigningInfoList !== undefined && object.validatorSigningInfoList !== null) {
      for (const e of object.validatorSigningInfoList) {
        message.validatorSigningInfoList.push(ValidatorSigningInfo.fromJSON(e))
      }
    }
    if (object.validatorMissedBlockBitArrayList !== undefined && object.validatorMissedBlockBitArrayList !== null) {
      for (const e of object.validatorMissedBlockBitArrayList) {
        message.validatorMissedBlockBitArrayList.push(ValidatorMissedBlockBitArray.fromJSON(e))
      }
    }
    if (object.validatorOwnerList !== undefined && object.validatorOwnerList !== null) {
      for (const e of object.validatorOwnerList) {
        message.validatorOwnerList.push(ValidatorOwner.fromJSON(e))
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
    if (message.validatorSigningInfoList) {
      obj.validatorSigningInfoList = message.validatorSigningInfoList.map((e) => (e ? ValidatorSigningInfo.toJSON(e) : undefined))
    } else {
      obj.validatorSigningInfoList = []
    }
    if (message.validatorMissedBlockBitArrayList) {
      obj.validatorMissedBlockBitArrayList = message.validatorMissedBlockBitArrayList.map((e) => (e ? ValidatorMissedBlockBitArray.toJSON(e) : undefined))
    } else {
      obj.validatorMissedBlockBitArrayList = []
    }
    if (message.validatorOwnerList) {
      obj.validatorOwnerList = message.validatorOwnerList.map((e) => (e ? ValidatorOwner.toJSON(e) : undefined))
    } else {
      obj.validatorOwnerList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.validatorList = []
    message.lastValidatorPowerList = []
    message.validatorSigningInfoList = []
    message.validatorMissedBlockBitArrayList = []
    message.validatorOwnerList = []
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
    if (object.validatorSigningInfoList !== undefined && object.validatorSigningInfoList !== null) {
      for (const e of object.validatorSigningInfoList) {
        message.validatorSigningInfoList.push(ValidatorSigningInfo.fromPartial(e))
      }
    }
    if (object.validatorMissedBlockBitArrayList !== undefined && object.validatorMissedBlockBitArrayList !== null) {
      for (const e of object.validatorMissedBlockBitArrayList) {
        message.validatorMissedBlockBitArrayList.push(ValidatorMissedBlockBitArray.fromPartial(e))
      }
    }
    if (object.validatorOwnerList !== undefined && object.validatorOwnerList !== null) {
      for (const e of object.validatorOwnerList) {
        message.validatorOwnerList.push(ValidatorOwner.fromPartial(e))
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
