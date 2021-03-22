import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { baseURL } from "../../apiURL/baseURL";
import { Observable } from "rxjs";

import { Producto } from "../../models/producto/producto";
import { Tienda} from "../../models/tienda/tienda";

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
}
