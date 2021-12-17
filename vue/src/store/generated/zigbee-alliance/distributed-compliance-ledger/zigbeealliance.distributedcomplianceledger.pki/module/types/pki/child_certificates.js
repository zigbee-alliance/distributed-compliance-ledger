/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseChildCertificates = { issuer: '', authorityKeyId: '', certIds: '' };
export const ChildCertificates = {
    encode(message, writer = Writer.create()) {
        if (message.issuer !== '') {
            writer.uint32(10).string(message.issuer);
        }
        if (message.authorityKeyId !== '') {
            writer.uint32(18).string(message.authorityKeyId);
        }
        for (const v of message.certIds) {
            writer.uint32(26).string(v);
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
                    message.certIds.push(reader.string());
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
                message.certIds.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.issuer !== undefined && (obj.issuer = message.issuer);
        message.authorityKeyId !== undefined && (obj.authorityKeyId = message.authorityKeyId);
        if (message.certIds) {
            obj.certIds = message.certIds.map((e) => e);
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
                message.certIds.push(e);
            }
        }
        return message;
    }
};
