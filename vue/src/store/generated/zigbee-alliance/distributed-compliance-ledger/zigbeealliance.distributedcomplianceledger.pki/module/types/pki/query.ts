/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { ApprovedCertificates } from '../pki/approved_certificates'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { ProposedCertificate } from '../pki/proposed_certificate'
import { ChildCertificates } from '../pki/child_certificates'
import { ProposedCertificateRevocation } from '../pki/proposed_certificate_revocation'
import { RevokedCertificates } from '../pki/revoked_certificates'
import { ApprovedRootCertificates } from '../pki/approved_root_certificates'
import { RevokedRootCertificates } from '../pki/revoked_root_certificates'
import { ApprovedCertificatesBySubject } from '../pki/approved_certificates_by_subject'
import { RejectedCertificate } from '../pki/rejected_certificate'
import { PkiRevocationDistributionPoint } from '../pki/pki_revocation_distribution_point'
import { PkiRevocationDistributionPointsByIssuerSubjectKeyID } from '../pki/pki_revocation_distribution_points_by_issuer_subject_key_id'
import { NocRootCertificates } from '../pki/noc_root_certificates'
import { NocIcaCertificates } from '../pki/noc_ica_certificates'
import { RevokedNocRootCertificates } from '../pki/revoked_noc_root_certificates'

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
  subjectKeyId: string
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

export interface QueryGetProposedCertificateRevocationRequest {
  subject: string
  subjectKeyId: string
  serialNumber: string
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

export interface QueryGetApprovedRootCertificatesRequest {}

export interface QueryGetApprovedRootCertificatesResponse {
  approvedRootCertificates: ApprovedRootCertificates | undefined
}

export interface QueryGetRevokedRootCertificatesRequest {}

export interface QueryGetRevokedRootCertificatesResponse {
  revokedRootCertificates: RevokedRootCertificates | undefined
}

export interface QueryGetApprovedCertificatesBySubjectRequest {
  subject: string
}

export interface QueryGetApprovedCertificatesBySubjectResponse {
  approvedCertificatesBySubject: ApprovedCertificatesBySubject | undefined
}

export interface QueryGetRejectedCertificatesRequest {
  subject: string
  subjectKeyId: string
}

export interface QueryGetRejectedCertificatesResponse {
  rejectedCertificate: RejectedCertificate | undefined
}

export interface QueryAllRejectedCertificatesRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllRejectedCertificatesResponse {
  rejectedCertificate: RejectedCertificate[]
  pagination: PageResponse | undefined
}

export interface QueryGetPkiRevocationDistributionPointRequest {
  vid: number
  label: string
  issuerSubjectKeyID: string
}

export interface QueryGetPkiRevocationDistributionPointResponse {
  PkiRevocationDistributionPoint: PkiRevocationDistributionPoint | undefined
}

export interface QueryAllPkiRevocationDistributionPointRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllPkiRevocationDistributionPointResponse {
  PkiRevocationDistributionPoint: PkiRevocationDistributionPoint[]
  pagination: PageResponse | undefined
}

export interface QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
  issuerSubjectKeyID: string
}

export interface QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
  pkiRevocationDistributionPointsByIssuerSubjectKeyID: PkiRevocationDistributionPointsByIssuerSubjectKeyID | undefined
}

export interface QueryGetNocRootCertificatesRequest {
  vid: number
}

export interface QueryGetNocRootCertificatesResponse {
  nocRootCertificates: NocRootCertificates | undefined
}

export interface QueryAllNocRootCertificatesRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllNocRootCertificatesResponse {
  nocRootCertificates: NocRootCertificates[]
  pagination: PageResponse | undefined
}

export interface QueryGetNocIcaCertificatesRequest {
  vid: number
}

export interface QueryGetNocIcaCertificatesResponse {
  nocIcaCertificates: NocIcaCertificates | undefined
}

export interface QueryAllNocIcaCertificatesRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllNocIcaCertificatesResponse {
  nocIcaCertificates: NocIcaCertificates[]
  pagination: PageResponse | undefined
}

export interface QueryGetRevokedNocRootCertificatesRequest {
  subject: string
  subjectKeyId: string
}

export interface QueryGetRevokedNocRootCertificatesResponse {
  revokedNocRootCertificates: RevokedNocRootCertificates | undefined
}

export interface QueryAllRevokedNocRootCertificatesRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllRevokedNocRootCertificatesResponse {
  revokedNocRootCertificates: RevokedNocRootCertificates[]
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

const baseQueryAllApprovedCertificatesRequest: object = { subjectKeyId: '' }

export const QueryAllApprovedCertificatesRequest = {
  encode(message: QueryAllApprovedCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
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

  fromJSON(object: any): QueryAllApprovedCertificatesRequest {
    const message = { ...baseQueryAllApprovedCertificatesRequest } as QueryAllApprovedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = String(object.subjectKeyId)
    } else {
      message.subjectKeyId = ''
    }
    return message
  },

  toJSON(message: QueryAllApprovedCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllApprovedCertificatesRequest>): QueryAllApprovedCertificatesRequest {
    const message = { ...baseQueryAllApprovedCertificatesRequest } as QueryAllApprovedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
      message.subjectKeyId = object.subjectKeyId
    } else {
      message.subjectKeyId = ''
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

const baseQueryGetProposedCertificateRevocationRequest: object = { subject: '', subjectKeyId: '', serialNumber: '' }

export const QueryGetProposedCertificateRevocationRequest = {
  encode(message: QueryGetProposedCertificateRevocationRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    if (message.serialNumber !== '') {
      writer.uint32(26).string(message.serialNumber)
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
        case 3:
          message.serialNumber = reader.string()
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
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = String(object.serialNumber)
    } else {
      message.serialNumber = ''
    }
    return message
  },

  toJSON(message: QueryGetProposedCertificateRevocationRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber)
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
    if (object.serialNumber !== undefined && object.serialNumber !== null) {
      message.serialNumber = object.serialNumber
    } else {
      message.serialNumber = ''
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

const baseQueryGetApprovedRootCertificatesRequest: object = {}

export const QueryGetApprovedRootCertificatesRequest = {
  encode(_: QueryGetApprovedRootCertificatesRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedRootCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedRootCertificatesRequest } as QueryGetApprovedRootCertificatesRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryGetApprovedRootCertificatesRequest {
    const message = { ...baseQueryGetApprovedRootCertificatesRequest } as QueryGetApprovedRootCertificatesRequest
    return message
  },

  toJSON(_: QueryGetApprovedRootCertificatesRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetApprovedRootCertificatesRequest>): QueryGetApprovedRootCertificatesRequest {
    const message = { ...baseQueryGetApprovedRootCertificatesRequest } as QueryGetApprovedRootCertificatesRequest
    return message
  }
}

const baseQueryGetApprovedRootCertificatesResponse: object = {}

export const QueryGetApprovedRootCertificatesResponse = {
  encode(message: QueryGetApprovedRootCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.approvedRootCertificates !== undefined) {
      ApprovedRootCertificates.encode(message.approvedRootCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedRootCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedRootCertificatesResponse } as QueryGetApprovedRootCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.approvedRootCertificates = ApprovedRootCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetApprovedRootCertificatesResponse {
    const message = { ...baseQueryGetApprovedRootCertificatesResponse } as QueryGetApprovedRootCertificatesResponse
    if (object.approvedRootCertificates !== undefined && object.approvedRootCertificates !== null) {
      message.approvedRootCertificates = ApprovedRootCertificates.fromJSON(object.approvedRootCertificates)
    } else {
      message.approvedRootCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetApprovedRootCertificatesResponse): unknown {
    const obj: any = {}
    message.approvedRootCertificates !== undefined &&
      (obj.approvedRootCertificates = message.approvedRootCertificates ? ApprovedRootCertificates.toJSON(message.approvedRootCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetApprovedRootCertificatesResponse>): QueryGetApprovedRootCertificatesResponse {
    const message = { ...baseQueryGetApprovedRootCertificatesResponse } as QueryGetApprovedRootCertificatesResponse
    if (object.approvedRootCertificates !== undefined && object.approvedRootCertificates !== null) {
      message.approvedRootCertificates = ApprovedRootCertificates.fromPartial(object.approvedRootCertificates)
    } else {
      message.approvedRootCertificates = undefined
    }
    return message
  }
}

const baseQueryGetRevokedRootCertificatesRequest: object = {}

export const QueryGetRevokedRootCertificatesRequest = {
  encode(_: QueryGetRevokedRootCertificatesRequest, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedRootCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedRootCertificatesRequest } as QueryGetRevokedRootCertificatesRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): QueryGetRevokedRootCertificatesRequest {
    const message = { ...baseQueryGetRevokedRootCertificatesRequest } as QueryGetRevokedRootCertificatesRequest
    return message
  },

  toJSON(_: QueryGetRevokedRootCertificatesRequest): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<QueryGetRevokedRootCertificatesRequest>): QueryGetRevokedRootCertificatesRequest {
    const message = { ...baseQueryGetRevokedRootCertificatesRequest } as QueryGetRevokedRootCertificatesRequest
    return message
  }
}

const baseQueryGetRevokedRootCertificatesResponse: object = {}

export const QueryGetRevokedRootCertificatesResponse = {
  encode(message: QueryGetRevokedRootCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.revokedRootCertificates !== undefined) {
      RevokedRootCertificates.encode(message.revokedRootCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedRootCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedRootCertificatesResponse } as QueryGetRevokedRootCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.revokedRootCertificates = RevokedRootCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRevokedRootCertificatesResponse {
    const message = { ...baseQueryGetRevokedRootCertificatesResponse } as QueryGetRevokedRootCertificatesResponse
    if (object.revokedRootCertificates !== undefined && object.revokedRootCertificates !== null) {
      message.revokedRootCertificates = RevokedRootCertificates.fromJSON(object.revokedRootCertificates)
    } else {
      message.revokedRootCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetRevokedRootCertificatesResponse): unknown {
    const obj: any = {}
    message.revokedRootCertificates !== undefined &&
      (obj.revokedRootCertificates = message.revokedRootCertificates ? RevokedRootCertificates.toJSON(message.revokedRootCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRevokedRootCertificatesResponse>): QueryGetRevokedRootCertificatesResponse {
    const message = { ...baseQueryGetRevokedRootCertificatesResponse } as QueryGetRevokedRootCertificatesResponse
    if (object.revokedRootCertificates !== undefined && object.revokedRootCertificates !== null) {
      message.revokedRootCertificates = RevokedRootCertificates.fromPartial(object.revokedRootCertificates)
    } else {
      message.revokedRootCertificates = undefined
    }
    return message
  }
}

const baseQueryGetApprovedCertificatesBySubjectRequest: object = { subject: '' }

export const QueryGetApprovedCertificatesBySubjectRequest = {
  encode(message: QueryGetApprovedCertificatesBySubjectRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesBySubjectRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedCertificatesBySubjectRequest } as QueryGetApprovedCertificatesBySubjectRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetApprovedCertificatesBySubjectRequest {
    const message = { ...baseQueryGetApprovedCertificatesBySubjectRequest } as QueryGetApprovedCertificatesBySubjectRequest
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = String(object.subject)
    } else {
      message.subject = ''
    }
    return message
  },

  toJSON(message: QueryGetApprovedCertificatesBySubjectRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetApprovedCertificatesBySubjectRequest>): QueryGetApprovedCertificatesBySubjectRequest {
    const message = { ...baseQueryGetApprovedCertificatesBySubjectRequest } as QueryGetApprovedCertificatesBySubjectRequest
    if (object.subject !== undefined && object.subject !== null) {
      message.subject = object.subject
    } else {
      message.subject = ''
    }
    return message
  }
}

const baseQueryGetApprovedCertificatesBySubjectResponse: object = {}

export const QueryGetApprovedCertificatesBySubjectResponse = {
  encode(message: QueryGetApprovedCertificatesBySubjectResponse, writer: Writer = Writer.create()): Writer {
    if (message.approvedCertificatesBySubject !== undefined) {
      ApprovedCertificatesBySubject.encode(message.approvedCertificatesBySubject, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesBySubjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetApprovedCertificatesBySubjectResponse } as QueryGetApprovedCertificatesBySubjectResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificatesBySubject = ApprovedCertificatesBySubject.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetApprovedCertificatesBySubjectResponse {
    const message = { ...baseQueryGetApprovedCertificatesBySubjectResponse } as QueryGetApprovedCertificatesBySubjectResponse
    if (object.approvedCertificatesBySubject !== undefined && object.approvedCertificatesBySubject !== null) {
      message.approvedCertificatesBySubject = ApprovedCertificatesBySubject.fromJSON(object.approvedCertificatesBySubject)
    } else {
      message.approvedCertificatesBySubject = undefined
    }
    return message
  },

  toJSON(message: QueryGetApprovedCertificatesBySubjectResponse): unknown {
    const obj: any = {}
    message.approvedCertificatesBySubject !== undefined &&
      (obj.approvedCertificatesBySubject = message.approvedCertificatesBySubject
        ? ApprovedCertificatesBySubject.toJSON(message.approvedCertificatesBySubject)
        : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetApprovedCertificatesBySubjectResponse>): QueryGetApprovedCertificatesBySubjectResponse {
    const message = { ...baseQueryGetApprovedCertificatesBySubjectResponse } as QueryGetApprovedCertificatesBySubjectResponse
    if (object.approvedCertificatesBySubject !== undefined && object.approvedCertificatesBySubject !== null) {
      message.approvedCertificatesBySubject = ApprovedCertificatesBySubject.fromPartial(object.approvedCertificatesBySubject)
    } else {
      message.approvedCertificatesBySubject = undefined
    }
    return message
  }
}

const baseQueryGetRejectedCertificatesRequest: object = { subject: '', subjectKeyId: '' }

export const QueryGetRejectedCertificatesRequest = {
  encode(message: QueryGetRejectedCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRejectedCertificatesRequest } as QueryGetRejectedCertificatesRequest
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

  fromJSON(object: any): QueryGetRejectedCertificatesRequest {
    const message = { ...baseQueryGetRejectedCertificatesRequest } as QueryGetRejectedCertificatesRequest
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

  toJSON(message: QueryGetRejectedCertificatesRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRejectedCertificatesRequest>): QueryGetRejectedCertificatesRequest {
    const message = { ...baseQueryGetRejectedCertificatesRequest } as QueryGetRejectedCertificatesRequest
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

const baseQueryGetRejectedCertificatesResponse: object = {}

export const QueryGetRejectedCertificatesResponse = {
  encode(message: QueryGetRejectedCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.rejectedCertificate !== undefined) {
      RejectedCertificate.encode(message.rejectedCertificate, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRejectedCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRejectedCertificatesResponse } as QueryGetRejectedCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.rejectedCertificate = RejectedCertificate.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRejectedCertificatesResponse {
    const message = { ...baseQueryGetRejectedCertificatesResponse } as QueryGetRejectedCertificatesResponse
    if (object.rejectedCertificate !== undefined && object.rejectedCertificate !== null) {
      message.rejectedCertificate = RejectedCertificate.fromJSON(object.rejectedCertificate)
    } else {
      message.rejectedCertificate = undefined
    }
    return message
  },

  toJSON(message: QueryGetRejectedCertificatesResponse): unknown {
    const obj: any = {}
    message.rejectedCertificate !== undefined &&
      (obj.rejectedCertificate = message.rejectedCertificate ? RejectedCertificate.toJSON(message.rejectedCertificate) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRejectedCertificatesResponse>): QueryGetRejectedCertificatesResponse {
    const message = { ...baseQueryGetRejectedCertificatesResponse } as QueryGetRejectedCertificatesResponse
    if (object.rejectedCertificate !== undefined && object.rejectedCertificate !== null) {
      message.rejectedCertificate = RejectedCertificate.fromPartial(object.rejectedCertificate)
    } else {
      message.rejectedCertificate = undefined
    }
    return message
  }
}

const baseQueryAllRejectedCertificatesRequest: object = {}

export const QueryAllRejectedCertificatesRequest = {
  encode(message: QueryAllRejectedCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRejectedCertificatesRequest } as QueryAllRejectedCertificatesRequest
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

  fromJSON(object: any): QueryAllRejectedCertificatesRequest {
    const message = { ...baseQueryAllRejectedCertificatesRequest } as QueryAllRejectedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRejectedCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRejectedCertificatesRequest>): QueryAllRejectedCertificatesRequest {
    const message = { ...baseQueryAllRejectedCertificatesRequest } as QueryAllRejectedCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllRejectedCertificatesResponse: object = {}

export const QueryAllRejectedCertificatesResponse = {
  encode(message: QueryAllRejectedCertificatesResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.rejectedCertificate) {
      RejectedCertificate.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRejectedCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRejectedCertificatesResponse } as QueryAllRejectedCertificatesResponse
    message.rejectedCertificate = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.rejectedCertificate.push(RejectedCertificate.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllRejectedCertificatesResponse {
    const message = { ...baseQueryAllRejectedCertificatesResponse } as QueryAllRejectedCertificatesResponse
    message.rejectedCertificate = []
    if (object.rejectedCertificate !== undefined && object.rejectedCertificate !== null) {
      for (const e of object.rejectedCertificate) {
        message.rejectedCertificate.push(RejectedCertificate.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRejectedCertificatesResponse): unknown {
    const obj: any = {}
    if (message.rejectedCertificate) {
      obj.rejectedCertificate = message.rejectedCertificate.map((e) => (e ? RejectedCertificate.toJSON(e) : undefined))
    } else {
      obj.rejectedCertificate = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRejectedCertificatesResponse>): QueryAllRejectedCertificatesResponse {
    const message = { ...baseQueryAllRejectedCertificatesResponse } as QueryAllRejectedCertificatesResponse
    message.rejectedCertificate = []
    if (object.rejectedCertificate !== undefined && object.rejectedCertificate !== null) {
      for (const e of object.rejectedCertificate) {
        message.rejectedCertificate.push(RejectedCertificate.fromPartial(e))
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

const baseQueryGetPkiRevocationDistributionPointRequest: object = { vid: 0, label: '', issuerSubjectKeyID: '' }

export const QueryGetPkiRevocationDistributionPointRequest = {
  encode(message: QueryGetPkiRevocationDistributionPointRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.label !== '') {
      writer.uint32(18).string(message.label)
    }
    if (message.issuerSubjectKeyID !== '') {
      writer.uint32(26).string(message.issuerSubjectKeyID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPkiRevocationDistributionPointRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPkiRevocationDistributionPointRequest } as QueryGetPkiRevocationDistributionPointRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.label = reader.string()
          break
        case 3:
          message.issuerSubjectKeyID = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointRequest {
    const message = { ...baseQueryGetPkiRevocationDistributionPointRequest } as QueryGetPkiRevocationDistributionPointRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label)
    } else {
      message.label = ''
    }
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = String(object.issuerSubjectKeyID)
    } else {
      message.issuerSubjectKeyID = ''
    }
    return message
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.label !== undefined && (obj.label = message.label)
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPkiRevocationDistributionPointRequest>): QueryGetPkiRevocationDistributionPointRequest {
    const message = { ...baseQueryGetPkiRevocationDistributionPointRequest } as QueryGetPkiRevocationDistributionPointRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label
    } else {
      message.label = ''
    }
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = object.issuerSubjectKeyID
    } else {
      message.issuerSubjectKeyID = ''
    }
    return message
  }
}

const baseQueryGetPkiRevocationDistributionPointResponse: object = {}

export const QueryGetPkiRevocationDistributionPointResponse = {
  encode(message: QueryGetPkiRevocationDistributionPointResponse, writer: Writer = Writer.create()): Writer {
    if (message.PkiRevocationDistributionPoint !== undefined) {
      PkiRevocationDistributionPoint.encode(message.PkiRevocationDistributionPoint, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPkiRevocationDistributionPointResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetPkiRevocationDistributionPointResponse } as QueryGetPkiRevocationDistributionPointResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.PkiRevocationDistributionPoint = PkiRevocationDistributionPoint.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointResponse {
    const message = { ...baseQueryGetPkiRevocationDistributionPointResponse } as QueryGetPkiRevocationDistributionPointResponse
    if (object.PkiRevocationDistributionPoint !== undefined && object.PkiRevocationDistributionPoint !== null) {
      message.PkiRevocationDistributionPoint = PkiRevocationDistributionPoint.fromJSON(object.PkiRevocationDistributionPoint)
    } else {
      message.PkiRevocationDistributionPoint = undefined
    }
    return message
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointResponse): unknown {
    const obj: any = {}
    message.PkiRevocationDistributionPoint !== undefined &&
      (obj.PkiRevocationDistributionPoint = message.PkiRevocationDistributionPoint
        ? PkiRevocationDistributionPoint.toJSON(message.PkiRevocationDistributionPoint)
        : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetPkiRevocationDistributionPointResponse>): QueryGetPkiRevocationDistributionPointResponse {
    const message = { ...baseQueryGetPkiRevocationDistributionPointResponse } as QueryGetPkiRevocationDistributionPointResponse
    if (object.PkiRevocationDistributionPoint !== undefined && object.PkiRevocationDistributionPoint !== null) {
      message.PkiRevocationDistributionPoint = PkiRevocationDistributionPoint.fromPartial(object.PkiRevocationDistributionPoint)
    } else {
      message.PkiRevocationDistributionPoint = undefined
    }
    return message
  }
}

const baseQueryAllPkiRevocationDistributionPointRequest: object = {}

export const QueryAllPkiRevocationDistributionPointRequest = {
  encode(message: QueryAllPkiRevocationDistributionPointRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPkiRevocationDistributionPointRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPkiRevocationDistributionPointRequest } as QueryAllPkiRevocationDistributionPointRequest
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

  fromJSON(object: any): QueryAllPkiRevocationDistributionPointRequest {
    const message = { ...baseQueryAllPkiRevocationDistributionPointRequest } as QueryAllPkiRevocationDistributionPointRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPkiRevocationDistributionPointRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPkiRevocationDistributionPointRequest>): QueryAllPkiRevocationDistributionPointRequest {
    const message = { ...baseQueryAllPkiRevocationDistributionPointRequest } as QueryAllPkiRevocationDistributionPointRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllPkiRevocationDistributionPointResponse: object = {}

export const QueryAllPkiRevocationDistributionPointResponse = {
  encode(message: QueryAllPkiRevocationDistributionPointResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.PkiRevocationDistributionPoint) {
      PkiRevocationDistributionPoint.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPkiRevocationDistributionPointResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllPkiRevocationDistributionPointResponse } as QueryAllPkiRevocationDistributionPointResponse
    message.PkiRevocationDistributionPoint = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.PkiRevocationDistributionPoint.push(PkiRevocationDistributionPoint.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllPkiRevocationDistributionPointResponse {
    const message = { ...baseQueryAllPkiRevocationDistributionPointResponse } as QueryAllPkiRevocationDistributionPointResponse
    message.PkiRevocationDistributionPoint = []
    if (object.PkiRevocationDistributionPoint !== undefined && object.PkiRevocationDistributionPoint !== null) {
      for (const e of object.PkiRevocationDistributionPoint) {
        message.PkiRevocationDistributionPoint.push(PkiRevocationDistributionPoint.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllPkiRevocationDistributionPointResponse): unknown {
    const obj: any = {}
    if (message.PkiRevocationDistributionPoint) {
      obj.PkiRevocationDistributionPoint = message.PkiRevocationDistributionPoint.map((e) => (e ? PkiRevocationDistributionPoint.toJSON(e) : undefined))
    } else {
      obj.PkiRevocationDistributionPoint = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllPkiRevocationDistributionPointResponse>): QueryAllPkiRevocationDistributionPointResponse {
    const message = { ...baseQueryAllPkiRevocationDistributionPointResponse } as QueryAllPkiRevocationDistributionPointResponse
    message.PkiRevocationDistributionPoint = []
    if (object.PkiRevocationDistributionPoint !== undefined && object.PkiRevocationDistributionPoint !== null) {
      for (const e of object.PkiRevocationDistributionPoint) {
        message.PkiRevocationDistributionPoint.push(PkiRevocationDistributionPoint.fromPartial(e))
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

const baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest: object = { issuerSubjectKeyID: '' }

export const QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest = {
  encode(message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest, writer: Writer = Writer.create()): Writer {
    if (message.issuerSubjectKeyID !== '') {
      writer.uint32(10).string(message.issuerSubjectKeyID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
    } as QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
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

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
    const message = {
      ...baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
    } as QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = String(object.issuerSubjectKeyID)
    } else {
      message.issuerSubjectKeyID = ''
    }
    return message
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest): unknown {
    const obj: any = {}
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID)
    return obj
  },

  fromPartial(
    object: DeepPartial<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest>
  ): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
    const message = {
      ...baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
    } as QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = object.issuerSubjectKeyID
    } else {
      message.issuerSubjectKeyID = ''
    }
    return message
  }
}

const baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse: object = {}

export const QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse = {
  encode(message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse, writer: Writer = Writer.create()): Writer {
    if (message.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== undefined) {
      PkiRevocationDistributionPointsByIssuerSubjectKeyID.encode(message.pkiRevocationDistributionPointsByIssuerSubjectKeyID, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = {
      ...baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
    } as QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pkiRevocationDistributionPointsByIssuerSubjectKeyID = PkiRevocationDistributionPointsByIssuerSubjectKeyID.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
    const message = {
      ...baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
    } as QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
    if (object.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== undefined && object.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== null) {
      message.pkiRevocationDistributionPointsByIssuerSubjectKeyID = PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromJSON(
        object.pkiRevocationDistributionPointsByIssuerSubjectKeyID
      )
    } else {
      message.pkiRevocationDistributionPointsByIssuerSubjectKeyID = undefined
    }
    return message
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse): unknown {
    const obj: any = {}
    message.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== undefined &&
      (obj.pkiRevocationDistributionPointsByIssuerSubjectKeyID = message.pkiRevocationDistributionPointsByIssuerSubjectKeyID
        ? PkiRevocationDistributionPointsByIssuerSubjectKeyID.toJSON(message.pkiRevocationDistributionPointsByIssuerSubjectKeyID)
        : undefined)
    return obj
  },

  fromPartial(
    object: DeepPartial<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse>
  ): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
    const message = {
      ...baseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
    } as QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse
    if (object.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== undefined && object.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== null) {
      message.pkiRevocationDistributionPointsByIssuerSubjectKeyID = PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromPartial(
        object.pkiRevocationDistributionPointsByIssuerSubjectKeyID
      )
    } else {
      message.pkiRevocationDistributionPointsByIssuerSubjectKeyID = undefined
    }
    return message
  }
}

const baseQueryGetNocRootCertificatesRequest: object = { vid: 0 }

export const QueryGetNocRootCertificatesRequest = {
  encode(message: QueryGetNocRootCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNocRootCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNocRootCertificatesRequest } as QueryGetNocRootCertificatesRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNocRootCertificatesRequest {
    const message = { ...baseQueryGetNocRootCertificatesRequest } as QueryGetNocRootCertificatesRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    return message
  },

  toJSON(message: QueryGetNocRootCertificatesRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNocRootCertificatesRequest>): QueryGetNocRootCertificatesRequest {
    const message = { ...baseQueryGetNocRootCertificatesRequest } as QueryGetNocRootCertificatesRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    return message
  }
}

const baseQueryGetNocRootCertificatesResponse: object = {}

export const QueryGetNocRootCertificatesResponse = {
  encode(message: QueryGetNocRootCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.nocRootCertificates !== undefined) {
      NocRootCertificates.encode(message.nocRootCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNocRootCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNocRootCertificatesResponse } as QueryGetNocRootCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.nocRootCertificates = NocRootCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNocRootCertificatesResponse {
    const message = { ...baseQueryGetNocRootCertificatesResponse } as QueryGetNocRootCertificatesResponse
    if (object.nocRootCertificates !== undefined && object.nocRootCertificates !== null) {
      message.nocRootCertificates = NocRootCertificates.fromJSON(object.nocRootCertificates)
    } else {
      message.nocRootCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetNocRootCertificatesResponse): unknown {
    const obj: any = {}
    message.nocRootCertificates !== undefined &&
      (obj.nocRootCertificates = message.nocRootCertificates ? NocRootCertificates.toJSON(message.nocRootCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNocRootCertificatesResponse>): QueryGetNocRootCertificatesResponse {
    const message = { ...baseQueryGetNocRootCertificatesResponse } as QueryGetNocRootCertificatesResponse
    if (object.nocRootCertificates !== undefined && object.nocRootCertificates !== null) {
      message.nocRootCertificates = NocRootCertificates.fromPartial(object.nocRootCertificates)
    } else {
      message.nocRootCertificates = undefined
    }
    return message
  }
}

const baseQueryAllNocRootCertificatesRequest: object = {}

export const QueryAllNocRootCertificatesRequest = {
  encode(message: QueryAllNocRootCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNocRootCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNocRootCertificatesRequest } as QueryAllNocRootCertificatesRequest
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

  fromJSON(object: any): QueryAllNocRootCertificatesRequest {
    const message = { ...baseQueryAllNocRootCertificatesRequest } as QueryAllNocRootCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNocRootCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNocRootCertificatesRequest>): QueryAllNocRootCertificatesRequest {
    const message = { ...baseQueryAllNocRootCertificatesRequest } as QueryAllNocRootCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllNocRootCertificatesResponse: object = {}

export const QueryAllNocRootCertificatesResponse = {
  encode(message: QueryAllNocRootCertificatesResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.nocRootCertificates) {
      NocRootCertificates.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNocRootCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNocRootCertificatesResponse } as QueryAllNocRootCertificatesResponse
    message.nocRootCertificates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.nocRootCertificates.push(NocRootCertificates.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllNocRootCertificatesResponse {
    const message = { ...baseQueryAllNocRootCertificatesResponse } as QueryAllNocRootCertificatesResponse
    message.nocRootCertificates = []
    if (object.nocRootCertificates !== undefined && object.nocRootCertificates !== null) {
      for (const e of object.nocRootCertificates) {
        message.nocRootCertificates.push(NocRootCertificates.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNocRootCertificatesResponse): unknown {
    const obj: any = {}
    if (message.nocRootCertificates) {
      obj.nocRootCertificates = message.nocRootCertificates.map((e) => (e ? NocRootCertificates.toJSON(e) : undefined))
    } else {
      obj.nocRootCertificates = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNocRootCertificatesResponse>): QueryAllNocRootCertificatesResponse {
    const message = { ...baseQueryAllNocRootCertificatesResponse } as QueryAllNocRootCertificatesResponse
    message.nocRootCertificates = []
    if (object.nocRootCertificates !== undefined && object.nocRootCertificates !== null) {
      for (const e of object.nocRootCertificates) {
        message.nocRootCertificates.push(NocRootCertificates.fromPartial(e))
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

const baseQueryGetNocIcaCertificatesRequest: object = { vid: 0 }

export const QueryGetNocIcaCertificatesRequest = {
  encode(message: QueryGetNocIcaCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNocIcaCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNocIcaCertificatesRequest } as QueryGetNocIcaCertificatesRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNocIcaCertificatesRequest {
    const message = { ...baseQueryGetNocIcaCertificatesRequest } as QueryGetNocIcaCertificatesRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    return message
  },

  toJSON(message: QueryGetNocIcaCertificatesRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNocIcaCertificatesRequest>): QueryGetNocIcaCertificatesRequest {
    const message = { ...baseQueryGetNocIcaCertificatesRequest } as QueryGetNocIcaCertificatesRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    return message
  }
}

const baseQueryGetNocIcaCertificatesResponse: object = {}

export const QueryGetNocIcaCertificatesResponse = {
  encode(message: QueryGetNocIcaCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.nocIcaCertificates !== undefined) {
      NocIcaCertificates.encode(message.nocIcaCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNocIcaCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNocIcaCertificatesResponse } as QueryGetNocIcaCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.nocIcaCertificates = NocIcaCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNocIcaCertificatesResponse {
    const message = { ...baseQueryGetNocIcaCertificatesResponse } as QueryGetNocIcaCertificatesResponse
    if (object.nocIcaCertificates !== undefined && object.nocIcaCertificates !== null) {
      message.nocIcaCertificates = NocIcaCertificates.fromJSON(object.nocIcaCertificates)
    } else {
      message.nocIcaCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetNocIcaCertificatesResponse): unknown {
    const obj: any = {}
    message.nocIcaCertificates !== undefined &&
      (obj.nocIcaCertificates = message.nocIcaCertificates ? NocIcaCertificates.toJSON(message.nocIcaCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNocIcaCertificatesResponse>): QueryGetNocIcaCertificatesResponse {
    const message = { ...baseQueryGetNocIcaCertificatesResponse } as QueryGetNocIcaCertificatesResponse
    if (object.nocIcaCertificates !== undefined && object.nocIcaCertificates !== null) {
      message.nocIcaCertificates = NocIcaCertificates.fromPartial(object.nocIcaCertificates)
    } else {
      message.nocIcaCertificates = undefined
    }
    return message
  }
}

const baseQueryAllNocIcaCertificatesRequest: object = {}

export const QueryAllNocIcaCertificatesRequest = {
  encode(message: QueryAllNocIcaCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNocIcaCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNocIcaCertificatesRequest } as QueryAllNocIcaCertificatesRequest
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

  fromJSON(object: any): QueryAllNocIcaCertificatesRequest {
    const message = { ...baseQueryAllNocIcaCertificatesRequest } as QueryAllNocIcaCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNocIcaCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNocIcaCertificatesRequest>): QueryAllNocIcaCertificatesRequest {
    const message = { ...baseQueryAllNocIcaCertificatesRequest } as QueryAllNocIcaCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllNocIcaCertificatesResponse: object = {}

export const QueryAllNocIcaCertificatesResponse = {
  encode(message: QueryAllNocIcaCertificatesResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.nocIcaCertificates) {
      NocIcaCertificates.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNocIcaCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNocIcaCertificatesResponse } as QueryAllNocIcaCertificatesResponse
    message.nocIcaCertificates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.nocIcaCertificates.push(NocIcaCertificates.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllNocIcaCertificatesResponse {
    const message = { ...baseQueryAllNocIcaCertificatesResponse } as QueryAllNocIcaCertificatesResponse
    message.nocIcaCertificates = []
    if (object.nocIcaCertificates !== undefined && object.nocIcaCertificates !== null) {
      for (const e of object.nocIcaCertificates) {
        message.nocIcaCertificates.push(NocIcaCertificates.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNocIcaCertificatesResponse): unknown {
    const obj: any = {}
    if (message.nocIcaCertificates) {
      obj.nocIcaCertificates = message.nocIcaCertificates.map((e) => (e ? NocIcaCertificates.toJSON(e) : undefined))
    } else {
      obj.nocIcaCertificates = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNocIcaCertificatesResponse>): QueryAllNocIcaCertificatesResponse {
    const message = { ...baseQueryAllNocIcaCertificatesResponse } as QueryAllNocIcaCertificatesResponse
    message.nocIcaCertificates = []
    if (object.nocIcaCertificates !== undefined && object.nocIcaCertificates !== null) {
      for (const e of object.nocIcaCertificates) {
        message.nocIcaCertificates.push(NocIcaCertificates.fromPartial(e))
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

const baseQueryGetRevokedNocRootCertificatesRequest: object = { subject: '', subjectKeyId: '' }

export const QueryGetRevokedNocRootCertificatesRequest = {
  encode(message: QueryGetRevokedNocRootCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.subject !== '') {
      writer.uint32(10).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(18).string(message.subjectKeyId)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedNocRootCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedNocRootCertificatesRequest } as QueryGetRevokedNocRootCertificatesRequest
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

  fromJSON(object: any): QueryGetRevokedNocRootCertificatesRequest {
    const message = { ...baseQueryGetRevokedNocRootCertificatesRequest } as QueryGetRevokedNocRootCertificatesRequest
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

  toJSON(message: QueryGetRevokedNocRootCertificatesRequest): unknown {
    const obj: any = {}
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRevokedNocRootCertificatesRequest>): QueryGetRevokedNocRootCertificatesRequest {
    const message = { ...baseQueryGetRevokedNocRootCertificatesRequest } as QueryGetRevokedNocRootCertificatesRequest
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

const baseQueryGetRevokedNocRootCertificatesResponse: object = {}

export const QueryGetRevokedNocRootCertificatesResponse = {
  encode(message: QueryGetRevokedNocRootCertificatesResponse, writer: Writer = Writer.create()): Writer {
    if (message.revokedNocRootCertificates !== undefined) {
      RevokedNocRootCertificates.encode(message.revokedNocRootCertificates, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedNocRootCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetRevokedNocRootCertificatesResponse } as QueryGetRevokedNocRootCertificatesResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.revokedNocRootCertificates = RevokedNocRootCertificates.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetRevokedNocRootCertificatesResponse {
    const message = { ...baseQueryGetRevokedNocRootCertificatesResponse } as QueryGetRevokedNocRootCertificatesResponse
    if (object.revokedNocRootCertificates !== undefined && object.revokedNocRootCertificates !== null) {
      message.revokedNocRootCertificates = RevokedNocRootCertificates.fromJSON(object.revokedNocRootCertificates)
    } else {
      message.revokedNocRootCertificates = undefined
    }
    return message
  },

  toJSON(message: QueryGetRevokedNocRootCertificatesResponse): unknown {
    const obj: any = {}
    message.revokedNocRootCertificates !== undefined &&
      (obj.revokedNocRootCertificates = message.revokedNocRootCertificates ? RevokedNocRootCertificates.toJSON(message.revokedNocRootCertificates) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetRevokedNocRootCertificatesResponse>): QueryGetRevokedNocRootCertificatesResponse {
    const message = { ...baseQueryGetRevokedNocRootCertificatesResponse } as QueryGetRevokedNocRootCertificatesResponse
    if (object.revokedNocRootCertificates !== undefined && object.revokedNocRootCertificates !== null) {
      message.revokedNocRootCertificates = RevokedNocRootCertificates.fromPartial(object.revokedNocRootCertificates)
    } else {
      message.revokedNocRootCertificates = undefined
    }
    return message
  }
}

const baseQueryAllRevokedNocRootCertificatesRequest: object = {}

export const QueryAllRevokedNocRootCertificatesRequest = {
  encode(message: QueryAllRevokedNocRootCertificatesRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedNocRootCertificatesRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRevokedNocRootCertificatesRequest } as QueryAllRevokedNocRootCertificatesRequest
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

  fromJSON(object: any): QueryAllRevokedNocRootCertificatesRequest {
    const message = { ...baseQueryAllRevokedNocRootCertificatesRequest } as QueryAllRevokedNocRootCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRevokedNocRootCertificatesRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRevokedNocRootCertificatesRequest>): QueryAllRevokedNocRootCertificatesRequest {
    const message = { ...baseQueryAllRevokedNocRootCertificatesRequest } as QueryAllRevokedNocRootCertificatesRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllRevokedNocRootCertificatesResponse: object = {}

export const QueryAllRevokedNocRootCertificatesResponse = {
  encode(message: QueryAllRevokedNocRootCertificatesResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.revokedNocRootCertificates) {
      RevokedNocRootCertificates.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedNocRootCertificatesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllRevokedNocRootCertificatesResponse } as QueryAllRevokedNocRootCertificatesResponse
    message.revokedNocRootCertificates = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.revokedNocRootCertificates.push(RevokedNocRootCertificates.decode(reader, reader.uint32()))
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

  fromJSON(object: any): QueryAllRevokedNocRootCertificatesResponse {
    const message = { ...baseQueryAllRevokedNocRootCertificatesResponse } as QueryAllRevokedNocRootCertificatesResponse
    message.revokedNocRootCertificates = []
    if (object.revokedNocRootCertificates !== undefined && object.revokedNocRootCertificates !== null) {
      for (const e of object.revokedNocRootCertificates) {
        message.revokedNocRootCertificates.push(RevokedNocRootCertificates.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllRevokedNocRootCertificatesResponse): unknown {
    const obj: any = {}
    if (message.revokedNocRootCertificates) {
      obj.revokedNocRootCertificates = message.revokedNocRootCertificates.map((e) => (e ? RevokedNocRootCertificates.toJSON(e) : undefined))
    } else {
      obj.revokedNocRootCertificates = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllRevokedNocRootCertificatesResponse>): QueryAllRevokedNocRootCertificatesResponse {
    const message = { ...baseQueryAllRevokedNocRootCertificatesResponse } as QueryAllRevokedNocRootCertificatesResponse
    message.revokedNocRootCertificates = []
    if (object.revokedNocRootCertificates !== undefined && object.revokedNocRootCertificates !== null) {
      for (const e of object.revokedNocRootCertificates) {
        message.revokedNocRootCertificates.push(RevokedNocRootCertificates.fromPartial(e))
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
  /** Queries a ProposedCertificateRevocation by index. */
  ProposedCertificateRevocation(request: QueryGetProposedCertificateRevocationRequest): Promise<QueryGetProposedCertificateRevocationResponse>
  /** Queries a list of ProposedCertificateRevocation items. */
  ProposedCertificateRevocationAll(request: QueryAllProposedCertificateRevocationRequest): Promise<QueryAllProposedCertificateRevocationResponse>
  /** Queries a RevokedCertificates by index. */
  RevokedCertificates(request: QueryGetRevokedCertificatesRequest): Promise<QueryGetRevokedCertificatesResponse>
  /** Queries a list of RevokedCertificates items. */
  RevokedCertificatesAll(request: QueryAllRevokedCertificatesRequest): Promise<QueryAllRevokedCertificatesResponse>
  /** Queries a ApprovedRootCertificates by index. */
  ApprovedRootCertificates(request: QueryGetApprovedRootCertificatesRequest): Promise<QueryGetApprovedRootCertificatesResponse>
  /** Queries a RevokedRootCertificates by index. */
  RevokedRootCertificates(request: QueryGetRevokedRootCertificatesRequest): Promise<QueryGetRevokedRootCertificatesResponse>
  /** Queries a ApprovedCertificatesBySubject by index. */
  ApprovedCertificatesBySubject(request: QueryGetApprovedCertificatesBySubjectRequest): Promise<QueryGetApprovedCertificatesBySubjectResponse>
  /** Queries a RejectedCertificate by index. */
  RejectedCertificate(request: QueryGetRejectedCertificatesRequest): Promise<QueryGetRejectedCertificatesResponse>
  /** Queries a list of RejectedCertificate items. */
  RejectedCertificateAll(request: QueryAllRejectedCertificatesRequest): Promise<QueryAllRejectedCertificatesResponse>
  /** Queries a PkiRevocationDistributionPoint by index. */
  PkiRevocationDistributionPoint(request: QueryGetPkiRevocationDistributionPointRequest): Promise<QueryGetPkiRevocationDistributionPointResponse>
  /** Queries a list of PkiRevocationDistributionPoint items. */
  PkiRevocationDistributionPointAll(request: QueryAllPkiRevocationDistributionPointRequest): Promise<QueryAllPkiRevocationDistributionPointResponse>
  /** Queries a PkiRevocationDistributionPointsByIssuerSubjectKeyID by index. */
  PkiRevocationDistributionPointsByIssuerSubjectKeyID(
    request: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
  ): Promise<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse>
  /** Queries a NocRootCertificates by index. */
  NocRootCertificates(request: QueryGetNocRootCertificatesRequest): Promise<QueryGetNocRootCertificatesResponse>
  /** Queries a list of NocRootCertificates items. */
  NocRootCertificatesAll(request: QueryAllNocRootCertificatesRequest): Promise<QueryAllNocRootCertificatesResponse>
  /** Queries a NocIcaCertificates by index. */
  NocIcaCertificates(request: QueryGetNocIcaCertificatesRequest): Promise<QueryGetNocIcaCertificatesResponse>
  /** Queries a list of NocIcaCertificates items. */
  NocIcaCertificatesAll(request: QueryAllNocIcaCertificatesRequest): Promise<QueryAllNocIcaCertificatesResponse>
  /** Queries a RevokedNocRootCertificates by index. */
  RevokedNocRootCertificates(request: QueryGetRevokedNocRootCertificatesRequest): Promise<QueryGetRevokedNocRootCertificatesResponse>
  /** Queries a list of RevokedNocRootCertificates items. */
  RevokedNocRootCertificatesAll(request: QueryAllRevokedNocRootCertificatesRequest): Promise<QueryAllRevokedNocRootCertificatesResponse>
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

  ApprovedRootCertificates(request: QueryGetApprovedRootCertificatesRequest): Promise<QueryGetApprovedRootCertificatesResponse> {
    const data = QueryGetApprovedRootCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ApprovedRootCertificates', data)
    return promise.then((data) => QueryGetApprovedRootCertificatesResponse.decode(new Reader(data)))
  }

  RevokedRootCertificates(request: QueryGetRevokedRootCertificatesRequest): Promise<QueryGetRevokedRootCertificatesResponse> {
    const data = QueryGetRevokedRootCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RevokedRootCertificates', data)
    return promise.then((data) => QueryGetRevokedRootCertificatesResponse.decode(new Reader(data)))
  }

  ApprovedCertificatesBySubject(request: QueryGetApprovedCertificatesBySubjectRequest): Promise<QueryGetApprovedCertificatesBySubjectResponse> {
    const data = QueryGetApprovedCertificatesBySubjectRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ApprovedCertificatesBySubject', data)
    return promise.then((data) => QueryGetApprovedCertificatesBySubjectResponse.decode(new Reader(data)))
  }

  RejectedCertificate(request: QueryGetRejectedCertificatesRequest): Promise<QueryGetRejectedCertificatesResponse> {
    const data = QueryGetRejectedCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RejectedCertificate', data)
    return promise.then((data) => QueryGetRejectedCertificatesResponse.decode(new Reader(data)))
  }

  RejectedCertificateAll(request: QueryAllRejectedCertificatesRequest): Promise<QueryAllRejectedCertificatesResponse> {
    const data = QueryAllRejectedCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RejectedCertificateAll', data)
    return promise.then((data) => QueryAllRejectedCertificatesResponse.decode(new Reader(data)))
  }

  PkiRevocationDistributionPoint(request: QueryGetPkiRevocationDistributionPointRequest): Promise<QueryGetPkiRevocationDistributionPointResponse> {
    const data = QueryGetPkiRevocationDistributionPointRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'PkiRevocationDistributionPoint', data)
    return promise.then((data) => QueryGetPkiRevocationDistributionPointResponse.decode(new Reader(data)))
  }

  PkiRevocationDistributionPointAll(request: QueryAllPkiRevocationDistributionPointRequest): Promise<QueryAllPkiRevocationDistributionPointResponse> {
    const data = QueryAllPkiRevocationDistributionPointRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'PkiRevocationDistributionPointAll', data)
    return promise.then((data) => QueryAllPkiRevocationDistributionPointResponse.decode(new Reader(data)))
  }

  PkiRevocationDistributionPointsByIssuerSubjectKeyID(
    request: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest
  ): Promise<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse> {
    const data = QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'PkiRevocationDistributionPointsByIssuerSubjectKeyID', data)
    return promise.then((data) => QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse.decode(new Reader(data)))
  }

  NocRootCertificates(request: QueryGetNocRootCertificatesRequest): Promise<QueryGetNocRootCertificatesResponse> {
    const data = QueryGetNocRootCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'NocRootCertificates', data)
    return promise.then((data) => QueryGetNocRootCertificatesResponse.decode(new Reader(data)))
  }

  NocRootCertificatesAll(request: QueryAllNocRootCertificatesRequest): Promise<QueryAllNocRootCertificatesResponse> {
    const data = QueryAllNocRootCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'NocRootCertificatesAll', data)
    return promise.then((data) => QueryAllNocRootCertificatesResponse.decode(new Reader(data)))
  }

  NocIcaCertificates(request: QueryGetNocIcaCertificatesRequest): Promise<QueryGetNocIcaCertificatesResponse> {
    const data = QueryGetNocIcaCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'NocIcaCertificates', data)
    return promise.then((data) => QueryGetNocIcaCertificatesResponse.decode(new Reader(data)))
  }

  NocIcaCertificatesAll(request: QueryAllNocIcaCertificatesRequest): Promise<QueryAllNocIcaCertificatesResponse> {
    const data = QueryAllNocIcaCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'NocIcaCertificatesAll', data)
    return promise.then((data) => QueryAllNocIcaCertificatesResponse.decode(new Reader(data)))
  }

  RevokedNocRootCertificates(request: QueryGetRevokedNocRootCertificatesRequest): Promise<QueryGetRevokedNocRootCertificatesResponse> {
    const data = QueryGetRevokedNocRootCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RevokedNocRootCertificates', data)
    return promise.then((data) => QueryGetRevokedNocRootCertificatesResponse.decode(new Reader(data)))
  }

  RevokedNocRootCertificatesAll(request: QueryAllRevokedNocRootCertificatesRequest): Promise<QueryAllRevokedNocRootCertificatesResponse> {
    const data = QueryAllRevokedNocRootCertificatesRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RevokedNocRootCertificatesAll', data)
    return promise.then((data) => QueryAllRevokedNocRootCertificatesResponse.decode(new Reader(data)))
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
