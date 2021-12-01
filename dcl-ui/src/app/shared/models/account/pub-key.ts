import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('PubKey')
export class PubKey {
    @JsonProperty('type', String)
    type: string = undefined;

    @JsonProperty('value', String)
    value: string = undefined;
}
