import { Component, EventEmitter, OnInit, Output, Input} from '@angular/core';
import { ProductoService } from '../../services/producto/producto.service';

@Component({
  selector: 'app-cargar-inventario',
  templateUrl: './cargar-inventario.component.html',
  styleUrls: ['./cargar-inventario.component.css']
})
export class CargarInventarioComponent implements OnInit {

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
    this.tiendasService.CargarInventario(this.termino).subscribe(
      resp =>{
        //console.log(resp.toString);
      }, err => {
        this.error = true
      }
    );
    //console.log(this.termino)
    //this.onEnter.emit(this.termino);
  }

  constructor(private tiendasService : ProductoService) { }
  
  ngOnInit(): void {
  }
}
