import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from "rxjs";

import { Tienda } from "../../models/tienda/tienda";

@Injectable({
  providedIn: 'root'
})
export class GettiendasService {

  constructor(private http: HttpClient) { }

  getListaTiendas():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<Tienda[]>(baseURL + 'getListaTiendas', httpOptions);
  }

  mandarTienda(tiendaActual: Tienda):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      }),
    };
    return this.http.post<Tienda>(baseURL + 'getTiendaActual', tiendaActual, httpOptions);
  }
}
