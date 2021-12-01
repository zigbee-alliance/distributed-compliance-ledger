import { JsonObject, JsonProperty } from 'json2typescript';
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('TestingResult')
export class TestingResult {
  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('test_result', String)
  testResult = '';

  @JsonProperty('test_date', DateConverter)
  testDate: Date = new Date();

  constructor(vid: string, pid: string) {
    this.vid = parseInt(vid, 10);
    this.pid = parseInt(pid, 10);
  }
}
