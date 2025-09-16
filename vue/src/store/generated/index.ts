// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import ZigbeeallianceDistributedcomplianceledgerCompliance from './zigbeealliance.distributedcomplianceledger.compliance'
import ZigbeeallianceDistributedcomplianceledgerDclauth from './zigbeealliance.distributedcomplianceledger.dclauth'
import ZigbeeallianceDistributedcomplianceledgerDclgenutil from './zigbeealliance.distributedcomplianceledger.dclgenutil'
import ZigbeeallianceDistributedcomplianceledgerDclupgrade from './zigbeealliance.distributedcomplianceledger.dclupgrade'
import ZigbeeallianceDistributedcomplianceledgerModel from './zigbeealliance.distributedcomplianceledger.model'
import ZigbeeallianceDistributedcomplianceledgerPki from './zigbeealliance.distributedcomplianceledger.pki'
import ZigbeeallianceDistributedcomplianceledgerValidator from './zigbeealliance.distributedcomplianceledger.validator'
import ZigbeeallianceDistributedcomplianceledgerVendorinfo from './zigbeealliance.distributedcomplianceledger.vendorinfo'
import CosmosAuthV1Beta1 from './cosmos.auth.v1beta1'
import CosmosBaseTendermintV1Beta1 from './cosmos.base.tendermint.v1beta1'
import CosmosConsensusV1 from './cosmos.consensus.v1'
import CosmosParamsV1Beta1 from './cosmos.params.v1beta1'
import CosmosTxV1Beta1 from './cosmos.tx.v1beta1'
import CosmosUpgradeV1Beta1 from './cosmos.upgrade.v1beta1'
import CosmosVestingV1Beta1 from './cosmos.vesting.v1beta1'


export default { 
  ZigbeeallianceDistributedcomplianceledgerCompliance: load(ZigbeeallianceDistributedcomplianceledgerCompliance, 'zigbeealliance.distributedcomplianceledger.compliance'),
  ZigbeeallianceDistributedcomplianceledgerDclauth: load(ZigbeeallianceDistributedcomplianceledgerDclauth, 'zigbeealliance.distributedcomplianceledger.dclauth'),
  ZigbeeallianceDistributedcomplianceledgerDclgenutil: load(ZigbeeallianceDistributedcomplianceledgerDclgenutil, 'zigbeealliance.distributedcomplianceledger.dclgenutil'),
  ZigbeeallianceDistributedcomplianceledgerDclupgrade: load(ZigbeeallianceDistributedcomplianceledgerDclupgrade, 'zigbeealliance.distributedcomplianceledger.dclupgrade'),
  ZigbeeallianceDistributedcomplianceledgerModel: load(ZigbeeallianceDistributedcomplianceledgerModel, 'zigbeealliance.distributedcomplianceledger.model'),
  ZigbeeallianceDistributedcomplianceledgerPki: load(ZigbeeallianceDistributedcomplianceledgerPki, 'zigbeealliance.distributedcomplianceledger.pki'),
  ZigbeeallianceDistributedcomplianceledgerValidator: load(ZigbeeallianceDistributedcomplianceledgerValidator, 'zigbeealliance.distributedcomplianceledger.validator'),
  ZigbeeallianceDistributedcomplianceledgerVendorinfo: load(ZigbeeallianceDistributedcomplianceledgerVendorinfo, 'zigbeealliance.distributedcomplianceledger.vendorinfo'),
  CosmosAuthV1Beta1: load(CosmosAuthV1Beta1, 'cosmos.auth.v1beta1'),
  CosmosBaseTendermintV1Beta1: load(CosmosBaseTendermintV1Beta1, 'cosmos.base.tendermint.v1beta1'),
  CosmosConsensusV1: load(CosmosConsensusV1, 'cosmos.consensus.v1'),
  CosmosParamsV1Beta1: load(CosmosParamsV1Beta1, 'cosmos.params.v1beta1'),
  CosmosTxV1Beta1: load(CosmosTxV1Beta1, 'cosmos.tx.v1beta1'),
  CosmosUpgradeV1Beta1: load(CosmosUpgradeV1Beta1, 'cosmos.upgrade.v1beta1'),
  CosmosVestingV1Beta1: load(CosmosVestingV1Beta1, 'cosmos.vesting.v1beta1'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}