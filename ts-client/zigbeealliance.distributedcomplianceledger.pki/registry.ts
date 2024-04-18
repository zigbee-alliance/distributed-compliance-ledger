import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgProposeAddX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgApproveAddX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRemoveX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRevokeX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddPkiRevocationDistributionPoint } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAssignVid } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRevokeNocRootX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRejectAddX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRevokeNocX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddNocX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddNocX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgUpdatePkiRevocationDistributionPoint } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgDeletePkiRevocationDistributionPoint } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgApproveRevokeX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgProposeRevokeX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgProposeAddX509RootCert", MsgProposeAddX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgApproveAddX509RootCert", MsgApproveAddX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgRemoveX509Cert", MsgRemoveX509Cert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeX509Cert", MsgRevokeX509Cert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgAddPkiRevocationDistributionPoint", MsgAddPkiRevocationDistributionPoint],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgAssignVid", MsgAssignVid],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeNocRootX509Cert", MsgRevokeNocRootX509Cert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgRejectAddX509RootCert", MsgRejectAddX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeNocX509Cert", MsgRevokeNocX509Cert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgAddNocX509RootCert", MsgAddNocX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgAddNocX509Cert", MsgAddNocX509Cert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgUpdatePkiRevocationDistributionPoint", MsgUpdatePkiRevocationDistributionPoint],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgDeletePkiRevocationDistributionPoint", MsgDeletePkiRevocationDistributionPoint],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgApproveRevokeX509RootCert", MsgApproveRevokeX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgAddX509Cert", MsgAddX509Cert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgProposeRevokeX509RootCert", MsgProposeRevokeX509RootCert],
    
];

export { msgTypes }