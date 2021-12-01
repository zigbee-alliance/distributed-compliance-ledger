import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('ModelInfoHeader')
export class ModelInfoHeader {
  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('name', String)
  name: string = undefined;

  @JsonProperty('owner', String)
  owner: string = undefined;

  @JsonProperty('sku', String)
  sku: string = undefined;
}
