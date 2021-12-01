import { JsonObject, JsonProperty } from 'json2typescript';
import { ModelInfoHeader } from './model-info-header';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('ResultModelInfoHeaders')
export class ResultModelInfoHeaders {
  @JsonProperty('total', AminoNumberConverter)
  total: number = undefined;

  @JsonProperty('items', [ModelInfoHeader])
  items: ModelInfoHeader[] = undefined;
}
