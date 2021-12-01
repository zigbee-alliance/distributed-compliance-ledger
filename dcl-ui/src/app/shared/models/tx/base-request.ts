import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('BaseReq')
export class BaseReq {

  @JsonProperty('chain_id', String)
  chainId: string = undefined;

  @JsonProperty('account_number', AminoNumberConverter)
  accountNumber: number = undefined;

  @JsonProperty('sequence', AminoNumberConverter)
  sequence: number = undefined;

  @JsonProperty('from', String)
  from = '';

  constructor(init?: Partial<BaseReq>) {
    Object.assign(this, init);
  }

  setFrom(from: string) {
    this.from = from;
  }
}
