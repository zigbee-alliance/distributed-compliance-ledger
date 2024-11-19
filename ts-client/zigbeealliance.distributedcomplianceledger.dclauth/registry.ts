import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveRevokeAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgProposeAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgProposeRevokeAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgRejectAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgApproveAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", MsgApproveRevokeAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", MsgProposeAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", MsgProposeRevokeAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgRejectAddAccount", MsgRejectAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", MsgApproveAddAccount],
    
];

export { msgTypes }