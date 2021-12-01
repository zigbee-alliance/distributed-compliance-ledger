import { Component, OnInit } from '@angular/core';
import { PkiService } from '../../pki.service';
import { Observable } from 'rxjs';
import { ProposedCertificate } from '../../../../shared/models/pki/proposed-certificate';
import { pluck, share } from 'rxjs/operators';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CertificateSearch } from 'src/app/shared/models/pki/certificate-search';
import { Router } from '@angular/router';
import { CertificatesSearchComponent } from '../../components/certificates-search/certificates-search.component';

@Component({
  selector: 'app-proposed-certificates-list',
  templateUrl: './proposed-certificates-list.component.html',
  styleUrls: ['./proposed-certificates-list.component.scss']
})
export class ProposedCertificatesListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<ProposedCertificate[]>;

  constructor(public pkiService: PkiService,
              private router: Router,
              private modalService: NgbModal) {
  }

  ngOnInit() {
    this.getCertificates();
  }

  getCertificates(): void {
    const source = this.pkiService.getAllProposedX509RootCerts().pipe(share());

    this.total$ = source.pipe(pluck('total'));
    this.items$ = source.pipe(pluck('items'));
  }

  findCertificate() {
    this.modalService.open(CertificatesSearchComponent, { size: 'sm' }).
      result.then((result: CertificateSearch) => {
        if (result.subject.length > 0 && result.subjectKeyId.length > 0) {
          this.router.navigate([`proposed-certificates/${result.subject}/${result.subjectKeyId}`], {queryParams : {prev_height: true}});
        } else {
          this.getCertificates();
        }
      });
  }
}
