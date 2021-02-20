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

type BusquedaEspecifica struct {
	Departamento string
	Nombre string
	Calificacion int
}


