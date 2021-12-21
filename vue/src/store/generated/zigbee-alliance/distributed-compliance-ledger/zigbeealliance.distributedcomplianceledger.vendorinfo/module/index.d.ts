import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateVendorInfoType } from "./types/vendorinfo/tx";
import { MsgUpdateVendorInfoType } from "./types/vendorinfo/tx";
import { MsgDeleteVendorInfoType } from "./types/vendorinfo/tx";
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
    msgCreateVendorInfoType: (data: MsgCreateVendorInfoType) => EncodeObject;
    msgUpdateVendorInfoType: (data: MsgUpdateVendorInfoType) => EncodeObject;
    msgDeleteVendorInfoType: (data: MsgDeleteVendorInfoType) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
