import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('ModelInfo')
export class ModelInfo {
  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('cid', Number, true)
  cid = null;

  @JsonProperty('name', String)
  name = '';

  @JsonProperty('owner', String)
  owner = '';

  @JsonProperty('description', String)
  description = '';

  @JsonProperty('sku', String)
  sku = '';

  @JsonProperty('firmware_version', String)
  firmwareVersion = '';

  @JsonProperty('hardware_version', String)
  hardwareVersion = null;

  @JsonProperty('custom', String, true)
  custom = null;

  @JsonProperty('tis_or_trp_testing_completed', Boolean)
  tisOrTrpTestingCompleted = false;
}
