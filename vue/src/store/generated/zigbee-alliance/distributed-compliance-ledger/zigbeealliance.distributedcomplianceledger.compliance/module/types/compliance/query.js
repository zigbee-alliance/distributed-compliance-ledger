/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { ComplianceInfo } from '../compliance/compliance_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseQueryGetComplianceInfoRequest = { vid: 0, pid: 0, softwareVersion: 0, certificationType: '' };
export const QueryGetComplianceInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint32(message.softwareVersion);
        }
        if (message.certificationType !== '') {
            writer.uint32(34).string(message.certificationType);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetComplianceInfoRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.pid = reader.int32();
                    break;
                case 3:
                    message.softwareVersion = reader.uint32();
                    break;
                case 4:
                    message.certificationType = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetComplianceInfoRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetComplianceInfoRequest };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        return message;
    }
};
const baseQueryGetComplianceInfoResponse = {};
export const QueryGetComplianceInfoResponse = {
    encode(message, writer = Writer.create()) {
        if (message.complianceInfo !== undefined) {
            ComplianceInfo.encode(message.complianceInfo, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetComplianceInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.complianceInfo = ComplianceInfo.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetComplianceInfoResponse };
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            message.complianceInfo = ComplianceInfo.fromJSON(object.complianceInfo);
        }
        else {
            message.complianceInfo = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.complianceInfo !== undefined && (obj.complianceInfo = message.complianceInfo ? ComplianceInfo.toJSON(message.complianceInfo) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetComplianceInfoResponse };
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            message.complianceInfo = ComplianceInfo.fromPartial(object.complianceInfo);
        }
        else {
            message.complianceInfo = undefined;
        }
        return message;
    }
};
const baseQueryAllComplianceInfoRequest = {};
export const QueryAllComplianceInfoRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllComplianceInfoRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllComplianceInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllComplianceInfoRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllComplianceInfoResponse = {};
export const QueryAllComplianceInfoResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.complianceInfo) {
            ComplianceInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllComplianceInfoResponse };
        message.complianceInfo = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.complianceInfo.push(ComplianceInfo.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllComplianceInfoResponse };
        message.complianceInfo = [];
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(ComplianceInfo.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.complianceInfo) {
            obj.complianceInfo = message.complianceInfo.map((e) => (e ? ComplianceInfo.toJSON(e) : undefined));
        }
        else {
            obj.complianceInfo = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllComplianceInfoResponse };
        message.complianceInfo = [];
        if (object.complianceInfo !== undefined && object.complianceInfo !== null) {
            for (const e of object.complianceInfo) {
                message.complianceInfo.push(ComplianceInfo.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    ComplianceInfo(request) {
        const data = QueryGetComplianceInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfo', data);
        return promise.then((data) => QueryGetComplianceInfoResponse.decode(new Reader(data)));
    }
    ComplianceInfoAll(request) {
        const data = QueryAllComplianceInfoRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliance.Query', 'ComplianceInfoAll', data);
        return promise.then((data) => QueryAllComplianceInfoResponse.decode(new Reader(data)));
    }
}
