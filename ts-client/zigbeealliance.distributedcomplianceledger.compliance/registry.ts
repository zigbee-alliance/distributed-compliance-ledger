import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRevokeModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgCertifyModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgUpdateComplianceInfo } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgProvisionModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgDeleteComplianceInfo } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel", MsgRevokeModel],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel", MsgCertifyModel],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgUpdateComplianceInfo", MsgUpdateComplianceInfo],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgProvisionModel", MsgProvisionModel],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgDeleteComplianceInfo", MsgDeleteComplianceInfo],
    
];

export { msgTypes }