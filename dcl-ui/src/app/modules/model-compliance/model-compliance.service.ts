import { Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { ComplianceInfo } from '../../shared/models/model-compliance/compliance-info';
import { ComplianceInfoRecords } from '../../shared/models/model-compliance/compliance-info-records';
import { HttpParams } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ModelComplianceService {
  static certificationTypes = ['zb'];

  private certificationType = 'zb';
  private baseUrl = 'compliance';

  constructor(
    private http: HttpExtensionsService) {
  }

  getComplianceInfo(vid: string, pid: string, prevHeight: string): Observable<ComplianceInfo> {
    const params = prevHeight ? new HttpParams().append('prev_height', prevHeight) : null;
    return this.http.get(`${this.baseUrl}/${vid}/${pid}/${this.certificationType}`, ComplianceInfo, params);
  }

  getAllComplianceInfos(): Observable<ComplianceInfoRecords> {
    return this.http.get(`${this.baseUrl}`, ComplianceInfoRecords);
  }
}
