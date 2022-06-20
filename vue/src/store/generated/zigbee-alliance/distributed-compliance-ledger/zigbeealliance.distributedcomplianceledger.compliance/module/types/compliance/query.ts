/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { ComplianceInfo } from '../compliance/compliance_info'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { CertifiedModel } from '../compliance/certified_model'
import { RevokedModel } from '../compliance/revoked_model'
import { ProvisionalModel } from '../compliance/provisional_model'
import { DeviceSoftwareCompliance } from '../compliance/device_software_compliance'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance'

export interface QueryGetComplianceInfoRequest {
  vid: number
  pid: number
  softwareVersion: number
  certificationType: string
}

export interface QueryGetComplianceInfoResponse {
  complianceInfo: ComplianceInfo | undefined
}

export interface QueryAllComplianceInfoRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllComplianceInfoResponse {
  complianceInfo: ComplianceInfo[]
  pagination: PageResponse | undefined
}

export interface QueryGetCertifiedModelRequest {
  vid: number
  pid: number
  softwareVersion: number
  certificationType: string
}

export interface QueryGetCertifiedModelResponse {
  certifiedModel: CertifiedModel | undefined
}

export interface QueryAllCertifiedModelRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllCertifiedModelResponse {
  certifiedModel: CertifiedModel[]
  pagination: PageResponse | undefined
}

export interface QueryGetRevokedModelRequest {
  vid: number
  pid: number
  softwareVersion: number
  certificationType: string
}

export interface QueryGetRevokedModelResponse {
  revokedModel: RevokedModel | undefined
}

export interface QueryAllRevokedModelRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllRevokedModelResponse {
  revokedModel: RevokedModel[]
  pagination: PageResponse | undefined
}

export interface QueryGetProvisionalModelRequest {
  vid: number
  pid: number
  softwareVersion: number
  certificationType: string
}

export interface QueryGetProvisionalModelResponse {
  provisionalModel: ProvisionalModel | undefined
}

export interface QueryAllProvisionalModelRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllProvisionalModelResponse {
  provisionalModel: ProvisionalModel[]
  pagination: PageResponse | undefined
}

export interface QueryGetDeviceSoftwareComplianceRequest {
  cDCertificateId: string
}

export interface QueryGetDeviceSoftwareComplianceResponse {
  deviceSoftwareCompliance: DeviceSoftwareCompliance | undefined
}

export interface QueryAllDeviceSoftwareComplianceRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllDeviceSoftwareComplianceResponse {
  deviceSoftwareCompliance: DeviceSoftwareCompliance[]
  pagination: PageResponse | undefined
}

const baseQueryGetComplianceInfoRequest: object = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' }

export const QueryGetComplianceInfoRequest = {
  encode(message: QueryGetComplianceInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion)
    }
    if (message.certificationType !== '') {
      writer.uint32(34).string(message.certificationType)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetComplianceInfoRequest } as QueryGetComplianceInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        case 3:
          message.softwareVersion = reader.uint32()
          break
        case 4:
          message.certificationType = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetComplianceInfoRequest {
    const message = { ...baseQueryGetComplianceInfoRequest } as QueryGetComplianceInfoRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = String(object.certificationType)
    } else {
      message.certificationType = ''
    }
    return message
  },

  toJSON(message: QueryGetComplianceInfoRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    message.certificationType !== undefined && (obj.certificationType = message.certificationType)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetComplianceInfoRequest>): QueryGetComplianceInfoRequest {
    const message = { ...baseQueryGetComplianceInfoRequest } as QueryGetComplianceInfoRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = object.certificationType
    } else {
      message.certificationType = ''
    }
    return message
  }
}

const baseQueryGetComplianceInfoResponse: object = {}

export const QueryGetComplianceInfoResponse = {
  encode(message: QueryGetComplianceInfoResponse, writer: Writer = Writer.create()): Writer {
    if (message.complianceInfo !== undefined) {
      ComplianceInfo.encode(message.complianceInfo, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetComplianceInfoResponse } as QueryGetComplianceInfoResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.complianceInfo = ComplianceInfo.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetComplianceInfoResponse {
    const message = { ...baseQueryGetComplianceInfoResponse } as QueryGetComplianceInfoResponse
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      message.complianceInfo = ComplianceInfo.fromJSON(object.complianceInfo)
    } else {
      message.complianceInfo = undefined
    }
    return message
  },

  toJSON(message: QueryGetComplianceInfoResponse): unknown {
    const obj: any = {}
    message.complianceInfo !== undefined && (obj.complianceInfo = message.complianceInfo ? ComplianceInfo.toJSON(message.complianceInfo) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetComplianceInfoResponse>): QueryGetComplianceInfoResponse {
    const message = { ...baseQueryGetComplianceInfoResponse } as QueryGetComplianceInfoResponse
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      message.complianceInfo = ComplianceInfo.fromPartial(object.complianceInfo)
    } else {
      message.complianceInfo = undefined
    }
    return message
  }
}

const baseQueryAllComplianceInfoRequest: object = {}

export const QueryAllComplianceInfoRequest = {
  encode(message: QueryAllComplianceInfoRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllComplianceInfoRequest } as QueryAllComplianceInfoRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllComplianceInfoRequest {
    const message = { ...baseQueryAllComplianceInfoRequest } as QueryAllComplianceInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllComplianceInfoRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllComplianceInfoRequest>): QueryAllComplianceInfoRequest {
    const message = { ...baseQueryAllComplianceInfoRequest } as QueryAllComplianceInfoRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllComplianceInfoResponse: object = {}

export const QueryAllComplianceInfoResponse = {
  encode(message: QueryAllComplianceInfoResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.complianceInfo) {
      ComplianceInfo.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllComplianceInfoResponse } as QueryAllComplianceInfoResponse
    message.complianceInfo = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.complianceInfo.push(ComplianceInfo.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllComplianceInfoResponse {
    const message = { ...baseQueryAllComplianceInfoResponse } as QueryAllComplianceInfoResponse
    message.complianceInfo = []
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(ComplianceInfo.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllComplianceInfoResponse): unknown {
    const obj: any = {}
    if (message.complianceInfo) {
      obj.complianceInfo = message.complianceInfo.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined))
    } else {
      obj.complianceInfo = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllComplianceInfoResponse>): QueryAllComplianceInfoResponse {
    const message = { ...baseQueryAllComplianceInfoResponse } as QueryAllComplianceInfoResponse
    message.complianceInfo = []
    if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
      for (const e of object.complianceInfo) {
        message.complianceInfo.push(ComplianceInfo.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetCertifiedModelRequest: object = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' }

export const QueryGetCertifiedModelRequest = {
  encode(message: QueryGetCertifiedModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion)
    }
    if (message.certificationType !== '') {
      writer.uint32(34).string(message.certificationType)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCertifiedModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCertifiedModelRequest } as QueryGetCertifiedModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        case 3:
          message.softwareVersion = reader.uint32()
          break
        case 4:
          message.certificationType = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetCertifiedModelRequest {
    const message = { ...baseQueryGetCertifiedModelRequest } as QueryGetCertifiedModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = String(object.certificationType)
    } else {
      message.certificationType = ''
    }
    return message
  },

  toJSON(message: QueryGetCertifiedModelRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    message.certificationType !== undefined && (obj.certificationType = message.certificationType)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCertifiedModelRequest>): QueryGetCertifiedModelRequest {
    const message = { ...baseQueryGetCertifiedModelRequest } as QueryGetCertifiedModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = object.certificationType
    } else {
      message.certificationType = ''
    }
    return message
  }
}

const baseQueryGetCertifiedModelResponse: object = {}

export const QueryGetCertifiedModelResponse = {
  encode(message: QueryGetCertifiedModelResponse, writer: Writer = Writer.create()): Writer {
    if (message.certifiedModel !== undefined) {
      CertifiedModel.encode(message.certifiedModel, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetCertifiedModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetCertifiedModelResponse } as QueryGetCertifiedModelResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.certifiedModel = CertifiedModel.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetCertifiedModelResponse {
    const message = { ...baseQueryGetCertifiedModelResponse } as QueryGetCertifiedModelResponse
    if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
      message.certifiedModel = CertifiedModel.fromJSON(object.certifiedModel)
    } else {
      message.certifiedModel = undefined
    }
    return message
  },

  toJSON(message: QueryGetCertifiedModelResponse): unknown {
    const obj: any = {}
    message.certifiedModel !== undefined && (obj.certifiedModel = message.certifiedModel ? CertifiedModel.toJSON(message.certifiedModel) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetCertifiedModelResponse>): QueryGetCertifiedModelResponse {
    const message = { ...baseQueryGetCertifiedModelResponse } as QueryGetCertifiedModelResponse
    if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
      message.certifiedModel = CertifiedModel.fromPartial(object.certifiedModel)
    } else {
      message.certifiedModel = undefined
    }
    return message
  }
}

const baseQueryAllCertifiedModelRequest: object = {}

export const QueryAllCertifiedModelRequest = {
  encode(message: QueryAllCertifiedModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllCertifiedModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllCertifiedModelRequest } as QueryAllCertifiedModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllCertifiedModelRequest {
    const message = { ...baseQueryAllCertifiedModelRequest } as QueryAllCertifiedModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllCertifiedModelRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllCertifiedModelRequest>): QueryAllCertifiedModelRequest {
    const message = { ...baseQueryAllCertifiedModelRequest } as QueryAllCertifiedModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllCertifiedModelResponse: object = {}

export const QueryAllCertifiedModelResponse = {
  encode(message: QueryAllCertifiedModelResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.certifiedModel) {
      CertifiedModel.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllCertifiedModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllCertifiedModelResponse } as QueryAllCertifiedModelResponse
    message.certifiedModel = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.certifiedModel.push(CertifiedModel.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllCertifiedModelResponse {
    const message = { ...baseQueryAllCertifiedModelResponse } as QueryAllCertifiedModelResponse
    message.certifiedModel = []
    if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
      for (const e of object.certifiedModel) {
        message.certifiedModel.push(CertifiedModel.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllCertifiedModelResponse): unknown {
    const obj: any = {}
    if (message.certifiedModel) {
      obj.certifiedModel = message.certifiedModel.map((e) => (e ? CertifiedModel.toJSON(e) : undefined))
    } else {
      obj.certifiedModel = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllCertifiedModelResponse>): QueryAllCertifiedModelResponse {
    const message = { ...baseQueryAllCertifiedModelResponse } as QueryAllCertifiedModelResponse
    message.certifiedModel = []
    if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
      for (const e of object.certifiedModel) {
        message.certifiedModel.push(CertifiedModel.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetRevokedModelRequest: object = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' }

export const QueryGetRevokedModelRequest = {
  encode(message: QueryGetRevokedModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion)
    }
    if (message.certificationType !== '') {
      writer.uint32(34).string(message.certificationType)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedModelRequest } as QueryGetRevokedModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        case 3:
          message.softwareVersion = reader.uint32()
          break
        case 4:
          message.certificationType = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRevokedModelRequest {
    const message = { ...baseQueryGetRevokedModelRequest } as QueryGetRevokedModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = String(object.certificationType)
    } else {
      message.certificationType = ''
    }
    return message
  },

  toJSON(message: QueryGetRevokedModelRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    message.certificationType !== undefined && (obj.certificationType = message.certificationType)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRevokedModelRequest>): QueryGetRevokedModelRequest {
    const message = { ...baseQueryGetRevokedModelRequest } as QueryGetRevokedModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = object.certificationType
    } else {
      message.certificationType = ''
    }
    return message
  }
}

const baseQueryGetRevokedModelResponse: object = {}

export const QueryGetRevokedModelResponse = {
  encode(message: QueryGetRevokedModelResponse, writer: Writer = Writer.create()): Writer {
    if (message.revokedModel !== undefined) {
      RevokedModel.encode(message.revokedModel, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedModelResponse } as QueryGetRevokedModelResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.revokedModel = RevokedModel.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRevokedModelResponse {
    const message = { ...baseQueryGetRevokedModelResponse } as QueryGetRevokedModelResponse
    if (object.revokedModel !== undefined && object.revokedModel !== null) {
      message.revokedModel = RevokedModel.fromJSON(object.revokedModel)
    } else {
      message.revokedModel = undefined
    }
    return message
  },

  toJSON(message: QueryGetRevokedModelResponse): unknown {
    const obj: any = {}
    message.revokedModel !== undefined && (obj.revokedModel = message.revokedModel ? RevokedModel.toJSON(message.revokedModel) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRevokedModelResponse>): QueryGetRevokedModelResponse {
    const message = { ...baseQueryGetRevokedModelResponse } as QueryGetRevokedModelResponse
    if (object.revokedModel !== undefined && object.revokedModel !== null) {
      message.revokedModel = RevokedModel.fromPartial(object.revokedModel)
    } else {
      message.revokedModel = undefined
    }
    return message
  }
}

const baseQueryAllRevokedModelRequest: object = {}

export const QueryAllRevokedModelRequest = {
  encode(message: QueryAllRevokedModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRevokedModelRequest } as QueryAllRevokedModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllRevokedModelRequest {
    const message = { ...baseQueryAllRevokedModelRequest } as QueryAllRevokedModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRevokedModelRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRevokedModelRequest>): QueryAllRevokedModelRequest {
    const message = { ...baseQueryAllRevokedModelRequest } as QueryAllRevokedModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllRevokedModelResponse: object = {}

export const QueryAllRevokedModelResponse = {
  encode(message: QueryAllRevokedModelResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.revokedModel) {
      RevokedModel.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRevokedModelResponse } as QueryAllRevokedModelResponse
    message.revokedModel = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.revokedModel.push(RevokedModel.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllRevokedModelResponse {
    const message = { ...baseQueryAllRevokedModelResponse } as QueryAllRevokedModelResponse
    message.revokedModel = []
    if (object.revokedModel !== undefined && object.revokedModel !== null) {
      for (const e of object.revokedModel) {
        message.revokedModel.push(RevokedModel.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRevokedModelResponse): unknown {
    const obj: any = {}
    if (message.revokedModel) {
      obj.revokedModel = message.revokedModel.map((e) => (e ? RevokedModel.toJSON(e) : undefined))
    } else {
      obj.revokedModel = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRevokedModelResponse>): QueryAllRevokedModelResponse {
    const message = { ...baseQueryAllRevokedModelResponse } as QueryAllRevokedModelResponse
    message.revokedModel = []
    if (object.revokedModel !== undefined && object.revokedModel !== null) {
      for (const e of object.revokedModel) {
        message.revokedModel.push(RevokedModel.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetProvisionalModelRequest: object = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' }

export const QueryGetProvisionalModelRequest = {
  encode(message: QueryGetProvisionalModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion)
    }
    if (message.certificationType !== '') {
      writer.uint32(34).string(message.certificationType)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProvisionalModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProvisionalModelRequest } as QueryGetProvisionalModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        case 3:
          message.softwareVersion = reader.uint32()
          break
        case 4:
          message.certificationType = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProvisionalModelRequest {
    const message = { ...baseQueryGetProvisionalModelRequest } as QueryGetProvisionalModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = String(object.certificationType)
    } else {
      message.certificationType = ''
    }
    return message
  },

  toJSON(message: QueryGetProvisionalModelRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    message.certificationType !== undefined && (obj.certificationType = message.certificationType)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProvisionalModelRequest>): QueryGetProvisionalModelRequest {
    const message = { ...baseQueryGetProvisionalModelRequest } as QueryGetProvisionalModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    if (object.certificationType !== undefined && object.certificationType !== null) {
      message.certificationType = object.certificationType
    } else {
      message.certificationType = ''
    }
    return message
  }
}

const baseQueryGetProvisionalModelResponse: object = {}

export const QueryGetProvisionalModelResponse = {
  encode(message: QueryGetProvisionalModelResponse, writer: Writer = Writer.create()): Writer {
    if (message.provisionalModel !== undefined) {
      ProvisionalModel.encode(message.provisionalModel, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProvisionalModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProvisionalModelResponse } as QueryGetProvisionalModelResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.provisionalModel = ProvisionalModel.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProvisionalModelResponse {
    const message = { ...baseQueryGetProvisionalModelResponse } as QueryGetProvisionalModelResponse
    if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
      message.provisionalModel = ProvisionalModel.fromJSON(object.provisionalModel)
    } else {
      message.provisionalModel = undefined
    }
    return message
  },

  toJSON(message: QueryGetProvisionalModelResponse): unknown {
    const obj: any = {}
    message.provisionalModel !== undefined && (obj.provisionalModel = message.provisionalModel ? ProvisionalModel.toJSON(message.provisionalModel) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProvisionalModelResponse>): QueryGetProvisionalModelResponse {
    const message = { ...baseQueryGetProvisionalModelResponse } as QueryGetProvisionalModelResponse
    if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
      message.provisionalModel = ProvisionalModel.fromPartial(object.provisionalModel)
    } else {
      message.provisionalModel = undefined
    }
    return message
  }
}

const baseQueryAllProvisionalModelRequest: object = {}

export const QueryAllProvisionalModelRequest = {
  encode(message: QueryAllProvisionalModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProvisionalModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProvisionalModelRequest } as QueryAllProvisionalModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllProvisionalModelRequest {
    const message = { ...baseQueryAllProvisionalModelRequest } as QueryAllProvisionalModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProvisionalModelRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProvisionalModelRequest>): QueryAllProvisionalModelRequest {
    const message = { ...baseQueryAllProvisionalModelRequest } as QueryAllProvisionalModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllProvisionalModelResponse: object = {}

export const QueryAllProvisionalModelResponse = {
  encode(message: QueryAllProvisionalModelResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.provisionalModel) {
      ProvisionalModel.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProvisionalModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProvisionalModelResponse } as QueryAllProvisionalModelResponse
    message.provisionalModel = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.provisionalModel.push(ProvisionalModel.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllProvisionalModelResponse {
    const message = { ...baseQueryAllProvisionalModelResponse } as QueryAllProvisionalModelResponse
    message.provisionalModel = []
    if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
      for (const e of object.provisionalModel) {
        message.provisionalModel.push(ProvisionalModel.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProvisionalModelResponse): unknown {
    const obj: any = {}
    if (message.provisionalModel) {
      obj.provisionalModel = message.provisionalModel.map((e) => (e ? ProvisionalModel.toJSON(e) : undefined))
    } else {
      obj.provisionalModel = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProvisionalModelResponse>): QueryAllProvisionalModelResponse {
    const message = { ...baseQueryAllProvisionalModelResponse } as QueryAllProvisionalModelResponse
    message.provisionalModel = []
    if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
      for (const e of object.provisionalModel) {
        message.provisionalModel.push(ProvisionalModel.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetDeviceSoftwareComplianceRequest: object = { cDCertificateId: '' }

export const QueryGetDeviceSoftwareComplianceRequest = {
  encode(message: QueryGetDeviceSoftwareComplianceRequest, writer: Writer = Writer.create()): Writer {
    if (message.cDCertificateId !== '') {
      writer.uint32(10).string(message.cDCertificateId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDeviceSoftwareComplianceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetDeviceSoftwareComplianceRequest } as QueryGetDeviceSoftwareComplianceRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.cDCertificateId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetDeviceSoftwareComplianceRequest {
    const message = { ...baseQueryGetDeviceSoftwareComplianceRequest } as QueryGetDeviceSoftwareComplianceRequest
    if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
      message.cDCertificateId = String(object.cDCertificateId)
    } else {
      message.cDCertificateId = ''
    }
    return message
  },

  toJSON(message: QueryGetDeviceSoftwareComplianceRequest): unknown {
    const obj: any = {}
    message.cDCertificateId !== undefined && (obj.cDCertificateId = message.cDCertificateId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetDeviceSoftwareComplianceRequest>): QueryGetDeviceSoftwareComplianceRequest {
    const message = { ...baseQueryGetDeviceSoftwareComplianceRequest } as QueryGetDeviceSoftwareComplianceRequest
    if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
      message.cDCertificateId = object.cDCertificateId
    } else {
      message.cDCertificateId = ''
    }
    return message
  }
}

const baseQueryGetDeviceSoftwareComplianceResponse: object = {}

export const QueryGetDeviceSoftwareComplianceResponse = {
  encode(message: QueryGetDeviceSoftwareComplianceResponse, writer: Writer = Writer.create()): Writer {
    if (message.deviceSoftwareCompliance !== undefined) {
      DeviceSoftwareCompliance.encode(message.deviceSoftwareCompliance, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetDeviceSoftwareComplianceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetDeviceSoftwareComplianceResponse } as QueryGetDeviceSoftwareComplianceResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.deviceSoftwareCompliance = DeviceSoftwareCompliance.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetDeviceSoftwareComplianceResponse {
    const message = { ...baseQueryGetDeviceSoftwareComplianceResponse } as QueryGetDeviceSoftwareComplianceResponse
    if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
      message.deviceSoftwareCompliance = DeviceSoftwareCompliance.fromJSON(object.deviceSoftwareCompliance)
    } else {
      message.deviceSoftwareCompliance = undefined
    }
    return message
  },

  toJSON(message: QueryGetDeviceSoftwareComplianceResponse): unknown {
    const obj: any = {}
    message.deviceSoftwareCompliance !== undefined &&
      (obj.deviceSoftwareCompliance = message.deviceSoftwareCompliance ? DeviceSoftwareCompliance.toJSON(message.deviceSoftwareCompliance) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetDeviceSoftwareComplianceResponse>): QueryGetDeviceSoftwareComplianceResponse {
    const message = { ...baseQueryGetDeviceSoftwareComplianceResponse } as QueryGetDeviceSoftwareComplianceResponse
    if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
      message.deviceSoftwareCompliance = DeviceSoftwareCompliance.fromPartial(object.deviceSoftwareCompliance)
    } else {
      message.deviceSoftwareCompliance = undefined
    }
    return message
  }
}

const baseQueryAllDeviceSoftwareComplianceRequest: object = {}

export const QueryAllDeviceSoftwareComplianceRequest = {
  encode(message: QueryAllDeviceSoftwareComplianceRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDeviceSoftwareComplianceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllDeviceSoftwareComplianceRequest } as QueryAllDeviceSoftwareComplianceRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllDeviceSoftwareComplianceRequest {
    const message = { ...baseQueryAllDeviceSoftwareComplianceRequest } as QueryAllDeviceSoftwareComplianceRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllDeviceSoftwareComplianceRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllDeviceSoftwareComplianceRequest>): QueryAllDeviceSoftwareComplianceRequest {
    const message = { ...baseQueryAllDeviceSoftwareComplianceRequest } as QueryAllDeviceSoftwareComplianceRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllDeviceSoftwareComplianceResponse: object = {}

export const QueryAllDeviceSoftwareComplianceResponse = {
  encode(message: QueryAllDeviceSoftwareComplianceResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.deviceSoftwareCompliance) {
      DeviceSoftwareCompliance.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDeviceSoftwareComplianceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllDeviceSoftwareComplianceResponse } as QueryAllDeviceSoftwareComplianceResponse
    message.deviceSoftwareCompliance = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.deviceSoftwareCompliance.push(DeviceSoftwareCompliance.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllDeviceSoftwareComplianceResponse {
    const message = { ...baseQueryAllDeviceSoftwareComplianceResponse } as QueryAllDeviceSoftwareComplianceResponse
    message.deviceSoftwareCompliance = []
    if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
      for (const e of object.deviceSoftwareCompliance) {
        message.deviceSoftwareCompliance.push(DeviceSoftwareCompliance.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllDeviceSoftwareComplianceResponse): unknown {
    const obj: any = {}
    if (message.deviceSoftwareCompliance) {
      obj.deviceSoftwareCompliance = message.deviceSoftwareCompliance.map((e) => (e ? DeviceSoftwareCompliance.toJSON(e) : undefined))
    } else {
      obj.deviceSoftwareCompliance = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllDeviceSoftwareComplianceResponse>): QueryAllDeviceSoftwareComplianceResponse {
    const message = { ...baseQueryAllDeviceSoftwareComplianceResponse } as QueryAllDeviceSoftwareComplianceResponse
    message.deviceSoftwareCompliance = []
    if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
      for (const e of object.deviceSoftwareCompliance) {
        message.deviceSoftwareCompliance.push(DeviceSoftwareCompliance.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a ComplianceInfo by index. */
  ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse>
  /** Queries a list of ComplianceInfo items. */
  ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse>
  /** Queries a CertifiedModel by index. */
  CertifiedModel(request: QueryGetCertifiedModelRequest): Promise<QueryGetCertifiedModelResponse>
  /** Queries a list of CertifiedModel items. */
  CertifiedModelAll(request: QueryAllCertifiedModelRequest): Promise<QueryAllCertifiedModelResponse>
  /** Queries a RevokedModel by index. */
  RevokedModel(request: QueryGetRevokedModelRequest): Promise<QueryGetRevokedModelResponse>
  /** Queries a list of RevokedModel items. */
  RevokedModelAll(request: QueryAllRevokedModelRequest): Promise<QueryAllRevokedModelResponse>
  /** Queries a ProvisionalModel by index. */
  ProvisionalModel(request: QueryGetProvisionalModelRequest): Promise<QueryGetProvisionalModelResponse>
  /** Queries a list of ProvisionalModel items. */
  ProvisionalModelAll(request: QueryAllProvisionalModelRequest): Promise<QueryAllProvisionalModelResponse>
  /** Queries a DeviceSoftwareCompliance by index. */
  DeviceSoftwareCompliance(request: QueryGetDeviceSoftwareComplianceRequest): Promise<QueryGetDeviceSoftwareComplianceResponse>
  /** Queries a list of DeviceSoftwareCompliance items. */
  DeviceSoftwareComplianceAll(request: QueryAllDeviceSoftwareComplianceRequest): Promise<QueryAllDeviceSoftwareComplianceResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse> {
    const data = QueryGetComplianceInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfo', data)
    return promise.then((data) => QueryGetComplianceInfoResponse.decode(new Reader(data)))
  }

  ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse> {
    const data = QueryAllComplianceInfoRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfoAll', data)
    return promise.then((data) => QueryAllComplianceInfoResponse.decode(new Reader(data)))
  }

  CertifiedModel(request: QueryGetCertifiedModelRequest): Promise<QueryGetCertifiedModelResponse> {
    const data = QueryGetCertifiedModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'CertifiedModel', data)
    return promise.then((data) => QueryGetCertifiedModelResponse.decode(new Reader(data)))
  }

  CertifiedModelAll(request: QueryAllCertifiedModelRequest): Promise<QueryAllCertifiedModelResponse> {
    const data = QueryAllCertifiedModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'CertifiedModelAll', data)
    return promise.then((data) => QueryAllCertifiedModelResponse.decode(new Reader(data)))
  }

  RevokedModel(request: QueryGetRevokedModelRequest): Promise<QueryGetRevokedModelResponse> {
    const data = QueryGetRevokedModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'RevokedModel', data)
    return promise.then((data) => QueryGetRevokedModelResponse.decode(new Reader(data)))
  }

  RevokedModelAll(request: QueryAllRevokedModelRequest): Promise<QueryAllRevokedModelResponse> {
    const data = QueryAllRevokedModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'RevokedModelAll', data)
    return promise.then((data) => QueryAllRevokedModelResponse.decode(new Reader(data)))
  }

  ProvisionalModel(request: QueryGetProvisionalModelRequest): Promise<QueryGetProvisionalModelResponse> {
    const data = QueryGetProvisionalModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ProvisionalModel', data)
    return promise.then((data) => QueryGetProvisionalModelResponse.decode(new Reader(data)))
  }

  ProvisionalModelAll(request: QueryAllProvisionalModelRequest): Promise<QueryAllProvisionalModelResponse> {
    const data = QueryAllProvisionalModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ProvisionalModelAll', data)
    return promise.then((data) => QueryAllProvisionalModelResponse.decode(new Reader(data)))
  }

  DeviceSoftwareCompliance(request: QueryGetDeviceSoftwareComplianceRequest): Promise<QueryGetDeviceSoftwareComplianceResponse> {
    const data = QueryGetDeviceSoftwareComplianceRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'DeviceSoftwareCompliance', data)
    return promise.then((data) => QueryGetDeviceSoftwareComplianceResponse.decode(new Reader(data)))
  }

  DeviceSoftwareComplianceAll(request: QueryAllDeviceSoftwareComplianceRequest): Promise<QueryAllDeviceSoftwareComplianceResponse> {
    const data = QueryAllDeviceSoftwareComplianceRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'DeviceSoftwareComplianceAll', data)
    return promise.then((data) => QueryAllDeviceSoftwareComplianceResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
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
