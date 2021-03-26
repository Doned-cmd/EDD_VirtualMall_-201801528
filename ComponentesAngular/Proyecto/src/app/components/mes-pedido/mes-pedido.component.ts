import { Component, OnInit } from '@angular/core';
import {PedidoService} from '../../servicios/Pedidos/pedido.service';

@Component({
  selector: 'app-mes-pedido',
  templateUrl: './mes-pedido.component.html',
  styleUrls: ['./mes-pedido.component.css']
})
export class MesPedidoComponent implements OnInit {

  mostrarMensajeError:boolean = false;
  mensajeError:string;
  ListaMeses:number[]
  constructor(private ServicioDePedidos: PedidoService) { 
    this.ServicioDePedidos.DevolverMesesPedido().subscribe((dataList:number[])=>{
      this.ListaMeses=dataList      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }
  EstablecerMes(Mes:number){
    this.ServicioDePedidos.EstablecerMesBack(Mes).subscribe((dataList:any)=>{        
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }

  ngOnInit(): void {
  }

}
