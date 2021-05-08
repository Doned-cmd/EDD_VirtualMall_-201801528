import { Component, OnInit } from '@angular/core';
import {MerkleService} from '../../services/merkle/merkle.service';

@Component({
  selector: 'app-reportes-merkle',
  templateUrl: './reportes-merkle.component.html',
  styleUrls: ['./reportes-merkle.component.css']
})
export class ReportesMerkleComponent implements OnInit {

  mostrarMensajeError:boolean = false;
  mensajeError:string;
  numero:number ;
  constructor(private ServicioDePedidos: MerkleService) { 
   
  }
  GenerarReporte(){
    this.ServicioDePedidos.GenerarReportes().subscribe((dataList:number[])=>{
      this.CambiarIntervalo()
    },(err)=>{      
    }
    )
  }

  CambiarIntervalo(){
    if(this.numero > 0){
      this.ServicioDePedidos.CambiarIntervalo(this.numero).subscribe((dataList:number[])=>{
            
      },(err)=>{      
      }
      )
    }
  }
  

  ngOnInit(): void {
  }

}
