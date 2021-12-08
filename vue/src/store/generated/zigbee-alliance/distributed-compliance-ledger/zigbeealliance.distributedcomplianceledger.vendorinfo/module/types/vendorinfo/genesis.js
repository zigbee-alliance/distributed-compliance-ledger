/* eslint-disable */
import { NewVendorInfo } from '../vendorinfo/new_vendor_info';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.newVendorInfoList) {
            NewVendorInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.newVendorInfoList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.newVendorInfoList.push(NewVendorInfo.decode(reader, reader.uint32()));
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
        message.newVendorInfoList = [];
        if (object.newVendorInfoList !== undefined && object.newVendorInfoList !== null) {
            for (const e of object.newVendorInfoList) {
                message.newVendorInfoList.push(NewVendorInfo.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.newVendorInfoList) {
            obj.newVendorInfoList = message.newVendorInfoList.map((e) => (e ? NewVendorInfo.toJSON(e) : undefined));
        }
        else {
            obj.newVendorInfoList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.newVendorInfoList = [];
        if (object.newVendorInfoList !== undefined && object.newVendorInfoList !== null) {
            for (const e of object.newVendorInfoList) {
                message.newVendorInfoList.push(NewVendorInfo.fromPartial(e));
            }
        }
        return message;
    }
};
