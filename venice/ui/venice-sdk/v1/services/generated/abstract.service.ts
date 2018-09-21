/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */

import { HttpClient } from '../../../../webapp/node_modules/@angular/common/http';
import { Observable } from '../../../../webapp/node_modules/rxjs/Observable';
import { Injectable } from '../../../../webapp/node_modules/@angular/core';

@Injectable()
export class AbstractService {
  protected O_Tenant: string = "";

  constructor(protected _http: HttpClient) {
  }

  /**
   * Get the service class-name
   */
  getClassName(): string {
    return 'base_service';
  }

  // This function should be overwritten
  protected invokeAJAX(method: string, url: string, payload: any, eventPayloadID: any, forceReal: boolean = false): Observable<any> {
    let observer;
    switch (method) {
      case 'GET':
        if (payload == null) {
          observer = this._http.get(url, { observe: 'response' });
        } else {
          observer = this._http.get(url, { params: payload, observe: 'response' });
        }
        break;
      case 'POST':
        observer = this._http.post(url, payload, { observe: 'response' });
        break;
      case 'PUT':
        observer = this._http.put(url, payload, { observe: 'response' });
        break;
      case 'DELETE':
        observer = this._http.delete(url, { observe: 'response' });
        break;
      default:
        break;
    }
    return observer;
  }

  protected invokeAJAXGetCall(url: string, payload: any, eventPayloadID: any, forceReal: boolean = false): Observable<any> {
    return this.invokeAJAX('GET', url, payload, eventPayloadID, forceReal);
  }

  protected invokeAJAXPostCall(url: string, payload: any, eventPayloadID: any, forceReal: boolean = false): Observable<any> {
    return this.invokeAJAX('POST', url, payload, eventPayloadID, forceReal);
  }

  protected invokeAJAXPutCall(url: string, payload: any, eventPayloadID: any, forceReal: boolean = false): Observable<any> {
    return this.invokeAJAX('PUT', url, payload, eventPayloadID, forceReal);
  }

  protected invokeAJAXDeleteCall(url: string, eventPayloadID: any, forceReal: boolean = false): Observable<any> {
    return this.invokeAJAX('DELETE', url, null, eventPayloadID, forceReal);
  }
}
