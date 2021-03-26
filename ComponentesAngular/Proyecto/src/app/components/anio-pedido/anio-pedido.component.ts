import { Component, OnInit } from '@angular/core';
import {PedidoService} from '../../servicios/Pedidos/pedido.service';

@Component({
  selector: 'app-anio-pedido',
  templateUrl: './anio-pedido.component.html',
  styleUrls: ['./anio-pedido.component.css']
})
export class AnioPedidoComponent implements OnInit {

  mostrarMensajeError:boolean = false;
  mensajeError:string;
  ListaAnios:number[]

  constructor(private ServicioDePedidos: PedidoService) { 
    this.ServicioDePedidos.ObtenerAniosPedido().subscribe((dataList:number[])=>{
      this.ListaAnios=dataList      
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }

  DirigirAMes(Anio:number){
    this.ServicioDePedidos.ObtenerMesesPedido(Anio).subscribe((dataList:any)=>{           
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }

  GenerarReporteAnio(Anio:number){
    this.ServicioDePedidos.GenerarReporteMeses(Anio).subscribe((dataList:any)=>{     
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }
  GenerarReporteGlobal(){

  }

  ngOnInit(): void {
  }

}
