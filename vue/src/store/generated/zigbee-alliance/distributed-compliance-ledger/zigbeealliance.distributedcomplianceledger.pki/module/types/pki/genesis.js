/* eslint-disable */
import { ApprovedCertificates } from '../pki/approved_certificates';
import { ProposedCertificate } from '../pki/proposed_certificate';
import { ChildCertificates } from '../pki/child_certificates';
import { ProposedCertificateRevocation } from '../pki/proposed_certificate_revocation';
import { RevokedCertificates } from '../pki/revoked_certificates';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.approvedCertificatesList) {
            ApprovedCertificates.encode(v, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.proposedCertificateList) {
            ProposedCertificate.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.childCertificatesList) {
            ChildCertificates.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.proposedCertificateRevocationList) {
            ProposedCertificateRevocation.encode(v, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.revokedCertificatesList) {
            RevokedCertificates.encode(v, writer.uint32(42).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.approvedCertificatesList = [];
        message.proposedCertificateList = [];
        message.childCertificatesList = [];
        message.proposedCertificateRevocationList = [];
        message.revokedCertificatesList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.approvedCertificatesList.push(ApprovedCertificates.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.proposedCertificateList.push(ProposedCertificate.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.childCertificatesList.push(ChildCertificates.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.proposedCertificateRevocationList.push(ProposedCertificateRevocation.decode(reader, reader.uint32()));
                    break;
                case 5:
                    message.revokedCertificatesList.push(RevokedCertificates.decode(reader, reader.uint32()));
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
        message.approvedCertificatesList = [];
        message.proposedCertificateList = [];
        message.childCertificatesList = [];
        message.proposedCertificateRevocationList = [];
        message.revokedCertificatesList = [];
        if (object.approvedCertificatesList !== undefined && object.approvedCertificatesList !== null) {
            for (const e of object.approvedCertificatesList) {
                message.approvedCertificatesList.push(ApprovedCertificates.fromJSON(e));
            }
        }
        if (object.proposedCertificateList !== undefined && object.proposedCertificateList !== null) {
            for (const e of object.proposedCertificateList) {
                message.proposedCertificateList.push(ProposedCertificate.fromJSON(e));
            }
        }
        if (object.childCertificatesList !== undefined && object.childCertificatesList !== null) {
            for (const e of object.childCertificatesList) {
                message.childCertificatesList.push(ChildCertificates.fromJSON(e));
            }
        }
        if (object.proposedCertificateRevocationList !== undefined && object.proposedCertificateRevocationList !== null) {
            for (const e of object.proposedCertificateRevocationList) {
                message.proposedCertificateRevocationList.push(ProposedCertificateRevocation.fromJSON(e));
            }
        }
        if (object.revokedCertificatesList !== undefined && object.revokedCertificatesList !== null) {
            for (const e of object.revokedCertificatesList) {
                message.revokedCertificatesList.push(RevokedCertificates.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.approvedCertificatesList) {
            obj.approvedCertificatesList = message.approvedCertificatesList.map((e) => (e ? ApprovedCertificates.toJSON(e) : undefined));
        }
        else {
            obj.approvedCertificatesList = [];
        }
        if (message.proposedCertificateList) {
            obj.proposedCertificateList = message.proposedCertificateList.map((e) => (e ? ProposedCertificate.toJSON(e) : undefined));
        }
        else {
            obj.proposedCertificateList = [];
        }
        if (message.childCertificatesList) {
            obj.childCertificatesList = message.childCertificatesList.map((e) => (e ? ChildCertificates.toJSON(e) : undefined));
        }
        else {
            obj.childCertificatesList = [];
        }
        if (message.proposedCertificateRevocationList) {
            obj.proposedCertificateRevocationList = message.proposedCertificateRevocationList.map((e) => (e ? ProposedCertificateRevocation.toJSON(e) : undefined));
        }
        else {
            obj.proposedCertificateRevocationList = [];
        }
        if (message.revokedCertificatesList) {
            obj.revokedCertificatesList = message.revokedCertificatesList.map((e) => (e ? RevokedCertificates.toJSON(e) : undefined));
        }
        else {
            obj.revokedCertificatesList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.approvedCertificatesList = [];
        message.proposedCertificateList = [];
        message.childCertificatesList = [];
        message.proposedCertificateRevocationList = [];
        message.revokedCertificatesList = [];
        if (object.approvedCertificatesList !== undefined && object.approvedCertificatesList !== null) {
            for (const e of object.approvedCertificatesList) {
                message.approvedCertificatesList.push(ApprovedCertificates.fromPartial(e));
            }
        }
        if (object.proposedCertificateList !== undefined && object.proposedCertificateList !== null) {
            for (const e of object.proposedCertificateList) {
                message.proposedCertificateList.push(ProposedCertificate.fromPartial(e));
            }
        }
        if (object.childCertificatesList !== undefined && object.childCertificatesList !== null) {
            for (const e of object.childCertificatesList) {
                message.childCertificatesList.push(ChildCertificates.fromPartial(e));
            }
        }
        if (object.proposedCertificateRevocationList !== undefined && object.proposedCertificateRevocationList !== null) {
            for (const e of object.proposedCertificateRevocationList) {
                message.proposedCertificateRevocationList.push(ProposedCertificateRevocation.fromPartial(e));
            }
        }
        if (object.revokedCertificatesList !== undefined && object.revokedCertificatesList !== null) {
            for (const e of object.revokedCertificatesList) {
                message.revokedCertificatesList.push(RevokedCertificates.fromPartial(e));
            }
        }
        return message;
    }
};
