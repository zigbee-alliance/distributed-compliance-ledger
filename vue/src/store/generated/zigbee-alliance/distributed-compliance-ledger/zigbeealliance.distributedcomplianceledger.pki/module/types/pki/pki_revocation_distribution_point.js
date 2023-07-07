/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const basePkiRevocationDistributionPoint = {
    vid: 0,
    label: '',
    issuerSubjectKeyID: '',
    pid: 0,
    isPAA: false,
    crlSignerCertificate: '',
    dataURL: '',
    dataFileSize: 0,
    dataDigest: '',
    dataDigestType: 0,
    revocationType: 0
};
export const PkiRevocationDistributionPoint = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).uint64(message.vid);
        }
        if (message.label !== '') {
            writer.uint32(18).string(message.label);
        }
        if (message.issuerSubjectKeyID !== '') {
            writer.uint32(26).string(message.issuerSubjectKeyID);
        }
        if (message.pid !== 0) {
            writer.uint32(32).uint64(message.pid);
        }
        if (message.isPAA === true) {
            writer.uint32(40).bool(message.isPAA);
        }
        if (message.crlSignerCertificate !== '') {
            writer.uint32(50).string(message.crlSignerCertificate);
        }
        if (message.dataURL !== '') {
            writer.uint32(58).string(message.dataURL);
        }
        if (message.dataFileSize !== 0) {
            writer.uint32(64).uint64(message.dataFileSize);
        }
        if (message.dataDigest !== '') {
            writer.uint32(74).string(message.dataDigest);
        }
        if (message.dataDigestType !== 0) {
            writer.uint32(80).uint64(message.dataDigestType);
        }
        if (message.revocationType !== 0) {
            writer.uint32(88).uint64(message.revocationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...basePkiRevocationDistributionPoint };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = longToNumber(reader.uint64());
                    break;
                case 2:
                    message.label = reader.string();
                    break;
                case 3:
                    message.issuerSubjectKeyID = reader.string();
                    break;
                case 4:
                    message.pid = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.isPAA = reader.bool();
                    break;
                case 6:
                    message.crlSignerCertificate = reader.string();
                    break;
                case 7:
                    message.dataURL = reader.string();
                    break;
                case 8:
                    message.dataFileSize = longToNumber(reader.uint64());
                    break;
                case 9:
                    message.dataDigest = reader.string();
                    break;
                case 10:
                    message.dataDigestType = longToNumber(reader.uint64());
                    break;
                case 11:
                    message.revocationType = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...basePkiRevocationDistributionPoint };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = String(object.label);
        }
        else {
            message.label = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = String(object.issuerSubjectKeyID);
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.isPAA !== undefined && object.isPAA !== null) {
            message.isPAA = Boolean(object.isPAA);
        }
        else {
            message.isPAA = false;
        }
        if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
            message.crlSignerCertificate = String(object.crlSignerCertificate);
        }
        else {
            message.crlSignerCertificate = '';
        }
        if (object.dataURL !== undefined && object.dataURL !== null) {
            message.dataURL = String(object.dataURL);
        }
        else {
            message.dataURL = '';
        }
        if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
            message.dataFileSize = Number(object.dataFileSize);
        }
        else {
            message.dataFileSize = 0;
        }
        if (object.dataDigest !== undefined && object.dataDigest !== null) {
            message.dataDigest = String(object.dataDigest);
        }
        else {
            message.dataDigest = '';
        }
        if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
            message.dataDigestType = Number(object.dataDigestType);
        }
        else {
            message.dataDigestType = 0;
        }
        if (object.revocationType !== undefined && object.revocationType !== null) {
            message.revocationType = Number(object.revocationType);
        }
        else {
            message.revocationType = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.label !== undefined && (obj.label = message.label);
        message.issuerSubjectKeyID !== undefined && (obj.issuerSubjectKeyID = message.issuerSubjectKeyID);
        message.pid !== undefined && (obj.pid = message.pid);
        message.isPAA !== undefined && (obj.isPAA = message.isPAA);
        message.crlSignerCertificate !== undefined && (obj.crlSignerCertificate = message.crlSignerCertificate);
        message.dataURL !== undefined && (obj.dataURL = message.dataURL);
        message.dataFileSize !== undefined && (obj.dataFileSize = message.dataFileSize);
        message.dataDigest !== undefined && (obj.dataDigest = message.dataDigest);
        message.dataDigestType !== undefined && (obj.dataDigestType = message.dataDigestType);
        message.revocationType !== undefined && (obj.revocationType = message.revocationType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...basePkiRevocationDistributionPoint };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.label !== undefined && object.label !== null) {
            message.label = object.label;
        }
        else {
            message.label = '';
        }
        if (object.issuerSubjectKeyID !== undefined && object.issuerSubjectKeyID !== null) {
            message.issuerSubjectKeyID = object.issuerSubjectKeyID;
        }
        else {
            message.issuerSubjectKeyID = '';
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.isPAA !== undefined && object.isPAA !== null) {
            message.isPAA = object.isPAA;
        }
        else {
            message.isPAA = false;
        }
        if (object.crlSignerCertificate !== undefined && object.crlSignerCertificate !== null) {
            message.crlSignerCertificate = object.crlSignerCertificate;
        }
        else {
            message.crlSignerCertificate = '';
        }
        if (object.dataURL !== undefined && object.dataURL !== null) {
            message.dataURL = object.dataURL;
        }
        else {
            message.dataURL = '';
        }
        if (object.dataFileSize !== undefined && object.dataFileSize !== null) {
            message.dataFileSize = object.dataFileSize;
        }
        else {
            message.dataFileSize = 0;
        }
        if (object.dataDigest !== undefined && object.dataDigest !== null) {
            message.dataDigest = object.dataDigest;
        }
        else {
            message.dataDigest = '';
        }
        if (object.dataDigestType !== undefined && object.dataDigestType !== null) {
            message.dataDigestType = object.dataDigestType;
        }
        else {
            message.dataDigestType = 0;
        }
        if (object.revocationType !== undefined && object.revocationType !== null) {
            message.revocationType = object.revocationType;
        }
        else {
            message.revocationType = 0;
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
