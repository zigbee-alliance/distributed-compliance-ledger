/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { CertifiedModel } from "./certified_model";
import { ComplianceInfo } from "./compliance_info";
import { DeviceSoftwareCompliance } from "./device_software_compliance";
import { ProvisionalModel } from "./provisional_model";
import { RevokedModel } from "./revoked_model";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";

export interface QueryGetComplianceInfoRequest {
  vid: number;
  pid: number;
  softwareVersion: number;
  certificationType: string;
}

export interface QueryGetComplianceInfoResponse {
  complianceInfo: ComplianceInfo | undefined;
}

export interface QueryAllComplianceInfoRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllComplianceInfoResponse {
  complianceInfo: ComplianceInfo[];
  pagination: PageResponse | undefined;
}

export interface QueryGetCertifiedModelRequest {
  vid: number;
  pid: number;
  softwareVersion: number;
  certificationType: string;
}

export interface QueryGetCertifiedModelResponse {
  certifiedModel: CertifiedModel | undefined;
}

export interface QueryAllCertifiedModelRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCertifiedModelResponse {
  certifiedModel: CertifiedModel[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRevokedModelRequest {
  vid: number;
  pid: number;
  softwareVersion: number;
  certificationType: string;
}

export interface QueryGetRevokedModelResponse {
  revokedModel: RevokedModel | undefined;
}

export interface QueryAllRevokedModelRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRevokedModelResponse {
  revokedModel: RevokedModel[];
  pagination: PageResponse | undefined;
}

export interface QueryGetProvisionalModelRequest {
  vid: number;
  pid: number;
  softwareVersion: number;
  certificationType: string;
}

export interface QueryGetProvisionalModelResponse {
  provisionalModel: ProvisionalModel | undefined;
}

export interface QueryAllProvisionalModelRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllProvisionalModelResponse {
  provisionalModel: ProvisionalModel[];
  pagination: PageResponse | undefined;
}

export interface QueryGetDeviceSoftwareComplianceRequest {
  cDCertificateId: string;
}

export interface QueryGetDeviceSoftwareComplianceResponse {
  deviceSoftwareCompliance: DeviceSoftwareCompliance | undefined;
}

export interface QueryAllDeviceSoftwareComplianceRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDeviceSoftwareComplianceResponse {
  deviceSoftwareCompliance: DeviceSoftwareCompliance[];
  pagination: PageResponse | undefined;
}

function createBaseQueryGetComplianceInfoRequest(): QueryGetComplianceInfoRequest {
  return { vid: 0, pid: 0, softwareVersion: 0, certificationType: "" };
}

export const QueryGetComplianceInfoRequest = {
  encode(message: QueryGetComplianceInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion);
    }
    if (message.certificationType !== "") {
      writer.uint32(34).string(message.certificationType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetComplianceInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetComplianceInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.pid = reader.int32();
          break;
        case 3:
          message.softwareVersion = reader.uint32();
          break;
        case 4:
          message.certificationType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetComplianceInfoRequest {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
      certificationType: isSet(object.certificationType) ? String(object.certificationType) : "",
    };
  },

  toJSON(message: QueryGetComplianceInfoRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    message.certificationType !== undefined && (obj.certificationType = message.certificationType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetComplianceInfoRequest>, I>>(
    object: I,
  ): QueryGetComplianceInfoRequest {
    const message = createBaseQueryGetComplianceInfoRequest();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    message.certificationType = object.certificationType ?? "";
    return message;
  },
};

function createBaseQueryGetComplianceInfoResponse(): QueryGetComplianceInfoResponse {
  return { complianceInfo: undefined };
}

export const QueryGetComplianceInfoResponse = {
  encode(message: QueryGetComplianceInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.complianceInfo !== undefined) {
      ComplianceInfo.encode(message.complianceInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetComplianceInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetComplianceInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.complianceInfo = ComplianceInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetComplianceInfoResponse {
    return {
      complianceInfo: isSet(object.complianceInfo) ? ComplianceInfo.fromJSON(object.complianceInfo) : undefined,
    };
  },

  toJSON(message: QueryGetComplianceInfoResponse): unknown {
    const obj: any = {};
    message.complianceInfo !== undefined
      && (obj.complianceInfo = message.complianceInfo ? ComplianceInfo.toJSON(message.complianceInfo) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetComplianceInfoResponse>, I>>(
    object: I,
  ): QueryGetComplianceInfoResponse {
    const message = createBaseQueryGetComplianceInfoResponse();
    message.complianceInfo = (object.complianceInfo !== undefined && object.complianceInfo !== null)
      ? ComplianceInfo.fromPartial(object.complianceInfo)
      : undefined;
    return message;
  },
};

function createBaseQueryAllComplianceInfoRequest(): QueryAllComplianceInfoRequest {
  return { pagination: undefined };
}

export const QueryAllComplianceInfoRequest = {
  encode(message: QueryAllComplianceInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllComplianceInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllComplianceInfoRequest();
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

  fromJSON(object: any): QueryAllComplianceInfoRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllComplianceInfoRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllComplianceInfoRequest>, I>>(
    object: I,
  ): QueryAllComplianceInfoRequest {
    const message = createBaseQueryAllComplianceInfoRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllComplianceInfoResponse(): QueryAllComplianceInfoResponse {
  return { complianceInfo: [], pagination: undefined };
}

export const QueryAllComplianceInfoResponse = {
  encode(message: QueryAllComplianceInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.complianceInfo) {
      ComplianceInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllComplianceInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllComplianceInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.complianceInfo.push(ComplianceInfo.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllComplianceInfoResponse {
    return {
      complianceInfo: Array.isArray(object?.complianceInfo)
        ? object.complianceInfo.map((e: any) => ComplianceInfo.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllComplianceInfoResponse): unknown {
    const obj: any = {};
    if (message.complianceInfo) {
      obj.complianceInfo = message.complianceInfo.map((e) => e ? ComplianceInfo.toJSON(e) : undefined);
    } else {
      obj.complianceInfo = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllComplianceInfoResponse>, I>>(
    object: I,
  ): QueryAllComplianceInfoResponse {
    const message = createBaseQueryAllComplianceInfoResponse();
    message.complianceInfo = object.complianceInfo?.map((e) => ComplianceInfo.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetCertifiedModelRequest(): QueryGetCertifiedModelRequest {
  return { vid: 0, pid: 0, softwareVersion: 0, certificationType: "" };
}

export const QueryGetCertifiedModelRequest = {
  encode(message: QueryGetCertifiedModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion);
    }
    if (message.certificationType !== "") {
      writer.uint32(34).string(message.certificationType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCertifiedModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCertifiedModelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.pid = reader.int32();
          break;
        case 3:
          message.softwareVersion = reader.uint32();
          break;
        case 4:
          message.certificationType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCertifiedModelRequest {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
      certificationType: isSet(object.certificationType) ? String(object.certificationType) : "",
    };
  },

  toJSON(message: QueryGetCertifiedModelRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    message.certificationType !== undefined && (obj.certificationType = message.certificationType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCertifiedModelRequest>, I>>(
    object: I,
  ): QueryGetCertifiedModelRequest {
    const message = createBaseQueryGetCertifiedModelRequest();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    message.certificationType = object.certificationType ?? "";
    return message;
  },
};

function createBaseQueryGetCertifiedModelResponse(): QueryGetCertifiedModelResponse {
  return { certifiedModel: undefined };
}

export const QueryGetCertifiedModelResponse = {
  encode(message: QueryGetCertifiedModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.certifiedModel !== undefined) {
      CertifiedModel.encode(message.certifiedModel, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCertifiedModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCertifiedModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.certifiedModel = CertifiedModel.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCertifiedModelResponse {
    return {
      certifiedModel: isSet(object.certifiedModel) ? CertifiedModel.fromJSON(object.certifiedModel) : undefined,
    };
  },

  toJSON(message: QueryGetCertifiedModelResponse): unknown {
    const obj: any = {};
    message.certifiedModel !== undefined
      && (obj.certifiedModel = message.certifiedModel ? CertifiedModel.toJSON(message.certifiedModel) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCertifiedModelResponse>, I>>(
    object: I,
  ): QueryGetCertifiedModelResponse {
    const message = createBaseQueryGetCertifiedModelResponse();
    message.certifiedModel = (object.certifiedModel !== undefined && object.certifiedModel !== null)
      ? CertifiedModel.fromPartial(object.certifiedModel)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCertifiedModelRequest(): QueryAllCertifiedModelRequest {
  return { pagination: undefined };
}

export const QueryAllCertifiedModelRequest = {
  encode(message: QueryAllCertifiedModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCertifiedModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCertifiedModelRequest();
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

  fromJSON(object: any): QueryAllCertifiedModelRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllCertifiedModelRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCertifiedModelRequest>, I>>(
    object: I,
  ): QueryAllCertifiedModelRequest {
    const message = createBaseQueryAllCertifiedModelRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCertifiedModelResponse(): QueryAllCertifiedModelResponse {
  return { certifiedModel: [], pagination: undefined };
}

export const QueryAllCertifiedModelResponse = {
  encode(message: QueryAllCertifiedModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.certifiedModel) {
      CertifiedModel.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCertifiedModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCertifiedModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.certifiedModel.push(CertifiedModel.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllCertifiedModelResponse {
    return {
      certifiedModel: Array.isArray(object?.certifiedModel)
        ? object.certifiedModel.map((e: any) => CertifiedModel.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllCertifiedModelResponse): unknown {
    const obj: any = {};
    if (message.certifiedModel) {
      obj.certifiedModel = message.certifiedModel.map((e) => e ? CertifiedModel.toJSON(e) : undefined);
    } else {
      obj.certifiedModel = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCertifiedModelResponse>, I>>(
    object: I,
  ): QueryAllCertifiedModelResponse {
    const message = createBaseQueryAllCertifiedModelResponse();
    message.certifiedModel = object.certifiedModel?.map((e) => CertifiedModel.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRevokedModelRequest(): QueryGetRevokedModelRequest {
  return { vid: 0, pid: 0, softwareVersion: 0, certificationType: "" };
}

export const QueryGetRevokedModelRequest = {
  encode(message: QueryGetRevokedModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion);
    }
    if (message.certificationType !== "") {
      writer.uint32(34).string(message.certificationType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedModelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.pid = reader.int32();
          break;
        case 3:
          message.softwareVersion = reader.uint32();
          break;
        case 4:
          message.certificationType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRevokedModelRequest {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
      certificationType: isSet(object.certificationType) ? String(object.certificationType) : "",
    };
  },

  toJSON(message: QueryGetRevokedModelRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    message.certificationType !== undefined && (obj.certificationType = message.certificationType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedModelRequest>, I>>(object: I): QueryGetRevokedModelRequest {
    const message = createBaseQueryGetRevokedModelRequest();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    message.certificationType = object.certificationType ?? "";
    return message;
  },
};

function createBaseQueryGetRevokedModelResponse(): QueryGetRevokedModelResponse {
  return { revokedModel: undefined };
}

export const QueryGetRevokedModelResponse = {
  encode(message: QueryGetRevokedModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.revokedModel !== undefined) {
      RevokedModel.encode(message.revokedModel, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRevokedModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRevokedModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revokedModel = RevokedModel.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRevokedModelResponse {
    return { revokedModel: isSet(object.revokedModel) ? RevokedModel.fromJSON(object.revokedModel) : undefined };
  },

  toJSON(message: QueryGetRevokedModelResponse): unknown {
    const obj: any = {};
    message.revokedModel !== undefined
      && (obj.revokedModel = message.revokedModel ? RevokedModel.toJSON(message.revokedModel) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRevokedModelResponse>, I>>(object: I): QueryGetRevokedModelResponse {
    const message = createBaseQueryGetRevokedModelResponse();
    message.revokedModel = (object.revokedModel !== undefined && object.revokedModel !== null)
      ? RevokedModel.fromPartial(object.revokedModel)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRevokedModelRequest(): QueryAllRevokedModelRequest {
  return { pagination: undefined };
}

export const QueryAllRevokedModelRequest = {
  encode(message: QueryAllRevokedModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRevokedModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRevokedModelRequest();
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

  fromJSON(object: any): QueryAllRevokedModelRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRevokedModelRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRevokedModelRequest>, I>>(object: I): QueryAllRevokedModelRequest {
    const message = createBaseQueryAllRevokedModelRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRevokedModelResponse(): QueryAllRevokedModelResponse {
  return { revokedModel: [], pagination: undefined };
}

export const QueryAllRevokedModelResponse = {
  encode(message: QueryAllRevokedModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.revokedModel) {
      RevokedModel.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRevokedModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRevokedModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.revokedModel.push(RevokedModel.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllRevokedModelResponse {
    return {
      revokedModel: Array.isArray(object?.revokedModel)
        ? object.revokedModel.map((e: any) => RevokedModel.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRevokedModelResponse): unknown {
    const obj: any = {};
    if (message.revokedModel) {
      obj.revokedModel = message.revokedModel.map((e) => e ? RevokedModel.toJSON(e) : undefined);
    } else {
      obj.revokedModel = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRevokedModelResponse>, I>>(object: I): QueryAllRevokedModelResponse {
    const message = createBaseQueryAllRevokedModelResponse();
    message.revokedModel = object.revokedModel?.map((e) => RevokedModel.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetProvisionalModelRequest(): QueryGetProvisionalModelRequest {
  return { vid: 0, pid: 0, softwareVersion: 0, certificationType: "" };
}

export const QueryGetProvisionalModelRequest = {
  encode(message: QueryGetProvisionalModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion);
    }
    if (message.certificationType !== "") {
      writer.uint32(34).string(message.certificationType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProvisionalModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProvisionalModelRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.pid = reader.int32();
          break;
        case 3:
          message.softwareVersion = reader.uint32();
          break;
        case 4:
          message.certificationType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProvisionalModelRequest {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      softwareVersion: isSet(object.softwareVersion) ? Number(object.softwareVersion) : 0,
      certificationType: isSet(object.certificationType) ? String(object.certificationType) : "",
    };
  },

  toJSON(message: QueryGetProvisionalModelRequest): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.softwareVersion !== undefined && (obj.softwareVersion = Math.round(message.softwareVersion));
    message.certificationType !== undefined && (obj.certificationType = message.certificationType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProvisionalModelRequest>, I>>(
    object: I,
  ): QueryGetProvisionalModelRequest {
    const message = createBaseQueryGetProvisionalModelRequest();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.softwareVersion = object.softwareVersion ?? 0;
    message.certificationType = object.certificationType ?? "";
    return message;
  },
};

function createBaseQueryGetProvisionalModelResponse(): QueryGetProvisionalModelResponse {
  return { provisionalModel: undefined };
}

export const QueryGetProvisionalModelResponse = {
  encode(message: QueryGetProvisionalModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.provisionalModel !== undefined) {
      ProvisionalModel.encode(message.provisionalModel, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProvisionalModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProvisionalModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.provisionalModel = ProvisionalModel.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProvisionalModelResponse {
    return {
      provisionalModel: isSet(object.provisionalModel) ? ProvisionalModel.fromJSON(object.provisionalModel) : undefined,
    };
  },

  toJSON(message: QueryGetProvisionalModelResponse): unknown {
    const obj: any = {};
    message.provisionalModel !== undefined && (obj.provisionalModel = message.provisionalModel
      ? ProvisionalModel.toJSON(message.provisionalModel)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProvisionalModelResponse>, I>>(
    object: I,
  ): QueryGetProvisionalModelResponse {
    const message = createBaseQueryGetProvisionalModelResponse();
    message.provisionalModel = (object.provisionalModel !== undefined && object.provisionalModel !== null)
      ? ProvisionalModel.fromPartial(object.provisionalModel)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProvisionalModelRequest(): QueryAllProvisionalModelRequest {
  return { pagination: undefined };
}

export const QueryAllProvisionalModelRequest = {
  encode(message: QueryAllProvisionalModelRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProvisionalModelRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProvisionalModelRequest();
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

  fromJSON(object: any): QueryAllProvisionalModelRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllProvisionalModelRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProvisionalModelRequest>, I>>(
    object: I,
  ): QueryAllProvisionalModelRequest {
    const message = createBaseQueryAllProvisionalModelRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProvisionalModelResponse(): QueryAllProvisionalModelResponse {
  return { provisionalModel: [], pagination: undefined };
}

export const QueryAllProvisionalModelResponse = {
  encode(message: QueryAllProvisionalModelResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.provisionalModel) {
      ProvisionalModel.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProvisionalModelResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProvisionalModelResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.provisionalModel.push(ProvisionalModel.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllProvisionalModelResponse {
    return {
      provisionalModel: Array.isArray(object?.provisionalModel)
        ? object.provisionalModel.map((e: any) => ProvisionalModel.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllProvisionalModelResponse): unknown {
    const obj: any = {};
    if (message.provisionalModel) {
      obj.provisionalModel = message.provisionalModel.map((e) => e ? ProvisionalModel.toJSON(e) : undefined);
    } else {
      obj.provisionalModel = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProvisionalModelResponse>, I>>(
    object: I,
  ): QueryAllProvisionalModelResponse {
    const message = createBaseQueryAllProvisionalModelResponse();
    message.provisionalModel = object.provisionalModel?.map((e) => ProvisionalModel.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDeviceSoftwareComplianceRequest(): QueryGetDeviceSoftwareComplianceRequest {
  return { cDCertificateId: "" };
}

export const QueryGetDeviceSoftwareComplianceRequest = {
  encode(message: QueryGetDeviceSoftwareComplianceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cDCertificateId !== "") {
      writer.uint32(10).string(message.cDCertificateId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDeviceSoftwareComplianceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDeviceSoftwareComplianceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cDCertificateId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDeviceSoftwareComplianceRequest {
    return { cDCertificateId: isSet(object.cDCertificateId) ? String(object.cDCertificateId) : "" };
  },

  toJSON(message: QueryGetDeviceSoftwareComplianceRequest): unknown {
    const obj: any = {};
    message.cDCertificateId !== undefined && (obj.cDCertificateId = message.cDCertificateId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDeviceSoftwareComplianceRequest>, I>>(
    object: I,
  ): QueryGetDeviceSoftwareComplianceRequest {
    const message = createBaseQueryGetDeviceSoftwareComplianceRequest();
    message.cDCertificateId = object.cDCertificateId ?? "";
    return message;
  },
};

function createBaseQueryGetDeviceSoftwareComplianceResponse(): QueryGetDeviceSoftwareComplianceResponse {
  return { deviceSoftwareCompliance: undefined };
}

export const QueryGetDeviceSoftwareComplianceResponse = {
  encode(message: QueryGetDeviceSoftwareComplianceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceSoftwareCompliance !== undefined) {
      DeviceSoftwareCompliance.encode(message.deviceSoftwareCompliance, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDeviceSoftwareComplianceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDeviceSoftwareComplianceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceSoftwareCompliance = DeviceSoftwareCompliance.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDeviceSoftwareComplianceResponse {
    return {
      deviceSoftwareCompliance: isSet(object.deviceSoftwareCompliance)
        ? DeviceSoftwareCompliance.fromJSON(object.deviceSoftwareCompliance)
        : undefined,
    };
  },

  toJSON(message: QueryGetDeviceSoftwareComplianceResponse): unknown {
    const obj: any = {};
    message.deviceSoftwareCompliance !== undefined && (obj.deviceSoftwareCompliance = message.deviceSoftwareCompliance
      ? DeviceSoftwareCompliance.toJSON(message.deviceSoftwareCompliance)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDeviceSoftwareComplianceResponse>, I>>(
    object: I,
  ): QueryGetDeviceSoftwareComplianceResponse {
    const message = createBaseQueryGetDeviceSoftwareComplianceResponse();
    message.deviceSoftwareCompliance =
      (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null)
        ? DeviceSoftwareCompliance.fromPartial(object.deviceSoftwareCompliance)
        : undefined;
    return message;
  },
};

function createBaseQueryAllDeviceSoftwareComplianceRequest(): QueryAllDeviceSoftwareComplianceRequest {
  return { pagination: undefined };
}

export const QueryAllDeviceSoftwareComplianceRequest = {
  encode(message: QueryAllDeviceSoftwareComplianceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDeviceSoftwareComplianceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDeviceSoftwareComplianceRequest();
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

  fromJSON(object: any): QueryAllDeviceSoftwareComplianceRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDeviceSoftwareComplianceRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDeviceSoftwareComplianceRequest>, I>>(
    object: I,
  ): QueryAllDeviceSoftwareComplianceRequest {
    const message = createBaseQueryAllDeviceSoftwareComplianceRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDeviceSoftwareComplianceResponse(): QueryAllDeviceSoftwareComplianceResponse {
  return { deviceSoftwareCompliance: [], pagination: undefined };
}

export const QueryAllDeviceSoftwareComplianceResponse = {
  encode(message: QueryAllDeviceSoftwareComplianceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.deviceSoftwareCompliance) {
      DeviceSoftwareCompliance.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDeviceSoftwareComplianceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDeviceSoftwareComplianceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceSoftwareCompliance.push(DeviceSoftwareCompliance.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllDeviceSoftwareComplianceResponse {
    return {
      deviceSoftwareCompliance: Array.isArray(object?.deviceSoftwareCompliance)
        ? object.deviceSoftwareCompliance.map((e: any) => DeviceSoftwareCompliance.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDeviceSoftwareComplianceResponse): unknown {
    const obj: any = {};
    if (message.deviceSoftwareCompliance) {
      obj.deviceSoftwareCompliance = message.deviceSoftwareCompliance.map((e) =>
        e ? DeviceSoftwareCompliance.toJSON(e) : undefined
      );
    } else {
      obj.deviceSoftwareCompliance = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDeviceSoftwareComplianceResponse>, I>>(
    object: I,
  ): QueryAllDeviceSoftwareComplianceResponse {
    const message = createBaseQueryAllDeviceSoftwareComplianceResponse();
    message.deviceSoftwareCompliance =
      object.deviceSoftwareCompliance?.map((e) => DeviceSoftwareCompliance.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a ComplianceInfo by index. */
  ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse>;
  /** Queries a list of ComplianceInfo items. */
  ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse>;
  /** Queries a CertifiedModel by index. */
  CertifiedModel(request: QueryGetCertifiedModelRequest): Promise<QueryGetCertifiedModelResponse>;
  /** Queries a list of CertifiedModel items. */
  CertifiedModelAll(request: QueryAllCertifiedModelRequest): Promise<QueryAllCertifiedModelResponse>;
  /** Queries a RevokedModel by index. */
  RevokedModel(request: QueryGetRevokedModelRequest): Promise<QueryGetRevokedModelResponse>;
  /** Queries a list of RevokedModel items. */
  RevokedModelAll(request: QueryAllRevokedModelRequest): Promise<QueryAllRevokedModelResponse>;
  /** Queries a ProvisionalModel by index. */
  ProvisionalModel(request: QueryGetProvisionalModelRequest): Promise<QueryGetProvisionalModelResponse>;
  /** Queries a list of ProvisionalModel items. */
  ProvisionalModelAll(request: QueryAllProvisionalModelRequest): Promise<QueryAllProvisionalModelResponse>;
  /** Queries a DeviceSoftwareCompliance by index. */
  DeviceSoftwareCompliance(
    request: QueryGetDeviceSoftwareComplianceRequest,
  ): Promise<QueryGetDeviceSoftwareComplianceResponse>;
  /** Queries a list of DeviceSoftwareCompliance items. */
  DeviceSoftwareComplianceAll(
    request: QueryAllDeviceSoftwareComplianceRequest,
  ): Promise<QueryAllDeviceSoftwareComplianceResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ComplianceInfo = this.ComplianceInfo.bind(this);
    this.ComplianceInfoAll = this.ComplianceInfoAll.bind(this);
    this.CertifiedModel = this.CertifiedModel.bind(this);
    this.CertifiedModelAll = this.CertifiedModelAll.bind(this);
    this.RevokedModel = this.RevokedModel.bind(this);
    this.RevokedModelAll = this.RevokedModelAll.bind(this);
    this.ProvisionalModel = this.ProvisionalModel.bind(this);
    this.ProvisionalModelAll = this.ProvisionalModelAll.bind(this);
    this.DeviceSoftwareCompliance = this.DeviceSoftwareCompliance.bind(this);
    this.DeviceSoftwareComplianceAll = this.DeviceSoftwareComplianceAll.bind(this);
  }
  ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse> {
    const data = QueryGetComplianceInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "ComplianceInfo",
      data,
    );
    return promise.then((data) => QueryGetComplianceInfoResponse.decode(new _m0.Reader(data)));
  }

  ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse> {
    const data = QueryAllComplianceInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "ComplianceInfoAll",
      data,
    );
    return promise.then((data) => QueryAllComplianceInfoResponse.decode(new _m0.Reader(data)));
  }

  CertifiedModel(request: QueryGetCertifiedModelRequest): Promise<QueryGetCertifiedModelResponse> {
    const data = QueryGetCertifiedModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "CertifiedModel",
      data,
    );
    return promise.then((data) => QueryGetCertifiedModelResponse.decode(new _m0.Reader(data)));
  }

  CertifiedModelAll(request: QueryAllCertifiedModelRequest): Promise<QueryAllCertifiedModelResponse> {
    const data = QueryAllCertifiedModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "CertifiedModelAll",
      data,
    );
    return promise.then((data) => QueryAllCertifiedModelResponse.decode(new _m0.Reader(data)));
  }

  RevokedModel(request: QueryGetRevokedModelRequest): Promise<QueryGetRevokedModelResponse> {
    const data = QueryGetRevokedModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "RevokedModel",
      data,
    );
    return promise.then((data) => QueryGetRevokedModelResponse.decode(new _m0.Reader(data)));
  }

  RevokedModelAll(request: QueryAllRevokedModelRequest): Promise<QueryAllRevokedModelResponse> {
    const data = QueryAllRevokedModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "RevokedModelAll",
      data,
    );
    return promise.then((data) => QueryAllRevokedModelResponse.decode(new _m0.Reader(data)));
  }

  ProvisionalModel(request: QueryGetProvisionalModelRequest): Promise<QueryGetProvisionalModelResponse> {
    const data = QueryGetProvisionalModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "ProvisionalModel",
      data,
    );
    return promise.then((data) => QueryGetProvisionalModelResponse.decode(new _m0.Reader(data)));
  }

  ProvisionalModelAll(request: QueryAllProvisionalModelRequest): Promise<QueryAllProvisionalModelResponse> {
    const data = QueryAllProvisionalModelRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "ProvisionalModelAll",
      data,
    );
    return promise.then((data) => QueryAllProvisionalModelResponse.decode(new _m0.Reader(data)));
  }

  DeviceSoftwareCompliance(
    request: QueryGetDeviceSoftwareComplianceRequest,
  ): Promise<QueryGetDeviceSoftwareComplianceResponse> {
    const data = QueryGetDeviceSoftwareComplianceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "DeviceSoftwareCompliance",
      data,
    );
    return promise.then((data) => QueryGetDeviceSoftwareComplianceResponse.decode(new _m0.Reader(data)));
  }

  DeviceSoftwareComplianceAll(
    request: QueryAllDeviceSoftwareComplianceRequest,
  ): Promise<QueryAllDeviceSoftwareComplianceResponse> {
    const data = QueryAllDeviceSoftwareComplianceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.compliance.Query",
      "DeviceSoftwareComplianceAll",
      data,
    );
    return promise.then((data) => QueryAllDeviceSoftwareComplianceResponse.decode(new _m0.Reader(data)));
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
