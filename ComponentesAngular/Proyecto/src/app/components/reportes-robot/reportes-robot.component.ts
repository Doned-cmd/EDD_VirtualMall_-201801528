import { Component, OnInit } from '@angular/core';
import {CuentasService} from '../../services/Cuentas/cuentas.service';
@Component({
  selector: 'app-reportes-robot',
  templateUrl: './reportes-robot.component.html',
  styleUrls: ['./reportes-robot.component.css']
})
export class ReportesRobotComponent implements OnInit {

 
  
  mostrarMensajeError:boolean = false;
  mensajeError:string;
  Key:string
  constructor(private ServicioDePedidos: CuentasService) { 
    
  }

  

  ngOnInit(): void {
  }

}
