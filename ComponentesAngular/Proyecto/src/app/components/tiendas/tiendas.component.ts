import { Component, OnInit } from '@angular/core';
import {Tienda} from '../../models/tienda/tienda'
import {GettiendasService} from '../../services/tiendas/gettiendas.service'
@Component({
  selector: 'app-tiendas',
  templateUrl: './tiendas.component.html',
  styleUrls: ['./tiendas.component.css']
})
export class TiendasComponent implements OnInit {
  Tiendas: Tienda[];
  mostrarMensajeError:boolean = false;
  mensajeError:string;
  constructor(private gettiendaService: GettiendasService) { 
    this.gettiendaService.getListaTiendas().subscribe((dataList:Tienda[])=>{
      this.Tiendas=dataList
    },(err)=>{
      this.mostrarMensajeError=true
      this.mensajeError = 'Error al cargar las tiendas'
    }
    )}
  
    Comprar(tiendaactual: Tienda){
      console.log("Entro")
      this.gettiendaService.mandarTienda(tiendaactual).subscribe((res:any)=>{

      }, (err)=>{
        this.mostrarMensajeError=true
        this.mensajeError='No se pudo acceder a la tienda'
      })
      
    }

  ngOnInit(): void {
  }
  desactivarMensaje(){
    
    this.mostrarMensajeError=false
  }

}
