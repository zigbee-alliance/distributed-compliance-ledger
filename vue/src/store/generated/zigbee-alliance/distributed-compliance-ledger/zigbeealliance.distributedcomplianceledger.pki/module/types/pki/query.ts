/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { ApprovedCertificates } from '../pki/approved_certificates'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { ProposedCertificate } from '../pki/proposed_certificate'
import { ChildCertificates } from '../pki/child_certificates'
import { ProposedCertificateRevocation } from '../pki/proposed_certificate_revocation'
import { RevokedCertificates } from '../pki/revoked_certificates'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface QueryGetApprovedCertificatesRequest {
  subject: string
  subjectKeyId: string
}

export interface QueryGetApprovedCertificatesResponse {
  approvedCertificates: ApprovedCertificates | undefined
}

export interface QueryAllApprovedCertificatesRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllApprovedCertificatesResponse {
  approvedCertificates: ApprovedCertificates[]
  pagination: PageResponse | undefined
}

export interface QueryGetProposedCertificateRequest {
  subject: string
  subjectKeyId: string
}

export interface QueryGetProposedCertificateResponse {
  proposedCertificate: ProposedCertificate | undefined
}

export interface QueryAllProposedCertificateRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllProposedCertificateResponse {
  proposedCertificate: ProposedCertificate[]
  pagination: PageResponse | undefined
}

export interface QueryGetChildCertificatesRequest {
  issuer: string
  authorityKeyId: string
}

export interface QueryGetChildCertificatesResponse {
  childCertificates: ChildCertificates | undefined
}

export interface QueryAllChildCertificatesRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllChildCertificatesResponse {
  childCertificates: ChildCertificates[]
  pagination: PageResponse | undefined
}

export interface QueryGetProposedCertificateRevocationRequest {
  subject: string
  subjectKeyId: string
}

export interface QueryGetProposedCertificateRevocationResponse {
  proposedCertificateRevocation: ProposedCertificateRevocation | undefined
}

export interface QueryAllProposedCertificateRevocationRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllProposedCertificateRevocationResponse {
  proposedCertificateRevocation: ProposedCertificateRevocation[]
  pagination: PageResponse | undefined
}

export interface QueryGetRevokedCertificatesRequest {
  subject: string
  subjectKeyId: string
}

export interface QueryGetRevokedCertificatesResponse {
  revokedCertificates: RevokedCertificates | undefined
}

export interface QueryAllRevokedCertificatesRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllRevokedCertificatesResponse {
  revokedCertificates: RevokedCertificates[]
  pagination: PageResponse | undefined
}

const baseQueryGetApprovedCertificatesRequest: object = { subject: '', subjectKeyId: '' }

export const QueryGetApprovedCertificatesRequest = {
  encode(message: QueryGetApprovedCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedCertificatesRequest } as QueryGetApprovedCertificatesRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetApprovedCertificatesRequest {
    const message = { ...baseQueryGetApprovedCertificatesRequest } as QueryGetApprovedCertificatesRequest
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
    return message
  },

  toJSON(message: QueryGetApprovedCertificatesRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetApprovedCertificatesRequest>): QueryGetApprovedCertificatesRequest {
    const message = { ...baseQueryGetApprovedCertificatesRequest } as QueryGetApprovedCertificatesRequest
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
    return message
  }
}

const baseQueryGetApprovedCertificatesResponse: object = {}

export const QueryGetApprovedCertificatesResponse = {
  encode(message: QueryGetApprovedCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.approvedCertificates !== undefined) {
      ApprovedCertificates.encode(message.approvedCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedCertificatesResponse } as QueryGetApprovedCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificates = ApprovedCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetApprovedCertificatesResponse {
    const message = { ...baseQueryGetApprovedCertificatesResponse } as QueryGetApprovedCertificatesResponse
    if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
      message.approvedCertificates = ApprovedCertificates.fromJSON(object.approvedCertificates)
    } else {
      message.approvedCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetApprovedCertificatesResponse): unknown {
    const obj: any = {}
    message.approvedCertificates !== undefined &&
      (obj.approvedCertificates = message.approvedCertificates ? ApprovedCertificates.toJSON(message.approvedCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetApprovedCertificatesResponse>): QueryGetApprovedCertificatesResponse {
    const message = { ...baseQueryGetApprovedCertificatesResponse } as QueryGetApprovedCertificatesResponse
    if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
      message.approvedCertificates = ApprovedCertificates.fromPartial(object.approvedCertificates)
    } else {
      message.approvedCertificates = undefined
    }
    return message
  }
}

const baseQueryAllApprovedCertificatesRequest: object = {}

export const QueryAllApprovedCertificatesRequest = {
  encode(message: QueryAllApprovedCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllApprovedCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllApprovedCertificatesRequest } as QueryAllApprovedCertificatesRequest
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

  fromJSON(object: any): QueryAllApprovedCertificatesRequest {
    const message = { ...baseQueryAllApprovedCertificatesRequest } as QueryAllApprovedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllApprovedCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllApprovedCertificatesRequest>): QueryAllApprovedCertificatesRequest {
    const message = { ...baseQueryAllApprovedCertificatesRequest } as QueryAllApprovedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllApprovedCertificatesResponse: object = {}

export const QueryAllApprovedCertificatesResponse = {
  encode(message: QueryAllApprovedCertificatesResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.approvedCertificates) {
      ApprovedCertificates.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllApprovedCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllApprovedCertificatesResponse } as QueryAllApprovedCertificatesResponse
    message.approvedCertificates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificates.push(ApprovedCertificates.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllApprovedCertificatesResponse {
    const message = { ...baseQueryAllApprovedCertificatesResponse } as QueryAllApprovedCertificatesResponse
    message.approvedCertificates = []
    if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
      for (const e of object.approvedCertificates) {
        message.approvedCertificates.push(ApprovedCertificates.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllApprovedCertificatesResponse): unknown {
    const obj: any = {}
    if (message.approvedCertificates) {
      obj.approvedCertificates = message.approvedCertificates.map((e) => (e ? ApprovedCertificates.toJSON(e) : undefined))
    } else {
      obj.approvedCertificates = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllApprovedCertificatesResponse>): QueryAllApprovedCertificatesResponse {
    const message = { ...baseQueryAllApprovedCertificatesResponse } as QueryAllApprovedCertificatesResponse
    message.approvedCertificates = []
    if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
      for (const e of object.approvedCertificates) {
        message.approvedCertificates.push(ApprovedCertificates.fromPartial(e))
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

const baseQueryGetProposedCertificateRequest: object = { subject: '', subjectKeyId: '' }

export const QueryGetProposedCertificateRequest = {
  encode(message: QueryGetProposedCertificateRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedCertificateRequest } as QueryGetProposedCertificateRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProposedCertificateRequest {
    const message = { ...baseQueryGetProposedCertificateRequest } as QueryGetProposedCertificateRequest
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
    return message
  },

  toJSON(message: QueryGetProposedCertificateRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedCertificateRequest>): QueryGetProposedCertificateRequest {
    const message = { ...baseQueryGetProposedCertificateRequest } as QueryGetProposedCertificateRequest
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
    return message
  }
}

const baseQueryGetProposedCertificateResponse: object = {}

export const QueryGetProposedCertificateResponse = {
  encode(message: QueryGetProposedCertificateResponse, writer: Writer = Writer.create()): Writer {
    if (message.proposedCertificate !== undefined) {
      ProposedCertificate.encode(message.proposedCertificate, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedCertificateResponse } as QueryGetProposedCertificateResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificate = ProposedCertificate.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProposedCertificateResponse {
    const message = { ...baseQueryGetProposedCertificateResponse } as QueryGetProposedCertificateResponse
    if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
      message.proposedCertificate = ProposedCertificate.fromJSON(object.proposedCertificate)
    } else {
      message.proposedCertificate = undefined
    }
    return message
  },

  toJSON(message: QueryGetProposedCertificateResponse): unknown {
    const obj: any = {}
    message.proposedCertificate !== undefined &&
      (obj.proposedCertificate = message.proposedCertificate ? ProposedCertificate.toJSON(message.proposedCertificate) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedCertificateResponse>): QueryGetProposedCertificateResponse {
    const message = { ...baseQueryGetProposedCertificateResponse } as QueryGetProposedCertificateResponse
    if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
      message.proposedCertificate = ProposedCertificate.fromPartial(object.proposedCertificate)
    } else {
      message.proposedCertificate = undefined
    }
    return message
  }
}

const baseQueryAllProposedCertificateRequest: object = {}

export const QueryAllProposedCertificateRequest = {
  encode(message: QueryAllProposedCertificateRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedCertificateRequest } as QueryAllProposedCertificateRequest
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

  fromJSON(object: any): QueryAllProposedCertificateRequest {
    const message = { ...baseQueryAllProposedCertificateRequest } as QueryAllProposedCertificateRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedCertificateRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedCertificateRequest>): QueryAllProposedCertificateRequest {
    const message = { ...baseQueryAllProposedCertificateRequest } as QueryAllProposedCertificateRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllProposedCertificateResponse: object = {}

export const QueryAllProposedCertificateResponse = {
  encode(message: QueryAllProposedCertificateResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.proposedCertificate) {
      ProposedCertificate.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedCertificateResponse } as QueryAllProposedCertificateResponse
    message.proposedCertificate = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificate.push(ProposedCertificate.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllProposedCertificateResponse {
    const message = { ...baseQueryAllProposedCertificateResponse } as QueryAllProposedCertificateResponse
    message.proposedCertificate = []
    if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
      for (const e of object.proposedCertificate) {
        message.proposedCertificate.push(ProposedCertificate.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedCertificateResponse): unknown {
    const obj: any = {}
    if (message.proposedCertificate) {
      obj.proposedCertificate = message.proposedCertificate.map((e) => (e ? ProposedCertificate.toJSON(e) : undefined))
    } else {
      obj.proposedCertificate = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedCertificateResponse>): QueryAllProposedCertificateResponse {
    const message = { ...baseQueryAllProposedCertificateResponse } as QueryAllProposedCertificateResponse
    message.proposedCertificate = []
    if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
      for (const e of object.proposedCertificate) {
        message.proposedCertificate.push(ProposedCertificate.fromPartial(e))
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

const baseQueryGetChildCertificatesRequest: object = { issuer: '', authorityKeyId: '' }

export const QueryGetChildCertificatesRequest = {
  encode(message: QueryGetChildCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.issuer !== '') {
      writer.uint32(10).string(message.issuer)
    }
    if (message.authorityKeyId !== '') {
      writer.uint32(18).string(message.authorityKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChildCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChildCertificatesRequest } as QueryGetChildCertificatesRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string()
          break
        case 2:
          message.authorityKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetChildCertificatesRequest {
    const message = { ...baseQueryGetChildCertificatesRequest } as QueryGetChildCertificatesRequest
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = String(object.issuer)
    } else {
      message.issuer = ''
    }
    if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
      message.authorityKeyId = String(object.authorityKeyId)
    } else {
      message.authorityKeyId = ''
    }
    return message
  },

  toJSON(message: QueryGetChildCertificatesRequest): unknown {
    const obj: any = {}
    message.issuer !== undefined && (obj.issuer = message.issuer)
    message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChildCertificatesRequest>): QueryGetChildCertificatesRequest {
    const message = { ...baseQueryGetChildCertificatesRequest } as QueryGetChildCertificatesRequest
    if (object.issuer !== undefined && object.issuer !== null) {
      message.issuer = object.issuer
    } else {
      message.issuer = ''
    }
    if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
      message.authorityKeyId = object.authorityKeyId
    } else {
      message.authorityKeyId = ''
    }
    return message
  }
}

const baseQueryGetChildCertificatesResponse: object = {}

export const QueryGetChildCertificatesResponse = {
  encode(message: QueryGetChildCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.childCertificates !== undefined) {
      ChildCertificates.encode(message.childCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetChildCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetChildCertificatesResponse } as QueryGetChildCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.childCertificates = ChildCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetChildCertificatesResponse {
    const message = { ...baseQueryGetChildCertificatesResponse } as QueryGetChildCertificatesResponse
    if (object.childCertificates !== undefined && object.childCertificates !== null) {
      message.childCertificates = ChildCertificates.fromJSON(object.childCertificates)
    } else {
      message.childCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetChildCertificatesResponse): unknown {
    const obj: any = {}
    message.childCertificates !== undefined &&
      (obj.childCertificates = message.childCertificates ? ChildCertificates.toJSON(message.childCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetChildCertificatesResponse>): QueryGetChildCertificatesResponse {
    const message = { ...baseQueryGetChildCertificatesResponse } as QueryGetChildCertificatesResponse
    if (object.childCertificates !== undefined && object.childCertificates !== null) {
      message.childCertificates = ChildCertificates.fromPartial(object.childCertificates)
    } else {
      message.childCertificates = undefined
    }
    return message
  }
}

const baseQueryAllChildCertificatesRequest: object = {}

export const QueryAllChildCertificatesRequest = {
  encode(message: QueryAllChildCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllChildCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllChildCertificatesRequest } as QueryAllChildCertificatesRequest
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

  fromJSON(object: any): QueryAllChildCertificatesRequest {
    const message = { ...baseQueryAllChildCertificatesRequest } as QueryAllChildCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllChildCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllChildCertificatesRequest>): QueryAllChildCertificatesRequest {
    const message = { ...baseQueryAllChildCertificatesRequest } as QueryAllChildCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllChildCertificatesResponse: object = {}

export const QueryAllChildCertificatesResponse = {
  encode(message: QueryAllChildCertificatesResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.childCertificates) {
      ChildCertificates.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllChildCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllChildCertificatesResponse } as QueryAllChildCertificatesResponse
    message.childCertificates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.childCertificates.push(ChildCertificates.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllChildCertificatesResponse {
    const message = { ...baseQueryAllChildCertificatesResponse } as QueryAllChildCertificatesResponse
    message.childCertificates = []
    if (object.childCertificates !== undefined && object.childCertificates !== null) {
      for (const e of object.childCertificates) {
        message.childCertificates.push(ChildCertificates.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllChildCertificatesResponse): unknown {
    const obj: any = {}
    if (message.childCertificates) {
      obj.childCertificates = message.childCertificates.map((e) => (e ? ChildCertificates.toJSON(e) : undefined))
    } else {
      obj.childCertificates = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllChildCertificatesResponse>): QueryAllChildCertificatesResponse {
    const message = { ...baseQueryAllChildCertificatesResponse } as QueryAllChildCertificatesResponse
    message.childCertificates = []
    if (object.childCertificates !== undefined && object.childCertificates !== null) {
      for (const e of object.childCertificates) {
        message.childCertificates.push(ChildCertificates.fromPartial(e))
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

const baseQueryGetProposedCertificateRevocationRequest: object = { subject: '', subjectKeyId: '' }

export const QueryGetProposedCertificateRevocationRequest = {
  encode(message: QueryGetProposedCertificateRevocationRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateRevocationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedCertificateRevocationRequest } as QueryGetProposedCertificateRevocationRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProposedCertificateRevocationRequest {
    const message = { ...baseQueryGetProposedCertificateRevocationRequest } as QueryGetProposedCertificateRevocationRequest
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
    return message
  },

  toJSON(message: QueryGetProposedCertificateRevocationRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedCertificateRevocationRequest>): QueryGetProposedCertificateRevocationRequest {
    const message = { ...baseQueryGetProposedCertificateRevocationRequest } as QueryGetProposedCertificateRevocationRequest
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
    return message
  }
}

const baseQueryGetProposedCertificateRevocationResponse: object = {}

export const QueryGetProposedCertificateRevocationResponse = {
  encode(message: QueryGetProposedCertificateRevocationResponse, writer: Writer = Writer.create()): Writer {
    if (message.proposedCertificateRevocation !== undefined) {
      ProposedCertificateRevocation.encode(message.proposedCertificateRevocation, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateRevocationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetProposedCertificateRevocationResponse } as QueryGetProposedCertificateRevocationResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificateRevocation = ProposedCertificateRevocation.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetProposedCertificateRevocationResponse {
    const message = { ...baseQueryGetProposedCertificateRevocationResponse } as QueryGetProposedCertificateRevocationResponse
    if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
      message.proposedCertificateRevocation = ProposedCertificateRevocation.fromJSON(object.proposedCertificateRevocation)
    } else {
      message.proposedCertificateRevocation = undefined
    }
    return message
  },

  toJSON(message: QueryGetProposedCertificateRevocationResponse): unknown {
    const obj: any = {}
    message.proposedCertificateRevocation !== undefined &&
      (obj.proposedCertificateRevocation = message.proposedCertificateRevocation
        ? ProposedCertificateRevocation.toJSON(message.proposedCertificateRevocation)
        : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetProposedCertificateRevocationResponse>): QueryGetProposedCertificateRevocationResponse {
    const message = { ...baseQueryGetProposedCertificateRevocationResponse } as QueryGetProposedCertificateRevocationResponse
    if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
      message.proposedCertificateRevocation = ProposedCertificateRevocation.fromPartial(object.proposedCertificateRevocation)
    } else {
      message.proposedCertificateRevocation = undefined
    }
    return message
  }
}

const baseQueryAllProposedCertificateRevocationRequest: object = {}

export const QueryAllProposedCertificateRevocationRequest = {
  encode(message: QueryAllProposedCertificateRevocationRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateRevocationRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedCertificateRevocationRequest } as QueryAllProposedCertificateRevocationRequest
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

  fromJSON(object: any): QueryAllProposedCertificateRevocationRequest {
    const message = { ...baseQueryAllProposedCertificateRevocationRequest } as QueryAllProposedCertificateRevocationRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedCertificateRevocationRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedCertificateRevocationRequest>): QueryAllProposedCertificateRevocationRequest {
    const message = { ...baseQueryAllProposedCertificateRevocationRequest } as QueryAllProposedCertificateRevocationRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllProposedCertificateRevocationResponse: object = {}

export const QueryAllProposedCertificateRevocationResponse = {
  encode(message: QueryAllProposedCertificateRevocationResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.proposedCertificateRevocation) {
      ProposedCertificateRevocation.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateRevocationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllProposedCertificateRevocationResponse } as QueryAllProposedCertificateRevocationResponse
    message.proposedCertificateRevocation = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificateRevocation.push(ProposedCertificateRevocation.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllProposedCertificateRevocationResponse {
    const message = { ...baseQueryAllProposedCertificateRevocationResponse } as QueryAllProposedCertificateRevocationResponse
    message.proposedCertificateRevocation = []
    if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
      for (const e of object.proposedCertificateRevocation) {
        message.proposedCertificateRevocation.push(ProposedCertificateRevocation.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllProposedCertificateRevocationResponse): unknown {
    const obj: any = {}
    if (message.proposedCertificateRevocation) {
      obj.proposedCertificateRevocation = message.proposedCertificateRevocation.map((e) => (e ? ProposedCertificateRevocation.toJSON(e) : undefined))
    } else {
      obj.proposedCertificateRevocation = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllProposedCertificateRevocationResponse>): QueryAllProposedCertificateRevocationResponse {
    const message = { ...baseQueryAllProposedCertificateRevocationResponse } as QueryAllProposedCertificateRevocationResponse
    message.proposedCertificateRevocation = []
    if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
      for (const e of object.proposedCertificateRevocation) {
        message.proposedCertificateRevocation.push(ProposedCertificateRevocation.fromPartial(e))
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

const baseQueryGetRevokedCertificatesRequest: object = { subject: '', subjectKeyId: '' }

export const QueryGetRevokedCertificatesRequest = {
  encode(message: QueryGetRevokedCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedCertificatesRequest } as QueryGetRevokedCertificatesRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        case 2:
          message.subjectKeyId = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRevokedCertificatesRequest {
    const message = { ...baseQueryGetRevokedCertificatesRequest } as QueryGetRevokedCertificatesRequest
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
    return message
  },

  toJSON(message: QueryGetRevokedCertificatesRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRevokedCertificatesRequest>): QueryGetRevokedCertificatesRequest {
    const message = { ...baseQueryGetRevokedCertificatesRequest } as QueryGetRevokedCertificatesRequest
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
    return message
  }
}

const baseQueryGetRevokedCertificatesResponse: object = {}

export const QueryGetRevokedCertificatesResponse = {
  encode(message: QueryGetRevokedCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.revokedCertificates !== undefined) {
      RevokedCertificates.encode(message.revokedCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedCertificatesResponse } as QueryGetRevokedCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.revokedCertificates = RevokedCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRevokedCertificatesResponse {
    const message = { ...baseQueryGetRevokedCertificatesResponse } as QueryGetRevokedCertificatesResponse
    if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
      message.revokedCertificates = RevokedCertificates.fromJSON(object.revokedCertificates)
    } else {
      message.revokedCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetRevokedCertificatesResponse): unknown {
    const obj: any = {}
    message.revokedCertificates !== undefined &&
      (obj.revokedCertificates = message.revokedCertificates ? RevokedCertificates.toJSON(message.revokedCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRevokedCertificatesResponse>): QueryGetRevokedCertificatesResponse {
    const message = { ...baseQueryGetRevokedCertificatesResponse } as QueryGetRevokedCertificatesResponse
    if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
      message.revokedCertificates = RevokedCertificates.fromPartial(object.revokedCertificates)
    } else {
      message.revokedCertificates = undefined
    }
    return message
  }
}

const baseQueryAllRevokedCertificatesRequest: object = {}

export const QueryAllRevokedCertificatesRequest = {
  encode(message: QueryAllRevokedCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRevokedCertificatesRequest } as QueryAllRevokedCertificatesRequest
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

  fromJSON(object: any): QueryAllRevokedCertificatesRequest {
    const message = { ...baseQueryAllRevokedCertificatesRequest } as QueryAllRevokedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRevokedCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRevokedCertificatesRequest>): QueryAllRevokedCertificatesRequest {
    const message = { ...baseQueryAllRevokedCertificatesRequest } as QueryAllRevokedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllRevokedCertificatesResponse: object = {}

export const QueryAllRevokedCertificatesResponse = {
  encode(message: QueryAllRevokedCertificatesResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.revokedCertificates) {
      RevokedCertificates.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRevokedCertificatesResponse } as QueryAllRevokedCertificatesResponse
    message.revokedCertificates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.revokedCertificates.push(RevokedCertificates.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllRevokedCertificatesResponse {
    const message = { ...baseQueryAllRevokedCertificatesResponse } as QueryAllRevokedCertificatesResponse
    message.revokedCertificates = []
    if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
      for (const e of object.revokedCertificates) {
        message.revokedCertificates.push(RevokedCertificates.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRevokedCertificatesResponse): unknown {
    const obj: any = {}
    if (message.revokedCertificates) {
      obj.revokedCertificates = message.revokedCertificates.map((e) => (e ? RevokedCertificates.toJSON(e) : undefined))
    } else {
      obj.revokedCertificates = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRevokedCertificatesResponse>): QueryAllRevokedCertificatesResponse {
    const message = { ...baseQueryAllRevokedCertificatesResponse } as QueryAllRevokedCertificatesResponse
    message.revokedCertificates = []
    if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
      for (const e of object.revokedCertificates) {
        message.revokedCertificates.push(RevokedCertificates.fromPartial(e))
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
  /** Queries a ApprovedCertificates by index. */
  ApprovedCertificates(request: QueryGetApprovedCertificatesRequest): Promise<QueryGetApprovedCertificatesResponse>
  /** Queries a list of ApprovedCertificates items. */
  ApprovedCertificatesAll(request: QueryAllApprovedCertificatesRequest): Promise<QueryAllApprovedCertificatesResponse>
  /** Queries a ProposedCertificate by index. */
  ProposedCertificate(request: QueryGetProposedCertificateRequest): Promise<QueryGetProposedCertificateResponse>
  /** Queries a list of ProposedCertificate items. */
  ProposedCertificateAll(request: QueryAllProposedCertificateRequest): Promise<QueryAllProposedCertificateResponse>
  /** Queries a ChildCertificates by index. */
  ChildCertificates(request: QueryGetChildCertificatesRequest): Promise<QueryGetChildCertificatesResponse>
  /** Queries a list of ChildCertificates items. */
  ChildCertificatesAll(request: QueryAllChildCertificatesRequest): Promise<QueryAllChildCertificatesResponse>
  /** Queries a ProposedCertificateRevocation by index. */
  ProposedCertificateRevocation(request: QueryGetProposedCertificateRevocationRequest): Promise<QueryGetProposedCertificateRevocationResponse>
  /** Queries a list of ProposedCertificateRevocation items. */
  ProposedCertificateRevocationAll(request: QueryAllProposedCertificateRevocationRequest): Promise<QueryAllProposedCertificateRevocationResponse>
  /** Queries a RevokedCertificates by index. */
  RevokedCertificates(request: QueryGetRevokedCertificatesRequest): Promise<QueryGetRevokedCertificatesResponse>
  /** Queries a list of RevokedCertificates items. */
  RevokedCertificatesAll(request: QueryAllRevokedCertificatesRequest): Promise<QueryAllRevokedCertificatesResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ApprovedCertificates(request: QueryGetApprovedCertificatesRequest): Promise<QueryGetApprovedCertificatesResponse> {
    const data = QueryGetApprovedCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ApprovedCertificates', data)
    return promise.then((data) => QueryGetApprovedCertificatesResponse.decode(new Reader(data)))
  }

  ApprovedCertificatesAll(request: QueryAllApprovedCertificatesRequest): Promise<QueryAllApprovedCertificatesResponse> {
    const data = QueryAllApprovedCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ApprovedCertificatesAll', data)
    return promise.then((data) => QueryAllApprovedCertificatesResponse.decode(new Reader(data)))
  }

  ProposedCertificate(request: QueryGetProposedCertificateRequest): Promise<QueryGetProposedCertificateResponse> {
    const data = QueryGetProposedCertificateRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificate', data)
    return promise.then((data) => QueryGetProposedCertificateResponse.decode(new Reader(data)))
  }

  ProposedCertificateAll(request: QueryAllProposedCertificateRequest): Promise<QueryAllProposedCertificateResponse> {
    const data = QueryAllProposedCertificateRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificateAll', data)
    return promise.then((data) => QueryAllProposedCertificateResponse.decode(new Reader(data)))
  }

  ChildCertificates(request: QueryGetChildCertificatesRequest): Promise<QueryGetChildCertificatesResponse> {
    const data = QueryGetChildCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ChildCertificates', data)
    return promise.then((data) => QueryGetChildCertificatesResponse.decode(new Reader(data)))
  }

  ChildCertificatesAll(request: QueryAllChildCertificatesRequest): Promise<QueryAllChildCertificatesResponse> {
    const data = QueryAllChildCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ChildCertificatesAll', data)
    return promise.then((data) => QueryAllChildCertificatesResponse.decode(new Reader(data)))
  }

  ProposedCertificateRevocation(request: QueryGetProposedCertificateRevocationRequest): Promise<QueryGetProposedCertificateRevocationResponse> {
    const data = QueryGetProposedCertificateRevocationRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificateRevocation', data)
    return promise.then((data) => QueryGetProposedCertificateRevocationResponse.decode(new Reader(data)))
  }

  ProposedCertificateRevocationAll(request: QueryAllProposedCertificateRevocationRequest): Promise<QueryAllProposedCertificateRevocationResponse> {
    const data = QueryAllProposedCertificateRevocationRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificateRevocationAll', data)
    return promise.then((data) => QueryAllProposedCertificateRevocationResponse.decode(new Reader(data)))
  }

  RevokedCertificates(request: QueryGetRevokedCertificatesRequest): Promise<QueryGetRevokedCertificatesResponse> {
    const data = QueryGetRevokedCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RevokedCertificates', data)
    return promise.then((data) => QueryGetRevokedCertificatesResponse.decode(new Reader(data)))
  }

  RevokedCertificatesAll(request: QueryAllRevokedCertificatesRequest): Promise<QueryAllRevokedCertificatesResponse> {
    const data = QueryAllRevokedCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RevokedCertificatesAll', data)
    return promise.then((data) => QueryAllRevokedCertificatesResponse.decode(new Reader(data)))
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
