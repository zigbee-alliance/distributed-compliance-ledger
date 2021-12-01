import { Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { ValidatorRecords } from 'src/app/shared/models/validator/validator-records';

@Injectable({
  providedIn: 'root'
})
export class ValidatorService {
  private baseUrl = 'validators';

  constructor(
    private http: HttpExtensionsService) {
  }

  getAllValidators(): Observable<ValidatorRecords> {
    return this.http.get(`${this.baseUrl}?state=active`, ValidatorRecords);
  }
}
