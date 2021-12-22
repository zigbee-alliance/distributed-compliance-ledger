/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseApprovedCertificatesBySubject = { subject: '', subjectKeyIds: '' };
export const ApprovedCertificatesBySubject = {
    encode(message, writer = Writer.create()) {
        if (message.subject !== '') {
            writer.uint32(10).string(message.subject);
        }
        for (const v of message.subjectKeyIds) {
            writer.uint32(18).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseApprovedCertificatesBySubject };
        message.subjectKeyIds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.subject = reader.string();
                    break;
                case 2:
                    message.subjectKeyIds.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseApprovedCertificatesBySubject };
        message.subjectKeyIds = [];
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = String(object.subject);
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyIds !== undefined && object.subjectKeyIds !== null) {
            for (const e of object.subjectKeyIds) {
                message.subjectKeyIds.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.subject !== undefined && (obj.subject = message.subject);
        if (message.subjectKeyIds) {
            obj.subjectKeyIds = message.subjectKeyIds.map((e) => e);
        }
        else {
            obj.subjectKeyIds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseApprovedCertificatesBySubject };
        message.subjectKeyIds = [];
        if (object.subject !== undefined && object.subject !== null) {
            message.subject = object.subject;
        }
        else {
            message.subject = '';
        }
        if (object.subjectKeyIds !== undefined && object.subjectKeyIds !== null) {
            for (const e of object.subjectKeyIds) {
                message.subjectKeyIds.push(e);
            }
        }
        return message;
    }
};
