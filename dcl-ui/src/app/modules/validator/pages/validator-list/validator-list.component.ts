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
import { Validator } from 'src/app/shared/models/validator/validator';
import { ValidatorService } from '../../validator.service';

@Component({
  selector: 'app-validator-list',
  templateUrl: './validator-list.component.html',
  styleUrls: ['./validator-list.component.scss']
})
export class ValidatorListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<Validator[]>;

  constructor(public validatorService: ValidatorService) {
  }

  ngOnInit() {
    this.getValidatorHeaders();
  }

  getValidatorHeaders(): void {
    const source = this.validatorService.getAllValidators().pipe(share());

    this.total$ = source.pipe(pluck('total'));
    this.items$ = source.pipe(pluck('items'));
  }

}
