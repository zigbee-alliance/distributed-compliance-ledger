/* eslint-disable */
import { Reader, util, configure, Writer } from 'protobufjs/minimal';
import * as Long from 'long';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.vendorinfo';
const baseMsgCreateVendorInfoType = {
    creator: '',
    vendorID: 0,
    vendorName: '',
    companyLegalName: '',
    companyPrefferedName: '',
    vendorLandingPageURL: ''
};
export const MsgCreateVendorInfoType = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.vendorID !== 0) {
            writer.uint32(16).uint64(message.vendorID);
        }
        if (message.vendorName !== '') {
            writer.uint32(26).string(message.vendorName);
        }
        if (message.companyLegalName !== '') {
            writer.uint32(34).string(message.companyLegalName);
        }
        if (message.companyPrefferedName !== '') {
            writer.uint32(42).string(message.companyPrefferedName);
        }
        if (message.vendorLandingPageURL !== '') {
            writer.uint32(50).string(message.vendorLandingPageURL);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateVendorInfoType };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.vendorID = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.vendorName = reader.string();
                    break;
                case 4:
                    message.companyLegalName = reader.string();
                    break;
                case 5:
                    message.companyPrefferedName = reader.string();
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
        const message = { ...baseMsgCreateVendorInfoType };
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
        if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
            message.companyPrefferedName = String(object.companyPrefferedName);
        }
        else {
            message.companyPrefferedName = '';
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
        message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName);
        message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateVendorInfoType };
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
        if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
            message.companyPrefferedName = object.companyPrefferedName;
        }
        else {
            message.companyPrefferedName = '';
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
const baseMsgCreateVendorInfoTypeResponse = {};
export const MsgCreateVendorInfoTypeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateVendorInfoTypeResponse };
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
        const message = { ...baseMsgCreateVendorInfoTypeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgCreateVendorInfoTypeResponse };
        return message;
    }
};
const baseMsgUpdateVendorInfoType = {
    creator: '',
    vendorID: 0,
    vendorName: '',
    companyLegalName: '',
    companyPrefferedName: '',
    vendorLandingPageURL: ''
};
export const MsgUpdateVendorInfoType = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.vendorID !== 0) {
            writer.uint32(16).uint64(message.vendorID);
        }
        if (message.vendorName !== '') {
            writer.uint32(26).string(message.vendorName);
        }
        if (message.companyLegalName !== '') {
            writer.uint32(34).string(message.companyLegalName);
        }
        if (message.companyPrefferedName !== '') {
            writer.uint32(42).string(message.companyPrefferedName);
        }
        if (message.vendorLandingPageURL !== '') {
            writer.uint32(50).string(message.vendorLandingPageURL);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateVendorInfoType };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.vendorID = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.vendorName = reader.string();
                    break;
                case 4:
                    message.companyLegalName = reader.string();
                    break;
                case 5:
                    message.companyPrefferedName = reader.string();
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
        const message = { ...baseMsgUpdateVendorInfoType };
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
        if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
            message.companyPrefferedName = String(object.companyPrefferedName);
        }
        else {
            message.companyPrefferedName = '';
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
        message.companyPrefferedName !== undefined && (obj.companyPrefferedName = message.companyPrefferedName);
        message.vendorLandingPageURL !== undefined && (obj.vendorLandingPageURL = message.vendorLandingPageURL);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgUpdateVendorInfoType };
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
        if (object.companyPrefferedName !== undefined && object.companyPrefferedName !== null) {
            message.companyPrefferedName = object.companyPrefferedName;
        }
        else {
            message.companyPrefferedName = '';
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
const baseMsgUpdateVendorInfoTypeResponse = {};
export const MsgUpdateVendorInfoTypeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateVendorInfoTypeResponse };
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
        const message = { ...baseMsgUpdateVendorInfoTypeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgUpdateVendorInfoTypeResponse };
        return message;
    }
};
const baseMsgDeleteVendorInfoType = { creator: '', vendorID: 0 };
export const MsgDeleteVendorInfoType = {
    encode(message, writer = Writer.create()) {
        if (message.creator !== '') {
            writer.uint32(10).string(message.creator);
        }
        if (message.vendorID !== 0) {
            writer.uint32(16).uint64(message.vendorID);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteVendorInfoType };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.vendorID = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseMsgDeleteVendorInfoType };
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
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.vendorID !== undefined && (obj.vendorID = message.vendorID);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDeleteVendorInfoType };
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
        return message;
    }
};
const baseMsgDeleteVendorInfoTypeResponse = {};
export const MsgDeleteVendorInfoTypeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDeleteVendorInfoTypeResponse };
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
        const message = { ...baseMsgDeleteVendorInfoTypeResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDeleteVendorInfoTypeResponse };
        return message;
    }
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateVendorInfoType(request) {
        const data = MsgCreateVendorInfoType.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'CreateVendorInfoType', data);
        return promise.then((data) => MsgCreateVendorInfoTypeResponse.decode(new Reader(data)));
    }
    UpdateVendorInfoType(request) {
        const data = MsgUpdateVendorInfoType.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'UpdateVendorInfoType', data);
        return promise.then((data) => MsgUpdateVendorInfoTypeResponse.decode(new Reader(data)));
    }
    DeleteVendorInfoType(request) {
        const data = MsgDeleteVendorInfoType.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.vendorinfo.Msg', 'DeleteVendorInfoType', data);
        return promise.then((data) => MsgDeleteVendorInfoTypeResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
