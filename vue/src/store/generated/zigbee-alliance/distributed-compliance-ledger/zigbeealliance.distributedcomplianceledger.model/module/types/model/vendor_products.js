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
        for (const v of message.products) {
            Product.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVendorProducts };
        message.products = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.products.push(Product.decode(reader, reader.uint32()));
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
        message.products = [];
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.products !== undefined && object.products !== null) {
            for (const e of object.products) {
                message.products.push(Product.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        if (message.products) {
            obj.products = message.products.map((e) => (e ? Product.toJSON(e) : undefined));
        }
        else {
            obj.products = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVendorProducts };
        message.products = [];
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.products !== undefined && object.products !== null) {
            for (const e of object.products) {
                message.products.push(Product.fromPartial(e));
            }
        }
        return message;
    }
};
