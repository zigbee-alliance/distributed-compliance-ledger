import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateValidator } from "./types/validator/tx";
import { MsgDisableValidator } from "./types/validator/tx";
import { MsgEnableValidator } from "./types/validator/tx";
import { MsgApproveDisableValidator } from "./types/validator/tx";
import { MsgProposeDisableValidator } from "./types/validator/tx";
import { MsgRejectDisableValidator } from "./types/validator/tx";
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
    msgCreateValidator: (data: MsgCreateValidator) => EncodeObject;
    msgDisableValidator: (data: MsgDisableValidator) => EncodeObject;
    msgEnableValidator: (data: MsgEnableValidator) => EncodeObject;
    msgApproveDisableValidator: (data: MsgApproveDisableValidator) => EncodeObject;
    msgProposeDisableValidator: (data: MsgProposeDisableValidator) => EncodeObject;
    msgRejectDisableValidator: (data: MsgRejectDisableValidator) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
