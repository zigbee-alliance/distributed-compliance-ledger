/* eslint-disable */
import { ComplianceHistoryItem } from '../compliance/compliance_history_item';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.compliance';
const baseComplianceInfo = {
    vid: 0,
    pid: 0,
    softwareVersion: 0,
    certificationType: '',
    softwareVersionString: '',
    cDVersionNumber: 0,
    softwareVersionCertificationStatus: 0,
    date: '',
    reason: '',
    owner: '',
    programTypeVersion: '',
    certificationID: '',
    familyID: '',
    supportedClusters: '',
    compliancePlatformUsed: '',
    compliancePlatformVersion: '',
    OSVersion: ''
};
export const ComplianceInfo = {
    encode(message, writer = Writer.create()) {
        if (message.vid !== 0) {
            writer.uint32(8).int32(message.vid);
        }
        if (message.pid !== 0) {
            writer.uint32(16).int32(message.pid);
        }
        if (message.softwareVersion !== 0) {
            writer.uint32(24).uint32(message.softwareVersion);
        }
        if (message.certificationType !== '') {
            writer.uint32(34).string(message.certificationType);
        }
        if (message.softwareVersionString !== '') {
            writer.uint32(42).string(message.softwareVersionString);
        }
        if (message.cDVersionNumber !== 0) {
            writer.uint32(48).uint32(message.cDVersionNumber);
        }
        if (message.softwareVersionCertificationStatus !== 0) {
            writer.uint32(56).uint32(message.softwareVersionCertificationStatus);
        }
        if (message.date !== '') {
            writer.uint32(66).string(message.date);
        }
        if (message.reason !== '') {
            writer.uint32(74).string(message.reason);
        }
        if (message.owner !== '') {
            writer.uint32(82).string(message.owner);
        }
        for (const v of message.history) {
            ComplianceHistoryItem.encode(v, writer.uint32(90).fork()).ldelim();
        }
        if (message.programTypeVersion !== '') {
            writer.uint32(98).string(message.programTypeVersion);
        }
        if (message.certificationID !== '') {
            writer.uint32(106).string(message.certificationID);
        }
        if (message.familyID !== '') {
            writer.uint32(114).string(message.familyID);
        }
        if (message.supportedClusters !== '') {
            writer.uint32(122).string(message.supportedClusters);
        }
        if (message.compliancePlatformUsed !== '') {
            writer.uint32(130).string(message.compliancePlatformUsed);
        }
        if (message.compliancePlatformVersion !== '') {
            writer.uint32(138).string(message.compliancePlatformVersion);
        }
        if (message.OSVersion !== '') {
            writer.uint32(146).string(message.OSVersion);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseComplianceInfo };
        message.history = [];
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
                    message.softwareVersion = reader.uint32();
                    break;
                case 4:
                    message.certificationType = reader.string();
                    break;
                case 5:
                    message.softwareVersionString = reader.string();
                    break;
                case 6:
                    message.cDVersionNumber = reader.uint32();
                    break;
                case 7:
                    message.softwareVersionCertificationStatus = reader.uint32();
                    break;
                case 8:
                    message.date = reader.string();
                    break;
                case 9:
                    message.reason = reader.string();
                    break;
                case 10:
                    message.owner = reader.string();
                    break;
                case 11:
                    message.history.push(ComplianceHistoryItem.decode(reader, reader.uint32()));
                    break;
                case 12:
                    message.programTypeVersion = reader.string();
                    break;
                case 13:
                    message.certificationID = reader.string();
                    break;
                case 14:
                    message.familyID = reader.string();
                    break;
                case 15:
                    message.supportedClusters = reader.string();
                    break;
                case 16:
                    message.compliancePlatformUsed = reader.string();
                    break;
                case 17:
                    message.compliancePlatformVersion = reader.string();
                    break;
                case 18:
                    message.OSVersion = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseComplianceInfo };
        message.history = [];
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
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = Number(object.softwareVersion);
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = String(object.certificationType);
        }
        else {
            message.certificationType = '';
        }
        if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
            message.softwareVersionString = String(object.softwareVersionString);
        }
        else {
            message.softwareVersionString = '';
        }
        if (object.cDVersionNumber !== undefined && object.cDVersionNumber !== null) {
            message.cDVersionNumber = Number(object.cDVersionNumber);
        }
        else {
            message.cDVersionNumber = 0;
        }
        if (object.softwareVersionCertificationStatus !== undefined && object.softwareVersionCertificationStatus !== null) {
            message.softwareVersionCertificationStatus = Number(object.softwareVersionCertificationStatus);
        }
        else {
            message.softwareVersionCertificationStatus = 0;
        }
        if (object.date !== undefined && object.date !== null) {
            message.date = String(object.date);
        }
        else {
            message.date = '';
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = String(object.reason);
        }
        else {
            message.reason = '';
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = '';
        }
        if (object.history !== undefined && object.history !== null) {
            for (const e of object.history) {
                message.history.push(ComplianceHistoryItem.fromJSON(e));
            }
        }
        if (object.programTypeVersion !== undefined && object.programTypeVersion !== null) {
            message.programTypeVersion = String(object.programTypeVersion);
        }
        else {
            message.programTypeVersion = '';
        }
        if (object.certificationID !== undefined && object.certificationID !== null) {
            message.certificationID = String(object.certificationID);
        }
        else {
            message.certificationID = '';
        }
        if (object.familyID !== undefined && object.familyID !== null) {
            message.familyID = String(object.familyID);
        }
        else {
            message.familyID = '';
        }
        if (object.supportedClusters !== undefined && object.supportedClusters !== null) {
            message.supportedClusters = String(object.supportedClusters);
        }
        else {
            message.supportedClusters = '';
        }
        if (object.compliancePlatformUsed !== undefined && object.compliancePlatformUsed !== null) {
            message.compliancePlatformUsed = String(object.compliancePlatformUsed);
        }
        else {
            message.compliancePlatformUsed = '';
        }
        if (object.compliancePlatformVersion !== undefined && object.compliancePlatformVersion !== null) {
            message.compliancePlatformVersion = String(object.compliancePlatformVersion);
        }
        else {
            message.compliancePlatformVersion = '';
        }
        if (object.OSVersion !== undefined && object.OSVersion !== null) {
            message.OSVersion = String(object.OSVersion);
        }
        else {
            message.OSVersion = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.vid !== undefined && (obj.vid = message.vid);
        message.pid !== undefined && (obj.pid = message.pid);
        message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion);
        message.certificationType !== undefined && (obj.certificationType = message.certificationType);
        message.softwareVersionString !== undefined && (obj.softwareVersionString = message.softwareVersionString);
        message.cDVersionNumber !== undefined && (obj.cDVersionNumber = message.cDVersionNumber);
        message.softwareVersionCertificationStatus !== undefined && (obj.softwareVersionCertificationStatus = message.softwareVersionCertificationStatus);
        message.date !== undefined && (obj.date = message.date);
        message.reason !== undefined && (obj.reason = message.reason);
        message.owner !== undefined && (obj.owner = message.owner);
        if (message.history) {
            obj.history = message.history.map((e) => (e ? ComplianceHistoryItem.toJSON(e) : undefined));
        }
        else {
            obj.history = [];
        }
        message.programTypeVersion !== undefined && (obj.programTypeVersion = message.programTypeVersion);
        message.certificationID !== undefined && (obj.certificationID = message.certificationID);
        message.familyID !== undefined && (obj.familyID = message.familyID);
        message.supportedClusters !== undefined && (obj.supportedClusters = message.supportedClusters);
        message.compliancePlatformUsed !== undefined && (obj.compliancePlatformUsed = message.compliancePlatformUsed);
        message.compliancePlatformVersion !== undefined && (obj.compliancePlatformVersion = message.compliancePlatformVersion);
        message.OSVersion !== undefined && (obj.OSVersion = message.OSVersion);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseComplianceInfo };
        message.history = [];
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
        if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
            message.softwareVersion = object.softwareVersion;
        }
        else {
            message.softwareVersion = 0;
        }
        if (object.certificationType !== undefined && object.certificationType !== null) {
            message.certificationType = object.certificationType;
        }
        else {
            message.certificationType = '';
        }
        if (object.softwareVersionString !== undefined && object.softwareVersionString !== null) {
            message.softwareVersionString = object.softwareVersionString;
        }
        else {
            message.softwareVersionString = '';
        }
        if (object.cDVersionNumber !== undefined && object.cDVersionNumber !== null) {
            message.cDVersionNumber = object.cDVersionNumber;
        }
        else {
            message.cDVersionNumber = 0;
        }
        if (object.softwareVersionCertificationStatus !== undefined && object.softwareVersionCertificationStatus !== null) {
            message.softwareVersionCertificationStatus = object.softwareVersionCertificationStatus;
        }
        else {
            message.softwareVersionCertificationStatus = 0;
        }
        if (object.date !== undefined && object.date !== null) {
            message.date = object.date;
        }
        else {
            message.date = '';
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = object.reason;
        }
        else {
            message.reason = '';
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = '';
        }
        if (object.history !== undefined && object.history !== null) {
            for (const e of object.history) {
                message.history.push(ComplianceHistoryItem.fromPartial(e));
            }
        }
        if (object.programTypeVersion !== undefined && object.programTypeVersion !== null) {
            message.programTypeVersion = object.programTypeVersion;
        }
        else {
            message.programTypeVersion = '';
        }
        if (object.certificationID !== undefined && object.certificationID !== null) {
            message.certificationID = object.certificationID;
        }
        else {
            message.certificationID = '';
        }
        if (object.familyID !== undefined && object.familyID !== null) {
            message.familyID = object.familyID;
        }
        else {
            message.familyID = '';
        }
        if (object.supportedClusters !== undefined && object.supportedClusters !== null) {
            message.supportedClusters = object.supportedClusters;
        }
        else {
            message.supportedClusters = '';
        }
        if (object.compliancePlatformUsed !== undefined && object.compliancePlatformUsed !== null) {
            message.compliancePlatformUsed = object.compliancePlatformUsed;
        }
        else {
            message.compliancePlatformUsed = '';
        }
        if (object.compliancePlatformVersion !== undefined && object.compliancePlatformVersion !== null) {
            message.compliancePlatformVersion = object.compliancePlatformVersion;
        }
        else {
            message.compliancePlatformVersion = '';
        }
        if (object.OSVersion !== undefined && object.OSVersion !== null) {
            message.OSVersion = object.OSVersion;
        }
        else {
            message.OSVersion = '';
        }
        return message;
    }
};
