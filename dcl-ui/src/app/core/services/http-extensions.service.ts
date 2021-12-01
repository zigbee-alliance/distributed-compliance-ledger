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
