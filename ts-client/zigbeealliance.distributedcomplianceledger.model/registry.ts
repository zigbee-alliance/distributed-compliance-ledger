import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateModelVersion } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgCreateModelVersion } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgCreateModel } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgDeleteModel } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgDeleteModelVersion } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";
import { MsgUpdateModel } from "./types/zigbeealliance/distributedcomplianceledger/model/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModelVersion", MsgUpdateModelVersion],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModelVersion", MsgCreateModelVersion],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModel", MsgCreateModel],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModel", MsgDeleteModel],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModelVersion", MsgDeleteModelVersion],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModel", MsgUpdateModel],
    
];

export { msgTypes }