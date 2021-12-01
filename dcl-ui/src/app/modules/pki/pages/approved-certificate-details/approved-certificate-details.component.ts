import { Component, OnInit } from '@angular/core';
import { Observable, from } from 'rxjs';
import { ProposedCertificate } from '../../../../shared/models/pki/proposed-certificate';
import { PkiService } from '../../pki.service';
import { ActivatedRoute } from '@angular/router';
import { switchMap, map, pluck } from 'rxjs/operators';
import { ApprovedCertificate } from 'src/app/shared/models/pki/approved-certificate';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ModalWindowComponent } from '../../../../core/modal/modal.component';

@Component({
  selector: 'app-approved-certificate-details',
  templateUrl: './approved-certificate-details.component.html',
  styleUrls: ['./approved-certificate-details.component.scss']
})
export class ApprovedCertificateDetailsComponent implements OnInit {

  items$: Observable<ApprovedCertificate[]>;

  constructor(public pkiService: PkiService,
              private route: ActivatedRoute,
              private modalService: NgbModal) {
  }

  ngOnInit() {
    const source = this.route.paramMap.pipe(
      switchMap(params =>
        this.pkiService.getX509Cert(params.get('subject'), params.get('subjectKeyId'))
      )
    );

    this.items$ = source.pipe(pluck('items'));
  }

  viewCertificate(pemCert: string) {
    const modalRef = this.modalService.open(ModalWindowComponent, { size: 'lg' });
    modalRef.componentInstance.content = pemCert;
    modalRef.componentInstance.header = 'Certificate';
  }
}
