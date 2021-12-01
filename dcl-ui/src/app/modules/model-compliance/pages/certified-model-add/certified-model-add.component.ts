import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ActivatedRoute } from '@angular/router';
import { NgForm, NgModel } from '@angular/forms';
import { JsonConvert } from 'json2typescript';
import { TxService } from '../../../tx/tx.service';
import { map } from 'rxjs/operators';
import { ComplianceInfo } from 'src/app/shared/models/model-compliance/compliance-info';
import { MsgCertifyModel } from 'src/app/shared/models/model-compliance/message-certify-model';
import { ModelComplianceService } from '../../model-compliance.service';

@Component({
  selector: 'app-certified-model-add',
  templateUrl: './certified-model-add.component.html',
  styleUrls: ['./certified-model-add.component.scss']
})
export class CertifiedModelAddComponent implements OnInit {

  item$: Observable<ComplianceInfo>;
  certificationTypes = ModelComplianceService.certificationTypes;

  editIDs = true;
  showValidation = false;

  jsonConvert = new JsonConvert();

  constructor(private route: ActivatedRoute,
              private txService: TxService) {
  }

  ngOnInit() {
    this.item$ = this.route.paramMap.pipe(
      map(paramMap => {
        if (paramMap.get('vid') && paramMap.get('vid').length > 0) {
          this.editIDs = false;
        }
        return new ComplianceInfo(paramMap.get('vid'), paramMap.get('pid'));
      })
    );
  }

  onSubmit(form: NgForm, model: ComplianceInfo) {
    if (!form.valid) {
      this.showValidation = true;
      return;
    }

    const message = new MsgCertifyModel(model);
    this.txService.goPreview(message, '/model-info/' + model.vid + '/' + model.pid);
  }

  getValidityClasses(model: NgModel) {
    return {
      'is-valid': (model.touched || this.showValidation) && model.valid,
      'is-invalid': (model.touched || this.showValidation) && model.invalid
    };
  }

}
