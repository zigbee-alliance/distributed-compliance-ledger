import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgProvisionModel } from "./types/compliance/tx";
import { MsgDeleteComplianceInfo } from "./types/compliance/tx";
import { MsgUpdateComplianceInfo } from "./types/compliance/tx";
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
    msgProvisionModel: (data: MsgProvisionModel) => EncodeObject;
    msgDeleteComplianceInfo: (data: MsgDeleteComplianceInfo) => EncodeObject;
    msgUpdateComplianceInfo: (data: MsgUpdateComplianceInfo) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
