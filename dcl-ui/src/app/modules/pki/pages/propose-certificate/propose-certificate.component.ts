import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { NgForm, NgModel } from '@angular/forms';
import { TxService } from '../../../tx/tx.service';
import { PemCertificate } from 'src/app/shared/models/pki/pem-certificate';
import { MessageProposeAddX509RootCert } from 'src/app/shared/models/pki/message-propose-x509-root-certificate';

@Component({
  selector: 'app-propose-certificate',
  templateUrl: './propose-certificate.component.html',
  styleUrls: ['./propose-certificate.component.scss']
})
export class ProposeCertificateComponent implements OnInit {

  item: PemCertificate = new PemCertificate();

  showValidation = false;

  constructor(private txService: TxService) {}

  ngOnInit() {}

  onSubmit(form: NgForm, pemCertificate: PemCertificate) {
    if (!form.valid) {
      this.showValidation = true;
      return;
    }

    const message = new MessageProposeAddX509RootCert(pemCertificate);
    this.txService.goPreview(message, '/proposed-certificates');
  }

  getValidityClasses(model: NgModel) {
    return {
      'is-valid': (model.touched || this.showValidation) && model.valid,
      'is-invalid': (model.touched || this.showValidation) && model.invalid
    };
  }
}
