import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('ApprovedCertificate')
export class ApprovedCertificate {

  @JsonProperty('pem_cert', String)
  pemCert = '';

  @JsonProperty('subject', String)
  subject = '';

  @JsonProperty('subject_key_id', String)
  subjectKeyId = '';

  @JsonProperty('serial_number', String)
  serialNumber = '';

  @JsonProperty('root_subject', String, true)
  rootSubject = '';

  @JsonProperty('root_subject_key_id', String, true)
  rootSubjectKeyId = '';

  @JsonProperty('is_root', Boolean)
  isRoot = '';

  @JsonProperty('owner', String)
  owner = '';
}
