import { txClient, queryClient, MissingWalletError, registry } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { Description } from "./module/types/validator/description";
import { DisabledValidator } from "./module/types/validator/disabled_validator";
import { Grant } from "./module/types/validator/grant";
import { LastValidatorPower } from "./module/types/validator/last_validator_power";
import { ProposedDisableValidator } from "./module/types/validator/proposed_disable_validator";
import { RejectedDisableValidator } from "./module/types/validator/rejected_validator";
import { Validator } from "./module/types/validator/validator";
export { Description, DisabledValidator, Grant, LastValidatorPower, ProposedDisableValidator, RejectedDisableValidator, Validator };
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
        Validator: {},
        ValidatorAll: {},
        LastValidatorPower: {},
        LastValidatorPowerAll: {},
        ProposedDisableValidator: {},
        ProposedDisableValidatorAll: {},
        DisabledValidator: {},
        DisabledValidatorAll: {},
        RejectedDisableValidator: {},
        RejectedDisableValidatorAll: {},
        _Structure: {
            Description: getStructure(Description.fromPartial({})),
            DisabledValidator: getStructure(DisabledValidator.fromPartial({})),
            Grant: getStructure(Grant.fromPartial({})),
            LastValidatorPower: getStructure(LastValidatorPower.fromPartial({})),
            ProposedDisableValidator: getStructure(ProposedDisableValidator.fromPartial({})),
            RejectedDisableValidator: getStructure(RejectedDisableValidator.fromPartial({})),
            Validator: getStructure(Validator.fromPartial({})),
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
        getValidator: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Validator[JSON.stringify(params)] ?? {};
        },
        getValidatorAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ValidatorAll[JSON.stringify(params)] ?? {};
        },
        getLastValidatorPower: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.LastValidatorPower[JSON.stringify(params)] ?? {};
        },
        getLastValidatorPowerAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.LastValidatorPowerAll[JSON.stringify(params)] ?? {};
        },
        getProposedDisableValidator: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedDisableValidator[JSON.stringify(params)] ?? {};
        },
        getProposedDisableValidatorAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.ProposedDisableValidatorAll[JSON.stringify(params)] ?? {};
        },
        getDisabledValidator: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.DisabledValidator[JSON.stringify(params)] ?? {};
        },
        getDisabledValidatorAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.DisabledValidatorAll[JSON.stringify(params)] ?? {};
        },
        getRejectedDisableValidator: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RejectedDisableValidator[JSON.stringify(params)] ?? {};
        },
        getRejectedDisableValidatorAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.RejectedDisableValidatorAll[JSON.stringify(params)] ?? {};
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
            console.log('Vuex module: zigbeealliance.distributedcomplianceledger.validator initialized!');
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
        async QueryValidator({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryValidator(key.owner)).data;
                commit('QUERY', { query: 'Validator', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryValidator', payload: { options: { all }, params: { ...key }, query } });
                return getters['getValidator']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryValidator', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryValidatorAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryValidatorAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ValidatorAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryValidatorAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getValidatorAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryValidatorAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryLastValidatorPower({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryLastValidatorPower(key.owner)).data;
                commit('QUERY', { query: 'LastValidatorPower', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryLastValidatorPower', payload: { options: { all }, params: { ...key }, query } });
                return getters['getLastValidatorPower']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryLastValidatorPower', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryLastValidatorPowerAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryLastValidatorPowerAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryLastValidatorPowerAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'LastValidatorPowerAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryLastValidatorPowerAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getLastValidatorPowerAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryLastValidatorPowerAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProposedDisableValidator({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedDisableValidator(key.address)).data;
                commit('QUERY', { query: 'ProposedDisableValidator', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedDisableValidator', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedDisableValidator']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedDisableValidator', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryProposedDisableValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryProposedDisableValidatorAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryProposedDisableValidatorAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'ProposedDisableValidatorAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryProposedDisableValidatorAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getProposedDisableValidatorAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryProposedDisableValidatorAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryDisabledValidator({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryDisabledValidator(key.address)).data;
                commit('QUERY', { query: 'DisabledValidator', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryDisabledValidator', payload: { options: { all }, params: { ...key }, query } });
                return getters['getDisabledValidator']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryDisabledValidator', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryDisabledValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryDisabledValidatorAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryDisabledValidatorAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'DisabledValidatorAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryDisabledValidatorAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getDisabledValidatorAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryDisabledValidatorAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRejectedDisableValidator({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRejectedDisableValidator(key.owner)).data;
                commit('QUERY', { query: 'RejectedDisableValidator', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRejectedDisableValidator', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRejectedDisableValidator']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRejectedDisableValidator', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryRejectedDisableValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryRejectedDisableValidatorAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryRejectedDisableValidatorAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'RejectedDisableValidatorAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryRejectedDisableValidatorAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getRejectedDisableValidatorAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryRejectedDisableValidatorAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgEnableValidator({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgEnableValidator(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgEnableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgEnableValidator:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgProposeDisableValidator({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeDisableValidator(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeDisableValidator:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgCreateValidator({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateValidator(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateValidator:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgRejectDisableValidator({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRejectDisableValidator(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRejectDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRejectDisableValidator:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgApproveDisableValidator({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveDisableValidator(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveDisableValidator:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgDisableValidator({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDisableValidator(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgDisableValidator:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgEnableValidator({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgEnableValidator(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgEnableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgEnableValidator:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgProposeDisableValidator({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeDisableValidator(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeDisableValidator:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgCreateValidator({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateValidator(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateValidator:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgRejectDisableValidator({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRejectDisableValidator(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRejectDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRejectDisableValidator:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgApproveDisableValidator({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveDisableValidator(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveDisableValidator:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgDisableValidator({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgDisableValidator(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgDisableValidator:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgDisableValidator:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
