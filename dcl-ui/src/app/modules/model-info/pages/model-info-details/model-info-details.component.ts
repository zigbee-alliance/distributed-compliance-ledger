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
import { Observable, of, forkJoin } from 'rxjs';
import { ModelInfo } from '../../../../shared/models/model-info/model-info';
import { ModelInfoService } from '../../model-info.service';
import { ModelComplianceService } from '../../../model-compliance/model-compliance.service';
import { ActivatedRoute } from '@angular/router';
import { ComplianceInfo } from 'src/app/shared/models/model-compliance/compliance-info';
import { map, catchError } from 'rxjs/operators';

@Component({
  selector: 'app-model-info-details',
  templateUrl: './model-info-details.component.html',
  styleUrls: ['./model-info-details.component.scss']
})
export class ModelInfoDetailsComponent implements OnInit {

  item$: Observable<ModelInfo>;
  certificationInfo$: Observable<ComplianceInfo>;

  showTestingResult = false;

  constructor(public modelInfoService: ModelInfoService,
              private complianceService: ModelComplianceService,
              private route: ActivatedRoute) {
  }

  ngOnInit() {
    const params = this.route.snapshot.paramMap;
    const queryParamMap = this.route.snapshot.queryParamMap;

    this.item$ = this.modelInfoService.getModelInfo(params.get('vid'), params.get('pid'), queryParamMap.get('prev_height'));
    this.certificationInfo$ = this.complianceService.getComplianceInfo(
      params.get('vid'), params.get('pid'), queryParamMap.get('prev_height'))
      .pipe(
        catchError(() => of(new ComplianceInfo(params.get('vid'), params.get('pid')))),
      );
  }
}
