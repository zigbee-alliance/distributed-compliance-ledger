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
import { HttpParams } from '@angular/common/http';
import { ResultAccountHeaders } from '../../shared/models/account/result-account-headers';
import { Account } from '../../shared/models/account/acount';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  private accountBaseUrl = 'auth/accounts';

  constructor(
    private http: HttpExtensionsService) {
  }

  getAccountHeaders(skip: number = 0, take: number = 0): Observable<ResultAccountHeaders> {
    const params = new HttpParams()
      .append('skip', skip.toString())
      .append('take', take.toString());

    return this.http.get(this.accountBaseUrl, ResultAccountHeaders, params);
  }

  getAccount(address: string): Observable<Account> {
    return this.http.get(`${this.accountBaseUrl}/${address}`, Account);
  }
}
