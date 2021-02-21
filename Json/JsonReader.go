package Json

import(

)

//Conseguir datos Json
type Tienda struct {
	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
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
