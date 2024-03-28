import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveDisableValidator } from "./types/zigbeealliance/distributedcomplianceledger/validator/tx";
import { MsgDisableValidator } from "./types/zigbeealliance/distributedcomplianceledger/validator/tx";
import { MsgProposeDisableValidator } from "./types/zigbeealliance/distributedcomplianceledger/validator/tx";
import { MsgRejectDisableValidator } from "./types/zigbeealliance/distributedcomplianceledger/validator/tx";
import { MsgCreateValidator } from "./types/zigbeealliance/distributedcomplianceledger/validator/tx";
import { MsgEnableValidator } from "./types/zigbeealliance/distributedcomplianceledger/validator/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgApproveDisableValidator", MsgApproveDisableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgDisableValidator", MsgDisableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgProposeDisableValidator", MsgProposeDisableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgRejectDisableValidator", MsgRejectDisableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgCreateValidator", MsgCreateValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgEnableValidator", MsgEnableValidator],
    
];

export { msgTypes }