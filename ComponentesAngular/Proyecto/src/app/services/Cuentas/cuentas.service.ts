import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";

import { Observable } from "rxjs";

import { Usuario,MostrarUsuario } from "../../models/Usuario/usuario";


@Injectable({
  providedIn: 'root'
})
export class CuentasService {

  constructor(private http: HttpClient) { }

  CompararDatos(Comparable: Usuario):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<boolean>(baseURL + 'VerificarCuenta', Comparable, httpOptions);
  }

  VerificarSiAdmin():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<boolean>(baseURL + 'VerificarSiAdmin', httpOptions);
  }

  RegresarCuentaLoged():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<MostrarUsuario>(baseURL + 'DevolverCuenta', httpOptions);
  }

  
}
