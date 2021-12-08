/* eslint-disable */
import { VendorInfo } from '../vendorinfo/vendor_info';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseNewVendorInfo = { index: '', creator: '' };
export const NewVendorInfo = {
    encode(message, writer = Writer.create()) {
        if (message.index !== '') {
            writer.uint32(10).string(message.index);
        }
        if (message.vendorInfo !== undefined) {
            VendorInfo.encode(message.vendorInfo, writer.uint32(18).fork()).ldelim();
        }
        if (message.creator !== '') {
            writer.uint32(26).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNewVendorInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.index = reader.string();
                    break;
                case 2:
                    message.vendorInfo = VendorInfo.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseNewVendorInfo };
        if (object.index !== undefined && object.index !== null) {
            message.index = String(object.index);
        }
        else {
            message.index = '';
        }
        if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
            message.vendorInfo = VendorInfo.fromJSON(object.vendorInfo);
        }
        else {
            message.vendorInfo = undefined;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.index !== undefined && (obj.index = message.index);
        message.vendorInfo !== undefined && (obj.vendorInfo = message.vendorInfo ? VendorInfo.toJSON(message.vendorInfo) : undefined);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseNewVendorInfo };
        if (object.index !== undefined && object.index !== null) {
            message.index = object.index;
        }
        else {
            message.index = '';
        }
        if (object.vendorInfo !== undefined && object.vendorInfo !== null) {
            message.vendorInfo = VendorInfo.fromPartial(object.vendorInfo);
        }
        else {
            message.vendorInfo = undefined;
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
        return message;
    }
};
