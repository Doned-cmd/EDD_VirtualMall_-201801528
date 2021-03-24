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
  Total: Number;

  constructor(private ServicioDeProductos: ProductoService) { 
    this.ServicioDeProductos.getProductosCarro().subscribe((dataList:ProductoAngular[])=>{
      this.listaCarro=dataList
      console.log(this.listaCarro)
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )}

  ngOnInit(): void {
  }

  checkValue(pro:Producto) {
    
    for (let index = 0; index < this.listaCarro.length; index++) {
      if(pro.Codigo==this.listaCarro[index].Codigo){
        if (this.listaCarro[index].Cantidad > this.listaCarro[index].CantidadMax){
          this.listaCarro[index].Cantidad=this.listaCarro[index].CantidadMax
        }else if(this.listaCarro[index].Cantidad < 1){
          this.listaCarro[index].Cantidad=0
        }
      }
    }
  }

  desactivarMensaje(){
    
    this.mostrarMensajeError=false
  }
  
}
