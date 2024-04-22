/* eslint-disable */
import { ApprovedCertificates } from '../pki/approved_certificates'
import { ProposedCertificate } from '../pki/proposed_certificate'
import { ChildCertificates } from '../pki/child_certificates'
import { ProposedCertificateRevocation } from '../pki/proposed_certificate_revocation'
import { RevokedCertificates } from '../pki/revoked_certificates'
import { UniqueCertificate } from '../pki/unique_certificate'
import { ApprovedRootCertificates } from '../pki/approved_root_certificates'
import { RevokedRootCertificates } from '../pki/revoked_root_certificates'
import { ApprovedCertificatesBySubject } from '../pki/approved_certificates_by_subject'
import { RejectedCertificate } from '../pki/rejected_certificate'
import { PkiRevocationDistributionPoint } from '../pki/pki_revocation_distribution_point'
import { PkiRevocationDistributionPointsByIssuerSubjectKeyID } from '../pki/pki_revocation_distribution_points_by_issuer_subject_key_id'
import { ApprovedCertificatesBySubjectKeyId } from '../pki/approved_certificates_by_subject_key_id'
import { NocRootCertificates } from '../pki/noc_root_certificates'
import { NocIcaCertificates } from '../pki/noc_ica_certificates'
import { RevokedNocRootCertificates } from '../pki/revoked_noc_root_certificates'
import { NocRootCertificatesByVidAndSkid } from '../pki/noc_root_certificates_by_vid_and_skid'
import { Writer, Reader } from 'protobufjs/minimal'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.pki'

/** GenesisState defines the pki module's genesis state. */
export interface GenesisState {
  approvedCertificatesList: ApprovedCertificates[]
  proposedCertificateList: ProposedCertificate[]
  childCertificatesList: ChildCertificates[]
  proposedCertificateRevocationList: ProposedCertificateRevocation[]
  revokedCertificatesList: RevokedCertificates[]
  uniqueCertificateList: UniqueCertificate[]
  approvedRootCertificates: ApprovedRootCertificates | undefined
  revokedRootCertificates: RevokedRootCertificates | undefined
  approvedCertificatesBySubjectList: ApprovedCertificatesBySubject[]
  rejectedCertificateList: RejectedCertificate[]
  PkiRevocationDistributionPointList: PkiRevocationDistributionPoint[]
  pkiRevocationDistributionPointsByIssuerSubjectKeyIDList: PkiRevocationDistributionPointsByIssuerSubjectKeyID[]
  approvedCertificatesBySubjectKeyIdList: ApprovedCertificatesBySubjectKeyId[]
  nocRootCertificatesList: NocRootCertificates[]
  nocIcaCertificatesList: NocIcaCertificates[]
  revokedNocRootCertificatesList: RevokedNocRootCertificates[]
  /** this line is used by starport scaffolding # genesis/proto/state */
  nocRootCertificatesByVidAndSkidList: NocRootCertificatesByVidAndSkid[]
}

const baseGenesisState: object = {}

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.approvedCertificatesList) {
      ApprovedCertificates.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    for (const v of message.proposedCertificateList) {
      ProposedCertificate.encode(v!, writer.uint32(18).fork()).ldelim()
    }
    for (const v of message.childCertificatesList) {
      ChildCertificates.encode(v!, writer.uint32(26).fork()).ldelim()
    }
    for (const v of message.proposedCertificateRevocationList) {
      ProposedCertificateRevocation.encode(v!, writer.uint32(34).fork()).ldelim()
    }
    for (const v of message.revokedCertificatesList) {
      RevokedCertificates.encode(v!, writer.uint32(42).fork()).ldelim()
    }
    for (const v of message.uniqueCertificateList) {
      UniqueCertificate.encode(v!, writer.uint32(50).fork()).ldelim()
    }
    if (message.approvedRootCertificates !== undefined) {
      ApprovedRootCertificates.encode(message.approvedRootCertificates, writer.uint32(58).fork()).ldelim()
    }
    if (message.revokedRootCertificates !== undefined) {
      RevokedRootCertificates.encode(message.revokedRootCertificates, writer.uint32(66).fork()).ldelim()
    }
    for (const v of message.approvedCertificatesBySubjectList) {
      ApprovedCertificatesBySubject.encode(v!, writer.uint32(74).fork()).ldelim()
    }
    for (const v of message.rejectedCertificateList) {
      RejectedCertificate.encode(v!, writer.uint32(82).fork()).ldelim()
    }
    for (const v of message.PkiRevocationDistributionPointList) {
      PkiRevocationDistributionPoint.encode(v!, writer.uint32(90).fork()).ldelim()
    }
    for (const v of message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList) {
      PkiRevocationDistributionPointsByIssuerSubjectKeyID.encode(v!, writer.uint32(98).fork()).ldelim()
    }
    for (const v of message.approvedCertificatesBySubjectKeyIdList) {
      ApprovedCertificatesBySubjectKeyId.encode(v!, writer.uint32(106).fork()).ldelim()
    }
    for (const v of message.nocRootCertificatesList) {
      NocRootCertificates.encode(v!, writer.uint32(114).fork()).ldelim()
    }
    for (const v of message.nocIcaCertificatesList) {
      NocIcaCertificates.encode(v!, writer.uint32(122).fork()).ldelim()
    }
    for (const v of message.revokedNocRootCertificatesList) {
      RevokedNocRootCertificates.encode(v!, writer.uint32(130).fork()).ldelim()
    }
    for (const v of message.nocRootCertificatesByVidAndSkidList) {
      NocRootCertificatesByVidAndSkid.encode(v!, writer.uint32(138).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseGenesisState } as GenesisState
    message.approvedCertificatesList = []
    message.proposedCertificateList = []
    message.childCertificatesList = []
    message.proposedCertificateRevocationList = []
    message.revokedCertificatesList = []
    message.uniqueCertificateList = []
    message.approvedCertificatesBySubjectList = []
    message.rejectedCertificateList = []
    message.PkiRevocationDistributionPointList = []
    message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList = []
    message.approvedCertificatesBySubjectKeyIdList = []
    message.nocRootCertificatesList = []
    message.nocIcaCertificatesList = []
    message.revokedNocRootCertificatesList = []
    message.nocRootCertificatesByVidAndSkidList = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificatesList.push(ApprovedCertificates.decode(reader, reader.uint32()))
          break
        case 2:
          message.proposedCertificateList.push(ProposedCertificate.decode(reader, reader.uint32()))
          break
        case 3:
          message.childCertificatesList.push(ChildCertificates.decode(reader, reader.uint32()))
          break
        case 4:
          message.proposedCertificateRevocationList.push(ProposedCertificateRevocation.decode(reader, reader.uint32()))
          break
        case 5:
          message.revokedCertificatesList.push(RevokedCertificates.decode(reader, reader.uint32()))
          break
        case 6:
          message.uniqueCertificateList.push(UniqueCertificate.decode(reader, reader.uint32()))
          break
        case 7:
          message.approvedRootCertificates = ApprovedRootCertificates.decode(reader, reader.uint32())
          break
        case 8:
          message.revokedRootCertificates = RevokedRootCertificates.decode(reader, reader.uint32())
          break
        case 9:
          message.approvedCertificatesBySubjectList.push(ApprovedCertificatesBySubject.decode(reader, reader.uint32()))
          break
        case 10:
          message.rejectedCertificateList.push(RejectedCertificate.decode(reader, reader.uint32()))
          break
        case 11:
          message.PkiRevocationDistributionPointList.push(PkiRevocationDistributionPoint.decode(reader, reader.uint32()))
          break
        case 12:
          message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList.push(
            PkiRevocationDistributionPointsByIssuerSubjectKeyID.decode(reader, reader.uint32())
          )
          break
        case 13:
          message.approvedCertificatesBySubjectKeyIdList.push(ApprovedCertificatesBySubjectKeyId.decode(reader, reader.uint32()))
          break
        case 14:
          message.nocRootCertificatesList.push(NocRootCertificates.decode(reader, reader.uint32()))
          break
        case 15:
          message.nocIcaCertificatesList.push(NocIcaCertificates.decode(reader, reader.uint32()))
          break
        case 16:
          message.revokedNocRootCertificatesList.push(RevokedNocRootCertificates.decode(reader, reader.uint32()))
          break
        case 17:
          message.nocRootCertificatesByVidAndSkidList.push(NocRootCertificatesByVidAndSkid.decode(reader, reader.uint32()))
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.approvedCertificatesList = []
    message.proposedCertificateList = []
    message.childCertificatesList = []
    message.proposedCertificateRevocationList = []
    message.revokedCertificatesList = []
    message.uniqueCertificateList = []
    message.approvedCertificatesBySubjectList = []
    message.rejectedCertificateList = []
    message.PkiRevocationDistributionPointList = []
    message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList = []
    message.approvedCertificatesBySubjectKeyIdList = []
    message.nocRootCertificatesList = []
    message.nocIcaCertificatesList = []
    message.revokedNocRootCertificatesList = []
    message.nocRootCertificatesByVidAndSkidList = []
    if (object.approvedCertificatesList !== undefined && object.approvedCertificatesList !== null) {
      for (const e of object.approvedCertificatesList) {
        message.approvedCertificatesList.push(ApprovedCertificates.fromJSON(e))
      }
    }
    if (object.proposedCertificateList !== undefined && object.proposedCertificateList !== null) {
      for (const e of object.proposedCertificateList) {
        message.proposedCertificateList.push(ProposedCertificate.fromJSON(e))
      }
    }
    if (object.childCertificatesList !== undefined && object.childCertificatesList !== null) {
      for (const e of object.childCertificatesList) {
        message.childCertificatesList.push(ChildCertificates.fromJSON(e))
      }
    }
    if (object.proposedCertificateRevocationList !== undefined && object.proposedCertificateRevocationList !== null) {
      for (const e of object.proposedCertificateRevocationList) {
        message.proposedCertificateRevocationList.push(ProposedCertificateRevocation.fromJSON(e))
      }
    }
    if (object.revokedCertificatesList !== undefined && object.revokedCertificatesList !== null) {
      for (const e of object.revokedCertificatesList) {
        message.revokedCertificatesList.push(RevokedCertificates.fromJSON(e))
      }
    }
    if (object.uniqueCertificateList !== undefined && object.uniqueCertificateList !== null) {
      for (const e of object.uniqueCertificateList) {
        message.uniqueCertificateList.push(UniqueCertificate.fromJSON(e))
      }
    }
    if (object.approvedRootCertificates !== undefined && object.approvedRootCertificates !== null) {
      message.approvedRootCertificates = ApprovedRootCertificates.fromJSON(object.approvedRootCertificates)
    } else {
      message.approvedRootCertificates = undefined
    }
    if (object.revokedRootCertificates !== undefined && object.revokedRootCertificates !== null) {
      message.revokedRootCertificates = RevokedRootCertificates.fromJSON(object.revokedRootCertificates)
    } else {
      message.revokedRootCertificates = undefined
    }
    if (object.approvedCertificatesBySubjectList !== undefined && object.approvedCertificatesBySubjectList !== null) {
      for (const e of object.approvedCertificatesBySubjectList) {
        message.approvedCertificatesBySubjectList.push(ApprovedCertificatesBySubject.fromJSON(e))
      }
    }
    if (object.rejectedCertificateList !== undefined && object.rejectedCertificateList !== null) {
      for (const e of object.rejectedCertificateList) {
        message.rejectedCertificateList.push(RejectedCertificate.fromJSON(e))
      }
    }
    if (object.PkiRevocationDistributionPointList !== undefined && object.PkiRevocationDistributionPointList !== null) {
      for (const e of object.PkiRevocationDistributionPointList) {
        message.PkiRevocationDistributionPointList.push(PkiRevocationDistributionPoint.fromJSON(e))
      }
    }
    if (
      object.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList !== undefined &&
      object.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList !== null
    ) {
      for (const e of object.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList) {
        message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList.push(PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromJSON(e))
      }
    }
    if (object.approvedCertificatesBySubjectKeyIdList !== undefined && object.approvedCertificatesBySubjectKeyIdList !== null) {
      for (const e of object.approvedCertificatesBySubjectKeyIdList) {
        message.approvedCertificatesBySubjectKeyIdList.push(ApprovedCertificatesBySubjectKeyId.fromJSON(e))
      }
    }
    if (object.nocRootCertificatesList !== undefined && object.nocRootCertificatesList !== null) {
      for (const e of object.nocRootCertificatesList) {
        message.nocRootCertificatesList.push(NocRootCertificates.fromJSON(e))
      }
    }
    if (object.nocIcaCertificatesList !== undefined && object.nocIcaCertificatesList !== null) {
      for (const e of object.nocIcaCertificatesList) {
        message.nocIcaCertificatesList.push(NocIcaCertificates.fromJSON(e))
      }
    }
    if (object.revokedNocRootCertificatesList !== undefined && object.revokedNocRootCertificatesList !== null) {
      for (const e of object.revokedNocRootCertificatesList) {
        message.revokedNocRootCertificatesList.push(RevokedNocRootCertificates.fromJSON(e))
      }
    }
    if (object.nocRootCertificatesByVidAndSkidList !== undefined && object.nocRootCertificatesByVidAndSkidList !== null) {
      for (const e of object.nocRootCertificatesByVidAndSkidList) {
        message.nocRootCertificatesByVidAndSkidList.push(NocRootCertificatesByVidAndSkid.fromJSON(e))
      }
    }
    return message
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {}
    if (message.approvedCertificatesList) {
      obj.approvedCertificatesList = message.approvedCertificatesList.map((e) => (e ? ApprovedCertificates.toJSON(e) : undefined))
    } else {
      obj.approvedCertificatesList = []
    }
    if (message.proposedCertificateList) {
      obj.proposedCertificateList = message.proposedCertificateList.map((e) => (e ? ProposedCertificate.toJSON(e) : undefined))
    } else {
      obj.proposedCertificateList = []
    }
    if (message.childCertificatesList) {
      obj.childCertificatesList = message.childCertificatesList.map((e) => (e ? ChildCertificates.toJSON(e) : undefined))
    } else {
      obj.childCertificatesList = []
    }
    if (message.proposedCertificateRevocationList) {
      obj.proposedCertificateRevocationList = message.proposedCertificateRevocationList.map((e) => (e ? ProposedCertificateRevocation.toJSON(e) : undefined))
    } else {
      obj.proposedCertificateRevocationList = []
    }
    if (message.revokedCertificatesList) {
      obj.revokedCertificatesList = message.revokedCertificatesList.map((e) => (e ? RevokedCertificates.toJSON(e) : undefined))
    } else {
      obj.revokedCertificatesList = []
    }
    if (message.uniqueCertificateList) {
      obj.uniqueCertificateList = message.uniqueCertificateList.map((e) => (e ? UniqueCertificate.toJSON(e) : undefined))
    } else {
      obj.uniqueCertificateList = []
    }
    message.approvedRootCertificates !== undefined &&
      (obj.approvedRootCertificates = message.approvedRootCertificates ? ApprovedRootCertificates.toJSON(message.approvedRootCertificates) : undefined)
    message.revokedRootCertificates !== undefined &&
      (obj.revokedRootCertificates = message.revokedRootCertificates ? RevokedRootCertificates.toJSON(message.revokedRootCertificates) : undefined)
    if (message.approvedCertificatesBySubjectList) {
      obj.approvedCertificatesBySubjectList = message.approvedCertificatesBySubjectList.map((e) => (e ? ApprovedCertificatesBySubject.toJSON(e) : undefined))
    } else {
      obj.approvedCertificatesBySubjectList = []
    }
    if (message.rejectedCertificateList) {
      obj.rejectedCertificateList = message.rejectedCertificateList.map((e) => (e ? RejectedCertificate.toJSON(e) : undefined))
    } else {
      obj.rejectedCertificateList = []
    }
    if (message.PkiRevocationDistributionPointList) {
      obj.PkiRevocationDistributionPointList = message.PkiRevocationDistributionPointList.map((e) => (e ? PkiRevocationDistributionPoint.toJSON(e) : undefined))
    } else {
      obj.PkiRevocationDistributionPointList = []
    }
    if (message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList) {
      obj.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList = message.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList.map((e) =>
        e ? PkiRevocationDistributionPointsByIssuerSubjectKeyID.toJSON(e) : undefined
      )
    } else {
      obj.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList = []
    }
    if (message.approvedCertificatesBySubjectKeyIdList) {
      obj.approvedCertificatesBySubjectKeyIdList = message.approvedCertificatesBySubjectKeyIdList.map((e) =>
        e ? ApprovedCertificatesBySubjectKeyId.toJSON(e) : undefined
      )
    } else {
      obj.approvedCertificatesBySubjectKeyIdList = []
    }
    if (message.nocRootCertificatesList) {
      obj.nocRootCertificatesList = message.nocRootCertificatesList.map((e) => (e ? NocRootCertificates.toJSON(e) : undefined))
    } else {
      obj.nocRootCertificatesList = []
    }
    if (message.nocIcaCertificatesList) {
      obj.nocIcaCertificatesList = message.nocIcaCertificatesList.map((e) => (e ? NocIcaCertificates.toJSON(e) : undefined))
    } else {
      obj.nocIcaCertificatesList = []
    }
    if (message.revokedNocRootCertificatesList) {
      obj.revokedNocRootCertificatesList = message.revokedNocRootCertificatesList.map((e) => (e ? RevokedNocRootCertificates.toJSON(e) : undefined))
    } else {
      obj.revokedNocRootCertificatesList = []
    }
    if (message.nocRootCertificatesByVidAndSkidList) {
      obj.nocRootCertificatesByVidAndSkidList = message.nocRootCertificatesByVidAndSkidList.map((e) =>
        e ? NocRootCertificatesByVidAndSkid.toJSON(e) : undefined
      )
    } else {
      obj.nocRootCertificatesByVidAndSkidList = []
    }
    return obj
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState
    message.approvedCertificatesList = []
    message.proposedCertificateList = []
    message.childCertificatesList = []
    message.proposedCertificateRevocationList = []
    message.revokedCertificatesList = []
    message.uniqueCertificateList = []
    message.approvedCertificatesBySubjectList = []
    message.rejectedCertificateList = []
    message.PkiRevocationDistributionPointList = []
    message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList = []
    message.approvedCertificatesBySubjectKeyIdList = []
    message.nocRootCertificatesList = []
    message.nocIcaCertificatesList = []
    message.revokedNocRootCertificatesList = []
    message.nocRootCertificatesByVidAndSkidList = []
    if (object.approvedCertificatesList !== undefined && object.approvedCertificatesList !== null) {
      for (const e of object.approvedCertificatesList) {
        message.approvedCertificatesList.push(ApprovedCertificates.fromPartial(e))
      }
    }
    if (object.proposedCertificateList !== undefined && object.proposedCertificateList !== null) {
      for (const e of object.proposedCertificateList) {
        message.proposedCertificateList.push(ProposedCertificate.fromPartial(e))
      }
    }
    if (object.childCertificatesList !== undefined && object.childCertificatesList !== null) {
      for (const e of object.childCertificatesList) {
        message.childCertificatesList.push(ChildCertificates.fromPartial(e))
      }
    }
    if (object.proposedCertificateRevocationList !== undefined && object.proposedCertificateRevocationList !== null) {
      for (const e of object.proposedCertificateRevocationList) {
        message.proposedCertificateRevocationList.push(ProposedCertificateRevocation.fromPartial(e))
      }
    }
    if (object.revokedCertificatesList !== undefined && object.revokedCertificatesList !== null) {
      for (const e of object.revokedCertificatesList) {
        message.revokedCertificatesList.push(RevokedCertificates.fromPartial(e))
      }
    }
    if (object.uniqueCertificateList !== undefined && object.uniqueCertificateList !== null) {
      for (const e of object.uniqueCertificateList) {
        message.uniqueCertificateList.push(UniqueCertificate.fromPartial(e))
      }
    }
    if (object.approvedRootCertificates !== undefined && object.approvedRootCertificates !== null) {
      message.approvedRootCertificates = ApprovedRootCertificates.fromPartial(object.approvedRootCertificates)
    } else {
      message.approvedRootCertificates = undefined
    }
    if (object.revokedRootCertificates !== undefined && object.revokedRootCertificates !== null) {
      message.revokedRootCertificates = RevokedRootCertificates.fromPartial(object.revokedRootCertificates)
    } else {
      message.revokedRootCertificates = undefined
    }
    if (object.approvedCertificatesBySubjectList !== undefined && object.approvedCertificatesBySubjectList !== null) {
      for (const e of object.approvedCertificatesBySubjectList) {
        message.approvedCertificatesBySubjectList.push(ApprovedCertificatesBySubject.fromPartial(e))
      }
    }
    if (object.rejectedCertificateList !== undefined && object.rejectedCertificateList !== null) {
      for (const e of object.rejectedCertificateList) {
        message.rejectedCertificateList.push(RejectedCertificate.fromPartial(e))
      }
    }
    if (object.PkiRevocationDistributionPointList !== undefined && object.PkiRevocationDistributionPointList !== null) {
      for (const e of object.PkiRevocationDistributionPointList) {
        message.PkiRevocationDistributionPointList.push(PkiRevocationDistributionPoint.fromPartial(e))
      }
    }
    if (
      object.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList !== undefined &&
      object.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList !== null
    ) {
      for (const e of object.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList) {
        message.pkiRevocationDistributionPointsByIssuerSubjectKeyIDList.push(PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromPartial(e))
      }
    }
    if (object.approvedCertificatesBySubjectKeyIdList !== undefined && object.approvedCertificatesBySubjectKeyIdList !== null) {
      for (const e of object.approvedCertificatesBySubjectKeyIdList) {
        message.approvedCertificatesBySubjectKeyIdList.push(ApprovedCertificatesBySubjectKeyId.fromPartial(e))
      }
    }
    if (object.nocRootCertificatesList !== undefined && object.nocRootCertificatesList !== null) {
      for (const e of object.nocRootCertificatesList) {
        message.nocRootCertificatesList.push(NocRootCertificates.fromPartial(e))
      }
    }
    if (object.nocIcaCertificatesList !== undefined && object.nocIcaCertificatesList !== null) {
      for (const e of object.nocIcaCertificatesList) {
        message.nocIcaCertificatesList.push(NocIcaCertificates.fromPartial(e))
      }
    }
    if (object.revokedNocRootCertificatesList !== undefined && object.revokedNocRootCertificatesList !== null) {
      for (const e of object.revokedNocRootCertificatesList) {
        message.revokedNocRootCertificatesList.push(RevokedNocRootCertificates.fromPartial(e))
      }
    }
    if (object.nocRootCertificatesByVidAndSkidList !== undefined && object.nocRootCertificatesByVidAndSkidList !== null) {
      for (const e of object.nocRootCertificatesByVidAndSkidList) {
        message.nocRootCertificatesByVidAndSkidList.push(NocRootCertificatesByVidAndSkid.fromPartial(e))
      }
    }
    return message
  }
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
