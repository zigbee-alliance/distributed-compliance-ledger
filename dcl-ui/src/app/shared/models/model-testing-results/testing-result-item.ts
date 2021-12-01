import { JsonObject, JsonProperty } from 'json2typescript';
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('TestingResultItem')
export class TestingResultItem {
  @JsonProperty('owner', String)
  owner = '';

  @JsonProperty('test_result', String)
  testResult = '';

  @JsonProperty('test_date', DateConverter)
  testDate: Date = null;
}
