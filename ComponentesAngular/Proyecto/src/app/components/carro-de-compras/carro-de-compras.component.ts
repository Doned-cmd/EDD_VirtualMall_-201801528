import { Component, OnInit } from '@angular/core';
import { Producto } from "../../models/producto/producto";
import { ProductoAngular } from "../../models/producto/producto";
import {ProductoService} from '../../services/producto/producto.service';
@Component({
  selector: 'app-carro-de-compras',
  templateUrl: './carro-de-compras.component.html',
  styleUrls: ['./carro-de-compras.component.css']
})
export class CarroDeComprasComponent implements OnInit {

  mostrarMensajeError:boolean = false;
  mensajeError:string;
  listaCarro: ProductoAngular[];
  Total: number;

  constructor(private ServicioDeProductos: ProductoService) { 
    this.ServicioDeProductos.getProductosCarro().subscribe((dataList:ProductoAngular[])=>{
      this.listaCarro=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
    this.ServicioDeProductos.ObtenerSumaCarro().subscribe((dataList:number)=>{
      this.Total=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error sumar los totales'
    }
    )
  }

  ngOnInit(): void {
  }

  checkValue(pro:ProductoAngular) {
    
    for (let index = 0; index < this.listaCarro.length; index++) {
      if(pro.Codigo==this.listaCarro[index].Codigo){
        if (this.listaCarro[index].Cantidad >= this.listaCarro[index].CantidadMax){
          this.listaCarro[index].Cantidad=this.listaCarro[index].CantidadMax
        }else if(this.listaCarro[index].Cantidad <= 1){
          this.listaCarro[index].Cantidad=1
        }
      }
    }

    this.ServicioDeProductos.ObtenerSumaCarro().subscribe((dataList:number)=>{
      this.Total=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error sumar los totales'
    }
    )
  }

  Agregar(pro:ProductoAngular, num:number) {
    
    for (let index = 0; index < this.listaCarro.length; index++) {      
      if(pro.Codigo==this.listaCarro[index].Codigo){
        this.listaCarro[index].Cantidad = this.listaCarro[index].Cantidad + num;
        if (this.listaCarro[index].Cantidad >= this.listaCarro[index].CantidadMax){
          this.listaCarro[index].Cantidad=this.listaCarro[index].CantidadMax
        }else if(this.listaCarro[index].Cantidad <= 1){
          this.listaCarro[index].Cantidad=1
        }
      }
    }
    this.ServicioDeProductos.ActualizarCarroCambio(this.listaCarro).subscribe((dataList:ProductoAngular[])=>{      
      this.listaCarro = dataList
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error sumar los totales'
    }
    )
    this.ServicioDeProductos.ObtenerSumaCarro().subscribe((dataList:number)=>{
      this.Total=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error sumar los totales'
    }
    )
  }


  EliminarDeCarrito(pro:ProductoAngular){
    this.ServicioDeProductos.EliminarProductoCarro(pro).subscribe((dataList:ProductoAngular[])=>{
      this.listaCarro=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al eliminar producto'
    }
    )
    this.ServicioDeProductos.ObtenerSumaCarro().subscribe((dataList:number)=>{
      this.Total=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error sumar los totales'
    }
    )
  }

  TeriminarPedido(){
    this.ServicioDeProductos.ActualizarInventario().subscribe((dataList:any)=>{      
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al eliminar producto'
    }
    )
    this.ServicioDeProductos.getProductosCarro().subscribe((dataList:ProductoAngular[])=>{
      this.listaCarro=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }

  desactivarMensaje(){
    
    this.mostrarMensajeError=false
  }
  
}
