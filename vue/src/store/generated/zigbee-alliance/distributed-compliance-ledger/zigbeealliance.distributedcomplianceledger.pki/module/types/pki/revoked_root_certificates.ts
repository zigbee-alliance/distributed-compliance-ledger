/* eslint-disable */
import { CertificateIdentifier } from '../pki/certificate_identifier'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface RevokedRootCertificates {
  certs: CertificateIdentifier[]
}

const baseRevokedRootCertificates: object = {}

export const RevokedRootCertificates = {
  encode(message: RevokedRootCertificates, writer: Writer = Writer.create()): Writer {
    for (const v of message.certs) {
      CertificateIdentifier.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): RevokedRootCertificates {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseRevokedRootCertificates } as RevokedRootCertificates
    message.certs = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.certs.push(CertificateIdentifier.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): RevokedRootCertificates {
    const message = { ...baseRevokedRootCertificates } as RevokedRootCertificates
    message.certs = []
    if (object.certs !== undefined && object.certs !== null) {
      for (const e of object.certs) {
        message.certs.push(CertificateIdentifier.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: RevokedRootCertificates): unknown {
    const obj: any = {}
    if (message.certs) {
      obj.certs = message.certs.map((e) => (e ? CertificateIdentifier.toJSON(e) : undefined))
    } else {
      obj.certs = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<RevokedRootCertificates>): RevokedRootCertificates {
    const message = { ...baseRevokedRootCertificates } as RevokedRootCertificates
    message.certs = []
    if (object.certs !== undefined && object.certs !== null) {
      for (const e of object.certs) {
        message.certs.push(CertificateIdentifier.fromPartial(e))
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
