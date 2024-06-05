import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateVendorInfo } from "./types/zigbeealliance/distributedcomplianceledger/vendorinfo/tx";
import { MsgUpdateVendorInfo } from "./types/zigbeealliance/distributedcomplianceledger/vendorinfo/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgCreateVendorInfo", MsgCreateVendorInfo],
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgUpdateVendorInfo", MsgUpdateVendorInfo],
    
];

export { msgTypes }