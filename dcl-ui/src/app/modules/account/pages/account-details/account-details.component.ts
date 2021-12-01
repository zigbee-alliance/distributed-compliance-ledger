import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { ActivatedRoute } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { AccountService } from '../../account.service';
import { Account } from '../../../../shared/models/account/acount';

@Component({
  selector: 'app-account-details',
  templateUrl: './account-details.component.html',
  styleUrls: ['./account-details.component.scss']
})
export class AccountDetailsComponent implements OnInit {

  item$: Observable<Account>;

  constructor(private route: ActivatedRoute,
              private accountService: AccountService) {
  }

  ngOnInit() {
    this.item$ = this.route.paramMap.pipe(
      switchMap(params => this.accountService.getAccount(params.get('addr')))
    );
  }

}
