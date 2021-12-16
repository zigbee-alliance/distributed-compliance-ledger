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
import { DateConverter } from '../../json-converters/date-converter';

@JsonObject('BlockHeader')
export class BlockHeader {
  @JsonProperty('chain_id', String)
  chainId: string = undefined;

  @JsonProperty('height', AminoNumberConverter)
  height: number = undefined;

  @JsonProperty('time', DateConverter)
  time: Date = undefined;

  @JsonProperty('num_txs', AminoNumberConverter)
  numTxs: number = undefined;

  @JsonProperty('total_txs', AminoNumberConverter)
  totalTxs: number = undefined;
}
