/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface CertificateIdentifier {
  subject: string
  subjectKeyId: string
}

const baseCertificateIdentifier: object = { subject: '', subjectKeyId: '' }

export const CertificateIdentifier = {
  encode(message: CertificateIdentifier, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): CertificateIdentifier {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseCertificateIdentifier } as CertificateIdentifier
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): CertificateIdentifier {
    const message = { ...baseCertificateIdentifier } as CertificateIdentifier
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = String(object.subject)
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = String(object.subjectKeyId)
    } else {
      message.subjectKeyId = ''
    }
    return message
  },

  toJSON(message: CertificateIdentifier): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<CertificateIdentifier>): CertificateIdentifier {
    const message = { ...baseCertificateIdentifier } as CertificateIdentifier
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = object.subject
    } else {
      message.subject = ''
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = object.subjectKeyId
    } else {
      message.subjectKeyId = ''
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
