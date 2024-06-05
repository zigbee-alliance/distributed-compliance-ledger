/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { ApprovedCertificates } from "./approved_certificates";
import { ApprovedCertificatesBySubject } from "./approved_certificates_by_subject";
import { ApprovedRootCertificates } from "./approved_root_certificates";
import { ChildCertificates } from "./child_certificates";
import { NocIcaCertificates } from "./noc_ica_certificates";
import { NocRootCertificates } from "./noc_root_certificates";
import { NocRootCertificatesByVidAndSkid } from "./noc_root_certificates_by_vid_and_skid";
import { PkiRevocationDistributionPoint } from "./pki_revocation_distribution_point";
import { PkiRevocationDistributionPointsByIssuerSubjectKeyID } from "./pki_revocation_distribution_points_by_issuer_subject_key_id";
import { ProposedCertificate } from "./proposed_certificate";
import { ProposedCertificateRevocation } from "./proposed_certificate_revocation";
import { RejectedCertificate } from "./rejected_certificate";
import { RevokedCertificates } from "./revoked_certificates";
import { RevokedNocRootCertificates } from "./revoked_noc_root_certificates";
import { RevokedRootCertificates } from "./revoked_root_certificates";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface QueryGetApprovedCertificatesRequest {
  subject: string;
  subjectKeyId: string;
}

export interface QueryGetApprovedCertificatesResponse {
  approvedCertificates: ApprovedCertificates | undefined;
}

export interface QueryAllApprovedCertificatesRequest {
  pagination: PageRequest | undefined;
  subjectKeyId: string;
}

export interface QueryAllApprovedCertificatesResponse {
  approvedCertificates: ApprovedCertificates[];
  pagination: PageResponse | undefined;
}

export interface QueryGetProposedCertificateRequest {
  subject: string;
  subjectKeyId: string;
}

export interface QueryGetProposedCertificateResponse {
  proposedCertificate: ProposedCertificate | undefined;
}

export interface QueryAllProposedCertificateRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllProposedCertificateResponse {
  proposedCertificate: ProposedCertificate[];
  pagination: PageResponse | undefined;
}

export interface QueryGetChildCertificatesRequest {
  issuer: string;
  authorityKeyId: string;
}

export interface QueryGetChildCertificatesResponse {
  childCertificates: ChildCertificates | undefined;
}

export interface QueryGetProposedCertificateRevocationRequest {
  subject: string;
  subjectKeyId: string;
  serialNumber: string;
}

export interface QueryGetProposedCertificateRevocationResponse {
  proposedCertificateRevocation: ProposedCertificateRevocation | undefined;
}

export interface QueryAllProposedCertificateRevocationRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllProposedCertificateRevocationResponse {
  proposedCertificateRevocation: ProposedCertificateRevocation[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRevokedCertificatesRequest {
  subject: string;
  subjectKeyId: string;
}

export interface QueryGetRevokedCertificatesResponse {
  revokedCertificates: RevokedCertificates | undefined;
}

export interface QueryAllRevokedCertificatesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRevokedCertificatesResponse {
  revokedCertificates: RevokedCertificates[];
  pagination: PageResponse | undefined;
}

export interface QueryGetApprovedRootCertificatesRequest {
}

export interface QueryGetApprovedRootCertificatesResponse {
  approvedRootCertificates: ApprovedRootCertificates | undefined;
}

export interface QueryGetRevokedRootCertificatesRequest {
}

export interface QueryGetRevokedRootCertificatesResponse {
  revokedRootCertificates: RevokedRootCertificates | undefined;
}

export interface QueryGetApprovedCertificatesBySubjectRequest {
  subject: string;
}

export interface QueryGetApprovedCertificatesBySubjectResponse {
  approvedCertificatesBySubject: ApprovedCertificatesBySubject | undefined;
}

export interface QueryGetRejectedCertificatesRequest {
  subject: string;
  subjectKeyId: string;
}

export interface QueryGetRejectedCertificatesResponse {
  rejectedCertificate: RejectedCertificate | undefined;
}

export interface QueryAllRejectedCertificatesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRejectedCertificatesResponse {
  rejectedCertificate: RejectedCertificate[];
  pagination: PageResponse | undefined;
}

export interface QueryGetPkiRevocationDistributionPointRequest {
  vid: number;
  label: string;
  issuerSubjectKeyID: string;
}

export interface QueryGetPkiRevocationDistributionPointResponse {
  PkiRevocationDistributionPoint: PkiRevocationDistributionPoint | undefined;
}

export interface QueryAllPkiRevocationDistributionPointRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPkiRevocationDistributionPointResponse {
  PkiRevocationDistributionPoint: PkiRevocationDistributionPoint[];
  pagination: PageResponse | undefined;
}

export interface QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
  issuerSubjectKeyID: string;
}

export interface QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
  pkiRevocationDistributionPointsByIssuerSubjectKeyID: PkiRevocationDistributionPointsByIssuerSubjectKeyID | undefined;
}

export interface QueryGetNocRootCertificatesRequest {
  vid: number;
}

export interface QueryGetNocRootCertificatesResponse {
  nocRootCertificates: NocRootCertificates | undefined;
}

export interface QueryAllNocRootCertificatesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNocRootCertificatesResponse {
  nocRootCertificates: NocRootCertificates[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNocIcaCertificatesRequest {
  vid: number;
}

export interface QueryGetNocIcaCertificatesResponse {
  nocIcaCertificates: NocIcaCertificates | undefined;
}

export interface QueryAllNocIcaCertificatesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNocIcaCertificatesResponse {
  nocIcaCertificates: NocIcaCertificates[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRevokedNocRootCertificatesRequest {
  subject: string;
  subjectKeyId: string;
}

export interface QueryGetRevokedNocRootCertificatesResponse {
  revokedNocRootCertificates: RevokedNocRootCertificates | undefined;
}

export interface QueryAllRevokedNocRootCertificatesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRevokedNocRootCertificatesResponse {
  revokedNocRootCertificates: RevokedNocRootCertificates[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNocRootCertificatesByVidAndSkidRequest {
  vid: number;
  subjectKeyId: string;
}

export interface QueryGetNocRootCertificatesByVidAndSkidResponse {
  nocRootCertificatesByVidAndSkid: NocRootCertificatesByVidAndSkid | undefined;
}

function createBaseQueryGetApprovedCertificatesRequest(): QueryGetApprovedCertificatesRequest {
  return { subject: "", subjectKeyId: "" };
}

export const QueryGetApprovedCertificatesRequest = {
  encode(message: QueryGetApprovedCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetApprovedCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetApprovedCertificatesRequest {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: QueryGetApprovedCertificatesRequest): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetApprovedCertificatesRequest>, I>>(
    object: I,
  ): QueryGetApprovedCertificatesRequest {
    const message = createBaseQueryGetApprovedCertificatesRequest();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    return message;
  },
};

function createBaseQueryGetApprovedCertificatesResponse(): QueryGetApprovedCertificatesResponse {
  return { approvedCertificates: undefined };
}

export const QueryGetApprovedCertificatesResponse = {
  encode(message: QueryGetApprovedCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.approvedCertificates !== undefined) {
      ApprovedCertificates.encode(message.approvedCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetApprovedCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificates = ApprovedCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetApprovedCertificatesResponse {
    return {
      approvedCertificates: isSet(object.approvedCertificates)
        ? ApprovedCertificates.fromJSON(object.approvedCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetApprovedCertificatesResponse): unknown {
    const obj: any = {};
    message.approvedCertificates !== undefined && (obj.approvedCertificates = message.approvedCertificates
      ? ApprovedCertificates.toJSON(message.approvedCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetApprovedCertificatesResponse>, I>>(
    object: I,
  ): QueryGetApprovedCertificatesResponse {
    const message = createBaseQueryGetApprovedCertificatesResponse();
    message.approvedCertificates = (object.approvedCertificates !== undefined && object.approvedCertificates !== null)
      ? ApprovedCertificates.fromPartial(object.approvedCertificates)
      : undefined;
    return message;
  },
};

function createBaseQueryAllApprovedCertificatesRequest(): QueryAllApprovedCertificatesRequest {
  return { pagination: undefined, subjectKeyId: "" };
}

export const QueryAllApprovedCertificatesRequest = {
  encode(message: QueryAllApprovedCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllApprovedCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllApprovedCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllApprovedCertificatesRequest {
    return {
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: QueryAllApprovedCertificatesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllApprovedCertificatesRequest>, I>>(
    object: I,
  ): QueryAllApprovedCertificatesRequest {
    const message = createBaseQueryAllApprovedCertificatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    message.subjectKeyId = object.subjectKeyId ?? "";
    return message;
  },
};

function createBaseQueryAllApprovedCertificatesResponse(): QueryAllApprovedCertificatesResponse {
  return { approvedCertificates: [], pagination: undefined };
}

export const QueryAllApprovedCertificatesResponse = {
  encode(message: QueryAllApprovedCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.approvedCertificates) {
      ApprovedCertificates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllApprovedCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllApprovedCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificates.push(ApprovedCertificates.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllApprovedCertificatesResponse {
    return {
      approvedCertificates: Array.isArray(object?.approvedCertificates)
        ? object.approvedCertificates.map((e: any) => ApprovedCertificates.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllApprovedCertificatesResponse): unknown {
    const obj: any = {};
    if (message.approvedCertificates) {
      obj.approvedCertificates = message.approvedCertificates.map((e) =>
        e ? ApprovedCertificates.toJSON(e) : undefined
      );
    } else {
      obj.approvedCertificates = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllApprovedCertificatesResponse>, I>>(
    object: I,
  ): QueryAllApprovedCertificatesResponse {
    const message = createBaseQueryAllApprovedCertificatesResponse();
    message.approvedCertificates = object.approvedCertificates?.map((e) => ApprovedCertificates.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetProposedCertificateRequest(): QueryGetProposedCertificateRequest {
  return { subject: "", subjectKeyId: "" };
}

export const QueryGetProposedCertificateRequest = {
  encode(message: QueryGetProposedCertificateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProposedCertificateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProposedCertificateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProposedCertificateRequest {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: QueryGetProposedCertificateRequest): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProposedCertificateRequest>, I>>(
    object: I,
  ): QueryGetProposedCertificateRequest {
    const message = createBaseQueryGetProposedCertificateRequest();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    return message;
  },
};

function createBaseQueryGetProposedCertificateResponse(): QueryGetProposedCertificateResponse {
  return { proposedCertificate: undefined };
}

export const QueryGetProposedCertificateResponse = {
  encode(message: QueryGetProposedCertificateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.proposedCertificate !== undefined) {
      ProposedCertificate.encode(message.proposedCertificate, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProposedCertificateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProposedCertificateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificate = ProposedCertificate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProposedCertificateResponse {
    return {
      proposedCertificate: isSet(object.proposedCertificate)
        ? ProposedCertificate.fromJSON(object.proposedCertificate)
        : undefined,
    };
  },

  toJSON(message: QueryGetProposedCertificateResponse): unknown {
    const obj: any = {};
    message.proposedCertificate !== undefined && (obj.proposedCertificate = message.proposedCertificate
      ? ProposedCertificate.toJSON(message.proposedCertificate)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProposedCertificateResponse>, I>>(
    object: I,
  ): QueryGetProposedCertificateResponse {
    const message = createBaseQueryGetProposedCertificateResponse();
    message.proposedCertificate = (object.proposedCertificate !== undefined && object.proposedCertificate !== null)
      ? ProposedCertificate.fromPartial(object.proposedCertificate)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProposedCertificateRequest(): QueryAllProposedCertificateRequest {
  return { pagination: undefined };
}

export const QueryAllProposedCertificateRequest = {
  encode(message: QueryAllProposedCertificateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProposedCertificateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProposedCertificateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProposedCertificateRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllProposedCertificateRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProposedCertificateRequest>, I>>(
    object: I,
  ): QueryAllProposedCertificateRequest {
    const message = createBaseQueryAllProposedCertificateRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProposedCertificateResponse(): QueryAllProposedCertificateResponse {
  return { proposedCertificate: [], pagination: undefined };
}

export const QueryAllProposedCertificateResponse = {
  encode(message: QueryAllProposedCertificateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.proposedCertificate) {
      ProposedCertificate.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProposedCertificateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProposedCertificateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificate.push(ProposedCertificate.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProposedCertificateResponse {
    return {
      proposedCertificate: Array.isArray(object?.proposedCertificate)
        ? object.proposedCertificate.map((e: any) => ProposedCertificate.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllProposedCertificateResponse): unknown {
    const obj: any = {};
    if (message.proposedCertificate) {
      obj.proposedCertificate = message.proposedCertificate.map((e) => e ? ProposedCertificate.toJSON(e) : undefined);
    } else {
      obj.proposedCertificate = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProposedCertificateResponse>, I>>(
    object: I,
  ): QueryAllProposedCertificateResponse {
    const message = createBaseQueryAllProposedCertificateResponse();
    message.proposedCertificate = object.proposedCertificate?.map((e) => ProposedCertificate.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetChildCertificatesRequest(): QueryGetChildCertificatesRequest {
  return { issuer: "", authorityKeyId: "" };
}

export const QueryGetChildCertificatesRequest = {
  encode(message: QueryGetChildCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    if (message.authorityKeyId !== "") {
      writer.uint32(18).string(message.authorityKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetChildCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetChildCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string();
          break;
        case 2:
          message.authorityKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetChildCertificatesRequest {
    return {
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      authorityKeyId: isSet(object.authorityKeyId) ? String(object.authorityKeyId) : "",
    };
  },

  toJSON(message: QueryGetChildCertificatesRequest): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetChildCertificatesRequest>, I>>(
    object: I,
  ): QueryGetChildCertificatesRequest {
    const message = createBaseQueryGetChildCertificatesRequest();
    message.issuer = object.issuer ?? "";
    message.authorityKeyId = object.authorityKeyId ?? "";
    return message;
  },
};

function createBaseQueryGetChildCertificatesResponse(): QueryGetChildCertificatesResponse {
  return { childCertificates: undefined };
}

export const QueryGetChildCertificatesResponse = {
  encode(message: QueryGetChildCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.childCertificates !== undefined) {
      ChildCertificates.encode(message.childCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetChildCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetChildCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.childCertificates = ChildCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetChildCertificatesResponse {
    return {
      childCertificates: isSet(object.childCertificates)
        ? ChildCertificates.fromJSON(object.childCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetChildCertificatesResponse): unknown {
    const obj: any = {};
    message.childCertificates !== undefined && (obj.childCertificates = message.childCertificates
      ? ChildCertificates.toJSON(message.childCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetChildCertificatesResponse>, I>>(
    object: I,
  ): QueryGetChildCertificatesResponse {
    const message = createBaseQueryGetChildCertificatesResponse();
    message.childCertificates = (object.childCertificates !== undefined && object.childCertificates !== null)
      ? ChildCertificates.fromPartial(object.childCertificates)
      : undefined;
    return message;
  },
};

function createBaseQueryGetProposedCertificateRevocationRequest(): QueryGetProposedCertificateRevocationRequest {
  return { subject: "", subjectKeyId: "", serialNumber: "" };
}

export const QueryGetProposedCertificateRevocationRequest = {
  encode(message: QueryGetProposedCertificateRevocationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    if (message.serialNumber !== "") {
      writer.uint32(26).string(message.serialNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProposedCertificateRevocationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProposedCertificateRevocationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        case 3:
          message.serialNumber = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProposedCertificateRevocationRequest {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
      serialNumber: isSet(object.serialNumber) ? String(object.serialNumber) : "",
    };
  },

  toJSON(message: QueryGetProposedCertificateRevocationRequest): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProposedCertificateRevocationRequest>, I>>(
    object: I,
  ): QueryGetProposedCertificateRevocationRequest {
    const message = createBaseQueryGetProposedCertificateRevocationRequest();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    message.serialNumber = object.serialNumber ?? "";
    return message;
  },
};

function createBaseQueryGetProposedCertificateRevocationResponse(): QueryGetProposedCertificateRevocationResponse {
  return { proposedCertificateRevocation: undefined };
}

export const QueryGetProposedCertificateRevocationResponse = {
  encode(message: QueryGetProposedCertificateRevocationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.proposedCertificateRevocation !== undefined) {
      ProposedCertificateRevocation.encode(message.proposedCertificateRevocation, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProposedCertificateRevocationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProposedCertificateRevocationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificateRevocation = ProposedCertificateRevocation.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProposedCertificateRevocationResponse {
    return {
      proposedCertificateRevocation: isSet(object.proposedCertificateRevocation)
        ? ProposedCertificateRevocation.fromJSON(object.proposedCertificateRevocation)
        : undefined,
    };
  },

  toJSON(message: QueryGetProposedCertificateRevocationResponse): unknown {
    const obj: any = {};
    message.proposedCertificateRevocation !== undefined
      && (obj.proposedCertificateRevocation = message.proposedCertificateRevocation
        ? ProposedCertificateRevocation.toJSON(message.proposedCertificateRevocation)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProposedCertificateRevocationResponse>, I>>(
    object: I,
  ): QueryGetProposedCertificateRevocationResponse {
    const message = createBaseQueryGetProposedCertificateRevocationResponse();
    message.proposedCertificateRevocation =
      (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null)
        ? ProposedCertificateRevocation.fromPartial(object.proposedCertificateRevocation)
        : undefined;
    return message;
  },
};

function createBaseQueryAllProposedCertificateRevocationRequest(): QueryAllProposedCertificateRevocationRequest {
  return { pagination: undefined };
}

export const QueryAllProposedCertificateRevocationRequest = {
  encode(message: QueryAllProposedCertificateRevocationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProposedCertificateRevocationRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProposedCertificateRevocationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProposedCertificateRevocationRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllProposedCertificateRevocationRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProposedCertificateRevocationRequest>, I>>(
    object: I,
  ): QueryAllProposedCertificateRevocationRequest {
    const message = createBaseQueryAllProposedCertificateRevocationRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProposedCertificateRevocationResponse(): QueryAllProposedCertificateRevocationResponse {
  return { proposedCertificateRevocation: [], pagination: undefined };
}

export const QueryAllProposedCertificateRevocationResponse = {
  encode(message: QueryAllProposedCertificateRevocationResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.proposedCertificateRevocation) {
      ProposedCertificateRevocation.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProposedCertificateRevocationResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProposedCertificateRevocationResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposedCertificateRevocation.push(ProposedCertificateRevocation.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProposedCertificateRevocationResponse {
    return {
      proposedCertificateRevocation: Array.isArray(object?.proposedCertificateRevocation)
        ? object.proposedCertificateRevocation.map((e: any) => ProposedCertificateRevocation.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllProposedCertificateRevocationResponse): unknown {
    const obj: any = {};
    if (message.proposedCertificateRevocation) {
      obj.proposedCertificateRevocation = message.proposedCertificateRevocation.map((e) =>
        e ? ProposedCertificateRevocation.toJSON(e) : undefined
      );
    } else {
      obj.proposedCertificateRevocation = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProposedCertificateRevocationResponse>, I>>(
    object: I,
  ): QueryAllProposedCertificateRevocationResponse {
    const message = createBaseQueryAllProposedCertificateRevocationResponse();
    message.proposedCertificateRevocation =
      object.proposedCertificateRevocation?.map((e) => ProposedCertificateRevocation.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRevokedCertificatesRequest(): QueryGetRevokedCertificatesRequest {
  return { subject: "", subjectKeyId: "" };
}

export const QueryGetRevokedCertificatesRequest = {
  encode(message: QueryGetRevokedCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRevokedCertificatesRequest {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: QueryGetRevokedCertificatesRequest): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedCertificatesRequest>, I>>(
    object: I,
  ): QueryGetRevokedCertificatesRequest {
    const message = createBaseQueryGetRevokedCertificatesRequest();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    return message;
  },
};

function createBaseQueryGetRevokedCertificatesResponse(): QueryGetRevokedCertificatesResponse {
  return { revokedCertificates: undefined };
}

export const QueryGetRevokedCertificatesResponse = {
  encode(message: QueryGetRevokedCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.revokedCertificates !== undefined) {
      RevokedCertificates.encode(message.revokedCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revokedCertificates = RevokedCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRevokedCertificatesResponse {
    return {
      revokedCertificates: isSet(object.revokedCertificates)
        ? RevokedCertificates.fromJSON(object.revokedCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetRevokedCertificatesResponse): unknown {
    const obj: any = {};
    message.revokedCertificates !== undefined && (obj.revokedCertificates = message.revokedCertificates
      ? RevokedCertificates.toJSON(message.revokedCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedCertificatesResponse>, I>>(
    object: I,
  ): QueryGetRevokedCertificatesResponse {
    const message = createBaseQueryGetRevokedCertificatesResponse();
    message.revokedCertificates = (object.revokedCertificates !== undefined && object.revokedCertificates !== null)
      ? RevokedCertificates.fromPartial(object.revokedCertificates)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRevokedCertificatesRequest(): QueryAllRevokedCertificatesRequest {
  return { pagination: undefined };
}

export const QueryAllRevokedCertificatesRequest = {
  encode(message: QueryAllRevokedCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRevokedCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRevokedCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRevokedCertificatesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRevokedCertificatesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRevokedCertificatesRequest>, I>>(
    object: I,
  ): QueryAllRevokedCertificatesRequest {
    const message = createBaseQueryAllRevokedCertificatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRevokedCertificatesResponse(): QueryAllRevokedCertificatesResponse {
  return { revokedCertificates: [], pagination: undefined };
}

export const QueryAllRevokedCertificatesResponse = {
  encode(message: QueryAllRevokedCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.revokedCertificates) {
      RevokedCertificates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRevokedCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRevokedCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revokedCertificates.push(RevokedCertificates.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRevokedCertificatesResponse {
    return {
      revokedCertificates: Array.isArray(object?.revokedCertificates)
        ? object.revokedCertificates.map((e: any) => RevokedCertificates.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRevokedCertificatesResponse): unknown {
    const obj: any = {};
    if (message.revokedCertificates) {
      obj.revokedCertificates = message.revokedCertificates.map((e) => e ? RevokedCertificates.toJSON(e) : undefined);
    } else {
      obj.revokedCertificates = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRevokedCertificatesResponse>, I>>(
    object: I,
  ): QueryAllRevokedCertificatesResponse {
    const message = createBaseQueryAllRevokedCertificatesResponse();
    message.revokedCertificates = object.revokedCertificates?.map((e) => RevokedCertificates.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetApprovedRootCertificatesRequest(): QueryGetApprovedRootCertificatesRequest {
  return {};
}

export const QueryGetApprovedRootCertificatesRequest = {
  encode(_: QueryGetApprovedRootCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetApprovedRootCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetApprovedRootCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetApprovedRootCertificatesRequest {
    return {};
  },

  toJSON(_: QueryGetApprovedRootCertificatesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetApprovedRootCertificatesRequest>, I>>(
    _: I,
  ): QueryGetApprovedRootCertificatesRequest {
    const message = createBaseQueryGetApprovedRootCertificatesRequest();
    return message;
  },
};

function createBaseQueryGetApprovedRootCertificatesResponse(): QueryGetApprovedRootCertificatesResponse {
  return { approvedRootCertificates: undefined };
}

export const QueryGetApprovedRootCertificatesResponse = {
  encode(message: QueryGetApprovedRootCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.approvedRootCertificates !== undefined) {
      ApprovedRootCertificates.encode(message.approvedRootCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetApprovedRootCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetApprovedRootCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.approvedRootCertificates = ApprovedRootCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetApprovedRootCertificatesResponse {
    return {
      approvedRootCertificates: isSet(object.approvedRootCertificates)
        ? ApprovedRootCertificates.fromJSON(object.approvedRootCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetApprovedRootCertificatesResponse): unknown {
    const obj: any = {};
    message.approvedRootCertificates !== undefined && (obj.approvedRootCertificates = message.approvedRootCertificates
      ? ApprovedRootCertificates.toJSON(message.approvedRootCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetApprovedRootCertificatesResponse>, I>>(
    object: I,
  ): QueryGetApprovedRootCertificatesResponse {
    const message = createBaseQueryGetApprovedRootCertificatesResponse();
    message.approvedRootCertificates =
      (object.approvedRootCertificates !== undefined && object.approvedRootCertificates !== null)
        ? ApprovedRootCertificates.fromPartial(object.approvedRootCertificates)
        : undefined;
    return message;
  },
};

function createBaseQueryGetRevokedRootCertificatesRequest(): QueryGetRevokedRootCertificatesRequest {
  return {};
}

export const QueryGetRevokedRootCertificatesRequest = {
  encode(_: QueryGetRevokedRootCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedRootCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedRootCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetRevokedRootCertificatesRequest {
    return {};
  },

  toJSON(_: QueryGetRevokedRootCertificatesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedRootCertificatesRequest>, I>>(
    _: I,
  ): QueryGetRevokedRootCertificatesRequest {
    const message = createBaseQueryGetRevokedRootCertificatesRequest();
    return message;
  },
};

function createBaseQueryGetRevokedRootCertificatesResponse(): QueryGetRevokedRootCertificatesResponse {
  return { revokedRootCertificates: undefined };
}

export const QueryGetRevokedRootCertificatesResponse = {
  encode(message: QueryGetRevokedRootCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.revokedRootCertificates !== undefined) {
      RevokedRootCertificates.encode(message.revokedRootCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedRootCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedRootCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revokedRootCertificates = RevokedRootCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRevokedRootCertificatesResponse {
    return {
      revokedRootCertificates: isSet(object.revokedRootCertificates)
        ? RevokedRootCertificates.fromJSON(object.revokedRootCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetRevokedRootCertificatesResponse): unknown {
    const obj: any = {};
    message.revokedRootCertificates !== undefined && (obj.revokedRootCertificates = message.revokedRootCertificates
      ? RevokedRootCertificates.toJSON(message.revokedRootCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedRootCertificatesResponse>, I>>(
    object: I,
  ): QueryGetRevokedRootCertificatesResponse {
    const message = createBaseQueryGetRevokedRootCertificatesResponse();
    message.revokedRootCertificates =
      (object.revokedRootCertificates !== undefined && object.revokedRootCertificates !== null)
        ? RevokedRootCertificates.fromPartial(object.revokedRootCertificates)
        : undefined;
    return message;
  },
};

function createBaseQueryGetApprovedCertificatesBySubjectRequest(): QueryGetApprovedCertificatesBySubjectRequest {
  return { subject: "" };
}

export const QueryGetApprovedCertificatesBySubjectRequest = {
  encode(message: QueryGetApprovedCertificatesBySubjectRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesBySubjectRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetApprovedCertificatesBySubjectRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetApprovedCertificatesBySubjectRequest {
    return { subject: isSet(object.subject) ? String(object.subject) : "" };
  },

  toJSON(message: QueryGetApprovedCertificatesBySubjectRequest): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetApprovedCertificatesBySubjectRequest>, I>>(
    object: I,
  ): QueryGetApprovedCertificatesBySubjectRequest {
    const message = createBaseQueryGetApprovedCertificatesBySubjectRequest();
    message.subject = object.subject ?? "";
    return message;
  },
};

function createBaseQueryGetApprovedCertificatesBySubjectResponse(): QueryGetApprovedCertificatesBySubjectResponse {
  return { approvedCertificatesBySubject: undefined };
}

export const QueryGetApprovedCertificatesBySubjectResponse = {
  encode(message: QueryGetApprovedCertificatesBySubjectResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.approvedCertificatesBySubject !== undefined) {
      ApprovedCertificatesBySubject.encode(message.approvedCertificatesBySubject, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesBySubjectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetApprovedCertificatesBySubjectResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificatesBySubject = ApprovedCertificatesBySubject.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetApprovedCertificatesBySubjectResponse {
    return {
      approvedCertificatesBySubject: isSet(object.approvedCertificatesBySubject)
        ? ApprovedCertificatesBySubject.fromJSON(object.approvedCertificatesBySubject)
        : undefined,
    };
  },

  toJSON(message: QueryGetApprovedCertificatesBySubjectResponse): unknown {
    const obj: any = {};
    message.approvedCertificatesBySubject !== undefined
      && (obj.approvedCertificatesBySubject = message.approvedCertificatesBySubject
        ? ApprovedCertificatesBySubject.toJSON(message.approvedCertificatesBySubject)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetApprovedCertificatesBySubjectResponse>, I>>(
    object: I,
  ): QueryGetApprovedCertificatesBySubjectResponse {
    const message = createBaseQueryGetApprovedCertificatesBySubjectResponse();
    message.approvedCertificatesBySubject =
      (object.approvedCertificatesBySubject !== undefined && object.approvedCertificatesBySubject !== null)
        ? ApprovedCertificatesBySubject.fromPartial(object.approvedCertificatesBySubject)
        : undefined;
    return message;
  },
};

function createBaseQueryGetRejectedCertificatesRequest(): QueryGetRejectedCertificatesRequest {
  return { subject: "", subjectKeyId: "" };
}

export const QueryGetRejectedCertificatesRequest = {
  encode(message: QueryGetRejectedCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRejectedCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRejectedCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRejectedCertificatesRequest {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: QueryGetRejectedCertificatesRequest): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRejectedCertificatesRequest>, I>>(
    object: I,
  ): QueryGetRejectedCertificatesRequest {
    const message = createBaseQueryGetRejectedCertificatesRequest();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    return message;
  },
};

function createBaseQueryGetRejectedCertificatesResponse(): QueryGetRejectedCertificatesResponse {
  return { rejectedCertificate: undefined };
}

export const QueryGetRejectedCertificatesResponse = {
  encode(message: QueryGetRejectedCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.rejectedCertificate !== undefined) {
      RejectedCertificate.encode(message.rejectedCertificate, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRejectedCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRejectedCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rejectedCertificate = RejectedCertificate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRejectedCertificatesResponse {
    return {
      rejectedCertificate: isSet(object.rejectedCertificate)
        ? RejectedCertificate.fromJSON(object.rejectedCertificate)
        : undefined,
    };
  },

  toJSON(message: QueryGetRejectedCertificatesResponse): unknown {
    const obj: any = {};
    message.rejectedCertificate !== undefined && (obj.rejectedCertificate = message.rejectedCertificate
      ? RejectedCertificate.toJSON(message.rejectedCertificate)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRejectedCertificatesResponse>, I>>(
    object: I,
  ): QueryGetRejectedCertificatesResponse {
    const message = createBaseQueryGetRejectedCertificatesResponse();
    message.rejectedCertificate = (object.rejectedCertificate !== undefined && object.rejectedCertificate !== null)
      ? RejectedCertificate.fromPartial(object.rejectedCertificate)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRejectedCertificatesRequest(): QueryAllRejectedCertificatesRequest {
  return { pagination: undefined };
}

export const QueryAllRejectedCertificatesRequest = {
  encode(message: QueryAllRejectedCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRejectedCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRejectedCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRejectedCertificatesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRejectedCertificatesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRejectedCertificatesRequest>, I>>(
    object: I,
  ): QueryAllRejectedCertificatesRequest {
    const message = createBaseQueryAllRejectedCertificatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRejectedCertificatesResponse(): QueryAllRejectedCertificatesResponse {
  return { rejectedCertificate: [], pagination: undefined };
}

export const QueryAllRejectedCertificatesResponse = {
  encode(message: QueryAllRejectedCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.rejectedCertificate) {
      RejectedCertificate.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRejectedCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRejectedCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rejectedCertificate.push(RejectedCertificate.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRejectedCertificatesResponse {
    return {
      rejectedCertificate: Array.isArray(object?.rejectedCertificate)
        ? object.rejectedCertificate.map((e: any) => RejectedCertificate.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRejectedCertificatesResponse): unknown {
    const obj: any = {};
    if (message.rejectedCertificate) {
      obj.rejectedCertificate = message.rejectedCertificate.map((e) => e ? RejectedCertificate.toJSON(e) : undefined);
    } else {
      obj.rejectedCertificate = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRejectedCertificatesResponse>, I>>(
    object: I,
  ): QueryAllRejectedCertificatesResponse {
    const message = createBaseQueryAllRejectedCertificatesResponse();
    message.rejectedCertificate = object.rejectedCertificate?.map((e) => RejectedCertificate.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetPkiRevocationDistributionPointRequest(): QueryGetPkiRevocationDistributionPointRequest {
  return { vid: 0, label: "", issuerSubjectKeyID: "" };
}

export const QueryGetPkiRevocationDistributionPointRequest = {
  encode(message: QueryGetPkiRevocationDistributionPointRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.label !== "") {
      writer.uint32(18).string(message.label);
    }
    if (message.issuerSubjectKeyID !== "") {
      writer.uint32(26).string(message.issuerSubjectKeyID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetPkiRevocationDistributionPointRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetPkiRevocationDistributionPointRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.label = reader.string();
          break;
        case 3:
          message.issuerSubjectKeyID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointRequest {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      label: isSet(object.label) ? String(object.label) : "",
      issuerSubjectKeyID: isSet(object.issuerSubjectKeyID) ? String(object.issuerSubjectKeyID) : "",
    };
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.label !== undefined && (obj.label = message.label);
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetPkiRevocationDistributionPointRequest>, I>>(
    object: I,
  ): QueryGetPkiRevocationDistributionPointRequest {
    const message = createBaseQueryGetPkiRevocationDistributionPointRequest();
    message.vid = object.vid ?? 0;
    message.label = object.label ?? "";
    message.issuerSubjectKeyID = object.issuerSubjectKeyID ?? "";
    return message;
  },
};

function createBaseQueryGetPkiRevocationDistributionPointResponse(): QueryGetPkiRevocationDistributionPointResponse {
  return { PkiRevocationDistributionPoint: undefined };
}

export const QueryGetPkiRevocationDistributionPointResponse = {
  encode(
    message: QueryGetPkiRevocationDistributionPointResponse,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.PkiRevocationDistributionPoint !== undefined) {
      PkiRevocationDistributionPoint.encode(message.PkiRevocationDistributionPoint, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetPkiRevocationDistributionPointResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetPkiRevocationDistributionPointResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.PkiRevocationDistributionPoint = PkiRevocationDistributionPoint.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointResponse {
    return {
      PkiRevocationDistributionPoint: isSet(object.PkiRevocationDistributionPoint)
        ? PkiRevocationDistributionPoint.fromJSON(object.PkiRevocationDistributionPoint)
        : undefined,
    };
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointResponse): unknown {
    const obj: any = {};
    message.PkiRevocationDistributionPoint !== undefined
      && (obj.PkiRevocationDistributionPoint = message.PkiRevocationDistributionPoint
        ? PkiRevocationDistributionPoint.toJSON(message.PkiRevocationDistributionPoint)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetPkiRevocationDistributionPointResponse>, I>>(
    object: I,
  ): QueryGetPkiRevocationDistributionPointResponse {
    const message = createBaseQueryGetPkiRevocationDistributionPointResponse();
    message.PkiRevocationDistributionPoint =
      (object.PkiRevocationDistributionPoint !== undefined && object.PkiRevocationDistributionPoint !== null)
        ? PkiRevocationDistributionPoint.fromPartial(object.PkiRevocationDistributionPoint)
        : undefined;
    return message;
  },
};

function createBaseQueryAllPkiRevocationDistributionPointRequest(): QueryAllPkiRevocationDistributionPointRequest {
  return { pagination: undefined };
}

export const QueryAllPkiRevocationDistributionPointRequest = {
  encode(message: QueryAllPkiRevocationDistributionPointRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllPkiRevocationDistributionPointRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllPkiRevocationDistributionPointRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPkiRevocationDistributionPointRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllPkiRevocationDistributionPointRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllPkiRevocationDistributionPointRequest>, I>>(
    object: I,
  ): QueryAllPkiRevocationDistributionPointRequest {
    const message = createBaseQueryAllPkiRevocationDistributionPointRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllPkiRevocationDistributionPointResponse(): QueryAllPkiRevocationDistributionPointResponse {
  return { PkiRevocationDistributionPoint: [], pagination: undefined };
}

export const QueryAllPkiRevocationDistributionPointResponse = {
  encode(
    message: QueryAllPkiRevocationDistributionPointResponse,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    for (const v of message.PkiRevocationDistributionPoint) {
      PkiRevocationDistributionPoint.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllPkiRevocationDistributionPointResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllPkiRevocationDistributionPointResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.PkiRevocationDistributionPoint.push(PkiRevocationDistributionPoint.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPkiRevocationDistributionPointResponse {
    return {
      PkiRevocationDistributionPoint: Array.isArray(object?.PkiRevocationDistributionPoint)
        ? object.PkiRevocationDistributionPoint.map((e: any) => PkiRevocationDistributionPoint.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllPkiRevocationDistributionPointResponse): unknown {
    const obj: any = {};
    if (message.PkiRevocationDistributionPoint) {
      obj.PkiRevocationDistributionPoint = message.PkiRevocationDistributionPoint.map((e) =>
        e ? PkiRevocationDistributionPoint.toJSON(e) : undefined
      );
    } else {
      obj.PkiRevocationDistributionPoint = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllPkiRevocationDistributionPointResponse>, I>>(
    object: I,
  ): QueryAllPkiRevocationDistributionPointResponse {
    const message = createBaseQueryAllPkiRevocationDistributionPointResponse();
    message.PkiRevocationDistributionPoint =
      object.PkiRevocationDistributionPoint?.map((e) => PkiRevocationDistributionPoint.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest(): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
  return { issuerSubjectKeyID: "" };
}

export const QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest = {
  encode(
    message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.issuerSubjectKeyID !== "") {
      writer.uint32(10).string(message.issuerSubjectKeyID);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuerSubjectKeyID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
    return { issuerSubjectKeyID: isSet(object.issuerSubjectKeyID) ? String(object.issuerSubjectKeyID) : "" };
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest): unknown {
    const obj: any = {};
    message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest>, I>>(
    object: I,
  ): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest {
    const message = createBaseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest();
    message.issuerSubjectKeyID = object.issuerSubjectKeyID ?? "";
    return message;
  },
};

function createBaseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse(): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
  return { pkiRevocationDistributionPointsByIssuerSubjectKeyID: undefined };
}

export const QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse = {
  encode(
    message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== undefined) {
      PkiRevocationDistributionPointsByIssuerSubjectKeyID.encode(
        message.pkiRevocationDistributionPointsByIssuerSubjectKeyID,
        writer.uint32(10).fork(),
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number,
  ): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pkiRevocationDistributionPointsByIssuerSubjectKeyID =
            PkiRevocationDistributionPointsByIssuerSubjectKeyID.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
    return {
      pkiRevocationDistributionPointsByIssuerSubjectKeyID:
        isSet(object.pkiRevocationDistributionPointsByIssuerSubjectKeyID)
          ? PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromJSON(
            object.pkiRevocationDistributionPointsByIssuerSubjectKeyID,
          )
          : undefined,
    };
  },

  toJSON(message: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse): unknown {
    const obj: any = {};
    message.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== undefined
      && (obj.pkiRevocationDistributionPointsByIssuerSubjectKeyID =
        message.pkiRevocationDistributionPointsByIssuerSubjectKeyID
          ? PkiRevocationDistributionPointsByIssuerSubjectKeyID.toJSON(
            message.pkiRevocationDistributionPointsByIssuerSubjectKeyID,
          )
          : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse>, I>>(
    object: I,
  ): QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
    const message = createBaseQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse();
    message.pkiRevocationDistributionPointsByIssuerSubjectKeyID =
      (object.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== undefined
          && object.pkiRevocationDistributionPointsByIssuerSubjectKeyID !== null)
        ? PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromPartial(
          object.pkiRevocationDistributionPointsByIssuerSubjectKeyID,
        )
        : undefined;
    return message;
  },
};

function createBaseQueryGetNocRootCertificatesRequest(): QueryGetNocRootCertificatesRequest {
  return { vid: 0 };
}

export const QueryGetNocRootCertificatesRequest = {
  encode(message: QueryGetNocRootCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNocRootCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetNocRootCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNocRootCertificatesRequest {
    return { vid: isSet(object.vid) ? Number(object.vid) : 0 };
  },

  toJSON(message: QueryGetNocRootCertificatesRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetNocRootCertificatesRequest>, I>>(
    object: I,
  ): QueryGetNocRootCertificatesRequest {
    const message = createBaseQueryGetNocRootCertificatesRequest();
    message.vid = object.vid ?? 0;
    return message;
  },
};

function createBaseQueryGetNocRootCertificatesResponse(): QueryGetNocRootCertificatesResponse {
  return { nocRootCertificates: undefined };
}

export const QueryGetNocRootCertificatesResponse = {
  encode(message: QueryGetNocRootCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.nocRootCertificates !== undefined) {
      NocRootCertificates.encode(message.nocRootCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNocRootCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetNocRootCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nocRootCertificates = NocRootCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNocRootCertificatesResponse {
    return {
      nocRootCertificates: isSet(object.nocRootCertificates)
        ? NocRootCertificates.fromJSON(object.nocRootCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetNocRootCertificatesResponse): unknown {
    const obj: any = {};
    message.nocRootCertificates !== undefined && (obj.nocRootCertificates = message.nocRootCertificates
      ? NocRootCertificates.toJSON(message.nocRootCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetNocRootCertificatesResponse>, I>>(
    object: I,
  ): QueryGetNocRootCertificatesResponse {
    const message = createBaseQueryGetNocRootCertificatesResponse();
    message.nocRootCertificates = (object.nocRootCertificates !== undefined && object.nocRootCertificates !== null)
      ? NocRootCertificates.fromPartial(object.nocRootCertificates)
      : undefined;
    return message;
  },
};

function createBaseQueryAllNocRootCertificatesRequest(): QueryAllNocRootCertificatesRequest {
  return { pagination: undefined };
}

export const QueryAllNocRootCertificatesRequest = {
  encode(message: QueryAllNocRootCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllNocRootCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllNocRootCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNocRootCertificatesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllNocRootCertificatesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllNocRootCertificatesRequest>, I>>(
    object: I,
  ): QueryAllNocRootCertificatesRequest {
    const message = createBaseQueryAllNocRootCertificatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllNocRootCertificatesResponse(): QueryAllNocRootCertificatesResponse {
  return { nocRootCertificates: [], pagination: undefined };
}

export const QueryAllNocRootCertificatesResponse = {
  encode(message: QueryAllNocRootCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.nocRootCertificates) {
      NocRootCertificates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllNocRootCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllNocRootCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nocRootCertificates.push(NocRootCertificates.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNocRootCertificatesResponse {
    return {
      nocRootCertificates: Array.isArray(object?.nocRootCertificates)
        ? object.nocRootCertificates.map((e: any) => NocRootCertificates.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllNocRootCertificatesResponse): unknown {
    const obj: any = {};
    if (message.nocRootCertificates) {
      obj.nocRootCertificates = message.nocRootCertificates.map((e) => e ? NocRootCertificates.toJSON(e) : undefined);
    } else {
      obj.nocRootCertificates = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllNocRootCertificatesResponse>, I>>(
    object: I,
  ): QueryAllNocRootCertificatesResponse {
    const message = createBaseQueryAllNocRootCertificatesResponse();
    message.nocRootCertificates = object.nocRootCertificates?.map((e) => NocRootCertificates.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetNocIcaCertificatesRequest(): QueryGetNocIcaCertificatesRequest {
  return { vid: 0 };
}

export const QueryGetNocIcaCertificatesRequest = {
  encode(message: QueryGetNocIcaCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNocIcaCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetNocIcaCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNocIcaCertificatesRequest {
    return { vid: isSet(object.vid) ? Number(object.vid) : 0 };
  },

  toJSON(message: QueryGetNocIcaCertificatesRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetNocIcaCertificatesRequest>, I>>(
    object: I,
  ): QueryGetNocIcaCertificatesRequest {
    const message = createBaseQueryGetNocIcaCertificatesRequest();
    message.vid = object.vid ?? 0;
    return message;
  },
};

function createBaseQueryGetNocIcaCertificatesResponse(): QueryGetNocIcaCertificatesResponse {
  return { nocIcaCertificates: undefined };
}

export const QueryGetNocIcaCertificatesResponse = {
  encode(message: QueryGetNocIcaCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.nocIcaCertificates !== undefined) {
      NocIcaCertificates.encode(message.nocIcaCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNocIcaCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetNocIcaCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nocIcaCertificates = NocIcaCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNocIcaCertificatesResponse {
    return {
      nocIcaCertificates: isSet(object.nocIcaCertificates)
        ? NocIcaCertificates.fromJSON(object.nocIcaCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetNocIcaCertificatesResponse): unknown {
    const obj: any = {};
    message.nocIcaCertificates !== undefined && (obj.nocIcaCertificates = message.nocIcaCertificates
      ? NocIcaCertificates.toJSON(message.nocIcaCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetNocIcaCertificatesResponse>, I>>(
    object: I,
  ): QueryGetNocIcaCertificatesResponse {
    const message = createBaseQueryGetNocIcaCertificatesResponse();
    message.nocIcaCertificates = (object.nocIcaCertificates !== undefined && object.nocIcaCertificates !== null)
      ? NocIcaCertificates.fromPartial(object.nocIcaCertificates)
      : undefined;
    return message;
  },
};

function createBaseQueryAllNocIcaCertificatesRequest(): QueryAllNocIcaCertificatesRequest {
  return { pagination: undefined };
}

export const QueryAllNocIcaCertificatesRequest = {
  encode(message: QueryAllNocIcaCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllNocIcaCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllNocIcaCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNocIcaCertificatesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllNocIcaCertificatesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllNocIcaCertificatesRequest>, I>>(
    object: I,
  ): QueryAllNocIcaCertificatesRequest {
    const message = createBaseQueryAllNocIcaCertificatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllNocIcaCertificatesResponse(): QueryAllNocIcaCertificatesResponse {
  return { nocIcaCertificates: [], pagination: undefined };
}

export const QueryAllNocIcaCertificatesResponse = {
  encode(message: QueryAllNocIcaCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.nocIcaCertificates) {
      NocIcaCertificates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllNocIcaCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllNocIcaCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nocIcaCertificates.push(NocIcaCertificates.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNocIcaCertificatesResponse {
    return {
      nocIcaCertificates: Array.isArray(object?.nocIcaCertificates)
        ? object.nocIcaCertificates.map((e: any) => NocIcaCertificates.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllNocIcaCertificatesResponse): unknown {
    const obj: any = {};
    if (message.nocIcaCertificates) {
      obj.nocIcaCertificates = message.nocIcaCertificates.map((e) => e ? NocIcaCertificates.toJSON(e) : undefined);
    } else {
      obj.nocIcaCertificates = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllNocIcaCertificatesResponse>, I>>(
    object: I,
  ): QueryAllNocIcaCertificatesResponse {
    const message = createBaseQueryAllNocIcaCertificatesResponse();
    message.nocIcaCertificates = object.nocIcaCertificates?.map((e) => NocIcaCertificates.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRevokedNocRootCertificatesRequest(): QueryGetRevokedNocRootCertificatesRequest {
  return { subject: "", subjectKeyId: "" };
}

export const QueryGetRevokedNocRootCertificatesRequest = {
  encode(message: QueryGetRevokedNocRootCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== "") {
      writer.uint32(10).string(message.subject);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedNocRootCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedNocRootCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subject = reader.string();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRevokedNocRootCertificatesRequest {
    return {
      subject: isSet(object.subject) ? String(object.subject) : "",
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: QueryGetRevokedNocRootCertificatesRequest): unknown {
    const obj: any = {};
    message.subject !== undefined && (obj.subject = message.subject);
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedNocRootCertificatesRequest>, I>>(
    object: I,
  ): QueryGetRevokedNocRootCertificatesRequest {
    const message = createBaseQueryGetRevokedNocRootCertificatesRequest();
    message.subject = object.subject ?? "";
    message.subjectKeyId = object.subjectKeyId ?? "";
    return message;
  },
};

function createBaseQueryGetRevokedNocRootCertificatesResponse(): QueryGetRevokedNocRootCertificatesResponse {
  return { revokedNocRootCertificates: undefined };
}

export const QueryGetRevokedNocRootCertificatesResponse = {
  encode(message: QueryGetRevokedNocRootCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.revokedNocRootCertificates !== undefined) {
      RevokedNocRootCertificates.encode(message.revokedNocRootCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedNocRootCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedNocRootCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revokedNocRootCertificates = RevokedNocRootCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRevokedNocRootCertificatesResponse {
    return {
      revokedNocRootCertificates: isSet(object.revokedNocRootCertificates)
        ? RevokedNocRootCertificates.fromJSON(object.revokedNocRootCertificates)
        : undefined,
    };
  },

  toJSON(message: QueryGetRevokedNocRootCertificatesResponse): unknown {
    const obj: any = {};
    message.revokedNocRootCertificates !== undefined
      && (obj.revokedNocRootCertificates = message.revokedNocRootCertificates
        ? RevokedNocRootCertificates.toJSON(message.revokedNocRootCertificates)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedNocRootCertificatesResponse>, I>>(
    object: I,
  ): QueryGetRevokedNocRootCertificatesResponse {
    const message = createBaseQueryGetRevokedNocRootCertificatesResponse();
    message.revokedNocRootCertificates =
      (object.revokedNocRootCertificates !== undefined && object.revokedNocRootCertificates !== null)
        ? RevokedNocRootCertificates.fromPartial(object.revokedNocRootCertificates)
        : undefined;
    return message;
  },
};

function createBaseQueryAllRevokedNocRootCertificatesRequest(): QueryAllRevokedNocRootCertificatesRequest {
  return { pagination: undefined };
}

export const QueryAllRevokedNocRootCertificatesRequest = {
  encode(message: QueryAllRevokedNocRootCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRevokedNocRootCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRevokedNocRootCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRevokedNocRootCertificatesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRevokedNocRootCertificatesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRevokedNocRootCertificatesRequest>, I>>(
    object: I,
  ): QueryAllRevokedNocRootCertificatesRequest {
    const message = createBaseQueryAllRevokedNocRootCertificatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRevokedNocRootCertificatesResponse(): QueryAllRevokedNocRootCertificatesResponse {
  return { revokedNocRootCertificates: [], pagination: undefined };
}

export const QueryAllRevokedNocRootCertificatesResponse = {
  encode(message: QueryAllRevokedNocRootCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.revokedNocRootCertificates) {
      RevokedNocRootCertificates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRevokedNocRootCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRevokedNocRootCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revokedNocRootCertificates.push(RevokedNocRootCertificates.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRevokedNocRootCertificatesResponse {
    return {
      revokedNocRootCertificates: Array.isArray(object?.revokedNocRootCertificates)
        ? object.revokedNocRootCertificates.map((e: any) => RevokedNocRootCertificates.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRevokedNocRootCertificatesResponse): unknown {
    const obj: any = {};
    if (message.revokedNocRootCertificates) {
      obj.revokedNocRootCertificates = message.revokedNocRootCertificates.map((e) =>
        e ? RevokedNocRootCertificates.toJSON(e) : undefined
      );
    } else {
      obj.revokedNocRootCertificates = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRevokedNocRootCertificatesResponse>, I>>(
    object: I,
  ): QueryAllRevokedNocRootCertificatesResponse {
    const message = createBaseQueryAllRevokedNocRootCertificatesResponse();
    message.revokedNocRootCertificates =
      object.revokedNocRootCertificates?.map((e) => RevokedNocRootCertificates.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetNocRootCertificatesByVidAndSkidRequest(): QueryGetNocRootCertificatesByVidAndSkidRequest {
  return { vid: 0, subjectKeyId: "" };
}

export const QueryGetNocRootCertificatesByVidAndSkidRequest = {
  encode(
    message: QueryGetNocRootCertificatesByVidAndSkidRequest,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.subjectKeyId !== "") {
      writer.uint32(18).string(message.subjectKeyId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNocRootCertificatesByVidAndSkidRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetNocRootCertificatesByVidAndSkidRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.subjectKeyId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNocRootCertificatesByVidAndSkidRequest {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      subjectKeyId: isSet(object.subjectKeyId) ? String(object.subjectKeyId) : "",
    };
  },

  toJSON(message: QueryGetNocRootCertificatesByVidAndSkidRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetNocRootCertificatesByVidAndSkidRequest>, I>>(
    object: I,
  ): QueryGetNocRootCertificatesByVidAndSkidRequest {
    const message = createBaseQueryGetNocRootCertificatesByVidAndSkidRequest();
    message.vid = object.vid ?? 0;
    message.subjectKeyId = object.subjectKeyId ?? "";
    return message;
  },
};

function createBaseQueryGetNocRootCertificatesByVidAndSkidResponse(): QueryGetNocRootCertificatesByVidAndSkidResponse {
  return { nocRootCertificatesByVidAndSkid: undefined };
}

export const QueryGetNocRootCertificatesByVidAndSkidResponse = {
  encode(
    message: QueryGetNocRootCertificatesByVidAndSkidResponse,
    writer: _m0.Writer = _m0.Writer.create(),
  ): _m0.Writer {
    if (message.nocRootCertificatesByVidAndSkid !== undefined) {
      NocRootCertificatesByVidAndSkid.encode(message.nocRootCertificatesByVidAndSkid, writer.uint32(10).fork())
        .ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetNocRootCertificatesByVidAndSkidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetNocRootCertificatesByVidAndSkidResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.nocRootCertificatesByVidAndSkid = NocRootCertificatesByVidAndSkid.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNocRootCertificatesByVidAndSkidResponse {
    return {
      nocRootCertificatesByVidAndSkid: isSet(object.nocRootCertificatesByVidAndSkid)
        ? NocRootCertificatesByVidAndSkid.fromJSON(object.nocRootCertificatesByVidAndSkid)
        : undefined,
    };
  },

  toJSON(message: QueryGetNocRootCertificatesByVidAndSkidResponse): unknown {
    const obj: any = {};
    message.nocRootCertificatesByVidAndSkid !== undefined
      && (obj.nocRootCertificatesByVidAndSkid = message.nocRootCertificatesByVidAndSkid
        ? NocRootCertificatesByVidAndSkid.toJSON(message.nocRootCertificatesByVidAndSkid)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetNocRootCertificatesByVidAndSkidResponse>, I>>(
    object: I,
  ): QueryGetNocRootCertificatesByVidAndSkidResponse {
    const message = createBaseQueryGetNocRootCertificatesByVidAndSkidResponse();
    message.nocRootCertificatesByVidAndSkid =
      (object.nocRootCertificatesByVidAndSkid !== undefined && object.nocRootCertificatesByVidAndSkid !== null)
        ? NocRootCertificatesByVidAndSkid.fromPartial(object.nocRootCertificatesByVidAndSkid)
        : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a ApprovedCertificates by index. */
  ApprovedCertificates(request: QueryGetApprovedCertificatesRequest): Promise<QueryGetApprovedCertificatesResponse>;
  /** Queries a list of ApprovedCertificates items. */
  ApprovedCertificatesAll(request: QueryAllApprovedCertificatesRequest): Promise<QueryAllApprovedCertificatesResponse>;
  /** Queries a ProposedCertificate by index. */
  ProposedCertificate(request: QueryGetProposedCertificateRequest): Promise<QueryGetProposedCertificateResponse>;
  /** Queries a list of ProposedCertificate items. */
  ProposedCertificateAll(request: QueryAllProposedCertificateRequest): Promise<QueryAllProposedCertificateResponse>;
  /** Queries a ChildCertificates by index. */
  ChildCertificates(request: QueryGetChildCertificatesRequest): Promise<QueryGetChildCertificatesResponse>;
  /** Queries a ProposedCertificateRevocation by index. */
  ProposedCertificateRevocation(
    request: QueryGetProposedCertificateRevocationRequest,
  ): Promise<QueryGetProposedCertificateRevocationResponse>;
  /** Queries a list of ProposedCertificateRevocation items. */
  ProposedCertificateRevocationAll(
    request: QueryAllProposedCertificateRevocationRequest,
  ): Promise<QueryAllProposedCertificateRevocationResponse>;
  /** Queries a RevokedCertificates by index. */
  RevokedCertificates(request: QueryGetRevokedCertificatesRequest): Promise<QueryGetRevokedCertificatesResponse>;
  /** Queries a list of RevokedCertificates items. */
  RevokedCertificatesAll(request: QueryAllRevokedCertificatesRequest): Promise<QueryAllRevokedCertificatesResponse>;
  /** Queries a ApprovedRootCertificates by index. */
  ApprovedRootCertificates(
    request: QueryGetApprovedRootCertificatesRequest,
  ): Promise<QueryGetApprovedRootCertificatesResponse>;
  /** Queries a RevokedRootCertificates by index. */
  RevokedRootCertificates(
    request: QueryGetRevokedRootCertificatesRequest,
  ): Promise<QueryGetRevokedRootCertificatesResponse>;
  /** Queries a ApprovedCertificatesBySubject by index. */
  ApprovedCertificatesBySubject(
    request: QueryGetApprovedCertificatesBySubjectRequest,
  ): Promise<QueryGetApprovedCertificatesBySubjectResponse>;
  /** Queries a RejectedCertificate by index. */
  RejectedCertificate(request: QueryGetRejectedCertificatesRequest): Promise<QueryGetRejectedCertificatesResponse>;
  /** Queries a list of RejectedCertificate items. */
  RejectedCertificateAll(request: QueryAllRejectedCertificatesRequest): Promise<QueryAllRejectedCertificatesResponse>;
  /** Queries a PkiRevocationDistributionPoint by index. */
  PkiRevocationDistributionPoint(
    request: QueryGetPkiRevocationDistributionPointRequest,
  ): Promise<QueryGetPkiRevocationDistributionPointResponse>;
  /** Queries a list of PkiRevocationDistributionPoint items. */
  PkiRevocationDistributionPointAll(
    request: QueryAllPkiRevocationDistributionPointRequest,
  ): Promise<QueryAllPkiRevocationDistributionPointResponse>;
  /** Queries a PkiRevocationDistributionPointsByIssuerSubjectKeyID by index. */
  PkiRevocationDistributionPointsByIssuerSubjectKeyID(
    request: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest,
  ): Promise<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse>;
  /** Queries a NocRootCertificates by index. */
  NocRootCertificates(request: QueryGetNocRootCertificatesRequest): Promise<QueryGetNocRootCertificatesResponse>;
  /** Queries a NocRootCertificatesByVidAndSkid by index. */
  NocRootCertificatesByVidAndSkid(
    request: QueryGetNocRootCertificatesByVidAndSkidRequest,
  ): Promise<QueryGetNocRootCertificatesByVidAndSkidResponse>;
  /** Queries a list of NocRootCertificates items. */
  NocRootCertificatesAll(request: QueryAllNocRootCertificatesRequest): Promise<QueryAllNocRootCertificatesResponse>;
  /** Queries a NocIcaCertificates by index. */
  NocIcaCertificates(request: QueryGetNocIcaCertificatesRequest): Promise<QueryGetNocIcaCertificatesResponse>;
  /** Queries a list of NocIcaCertificates items. */
  NocIcaCertificatesAll(request: QueryAllNocIcaCertificatesRequest): Promise<QueryAllNocIcaCertificatesResponse>;
  /** Queries a RevokedNocRootCertificates by index. */
  RevokedNocRootCertificates(
    request: QueryGetRevokedNocRootCertificatesRequest,
  ): Promise<QueryGetRevokedNocRootCertificatesResponse>;
  /** Queries a list of RevokedNocRootCertificates items. */
  RevokedNocRootCertificatesAll(
    request: QueryAllRevokedNocRootCertificatesRequest,
  ): Promise<QueryAllRevokedNocRootCertificatesResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ApprovedCertificates = this.ApprovedCertificates.bind(this);
    this.ApprovedCertificatesAll = this.ApprovedCertificatesAll.bind(this);
    this.ProposedCertificate = this.ProposedCertificate.bind(this);
    this.ProposedCertificateAll = this.ProposedCertificateAll.bind(this);
    this.ChildCertificates = this.ChildCertificates.bind(this);
    this.ProposedCertificateRevocation = this.ProposedCertificateRevocation.bind(this);
    this.ProposedCertificateRevocationAll = this.ProposedCertificateRevocationAll.bind(this);
    this.RevokedCertificates = this.RevokedCertificates.bind(this);
    this.RevokedCertificatesAll = this.RevokedCertificatesAll.bind(this);
    this.ApprovedRootCertificates = this.ApprovedRootCertificates.bind(this);
    this.RevokedRootCertificates = this.RevokedRootCertificates.bind(this);
    this.ApprovedCertificatesBySubject = this.ApprovedCertificatesBySubject.bind(this);
    this.RejectedCertificate = this.RejectedCertificate.bind(this);
    this.RejectedCertificateAll = this.RejectedCertificateAll.bind(this);
    this.PkiRevocationDistributionPoint = this.PkiRevocationDistributionPoint.bind(this);
    this.PkiRevocationDistributionPointAll = this.PkiRevocationDistributionPointAll.bind(this);
    this.PkiRevocationDistributionPointsByIssuerSubjectKeyID = this.PkiRevocationDistributionPointsByIssuerSubjectKeyID
      .bind(this);
    this.NocRootCertificates = this.NocRootCertificates.bind(this);
    this.NocRootCertificatesByVidAndSkid = this.NocRootCertificatesByVidAndSkid.bind(this);
    this.NocRootCertificatesAll = this.NocRootCertificatesAll.bind(this);
    this.NocIcaCertificates = this.NocIcaCertificates.bind(this);
    this.NocIcaCertificatesAll = this.NocIcaCertificatesAll.bind(this);
    this.RevokedNocRootCertificates = this.RevokedNocRootCertificates.bind(this);
    this.RevokedNocRootCertificatesAll = this.RevokedNocRootCertificatesAll.bind(this);
  }
  ApprovedCertificates(request: QueryGetApprovedCertificatesRequest): Promise<QueryGetApprovedCertificatesResponse> {
    const data = QueryGetApprovedCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ApprovedCertificates",
      data,
    );
    return promise.then((data) => QueryGetApprovedCertificatesResponse.decode(new _m0.Reader(data)));
  }

  ApprovedCertificatesAll(request: QueryAllApprovedCertificatesRequest): Promise<QueryAllApprovedCertificatesResponse> {
    const data = QueryAllApprovedCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ApprovedCertificatesAll",
      data,
    );
    return promise.then((data) => QueryAllApprovedCertificatesResponse.decode(new _m0.Reader(data)));
  }

  ProposedCertificate(request: QueryGetProposedCertificateRequest): Promise<QueryGetProposedCertificateResponse> {
    const data = QueryGetProposedCertificateRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ProposedCertificate",
      data,
    );
    return promise.then((data) => QueryGetProposedCertificateResponse.decode(new _m0.Reader(data)));
  }

  ProposedCertificateAll(request: QueryAllProposedCertificateRequest): Promise<QueryAllProposedCertificateResponse> {
    const data = QueryAllProposedCertificateRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ProposedCertificateAll",
      data,
    );
    return promise.then((data) => QueryAllProposedCertificateResponse.decode(new _m0.Reader(data)));
  }

  ChildCertificates(request: QueryGetChildCertificatesRequest): Promise<QueryGetChildCertificatesResponse> {
    const data = QueryGetChildCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.pki.Query", "ChildCertificates", data);
    return promise.then((data) => QueryGetChildCertificatesResponse.decode(new _m0.Reader(data)));
  }

  ProposedCertificateRevocation(
    request: QueryGetProposedCertificateRevocationRequest,
  ): Promise<QueryGetProposedCertificateRevocationResponse> {
    const data = QueryGetProposedCertificateRevocationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ProposedCertificateRevocation",
      data,
    );
    return promise.then((data) => QueryGetProposedCertificateRevocationResponse.decode(new _m0.Reader(data)));
  }

  ProposedCertificateRevocationAll(
    request: QueryAllProposedCertificateRevocationRequest,
  ): Promise<QueryAllProposedCertificateRevocationResponse> {
    const data = QueryAllProposedCertificateRevocationRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ProposedCertificateRevocationAll",
      data,
    );
    return promise.then((data) => QueryAllProposedCertificateRevocationResponse.decode(new _m0.Reader(data)));
  }

  RevokedCertificates(request: QueryGetRevokedCertificatesRequest): Promise<QueryGetRevokedCertificatesResponse> {
    const data = QueryGetRevokedCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "RevokedCertificates",
      data,
    );
    return promise.then((data) => QueryGetRevokedCertificatesResponse.decode(new _m0.Reader(data)));
  }

  RevokedCertificatesAll(request: QueryAllRevokedCertificatesRequest): Promise<QueryAllRevokedCertificatesResponse> {
    const data = QueryAllRevokedCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "RevokedCertificatesAll",
      data,
    );
    return promise.then((data) => QueryAllRevokedCertificatesResponse.decode(new _m0.Reader(data)));
  }

  ApprovedRootCertificates(
    request: QueryGetApprovedRootCertificatesRequest,
  ): Promise<QueryGetApprovedRootCertificatesResponse> {
    const data = QueryGetApprovedRootCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ApprovedRootCertificates",
      data,
    );
    return promise.then((data) => QueryGetApprovedRootCertificatesResponse.decode(new _m0.Reader(data)));
  }

  RevokedRootCertificates(
    request: QueryGetRevokedRootCertificatesRequest,
  ): Promise<QueryGetRevokedRootCertificatesResponse> {
    const data = QueryGetRevokedRootCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "RevokedRootCertificates",
      data,
    );
    return promise.then((data) => QueryGetRevokedRootCertificatesResponse.decode(new _m0.Reader(data)));
  }

  ApprovedCertificatesBySubject(
    request: QueryGetApprovedCertificatesBySubjectRequest,
  ): Promise<QueryGetApprovedCertificatesBySubjectResponse> {
    const data = QueryGetApprovedCertificatesBySubjectRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "ApprovedCertificatesBySubject",
      data,
    );
    return promise.then((data) => QueryGetApprovedCertificatesBySubjectResponse.decode(new _m0.Reader(data)));
  }

  RejectedCertificate(request: QueryGetRejectedCertificatesRequest): Promise<QueryGetRejectedCertificatesResponse> {
    const data = QueryGetRejectedCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "RejectedCertificate",
      data,
    );
    return promise.then((data) => QueryGetRejectedCertificatesResponse.decode(new _m0.Reader(data)));
  }

  RejectedCertificateAll(request: QueryAllRejectedCertificatesRequest): Promise<QueryAllRejectedCertificatesResponse> {
    const data = QueryAllRejectedCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "RejectedCertificateAll",
      data,
    );
    return promise.then((data) => QueryAllRejectedCertificatesResponse.decode(new _m0.Reader(data)));
  }

  PkiRevocationDistributionPoint(
    request: QueryGetPkiRevocationDistributionPointRequest,
  ): Promise<QueryGetPkiRevocationDistributionPointResponse> {
    const data = QueryGetPkiRevocationDistributionPointRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "PkiRevocationDistributionPoint",
      data,
    );
    return promise.then((data) => QueryGetPkiRevocationDistributionPointResponse.decode(new _m0.Reader(data)));
  }

  PkiRevocationDistributionPointAll(
    request: QueryAllPkiRevocationDistributionPointRequest,
  ): Promise<QueryAllPkiRevocationDistributionPointResponse> {
    const data = QueryAllPkiRevocationDistributionPointRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "PkiRevocationDistributionPointAll",
      data,
    );
    return promise.then((data) => QueryAllPkiRevocationDistributionPointResponse.decode(new _m0.Reader(data)));
  }

  PkiRevocationDistributionPointsByIssuerSubjectKeyID(
    request: QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest,
  ): Promise<QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse> {
    const data = QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "PkiRevocationDistributionPointsByIssuerSubjectKeyID",
      data,
    );
    return promise.then((data) =>
      QueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse.decode(new _m0.Reader(data))
    );
  }

  NocRootCertificates(request: QueryGetNocRootCertificatesRequest): Promise<QueryGetNocRootCertificatesResponse> {
    const data = QueryGetNocRootCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "NocRootCertificates",
      data,
    );
    return promise.then((data) => QueryGetNocRootCertificatesResponse.decode(new _m0.Reader(data)));
  }

  NocRootCertificatesByVidAndSkid(
    request: QueryGetNocRootCertificatesByVidAndSkidRequest,
  ): Promise<QueryGetNocRootCertificatesByVidAndSkidResponse> {
    const data = QueryGetNocRootCertificatesByVidAndSkidRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "NocRootCertificatesByVidAndSkid",
      data,
    );
    return promise.then((data) => QueryGetNocRootCertificatesByVidAndSkidResponse.decode(new _m0.Reader(data)));
  }

  NocRootCertificatesAll(request: QueryAllNocRootCertificatesRequest): Promise<QueryAllNocRootCertificatesResponse> {
    const data = QueryAllNocRootCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "NocRootCertificatesAll",
      data,
    );
    return promise.then((data) => QueryAllNocRootCertificatesResponse.decode(new _m0.Reader(data)));
  }

  NocIcaCertificates(request: QueryGetNocIcaCertificatesRequest): Promise<QueryGetNocIcaCertificatesResponse> {
    const data = QueryGetNocIcaCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "NocIcaCertificates",
      data,
    );
    return promise.then((data) => QueryGetNocIcaCertificatesResponse.decode(new _m0.Reader(data)));
  }

  NocIcaCertificatesAll(request: QueryAllNocIcaCertificatesRequest): Promise<QueryAllNocIcaCertificatesResponse> {
    const data = QueryAllNocIcaCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "NocIcaCertificatesAll",
      data,
    );
    return promise.then((data) => QueryAllNocIcaCertificatesResponse.decode(new _m0.Reader(data)));
  }

  RevokedNocRootCertificates(
    request: QueryGetRevokedNocRootCertificatesRequest,
  ): Promise<QueryGetRevokedNocRootCertificatesResponse> {
    const data = QueryGetRevokedNocRootCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "RevokedNocRootCertificates",
      data,
    );
    return promise.then((data) => QueryGetRevokedNocRootCertificatesResponse.decode(new _m0.Reader(data)));
  }

  RevokedNocRootCertificatesAll(
    request: QueryAllRevokedNocRootCertificatesRequest,
  ): Promise<QueryAllRevokedNocRootCertificatesResponse> {
    const data = QueryAllRevokedNocRootCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.pki.Query",
      "RevokedNocRootCertificatesAll",
      data,
    );
    return promise.then((data) => QueryAllRevokedNocRootCertificatesResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
