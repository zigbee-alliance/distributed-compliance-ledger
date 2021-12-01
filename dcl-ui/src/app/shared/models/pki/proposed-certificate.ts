import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('ProposedCertificate')
export class ProposedCertificate {

  @JsonProperty('pem_cert', String)
  pemCert = '';

  @JsonProperty('subject', String)
  subject = '';

  @JsonProperty('subject_key_id', String)
  subjectKeyId = '';

  @JsonProperty('serial_number', String)
  serialNumber = '';

  @JsonProperty('approvals', [String])
  approvals: string[] = [];

  @JsonProperty('owner', String)
  owner = '';
}
