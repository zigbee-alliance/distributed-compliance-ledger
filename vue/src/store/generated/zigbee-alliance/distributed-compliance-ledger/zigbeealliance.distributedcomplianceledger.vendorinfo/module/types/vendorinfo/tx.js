/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseMsgCreateVendorInfo = { creator: '', vendorID: 0, vendorName: '', companyLegalName: '', companyPreferredName: '', vendorLandingPageURL: '' };
export const MsgCreateVendorInfo = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.vendorID !== 0) {
            writer.uint32(16).int32(message.vendorID);
        }
        if (message.vendorName !== '') {
            writer.uint32(26).string(message.vendorName);
        }
        if (message.companyLegalName !== '') {
            writer.uint32(34).string(message.companyLegalName);
        }
        if (message.companyPreferredName !== '') {
            writer.uint32(42).string(message.companyPreferredName);
        }
        if (message.vendorLandingPageURL !== '') {
            writer.uint32(50).string(message.vendorLandingPageURL);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateVendorInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.vendorID = reader.int32();
                    break;
                case 3:
                    message.vendorName = reader.string();
                    break;
                case 4:
                    message.companyLegalName = reader.string();
                    break;
                case 5:
                    message.companyPreferredName = reader.string();
                    break;
                case 6:
                    message.vendorLandingPageURL = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgCreateVendorInfo };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.vendorID !== undefined && (obj.vendorID = message.vendorID);
        message.vendorName !== undefined && (obj.vendorName = message.vendorName);
        message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName);
        message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName);
        message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateVendorInfo };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
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
        return message;
    }
};
const baseMsgCreateVendorInfoResponse = {};
export const MsgCreateVendorInfoResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateVendorInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgCreateVendorInfoResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgCreateVendorInfoResponse };
        return message;
    }
};
const baseMsgUpdateVendorInfo = { creator: '', vendorID: 0, vendorName: '', companyLegalName: '', companyPreferredName: '', vendorLandingPageURL: '' };
export const MsgUpdateVendorInfo = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.vendorID !== 0) {
            writer.uint32(16).int32(message.vendorID);
        }
        if (message.vendorName !== '') {
            writer.uint32(26).string(message.vendorName);
        }
        if (message.companyLegalName !== '') {
            writer.uint32(34).string(message.companyLegalName);
        }
        if (message.companyPreferredName !== '') {
            writer.uint32(42).string(message.companyPreferredName);
        }
        if (message.vendorLandingPageURL !== '') {
            writer.uint32(50).string(message.vendorLandingPageURL);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateVendorInfo };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.vendorID = reader.int32();
                    break;
                case 3:
                    message.vendorName = reader.string();
                    break;
                case 4:
                    message.companyLegalName = reader.string();
                    break;
                case 5:
                    message.companyPreferredName = reader.string();
                    break;
                case 6:
                    message.vendorLandingPageURL = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgUpdateVendorInfo };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = '';
        }
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.vendorID !== undefined && (obj.vendorID = message.vendorID);
        message.vendorName !== undefined && (obj.vendorName = message.vendorName);
        message.companyLegalName !== undefined && (obj.companyLegalName = message.companyLegalName);
        message.companyPreferredName !== undefined && (obj.companyPreferredName = message.companyPreferredName);
        message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateVendorInfo };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = '';
        }
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
        return message;
    }
};
const baseMsgUpdateVendorInfoResponse = {};
export const MsgUpdateVendorInfoResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateVendorInfoResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseMsgUpdateVendorInfoResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUpdateVendorInfoResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateVendorInfo(request) {
        const data = MsgCreateVendorInfo.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'CreateVendorInfo', data);
        return promise.then((data) => MsgCreateVendorInfoResponse.decode(new Reader(data)));
    }
    UpdateVendorInfo(request) {
        const data = MsgUpdateVendorInfo.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'UpdateVendorInfo', data);
        return promise.then((data) => MsgUpdateVendorInfoResponse.decode(new Reader(data)));
    }
}
