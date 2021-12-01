import { JsonObject, JsonProperty } from 'json2typescript';
import { TestingResultItem } from './testing-result-item';


@JsonObject('TestingResults')
export class TestingResults {
  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('results', [TestingResultItem])
  results: TestingResultItem[] = [];
}
