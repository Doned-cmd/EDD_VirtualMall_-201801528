import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';


import { InicioComponent } from './components/inicio/inicio.component';
import {TiendasComponent}from './components/tiendas/tiendas.component'
import {ProductosComponent} from './components/productos/productos.component'
import {CarroDeComprasComponent} from './components/carro-de-compras/carro-de-compras.component'
import {AnioPedidoComponent} from './components/anio-pedido/anio-pedido.component'
import {MesPedidoComponent} from './components/mes-pedido/mes-pedido.component'
import {DiaPedidoComponent} from './components/dia-pedido/dia-pedido.component'
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
  },
  {
    path: 'pedidosAnio',
    component: AnioPedidoComponent,
  },
  {
    path: 'pedidosMes',
    component: MesPedidoComponent,
  },
  {
    path: 'pedidosDia',
    component: MesPedidoComponent,
  }    
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
