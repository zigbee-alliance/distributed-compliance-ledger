/* eslint-disable */
import { Product } from '../model/product';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model';
const baseVendorProducts = { vid: 0 };
export const VendorProducts = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.products !== undefined) {
            Product.encode(message.products, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVendorProducts };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.products = Product.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseVendorProducts };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.products !== undefined && object.products !== null) {
            message.products = Product.fromJSON(object.products);
        }
        else {
            message.products = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.products !== undefined && (obj.products = message.products ? Product.toJSON(message.products) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVendorProducts };
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.products !== undefined && object.products !== null) {
            message.products = Product.fromPartial(object.products);
        }
        else {
            message.products = undefined;
        }
        return message;
    }
};
