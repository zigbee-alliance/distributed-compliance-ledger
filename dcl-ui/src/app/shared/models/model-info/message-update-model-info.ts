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
import { ModelInfo } from './model-info';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';

@JsonObject('MessageUpdateModelInfo')
export class MessageUpdateModelInfo extends Message {

  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('cid', Number, true)
  cid = null;

  @JsonProperty('name', String)
  name = '';

  @JsonProperty('owner', String)
  owner = '';

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
    super(MessageType.UpdateModelInfo);

    if (modelInfo) {
      this.vid = modelInfo.vid;
      this.pid = modelInfo.pid;
      this.cid = modelInfo.cid;
      this.name = modelInfo.name;
      this.owner = modelInfo.owner;
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
