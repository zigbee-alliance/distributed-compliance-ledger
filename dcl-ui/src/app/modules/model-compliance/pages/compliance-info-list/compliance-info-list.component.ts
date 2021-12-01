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
