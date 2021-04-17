import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";

import { Observable } from "rxjs";

import { Usuario,MostrarUsuario, UsuarioLogeado } from "../../models/Usuario/usuario";


@Injectable({
  providedIn: 'root'
})
export class CuentasService {

  constructor(private http: HttpClient) { }

  CargarUsuarios(tiendaActual: String):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'CargarUsuarios', tiendaActual, httpOptions);
  }

  CompararDatos(Comparable: Usuario):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<boolean>(baseURL + 'VerificarCuenta', Comparable, httpOptions);
  }

  RegistrarUsuario(Comparable: UsuarioLogeado):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<boolean>(baseURL + 'RegistrarUsuario', Comparable, httpOptions);
  }

  EliminarCuenta(Comparable:Usuario):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'EliminarUsuario', Comparable, httpOptions);
  }

  VerificarSiAdmin():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'VerificarSiAdmin', httpOptions);
  }

  RegresarCuentaLoged():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<MostrarUsuario>(baseURL + 'DevolverCuenta', httpOptions);
  }

  GenerarReporte(llave:string):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'GenerarReporte', llave, httpOptions);
  }

}
