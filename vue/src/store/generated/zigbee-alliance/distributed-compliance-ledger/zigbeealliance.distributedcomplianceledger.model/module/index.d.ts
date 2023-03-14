import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateModelVersion } from "./types/model/tx";
import { MsgUpdateModel } from "./types/model/tx";
import { MsgCreateModel } from "./types/model/tx";
import { MsgDeleteModel } from "./types/model/tx";
import { MsgCreateModelVersion } from "./types/model/tx";
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
    msgUpdateModelVersion: (data: MsgUpdateModelVersion) => EncodeObject;
    msgUpdateModel: (data: MsgUpdateModel) => EncodeObject;
    msgCreateModel: (data: MsgCreateModel) => EncodeObject;
    msgDeleteModel: (data: MsgDeleteModel) => EncodeObject;
    msgCreateModelVersion: (data: MsgCreateModelVersion) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
