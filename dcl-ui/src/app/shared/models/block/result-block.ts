import { JsonObject, JsonProperty } from 'json2typescript';
import { BlockMeta } from './block-meta';
import { Block } from './block';

@JsonObject('ResultBlock')
export class ResultBlock {
  @JsonProperty('block_meta', BlockMeta)
  blockMeta: BlockMeta = undefined;

  @JsonProperty('block', Block)
  block: Block = undefined;
}
