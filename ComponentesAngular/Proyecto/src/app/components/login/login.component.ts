import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';


import { Usuario } from "../../models/Usuario/usuario";
import {CuentasService} from '../../services/Cuentas/cuentas.service'
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  Dpi:number
  Password:string
  constructor(private router: Router, private ServicioDePedidos: CuentasService)  {}
  //this.router.navigate(['/'])
  ngOnInit(): void {
  }

    Verificar(){
    let usuario:Usuario = {
      Dpi: this.Dpi,
      Password: this.Password
    }
    

    this.ServicioDePedidos.CompararDatos(usuario).subscribe((dataList:boolean)=>{ 
      if (dataList){
        this.router.navigate(['/'])
      }else{ 
        alert('Datos invalidos, intente de nuevo')
      }
    },(err)=>{
      //this.mostrarMensajeError=true
      //this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }

}
