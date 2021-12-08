import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateModelVersion } from "./types/model/tx";
import { MsgUpdateModelVersion } from "./types/model/tx";
import { MsgCreateModel } from "./types/model/tx";
import { MsgUpdateModelVersions } from "./types/model/tx";
import { MsgUpdateModel } from "./types/model/tx";
import { MsgDeleteModel } from "./types/model/tx";
import { MsgDeleteModelVersion } from "./types/model/tx";
import { MsgCreateModelVersions } from "./types/model/tx";
import { MsgDeleteModelVersions } from "./types/model/tx";
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
    msgCreateModelVersion: (data: MsgCreateModelVersion) => EncodeObject;
    msgUpdateModelVersion: (data: MsgUpdateModelVersion) => EncodeObject;
    msgCreateModel: (data: MsgCreateModel) => EncodeObject;
    msgUpdateModelVersions: (data: MsgUpdateModelVersions) => EncodeObject;
    msgUpdateModel: (data: MsgUpdateModel) => EncodeObject;
    msgDeleteModel: (data: MsgDeleteModel) => EncodeObject;
    msgDeleteModelVersion: (data: MsgDeleteModelVersion) => EncodeObject;
    msgCreateModelVersions: (data: MsgCreateModelVersions) => EncodeObject;
    msgDeleteModelVersions: (data: MsgDeleteModelVersions) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
