import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ResultBlockchainInfo } from '../../shared/models/block/result-blockchain-info';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { HttpParams } from '@angular/common/http';
import { ResultBlock } from '../../shared/models/block/result-block';

@Injectable({
  providedIn: 'root'
})
export class BlockService {

  private blockBaseUrl = 'blocks';

  constructor(
    private http: HttpExtensionsService) {
  }

  getBlockchainInfo(minHeight: number = 0, maxHeight: number = 0): Observable<ResultBlockchainInfo> {
    const params = new HttpParams()
      .append('minHeight', minHeight.toString())
      .append('maxHeight', maxHeight.toString());

    return this.http.get(this.blockBaseUrl, ResultBlockchainInfo, params);
  }

  getBlock(height: number = 0): Observable<ResultBlock> {
    return this.http.get(`${this.blockBaseUrl}/${height}`, ResultBlock);
  }
}
