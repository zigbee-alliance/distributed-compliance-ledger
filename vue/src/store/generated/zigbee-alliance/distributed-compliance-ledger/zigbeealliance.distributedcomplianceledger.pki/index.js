import { txClient, queryClient, MissingWalletError, registry } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { ApprovedCertificates } from "./module/types/pki/approved_certificates";
import { ApprovedCertificatesBySubject } from "./module/types/pki/approved_certificates_by_subject";
import { ApprovedRootCertificates } from "./module/types/pki/approved_root_certificates";
import { Certificate } from "./module/types/pki/certificate";
import { CertificateIdentifier } from "./module/types/pki/certificate_identifier";
import { ChildCertificates } from "./module/types/pki/child_certificates";
import { Grant } from "./module/types/pki/grant";
import { PkiRevocationDistributionPoint } from "./module/types/pki/pki_revocation_distribution_point";
import { ProposedCertificate } from "./module/types/pki/proposed_certificate";
import { ProposedCertificateRevocation } from "./module/types/pki/proposed_certificate_revocation";
import { RejectedCertificate } from "./module/types/pki/rejected_certificate";
import { RevokedCertificates } from "./module/types/pki/revoked_certificates";
import { RevokedRootCertificates } from "./module/types/pki/revoked_root_certificates";
import { UniqueCertificate } from "./module/types/pki/unique_certificate";
export { ApprovedCertificates, ApprovedCertificatesBySubject, ApprovedRootCertificates, Certificate, CertificateIdentifier, ChildCertificates, Grant, PkiRevocationDistributionPoint, ProposedCertificate, ProposedCertificateRevocation, RejectedCertificate, RevokedCertificates, RevokedRootCertificates, UniqueCertificate };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        ApprovedCertificates: {},
        ApprovedCertificatesAll: {},
        ProposedCertificate: {},
        ProposedCertificateAll: {},
        ChildCertificates: {},
        ProposedCertificateRevocation: {},
        ProposedCertificateRevocationAll: {},
        RevokedCertificates: {},
        RevokedCertificatesAll: {},
        ApprovedRootCertificates: {},
        RevokedRootCertificates: {},
        ApprovedCertificatesBySubject: {},
        RejectedCertificate: {},
        RejectedCertificateAll: {},
        PkiRevocationDistributionPoint: {},
        PkiRevocationDistributionPointAll: {},
        _Structure: {
            ApprovedCertificates: getStructure(ApprovedCertificates.fromPartial({})),
            ApprovedCertificatesBySubject: getStructure(ApprovedCertificatesBySubject.fromPartial({})),
            ApprovedRootCertificates: getStructure(ApprovedRootCertificates.fromPartial({})),
            Certificate: getStructure(Certificate.fromPartial({})),
            CertificateIdentifier: getStructure(CertificateIdentifier.fromPartial({})),
            ChildCertificates: getStructure(ChildCertificates.fromPartial({})),
            Grant: getStructure(Grant.fromPartial({})),
            PkiRevocationDistributionPoint: getStructure(PkiRevocationDistributionPoint.fromPartial({})),
            ProposedCertificate: getStructure(ProposedCertificate.fromPartial({})),
            ProposedCertificateRevocation: getStructure(ProposedCertificateRevocation.fromPartial({})),
            RejectedCertificate: getStructure(RejectedCertificate.fromPartial({})),
            RevokedCertificates: getStructure(RevokedCertificates.fromPartial({})),
            RevokedRootCertificates: getStructure(RevokedRootCertificates.fromPartial({})),
            UniqueCertificate: getStructure(UniqueCertificate.fromPartial({})),
        },
        _Registry: registry,
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(JSON.stringify(subscription));
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(JSON.stringify(subscription));
        }
    },
    getters: {
        getApprovedCertificates: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ApprovedCertificates[JSON.stringify(params)] ?? {};
        },
        getApprovedCertificatesAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ApprovedCertificatesAll[JSON.stringify(params)] ?? {};
        },
        getProposedCertificate: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedCertificate[JSON.stringify(params)] ?? {};
        },
        getProposedCertificateAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedCertificateAll[JSON.stringify(params)] ?? {};
        },
        getChildCertificates: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ChildCertificates[JSON.stringify(params)] ?? {};
        },
        getProposedCertificateRevocation: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedCertificateRevocation[JSON.stringify(params)] ?? {};
        },
        getProposedCertificateRevocationAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedCertificateRevocationAll[JSON.stringify(params)] ?? {};
        },
        getRevokedCertificates: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RevokedCertificates[JSON.stringify(params)] ?? {};
        },
        getRevokedCertificatesAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RevokedCertificatesAll[JSON.stringify(params)] ?? {};
        },
        getApprovedRootCertificates: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ApprovedRootCertificates[JSON.stringify(params)] ?? {};
        },
        getRevokedRootCertificates: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RevokedRootCertificates[JSON.stringify(params)] ?? {};
        },
        getApprovedCertificatesBySubject: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ApprovedCertificatesBySubject[JSON.stringify(params)] ?? {};
        },
        getRejectedCertificate: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RejectedCertificate[JSON.stringify(params)] ?? {};
        },
        getRejectedCertificateAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RejectedCertificateAll[JSON.stringify(params)] ?? {};
        },
        getPkiRevocationDistributionPoint: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PkiRevocationDistributionPoint[JSON.stringify(params)] ?? {};
        },
        getPkiRevocationDistributionPointAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PkiRevocationDistributionPointAll[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        },
        getRegistry: (state) => {
            return state._Registry;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: zigbeealliance.distributedcomplianceledger.pki initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    const sub = JSON.parse(subscription);
                    await dispatch(sub.action, sub.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryApprovedCertificates({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryApprovedCertificates(key.subject, key.subjectKeyId)).data;
                commit('QUERY', { query: 'ApprovedCertificates', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryApprovedCertificates', payload: { options: { all }, params: { ...key }, query } });
                return getters['getApprovedCertificates']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryApprovedCertificates', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryApprovedCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryApprovedCertificatesAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryApprovedCertificatesAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ApprovedCertificatesAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryApprovedCertificatesAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getApprovedCertificatesAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryApprovedCertificatesAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProposedCertificate({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedCertificate(key.subject, key.subjectKeyId)).data;
                commit('QUERY', { query: 'ProposedCertificate', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedCertificate', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedCertificate']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedCertificate', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProposedCertificateAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedCertificateAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryProposedCertificateAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ProposedCertificateAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedCertificateAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedCertificateAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedCertificateAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryChildCertificates({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryChildCertificates(key.issuer, key.authorityKeyId)).data;
                commit('QUERY', { query: 'ChildCertificates', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryChildCertificates', payload: { options: { all }, params: { ...key }, query } });
                return getters['getChildCertificates']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryChildCertificates', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProposedCertificateRevocation({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedCertificateRevocation(key.subject, key.subjectKeyId)).data;
                commit('QUERY', { query: 'ProposedCertificateRevocation', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedCertificateRevocation', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedCertificateRevocation']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedCertificateRevocation', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProposedCertificateRevocationAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedCertificateRevocationAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryProposedCertificateRevocationAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ProposedCertificateRevocationAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedCertificateRevocationAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedCertificateRevocationAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedCertificateRevocationAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRevokedCertificates({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRevokedCertificates(key.subject, key.subjectKeyId)).data;
                commit('QUERY', { query: 'RevokedCertificates', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRevokedCertificates', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRevokedCertificates']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRevokedCertificates', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRevokedCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRevokedCertificatesAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryRevokedCertificatesAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'RevokedCertificatesAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRevokedCertificatesAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRevokedCertificatesAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRevokedCertificatesAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryApprovedRootCertificates({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryApprovedRootCertificates()).data;
                commit('QUERY', { query: 'ApprovedRootCertificates', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryApprovedRootCertificates', payload: { options: { all }, params: { ...key }, query } });
                return getters['getApprovedRootCertificates']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryApprovedRootCertificates', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRevokedRootCertificates({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRevokedRootCertificates()).data;
                commit('QUERY', { query: 'RevokedRootCertificates', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRevokedRootCertificates', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRevokedRootCertificates']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRevokedRootCertificates', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryApprovedCertificatesBySubject({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryApprovedCertificatesBySubject(key.subject)).data;
                commit('QUERY', { query: 'ApprovedCertificatesBySubject', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryApprovedCertificatesBySubject', payload: { options: { all }, params: { ...key }, query } });
                return getters['getApprovedCertificatesBySubject']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryApprovedCertificatesBySubject', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRejectedCertificate({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRejectedCertificate(key.subject, key.subjectKeyId)).data;
                commit('QUERY', { query: 'RejectedCertificate', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRejectedCertificate', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRejectedCertificate']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRejectedCertificate', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRejectedCertificateAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRejectedCertificateAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryRejectedCertificateAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'RejectedCertificateAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRejectedCertificateAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRejectedCertificateAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRejectedCertificateAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgApproveAddX509RootCert({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveAddX509RootCert(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveAddX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveAddX509RootCert:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgApproveRevokeX509RootCert({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveRevokeX509RootCert(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveRevokeX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveRevokeX509RootCert:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgProposeRevokeX509RootCert({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeRevokeX509RootCert(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeRevokeX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeRevokeX509RootCert:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgRevokeX509Cert({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRevokeX509Cert(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRevokeX509Cert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRevokeX509Cert:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgRejectAddX509RootCert({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRejectAddX509RootCert(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRejectAddX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRejectAddX509RootCert:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgProposeAddX509RootCert({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeAddX509RootCert(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeAddX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeAddX509RootCert:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgAddX509Cert({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgAddX509Cert(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgAddX509Cert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgAddX509Cert:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgApproveAddX509RootCert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveAddX509RootCert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveAddX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveAddX509RootCert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgApproveRevokeX509RootCert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveRevokeX509RootCert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveRevokeX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveRevokeX509RootCert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgProposeRevokeX509RootCert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeRevokeX509RootCert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeRevokeX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeRevokeX509RootCert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgRevokeX509Cert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRevokeX509Cert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRevokeX509Cert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRevokeX509Cert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgRejectAddX509RootCert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRejectAddX509RootCert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRejectAddX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRejectAddX509RootCert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgProposeAddX509RootCert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeAddX509RootCert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeAddX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeAddX509RootCert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgAddX509Cert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgAddX509Cert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgAddX509Cert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgAddX509Cert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgAddX509Cert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgAddX509Cert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgAddX509Cert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgAddX509Cert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgApproveRevokeX509RootCert({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveRevokeX509RootCert(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveRevokeX509RootCert:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveRevokeX509RootCert:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgUpdatePkiRevocationDistributionPoint({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgUpdatePkiRevocationDistributionPoint(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgUpdatePkiRevocationDistributionPoint:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgUpdatePkiRevocationDistributionPoint:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
