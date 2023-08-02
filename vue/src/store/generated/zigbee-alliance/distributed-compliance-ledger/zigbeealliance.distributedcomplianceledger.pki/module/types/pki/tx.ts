/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal'
import * as Long from 'long'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

export interface MsgProposeAddX509RootCert {
  signer: string
  cert: string
  info: string
  time: number
  vid: number
}

export interface MsgProposeAddX509RootCertResponse {}

export interface MsgApproveAddX509RootCert {
  signer: string
  subject: string
  subjectKeyId: string
  info: string
  time: number
}

export interface MsgApproveAddX509RootCertResponse {}

export interface MsgAddX509Cert {
  signer: string
  cert: string
  info: string
  time: number
}

export interface MsgAddX509CertResponse {}

export interface MsgProposeRevokeX509RootCert {
  signer: string
  subject: string
  subjectKeyId: string
  info: string
  time: number
}

export interface MsgProposeRevokeX509RootCertResponse {}

export interface MsgApproveRevokeX509RootCert {
  signer: string
  subject: string
  subjectKeyId: string
  info: string
  time: number
}

export interface MsgApproveRevokeX509RootCertResponse {}

export interface MsgRevokeX509Cert {
  signer: string
  subject: string
  subjectKeyId: string
  info: string
  time: number
}

export interface MsgRevokeX509CertResponse {}

export interface MsgRejectAddX509RootCert {
  signer: string
  subject: string
  subjectKeyId: string
  info: string
  time: number
}

export interface MsgRejectAddX509RootCertResponse {}

export interface MsgAddPkiRevocationDistributionPoint {
  signer: string
  vid: number
  pid: number
  isPAA: boolean
  label: string
  crlSignerCertificate: string
  issuerSubjectKeyID: string
  dataURL: string
  dataFileSize: number
  dataDigest: string
  dataDigestType: number
  revocationType: number
}

export interface MsgAddPkiRevocationDistributionPointResponse {}

export interface MsgUpdatePkiRevocationDistributionPoint {
  signer: string
  vid: number
  label: string
  crlSignerCertificate: string
  issuerSubjectKeyID: string
  dataURL: string
  dataFileSize: number
  dataDigest: string
  dataDigestType: number
}

export interface MsgUpdatePkiRevocationDistributionPointResponse {}

export interface MsgDeletePkiRevocationDistributionPoint {
  signer: string
  vid: number
  label: string
  issuerSubjectKeyID: string
}

export interface MsgDeletePkiRevocationDistributionPointResponse {}

export interface MsgAssignVid {
  signer: string
  subject: string
  subjectKeyId: string
  vid: number
}

export interface MsgAssignVidResponse {}

const baseMsgProposeAddX509RootCert: object = { signer: '', cert: '', info: '', time: 0, vid: 0 }

export const MsgProposeAddX509RootCert = {
  encode(message: MsgProposeAddX509RootCert, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.cert !== '') {
      writer.uint32(18).string(message.cert)
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    if (message.vid !== 0) {
      writer.uint32(40).int32(message.vid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeAddX509RootCert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeAddX509RootCert } as MsgProposeAddX509RootCert
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.cert = reader.string()
          break
        case 3:
          message.info = reader.string()
          break
        case 4:
          message.time = longToNumber(reader.int64() as Long)
          break
        case 5:
          message.vid = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgProposeAddX509RootCert {
    const message = { ...baseMsgProposeAddX509RootCert } as MsgProposeAddX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.cert !== undefined && object.cert !== null) {
      message.cert = String(object.cert)
    } else {
      message.cert = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    return message
  },

  toJSON(message: MsgProposeAddX509RootCert): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.cert !== undefined && (obj.cert = message.cert)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    message.vid !== undefined && (obj.vid = message.vid)
    return obj
  },

  fromPartial(object: DeepPartial<MsgProposeAddX509RootCert>): MsgProposeAddX509RootCert {
    const message = { ...baseMsgProposeAddX509RootCert } as MsgProposeAddX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.cert !== undefined && object.cert !== null) {
      message.cert = object.cert
    } else {
      message.cert = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    return message
  }
}

const baseMsgProposeAddX509RootCertResponse: object = {}

export const MsgProposeAddX509RootCertResponse = {
  encode(_: MsgProposeAddX509RootCertResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeAddX509RootCertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeAddX509RootCertResponse } as MsgProposeAddX509RootCertResponse
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

  fromJSON(_: any): MsgProposeAddX509RootCertResponse {
    const message = { ...baseMsgProposeAddX509RootCertResponse } as MsgProposeAddX509RootCertResponse
    return message
  },

  toJSON(_: MsgProposeAddX509RootCertResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgProposeAddX509RootCertResponse>): MsgProposeAddX509RootCertResponse {
    const message = { ...baseMsgProposeAddX509RootCertResponse } as MsgProposeAddX509RootCertResponse
    return message
  }
}

const baseMsgApproveAddX509RootCert: object = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 }

export const MsgApproveAddX509RootCert = {
  encode(message: MsgApproveAddX509RootCert, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.subject !== '') {
      writer.uint32(18).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(26).string(message.subjectKeyId)
    }
    if (message.info !== '') {
      writer.uint32(34).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(40).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveAddX509RootCert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveAddX509RootCert } as MsgApproveAddX509RootCert
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.subject = reader.string()
          break
        case 3:
          message.subjectKeyId = reader.string()
          break
        case 4:
          message.info = reader.string()
          break
        case 5:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgApproveAddX509RootCert {
    const message = { ...baseMsgApproveAddX509RootCert } as MsgApproveAddX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgApproveAddX509RootCert): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgApproveAddX509RootCert>): MsgApproveAddX509RootCert {
    const message = { ...baseMsgApproveAddX509RootCert } as MsgApproveAddX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgApproveAddX509RootCertResponse: object = {}

export const MsgApproveAddX509RootCertResponse = {
  encode(_: MsgApproveAddX509RootCertResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveAddX509RootCertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveAddX509RootCertResponse } as MsgApproveAddX509RootCertResponse
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

  fromJSON(_: any): MsgApproveAddX509RootCertResponse {
    const message = { ...baseMsgApproveAddX509RootCertResponse } as MsgApproveAddX509RootCertResponse
    return message
  },

  toJSON(_: MsgApproveAddX509RootCertResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgApproveAddX509RootCertResponse>): MsgApproveAddX509RootCertResponse {
    const message = { ...baseMsgApproveAddX509RootCertResponse } as MsgApproveAddX509RootCertResponse
    return message
  }
}

const baseMsgAddX509Cert: object = { signer: '', cert: '', info: '', time: 0 }

export const MsgAddX509Cert = {
  encode(message: MsgAddX509Cert, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.cert !== '') {
      writer.uint32(18).string(message.cert)
    }
    if (message.info !== '') {
      writer.uint32(26).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(32).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddX509Cert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddX509Cert } as MsgAddX509Cert
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.cert = reader.string()
          break
        case 3:
          message.info = reader.string()
          break
        case 4:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgAddX509Cert {
    const message = { ...baseMsgAddX509Cert } as MsgAddX509Cert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
    if (object.cert !== undefined && object.cert !== null) {
      message.cert = String(object.cert)
    } else {
      message.cert = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgAddX509Cert): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.cert !== undefined && (obj.cert = message.cert)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgAddX509Cert>): MsgAddX509Cert {
    const message = { ...baseMsgAddX509Cert } as MsgAddX509Cert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
    if (object.cert !== undefined && object.cert !== null) {
      message.cert = object.cert
    } else {
      message.cert = ''
    }
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgAddX509CertResponse: object = {}

export const MsgAddX509CertResponse = {
  encode(_: MsgAddX509CertResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddX509CertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddX509CertResponse } as MsgAddX509CertResponse
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

  fromJSON(_: any): MsgAddX509CertResponse {
    const message = { ...baseMsgAddX509CertResponse } as MsgAddX509CertResponse
    return message
  },

  toJSON(_: MsgAddX509CertResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgAddX509CertResponse>): MsgAddX509CertResponse {
    const message = { ...baseMsgAddX509CertResponse } as MsgAddX509CertResponse
    return message
  }
}

const baseMsgProposeRevokeX509RootCert: object = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 }

export const MsgProposeRevokeX509RootCert = {
  encode(message: MsgProposeRevokeX509RootCert, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.subject !== '') {
      writer.uint32(18).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(26).string(message.subjectKeyId)
    }
    if (message.info !== '') {
      writer.uint32(34).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(40).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeX509RootCert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeRevokeX509RootCert } as MsgProposeRevokeX509RootCert
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.subject = reader.string()
          break
        case 3:
          message.subjectKeyId = reader.string()
          break
        case 4:
          message.info = reader.string()
          break
        case 5:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgProposeRevokeX509RootCert {
    const message = { ...baseMsgProposeRevokeX509RootCert } as MsgProposeRevokeX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgProposeRevokeX509RootCert): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgProposeRevokeX509RootCert>): MsgProposeRevokeX509RootCert {
    const message = { ...baseMsgProposeRevokeX509RootCert } as MsgProposeRevokeX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgProposeRevokeX509RootCertResponse: object = {}

export const MsgProposeRevokeX509RootCertResponse = {
  encode(_: MsgProposeRevokeX509RootCertResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeX509RootCertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgProposeRevokeX509RootCertResponse } as MsgProposeRevokeX509RootCertResponse
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

  fromJSON(_: any): MsgProposeRevokeX509RootCertResponse {
    const message = { ...baseMsgProposeRevokeX509RootCertResponse } as MsgProposeRevokeX509RootCertResponse
    return message
  },

  toJSON(_: MsgProposeRevokeX509RootCertResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgProposeRevokeX509RootCertResponse>): MsgProposeRevokeX509RootCertResponse {
    const message = { ...baseMsgProposeRevokeX509RootCertResponse } as MsgProposeRevokeX509RootCertResponse
    return message
  }
}

const baseMsgApproveRevokeX509RootCert: object = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 }

export const MsgApproveRevokeX509RootCert = {
  encode(message: MsgApproveRevokeX509RootCert, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.subject !== '') {
      writer.uint32(18).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(26).string(message.subjectKeyId)
    }
    if (message.info !== '') {
      writer.uint32(42).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(48).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeX509RootCert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveRevokeX509RootCert } as MsgApproveRevokeX509RootCert
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.subject = reader.string()
          break
        case 3:
          message.subjectKeyId = reader.string()
          break
        case 5:
          message.info = reader.string()
          break
        case 6:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgApproveRevokeX509RootCert {
    const message = { ...baseMsgApproveRevokeX509RootCert } as MsgApproveRevokeX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgApproveRevokeX509RootCert): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgApproveRevokeX509RootCert>): MsgApproveRevokeX509RootCert {
    const message = { ...baseMsgApproveRevokeX509RootCert } as MsgApproveRevokeX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgApproveRevokeX509RootCertResponse: object = {}

export const MsgApproveRevokeX509RootCertResponse = {
  encode(_: MsgApproveRevokeX509RootCertResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeX509RootCertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgApproveRevokeX509RootCertResponse } as MsgApproveRevokeX509RootCertResponse
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

  fromJSON(_: any): MsgApproveRevokeX509RootCertResponse {
    const message = { ...baseMsgApproveRevokeX509RootCertResponse } as MsgApproveRevokeX509RootCertResponse
    return message
  },

  toJSON(_: MsgApproveRevokeX509RootCertResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgApproveRevokeX509RootCertResponse>): MsgApproveRevokeX509RootCertResponse {
    const message = { ...baseMsgApproveRevokeX509RootCertResponse } as MsgApproveRevokeX509RootCertResponse
    return message
  }
}

const baseMsgRevokeX509Cert: object = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 }

export const MsgRevokeX509Cert = {
  encode(message: MsgRevokeX509Cert, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.subject !== '') {
      writer.uint32(18).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(26).string(message.subjectKeyId)
    }
    if (message.info !== '') {
      writer.uint32(34).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(40).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRevokeX509Cert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRevokeX509Cert } as MsgRevokeX509Cert
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.subject = reader.string()
          break
        case 3:
          message.subjectKeyId = reader.string()
          break
        case 4:
          message.info = reader.string()
          break
        case 5:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgRevokeX509Cert {
    const message = { ...baseMsgRevokeX509Cert } as MsgRevokeX509Cert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgRevokeX509Cert): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgRevokeX509Cert>): MsgRevokeX509Cert {
    const message = { ...baseMsgRevokeX509Cert } as MsgRevokeX509Cert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgRevokeX509CertResponse: object = {}

export const MsgRevokeX509CertResponse = {
  encode(_: MsgRevokeX509CertResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRevokeX509CertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRevokeX509CertResponse } as MsgRevokeX509CertResponse
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

  fromJSON(_: any): MsgRevokeX509CertResponse {
    const message = { ...baseMsgRevokeX509CertResponse } as MsgRevokeX509CertResponse
    return message
  },

  toJSON(_: MsgRevokeX509CertResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgRevokeX509CertResponse>): MsgRevokeX509CertResponse {
    const message = { ...baseMsgRevokeX509CertResponse } as MsgRevokeX509CertResponse
    return message
  }
}

const baseMsgRejectAddX509RootCert: object = { signer: '', subject: '', subjectKeyId: '', info: '', time: 0 }

export const MsgRejectAddX509RootCert = {
  encode(message: MsgRejectAddX509RootCert, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.subject !== '') {
      writer.uint32(18).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(26).string(message.subjectKeyId)
    }
    if (message.info !== '') {
      writer.uint32(34).string(message.info)
    }
    if (message.time !== 0) {
      writer.uint32(40).int64(message.time)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectAddX509RootCert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRejectAddX509RootCert } as MsgRejectAddX509RootCert
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.subject = reader.string()
          break
        case 3:
          message.subjectKeyId = reader.string()
          break
        case 4:
          message.info = reader.string()
          break
        case 5:
          message.time = longToNumber(reader.int64() as Long)
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgRejectAddX509RootCert {
    const message = { ...baseMsgRejectAddX509RootCert } as MsgRejectAddX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = String(object.info)
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = Number(object.time)
    } else {
      message.time = 0
    }
    return message
  },

  toJSON(message: MsgRejectAddX509RootCert): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.info !== undefined && (obj.info = message.info)
    message.time !== undefined && (obj.time = message.time)
    return obj
  },

  fromPartial(object: DeepPartial<MsgRejectAddX509RootCert>): MsgRejectAddX509RootCert {
    const message = { ...baseMsgRejectAddX509RootCert } as MsgRejectAddX509RootCert
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.info !== undefined && object.info !== null) {
      message.info = object.info
    } else {
      message.info = ''
    }
    if (object.time !== undefined && object.time !== null) {
      message.time = object.time
    } else {
      message.time = 0
    }
    return message
  }
}

const baseMsgRejectAddX509RootCertResponse: object = {}

export const MsgRejectAddX509RootCertResponse = {
  encode(_: MsgRejectAddX509RootCertResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectAddX509RootCertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgRejectAddX509RootCertResponse } as MsgRejectAddX509RootCertResponse
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

  fromJSON(_: any): MsgRejectAddX509RootCertResponse {
    const message = { ...baseMsgRejectAddX509RootCertResponse } as MsgRejectAddX509RootCertResponse
    return message
  },

  toJSON(_: MsgRejectAddX509RootCertResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgRejectAddX509RootCertResponse>): MsgRejectAddX509RootCertResponse {
    const message = { ...baseMsgRejectAddX509RootCertResponse } as MsgRejectAddX509RootCertResponse
    return message
  }
}

const baseMsgAddPkiRevocationDistributionPoint: object = {
  signer: '',
  vid: 0,
  pid: 0,
  isPAA: false,
  label: '',
  crlSignerCertificate: '',
  issuerSubjectKeyID: '',
  dataURL: '',
  dataFileSize: 0,
  dataDigest: '',
  dataDigestType: 0,
  revocationType: 0
}

export const MsgAddPkiRevocationDistributionPoint = {
  encode(message: MsgAddPkiRevocationDistributionPoint, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.vid !== 0) {
      writer.uint32(16).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(24).int32(message.pid)
    }
    if (message.isPAA === true) {
      writer.uint32(32).bool(message.isPAA)
    }
    if (message.label !== '') {
      writer.uint32(42).string(message.label)
    }
    if (message.crlSignerCertificate !== '') {
      writer.uint32(50).string(message.crlSignerCertificate)
    }
    if (message.issuerSubjectKeyID !== '') {
      writer.uint32(58).string(message.issuerSubjectKeyID)
    }
    if (message.dataURL !== '') {
      writer.uint32(66).string(message.dataURL)
    }
    if (message.dataFileSize !== 0) {
      writer.uint32(72).uint64(message.dataFileSize)
    }
    if (message.dataDigest !== '') {
      writer.uint32(82).string(message.dataDigest)
    }
    if (message.dataDigestType !== 0) {
      writer.uint32(88).uint32(message.dataDigestType)
    }
    if (message.revocationType !== 0) {
      writer.uint32(96).uint32(message.revocationType)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddPkiRevocationDistributionPoint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddPkiRevocationDistributionPoint } as MsgAddPkiRevocationDistributionPoint
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.vid = reader.int32()
          break
        case 3:
          message.pid = reader.int32()
          break
        case 4:
          message.isPAA = reader.bool()
          break
        case 5:
          message.label = reader.string()
          break
        case 6:
          message.crlSignerCertificate = reader.string()
          break
        case 7:
          message.issuerSubjectKeyID = reader.string()
          break
        case 8:
          message.dataURL = reader.string()
          break
        case 9:
          message.dataFileSize = longToNumber(reader.uint64() as Long)
          break
        case 10:
          message.dataDigest = reader.string()
          break
        case 11:
          message.dataDigestType = reader.uint32()
          break
        case 12:
          message.revocationType = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgAddPkiRevocationDistributionPoint {
    const message = { ...baseMsgAddPkiRevocationDistributionPoint } as MsgAddPkiRevocationDistributionPoint
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.isPAA !== undefined && object.isPAA !== null) {
      message.isPAA = Boolean(object.isPAA)
    } else {
      message.isPAA = false
    }
    if (object.label !== undefined && object.label !== null) {
      message.label = String(object.label)
    } else {
      message.label = ''
    }
    if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
      message.crlSignerCertificate = String(object.crlSignerCertificate)
    } else {
      message.crlSignerCertificate = ''
    }
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = String(object.issuerSubjectKeyID)
    } else {
      message.issuerSubjectKeyID = ''
    }
    if (object.dataURL !== undefined && object.dataURL !== null) {
      message.dataURL = String(object.dataURL)
    } else {
      message.dataURL = ''
    }
    if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
      message.dataFileSize = Number(object.dataFileSize)
    } else {
      message.dataFileSize = 0
    }
    if (object.dataDigest !== undefined && object.dataDigest !== null) {
      message.dataDigest = String(object.dataDigest)
    } else {
      message.dataDigest = ''
    }
    if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
      message.dataDigestType = Number(object.dataDigestType)
    } else {
      message.dataDigestType = 0
    }
    if (object.revocationType !== undefined && object.revocationType !== null) {
      message.revocationType = Number(object.revocationType)
    } else {
      message.revocationType = 0
    }
    return message
  },

  toJSON(message: MsgAddPkiRevocationDistributionPoint): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.isPAA !== undefined && (obj.isPAA = message.isPAA)
    message.label !== undefined && (obj.label = message.label)
    message.crlSignerCertificate !== undefined && (obj.crlSignerCertificate = message.crlSignerCertificate)
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID)
    message.dataURL !== undefined && (obj.dataURL = message.dataURL)
    message.dataFileSize !== undefined && (obj.dataFileSize = message.dataFileSize)
    message.dataDigest !== undefined && (obj.dataDigest = message.dataDigest)
    message.dataDigestType !== undefined && (obj.dataDigestType = message.dataDigestType)
    message.revocationType !== undefined && (obj.revocationType = message.revocationType)
    return obj
  },

  fromPartial(object: DeepPartial<MsgAddPkiRevocationDistributionPoint>): MsgAddPkiRevocationDistributionPoint {
    const message = { ...baseMsgAddPkiRevocationDistributionPoint } as MsgAddPkiRevocationDistributionPoint
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.isPAA !== undefined && object.isPAA !== null) {
      message.isPAA = object.isPAA
    } else {
      message.isPAA = false
    }
    if (object.label !== undefined && object.label !== null) {
      message.label = object.label
    } else {
      message.label = ''
    }
    if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
      message.crlSignerCertificate = object.crlSignerCertificate
    } else {
      message.crlSignerCertificate = ''
    }
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = object.issuerSubjectKeyID
    } else {
      message.issuerSubjectKeyID = ''
    }
    if (object.dataURL !== undefined && object.dataURL !== null) {
      message.dataURL = object.dataURL
    } else {
      message.dataURL = ''
    }
    if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
      message.dataFileSize = object.dataFileSize
    } else {
      message.dataFileSize = 0
    }
    if (object.dataDigest !== undefined && object.dataDigest !== null) {
      message.dataDigest = object.dataDigest
    } else {
      message.dataDigest = ''
    }
    if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
      message.dataDigestType = object.dataDigestType
    } else {
      message.dataDigestType = 0
    }
    if (object.revocationType !== undefined && object.revocationType !== null) {
      message.revocationType = object.revocationType
    } else {
      message.revocationType = 0
    }
    return message
  }
}

const baseMsgAddPkiRevocationDistributionPointResponse: object = {}

export const MsgAddPkiRevocationDistributionPointResponse = {
  encode(_: MsgAddPkiRevocationDistributionPointResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddPkiRevocationDistributionPointResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAddPkiRevocationDistributionPointResponse } as MsgAddPkiRevocationDistributionPointResponse
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

  fromJSON(_: any): MsgAddPkiRevocationDistributionPointResponse {
    const message = { ...baseMsgAddPkiRevocationDistributionPointResponse } as MsgAddPkiRevocationDistributionPointResponse
    return message
  },

  toJSON(_: MsgAddPkiRevocationDistributionPointResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgAddPkiRevocationDistributionPointResponse>): MsgAddPkiRevocationDistributionPointResponse {
    const message = { ...baseMsgAddPkiRevocationDistributionPointResponse } as MsgAddPkiRevocationDistributionPointResponse
    return message
  }
}

const baseMsgUpdatePkiRevocationDistributionPoint: object = {
  signer: '',
  vid: 0,
  label: '',
  crlSignerCertificate: '',
  issuerSubjectKeyID: '',
  dataURL: '',
  dataFileSize: 0,
  dataDigest: '',
  dataDigestType: 0
}

export const MsgUpdatePkiRevocationDistributionPoint = {
  encode(message: MsgUpdatePkiRevocationDistributionPoint, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.vid !== 0) {
      writer.uint32(16).int32(message.vid)
    }
    if (message.label !== '') {
      writer.uint32(26).string(message.label)
    }
    if (message.crlSignerCertificate !== '') {
      writer.uint32(34).string(message.crlSignerCertificate)
    }
    if (message.issuerSubjectKeyID !== '') {
      writer.uint32(42).string(message.issuerSubjectKeyID)
    }
    if (message.dataURL !== '') {
      writer.uint32(50).string(message.dataURL)
    }
    if (message.dataFileSize !== 0) {
      writer.uint32(56).uint64(message.dataFileSize)
    }
    if (message.dataDigest !== '') {
      writer.uint32(66).string(message.dataDigest)
    }
    if (message.dataDigestType !== 0) {
      writer.uint32(72).uint32(message.dataDigestType)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdatePkiRevocationDistributionPoint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdatePkiRevocationDistributionPoint } as MsgUpdatePkiRevocationDistributionPoint
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.vid = reader.int32()
          break
        case 3:
          message.label = reader.string()
          break
        case 4:
          message.crlSignerCertificate = reader.string()
          break
        case 5:
          message.issuerSubjectKeyID = reader.string()
          break
        case 6:
          message.dataURL = reader.string()
          break
        case 7:
          message.dataFileSize = longToNumber(reader.uint64() as Long)
          break
        case 8:
          message.dataDigest = reader.string()
          break
        case 9:
          message.dataDigestType = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgUpdatePkiRevocationDistributionPoint {
    const message = { ...baseMsgUpdatePkiRevocationDistributionPoint } as MsgUpdatePkiRevocationDistributionPoint
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
      message.crlSignerCertificate = String(object.crlSignerCertificate)
    } else {
      message.crlSignerCertificate = ''
    }
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = String(object.issuerSubjectKeyID)
    } else {
      message.issuerSubjectKeyID = ''
    }
    if (object.dataURL !== undefined && object.dataURL !== null) {
      message.dataURL = String(object.dataURL)
    } else {
      message.dataURL = ''
    }
    if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
      message.dataFileSize = Number(object.dataFileSize)
    } else {
      message.dataFileSize = 0
    }
    if (object.dataDigest !== undefined && object.dataDigest !== null) {
      message.dataDigest = String(object.dataDigest)
    } else {
      message.dataDigest = ''
    }
    if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
      message.dataDigestType = Number(object.dataDigestType)
    } else {
      message.dataDigestType = 0
    }
    return message
  },

  toJSON(message: MsgUpdatePkiRevocationDistributionPoint): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.vid !== undefined && (obj.vid = message.vid)
    message.label !== undefined && (obj.label = message.label)
    message.crlSignerCertificate !== undefined && (obj.crlSignerCertificate = message.crlSignerCertificate)
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID)
    message.dataURL !== undefined && (obj.dataURL = message.dataURL)
    message.dataFileSize !== undefined && (obj.dataFileSize = message.dataFileSize)
    message.dataDigest !== undefined && (obj.dataDigest = message.dataDigest)
    message.dataDigestType !== undefined && (obj.dataDigestType = message.dataDigestType)
    return obj
  },

  fromPartial(object: DeepPartial<MsgUpdatePkiRevocationDistributionPoint>): MsgUpdatePkiRevocationDistributionPoint {
    const message = { ...baseMsgUpdatePkiRevocationDistributionPoint } as MsgUpdatePkiRevocationDistributionPoint
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
      message.crlSignerCertificate = object.crlSignerCertificate
    } else {
      message.crlSignerCertificate = ''
    }
    if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
      message.issuerSubjectKeyID = object.issuerSubjectKeyID
    } else {
      message.issuerSubjectKeyID = ''
    }
    if (object.dataURL !== undefined && object.dataURL !== null) {
      message.dataURL = object.dataURL
    } else {
      message.dataURL = ''
    }
    if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
      message.dataFileSize = object.dataFileSize
    } else {
      message.dataFileSize = 0
    }
    if (object.dataDigest !== undefined && object.dataDigest !== null) {
      message.dataDigest = object.dataDigest
    } else {
      message.dataDigest = ''
    }
    if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
      message.dataDigestType = object.dataDigestType
    } else {
      message.dataDigestType = 0
    }
    return message
  }
}

const baseMsgUpdatePkiRevocationDistributionPointResponse: object = {}

export const MsgUpdatePkiRevocationDistributionPointResponse = {
  encode(_: MsgUpdatePkiRevocationDistributionPointResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdatePkiRevocationDistributionPointResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgUpdatePkiRevocationDistributionPointResponse } as MsgUpdatePkiRevocationDistributionPointResponse
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

  fromJSON(_: any): MsgUpdatePkiRevocationDistributionPointResponse {
    const message = { ...baseMsgUpdatePkiRevocationDistributionPointResponse } as MsgUpdatePkiRevocationDistributionPointResponse
    return message
  },

  toJSON(_: MsgUpdatePkiRevocationDistributionPointResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgUpdatePkiRevocationDistributionPointResponse>): MsgUpdatePkiRevocationDistributionPointResponse {
    const message = { ...baseMsgUpdatePkiRevocationDistributionPointResponse } as MsgUpdatePkiRevocationDistributionPointResponse
    return message
  }
}

const baseMsgDeletePkiRevocationDistributionPoint: object = { signer: '', vid: 0, label: '', issuerSubjectKeyID: '' }

export const MsgDeletePkiRevocationDistributionPoint = {
  encode(message: MsgDeletePkiRevocationDistributionPoint, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.vid !== 0) {
      writer.uint32(16).int32(message.vid)
    }
    if (message.label !== '') {
      writer.uint32(26).string(message.label)
    }
    if (message.issuerSubjectKeyID !== '') {
      writer.uint32(34).string(message.issuerSubjectKeyID)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeletePkiRevocationDistributionPoint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeletePkiRevocationDistributionPoint } as MsgDeletePkiRevocationDistributionPoint
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.vid = reader.int32()
          break
        case 3:
          message.label = reader.string()
          break
        case 4:
          message.issuerSubjectKeyID = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgDeletePkiRevocationDistributionPoint {
    const message = { ...baseMsgDeletePkiRevocationDistributionPoint } as MsgDeletePkiRevocationDistributionPoint
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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

  toJSON(message: MsgDeletePkiRevocationDistributionPoint): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.vid !== undefined && (obj.vid = message.vid)
    message.label !== undefined && (obj.label = message.label)
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID)
    return obj
  },

  fromPartial(object: DeepPartial<MsgDeletePkiRevocationDistributionPoint>): MsgDeletePkiRevocationDistributionPoint {
    const message = { ...baseMsgDeletePkiRevocationDistributionPoint } as MsgDeletePkiRevocationDistributionPoint
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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

const baseMsgDeletePkiRevocationDistributionPointResponse: object = {}

export const MsgDeletePkiRevocationDistributionPointResponse = {
  encode(_: MsgDeletePkiRevocationDistributionPointResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeletePkiRevocationDistributionPointResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgDeletePkiRevocationDistributionPointResponse } as MsgDeletePkiRevocationDistributionPointResponse
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

  fromJSON(_: any): MsgDeletePkiRevocationDistributionPointResponse {
    const message = { ...baseMsgDeletePkiRevocationDistributionPointResponse } as MsgDeletePkiRevocationDistributionPointResponse
    return message
  },

  toJSON(_: MsgDeletePkiRevocationDistributionPointResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgDeletePkiRevocationDistributionPointResponse>): MsgDeletePkiRevocationDistributionPointResponse {
    const message = { ...baseMsgDeletePkiRevocationDistributionPointResponse } as MsgDeletePkiRevocationDistributionPointResponse
    return message
  }
}

const baseMsgAssignVid: object = { signer: '', subject: '', subjectKeyId: '', vid: 0 }

export const MsgAssignVid = {
  encode(message: MsgAssignVid, writer: Writer = Writer.create()): Writer {
    if (message.signer !== '') {
      writer.uint32(10).string(message.signer)
    }
    if (message.subject !== '') {
      writer.uint32(18).string(message.subject)
    }
    if (message.subjectKeyId !== '') {
      writer.uint32(26).string(message.subjectKeyId)
    }
    if (message.vid !== 0) {
      writer.uint32(32).int32(message.vid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAssignVid {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAssignVid } as MsgAssignVid
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.signer = reader.string()
          break
        case 2:
          message.subject = reader.string()
          break
        case 3:
          message.subjectKeyId = reader.string()
          break
        case 4:
          message.vid = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgAssignVid {
    const message = { ...baseMsgAssignVid } as MsgAssignVid
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = String(object.signer)
    } else {
      message.signer = ''
    }
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
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    return message
  },

  toJSON(message: MsgAssignVid): unknown {
    const obj: any = {}
    message.signer !== undefined && (obj.signer = message.signer)
    message.subject !== undefined && (obj.subject = message.subject)
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId)
    message.vid !== undefined && (obj.vid = message.vid)
    return obj
  },

  fromPartial(object: DeepPartial<MsgAssignVid>): MsgAssignVid {
    const message = { ...baseMsgAssignVid } as MsgAssignVid
    if (object.signer !== undefined && object.signer !== null) {
      message.signer = object.signer
    } else {
      message.signer = ''
    }
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
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    return message
  }
}

const baseMsgAssignVidResponse: object = {}

export const MsgAssignVidResponse = {
  encode(_: MsgAssignVidResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAssignVidResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAssignVidResponse } as MsgAssignVidResponse
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

  fromJSON(_: any): MsgAssignVidResponse {
    const message = { ...baseMsgAssignVidResponse } as MsgAssignVidResponse
    return message
  },

  toJSON(_: MsgAssignVidResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgAssignVidResponse>): MsgAssignVidResponse {
    const message = { ...baseMsgAssignVidResponse } as MsgAssignVidResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  ProposeAddX509RootCert(request: MsgProposeAddX509RootCert): Promise<MsgProposeAddX509RootCertResponse>
  ApproveAddX509RootCert(request: MsgApproveAddX509RootCert): Promise<MsgApproveAddX509RootCertResponse>
  AddX509Cert(request: MsgAddX509Cert): Promise<MsgAddX509CertResponse>
  ProposeRevokeX509RootCert(request: MsgProposeRevokeX509RootCert): Promise<MsgProposeRevokeX509RootCertResponse>
  ApproveRevokeX509RootCert(request: MsgApproveRevokeX509RootCert): Promise<MsgApproveRevokeX509RootCertResponse>
  RevokeX509Cert(request: MsgRevokeX509Cert): Promise<MsgRevokeX509CertResponse>
  RejectAddX509RootCert(request: MsgRejectAddX509RootCert): Promise<MsgRejectAddX509RootCertResponse>
  AddPkiRevocationDistributionPoint(request: MsgAddPkiRevocationDistributionPoint): Promise<MsgAddPkiRevocationDistributionPointResponse>
  UpdatePkiRevocationDistributionPoint(request: MsgUpdatePkiRevocationDistributionPoint): Promise<MsgUpdatePkiRevocationDistributionPointResponse>
  DeletePkiRevocationDistributionPoint(request: MsgDeletePkiRevocationDistributionPoint): Promise<MsgDeletePkiRevocationDistributionPointResponse>
  /** this line is used by starport scaffolding # proto/tx/rpc */
  AssignVid(request: MsgAssignVid): Promise<MsgAssignVidResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  ProposeAddX509RootCert(request: MsgProposeAddX509RootCert): Promise<MsgProposeAddX509RootCertResponse> {
    const data = MsgProposeAddX509RootCert.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ProposeAddX509RootCert', data)
    return promise.then((data) => MsgProposeAddX509RootCertResponse.decode(new Reader(data)))
  }

  ApproveAddX509RootCert(request: MsgApproveAddX509RootCert): Promise<MsgApproveAddX509RootCertResponse> {
    const data = MsgApproveAddX509RootCert.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ApproveAddX509RootCert', data)
    return promise.then((data) => MsgApproveAddX509RootCertResponse.decode(new Reader(data)))
  }

  AddX509Cert(request: MsgAddX509Cert): Promise<MsgAddX509CertResponse> {
    const data = MsgAddX509Cert.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'AddX509Cert', data)
    return promise.then((data) => MsgAddX509CertResponse.decode(new Reader(data)))
  }

  ProposeRevokeX509RootCert(request: MsgProposeRevokeX509RootCert): Promise<MsgProposeRevokeX509RootCertResponse> {
    const data = MsgProposeRevokeX509RootCert.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ProposeRevokeX509RootCert', data)
    return promise.then((data) => MsgProposeRevokeX509RootCertResponse.decode(new Reader(data)))
  }

  ApproveRevokeX509RootCert(request: MsgApproveRevokeX509RootCert): Promise<MsgApproveRevokeX509RootCertResponse> {
    const data = MsgApproveRevokeX509RootCert.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'ApproveRevokeX509RootCert', data)
    return promise.then((data) => MsgApproveRevokeX509RootCertResponse.decode(new Reader(data)))
  }

  RevokeX509Cert(request: MsgRevokeX509Cert): Promise<MsgRevokeX509CertResponse> {
    const data = MsgRevokeX509Cert.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'RevokeX509Cert', data)
    return promise.then((data) => MsgRevokeX509CertResponse.decode(new Reader(data)))
  }

  RejectAddX509RootCert(request: MsgRejectAddX509RootCert): Promise<MsgRejectAddX509RootCertResponse> {
    const data = MsgRejectAddX509RootCert.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'RejectAddX509RootCert', data)
    return promise.then((data) => MsgRejectAddX509RootCertResponse.decode(new Reader(data)))
  }

  AddPkiRevocationDistributionPoint(request: MsgAddPkiRevocationDistributionPoint): Promise<MsgAddPkiRevocationDistributionPointResponse> {
    const data = MsgAddPkiRevocationDistributionPoint.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'AddPkiRevocationDistributionPoint', data)
    return promise.then((data) => MsgAddPkiRevocationDistributionPointResponse.decode(new Reader(data)))
  }

  UpdatePkiRevocationDistributionPoint(request: MsgUpdatePkiRevocationDistributionPoint): Promise<MsgUpdatePkiRevocationDistributionPointResponse> {
    const data = MsgUpdatePkiRevocationDistributionPoint.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'UpdatePkiRevocationDistributionPoint', data)
    return promise.then((data) => MsgUpdatePkiRevocationDistributionPointResponse.decode(new Reader(data)))
  }

  DeletePkiRevocationDistributionPoint(request: MsgDeletePkiRevocationDistributionPoint): Promise<MsgDeletePkiRevocationDistributionPointResponse> {
    const data = MsgDeletePkiRevocationDistributionPoint.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'DeletePkiRevocationDistributionPoint', data)
    return promise.then((data) => MsgDeletePkiRevocationDistributionPointResponse.decode(new Reader(data)))
  }

  AssignVid(request: MsgAssignVid): Promise<MsgAssignVidResponse> {
    const data = MsgAssignVid.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Msg', 'AssignVid', data)
    return promise.then((data) => MsgAssignVidResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') return globalThis
  if (typeof self !== 'undefined') return self
  if (typeof window !== 'undefined') return window
  if (typeof global !== 'undefined') return global
  throw 'Unable to locate global object'
})()

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (util.Long !== Long) {
  util.Long = Long as any
  configure()
}
