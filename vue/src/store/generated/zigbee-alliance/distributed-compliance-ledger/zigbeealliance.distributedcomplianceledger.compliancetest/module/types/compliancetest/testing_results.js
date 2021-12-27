/* eslint-disable */
import { TestingResult } from '../compliancetest/testing_result';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliancetest';
const baseTestingResults = { vid: 0, pid: 0, softwareVersion: 0, softwareVersionString: '' };
export const TestingResults = {
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
        for (const v of message.results) {
            TestingResult.encode(v, writer.uint32(34).fork()).ldelim();
        }
        if (message.softwareVersionString !== '') {
            writer.uint32(42).string(message.softwareVersionString);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseTestingResults };
        message.results = [];
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
                    message.results.push(TestingResult.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.softwareVersionString = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseTestingResults };
        message.results = [];
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
        if (object.results !== undefined && object.results !== null) {
            for (const e of object.results) {
                message.results.push(TestingResult.fromJSON(e));
            }
        }
        if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
            message.softwareVersionString = String(object.softwareVersionString);
        }
        else {
            message.softwareVersionString = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        if (message.results) {
            obj.results = message.results.map((e) => (e ? TestingResult.toJSON(e) : undefined));
        }
        else {
            obj.results = [];
        }
        message.softwareVersionString !== undefined && (obj.softwareVersionString = message.softwareVersionString);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseTestingResults };
        message.results = [];
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
        if (object.results !== undefined && object.results !== null) {
            for (const e of object.results) {
                message.results.push(TestingResult.fromPartial(e));
            }
        }
        if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
            message.softwareVersionString = object.softwareVersionString;
        }
        else {
            message.softwareVersionString = '';
        }
        return message;
    }
};
