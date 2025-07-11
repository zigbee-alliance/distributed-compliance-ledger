import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateModel } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgUpdateModel } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgUpdateModelVersion } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgDeleteModel } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgDeleteModelVersion } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgCreateModelVersion } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModel", MsgCreateModel],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModel", MsgUpdateModel],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModelVersion", MsgUpdateModelVersion],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModel", MsgDeleteModel],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModelVersion", MsgDeleteModelVersion],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModelVersion", MsgCreateModelVersion],
    
];

export { msgTypes }