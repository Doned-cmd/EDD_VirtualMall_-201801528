import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from "rxjs";

import { Producto } from "../../models/producto/producto";
import { ProductoAngular } from "../../models/producto/producto";

@Injectable({
  providedIn: 'root'
})
export class ProductoService {

  constructor(private http: HttpClient) { }

  getProductosDeTienda():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<Producto[]>(baseURL + 'getListaProductos', httpOptions);
  }

  getProductosCarro():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<ProductoAngular[]>(baseURL + 'getListaCarrito', httpOptions);
  }

  ActualizarListaCarrito(product:Producto):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'ActualizarListaCarrito', product, httpOptions);
  }

  ActualizarCarroCambio(producto:ProductoAngular[]):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<number>(baseURL + 'ActualizarCarroCambio', producto, httpOptions);
  }

  ActualizarInventario():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'ActualizarInventario', httpOptions);
  }

  EliminarProductoCarro(producto:ProductoAngular):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<ProductoAngular[]>(baseURL + 'EliminarProductoCarro', producto,httpOptions);
  }

  ObtenerSumaCarro():Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.get<number>(baseURL + 'DevolversumaCarro', httpOptions);
  }

  CargarInventario(tiendaActual: String):Observable<any>{
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      }),
    };
  return this.http.post<any>(baseURL + 'DevolverListaProductPedidos', tiendaActual, httpOptions);
  }
}
