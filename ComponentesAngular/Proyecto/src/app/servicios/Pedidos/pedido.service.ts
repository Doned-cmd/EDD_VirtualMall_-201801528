import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from "rxjs";

import { ProductosPedido } from "../../models/producto/producto";

@Injectable({
  providedIn: 'root'
})
export class PedidoService {

  constructor(private http: HttpClient) { }
  
  ObtenerAniosPedido():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<number[]>(baseURL + 'ObtenerAniosPedido', httpOptions);
  }

  ObtenerMesesPedido(anio:number):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'ObtenerAnioPedido', anio, httpOptions);
  }
  GenerarReporteMeses(anio:number):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + '', anio, httpOptions);
  }

  DevolverMesesPedido():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<number[]>(baseURL + 'DevolverMesesPedido', httpOptions);
  }
  
  EstablecerMesBack(Mes:number):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<number[]>(baseURL + 'DevolverMesesPedido', Mes,httpOptions);
  }

  

  DevolverLista():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<ProductosPedido[]>(baseURL + 'DevolverDias', httpOptions);
  }
}

