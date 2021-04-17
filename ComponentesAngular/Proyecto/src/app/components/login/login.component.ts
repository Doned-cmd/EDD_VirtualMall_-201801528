import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';


import { Usuario,UsuarioLogeado,MostrarUsuario} from "../../models/Usuario/usuario";
import {CuentasService} from '../../services/Cuentas/cuentas.service'
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  Dpi:number
  Password:string

  DpiReg:number
  PasswordReg:string
  NombreReg:string
  CorreoReg:string
  CuentaReg:string = "Usuario"

  PasswordDel:string 

  Usero:MostrarUsuario
  mostrar:boolean = false
  constructor(private router: Router, private ServicioDePedidos: CuentasService)  {
    this.ServicioDePedidos.RegresarCuentaLoged().subscribe((dataList:MostrarUsuario)=>{
      this.mostrar = dataList.Buleano
      this.Usero = dataList
    },(err)=>{
      //this.mostrarMensajeError=true
      //this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }
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
  Registrar(){
    let usuariolog:UsuarioLogeado = {
      Dpi:this.DpiReg,
      Nombre:this.NombreReg,
      Correo: this.CorreoReg,
      Password:this.PasswordReg,
      Cuenta: this.CuentaReg
    }
    this.ServicioDePedidos.RegistrarUsuario(usuariolog).subscribe((dataList:boolean)=>{ 
      alert('Usuario Registrado Exitosamente1')
    },(err)=>{
      //this.mostrarMensajeError=true
      //this.mensajeError = 'Error al cargar las tiendas'
    }
    )
  }
  Eliminar(){
    let usuario:Usuario = {
      Dpi: 0,
      Password: this.PasswordDel
    }
    this.ServicioDePedidos.EliminarCuenta(usuario).subscribe((dataList:boolean)=>{ 
      if (dataList){
        alert('Usuario Eliminado Exitosamente')
        this.router.navigate(['/Login'])
      }else{
        alert('La contraseña no es correcta')
      }
      
    },(err)=>{
      //this.mostrarMensajeError=true
      //this.mensajeError = 'Error al cargar las tiendas'
    }
    )
    
    
  }


}
