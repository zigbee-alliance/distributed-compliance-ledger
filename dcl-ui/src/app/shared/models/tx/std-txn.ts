import { JsonObject, JsonProperty } from 'json2typescript';
import { StdTxnValue } from './std-txn-value';

@JsonObject('StdTxn')
export class StdTxn {

  @JsonProperty('type', String)
  type = 'cosmos-sdk/StdTx';

  @JsonProperty('value', StdTxnValue)
  value: StdTxnValue = new StdTxnValue();

  constructor(init?: Partial<StdTxn>) {
    Object.assign(this, init);
  }
}
