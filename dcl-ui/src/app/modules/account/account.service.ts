import { Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { HttpParams } from '@angular/common/http';
import { ResultAccountHeaders } from '../../shared/models/account/result-account-headers';
import { Account } from '../../shared/models/account/acount';

@Injectable({
  providedIn: 'root'
})
export class AccountService {

  private accountBaseUrl = 'auth/accounts';

  constructor(
    private http: HttpExtensionsService) {
  }

  getAccountHeaders(skip: number = 0, take: number = 0): Observable<ResultAccountHeaders> {
    const params = new HttpParams()
      .append('skip', skip.toString())
      .append('take', take.toString());

    return this.http.get(this.accountBaseUrl, ResultAccountHeaders, params);
  }

  getAccount(address: string): Observable<Account> {
    return this.http.get(`${this.accountBaseUrl}/${address}`, Account);
  }
}
