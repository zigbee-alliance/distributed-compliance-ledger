import {JsonObject, JsonProperty} from 'json2typescript';

@JsonObject('BlockData')
export class BlockData {
  @JsonProperty('txs', [String])
  txs: string[] = undefined;
}
