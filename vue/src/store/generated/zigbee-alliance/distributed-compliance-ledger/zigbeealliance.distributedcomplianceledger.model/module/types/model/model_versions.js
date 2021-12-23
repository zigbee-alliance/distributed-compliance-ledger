/* eslint-disable */
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model';
const baseModelVersions = { vid: 0, pid: 0, softwareVersions: 0 };
export const ModelVersions = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        writer.uint32(26).fork();
        for (const v of message.softwareVersions) {
            writer.uint32(v);
        }
        writer.ldelim();
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseModelVersions };
        message.softwareVersions = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.vid = reader.int32();
                    break;
                case 2:
                    message.pid = reader.int32();
                    break;
                case 3:
                    if ((tag & 7) === 2) {
                        const end2 = reader.uint32() + reader.pos;
                        while (reader.pos < end2) {
                            message.softwareVersions.push(reader.uint32());
                        }
                    }
                    else {
                        message.softwareVersions.push(reader.uint32());
                    }
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseModelVersions };
        message.softwareVersions = [];
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = Number(object.vid);
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = Number(object.pid);
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersions !== undefined && object.softwareVersions !== null) {
            for (const e of object.softwareVersions) {
                message.softwareVersions.push(Number(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        if (message.softwareVersions) {
            obj.softwareVersions = message.softwareVersions.map((e) => e);
        }
        else {
            obj.softwareVersions = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseModelVersions };
        message.softwareVersions = [];
        if (object.vid !== undefined && object.vid !== null) {
            message.vid = object.vid;
        }
        else {
            message.vid = 0;
        }
        if (object.pid !== undefined && object.pid !== null) {
            message.pid = object.pid;
        }
        else {
            message.pid = 0;
        }
        if (object.softwareVersions !== undefined && object.softwareVersions !== null) {
            for (const e of object.softwareVersions) {
                message.softwareVersions.push(e);
            }
        }
        return message;
    }
};
