/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ApprovedCertificates } from '../pki/approved_certificates';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { ProposedCertificate } from '../pki/proposed_certificate';
import { ChildCertificates } from '../pki/child_certificates';
import { ProposedCertificateRevocation } from '../pki/proposed_certificate_revocation';
import { RevokedCertificates } from '../pki/revoked_certificates';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseQueryGetApprovedCertificatesRequest = { subject: '', subjectKeyId: '' };
export const QueryGetApprovedCertificatesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.subject !== '') {
            writer.uint32(10).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(18).string(message.subjectKeyId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetApprovedCertificatesRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetApprovedCertificatesRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetApprovedCertificatesRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    }
};
const baseQueryGetApprovedCertificatesResponse = {};
export const QueryGetApprovedCertificatesResponse = {
    encode(message, writer = Writer.create()) {
        if (message.approvedCertificates !== undefined) {
            ApprovedCertificates.encode(message.approvedCertificates, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetApprovedCertificatesResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetApprovedCertificatesResponse };
        if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
            message.approvedCertificates = ApprovedCertificates.fromJSON(object.approvedCertificates);
        }
        else {
            message.approvedCertificates = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.approvedCertificates !== undefined &&
            (obj.approvedCertificates = message.approvedCertificates ? ApprovedCertificates.toJSON(message.approvedCertificates) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetApprovedCertificatesResponse };
        if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
            message.approvedCertificates = ApprovedCertificates.fromPartial(object.approvedCertificates);
        }
        else {
            message.approvedCertificates = undefined;
        }
        return message;
    }
};
const baseQueryAllApprovedCertificatesRequest = {};
export const QueryAllApprovedCertificatesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllApprovedCertificatesRequest };
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
        const message = { ...baseQueryAllApprovedCertificatesRequest };
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
        const message = { ...baseQueryAllApprovedCertificatesRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllApprovedCertificatesResponse = {};
export const QueryAllApprovedCertificatesResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.approvedCertificates) {
            ApprovedCertificates.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllApprovedCertificatesResponse };
        message.approvedCertificates = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllApprovedCertificatesResponse };
        message.approvedCertificates = [];
        if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
            for (const e of object.approvedCertificates) {
                message.approvedCertificates.push(ApprovedCertificates.fromJSON(e));
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
        if (message.approvedCertificates) {
            obj.approvedCertificates = message.approvedCertificates.map((e) => (e ? ApprovedCertificates.toJSON(e) : undefined));
        }
        else {
            obj.approvedCertificates = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllApprovedCertificatesResponse };
        message.approvedCertificates = [];
        if (object.approvedCertificates !== undefined && object.approvedCertificates !== null) {
            for (const e of object.approvedCertificates) {
                message.approvedCertificates.push(ApprovedCertificates.fromPartial(e));
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
const baseQueryGetProposedCertificateRequest = { subject: '', subjectKeyId: '' };
export const QueryGetProposedCertificateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.subject !== '') {
            writer.uint32(10).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(18).string(message.subjectKeyId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedCertificateRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetProposedCertificateRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProposedCertificateRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    }
};
const baseQueryGetProposedCertificateResponse = {};
export const QueryGetProposedCertificateResponse = {
    encode(message, writer = Writer.create()) {
        if (message.proposedCertificate !== undefined) {
            ProposedCertificate.encode(message.proposedCertificate, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedCertificateResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetProposedCertificateResponse };
        if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
            message.proposedCertificate = ProposedCertificate.fromJSON(object.proposedCertificate);
        }
        else {
            message.proposedCertificate = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposedCertificate !== undefined &&
            (obj.proposedCertificate = message.proposedCertificate ? ProposedCertificate.toJSON(message.proposedCertificate) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProposedCertificateResponse };
        if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
            message.proposedCertificate = ProposedCertificate.fromPartial(object.proposedCertificate);
        }
        else {
            message.proposedCertificate = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedCertificateRequest = {};
export const QueryAllProposedCertificateRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedCertificateRequest };
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
        const message = { ...baseQueryAllProposedCertificateRequest };
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
        const message = { ...baseQueryAllProposedCertificateRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedCertificateResponse = {};
export const QueryAllProposedCertificateResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.proposedCertificate) {
            ProposedCertificate.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedCertificateResponse };
        message.proposedCertificate = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllProposedCertificateResponse };
        message.proposedCertificate = [];
        if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
            for (const e of object.proposedCertificate) {
                message.proposedCertificate.push(ProposedCertificate.fromJSON(e));
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
        if (message.proposedCertificate) {
            obj.proposedCertificate = message.proposedCertificate.map((e) => (e ? ProposedCertificate.toJSON(e) : undefined));
        }
        else {
            obj.proposedCertificate = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllProposedCertificateResponse };
        message.proposedCertificate = [];
        if (object.proposedCertificate !== undefined && object.proposedCertificate !== null) {
            for (const e of object.proposedCertificate) {
                message.proposedCertificate.push(ProposedCertificate.fromPartial(e));
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
const baseQueryGetChildCertificatesRequest = { issuer: '', authorityKeyId: '' };
export const QueryGetChildCertificatesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.issuer !== '') {
            writer.uint32(10).string(message.issuer);
        }
        if (message.authorityKeyId !== '') {
            writer.uint32(18).string(message.authorityKeyId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChildCertificatesRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetChildCertificatesRequest };
        if (object.issuer !== undefined && object.issuer !== null) {
            message.issuer = String(object.issuer);
        }
        else {
            message.issuer = '';
        }
        if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
            message.authorityKeyId = String(object.authorityKeyId);
        }
        else {
            message.authorityKeyId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.issuer !== undefined && (obj.issuer = message.issuer);
        message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChildCertificatesRequest };
        if (object.issuer !== undefined && object.issuer !== null) {
            message.issuer = object.issuer;
        }
        else {
            message.issuer = '';
        }
        if (object.authorityKeyId !== undefined && object.authorityKeyId !== null) {
            message.authorityKeyId = object.authorityKeyId;
        }
        else {
            message.authorityKeyId = '';
        }
        return message;
    }
};
const baseQueryGetChildCertificatesResponse = {};
export const QueryGetChildCertificatesResponse = {
    encode(message, writer = Writer.create()) {
        if (message.childCertificates !== undefined) {
            ChildCertificates.encode(message.childCertificates, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetChildCertificatesResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetChildCertificatesResponse };
        if (object.childCertificates !== undefined && object.childCertificates !== null) {
            message.childCertificates = ChildCertificates.fromJSON(object.childCertificates);
        }
        else {
            message.childCertificates = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.childCertificates !== undefined &&
            (obj.childCertificates = message.childCertificates ? ChildCertificates.toJSON(message.childCertificates) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetChildCertificatesResponse };
        if (object.childCertificates !== undefined && object.childCertificates !== null) {
            message.childCertificates = ChildCertificates.fromPartial(object.childCertificates);
        }
        else {
            message.childCertificates = undefined;
        }
        return message;
    }
};
const baseQueryAllChildCertificatesRequest = {};
export const QueryAllChildCertificatesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChildCertificatesRequest };
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
        const message = { ...baseQueryAllChildCertificatesRequest };
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
        const message = { ...baseQueryAllChildCertificatesRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllChildCertificatesResponse = {};
export const QueryAllChildCertificatesResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.childCertificates) {
            ChildCertificates.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllChildCertificatesResponse };
        message.childCertificates = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.childCertificates.push(ChildCertificates.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllChildCertificatesResponse };
        message.childCertificates = [];
        if (object.childCertificates !== undefined && object.childCertificates !== null) {
            for (const e of object.childCertificates) {
                message.childCertificates.push(ChildCertificates.fromJSON(e));
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
        if (message.childCertificates) {
            obj.childCertificates = message.childCertificates.map((e) => (e ? ChildCertificates.toJSON(e) : undefined));
        }
        else {
            obj.childCertificates = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllChildCertificatesResponse };
        message.childCertificates = [];
        if (object.childCertificates !== undefined && object.childCertificates !== null) {
            for (const e of object.childCertificates) {
                message.childCertificates.push(ChildCertificates.fromPartial(e));
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
const baseQueryGetProposedCertificateRevocationRequest = { subject: '', subjectKeyId: '' };
export const QueryGetProposedCertificateRevocationRequest = {
    encode(message, writer = Writer.create()) {
        if (message.subject !== '') {
            writer.uint32(10).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(18).string(message.subjectKeyId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedCertificateRevocationRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetProposedCertificateRevocationRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProposedCertificateRevocationRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    }
};
const baseQueryGetProposedCertificateRevocationResponse = {};
export const QueryGetProposedCertificateRevocationResponse = {
    encode(message, writer = Writer.create()) {
        if (message.proposedCertificateRevocation !== undefined) {
            ProposedCertificateRevocation.encode(message.proposedCertificateRevocation, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetProposedCertificateRevocationResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetProposedCertificateRevocationResponse };
        if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
            message.proposedCertificateRevocation = ProposedCertificateRevocation.fromJSON(object.proposedCertificateRevocation);
        }
        else {
            message.proposedCertificateRevocation = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.proposedCertificateRevocation !== undefined &&
            (obj.proposedCertificateRevocation = message.proposedCertificateRevocation
                ? ProposedCertificateRevocation.toJSON(message.proposedCertificateRevocation)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetProposedCertificateRevocationResponse };
        if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
            message.proposedCertificateRevocation = ProposedCertificateRevocation.fromPartial(object.proposedCertificateRevocation);
        }
        else {
            message.proposedCertificateRevocation = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedCertificateRevocationRequest = {};
export const QueryAllProposedCertificateRevocationRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedCertificateRevocationRequest };
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
        const message = { ...baseQueryAllProposedCertificateRevocationRequest };
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
        const message = { ...baseQueryAllProposedCertificateRevocationRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllProposedCertificateRevocationResponse = {};
export const QueryAllProposedCertificateRevocationResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.proposedCertificateRevocation) {
            ProposedCertificateRevocation.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllProposedCertificateRevocationResponse };
        message.proposedCertificateRevocation = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllProposedCertificateRevocationResponse };
        message.proposedCertificateRevocation = [];
        if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
            for (const e of object.proposedCertificateRevocation) {
                message.proposedCertificateRevocation.push(ProposedCertificateRevocation.fromJSON(e));
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
        if (message.proposedCertificateRevocation) {
            obj.proposedCertificateRevocation = message.proposedCertificateRevocation.map((e) => (e ? ProposedCertificateRevocation.toJSON(e) : undefined));
        }
        else {
            obj.proposedCertificateRevocation = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllProposedCertificateRevocationResponse };
        message.proposedCertificateRevocation = [];
        if (object.proposedCertificateRevocation !== undefined && object.proposedCertificateRevocation !== null) {
            for (const e of object.proposedCertificateRevocation) {
                message.proposedCertificateRevocation.push(ProposedCertificateRevocation.fromPartial(e));
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
const baseQueryGetRevokedCertificatesRequest = { subject: '', subjectKeyId: '' };
export const QueryGetRevokedCertificatesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.subject !== '') {
            writer.uint32(10).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(18).string(message.subjectKeyId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRevokedCertificatesRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryGetRevokedCertificatesRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = String(object.subjectKeyId);
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRevokedCertificatesRequest };
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyId !== undefined && object.subjectKeyId !== null) {
            message.subjectKeyId = object.subjectKeyId;
        }
        else {
            message.subjectKeyId = '';
        }
        return message;
    }
};
const baseQueryGetRevokedCertificatesResponse = {};
export const QueryGetRevokedCertificatesResponse = {
    encode(message, writer = Writer.create()) {
        if (message.revokedCertificates !== undefined) {
            RevokedCertificates.encode(message.revokedCertificates, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRevokedCertificatesResponse };
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
    fromJSON(object) {
        const message = { ...baseQueryGetRevokedCertificatesResponse };
        if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
            message.revokedCertificates = RevokedCertificates.fromJSON(object.revokedCertificates);
        }
        else {
            message.revokedCertificates = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.revokedCertificates !== undefined &&
            (obj.revokedCertificates = message.revokedCertificates ? RevokedCertificates.toJSON(message.revokedCertificates) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRevokedCertificatesResponse };
        if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
            message.revokedCertificates = RevokedCertificates.fromPartial(object.revokedCertificates);
        }
        else {
            message.revokedCertificates = undefined;
        }
        return message;
    }
};
const baseQueryAllRevokedCertificatesRequest = {};
export const QueryAllRevokedCertificatesRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRevokedCertificatesRequest };
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
        const message = { ...baseQueryAllRevokedCertificatesRequest };
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
        const message = { ...baseQueryAllRevokedCertificatesRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllRevokedCertificatesResponse = {};
export const QueryAllRevokedCertificatesResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.revokedCertificates) {
            RevokedCertificates.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRevokedCertificatesResponse };
        message.revokedCertificates = [];
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
    fromJSON(object) {
        const message = { ...baseQueryAllRevokedCertificatesResponse };
        message.revokedCertificates = [];
        if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
            for (const e of object.revokedCertificates) {
                message.revokedCertificates.push(RevokedCertificates.fromJSON(e));
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
        if (message.revokedCertificates) {
            obj.revokedCertificates = message.revokedCertificates.map((e) => (e ? RevokedCertificates.toJSON(e) : undefined));
        }
        else {
            obj.revokedCertificates = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRevokedCertificatesResponse };
        message.revokedCertificates = [];
        if (object.revokedCertificates !== undefined && object.revokedCertificates !== null) {
            for (const e of object.revokedCertificates) {
                message.revokedCertificates.push(RevokedCertificates.fromPartial(e));
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
    ApprovedCertificates(request) {
        const data = QueryGetApprovedCertificatesRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ApprovedCertificates', data);
        return promise.then((data) => QueryGetApprovedCertificatesResponse.decode(new Reader(data)));
    }
    ApprovedCertificatesAll(request) {
        const data = QueryAllApprovedCertificatesRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ApprovedCertificatesAll', data);
        return promise.then((data) => QueryAllApprovedCertificatesResponse.decode(new Reader(data)));
    }
    ProposedCertificate(request) {
        const data = QueryGetProposedCertificateRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificate', data);
        return promise.then((data) => QueryGetProposedCertificateResponse.decode(new Reader(data)));
    }
    ProposedCertificateAll(request) {
        const data = QueryAllProposedCertificateRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificateAll', data);
        return promise.then((data) => QueryAllProposedCertificateResponse.decode(new Reader(data)));
    }
    ChildCertificates(request) {
        const data = QueryGetChildCertificatesRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ChildCertificates', data);
        return promise.then((data) => QueryGetChildCertificatesResponse.decode(new Reader(data)));
    }
    ChildCertificatesAll(request) {
        const data = QueryAllChildCertificatesRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ChildCertificatesAll', data);
        return promise.then((data) => QueryAllChildCertificatesResponse.decode(new Reader(data)));
    }
    ProposedCertificateRevocation(request) {
        const data = QueryGetProposedCertificateRevocationRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificateRevocation', data);
        return promise.then((data) => QueryGetProposedCertificateRevocationResponse.decode(new Reader(data)));
    }
    ProposedCertificateRevocationAll(request) {
        const data = QueryAllProposedCertificateRevocationRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'ProposedCertificateRevocationAll', data);
        return promise.then((data) => QueryAllProposedCertificateRevocationResponse.decode(new Reader(data)));
    }
    RevokedCertificates(request) {
        const data = QueryGetRevokedCertificatesRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RevokedCertificates', data);
        return promise.then((data) => QueryGetRevokedCertificatesResponse.decode(new Reader(data)));
    }
    RevokedCertificatesAll(request) {
        const data = QueryAllRevokedCertificatesRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.pki.Query', 'RevokedCertificatesAll', data);
        return promise.then((data) => QueryAllRevokedCertificatesResponse.decode(new Reader(data)));
    }
}
