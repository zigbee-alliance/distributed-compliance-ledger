/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface PkiRevocationDistributionPointsByIssuerSubjectKeyID {
  issuerSubjectKeyID: string
}

const basePkiRevocationDistributionPointsByIssuerSubjectKeyID: object = { issuerSubjectKeyID: '' }

export const PkiRevocationDistributionPointsByIssuerSubjectKeyID = {
  encode(message: PkiRevocationDistributionPointsByIssuerSubjectKeyID, writer: Writer = Writer.create()): Writer {
    if (message.issuerSubjectKeyID !== '') {
      writer.uint32(10).string(message.issuerSubjectKeyID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): PkiRevocationDistributionPointsByIssuerSubjectKeyID {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...basePkiRevocationDistributionPointsByIssuerSubjectKeyID } as PkiRevocationDistributionPointsByIssuerSubjectKeyID
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.issuerSubjectKeyID = reader.string()
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
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = String(object.issuerSubjectKeyID)
    } else {
      message.issuerSubjectKeyID = ''
    }
    return message
  },

  toJSON(message: PkiRevocationDistributionPointsByIssuerSubjectKeyID): unknown {
    const obj: any = {}
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID)
    return obj
  },

  fromPartial(object: DeepPartial<PkiRevocationDistributionPointsByIssuerSubjectKeyID>): PkiRevocationDistributionPointsByIssuerSubjectKeyID {
    const message = { ...basePkiRevocationDistributionPointsByIssuerSubjectKeyID } as PkiRevocationDistributionPointsByIssuerSubjectKeyID
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = object.issuerSubjectKeyID
    } else {
      message.issuerSubjectKeyID = ''
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
