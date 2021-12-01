import { JsonObject, JsonProperty } from 'json2typescript';
import { BlockId } from './block-id';
import { BlockHeader } from './block-header';

@JsonObject('BlockMeta')
export class BlockMeta {
  @JsonProperty('block_id', BlockId)
  blockId: BlockId = undefined;

  @JsonProperty('header', BlockHeader)
  header: BlockHeader = undefined;
}
