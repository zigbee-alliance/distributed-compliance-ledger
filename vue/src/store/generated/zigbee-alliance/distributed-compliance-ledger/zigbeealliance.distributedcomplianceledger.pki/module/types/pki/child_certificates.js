/* eslint-disable */
import { CertificateIdentifier } from '../pki/certificate_identifier';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseChildCertificates = { issuer: '', authorityKeyId: '' };
export const ChildCertificates = {
    encode(message, writer = Writer.create()) {
        if (message.issuer !== '') {
            writer.uint32(10).string(message.issuer);
        }
        if (message.authorityKeyId !== '') {
            writer.uint32(18).string(message.authorityKeyId);
        }
        for (const v of message.certIds) {
            CertificateIdentifier.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseChildCertificates };
        message.certIds = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.issuer = reader.string();
                    break;
                case 2:
                    message.authorityKeyId = reader.string();
                    break;
                case 3:
                    message.certIds.push(CertificateIdentifier.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseChildCertificates };
        message.certIds = [];
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
        if (object.certIds !== undefined && object.certIds !== null) {
            for (const e of object.certIds) {
                message.certIds.push(CertificateIdentifier.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.issuer !== undefined && (obj.issuer = message.issuer);
        message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId);
        if (message.certIds) {
            obj.certIds = message.certIds.map((e) => (e ? CertificateIdentifier.toJSON(e) : undefined));
        }
        else {
            obj.certIds = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseChildCertificates };
        message.certIds = [];
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
        if (object.certIds !== undefined && object.certIds !== null) {
            for (const e of object.certIds) {
                message.certIds.push(CertificateIdentifier.fromPartial(e));
            }
        }
        return message;
    }
};
