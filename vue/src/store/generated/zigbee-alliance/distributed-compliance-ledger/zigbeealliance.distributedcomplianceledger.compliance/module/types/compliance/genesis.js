/* eslint-disable */
import { ComplianceInfo } from '../compliance/compliance_info';
import { CertifiedModel } from '../compliance/certified_model';
import { RevokedModel } from '../compliance/revoked_model';
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
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.complianceInfoList = [];
        message.certifiedModelList = [];
        message.revokedModelList = [];
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
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.complianceInfoList = [];
        message.certifiedModelList = [];
        message.revokedModelList = [];
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
        return message;
    }
};
