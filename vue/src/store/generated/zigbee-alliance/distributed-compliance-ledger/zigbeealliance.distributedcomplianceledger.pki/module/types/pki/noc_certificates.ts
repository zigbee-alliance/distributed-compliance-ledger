/* eslint-disable */
import { Certificate } from '../pki/certificate'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface NocCertificates {
  vid: number
  certs: Certificate[]
}

const baseNocCertificates: object = { vid: 0 }

export const NocCertificates = {
  encode(message: NocCertificates, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    for (const v of message.certs) {
      Certificate.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): NocCertificates {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseNocCertificates } as NocCertificates
    message.certs = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
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

  fromJSON(object: any): NocCertificates {
    const message = { ...baseNocCertificates } as NocCertificates
    message.certs = []
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.certs !== undefined && object.certs !== null) {
      for (const e of object.certs) {
        message.certs.push(Certificate.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: NocCertificates): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    if (message.certs) {
      obj.certs = message.certs.map((e) => (e ? Certificate.toJSON(e) : undefined))
    } else {
      obj.certs = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<NocCertificates>): NocCertificates {
    const message = { ...baseNocCertificates } as NocCertificates
    message.certs = []
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
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
