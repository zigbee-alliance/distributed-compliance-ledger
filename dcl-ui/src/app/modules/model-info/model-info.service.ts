import { Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { HttpParams } from '@angular/common/http';
import { ResultModelInfoHeaders } from '../../shared/models/model-info/result-model-info-headers';
import { ModelInfo } from '../../shared/models/model-info/model-info';
import { TxService } from '../tx/tx.service';
import { MessageDeleteModelInfo } from '../../shared/models/model-info/message-delete-model-info';

@Injectable({
  providedIn: 'root'
})
export class ModelInfoService {

  private baseUrl = 'modelinfo/models';

  constructor(
    private http: HttpExtensionsService,
    private txService: TxService) {
  }

  getModelInfoHeaders(skip: number = 0, take: number = 0): Observable<ResultModelInfoHeaders> {
    const params = new HttpParams()
      .append('skip', skip.toString())
      .append('take', take.toString());

    return this.http.get(this.baseUrl, ResultModelInfoHeaders, params);
  }

  getModelInfo(vid: string, pid: string, prevHeight: string): Observable<ModelInfo> {
    const params = prevHeight ? new HttpParams().append('prev_height', prevHeight) : null;
    return this.http.get(`${this.baseUrl}/${vid}/${pid}`, ModelInfo, params);
  }

  goDeleteModelInfo(vid: string, pid: string) {
    const message = new MessageDeleteModelInfo({vid: parseInt(vid, 10), pid: parseInt(pid, 10)});
    this.txService.goPreview(message, '/model-info');
  }
}
