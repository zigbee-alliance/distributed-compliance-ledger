import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('BlockHeader')
export class BlockHeader {
  @JsonProperty('chain_id', String)
  chainId: string = undefined;

  @JsonProperty('height', AminoNumberConverter)
  height: number = undefined;

  @JsonProperty('time', DateConverter)
  time: Date = undefined;

  @JsonProperty('num_txs', AminoNumberConverter)
  numTxs: number = undefined;

  @JsonProperty('total_txs', AminoNumberConverter)
  totalTxs: number = undefined;
}
