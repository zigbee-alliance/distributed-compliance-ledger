/* eslint-disable */
import { VendorProducts } from '../model/vendor_products';
import { Model } from '../model/model';
import { ModelVersion } from '../model/model_version';
import { ModelVersions } from '../model/model_versions';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.vendorProductsList) {
            VendorProducts.encode(v, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.modelList) {
            Model.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.modelVersionList) {
            ModelVersion.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.modelVersionsList) {
            ModelVersions.encode(v, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.vendorProductsList = [];
        message.modelList = [];
        message.modelVersionList = [];
        message.modelVersionsList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vendorProductsList.push(VendorProducts.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.modelList.push(Model.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.modelVersionList.push(ModelVersion.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.modelVersionsList.push(ModelVersions.decode(reader, reader.uint32()));
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
        message.vendorProductsList = [];
        message.modelList = [];
        message.modelVersionList = [];
        message.modelVersionsList = [];
        if (object.vendorProductsList !== undefined && object.vendorProductsList !== null) {
            for (const e of object.vendorProductsList) {
                message.vendorProductsList.push(VendorProducts.fromJSON(e));
            }
        }
        if (object.modelList !== undefined && object.modelList !== null) {
            for (const e of object.modelList) {
                message.modelList.push(Model.fromJSON(e));
            }
        }
        if (object.modelVersionList !== undefined && object.modelVersionList !== null) {
            for (const e of object.modelVersionList) {
                message.modelVersionList.push(ModelVersion.fromJSON(e));
            }
        }
        if (object.modelVersionsList !== undefined && object.modelVersionsList !== null) {
            for (const e of object.modelVersionsList) {
                message.modelVersionsList.push(ModelVersions.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.vendorProductsList) {
            obj.vendorProductsList = message.vendorProductsList.map((e) => (e ? VendorProducts.toJSON(e) : undefined));
        }
        else {
            obj.vendorProductsList = [];
        }
        if (message.modelList) {
            obj.modelList = message.modelList.map((e) => (e ? Model.toJSON(e) : undefined));
        }
        else {
            obj.modelList = [];
        }
        if (message.modelVersionList) {
            obj.modelVersionList = message.modelVersionList.map((e) => (e ? ModelVersion.toJSON(e) : undefined));
        }
        else {
            obj.modelVersionList = [];
        }
        if (message.modelVersionsList) {
            obj.modelVersionsList = message.modelVersionsList.map((e) => (e ? ModelVersions.toJSON(e) : undefined));
        }
        else {
            obj.modelVersionsList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.vendorProductsList = [];
        message.modelList = [];
        message.modelVersionList = [];
        message.modelVersionsList = [];
        if (object.vendorProductsList !== undefined && object.vendorProductsList !== null) {
            for (const e of object.vendorProductsList) {
                message.vendorProductsList.push(VendorProducts.fromPartial(e));
            }
        }
        if (object.modelList !== undefined && object.modelList !== null) {
            for (const e of object.modelList) {
                message.modelList.push(Model.fromPartial(e));
            }
        }
        if (object.modelVersionList !== undefined && object.modelVersionList !== null) {
            for (const e of object.modelVersionList) {
                message.modelVersionList.push(ModelVersion.fromPartial(e));
            }
        }
        if (object.modelVersionsList !== undefined && object.modelVersionsList !== null) {
            for (const e of object.modelVersionsList) {
                message.modelVersionsList.push(ModelVersions.fromPartial(e));
            }
        }
        return message;
    }
};
