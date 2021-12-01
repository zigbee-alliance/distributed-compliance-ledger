import { Component, OnInit } from '@angular/core';
import { ModelInfoService } from '../../model-info.service';
import { Observable, of } from 'rxjs';
import { ModelInfoHeader } from '../../../../shared/models/model-info/model-info-header';
import { pluck, share } from 'rxjs/operators';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ModelInfoSearch } from 'src/app/shared/models/model-info/model-info-search';
import { ModelInfoSearchComponent } from '../../components/certificates-search/model-info-search.component';

@Component({
  selector: 'app-model-info-list',
  templateUrl: './model-info-list.component.html',
  styleUrls: ['./model-info-list.component.scss']
})
export class ModelInfoListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<ModelInfoHeader[]>;

  constructor(public modelInfoService: ModelInfoService,
              private router: Router,
              private modalService: NgbModal) {
  }

  ngOnInit() {
    this.getModelInfoHeaders();
  }

  getModelInfoHeaders(): void {
    const source = this.modelInfoService.getModelInfoHeaders().pipe(share());

    this.total$ = source.pipe(pluck('total'));
    this.items$ = source.pipe(pluck('items'));
  }

  findModel() {
    this.modalService.open(ModelInfoSearchComponent, { size: 'sm' }).
      result.then((result: ModelInfoSearch) => {
        if (result.vid > 0 && result.pid > 0) {
          this.router.navigate([`model-info/${result.vid}/${result.pid}`], {queryParams : {prev_height: true}});
        }
      });
  }
}
