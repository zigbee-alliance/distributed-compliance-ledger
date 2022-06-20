/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

export interface DeviceSoftwareCompliance {
  cdCertificateId: string
  complianceInfo: string[]
}

const baseDeviceSoftwareCompliance: object = { cdCertificateId: '', complianceInfo: '' }

export const DeviceSoftwareCompliance = {
  encode(message: DeviceSoftwareCompliance, writer: Writer = Writer.create()): Writer {
    if (message.cdCertificateId !== '') {
      writer.uint32(10).string(message.cdCertificateId)
    }
    for (const v of message.complianceInfo) {
      writer.uint32(18).string(v!)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): DeviceSoftwareCompliance {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseDeviceSoftwareCompliance } as DeviceSoftwareCompliance
    message.complianceInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.cdCertificateId = reader.string()
          break
        case 2:
          message.complianceInfo.push(reader.string())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): DeviceSoftwareCompliance {
    const message = { ...baseDeviceSoftwareCompliance } as DeviceSoftwareCompliance
    message.complianceInfo = []
    if (object.cdCertificateId !== undefined && object.cdCertificateId !== null) {
      message.cdCertificateId = String(object.cdCertificateId)
    } else {
      message.cdCertificateId = ''
    }
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(String(e))
      }
    }
    return message
  },

  toJSON(message: DeviceSoftwareCompliance): unknown {
    const obj: any = {}
    message.cdCertificateId !== undefined && (obj.cdCertificateId = message.cdCertificateId)
    if (message.complianceInfo) {
      obj.complianceInfo = message.complianceInfo.map((e) => e)
    } else {
      obj.complianceInfo = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<DeviceSoftwareCompliance>): DeviceSoftwareCompliance {
    const message = { ...baseDeviceSoftwareCompliance } as DeviceSoftwareCompliance
    message.complianceInfo = []
    if (object.cdCertificateId !== undefined && object.cdCertificateId !== null) {
      message.cdCertificateId = object.cdCertificateId
    } else {
      message.cdCertificateId = ''
    }
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(e)
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
