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

import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ResultBlockchainInfo } from '../../shared/models/block/result-blockchain-info';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { HttpParams } from '@angular/common/http';
import { ResultBlock } from '../../shared/models/block/result-block';

@Injectable({
  providedIn: 'root'
})
export class BlockService {

  private blockBaseUrl = 'blocks';

  constructor(
    private http: HttpExtensionsService) {
  }

  getBlockchainInfo(minHeight: number = 0, maxHeight: number = 0): Observable<ResultBlockchainInfo> {
    const params = new HttpParams()
      .append('minHeight', minHeight.toString())
      .append('maxHeight', maxHeight.toString());

    return this.http.get(this.blockBaseUrl, ResultBlockchainInfo, params);
  }

  getBlock(height: number = 0): Observable<ResultBlock> {
    return this.http.get(`${this.blockBaseUrl}/${height}`, ResultBlock);
  }
}
