import { Component, OnInit } from '@angular/core';
import {ProductoService} from '../../services/producto/producto.service';
import { Producto } from "../../models/producto/producto";

@Component({
  selector: 'app-productos',
  templateUrl: './productos.component.html',
  styleUrls: ['./productos.component.css']
})
export class ProductosComponent implements OnInit {
  mostrarMensajeError:boolean = false;
  mensajeError:string;
  listaProductos: Producto[];

  constructor(private gettiendaService: ProductoService) { 
    this.gettiendaService.getProductosDeTienda().subscribe((dataList:Producto[])=>{
      this.listaProductos=dataList
      console.log(this.listaProductos)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )}

  ngOnInit(): void {
  }

  desactivarMensaje(){
    
    this.mostrarMensajeError=false
  }
}
