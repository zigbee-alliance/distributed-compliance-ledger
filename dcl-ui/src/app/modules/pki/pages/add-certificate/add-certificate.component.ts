import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { NgForm, NgModel } from '@angular/forms';
import { TxService } from '../../../tx/tx.service';
import { PemCertificate } from 'src/app/shared/models/pki/pem-certificate';
import { MessageAddX509Cert } from 'src/app/shared/models/pki/message-add-x509-certificate';

@Component({
  selector: 'app-add-certificate',
  templateUrl: './add-certificate.component.html',
  styleUrls: ['./add-certificate.component.scss']
})
export class AddCertificateComponent implements OnInit {

  item$: Observable<PemCertificate>;

  showValidation = false;

  constructor(private txService: TxService) {}

  ngOnInit() {}

  onSubmit(form: NgForm, pemCertificate: PemCertificate) {
    if (!form.valid) {
      this.showValidation = true;
      return;
    }

    const message = new MessageAddX509Cert(pemCertificate);
    this.txService.goPreview(message, '/certificates');
  }

  getValidityClasses(model: NgModel) {
    return {
      'is-valid': (model.touched || this.showValidation) && model.valid,
      'is-invalid': (model.touched || this.showValidation) && model.invalid
    };
  }

}
