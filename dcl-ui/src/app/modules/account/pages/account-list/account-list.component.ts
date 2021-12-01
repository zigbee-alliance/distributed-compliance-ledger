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

import { Component, OnInit } from '@angular/core';
import { AccountService } from '../../account.service';
import { pluck, share } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { Account } from '../../../../shared/models/account/acount';

@Component({
  selector: 'app-account-list',
  templateUrl: './account-list.component.html',
  styleUrls: ['./account-list.component.scss']
})
export class AccountListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<Account[]>;

  constructor(private accountService: AccountService) {
  }

  ngOnInit() {
    const source = this.accountService.getAccountHeaders().pipe(
      share()
    );

    this.total$ = source.pipe(
      pluck('total')
    );

    this.items$ = source.pipe(
      pluck('items')
    );
  }

}
