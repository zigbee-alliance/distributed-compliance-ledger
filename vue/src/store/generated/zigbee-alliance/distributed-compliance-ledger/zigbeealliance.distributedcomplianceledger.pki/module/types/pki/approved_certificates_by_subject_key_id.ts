/* eslint-disable */
import { Certificate } from '../pki/certificate'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface ApprovedCertificatesBySubjectKeyId {
  subjectKeyId: string
  certs: Certificate[]
}

const baseApprovedCertificatesBySubjectKeyId: object = { subjectKeyId: '' }

export const ApprovedCertificatesBySubjectKeyId = {
  encode(message: ApprovedCertificatesBySubjectKeyId, writer: Writer = Writer.create()): Writer {
    if (message.subjectKeyId !== '') {
      writer.uint32(10).string(message.subjectKeyId)
    }
    for (const v of message.certs) {
      Certificate.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ApprovedCertificatesBySubjectKeyId {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseApprovedCertificatesBySubjectKeyId } as ApprovedCertificatesBySubjectKeyId
    message.certs = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subjectKeyId = reader.string()
          break
        case 2:
          message.certs.push(Certificate.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ApprovedCertificatesBySubjectKeyId {
    const message = { ...baseApprovedCertificatesBySubjectKeyId } as ApprovedCertificatesBySubjectKeyId
    message.certs = []
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
    return message
  },

  toJSON(message: ApprovedCertificatesBySubjectKeyId): unknown {
    const obj: any = {}
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    if (message.certs) {
      obj.certs = message.certs.map((e) => (e ? Certificate.toJSON(e) : undefined))
    } else {
      obj.certs = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<ApprovedCertificatesBySubjectKeyId>): ApprovedCertificatesBySubjectKeyId {
    const message = { ...baseApprovedCertificatesBySubjectKeyId } as ApprovedCertificatesBySubjectKeyId
    message.certs = []
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
