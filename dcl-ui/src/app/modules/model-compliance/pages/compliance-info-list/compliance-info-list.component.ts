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
import { ModelComplianceService } from '../../model-compliance.service';
import { Observable } from 'rxjs';
import { ComplianceInfo } from '../../../../shared/models/model-compliance/compliance-info';
import { pluck, share } from 'rxjs/operators';

@Component({
  selector: 'app-compliance-info-list',
  templateUrl: './compliance-info-list.component.html',
  styleUrls: ['./compliance-info-list.component.scss']
})
export class ComplianceInfoListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<ComplianceInfo[]>;

  constructor(public complianceService: ModelComplianceService) {
  }

  ngOnInit() {
    this.getComplianceInfoHeaders();
  }

  getComplianceInfoHeaders(): void {
    const source = this.complianceService.getAllComplianceInfos().pipe(share());

    this.total$ = source.pipe(pluck('total'));
    this.items$ = source.pipe(pluck('items'));
  }

}
