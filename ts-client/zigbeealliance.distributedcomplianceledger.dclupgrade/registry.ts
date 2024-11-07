import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgProposeUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgRejectUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgApproveUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgProposeUpgrade", MsgProposeUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgRejectUpgrade", MsgRejectUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgApproveUpgrade", MsgApproveUpgrade],
    
];

export { msgTypes }