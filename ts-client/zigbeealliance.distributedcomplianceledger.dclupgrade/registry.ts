import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRejectUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgProposeUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgApproveUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgRejectUpgrade", MsgRejectUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgProposeUpgrade", MsgProposeUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgApproveUpgrade", MsgApproveUpgrade],
    
];

export { msgTypes }