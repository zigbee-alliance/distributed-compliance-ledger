import { JsonObject, JsonProperty } from 'json2typescript';
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('ComplianceInfo')
export class ComplianceInfo {
  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('owner', String)
  owner = '';

  @JsonProperty('state', String, true)
  state = null;

  @JsonProperty('date', DateConverter)
  date: Date = new Date();

  @JsonProperty('certification_type', String, true)
  certificationType = 'zb';

  @JsonProperty('reason', String, true)
  reason = null;

  constructor(vid?: string, pid?: string) {
    if (vid && vid.length > 0) {
      this.vid = parseInt(vid, 10);
    }
    if (pid && pid.length > 0) {
      this.pid = parseInt(pid, 10);
    }
  }

  isCertified(): boolean {
    return this.state === 'certified';
  }
}
