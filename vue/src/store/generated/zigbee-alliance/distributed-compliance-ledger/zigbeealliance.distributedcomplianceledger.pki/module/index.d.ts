import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgProposeAddX509RootCert } from "./types/pki/tx";
import { MsgAddX509Cert } from "./types/pki/tx";
import { MsgApproveRevokeX509RootCert } from "./types/pki/tx";
import { MsgApproveAddX509RootCert } from "./types/pki/tx";
import { MsgProposeRevokeX509RootCert } from "./types/pki/tx";
import { MsgRevokeX509Cert } from "./types/pki/tx";
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
    msgProposeAddX509RootCert: (data: MsgProposeAddX509RootCert) => EncodeObject;
    msgAddX509Cert: (data: MsgAddX509Cert) => EncodeObject;
    msgApproveRevokeX509RootCert: (data: MsgApproveRevokeX509RootCert) => EncodeObject;
    msgApproveAddX509RootCert: (data: MsgApproveAddX509RootCert) => EncodeObject;
    msgProposeRevokeX509RootCert: (data: MsgProposeRevokeX509RootCert) => EncodeObject;
    msgRevokeX509Cert: (data: MsgRevokeX509Cert) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
