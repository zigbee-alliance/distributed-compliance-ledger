// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgApproveAddX509RootCert } from "./types/pki/tx";
import { MsgApproveRevokeX509RootCert } from "./types/pki/tx";
import { MsgProposeRevokeX509RootCert } from "./types/pki/tx";
import { MsgRevokeX509Cert } from "./types/pki/tx";
import { MsgRejectAddX509RootCert } from "./types/pki/tx";
import { MsgProposeAddX509RootCert } from "./types/pki/tx";
import { MsgAddX509Cert } from "./types/pki/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgApproveAddX509RootCert", MsgApproveAddX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgApproveRevokeX509RootCert", MsgApproveRevokeX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgProposeRevokeX509RootCert", MsgProposeRevokeX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeX509Cert", MsgRevokeX509Cert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgRejectAddX509RootCert", MsgRejectAddX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgProposeAddX509RootCert", MsgProposeAddX509RootCert],
    ["/zigbeealliance.distributedcomplianceledger.pki.MsgAddX509Cert", MsgAddX509Cert],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgApproveAddX509RootCert: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgApproveAddX509RootCert", value: MsgApproveAddX509RootCert.fromPartial(data) }),
        msgApproveRevokeX509RootCert: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgApproveRevokeX509RootCert", value: MsgApproveRevokeX509RootCert.fromPartial(data) }),
        msgProposeRevokeX509RootCert: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgProposeRevokeX509RootCert", value: MsgProposeRevokeX509RootCert.fromPartial(data) }),
        msgRevokeX509Cert: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeX509Cert", value: MsgRevokeX509Cert.fromPartial(data) }),
        msgRejectAddX509RootCert: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRejectAddX509RootCert", value: MsgRejectAddX509RootCert.fromPartial(data) }),
        msgProposeAddX509RootCert: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgProposeAddX509RootCert", value: MsgProposeAddX509RootCert.fromPartial(data) }),
        msgAddX509Cert: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgAddX509Cert", value: MsgAddX509Cert.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
