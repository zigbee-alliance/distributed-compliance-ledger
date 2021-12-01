import { Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { ResultKeyInfos } from '../../shared/models/key/result-key-infos';
import { KeyInfo } from '../../shared/models/key/key-info';

@Injectable({
  providedIn: 'root'
})
export class KeyService {

  private keyBaseUrl = 'key';

  constructor(
    private http: HttpExtensionsService) {
  }

  getKeyInfos(): Observable<ResultKeyInfos> {
    return this.http.get(this.keyBaseUrl, ResultKeyInfos);
  }

  getKeyInfo(name: string): Observable<KeyInfo> {
    return this.http.get(`${this.keyBaseUrl}/${name}`, KeyInfo);
  }
}
