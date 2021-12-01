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
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('ComplianceInfo')
export class ComplianceInfo {
  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('owner', String)
  owner = '';

  @JsonProperty('state', String, true)
  state = null;

  @JsonProperty('date', DateConverter)
  date: Date = new Date();

  @JsonProperty('certification_type', String, true)
  certificationType = 'zb';

  @JsonProperty('reason', String, true)
  reason = null;

  constructor(vid?: string, pid?: string) {
    if (vid && vid.length > 0) {
      this.vid = parseInt(vid, 10);
    }
    if (pid && pid.length > 0) {
      this.pid = parseInt(pid, 10);
    }
  }

  isCertified(): boolean {
    return this.state === 'certified';
  }
}
