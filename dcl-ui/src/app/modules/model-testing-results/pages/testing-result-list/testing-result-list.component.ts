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
import { pluck, share, switchMap } from 'rxjs/operators';
import { TestingResultItem } from 'src/app/shared/models/model-testing-results/testing-result-item';
import { TestingResults } from 'src/app/shared/models/model-testing-results/testing-results';
import { HttpErrorResponse } from '@angular/common/http';
import { ActivatedRoute } from '@angular/router';
import { ModelTestingResultsService } from '../../model-testing-results.service';
import { of } from 'rxjs';
import { ModalWindowComponent } from 'src/app/core/modal/modal.component';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-testing-result-list',
  templateUrl: './testing-result-list.component.html',
  styleUrls: ['./testing-result-list.component.scss'],
})
export class TestingResultListComponent implements OnInit {
  testingResults$: Observable<TestingResultItem[]>;

  constructor(
    public testingResultService: ModelTestingResultsService,
    private route: ActivatedRoute,
    private modalService: NgbModal
  ) {}

  ngOnInit() {
    this.getTestingResult();
  }

  getTestingResult(): void {
    this.testingResults$ = of(undefined);

    this.route.paramMap
    .pipe(
      switchMap(params =>
        this.testingResultService.getTestingResults(params.get('vid'), params.get('pid'))
      )
    )
    .subscribe(
      (testingResults: TestingResults) => {
        this.testingResults$ = of(testingResults.results);
      },
      (error: HttpErrorResponse) => {
        this.testingResults$ = of([]);
      }
    );
  }

  viewTestingResult(testingResult: string) {
    const modalRef = this.modalService.open(ModalWindowComponent, { size: 'lg' });
    modalRef.componentInstance.content = testingResult;
    modalRef.componentInstance.header = 'Testing Result';
  }
}
