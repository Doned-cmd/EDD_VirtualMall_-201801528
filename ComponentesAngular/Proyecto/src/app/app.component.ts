import { Component, OnInit  } from '@angular/core';
import {CuentasService} from '../app/services/Cuentas/cuentas.service'

//import { Usuario } from "../../models/Usuario/usuario";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'Proyecto';
  
  MostrarAdmin:boolean;
  ngOnInit(): void {}
  constructor( private ServicioDePedidos: CuentasService)  {     
    this.ServicioDePedidos.VerificarSiAdmin().subscribe((dataList:boolean)=>{
      this.MostrarAdmin=dataList
    },(err)=>{
    }
    )}
}
