/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ComplianceInfo } from '../compliance/compliance_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { CertifiedModel } from '../compliance/certified_model';
import { RevokedModel } from '../compliance/revoked_model';
import { ProvisionalModel } from '../compliance/provisional_model';
import { DeviceSoftwareCompliance } from '../compliance/device_software_compliance';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseQueryGetComplianceInfoRequest = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' };
export const QueryGetComplianceInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint32(message.softwareVersion);
        }
        if (message.certificationType !== '') {
            writer.uint32(34).string(message.certificationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetComplianceInfoRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetComplianceInfoRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetComplianceInfoRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        return message;
    }
};
const baseQueryGetComplianceInfoResponse = {};
export const QueryGetComplianceInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.complianceInfo !== undefined) {
            ComplianceInfo.encode(message.complianceInfo, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetComplianceInfoResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetComplianceInfoResponse };
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            message.complianceInfo = ComplianceInfo.fromJSON(object.complianceInfo);
        }
        else {
            message.complianceInfo = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.complianceInfo !== undefined && (obj.complianceInfo = message.complianceInfo ? ComplianceInfo.toJSON(message.complianceInfo) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetComplianceInfoResponse };
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            message.complianceInfo = ComplianceInfo.fromPartial(object.complianceInfo);
        }
        else {
            message.complianceInfo = undefined;
        }
        return message;
    }
};
const baseQueryAllComplianceInfoRequest = {};
export const QueryAllComplianceInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllComplianceInfoRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllComplianceInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllComplianceInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllComplianceInfoResponse = {};
export const QueryAllComplianceInfoResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.complianceInfo) {
            ComplianceInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllComplianceInfoResponse };
        message.complianceInfo = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllComplianceInfoResponse };
        message.complianceInfo = [];
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(ComplianceInfo.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.complianceInfo) {
            obj.complianceInfo = message.complianceInfo.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined));
        }
        else {
            obj.complianceInfo = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllComplianceInfoResponse };
        message.complianceInfo = [];
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(ComplianceInfo.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetCertifiedModelRequest = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' };
export const QueryGetCertifiedModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint32(message.softwareVersion);
        }
        if (message.certificationType !== '') {
            writer.uint32(34).string(message.certificationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCertifiedModelRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetCertifiedModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCertifiedModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        return message;
    }
};
const baseQueryGetCertifiedModelResponse = {};
export const QueryGetCertifiedModelResponse = {
    encode(message, writer = Writer.create()) {
        if (message.certifiedModel !== undefined) {
            CertifiedModel.encode(message.certifiedModel, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetCertifiedModelResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetCertifiedModelResponse };
        if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
            message.certifiedModel = CertifiedModel.fromJSON(object.certifiedModel);
        }
        else {
            message.certifiedModel = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.certifiedModel !== undefined && (obj.certifiedModel = message.certifiedModel ? CertifiedModel.toJSON(message.certifiedModel) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetCertifiedModelResponse };
        if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
            message.certifiedModel = CertifiedModel.fromPartial(object.certifiedModel);
        }
        else {
            message.certifiedModel = undefined;
        }
        return message;
    }
};
const baseQueryAllCertifiedModelRequest = {};
export const QueryAllCertifiedModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllCertifiedModelRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllCertifiedModelRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllCertifiedModelRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllCertifiedModelResponse = {};
export const QueryAllCertifiedModelResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.certifiedModel) {
            CertifiedModel.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllCertifiedModelResponse };
        message.certifiedModel = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllCertifiedModelResponse };
        message.certifiedModel = [];
        if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
            for (const e of object.certifiedModel) {
                message.certifiedModel.push(CertifiedModel.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.certifiedModel) {
            obj.certifiedModel = message.certifiedModel.map((e) => (e ? CertifiedModel.toJSON(e) : undefined));
        }
        else {
            obj.certifiedModel = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllCertifiedModelResponse };
        message.certifiedModel = [];
        if (object.certifiedModel !== undefined && object.certifiedModel !== null) {
            for (const e of object.certifiedModel) {
                message.certifiedModel.push(CertifiedModel.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetRevokedModelRequest = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' };
export const QueryGetRevokedModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint32(message.softwareVersion);
        }
        if (message.certificationType !== '') {
            writer.uint32(34).string(message.certificationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRevokedModelRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetRevokedModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRevokedModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        return message;
    }
};
const baseQueryGetRevokedModelResponse = {};
export const QueryGetRevokedModelResponse = {
    encode(message, writer = Writer.create()) {
        if (message.revokedModel !== undefined) {
            RevokedModel.encode(message.revokedModel, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRevokedModelResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetRevokedModelResponse };
        if (object.revokedModel !== undefined && object.revokedModel !== null) {
            message.revokedModel = RevokedModel.fromJSON(object.revokedModel);
        }
        else {
            message.revokedModel = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.revokedModel !== undefined && (obj.revokedModel = message.revokedModel ? RevokedModel.toJSON(message.revokedModel) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRevokedModelResponse };
        if (object.revokedModel !== undefined && object.revokedModel !== null) {
            message.revokedModel = RevokedModel.fromPartial(object.revokedModel);
        }
        else {
            message.revokedModel = undefined;
        }
        return message;
    }
};
const baseQueryAllRevokedModelRequest = {};
export const QueryAllRevokedModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRevokedModelRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllRevokedModelRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRevokedModelRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllRevokedModelResponse = {};
export const QueryAllRevokedModelResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.revokedModel) {
            RevokedModel.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRevokedModelResponse };
        message.revokedModel = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllRevokedModelResponse };
        message.revokedModel = [];
        if (object.revokedModel !== undefined && object.revokedModel !== null) {
            for (const e of object.revokedModel) {
                message.revokedModel.push(RevokedModel.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.revokedModel) {
            obj.revokedModel = message.revokedModel.map((e) => (e ? RevokedModel.toJSON(e) : undefined));
        }
        else {
            obj.revokedModel = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRevokedModelResponse };
        message.revokedModel = [];
        if (object.revokedModel !== undefined && object.revokedModel !== null) {
            for (const e of object.revokedModel) {
                message.revokedModel.push(RevokedModel.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetProvisionalModelRequest = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' };
export const QueryGetProvisionalModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint32(message.softwareVersion);
        }
        if (message.certificationType !== '') {
            writer.uint32(34).string(message.certificationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProvisionalModelRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetProvisionalModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProvisionalModelRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        return message;
    }
};
const baseQueryGetProvisionalModelResponse = {};
export const QueryGetProvisionalModelResponse = {
    encode(message, writer = Writer.create()) {
        if (message.provisionalModel !== undefined) {
            ProvisionalModel.encode(message.provisionalModel, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProvisionalModelResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetProvisionalModelResponse };
        if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
            message.provisionalModel = ProvisionalModel.fromJSON(object.provisionalModel);
        }
        else {
            message.provisionalModel = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.provisionalModel !== undefined && (obj.provisionalModel = message.provisionalModel ? ProvisionalModel.toJSON(message.provisionalModel) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProvisionalModelResponse };
        if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
            message.provisionalModel = ProvisionalModel.fromPartial(object.provisionalModel);
        }
        else {
            message.provisionalModel = undefined;
        }
        return message;
    }
};
const baseQueryAllProvisionalModelRequest = {};
export const QueryAllProvisionalModelRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProvisionalModelRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllProvisionalModelRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllProvisionalModelRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllProvisionalModelResponse = {};
export const QueryAllProvisionalModelResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.provisionalModel) {
            ProvisionalModel.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProvisionalModelResponse };
        message.provisionalModel = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllProvisionalModelResponse };
        message.provisionalModel = [];
        if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
            for (const e of object.provisionalModel) {
                message.provisionalModel.push(ProvisionalModel.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.provisionalModel) {
            obj.provisionalModel = message.provisionalModel.map((e) => (e ? ProvisionalModel.toJSON(e) : undefined));
        }
        else {
            obj.provisionalModel = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllProvisionalModelResponse };
        message.provisionalModel = [];
        if (object.provisionalModel !== undefined && object.provisionalModel !== null) {
            for (const e of object.provisionalModel) {
                message.provisionalModel.push(ProvisionalModel.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetDeviceSoftwareComplianceRequest = { cDCertificateId: '' };
export const QueryGetDeviceSoftwareComplianceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.cDCertificateId !== '') {
            writer.uint32(10).string(message.cDCertificateId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetDeviceSoftwareComplianceRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetDeviceSoftwareComplianceRequest };
        if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
            message.cDCertificateId = String(object.cDCertificateId);
        }
        else {
            message.cDCertificateId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cDCertificateId !== undefined && (obj.cDCertificateId = message.cDCertificateId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetDeviceSoftwareComplianceRequest };
        if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
            message.cDCertificateId = object.cDCertificateId;
        }
        else {
            message.cDCertificateId = '';
        }
        return message;
    }
};
const baseQueryGetDeviceSoftwareComplianceResponse = {};
export const QueryGetDeviceSoftwareComplianceResponse = {
    encode(message, writer = Writer.create()) {
        if (message.deviceSoftwareCompliance !== undefined) {
            DeviceSoftwareCompliance.encode(message.deviceSoftwareCompliance, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetDeviceSoftwareComplianceResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetDeviceSoftwareComplianceResponse };
        if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
            message.deviceSoftwareCompliance = DeviceSoftwareCompliance.fromJSON(object.deviceSoftwareCompliance);
        }
        else {
            message.deviceSoftwareCompliance = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.deviceSoftwareCompliance !== undefined &&
            (obj.deviceSoftwareCompliance = message.deviceSoftwareCompliance ? DeviceSoftwareCompliance.toJSON(message.deviceSoftwareCompliance) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetDeviceSoftwareComplianceResponse };
        if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
            message.deviceSoftwareCompliance = DeviceSoftwareCompliance.fromPartial(object.deviceSoftwareCompliance);
        }
        else {
            message.deviceSoftwareCompliance = undefined;
        }
        return message;
    }
};
const baseQueryAllDeviceSoftwareComplianceRequest = {};
export const QueryAllDeviceSoftwareComplianceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllDeviceSoftwareComplianceRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllDeviceSoftwareComplianceRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllDeviceSoftwareComplianceRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllDeviceSoftwareComplianceResponse = {};
export const QueryAllDeviceSoftwareComplianceResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.deviceSoftwareCompliance) {
            DeviceSoftwareCompliance.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllDeviceSoftwareComplianceResponse };
        message.deviceSoftwareCompliance = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllDeviceSoftwareComplianceResponse };
        message.deviceSoftwareCompliance = [];
        if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
            for (const e of object.deviceSoftwareCompliance) {
                message.deviceSoftwareCompliance.push(DeviceSoftwareCompliance.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.deviceSoftwareCompliance) {
            obj.deviceSoftwareCompliance = message.deviceSoftwareCompliance.map((e) => (e ? DeviceSoftwareCompliance.toJSON(e) : undefined));
        }
        else {
            obj.deviceSoftwareCompliance = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllDeviceSoftwareComplianceResponse };
        message.deviceSoftwareCompliance = [];
        if (object.deviceSoftwareCompliance !== undefined && object.deviceSoftwareCompliance !== null) {
            for (const e of object.deviceSoftwareCompliance) {
                message.deviceSoftwareCompliance.push(DeviceSoftwareCompliance.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ComplianceInfo(request) {
        const data = QueryGetComplianceInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfo', data);
        return promise.then((data) => QueryGetComplianceInfoResponse.decode(new Reader(data)));
    }
    ComplianceInfoAll(request) {
        const data = QueryAllComplianceInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfoAll', data);
        return promise.then((data) => QueryAllComplianceInfoResponse.decode(new Reader(data)));
    }
    CertifiedModel(request) {
        const data = QueryGetCertifiedModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'CertifiedModel', data);
        return promise.then((data) => QueryGetCertifiedModelResponse.decode(new Reader(data)));
    }
    CertifiedModelAll(request) {
        const data = QueryAllCertifiedModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'CertifiedModelAll', data);
        return promise.then((data) => QueryAllCertifiedModelResponse.decode(new Reader(data)));
    }
    RevokedModel(request) {
        const data = QueryGetRevokedModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'RevokedModel', data);
        return promise.then((data) => QueryGetRevokedModelResponse.decode(new Reader(data)));
    }
    RevokedModelAll(request) {
        const data = QueryAllRevokedModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'RevokedModelAll', data);
        return promise.then((data) => QueryAllRevokedModelResponse.decode(new Reader(data)));
    }
    ProvisionalModel(request) {
        const data = QueryGetProvisionalModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ProvisionalModel', data);
        return promise.then((data) => QueryGetProvisionalModelResponse.decode(new Reader(data)));
    }
    ProvisionalModelAll(request) {
        const data = QueryAllProvisionalModelRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ProvisionalModelAll', data);
        return promise.then((data) => QueryAllProvisionalModelResponse.decode(new Reader(data)));
    }
    DeviceSoftwareCompliance(request) {
        const data = QueryGetDeviceSoftwareComplianceRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'DeviceSoftwareCompliance', data);
        return promise.then((data) => QueryGetDeviceSoftwareComplianceResponse.decode(new Reader(data)));
    }
    DeviceSoftwareComplianceAll(request) {
        const data = QueryAllDeviceSoftwareComplianceRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'DeviceSoftwareComplianceAll', data);
        return promise.then((data) => QueryAllDeviceSoftwareComplianceResponse.decode(new Reader(data)));
    }
}
