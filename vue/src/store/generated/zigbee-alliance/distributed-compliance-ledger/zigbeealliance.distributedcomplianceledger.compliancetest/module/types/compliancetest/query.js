/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { TestingResults } from '../compliancetest/testing_results';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliancetest';
const baseQueryGetTestingResultsRequest = { vid: 0, pid: 0, softwareVersion: 0 };
export const QueryGetTestingResultsRequest = {
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
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetTestingResultsRequest };
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
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetTestingResultsRequest };
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetTestingResultsRequest };
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
        return message;
    }
};
const baseQueryGetTestingResultsResponse = {};
export const QueryGetTestingResultsResponse = {
    encode(message, writer = Writer.create()) {
        if (message.testingResults !== undefined) {
            TestingResults.encode(message.testingResults, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetTestingResultsResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.testingResults = TestingResults.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetTestingResultsResponse };
        if (object.testingResults !== undefined && object.testingResults !== null) {
            message.testingResults = TestingResults.fromJSON(object.testingResults);
        }
        else {
            message.testingResults = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.testingResults !== undefined && (obj.testingResults = message.testingResults ? TestingResults.toJSON(message.testingResults) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetTestingResultsResponse };
        if (object.testingResults !== undefined && object.testingResults !== null) {
            message.testingResults = TestingResults.fromPartial(object.testingResults);
        }
        else {
            message.testingResults = undefined;
        }
        return message;
    }
};
const baseQueryAllTestingResultsRequest = {};
export const QueryAllTestingResultsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllTestingResultsRequest };
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
        const message = { ...baseQueryAllTestingResultsRequest };
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
        const message = { ...baseQueryAllTestingResultsRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllTestingResultsResponse = {};
export const QueryAllTestingResultsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.testingResults) {
            TestingResults.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllTestingResultsResponse };
        message.testingResults = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.testingResults.push(TestingResults.decode(reader, reader.uint32()));
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
        const message = { ...baseQueryAllTestingResultsResponse };
        message.testingResults = [];
        if (object.testingResults !== undefined && object.testingResults !== null) {
            for (const e of object.testingResults) {
                message.testingResults.push(TestingResults.fromJSON(e));
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
        if (message.testingResults) {
            obj.testingResults = message.testingResults.map((e) => (e ? TestingResults.toJSON(e) : undefined));
        }
        else {
            obj.testingResults = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllTestingResultsResponse };
        message.testingResults = [];
        if (object.testingResults !== undefined && object.testingResults !== null) {
            for (const e of object.testingResults) {
                message.testingResults.push(TestingResults.fromPartial(e));
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
    TestingResults(request) {
        const data = QueryGetTestingResultsRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliancetest.Query', 'TestingResults', data);
        return promise.then((data) => QueryGetTestingResultsResponse.decode(new Reader(data)));
    }
    TestingResultsAll(request) {
        const data = QueryAllTestingResultsRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.compliancetest.Query', 'TestingResultsAll', data);
        return promise.then((data) => QueryAllTestingResultsResponse.decode(new Reader(data)));
    }
}
