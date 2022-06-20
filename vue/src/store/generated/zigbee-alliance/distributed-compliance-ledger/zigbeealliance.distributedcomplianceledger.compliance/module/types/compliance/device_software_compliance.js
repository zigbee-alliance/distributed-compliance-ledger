/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseDeviceSoftwareCompliance = { cdCertificateId: '', complianceInfo: '' };
export const DeviceSoftwareCompliance = {
    encode(message, writer = Writer.create()) {
        if (message.cdCertificateId !== '') {
            writer.uint32(10).string(message.cdCertificateId);
        }
        for (const v of message.complianceInfo) {
            writer.uint32(18).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseDeviceSoftwareCompliance };
        message.complianceInfo = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.cdCertificateId = reader.string();
                    break;
                case 2:
                    message.complianceInfo.push(reader.string());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseDeviceSoftwareCompliance };
        message.complianceInfo = [];
        if (object.cdCertificateId !== undefined && object.cdCertificateId !== null) {
            message.cdCertificateId = String(object.cdCertificateId);
        }
        else {
            message.cdCertificateId = '';
        }
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cdCertificateId !== undefined && (obj.cdCertificateId = message.cdCertificateId);
        if (message.complianceInfo) {
            obj.complianceInfo = message.complianceInfo.map((e) => e);
        }
        else {
            obj.complianceInfo = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseDeviceSoftwareCompliance };
        message.complianceInfo = [];
        if (object.cdCertificateId !== undefined && object.cdCertificateId !== null) {
            message.cdCertificateId = object.cdCertificateId;
        }
        else {
            message.cdCertificateId = '';
        }
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(e);
            }
        }
        return message;
    }
};
