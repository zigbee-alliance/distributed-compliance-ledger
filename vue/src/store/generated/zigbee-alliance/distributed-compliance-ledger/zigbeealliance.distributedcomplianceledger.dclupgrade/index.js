import { txClient, queryClient, MissingWalletError, registry } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { ApprovedUpgrade } from "./module/types/dclupgrade/approved_upgrade";
import { Grant } from "./module/types/dclupgrade/grant";
import { ProposedUpgrade } from "./module/types/dclupgrade/proposed_upgrade";
export { ApprovedUpgrade, Grant, ProposedUpgrade };
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
        ProposedUpgrade: {},
        ProposedUpgradeAll: {},
        ApprovedUpgrade: {},
        ApprovedUpgradeAll: {},
        _Structure: {
            ApprovedUpgrade: getStructure(ApprovedUpgrade.fromPartial({})),
            Grant: getStructure(Grant.fromPartial({})),
            ProposedUpgrade: getStructure(ProposedUpgrade.fromPartial({})),
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
        getProposedUpgrade: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedUpgrade[JSON.stringify(params)] ?? {};
        },
        getProposedUpgradeAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedUpgradeAll[JSON.stringify(params)] ?? {};
        },
        getApprovedUpgrade: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ApprovedUpgrade[JSON.stringify(params)] ?? {};
        },
        getApprovedUpgradeAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ApprovedUpgradeAll[JSON.stringify(params)] ?? {};
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
            console.log('Vuex module: zigbeealliance.distributedcomplianceledger.dclupgrade initialized!');
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
        async QueryProposedUpgrade({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedUpgrade(key.name)).data;
                commit('QUERY', { query: 'ProposedUpgrade', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedUpgrade', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedUpgrade']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedUpgrade', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProposedUpgradeAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedUpgradeAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryProposedUpgradeAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ProposedUpgradeAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedUpgradeAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedUpgradeAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedUpgradeAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryApprovedUpgrade({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryApprovedUpgrade(key.name)).data;
                commit('QUERY', { query: 'ApprovedUpgrade', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryApprovedUpgrade', payload: { options: { all }, params: { ...key }, query } });
                return getters['getApprovedUpgrade']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryApprovedUpgrade', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryApprovedUpgradeAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryApprovedUpgradeAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryApprovedUpgradeAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ApprovedUpgradeAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryApprovedUpgradeAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getApprovedUpgradeAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryApprovedUpgradeAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgApproveUpgrade({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveUpgrade(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveUpgrade:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveUpgrade:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgProposeUpgrade({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeUpgrade(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeUpgrade:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeUpgrade:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgApproveUpgrade({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveUpgrade(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveUpgrade:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveUpgrade:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgProposeUpgrade({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeUpgrade(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeUpgrade:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeUpgrade:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
