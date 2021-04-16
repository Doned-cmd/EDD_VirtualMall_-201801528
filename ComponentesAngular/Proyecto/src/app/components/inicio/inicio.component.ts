import { Component, OnInit } from '@angular/core';

import { UsuarioLogeado,MostrarUsuario } from "../../models/Usuario/usuario";

import {CuentasService} from '../../services/Cuentas/cuentas.service'

@Component({
  selector: 'app-inicio',
  templateUrl: './inicio.component.html',
  styleUrls: ['./inicio.component.css']
})
export class InicioComponent implements OnInit {
  Usero:MostrarUsuario
  mostrar:boolean = false
  //user:UsuarioLogeado
  constructor(private ServicioDePedidos: CuentasService)  {
    this.ServicioDePedidos.RegresarCuentaLoged().subscribe((dataList:MostrarUsuario)=>{
      this.mostrar = dataList.Buleano
      this.Usero = dataList
    },(err)=>{
      //this.mostrarMensajeError=true
      //this.mensajeError = 'Error al cargar las tiendas'
    }
    )

  }

  ngOnInit(): void {
  }

}
