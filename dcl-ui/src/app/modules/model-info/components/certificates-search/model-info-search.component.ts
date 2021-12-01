import { Component } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ModelInfoSearch } from 'src/app/shared/models/model-info/model-info-search';

@Component({
  selector: 'app-model-info-search',
  templateUrl: './model-info-search.component.html',
  styleUrls: ['./model-info-search.component.scss']
})
export class ModelInfoSearchComponent {
  item: ModelInfoSearch = new ModelInfoSearch();

  constructor(public activeModal: NgbActiveModal) {}
}
