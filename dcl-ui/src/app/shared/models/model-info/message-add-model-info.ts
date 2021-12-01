import { JsonObject, JsonProperty } from 'json2typescript';
import { ModelInfo } from './model-info';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';

@JsonObject('MessageAddModelInfo')
export class MessageAddModelInfo extends Message {

  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('cid', Number, true)
  cid = null;

  @JsonProperty('name', String)
  name = '';

  @JsonProperty('description', String)
  description = '';

  @JsonProperty('sku', String)
  sku = '';

  @JsonProperty('firmware_version', String)
  firmwareVersion = '';

  @JsonProperty('hardware_version', String)
  hardwareVersion = '';

  @JsonProperty('custom', String, true)
  custom = null;

  @JsonProperty('tis_or_trp_testing_completed', Boolean)
  tisOrTrpTestingCompleted = false;

  @JsonProperty('signer', String)
  signer = '';

  constructor(modelInfo?: ModelInfo, signer?: string) {
    super(MessageType.AddModelInfo);

    if (modelInfo) {
      this.vid = modelInfo.vid;
      this.pid = modelInfo.pid;
      this.cid = modelInfo.cid;
      this.name = modelInfo.name;
      this.description = modelInfo.description;
      this.sku = modelInfo.sku;
      this.firmwareVersion = modelInfo.firmwareVersion;
      this.hardwareVersion = modelInfo.hardwareVersion;
      this.custom = modelInfo.custom;
      this.tisOrTrpTestingCompleted = modelInfo.tisOrTrpTestingCompleted;
    }

    if (signer) {
      this.signer = signer;
    }
  }

  setSigner(signer: string) {
    this.signer = signer;
  }
}
