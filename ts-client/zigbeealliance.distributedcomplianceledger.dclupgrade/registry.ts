import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgProposeUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgRejectUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgApproveUpgrade", MsgApproveUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgProposeUpgrade", MsgProposeUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgRejectUpgrade", MsgRejectUpgrade],
    
];

export { msgTypes }