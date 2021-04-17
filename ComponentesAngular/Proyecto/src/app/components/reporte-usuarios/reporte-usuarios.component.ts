import { Component, OnInit } from '@angular/core';
import {CuentasService} from '../../services/Cuentas/cuentas.service';
@Component({
  selector: 'app-reporte-usuarios',
  templateUrl: './reporte-usuarios.component.html',
  styleUrls: ['./reporte-usuarios.component.css']
})
export class ReporteUsuariosComponent implements OnInit {

  mostrarMensajeError:boolean = false;
  mensajeError:string;
  Key:string
  constructor(private ServicioDePedidos: CuentasService) { 
    
  }
  GenerarReporte(){
    this.ServicioDePedidos.GenerarReporte(this.Key).subscribe((dataList:number[])=>{
          
    },(err)=>{      
    }
    )
  }
  

  ngOnInit(): void {
  }

}
