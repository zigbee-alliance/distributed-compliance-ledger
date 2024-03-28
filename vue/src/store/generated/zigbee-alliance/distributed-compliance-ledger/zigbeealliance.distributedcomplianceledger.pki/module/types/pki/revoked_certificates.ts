/* eslint-disable */
import { Certificate } from '../pki/certificate'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface RevokedCertificates {
  subject: string
  subjectKeyId: string
  certs: Certificate[]
  schemaVersion: number
}

const baseRevokedCertificates: object = { subject: '', subjectKeyId: '', schemaVersion: 0 }

export const RevokedCertificates = {
  encode(message: RevokedCertificates, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    for (const v of message.certs) {
      Certificate.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(32).uint32(message.schemaVersion)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RevokedCertificates {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRevokedCertificates } as RevokedCertificates
    message.certs = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        case 3:
          message.certs.push(Certificate.decode(reader, reader.uint32()))
          break
        case 4:
          message.schemaVersion = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RevokedCertificates {
    const message = { ...baseRevokedCertificates } as RevokedCertificates
    message.certs = []
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
    if (object.certs !== undefined && object.certs !== null) {
      for (const e of object.certs) {
        message.certs.push(Certificate.fromJSON(e))
      }
    }
    if (object.schemaVersion !== undefined && object.schemaVersion !== null) {
      message.schemaVersion = Number(object.schemaVersion)
    } else {
      message.schemaVersion = 0
    }
    return message
  },

  toJSON(message: RevokedCertificates): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    if (message.certs) {
      obj.certs = message.certs.map((e) => (e ? Certificate.toJSON(e) : undefined))
    } else {
      obj.certs = []
    }
    message.schemaVersion !== undefined && (obj.schemaVersion = message.schemaVersion)
    return obj
  },

  fromPartial(object: DeepPartial<RevokedCertificates>): RevokedCertificates {
    const message = { ...baseRevokedCertificates } as RevokedCertificates
    message.certs = []
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
    if (object.certs !== undefined && object.certs !== null) {
      for (const e of object.certs) {
        message.certs.push(Certificate.fromPartial(e))
      }
    }
    if (object.schemaVersion !== undefined && object.schemaVersion !== null) {
      message.schemaVersion = object.schemaVersion
    } else {
      message.schemaVersion = 0
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
