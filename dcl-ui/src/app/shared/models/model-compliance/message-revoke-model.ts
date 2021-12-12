/**
 * Copyright 2020 DSR Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { JsonObject, JsonProperty } from 'json2typescript';
import { ComplianceInfo } from './compliance-info';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('MsgRevokeModel')
export class MsgRevokeModel extends Message {

  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('revocation_date', DateConverter)
  revocationDate: Date = new Date();

  @JsonProperty('certification_type', String, true)
  certificationType = null;

  @JsonProperty('reason', String, true)
  reason = null;

  @JsonProperty('signer', String)
  signer = '';

  constructor(complianceInfo?: ComplianceInfo, signer?: string) {
    super(MessageType.RevokeModel);

    if (complianceInfo) {
      this.vid = complianceInfo.vid;
      this.pid = complianceInfo.pid;
      this.revocationDate = complianceInfo.date;
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
