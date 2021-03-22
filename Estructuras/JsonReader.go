package Estructuras

import(

)

//Conseguir datos Json
type Tienda struct {
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
	Logo string `json:"Logo"`
	Productos Avl
}
type Departamento struct {
	Nombre string
	Tiendas []Tienda
}

type Dato struct {
	Indice string
	Departamentos []Departamento
}

type Sobre struct {
	Datos []Dato
}
//Estructuras para Buscar coincidencias
type BusquedaEspecifica struct {
	Departamento string
	Nombre string
	Calificacion int
}

type EliminarTienda struct {
	Nombre string
	Categoria string
	Calificacion int
}
//Estructura para Agregar tiendas
type Invetarios struct {
	Tienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos []Producto `json:"Productos"`
}
type Producto struct {
	Nombre string `json:"Nombre"`
	Codigo int `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	Precio float32	`json:"Precio"`
	Cantidad int `json:"Cantidad"`
	Imagen string  `json:"Imagen"`

}
type SobreInventario struct {
	Invetarios []Invetarios
}

//Estructura para los pedidos
type SobrePedidos struct {
	Pedidos []Pedido
}

type Pedido struct {
	Fecha string `json:"Fecha"`
	Tienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos[] Producto  `json:"Productos"`
}


//Estructura para regresar productos al front end

type ProductoAngular struct {
	Tienda string
	Departamento string
	Nombre string
	Codigo int
	Descripcion string
	Precio int
	Cantidad int
	Imagen string
}

type TiendaAngular struct {
	Nombre string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
	Contacto string `json:"Contacto"`
	Calificacion int `json:"Calificacion"`
	Logo string `json:"Logo"`
	Departamento string
	IndiceNombre string
}