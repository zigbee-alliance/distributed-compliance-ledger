import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgProposeRevokeAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgApproveAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgRejectAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgProposeAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgApproveRevokeAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", MsgProposeRevokeAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", MsgApproveAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgRejectAddAccount", MsgRejectAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", MsgProposeAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", MsgApproveRevokeAccount],
    
];

export { msgTypes }