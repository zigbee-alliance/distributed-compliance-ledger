/**
 * Copyright 2020 DSR Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { JsonConvert } from 'json2typescript';

@Injectable({
  providedIn: 'root'
})
export class HttpExtensionsService {

  private jsonConvert: JsonConvert = new JsonConvert();

  constructor(
    private http: HttpClient) {
  }

  private static removeWrapper(jsonObj: any): any {
    if (jsonObj.height && jsonObj.result) {
      return jsonObj.result;
    }

    return jsonObj;
  }

  getAny(url: string, params?: HttpParams): Observable<any> {
    return this.http.get(url, {
      responseType: 'text',
      params
    }).pipe(map(text => {
      const jsonObj: object = JSON.parse(text);
      return HttpExtensionsService.removeWrapper(jsonObj);
    }));
  }

  get<T>(url: string, classReference: (new() => T), params?: HttpParams): Observable<T> {
    return this.getAny(url, params).pipe(map(jsonObj => {
      return this.jsonConvert.deserializeObject(jsonObj, classReference);
    }));
  }

  postAny(url: string, body: any | null, params?: HttpParams, httpHeaders?: HttpHeaders): Observable<any> {
    return this.http.post(url, body, {
      responseType: 'text',
      params,
      headers: httpHeaders
    }).pipe(map(text => {
      const jsonObj: object = JSON.parse(text);
      return HttpExtensionsService.removeWrapper(jsonObj);
    }));
  }

  post<T>(url: string, body: any | null, classReference: (new() => T), params?: HttpParams): Observable<T> {
    body = this.jsonConvert.serializeObject(body);

    return this.postAny(url, body, params).pipe(map(jsonObj => {
      return this.jsonConvert.deserializeObject(jsonObj, classReference);
    }));
  }
}
