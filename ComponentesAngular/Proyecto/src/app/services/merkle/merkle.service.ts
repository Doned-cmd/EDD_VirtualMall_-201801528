import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";

import { Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class MerkleService {

  constructor(private http: HttpClient) { }

  CambiarIntervalo(Grafos: number):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'CambiarReloj', Grafos, httpOptions);
  }

  GenerarReportes():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<any>(baseURL + 'GenerarReportesMerkle', httpOptions);
  }
}
