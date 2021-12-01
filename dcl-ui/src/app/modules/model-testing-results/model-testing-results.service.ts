import { Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { TestingResults } from '../../shared/models/model-testing-results/testing-results';
import { HttpParams } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ModelTestingResultsService {

  private baseUrl = 'compliancetest/testresults';

  constructor(
    private http: HttpExtensionsService) {
  }

  getTestingResults(vid: string, pid: string): Observable<TestingResults> {
    const params = new HttpParams().append('prev_height', 'true');
    return this.http.get(`${this.baseUrl}/${vid}/${pid}`, TestingResults, params);
  }
}
