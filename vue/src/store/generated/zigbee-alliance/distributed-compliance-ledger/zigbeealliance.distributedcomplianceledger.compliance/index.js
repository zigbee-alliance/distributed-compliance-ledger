import { txClient, queryClient, MissingWalletError, registry } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { CertifiedModel } from "./module/types/compliance/certified_model";
import { ComplianceHistoryItem } from "./module/types/compliance/compliance_history_item";
import { ComplianceInfo } from "./module/types/compliance/compliance_info";
import { ProvisionalModel } from "./module/types/compliance/provisional_model";
import { RevokedModel } from "./module/types/compliance/revoked_model";
export { CertifiedModel, ComplianceHistoryItem, ComplianceInfo, ProvisionalModel, RevokedModel };
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
        ComplianceInfo: {},
        ComplianceInfoAll: {},
        CertifiedModel: {},
        CertifiedModelAll: {},
        RevokedModel: {},
        RevokedModelAll: {},
        ProvisionalModel: {},
        ProvisionalModelAll: {},
        _Structure: {
            CertifiedModel: getStructure(CertifiedModel.fromPartial({})),
            ComplianceHistoryItem: getStructure(ComplianceHistoryItem.fromPartial({})),
            ComplianceInfo: getStructure(ComplianceInfo.fromPartial({})),
            ProvisionalModel: getStructure(ProvisionalModel.fromPartial({})),
            RevokedModel: getStructure(RevokedModel.fromPartial({})),
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
        getComplianceInfo: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ComplianceInfo[JSON.stringify(params)] ?? {};
        },
        getComplianceInfoAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ComplianceInfoAll[JSON.stringify(params)] ?? {};
        },
        getCertifiedModel: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.CertifiedModel[JSON.stringify(params)] ?? {};
        },
        getCertifiedModelAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.CertifiedModelAll[JSON.stringify(params)] ?? {};
        },
        getRevokedModel: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RevokedModel[JSON.stringify(params)] ?? {};
        },
        getRevokedModelAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RevokedModelAll[JSON.stringify(params)] ?? {};
        },
        getProvisionalModel: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProvisionalModel[JSON.stringify(params)] ?? {};
        },
        getProvisionalModelAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProvisionalModelAll[JSON.stringify(params)] ?? {};
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
            console.log('Vuex module: zigbeealliance.distributedcomplianceledger.compliance initialized!');
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
        async QueryComplianceInfo({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryComplianceInfo(key.vid, key.pid, key.software_version, key.certification_type)).data;
                commit('QUERY', { query: 'ComplianceInfo', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryComplianceInfo', payload: { options: { all }, params: { ...key }, query } });
                return getters['getComplianceInfo']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryComplianceInfo', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryComplianceInfoAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryComplianceInfoAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryComplianceInfoAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ComplianceInfoAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryComplianceInfoAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getComplianceInfoAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryComplianceInfoAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryCertifiedModel({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryCertifiedModel(key.vid, key.pid, key.software_version, key.certification_type)).data;
                commit('QUERY', { query: 'CertifiedModel', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryCertifiedModel', payload: { options: { all }, params: { ...key }, query } });
                return getters['getCertifiedModel']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryCertifiedModel', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryCertifiedModelAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryCertifiedModelAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryCertifiedModelAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'CertifiedModelAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryCertifiedModelAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getCertifiedModelAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryCertifiedModelAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRevokedModel({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRevokedModel(key.vid, key.pid, key.software_version, key.certification_type)).data;
                commit('QUERY', { query: 'RevokedModel', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRevokedModel', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRevokedModel']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRevokedModel', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRevokedModelAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRevokedModelAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryRevokedModelAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'RevokedModelAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRevokedModelAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRevokedModelAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRevokedModelAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProvisionalModel({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProvisionalModel(key.vid, key.pid, key.software_version, key.certification_type)).data;
                commit('QUERY', { query: 'ProvisionalModel', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProvisionalModel', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProvisionalModel']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProvisionalModel', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProvisionalModelAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProvisionalModelAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryProvisionalModelAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ProvisionalModelAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProvisionalModelAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProvisionalModelAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProvisionalModelAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgRevokeModel({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRevokeModel(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRevokeModel:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRevokeModel:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgCertifyModel({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCertifyModel(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCertifyModel:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCertifyModel:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgProvisionModel({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProvisionModel(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProvisionModel:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProvisionModel:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgRevokeModel({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRevokeModel(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRevokeModel:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRevokeModel:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgCertifyModel({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCertifyModel(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCertifyModel:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCertifyModel:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgProvisionModel({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProvisionModel(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProvisionModel:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProvisionModel:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
