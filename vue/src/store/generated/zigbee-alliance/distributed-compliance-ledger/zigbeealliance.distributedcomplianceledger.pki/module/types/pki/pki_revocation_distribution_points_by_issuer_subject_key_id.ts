/* eslint-disable */
import { PkiRevocationDistributionPoint } from '../pki/pki_revocation_distribution_point'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface PkiRevocationDistributionPointsByIssuerSubjectKeyID {
  issuerSubjectKeyID: string
  points: PkiRevocationDistributionPoint[]
}

const basePkiRevocationDistributionPointsByIssuerSubjectKeyID: object = { issuerSubjectKeyID: '' }

export const PkiRevocationDistributionPointsByIssuerSubjectKeyID = {
  encode(message: PkiRevocationDistributionPointsByIssuerSubjectKeyID, writer: Writer = Writer.create()): Writer {
    if (message.issuerSubjectKeyID !== '') {
      writer.uint32(10).string(message.issuerSubjectKeyID)
    }
    for (const v of message.points) {
      PkiRevocationDistributionPoint.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): PkiRevocationDistributionPointsByIssuerSubjectKeyID {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePkiRevocationDistributionPointsByIssuerSubjectKeyID } as PkiRevocationDistributionPointsByIssuerSubjectKeyID
    message.points = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.issuerSubjectKeyID = reader.string()
          break
        case 2:
          message.points.push(PkiRevocationDistributionPoint.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): PkiRevocationDistributionPointsByIssuerSubjectKeyID {
    const message = { ...basePkiRevocationDistributionPointsByIssuerSubjectKeyID } as PkiRevocationDistributionPointsByIssuerSubjectKeyID
    message.points = []
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = String(object.issuerSubjectKeyID)
    } else {
      message.issuerSubjectKeyID = ''
    }
    if (object.points !== undefined && object.points !== null) {
      for (const e of object.points) {
        message.points.push(PkiRevocationDistributionPoint.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: PkiRevocationDistributionPointsByIssuerSubjectKeyID): unknown {
    const obj: any = {}
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID)
    if (message.points) {
      obj.points = message.points.map((e) => (e ? PkiRevocationDistributionPoint.toJSON(e) : undefined))
    } else {
      obj.points = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<PkiRevocationDistributionPointsByIssuerSubjectKeyID>): PkiRevocationDistributionPointsByIssuerSubjectKeyID {
    const message = { ...basePkiRevocationDistributionPointsByIssuerSubjectKeyID } as PkiRevocationDistributionPointsByIssuerSubjectKeyID
    message.points = []
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = object.issuerSubjectKeyID
    } else {
      message.issuerSubjectKeyID = ''
    }
    if (object.points !== undefined && object.points !== null) {
      for (const e of object.points) {
        message.points.push(PkiRevocationDistributionPoint.fromPartial(e))
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
