import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteVendorInfo } from "./types/vendorinfo/tx";
import { MsgUpdateVendorInfo } from "./types/vendorinfo/tx";
import { MsgCreateVendorInfo } from "./types/vendorinfo/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgDeleteVendorInfo: (data: MsgDeleteVendorInfo) => EncodeObject;
    msgUpdateVendorInfo: (data: MsgUpdateVendorInfo) => EncodeObject;
    msgCreateVendorInfo: (data: MsgCreateVendorInfo) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
