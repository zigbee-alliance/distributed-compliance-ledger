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
import { Observable, of } from 'rxjs';
import { ModelInfo } from '../../../../shared/models/model-info/model-info';
import { ModelInfoService } from '../../model-info.service';
import { ActivatedRoute, Router } from '@angular/router';
import { NgForm, NgModel } from '@angular/forms';
import { JsonConvert } from 'json2typescript';
import { MessageAddModelInfo } from '../../../../shared/models/model-info/message-add-model-info';
import { TxService } from '../../../tx/tx.service';
import { MessageUpdateModelInfo } from '../../../../shared/models/model-info/message-update-model-info';
import { AccountService } from '../../../account/account.service';
import { Account } from '../../../../shared/models/account/acount';
import { pluck } from 'rxjs/operators';

@Component({
  selector: 'app-model-info-edit',
  templateUrl: './model-info-edit.component.html',
  styleUrls: ['./model-info-edit.component.scss']
})
export class ModelInfoEditComponent implements OnInit {

  item$: Observable<ModelInfo>;
  accounts$: Observable<Account[]>;

  isNew: boolean;
  showValidation = false;

  jsonConvert = new JsonConvert();

  constructor(private modelInfoService: ModelInfoService,
              private accountService: AccountService,
              private route: ActivatedRoute,
              private txService: TxService) {
  }

  ngOnInit() {
    this.isNew = this.route.snapshot.url.reverse()[0].path === 'new';

    if (this.isNew) {
      this.item$ = of(new ModelInfo());
    } else {
      const params = this.route.snapshot.paramMap;
      const queryParamMap = this.route.snapshot.queryParamMap;
      this.item$ = this.modelInfoService.getModelInfo(params.get('vid'), params.get('pid'), queryParamMap.get('prev_height'));
    }

    this.accounts$ = this.accountService.getAccountHeaders().pipe(
      pluck('items')
    );
  }

  onSubmit(form: NgForm, modelInfo: ModelInfo) {
    if (!form.valid) {
      this.showValidation = true;
      return;
    }

    const message = this.isNew ?
      new MessageAddModelInfo(modelInfo) :
      new MessageUpdateModelInfo(modelInfo);

    this.txService.goPreview(message, '/model-info/' + modelInfo.vid + '/' + modelInfo.pid);
  }

  getValidityClasses(model: NgModel) {
    return {
      'is-valid': (model.touched || this.showValidation) && model.valid,
      'is-invalid': (model.touched || this.showValidation) && model.invalid
    };
  }

}
