import { Pipe, PipeTransform } from '@angular/core';
import {KeyInfo} from '../models/key/key-info';

@Pipe({name: 'keyinfo'})
export class KeyInfoPipe implements PipeTransform {
  transform(user: KeyInfo): string {
    return user.name + ' (' + user.address + ')';
  }
}
