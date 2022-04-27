/* eslint-disable */
import { Grant } from '../pki/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseCertificate = {
    pemCert: '',
    serialNumber: '',
    issuer: '',
    authorityKeyId: '',
    rootSubject: '',
    rootSubjectKeyId: '',
    isRoot: false,
    owner: '',
    subject: '',
    subjectKeyId: '',
    subjectAsText: ''
};
export const Certificate = {
    encode(message, writer = Writer.create()) {
        if (message.pemCert !== '') {
            writer.uint32(10).string(message.pemCert);
        }
        if (message.serialNumber !== '') {
            writer.uint32(18).string(message.serialNumber);
        }
        if (message.issuer !== '') {
            writer.uint32(26).string(message.issuer);
        }
        if (message.authorityKeyId !== '') {
            writer.uint32(34).string(message.authorityKeyId);
        }
        if (message.rootSubject !== '') {
            writer.uint32(42).string(message.rootSubject);
        }
        if (message.rootSubjectKeyId !== '') {
            writer.uint32(50).string(message.rootSubjectKeyId);
        }
        if (message.isRoot === true) {
            writer.uint32(56).bool(message.isRoot);
        }
        if (message.owner !== '') {
            writer.uint32(66).string(message.owner);
        }
        if (message.subject !== '') {
            writer.uint32(74).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(82).string(message.subjectKeyId);
        }
        for (const v of message.approvals) {
            Grant.encode(v, writer.uint32(90).fork()).ldelim();
        }
        if (message.subjectAsText !== '') {
            writer.uint32(98).string(message.subjectAsText);
        }
        for (const v of message.rejects) {
            Grant.encode(v, writer.uint32(106).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseCertificate };
        message.approvals = [];
        message.rejects = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pemCert = reader.string();
                    break;
                case 2:
                    message.serialNumber = reader.string();
                    break;
                case 3:
                    message.issuer = reader.string();
                    break;
                case 4:
                    message.authorityKeyId = reader.string();
                    break;
                case 5:
                    message.rootSubject = reader.string();
                    break;
                case 6:
                    message.rootSubjectKeyId = reader.string();
                    break;
                case 7:
                    message.isRoot = reader.bool();
                    break;
                case 8:
                    message.owner = reader.string();
                    break;
                case 9:
                    message.subject = reader.string();
                    break;
                case 10:
                    message.subjectKeyId = reader.string();
                    break;
                case 11:
                    message.approvals.push(Grant.decode(reader, reader.uint32()));
                    break;
                case 12:
                    message.subjectAsText = reader.string();
                    break;
                case 13:
                    message.rejects.push(Grant.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseCertificate };
        message.approvals = [];
        message.rejects = [];
        if (object.pemCert !== undefined && object.pemCert !== null) {
            message.pemCert = String(object.pemCert);
        }
        else {
            message.pemCert = '';
        }
        if (object.serialNumber !== undefined && object.serialNumber !== null) {
            message.serialNumber = String(object.serialNumber);
        }
        else {
            message.serialNumber = '';
        }
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
        if (object.rootSubject !== undefined && object.rootSubject !== null) {
            message.rootSubject = String(object.rootSubject);
        }
        else {
            message.rootSubject = '';
        }
        if (object.rootSubjectKeyId !== undefined && object.rootSubjectKeyId !== null) {
            message.rootSubjectKeyId = String(object.rootSubjectKeyId);
        }
        else {
            message.rootSubjectKeyId = '';
        }
        if (object.isRoot !== undefined && object.isRoot !== null) {
            message.isRoot = Boolean(object.isRoot);
        }
        else {
            message.isRoot = false;
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = '';
        }
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
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(Grant.fromJSON(e));
            }
        }
        if (object.subjectAsText !== undefined && object.subjectAsText !== null) {
            message.subjectAsText = String(object.subjectAsText);
        }
        else {
            message.subjectAsText = '';
        }
        if (object.rejects !== undefined && object.rejects !== null) {
            for (const e of object.rejects) {
                message.rejects.push(Grant.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pemCert !== undefined && (obj.pemCert = message.pemCert);
        message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber);
        message.issuer !== undefined && (obj.issuer = message.issuer);
        message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId);
        message.rootSubject !== undefined && (obj.rootSubject = message.rootSubject);
        message.rootSubjectKeyId !== undefined && (obj.rootSubjectKeyId = message.rootSubjectKeyId);
        message.isRoot !== undefined && (obj.isRoot = message.isRoot);
        message.owner !== undefined && (obj.owner = message.owner);
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        if (message.approvals) {
            obj.approvals = message.approvals.map((e) => (e ? Grant.toJSON(e) : undefined));
        }
        else {
            obj.approvals = [];
        }
        message.subjectAsText !== undefined && (obj.subjectAsText = message.subjectAsText);
        if (message.rejects) {
            obj.rejects = message.rejects.map((e) => (e ? Grant.toJSON(e) : undefined));
        }
        else {
            obj.rejects = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseCertificate };
        message.approvals = [];
        message.rejects = [];
        if (object.pemCert !== undefined && object.pemCert !== null) {
            message.pemCert = object.pemCert;
        }
        else {
            message.pemCert = '';
        }
        if (object.serialNumber !== undefined && object.serialNumber !== null) {
            message.serialNumber = object.serialNumber;
        }
        else {
            message.serialNumber = '';
        }
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
        if (object.rootSubject !== undefined && object.rootSubject !== null) {
            message.rootSubject = object.rootSubject;
        }
        else {
            message.rootSubject = '';
        }
        if (object.rootSubjectKeyId !== undefined && object.rootSubjectKeyId !== null) {
            message.rootSubjectKeyId = object.rootSubjectKeyId;
        }
        else {
            message.rootSubjectKeyId = '';
        }
        if (object.isRoot !== undefined && object.isRoot !== null) {
            message.isRoot = object.isRoot;
        }
        else {
            message.isRoot = false;
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = '';
        }
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
        if (object.approvals !== undefined && object.approvals !== null) {
            for (const e of object.approvals) {
                message.approvals.push(Grant.fromPartial(e));
            }
        }
        if (object.subjectAsText !== undefined && object.subjectAsText !== null) {
            message.subjectAsText = object.subjectAsText;
        }
        else {
            message.subjectAsText = '';
        }
        if (object.rejects !== undefined && object.rejects !== null) {
            for (const e of object.rejects) {
                message.rejects.push(Grant.fromPartial(e));
            }
        }
        return message;
    }
};
