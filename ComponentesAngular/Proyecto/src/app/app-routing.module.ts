import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';


import { InicioComponent } from './components/inicio/inicio.component';
import {TiendasComponent}from './components/tiendas/tiendas.component'
import {ProductosComponent} from './components/productos/productos.component'
import {CarroDeComprasComponent} from './components/carro-de-compras/carro-de-compras.component'
/*import { CrearEstudianteComponent } from "./components/crear-estudiante/crear-estudiante.component";

import { AgregarCursosComponent } from "./components/agregar-cursos/agregar-cursos.component";*/

const routes: Routes = [
  {
    path: '',
    component: InicioComponent,
  },
  {
    path: 'vertiendas',
    component: TiendasComponent,
  },
  {
    path: 'productos',
    component: ProductosComponent,
  },
  {
    path: 'carro',
    component: CarroDeComprasComponent,
  }
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
