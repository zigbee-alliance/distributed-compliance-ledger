/* eslint-disable */
import { Grant } from '../pki/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseProposedCertificateRevocation = { subject: '', subjectKeyId: '', subjectAsText: '' };
export const ProposedCertificateRevocation = {
    encode(message, writer = Writer.create()) {
        if (message.subject !== '') {
            writer.uint32(10).string(message.subject);
        }
        if (message.subjectKeyId !== '') {
            writer.uint32(18).string(message.subjectKeyId);
        }
        for (const v of message.approvals) {
            Grant.encode(v, writer.uint32(26).fork()).ldelim();
        }
        if (message.subjectAsText !== '') {
            writer.uint32(34).string(message.subjectAsText);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseProposedCertificateRevocation };
        message.approvals = [];
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
                    message.approvals.push(Grant.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.subjectAsText = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseProposedCertificateRevocation };
        message.approvals = [];
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.subject !== undefined && (obj.subject = message.subject);
        message.subjectKeyId !== undefined && (obj.subjectKeyId = message.subjectKeyId);
        if (message.approvals) {
            obj.approvals = message.approvals.map((e) => (e ? Grant.toJSON(e) : undefined));
        }
        else {
            obj.approvals = [];
        }
        message.subjectAsText !== undefined && (obj.subjectAsText = message.subjectAsText);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseProposedCertificateRevocation };
        message.approvals = [];
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
        return message;
    }
};
