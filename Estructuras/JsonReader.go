package Estructuras

import(
	"strconv"
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

type Pedidos struct {
	Fecha string `json:"Fecha"`
	Tienda string `json:"Tienda"`
	Departamento string `json:"Departamento"`
	Calificacion int `json:"Calificacion"`
	Productos []Productos `json:"Productos"`
}
type Productos struct {
	Codigo int `json:"Codigo"`
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
	Pedidos []Pedidos
}


//Estructura para regresar productos al front end

type ProductoAngular struct {
	Tienda string
	Departamento string
	Calificacion int
	Nombre string
	Codigo int
	Descripcion string
	Precio float32
	Cantidad int
	CantidadMax int
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

type ProductosPedido struct {
	Dia int
	Categoria string
}



//EStructura para guardar los pedidos en el arbol

//Estructura para guardar usuarios

type Usuario struct {
	Dpi int `json:"Dpi"`
	Nombre string `json:"Nombre"`
	Correo string `json:"Correo"`
	Password string `json:"Password"`
	Cuenta string `json:"Cuenta"`
}

type MostrarUsuario struct {
	Buleano bool
	Usuario Usuario
}

func (u Usuario) toDOT() string {
	return strconv.Itoa(u.Dpi) + "\\n" + u.Nombre + "\\n" + u.Correo + "\\n" + u.Password + "\\n" + u.Cuenta
}
//Estructura para recibir usuarios

type UsuarioJson struct {
	Dpi int `json:"Dpi"`
	Password string `json:"Password"`
}
