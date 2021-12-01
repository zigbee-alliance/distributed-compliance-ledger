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
