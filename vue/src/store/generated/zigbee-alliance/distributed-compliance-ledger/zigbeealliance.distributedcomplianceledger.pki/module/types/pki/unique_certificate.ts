/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface UniqueCertificate {
  issuer: string
  serialNumber: string
  present: boolean
}

const baseUniqueCertificate: object = { issuer: '', serialNumber: '', present: false }

export const UniqueCertificate = {
  encode(message: UniqueCertificate, writer: Writer = Writer.create()): Writer {
    if (message.issuer !== '') {
      writer.uint32(10).string(message.issuer)
    }
    if (message.serialNumber !== '') {
      writer.uint32(18).string(message.serialNumber)
    }
    if (message.present === true) {
      writer.uint32(24).bool(message.present)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): UniqueCertificate {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseUniqueCertificate } as UniqueCertificate
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string()
          break
        case 2:
          message.serialNumber = reader.string()
          break
        case 3:
          message.present = reader.bool()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): UniqueCertificate {
    const message = { ...baseUniqueCertificate } as UniqueCertificate
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = String(object.issuer)
    } else {
      message.issuer = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = String(object.serialNumber)
    } else {
      message.serialNumber = ''
    }
    if (object.present !== undefined && object.present !== null) {
      message.present = Boolean(object.present)
    } else {
      message.present = false
    }
    return message
  },

  toJSON(message: UniqueCertificate): unknown {
    const obj: any = {}
    message.issuer !== undefined && (obj.issuer = message.issuer)
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber)
    message.present !== undefined && (obj.present = message.present)
    return obj
  },

  fromPartial(object: DeepPartial<UniqueCertificate>): UniqueCertificate {
    const message = { ...baseUniqueCertificate } as UniqueCertificate
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = object.issuer
    } else {
      message.issuer = ''
    }
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = object.serialNumber
    } else {
      message.serialNumber = ''
    }
    if (object.present !== undefined && object.present !== null) {
      message.present = object.present
    } else {
      message.present = false
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
