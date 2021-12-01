import { JsonObject, JsonProperty } from 'json2typescript';
import { BaseReq } from './base-request';
import { StdTxn } from './std-txn';

@JsonObject('StdSignMsg')
export class StdSignMsg {

  @JsonProperty('base_req', BaseReq)
  baseReq: BaseReq = undefined;

  @JsonProperty('txn', StdTxn)
  txn: StdTxn = undefined;

  constructor(init?: Partial<StdSignMsg>) {
    Object.assign(this, init);
  }
}
