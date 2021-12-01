import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('Description')
export class Description {
  @JsonProperty('name', String)
  name = '';

  @JsonProperty('identity', String, true)
  identity = null;

  @JsonProperty('website', String, true)
  website = null;

  @JsonProperty('details', String, true)
  details = null;

  constructor(vid?: string, pid?: string) {
  }
}
