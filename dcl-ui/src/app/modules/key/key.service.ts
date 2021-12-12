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
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { ResultKeyInfos } from '../../shared/models/key/result-key-infos';
import { KeyInfo } from '../../shared/models/key/key-info';

@Injectable({
  providedIn: 'root'
})
export class KeyService {

  private keyBaseUrl = 'key';

  constructor(
    private http: HttpExtensionsService) {
  }

  getKeyInfos(): Observable<ResultKeyInfos> {
    return this.http.get(this.keyBaseUrl, ResultKeyInfos);
  }

  getKeyInfo(name: string): Observable<KeyInfo> {
    return this.http.get(`${this.keyBaseUrl}/${name}`, KeyInfo);
  }
}
