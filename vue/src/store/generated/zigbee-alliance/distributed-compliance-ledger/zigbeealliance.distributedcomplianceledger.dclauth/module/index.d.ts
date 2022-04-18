import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgProposeRevokeAccount } from "./types/dclauth/tx";
import { MsgProposeAddAccount } from "./types/dclauth/tx";
import { MsgApproveRevokeAccount } from "./types/dclauth/tx";
import { MsgApproveAddAccount } from "./types/dclauth/tx";
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
    msgProposeRevokeAccount: (data: MsgProposeRevokeAccount) => EncodeObject;
    msgProposeAddAccount: (data: MsgProposeAddAccount) => EncodeObject;
    msgApproveRevokeAccount: (data: MsgApproveRevokeAccount) => EncodeObject;
    msgApproveAddAccount: (data: MsgApproveAddAccount) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
