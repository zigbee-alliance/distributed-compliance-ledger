/* eslint-disable */
import { Certificate } from '../pki/certificate'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface NocRootCertificatesByVidAndSkid {
  vid: number
  subjectKeyId: string
  certs: Certificate[]
  tq: number
}

const baseNocRootCertificatesByVidAndSkid: object = { vid: 0, subjectKeyId: '', tq: 0 }

export const NocRootCertificatesByVidAndSkid = {
  encode(message: NocRootCertificatesByVidAndSkid, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    for (const v of message.certs) {
      Certificate.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    if (message.tq !== 0) {
      writer.uint32(37).float(message.tq)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): NocRootCertificatesByVidAndSkid {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseNocRootCertificatesByVidAndSkid } as NocRootCertificatesByVidAndSkid
    message.certs = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        case 3:
          message.certs.push(Certificate.decode(reader, reader.uint32()))
          break
        case 4:
          message.tq = reader.float()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): NocRootCertificatesByVidAndSkid {
    const message = { ...baseNocRootCertificatesByVidAndSkid } as NocRootCertificatesByVidAndSkid
    message.certs = []
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
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
    if (object.tq !== undefined && object.tq !== null) {
      message.tq = Number(object.tq)
    } else {
      message.tq = 0
    }
    return message
  },

  toJSON(message: NocRootCertificatesByVidAndSkid): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    if (message.certs) {
      obj.certs = message.certs.map((e) => (e ? Certificate.toJSON(e) : undefined))
    } else {
      obj.certs = []
    }
    message.tq !== undefined && (obj.tq = message.tq)
    return obj
  },

  fromPartial(object: DeepPartial<NocRootCertificatesByVidAndSkid>): NocRootCertificatesByVidAndSkid {
    const message = { ...baseNocRootCertificatesByVidAndSkid } as NocRootCertificatesByVidAndSkid
    message.certs = []
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
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
    if (object.tq !== undefined && object.tq !== null) {
      message.tq = object.tq
    } else {
      message.tq = 0
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
