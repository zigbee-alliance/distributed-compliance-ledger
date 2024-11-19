import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateVendorInfo } from "./types/zigbeealliance/distributedcomplianceledger/vendorinfo/tx";
import { MsgCreateVendorInfo } from "./types/zigbeealliance/distributedcomplianceledger/vendorinfo/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgUpdateVendorInfo", MsgUpdateVendorInfo],
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgCreateVendorInfo", MsgCreateVendorInfo],
    
];

export { msgTypes }