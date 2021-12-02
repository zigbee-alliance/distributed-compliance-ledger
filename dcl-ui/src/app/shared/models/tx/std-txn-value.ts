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
import { StdFee } from './std-fee';
import { MessageArrayConverter } from '../../json-converters/message-array-converter';
import { Message } from './message';

@JsonObject('StdTxnValue')
export class StdTxnValue {

  @JsonProperty('fee', StdFee)
  fee: StdFee = new StdFee();

  @JsonProperty('msg', MessageArrayConverter)
  messages: Message[] = [];

  @JsonProperty('memo', String)
  memo = '';

  constructor(init?: Partial<StdTxnValue>) {
    Object.assign(this, init);
  }
}
