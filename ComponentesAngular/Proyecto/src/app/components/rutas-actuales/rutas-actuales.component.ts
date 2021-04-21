import { Component, OnInit } from '@angular/core';
import {CuentasService} from '../../services/Cuentas/cuentas.service';
@Component({
  selector: 'app-rutas-actuales',
  templateUrl: './rutas-actuales.component.html',
  styleUrls: ['./rutas-actuales.component.css']
})
export class RutasActualesComponent implements OnInit {

  
  mostrarMensajeError:boolean = false;
  mensajeError:string;
  Key:string
  constructor(private ServicioDePedidos: CuentasService) { 
    
  }

  

  ngOnInit(): void {
  }
}
