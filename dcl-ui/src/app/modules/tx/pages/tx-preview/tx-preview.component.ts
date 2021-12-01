import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Message } from '../../../../shared/models/tx/message';
import { JsonConvert } from 'json2typescript';
import { map, pluck, switchMap } from 'rxjs/operators';
import { ActivatedRoute, Router } from '@angular/router';
import { NgForm, NgModel } from '@angular/forms';
import { TxService } from '../../tx.service';
import { HttpErrorResponse } from '@angular/common/http';
import { TxResponse } from '../../../../shared/models/tx/tx-response';
import { ToastService } from '../../../../core/services/toast.service';
import { KeyInfo } from '../../../../shared/models/key/key-info';
import { KeyService } from '../../../key/key.service';
import { AccountService } from '../../../account/account.service';

@Component({
  selector: 'app-tx-sign-broadcast',
  templateUrl: './tx-preview.component.html',
  styleUrls: ['./tx-preview.component.scss']
})
export class TxPreviewComponent implements OnInit {

  message$: Observable<Message>;
  keyInfos$: Observable<KeyInfo[]>;

  signer: KeyInfo;

  showValidation = false;
  isBusy = false;
  isCollapsed = true;

  jsonConvert: JsonConvert = new JsonConvert();

  constructor(private route: ActivatedRoute,
              private router: Router,
              private keyService: KeyService,
              private accountService: AccountService,
              private txService: TxService,
              private toastService: ToastService) {
  }

  ngOnInit() {
    this.message$ = this.route.paramMap.pipe(
      map(paramMap => {
        const str = paramMap.get('message');
        return this.txService.decodeMessage(str);
      })
    );

    this.keyInfos$ = this.keyService.getKeyInfos().pipe(
      pluck('items')
    );
  }

  onSubmit(signatureForm: NgForm, message: Message, signer: KeyInfo) {
    if (!signatureForm.valid) {
      this.showValidation = true;
      return;
    }

    this.isBusy = true;

    message.setSigner(signer.address);

    this.accountService.getAccount(signer.address).pipe( // Used to get fresh sequence number
      map(account => this.txService.makeSignMsg(message, account)),
      switchMap(tx => this.txService.signTx(tx, signer)),
      switchMap(signedTx => this.txService.broadcastTx(signedTx))
    ).subscribe((resp: TxResponse) => {
      this.isBusy = false;
      this.handleTxResponse(resp);
    }, (error: HttpErrorResponse) => {
      this.isBusy = false;
      this.handleError(error);
    });
  }

  handleError(error: any) {
    let errorMessage = error.error || error.message || error;

    try {
      const errJson = JSON.parse(errorMessage);
      errorMessage = errJson.error || errorMessage;
    } catch (e) {
      // Ignore
    }

    this.showErrorMessage(errorMessage);
  }

  handleTxResponse(txResponse: TxResponse) {
    const success = txResponse.logs ? txResponse.logs[0].success : false;

    if (success) {
      this.showSuccessMessage();
      this.goBack();
      return;
    }

    let errorMessage = txResponse.logs ? txResponse.logs[0].log : txResponse.rawLog;

    try {
      const errJson = JSON.parse(errorMessage);
      errorMessage = errJson.message;
    } catch (e) {
      // Ignore
    }

    this.showErrorMessage(errorMessage);
  }

  showErrorMessage(message: string) {
    this.toastService.show('Error: ' + message, { classname: 'bg-danger text-light' });
  }

  showSuccessMessage() {
    this.toastService.show('Transaction successfully committed', { classname: 'bg-success text-light' });
  }

  goBack() {
    const backRoute = this.route.snapshot.paramMap.get('back');

    if (backRoute) {
      this.router.navigate([backRoute]);
    }
  }

  getValidityClasses(model: NgModel) {
    return {
      'is-valid': (model.touched || this.showValidation) && model.valid,
      'is-invalid': (model.touched || this.showValidation) && model.invalid
    };
  }
}


