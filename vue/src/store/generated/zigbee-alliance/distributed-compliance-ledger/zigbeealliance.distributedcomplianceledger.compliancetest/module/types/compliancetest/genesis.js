/* eslint-disable */
import { TestingResults } from '../compliancetest/testing_results';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliancetest';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.testingResultsList) {
            TestingResults.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.testingResultsList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.testingResultsList.push(TestingResults.decode(reader, reader.uint32()));
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
        message.testingResultsList = [];
        if (object.testingResultsList !== undefined && object.testingResultsList !== null) {
            for (const e of object.testingResultsList) {
                message.testingResultsList.push(TestingResults.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.testingResultsList) {
            obj.testingResultsList = message.testingResultsList.map((e) => (e ? TestingResults.toJSON(e) : undefined));
        }
        else {
            obj.testingResultsList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.testingResultsList = [];
        if (object.testingResultsList !== undefined && object.testingResultsList !== null) {
            for (const e of object.testingResultsList) {
                message.testingResultsList.push(TestingResults.fromPartial(e));
            }
        }
        return message;
    }
};
