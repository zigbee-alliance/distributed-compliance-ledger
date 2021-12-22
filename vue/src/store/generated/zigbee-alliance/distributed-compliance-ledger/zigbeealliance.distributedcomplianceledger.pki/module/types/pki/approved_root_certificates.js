/* eslint-disable */
import { CertificateIdentifier } from '../pki/certificate_identifier';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseApprovedRootCertificates = {};
export const ApprovedRootCertificates = {
    encode(message, writer = Writer.create()) {
        for (const v of message.certs) {
            CertificateIdentifier.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseApprovedRootCertificates };
        message.certs = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.certs.push(CertificateIdentifier.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseApprovedRootCertificates };
        message.certs = [];
        if (object.certs !== undefined && object.certs !== null) {
            for (const e of object.certs) {
                message.certs.push(CertificateIdentifier.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.certs) {
            obj.certs = message.certs.map((e) => (e ? CertificateIdentifier.toJSON(e) : undefined));
        }
        else {
            obj.certs = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseApprovedRootCertificates };
        message.certs = [];
        if (object.certs !== undefined && object.certs !== null) {
            for (const e of object.certs) {
                message.certs.push(CertificateIdentifier.fromPartial(e));
            }
        }
        return message;
    }
};
