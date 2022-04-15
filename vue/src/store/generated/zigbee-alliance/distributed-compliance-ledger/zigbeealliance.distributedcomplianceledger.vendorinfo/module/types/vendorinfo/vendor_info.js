/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseVendorInfo = { vendorID: 0, vendorName: '', companyLegalName: '', companyPreferredName: '', vendorLandingPageURL: '', creator: '' };
export const VendorInfo = {
    encode(message, writer = Writer.create()) {
        if (message.vendorID !== 0) {
            writer.uint32(8).int32(message.vendorID);
        }
        if (message.vendorName !== '') {
            writer.uint32(18).string(message.vendorName);
        }
        if (message.companyLegalName !== '') {
            writer.uint32(26).string(message.companyLegalName);
        }
        if (message.companyPreferredName !== '') {
            writer.uint32(34).string(message.companyPreferredName);
        }
        if (message.vendorLandingPageURL !== '') {
            writer.uint32(42).string(message.vendorLandingPageURL);
        }
        if (message.creator !== '') {
            writer.uint32(50).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseVendorInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorID = reader.int32();
                    break;
                case 2:
                    message.vendorName = reader.string();
                    break;
                case 3:
                    message.companyLegalName = reader.string();
                    break;
                case 4:
                    message.companyPreferredName = reader.string();
                    break;
                case 5:
                    message.vendorLandingPageURL = reader.string();
                    break;
                case 6:
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
        const message = { ...baseVendorInfo };
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = Number(object.vendorID);
        }
        else {
            message.vendorID = 0;
        }
        if (object.vendorName !== undefined && object.vendorName !== null) {
            message.vendorName = String(object.vendorName);
        }
        else {
            message.vendorName = '';
        }
        if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
            message.companyLegalName = String(object.companyLegalName);
        }
        else {
            message.companyLegalName = '';
        }
        if (object.companyPreferredName !== undefined && object.companyPreferredName !== null) {
            message.companyPreferredName = String(object.companyPreferredName);
        }
        else {
            message.companyPreferredName = '';
        }
        if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
            message.vendorLandingPageURL = String(object.vendorLandingPageURL);
        }
        else {
            message.vendorLandingPageURL = '';
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
        message.vendorID !== undefined && (obj.vendorID = message.vendorID);
        message.vendorName !== undefined && (obj.vendorName = message.vendorName);
        message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName);
        message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName);
        message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseVendorInfo };
        if (object.vendorID !== undefined && object.vendorID !== null) {
            message.vendorID = object.vendorID;
        }
        else {
            message.vendorID = 0;
        }
        if (object.vendorName !== undefined && object.vendorName !== null) {
            message.vendorName = object.vendorName;
        }
        else {
            message.vendorName = '';
        }
        if (object.companyLegalName !== undefined && object.companyLegalName !== null) {
            message.companyLegalName = object.companyLegalName;
        }
        else {
            message.companyLegalName = '';
        }
        if (object.companyPreferredName !== undefined && object.companyPreferredName !== null) {
            message.companyPreferredName = object.companyPreferredName;
        }
        else {
            message.companyPreferredName = '';
        }
        if (object.vendorLandingPageURL !== undefined && object.vendorLandingPageURL !== null) {
            message.vendorLandingPageURL = object.vendorLandingPageURL;
        }
        else {
            message.vendorLandingPageURL = '';
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
