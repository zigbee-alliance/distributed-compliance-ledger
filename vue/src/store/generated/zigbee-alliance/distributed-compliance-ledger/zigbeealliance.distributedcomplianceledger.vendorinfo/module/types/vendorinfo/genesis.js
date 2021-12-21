/* eslint-disable */
import { VendorInfoType } from '../vendorinfo/vendor_info_type';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vendorInfoTypeList) {
            VendorInfoType.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.vendorInfoTypeList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorInfoTypeList.push(VendorInfoType.decode(reader, reader.uint32()));
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
        message.vendorInfoTypeList = [];
        if (object.vendorInfoTypeList !== undefined && object.vendorInfoTypeList !== null) {
            for (const e of object.vendorInfoTypeList) {
                message.vendorInfoTypeList.push(VendorInfoType.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.vendorInfoTypeList) {
            obj.vendorInfoTypeList = message.vendorInfoTypeList.map((e) => (e ? VendorInfoType.toJSON(e) : undefined));
        }
        else {
            obj.vendorInfoTypeList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.vendorInfoTypeList = [];
        if (object.vendorInfoTypeList !== undefined && object.vendorInfoTypeList !== null) {
            for (const e of object.vendorInfoTypeList) {
                message.vendorInfoTypeList.push(VendorInfoType.fromPartial(e));
            }
        }
        return message;
    }
};
