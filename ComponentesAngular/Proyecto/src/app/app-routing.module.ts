import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';


import { InicioComponent } from './components/inicio/inicio.component';
import {TiendasComponent}from './components/tiendas/tiendas.component'
import {ProductosComponent} from './components/productos/productos.component'
import {CarroDeComprasComponent} from './components/carro-de-compras/carro-de-compras.component'
import {AnioPedidoComponent} from './components/anio-pedido/anio-pedido.component'
import {MesPedidoComponent} from './components/mes-pedido/mes-pedido.component'
import {DiaPedidoComponent} from './components/dia-pedido/dia-pedido.component'
import {ListaProductosComponent} from './components/lista-productos/lista-productos.component'
import {CargarInventarioComponent} from './components/cargar-inventario/cargar-inventario.component'
import {CargarUsuariosComponent} from './components/cargar-usuarios/cargar-usuarios.component'
import {CargarPedidosComponent} from './components/cargar-pedidos/cargar-pedidos.component'
import {CargarTiendasComponent} from './components/cargar-tiendas/cargar-tiendas.component'
import {LoginComponent} from './components/login/login.component'
import {ReporteUsuariosComponent} from './components/reporte-usuarios/reporte-usuarios.component'
import { CargarEntregasComponent } from './components/cargar-entregas/cargar-entregas.component'
import { RutasActualesComponent } from './components/rutas-actuales/rutas-actuales.component'
import { ReportesRobotComponent } from './components/reportes-robot/reportes-robot.component'
/*import { CrearEstudianteComponent } from "./components/crear-estudiante/crear-estudiante.component";

import { AgregarCursosComponent } from "./components/agregar-cursos/agregar-cursos.component";*/

const routes: Routes = [
  {
    path: '',
    component: InicioComponent,
  },
  {
    path: 'CargarInventario',
    component: CargarInventarioComponent,
  },
  {
    path: 'CargarUsuarios',
    component: CargarUsuariosComponent,
  },
  {
    path: 'ReportarUsuarios',
    component: ReporteUsuariosComponent,
  },
  {
    path: 'Login',
    component: LoginComponent,
  },
  {
    path: 'CargarPedidos',
    component: CargarPedidosComponent,
  },
  {
    path: 'CargarTiendas',
    component: CargarTiendasComponent,
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
    component: DiaPedidoComponent,
  },
  {
    path: 'ListaPedido',
    component: ListaProductosComponent,
  },
  {
    path: 'CargarRutas',
    component: CargarEntregasComponent,
  },
  {
    path: 'RutasActuales',
    component: RutasActualesComponent,
  },
  {
    path: 'ReportesRobot',
    component: ReportesRobotComponent,
  }    
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
