/* eslint-disable */
import { ComplianceInfo } from '../compliance/compliance_info';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseDeviceSoftwareCompliance = { cDCertificateId: '' };
export const DeviceSoftwareCompliance = {
    encode(message, writer = Writer.create()) {
        if (message.cDCertificateId !== '') {
            writer.uint32(10).string(message.cDCertificateId);
        }
        for (const v of message.complianceInfo) {
            ComplianceInfo.encode(v, writer.uint32(18).fork()).ldelim();
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
                    message.cDCertificateId = reader.string();
                    break;
                case 2:
                    message.complianceInfo.push(ComplianceInfo.decode(reader, reader.uint32()));
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
        if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
            message.cDCertificateId = String(object.cDCertificateId);
        }
        else {
            message.cDCertificateId = '';
        }
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(ComplianceInfo.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.cDCertificateId !== undefined && (obj.cDCertificateId = message.cDCertificateId);
        if (message.complianceInfo) {
            obj.complianceInfo = message.complianceInfo.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined));
        }
        else {
            obj.complianceInfo = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseDeviceSoftwareCompliance };
        message.complianceInfo = [];
        if (object.cDCertificateId !== undefined && object.cDCertificateId !== null) {
            message.cDCertificateId = object.cDCertificateId;
        }
        else {
            message.cDCertificateId = '';
        }
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(ComplianceInfo.fromPartial(e));
            }
        }
        return message;
    }
};
