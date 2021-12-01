import { JsonObject, JsonProperty } from 'json2typescript';
import { BlockHeader } from './block-header';
import { BlockData } from './block-data';

@JsonObject('Block')
export class Block {
  @JsonProperty('header', BlockHeader)
  header: BlockHeader = undefined;

  @JsonProperty('data', BlockData)
  data: BlockData = undefined;
}
