import { Component, OnInit } from '@angular/core';
import { ProductoPedido } from "../../models/producto/producto";
import {PedidoService} from '../../servicios/Pedidos/pedido.service';

@Component({
  selector: 'app-lista-productos',
  templateUrl: './lista-productos.component.html',
  styleUrls: ['./lista-productos.component.css']
})
export class ListaProductosComponent implements OnInit {

  mostrarMensajeError:boolean = false;
  mensajeError:string;
  ListaProductos:ProductoPedido[]
  constructor(private ServicioDePedidos: PedidoService) { 
    this.ServicioDePedidos.DevolverListaProductos().subscribe((dataList:ProductoPedido[])=>{
      this.ListaProductos=dataList      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }
  ngOnInit(): void {
  }


}
