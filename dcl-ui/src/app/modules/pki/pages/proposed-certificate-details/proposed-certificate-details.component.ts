import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ProposedCertificate } from '../../../../shared/models/pki/proposed-certificate';
import { PkiService } from '../../pki.service';
import { ActivatedRoute } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ModalWindowComponent } from '../../../../core/modal/modal.component';

@Component({
  selector: 'app-proposed-certificate-details',
  templateUrl: './proposed-certificate-details.component.html',
  styleUrls: ['./proposed-certificate-details.component.scss']
})
export class ProposedCertificateDetailsComponent implements OnInit {

  item$: Observable<ProposedCertificate>;

  constructor(public pkiService: PkiService,
              private route: ActivatedRoute,
              private modalService: NgbModal) {
  }

  ngOnInit() {
    this.item$ = this.route.paramMap.pipe(
      switchMap(params => this.pkiService.getProposedX509RootCert(params.get('subject'), params.get('subjectKeyId')))
    );
  }

  viewCertificate(pemCert: string) {
    const modalRef = this.modalService.open(ModalWindowComponent, { size: 'lg' });
    modalRef.componentInstance.content = pemCert;
    modalRef.componentInstance.header = 'Proposed Root Certificate';
  }
}
