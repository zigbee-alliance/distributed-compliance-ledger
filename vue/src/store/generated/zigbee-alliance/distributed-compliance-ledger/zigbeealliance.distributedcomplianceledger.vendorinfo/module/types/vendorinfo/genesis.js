/* eslint-disable */
import { VendorInfo } from '../vendorinfo/vendor_info';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vendorInfoList) {
            VendorInfo.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.vendorInfoList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorInfoList.push(VendorInfo.decode(reader, reader.uint32()));
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
        message.vendorInfoList = [];
        if (object.vendorInfoList !== undefined && object.vendorInfoList !== null) {
            for (const e of object.vendorInfoList) {
                message.vendorInfoList.push(VendorInfo.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.vendorInfoList) {
            obj.vendorInfoList = message.vendorInfoList.map((e) => (e ? VendorInfo.toJSON(e) : undefined));
        }
        else {
            obj.vendorInfoList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.vendorInfoList = [];
        if (object.vendorInfoList !== undefined && object.vendorInfoList !== null) {
            for (const e of object.vendorInfoList) {
                message.vendorInfoList.push(VendorInfo.fromPartial(e));
            }
        }
        return message;
    }
};
