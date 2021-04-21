import { Component, OnInit } from '@angular/core';
import {PaqueteriaService} from '../../services/Paqueteria/paqueteria.service';
@Component({
  selector: 'app-reportes-robot',
  templateUrl: './reportes-robot.component.html',
  styleUrls: ['./reportes-robot.component.css']
})
export class ReportesRobotComponent implements OnInit {

 
  
  mostrarMensajeError:boolean = false;
  mensajeError:string;
  Key:string
  listado:number[]
  constructor(private ServicioDePedidos: PaqueteriaService) { 
    this.ServicioDePedidos.MostrarMovimientosRobot().subscribe((dataList:number[])=>{
      this.listado = dataList      
    },(err)=>{
      //this.mostrarMensajeError=true
      //this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }
  Descargar(Descarga,num:number){
    Descarga.href = "assets/archivosd/CaminoCorto"+ num+".svg"
    Descarga.download = "CaminoCorto"+ num+".svg"
  }
  

  ngOnInit(): void {
  }

}
