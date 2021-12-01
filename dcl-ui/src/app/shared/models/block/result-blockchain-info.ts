import { JsonObject, JsonProperty } from 'json2typescript';
import { BlockMeta } from './block-meta';

@JsonObject('ResultBlockchainInfo')
export class ResultBlockchainInfo {
  @JsonProperty('last_height', String)
  lastHeight: number = undefined;

  @JsonProperty('block_metas', [BlockMeta])
  blockMetas: BlockMeta[] = undefined;
}
