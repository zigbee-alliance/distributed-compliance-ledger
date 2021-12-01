import { Pipe, PipeTransform } from '@angular/core';
import {KeyInfo} from '../models/key/key-info';
import { Account } from '../models/account/acount';

@Pipe({name: 'account'})
export class AccountPipe implements PipeTransform {
  transform(account: Account): string {
    return account.address + ' (' + account.pubKey + ')';
  }
}
