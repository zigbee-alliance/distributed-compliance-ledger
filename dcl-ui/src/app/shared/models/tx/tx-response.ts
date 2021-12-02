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
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';
import { ABCIMessageLog } from './abci-message-log';

@JsonObject('TxResponse')
export class TxResponse {
  @JsonProperty('height', AminoNumberConverter)
  height: number = undefined;

  @JsonProperty('txhash', String)
  txHash: string = undefined;

  @JsonProperty('code', AminoNumberConverter, true)
  code: number = undefined;

  @JsonProperty('codespace', String, true)
  codespace: string = undefined;

  @JsonProperty('logs', [ABCIMessageLog], true)
  logs: ABCIMessageLog[] = undefined;

  @JsonProperty('raw_log', String, true)
  rawLog: string = undefined;
}
