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

  constructor(private ServicioDeProductos: ProductoService) { 
    this.ServicioDeProductos.getProductosDeTienda().subscribe((dataList:Producto[])=>{
      this.listaProductos=dataList
      console.log(this.listaProductos)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )}

  ngOnInit(): void {
  }

  AgregarCarrito(Product:Producto){
    if(Product.Cantidad > 0){
      
      this.ServicioDeProductos.ActualizarListaCarrito(Product).subscribe((dataList:any)=>{        
        console.log(this.listaProductos)
      },(err)=>{
        this.mostrarMensajeError=true
        this.mensajeError = 'Error al cargar las tiendas'
      })     
    }else{
      this.mostrarMensajeError = true
      this.mensajeError = "Error, no hay suficiente existencias del producto"
    }
  }
  desactivarMensaje(){
    
    this.mostrarMensajeError=false
  }
}
