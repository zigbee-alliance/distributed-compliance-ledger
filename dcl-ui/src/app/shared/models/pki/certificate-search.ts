import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('CertificateSearch')
export class CertificateSearch {

  @JsonProperty('subject', String)
  subject = '';

  @JsonProperty('subject_key_id', String, true)
  subjectKeyId = '';
}
