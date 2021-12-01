import { Component, OnInit } from '@angular/core';
import { PkiService } from '../../pki.service';
import { Observable } from 'rxjs';
import { pluck, share } from 'rxjs/operators';
import { ApprovedCertificate } from 'src/app/shared/models/pki/approved-certificate';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CertificateSearch } from 'src/app/shared/models/pki/certificate-search';
import { CertificatesSearchComponent } from '../../components/certificates-search/certificates-search.component';

@Component({
  selector: 'app-approved-certificates-list',
  templateUrl: './approved-certificates-list.component.html',
  styleUrls: ['./approved-certificates-list.component.scss']
})
export class ApprovedCertificatesListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<ApprovedCertificate[]>;

  constructor(public pkiService: PkiService,
              private modalService: NgbModal) {
  }

  ngOnInit() {
    this.getCertificates();
  }

  getCertificates(): void {
    const source = this.pkiService.getAllX509Certs().pipe(share());

    this.total$ = source.pipe(pluck('total'));
    this.items$ = source.pipe(pluck('items'));
  }

  findCertificate() {
    this.modalService.open(CertificatesSearchComponent, { size: 'sm' }).
      result.then((result: CertificateSearch) => {
        let source: any;
        if (result.subject.length > 0 && result.subjectKeyId.length > 0) {
          source = this.pkiService.getX509Cert(result.subject, result.subjectKeyId);
        } else if (result.subject.length > 0) {
          source = this.pkiService.getAllSubjectX509Certs(result.subject);
        } else {
          source = this.pkiService.getAllX509Certs();
        }

        this.total$ = source.pipe(pluck('total'));
        this.items$ = source.pipe(pluck('items'));

      });
  }
}
