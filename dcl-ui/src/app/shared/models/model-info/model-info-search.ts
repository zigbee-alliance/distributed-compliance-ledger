import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('ModelInfoSearch')
export class ModelInfoSearch {
  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;
}
