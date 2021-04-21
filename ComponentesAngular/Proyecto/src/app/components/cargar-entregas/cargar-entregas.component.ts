import { Component, EventEmitter, OnInit, Output, Input} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { PaqueteriaService } from '../../services/Paqueteria/paqueteria.service';
@Component({
  selector: 'app-cargar-entregas',
  templateUrl: './cargar-entregas.component.html',
  styleUrls: ['./cargar-entregas.component.css']
})
export class CargarEntregasComponent implements OnInit {

  
  @Input() opcion : string = ""
  @Output() onEnter : EventEmitter<string> = new EventEmitter();

  termino : string = ""
  error : boolean = false

  
  // Obtiene el contenido del archivo JSON
  onFileChange(event: any){
      let files = event.target.files;
    
      var lector = new FileReader();
      lector.readAsText(files[0]);
      lector.onload = () => {
      let texto:any = lector.result;
      if(texto){
        this.termino = texto;
      }
    }
  }
    //Envia el archivo al padre
  enviar(){
    this.tiendasService.CargarGrafos(this.termino).subscribe(
      resp =>{
        //console.log(resp.toString);
      }, err => {
        this.error = true
      }
    );
    //console.log(this.termino)
    //this.onEnter.emit(this.termino);
  }

  constructor(private tiendasService : PaqueteriaService) { }
  
  ngOnInit(): void {    
  }

}
