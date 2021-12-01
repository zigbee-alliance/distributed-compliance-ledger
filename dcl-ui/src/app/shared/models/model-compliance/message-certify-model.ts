import { JsonObject, JsonProperty } from 'json2typescript';
import { ComplianceInfo } from './compliance-info';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('MsgCertifyModel')
export class MsgCertifyModel extends Message {

  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('certification_date', DateConverter)
  certificationDate: Date = new Date();

  @JsonProperty('certification_type', String, true)
  certificationType = 'null';

  @JsonProperty('reason', String, true)
  reason = null;

  @JsonProperty('signer', String)
  signer = '';

  constructor(complianceInfo?: ComplianceInfo, signer?: string) {
    super(MessageType.CertifyModel);

    if (complianceInfo) {
      this.vid = complianceInfo.vid;
      this.pid = complianceInfo.pid;
      this.certificationDate = complianceInfo.date;
      this.certificationType = complianceInfo.certificationType;
    }

    if (signer) {
      this.signer = signer;
    }
  }

  setSigner(signer: string) {
    this.signer = signer;
  }
}
