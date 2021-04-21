import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";

import { Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class PaqueteriaService {

  constructor(private http: HttpClient) { }

  CargarGrafos(Grafos: String):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'CargarEntregas', Grafos, httpOptions);
  }

  
}
