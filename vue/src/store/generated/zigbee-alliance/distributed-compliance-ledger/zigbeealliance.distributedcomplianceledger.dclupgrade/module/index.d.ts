import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgProposeUpgrade } from "./types/dclupgrade/tx";
import { MsgApproveUpgrade } from "./types/dclupgrade/tx";
import { MsgRejectUpgrade } from "./types/dclupgrade/tx";
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
    msgProposeUpgrade: (data: MsgProposeUpgrade) => EncodeObject;
    msgApproveUpgrade: (data: MsgApproveUpgrade) => EncodeObject;
    msgRejectUpgrade: (data: MsgRejectUpgrade) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
