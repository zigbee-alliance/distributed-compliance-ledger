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
import { ActivatedRoute } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { KeyInfo } from '../../../../shared/models/key/key-info';
import { KeyService } from '../../key.service';

@Component({
  selector: 'app-key-details',
  templateUrl: './key-details.component.html',
  styleUrls: ['./key-details.component.scss']
})
export class KeyDetailsComponent implements OnInit {

  item$: Observable<KeyInfo>;

  constructor(private route: ActivatedRoute,
              private keyService: KeyService) {
  }

  ngOnInit() {
    this.item$ = this.route.paramMap.pipe(
      switchMap(params => this.keyService.getKeyInfo(params.get('name')))
    );
  }

}
