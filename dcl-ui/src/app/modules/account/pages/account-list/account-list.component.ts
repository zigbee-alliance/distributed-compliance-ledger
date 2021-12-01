import { Component, OnInit } from '@angular/core';
import { AccountService } from '../../account.service';
import { pluck, share } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { Account } from '../../../../shared/models/account/acount';

@Component({
  selector: 'app-account-list',
  templateUrl: './account-list.component.html',
  styleUrls: ['./account-list.component.scss']
})
export class AccountListComponent implements OnInit {

  total$: Observable<number>;
  items$: Observable<Account[]>;

  constructor(private accountService: AccountService) {
  }

  ngOnInit() {
    const source = this.accountService.getAccountHeaders().pipe(
      share()
    );

    this.total$ = source.pipe(
      pluck('total')
    );

    this.items$ = source.pipe(
      pluck('items')
    );
  }

}
