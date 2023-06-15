import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgEnableValidator } from "./types/validator/tx";
import { MsgProposeDisableValidator } from "./types/validator/tx";
import { MsgCreateValidator } from "./types/validator/tx";
import { MsgRejectDisableValidator } from "./types/validator/tx";
import { MsgApproveDisableValidator } from "./types/validator/tx";
import { MsgDisableValidator } from "./types/validator/tx";
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
    msgEnableValidator: (data: MsgEnableValidator) => EncodeObject;
    msgProposeDisableValidator: (data: MsgProposeDisableValidator) => EncodeObject;
    msgCreateValidator: (data: MsgCreateValidator) => EncodeObject;
    msgRejectDisableValidator: (data: MsgRejectDisableValidator) => EncodeObject;
    msgApproveDisableValidator: (data: MsgApproveDisableValidator) => EncodeObject;
    msgDisableValidator: (data: MsgDisableValidator) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
