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
import { Observable } from 'rxjs';
import { pluck, share } from 'rxjs/operators';
import { KeyService } from '../../key.service';
import { KeyInfo } from '../../../../shared/models/key/key-info';

@Component({
  selector: 'app-key-list',
  templateUrl: './key-list.component.html',
  styleUrls: ['./key-list.component.scss']
})
export class KeyListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<KeyInfo[]>;

  constructor(private keyService: KeyService) {
  }

  ngOnInit() {
    const source = this.keyService.getKeyInfos().pipe(
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
