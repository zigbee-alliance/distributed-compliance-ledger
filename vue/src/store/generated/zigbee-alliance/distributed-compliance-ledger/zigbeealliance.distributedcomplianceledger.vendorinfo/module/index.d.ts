import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateVendorInfo } from "./types/vendorinfo/tx";
import { MsgUpdateVendorInfo } from "./types/vendorinfo/tx";
import { MsgDeleteVendorInfo } from "./types/vendorinfo/tx";
export declare const MissingWalletError: Error;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgCreateVendorInfo: (data: MsgCreateVendorInfo) => EncodeObject;
    msgUpdateVendorInfo: (data: MsgUpdateVendorInfo) => EncodeObject;
    msgDeleteVendorInfo: (data: MsgDeleteVendorInfo) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
