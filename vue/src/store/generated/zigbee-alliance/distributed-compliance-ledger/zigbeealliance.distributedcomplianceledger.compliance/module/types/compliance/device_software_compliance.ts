/* eslint-disable */
import { ComplianceInfo } from '../compliance/compliance_info'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

export interface DeviceSoftwareCompliance {
  cDCertificateId: string
  complianceInfo: ComplianceInfo[]
}

const baseDeviceSoftwareCompliance: object = { cDCertificateId: '' }

export const DeviceSoftwareCompliance = {
  encode(message: DeviceSoftwareCompliance, writer: Writer = Writer.create()): Writer {
    if (message.cDCertificateId !== '') {
      writer.uint32(10).string(message.cDCertificateId)
    }
    for (const v of message.complianceInfo) {
      ComplianceInfo.encode(v!, writer.uint32(18).fork()).ldelim()
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
          message.cDCertificateId = reader.string()
          break
        case 2:
          message.complianceInfo.push(ComplianceInfo.decode(reader, reader.uint32()))
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
    if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
      message.cDCertificateId = String(object.cDCertificateId)
    } else {
      message.cDCertificateId = ''
    }
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(ComplianceInfo.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: DeviceSoftwareCompliance): unknown {
    const obj: any = {}
    message.cDCertificateId !== undefined && (obj.cDCertificateId = message.cDCertificateId)
    if (message.complianceInfo) {
      obj.complianceInfo = message.complianceInfo.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined))
    } else {
      obj.complianceInfo = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<DeviceSoftwareCompliance>): DeviceSoftwareCompliance {
    const message = { ...baseDeviceSoftwareCompliance } as DeviceSoftwareCompliance
    message.complianceInfo = []
    if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
      message.cDCertificateId = object.cDCertificateId
    } else {
      message.cDCertificateId = ''
    }
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(ComplianceInfo.fromPartial(e))
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
