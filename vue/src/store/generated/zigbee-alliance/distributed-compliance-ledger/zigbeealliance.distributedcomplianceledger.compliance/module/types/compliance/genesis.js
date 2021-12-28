/* eslint-disable */
import { ComplianceInfo } from '../compliance/compliance_info';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.complianceInfoList) {
            ComplianceInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.complianceInfoList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.complianceInfoList.push(ComplianceInfo.decode(reader, reader.uint32()));
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
        if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
            for (const e of object.complianceInfoList) {
                message.complianceInfoList.push(ComplianceInfo.fromJSON(e));
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
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.complianceInfoList = [];
        if (object.complianceInfoList !== undefined && object.complianceInfoList !== null) {
            for (const e of object.complianceInfoList) {
                message.complianceInfoList.push(ComplianceInfo.fromPartial(e));
            }
        }
        return message;
    }
};
