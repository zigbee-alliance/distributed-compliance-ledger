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
import { TestingResult } from '../../../../shared/models/model-testing-results/testing-result';
import { MsgAddTestingResult } from '../../../../shared/models/model-testing-results/message-add-testing-result';
import { ActivatedRoute } from '@angular/router';
import { NgForm, NgModel } from '@angular/forms';
import { JsonConvert } from 'json2typescript';
import { TxService } from '../../../tx/tx.service';
import { map } from 'rxjs/operators';

@Component({
  selector: 'app-testing-result-add',
  templateUrl: './testing-result-add.component.html',
  styleUrls: ['./testing-result-add.component.scss']
})
export class TestingResultAddComponent implements OnInit {

  item$: Observable<TestingResult>;

  showValidation = false;

  jsonConvert = new JsonConvert();

  constructor(private route: ActivatedRoute,
              private txService: TxService) {
  }

  ngOnInit() {
    this.item$ = this.route.paramMap.pipe(
      map(paramMap => {
        return new TestingResult(paramMap.get('vid'), paramMap.get('pid'));
      })
    );
  }

  onSubmit(form: NgForm, testingResult: TestingResult) {
    if (!form.valid) {
      this.showValidation = true;
      return;
    }

    const message = new MsgAddTestingResult(testingResult);
    this.txService.goPreview(message, '/model-info/' + testingResult.vid + '/' + testingResult.pid);
  }

  getValidityClasses(model: NgModel) {
    return {
      'is-valid': (model.touched || this.showValidation) && model.valid,
      'is-invalid': (model.touched || this.showValidation) && model.invalid
    };
  }

}
