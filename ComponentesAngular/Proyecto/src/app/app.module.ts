import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import { InicioComponent } from './components/inicio/inicio.component';
import { TiendasComponent } from './components/tiendas/tiendas.component';
import { ProductosComponent } from './components/productos/productos.component';
import { CarroDeComprasComponent } from './components/carro-de-compras/carro-de-compras.component';
import { AnioPedidoComponent } from './components/anio-pedido/anio-pedido.component';
import { MesPedidoComponent } from './components/mes-pedido/mes-pedido.component';
import { DiaPedidoComponent } from './components/dia-pedido/dia-pedido.component';
import { ListaProductosComponent } from './components/lista-productos/lista-productos.component';

@NgModule({
  declarations: [
    AppComponent,
    InicioComponent,
    TiendasComponent,
    ProductosComponent,
    CarroDeComprasComponent,
    AnioPedidoComponent,
    MesPedidoComponent,
    DiaPedidoComponent,
    ListaProductosComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
