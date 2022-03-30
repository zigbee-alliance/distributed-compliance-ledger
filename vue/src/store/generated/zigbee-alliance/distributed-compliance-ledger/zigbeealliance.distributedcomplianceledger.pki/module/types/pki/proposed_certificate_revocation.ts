/* eslint-disable */
import { Grant } from '../pki/grant'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface ProposedCertificateRevocation {
  subject: string
  subjectKeyId: string
  approvals: Grant[]
  subjectAsText: string
}

const baseProposedCertificateRevocation: object = { subject: '', subjectKeyId: '', subjectAsText: '' }

export const ProposedCertificateRevocation = {
  encode(message: ProposedCertificateRevocation, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    for (const v of message.approvals) {
      Grant.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    if (message.subjectAsText !== '') {
      writer.uint32(34).string(message.subjectAsText)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): ProposedCertificateRevocation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseProposedCertificateRevocation } as ProposedCertificateRevocation
    message.approvals = []
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
          message.approvals.push(Grant.decode(reader, reader.uint32()))
          break
        case 4:
          message.subjectAsText = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): ProposedCertificateRevocation {
    const message = { ...baseProposedCertificateRevocation } as ProposedCertificateRevocation
    message.approvals = []
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
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromJSON(e))
      }
    }
    if (object.subjectAsText !== undefined && object.subjectAsText !== null) {
      message.subjectAsText = String(object.subjectAsText)
    } else {
      message.subjectAsText = ''
    }
    return message
  },

  toJSON(message: ProposedCertificateRevocation): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    if (message.approvals) {
      obj.approvals = message.approvals.map((e) => (e ? Grant.toJSON(e) : undefined))
    } else {
      obj.approvals = []
    }
    message.subjectAsText !== undefined && (obj.subjectAsText = message.subjectAsText)
    return obj
  },

  fromPartial(object: DeepPartial<ProposedCertificateRevocation>): ProposedCertificateRevocation {
    const message = { ...baseProposedCertificateRevocation } as ProposedCertificateRevocation
    message.approvals = []
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
    if (object.approvals !== undefined && object.approvals !== null) {
      for (const e of object.approvals) {
        message.approvals.push(Grant.fromPartial(e))
      }
    }
    if (object.subjectAsText !== undefined && object.subjectAsText !== null) {
      message.subjectAsText = object.subjectAsText
    } else {
      message.subjectAsText = ''
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
