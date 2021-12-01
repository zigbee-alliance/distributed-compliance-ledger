import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('StdFee')
export class StdFee {
  @JsonProperty('gas', AminoNumberConverter)
  gas = 0;

  constructor(init?: Partial<StdFee>) {
    Object.assign(this, init);
  }
}
