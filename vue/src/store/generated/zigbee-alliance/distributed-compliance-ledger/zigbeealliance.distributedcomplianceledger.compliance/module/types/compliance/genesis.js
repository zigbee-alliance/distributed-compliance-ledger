/* eslint-disable */
import { ComplianceInfo } from '../compliance/compliance_info';
import { CertifiedModel } from '../compliance/certified_model';
import { RevokedModel } from '../compliance/revoked_model';
import { ProvisionalModel } from '../compliance/provisional_model';
import { DeviceSoftwareCompliance } from '../compliance/device_software_compliance';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.complianceInfoList) {
            ComplianceInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.certifiedModelList) {
            CertifiedModel.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.revokedModelList) {
            RevokedModel.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.provisionalModelList) {
            ProvisionalModel.encode(v, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.deviceSoftwareComplianceList) {
            DeviceSoftwareCompliance.encode(v, writer.uint32(42).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.complianceInfoList = [];
        message.certifiedModelList = [];
        message.revokedModelList = [];
        message.provisionalModelList = [];
        message.deviceSoftwareComplianceList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.complianceInfoList.push(ComplianceInfo.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.certifiedModelList.push(CertifiedModel.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.revokedModelList.push(RevokedModel.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.provisionalModelList.push(ProvisionalModel.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.deviceSoftwareComplianceList.push(DeviceSoftwareCompliance.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.complianceInfoList = [];
        message.certifiedModelList = [];
        message.revokedModelList = [];
        message.provisionalModelList = [];
        message.deviceSoftwareComplianceList = [];
        if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
            for (const e of object.complianceInfoList) {
                message.complianceInfoList.push(ComplianceInfo.fromJSON(e));
            }
        }
        if (object.certifiedModelList !== undefined && object.certifiedModelList !== null) {
            for (const e of object.certifiedModelList) {
                message.certifiedModelList.push(CertifiedModel.fromJSON(e));
            }
        }
        if (object.revokedModelList !== undefined && object.revokedModelList !== null) {
            for (const e of object.revokedModelList) {
                message.revokedModelList.push(RevokedModel.fromJSON(e));
            }
        }
        if (object.provisionalModelList !== undefined && object.provisionalModelList !== null) {
            for (const e of object.provisionalModelList) {
                message.provisionalModelList.push(ProvisionalModel.fromJSON(e));
            }
        }
        if (object.deviceSoftwareComplianceList !== undefined && object.deviceSoftwareComplianceList !== null) {
            for (const e of object.deviceSoftwareComplianceList) {
                message.deviceSoftwareComplianceList.push(DeviceSoftwareCompliance.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.complianceInfoList) {
            obj.complianceInfoList = message.complianceInfoList.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined));
        }
        else {
            obj.complianceInfoList = [];
        }
        if (message.certifiedModelList) {
            obj.certifiedModelList = message.certifiedModelList.map((e) => (e ? CertifiedModel.toJSON(e) : undefined));
        }
        else {
            obj.certifiedModelList = [];
        }
        if (message.revokedModelList) {
            obj.revokedModelList = message.revokedModelList.map((e) => (e ? RevokedModel.toJSON(e) : undefined));
        }
        else {
            obj.revokedModelList = [];
        }
        if (message.provisionalModelList) {
            obj.provisionalModelList = message.provisionalModelList.map((e) => (e ? ProvisionalModel.toJSON(e) : undefined));
        }
        else {
            obj.provisionalModelList = [];
        }
        if (message.deviceSoftwareComplianceList) {
            obj.deviceSoftwareComplianceList = message.deviceSoftwareComplianceList.map((e) => (e ? DeviceSoftwareCompliance.toJSON(e) : undefined));
        }
        else {
            obj.deviceSoftwareComplianceList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.complianceInfoList = [];
        message.certifiedModelList = [];
        message.revokedModelList = [];
        message.provisionalModelList = [];
        message.deviceSoftwareComplianceList = [];
        if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
            for (const e of object.complianceInfoList) {
                message.complianceInfoList.push(ComplianceInfo.fromPartial(e));
            }
        }
        if (object.certifiedModelList !== undefined && object.certifiedModelList !== null) {
            for (const e of object.certifiedModelList) {
                message.certifiedModelList.push(CertifiedModel.fromPartial(e));
            }
        }
        if (object.revokedModelList !== undefined && object.revokedModelList !== null) {
            for (const e of object.revokedModelList) {
                message.revokedModelList.push(RevokedModel.fromPartial(e));
            }
        }
        if (object.provisionalModelList !== undefined && object.provisionalModelList !== null) {
            for (const e of object.provisionalModelList) {
                message.provisionalModelList.push(ProvisionalModel.fromPartial(e));
            }
        }
        if (object.deviceSoftwareComplianceList !== undefined && object.deviceSoftwareComplianceList !== null) {
            for (const e of object.deviceSoftwareComplianceList) {
                message.deviceSoftwareComplianceList.push(DeviceSoftwareCompliance.fromPartial(e));
            }
        }
        return message;
    }
};
