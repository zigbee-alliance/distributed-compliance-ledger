import { Component } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { CertificateSearch } from 'src/app/shared/models/pki/certificate-search';

@Component({
  selector: 'app-certificates-search',
  templateUrl: './certificates-search.component.html',
  styleUrls: ['./certificates-search.component.scss']
})
export class CertificatesSearchComponent {
  item: CertificateSearch = new CertificateSearch();

  constructor(public activeModal: NgbActiveModal) {}
}
