import { Component, OnInit } from '@angular/core';
import { ProductosPedido } from "../../models/producto/producto";
import {PedidoService} from '../../servicios/Pedidos/pedido.service';

@Component({
  selector: 'app-dia-pedido',
  templateUrl: './dia-pedido.component.html',
  styleUrls: ['./dia-pedido.component.css']
})
export class DiaPedidoComponent implements OnInit {

  mostrarMensajeError:boolean = false;
  mensajeError:string;
  ListaDias:ProductosPedido[]
  constructor(private ServicioDePedidos: PedidoService) { 
    this.ServicioDePedidos.DevolverDias().subscribe((dataList:ProductosPedido[])=>{
      this.ListaDias=dataList      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }
  DirigirALista(Mes:ProductosPedido){
    this.ServicioDePedidos.EstablecerDia(Mes).subscribe((dataList:any)=>{        
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }

  ngOnInit(): void {
  }

}
