import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('PemCertificate')
export class PemCertificate {

  @JsonProperty('cert', String)
  cert = '';
}
