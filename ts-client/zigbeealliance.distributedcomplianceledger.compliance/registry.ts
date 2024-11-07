import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateComplianceInfo } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgDeleteComplianceInfo } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgCertifyModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgRevokeModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgProvisionModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgUpdateComplianceInfo", MsgUpdateComplianceInfo],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgDeleteComplianceInfo", MsgDeleteComplianceInfo],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel", MsgCertifyModel],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel", MsgRevokeModel],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgProvisionModel", MsgProvisionModel],
    
];

export { msgTypes }