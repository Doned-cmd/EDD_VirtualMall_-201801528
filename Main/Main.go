package main

import (
	"../Estructuras"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"os"
	"os/exec"
	"strconv"
)


//Matriz 3D
type Indice struct {
	Nombre string
	Departamentos []DepartamentoMatriz
}
type DepartamentoMatriz struct {
	Nombre string
	Calificaciones [5]Calificacion
}
type Calificacion struct {
	Calificacion int
	Lista *Estructuras.ListaDoble
}


//Arreglo Linealizado

type ArregloLinea struct {
	IndiceNombre string
	Departamento string
	Calificacion int

	Lista *Estructuras.ListaDoble
}

//Arreglos de datos
var matricita []Indice
var ArregloCali []ArregloLinea


var ArbolUsuarios *Estructuras.BTree

func main() {
	ArbolUsuarios = Estructuras.NewBTree(5)
	UsuarioAdmin := Estructuras.Usuario{
		Dpi:      1234567890101,
		Nombre:   "EDD2021",
		Correo:   "auxiliar@edd.com",
		Password: "12343",
		Cuenta:   "Admin",
	}


	ArbolUsuarios.Insert(UsuarioAdmin)
	//ArbolUsuarios.Gragicar()
	//EncriptacionFernet.EncriptarString("hola",Estructuras.LLaveEncriptado)
	//EncriptacionFernet.Desencriptar(EncriptacionFernet.EncriptarString("hola",Estructuras.LLaveEncriptado))

	//ArbolUsuarios.Insert(Usuario)
	//ArbolUsuarios.Insert(Usuario2)
	//ArbolUsuarios.Insert(Usuario3)
	//ArbolUsuarios.Insert(Usuario4)
	//ArbolUsuarios.Insert(Usuario5)
	//ArbolUsuarios.Insert(Usuario6)
	//ArbolUsuarios.Insert(Usuario7)
	//ArbolUsuarios.Insert(Usuario8)
	//ArbolUsuarios.Insert(Usuario9)
	//ArbolUsuarios.Insert(Usuario10)


	//ArbolUsuarios.Insert(Nuevo4)
	//ArbolUsuarios.ImprimirArbol()
	//fmt.Println(ArbolUsuarios.Gragicar())
	//ArbolUsuarios.Gragicar()
	//numeros := []int{5}
	//numeros = append(numeros, 7)
    //numeros = append(numeros[:0], numeros[0+1:]...) //cola
	//numeros = append(numeros[:len(numeros)-1], numeros[len(numeros)-1+1:]...) //pila


	//fmt.Print(numeros)
	//numeros = append(numeros, 1)
	//fmt.Print()
	request()


	//print(Regresar(3))
	//Regresar(&(numeros))
	Estructuras.NumeroGrafo = 0
	ContarMovRobot = -1
}

func Regresar(numeros *[]int)  {
	*numeros = append(*numeros, 5)
}

func request(){
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/getArreglo",getArreglo).Methods("GET")
	myrouter.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	myrouter.HandleFunc("/TiendaEspecifica", TiendaEspecifica).Methods("POST")
	myrouter.HandleFunc("/Eliminar", Eliminar).Methods("DELETE")
	myrouter.HandleFunc("/Guardar", Guardar).Methods("GET")
	myrouter.HandleFunc("/id/{indice}", ShowByID).Methods("GET")
	myrouter.HandleFunc("/cargarInventario", CargarInventario).Methods("POST")
	myrouter.HandleFunc("/getListaTiendas", DevolverlistaTiendas).Methods("GET")
	myrouter.HandleFunc("/getTiendaActual", TiendaActual).Methods("POST")
	myrouter.HandleFunc("/getListaProductos", DevolverlistaProductos).Methods("GET")

	myrouter.HandleFunc("/cargarPedido", CargarPedido).Methods("POST")
	myrouter.HandleFunc("/getListaCarrito", DevolverlistaCarrito).Methods("GET")
	myrouter.HandleFunc("/ActualizarInventario",ActualizarInventario).Methods("POST")
	myrouter.HandleFunc("/ActualizarListaCarrito",ActualizarListaCarro).Methods("POST")
	myrouter.HandleFunc("/EliminarProductoCarro",EliminarProductoCarro).Methods("POST")
	myrouter.HandleFunc("/DevolversumaCarro",DevolversumaCarro).Methods("GET")
	myrouter.HandleFunc("/ActualizarCarroCambio",ActualizarCarroCambio).Methods("POST")

	myrouter.HandleFunc("/ObtenerAniosPedido",ObtenerAniosPedido).Methods("GET")
	myrouter.HandleFunc("/ObtenerAnioPedido",ObtenerMesPedido).Methods("POST")
	myrouter.HandleFunc("/DevolverMesesPedido",DevolverMesesPedido).Methods("GET")
	myrouter.HandleFunc("/EstablecerMesBack",EstablecerMesBack).Methods("POST")
	myrouter.HandleFunc("/DevolverDias",DevolverDias).Methods("GET")

	myrouter.HandleFunc("/EstablecerDia",EstablecerDia).Methods("POST")
	myrouter.HandleFunc("/DevolverListaProductPedidos",DevolverListaProductPedidos).Methods("GET")


	myrouter.HandleFunc("/CargarUsuarios",CargarUsuarios).Methods("POST")
	myrouter.HandleFunc("/RegistrarUsuario",RegistrarUsuario).Methods("POST")

	myrouter.HandleFunc("/VerificarCuenta",VerificarCuenta).Methods("POST")
	myrouter.HandleFunc("/VerificarSiAdmin",VerificarSiAdmin).Methods("POST")
	myrouter.HandleFunc("/DevolverCuenta",DevolverCuenta).Methods("GET")
	myrouter.HandleFunc("/EliminarUsuario",EliminarUsuario).Methods("POST")
	myrouter.HandleFunc("/GenerarReporte",GenerarReporte).Methods("POST")

	myrouter.HandleFunc("/CargarEntregas",CargarEntregas).Methods("POST")
	myrouter.HandleFunc("/MostrarMovimientosRobot",MostrarMovimientosRobot).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myrouter)))
}




func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Servidor en Go")
}



//REGRESAR GRAFICO DE COMO ESTA LA MEMORIA
func getArreglo(w http.ResponseWriter, r *http.Request)  {
	var path = "D:/Escritorio/USAC/EDD/Proyecto/Practica 1/Grafico.dot"
	texto:="digraph G{" + "\n"

	contador := 0
	for _,cali := range ArregloCali{

		if cali.Lista.IsEmpty(){

			texto = texto + "node" + strconv.Itoa(contador) +`[label="` + "nulo"+`"` +"]" + "\n"
			contador++
		}else{
			nodos := cali.Lista.ReturnNodes()

			for j, recorriendo := range nodos{

				texto = texto + "node" + strconv.Itoa(contador) +`[label="` + recorriendo.Nombre + `"` +"]" + "\n"
				if j != len(nodos)-1{
					texto = texto + "node" + strconv.Itoa(contador)+ " -> " + "node" + strconv.Itoa(contador+1) + "\n"
				}
				contador++
			}
		}


		texto = texto + "\n"
	}

	texto = texto + "}"

	almacenarEnArchivoDot(path, texto)
}

func almacenarEnArchivoDot(path string, text string){
	var _, erro = os.Stat(path)

	//Crea el archivo si no existe
	if os.IsNotExist(erro) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()

	// Escribe el texto en el archivo
	_, err = file.WriteString(text)
	if existeError(err) {
		return
	}

	// Guarda los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}

	//Creación de la imagen
	path1, _ := exec.LookPath("dot")

	cmd, _ := exec.Command(path1, "-Tpng", "Grafico.dot").Output() //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777) //Se mantiene igual
	ioutil.WriteFile("List.png", cmd, os.FileMode(mode)) //Creacion de la imagen
}

//CARGAR JSON CON TIENDAS-----------------------------------------------------------------------------------------
func cargartienda(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Estructuras.Sobre
	json.Unmarshal(body, &re)
	fmt.Println(re.Datos)
	ConvertToMatrix(re)
	ConvertToArray(matricita)
}
//convertir a matriz
func  ConvertToMatrix (sbr Estructuras.Sobre) {

	var Matriz []Indice
	for  ind,r := range sbr.Datos {

		Indice := Indice{
			Nombre:        r.Indice,
			Departamentos: nil,
		}
		Matriz = append(Matriz, Indice)
		var Departamentos []DepartamentoMatriz

		for  depa,s := range r.Departamentos {
			Departamento := DepartamentoMatriz{
				Nombre:         s.Nombre,
			}

			Departamentos = append(Departamentos, Departamento)
			var  calificaciones  [5]Calificacion
			for i :=0; i<5; i++ {
				calificaciones[i]= Calificacion{
					Calificacion: i,
					Lista:        Estructuras.NewListaDoble(),
				}
			}

			for  _,t := range s.Tiendas {
				if t.Calificacion == 1 {
					calificaciones[0].Lista.Add(t)
					//calificaciones[0].Lista.ImprimirLista()
				}else if t.Calificacion ==2{
					calificaciones[1].Lista.Add(t)
					//calificaciones[1].Lista.ImprimirLista()
				}else if t.Calificacion ==3{
					calificaciones[2].Lista.Add(t)
					//calificaciones[2].Lista.ImprimirLista()
				}else if t.Calificacion ==4{
					calificaciones[3].Lista.Add(t)
					//print(t.Calificacion, t.Nombre)
					//calificaciones[3].Lista.ImprimirLista()
				}else if t.Calificacion ==5{
					calificaciones[4].Lista.Add(t)
					//print(t.Calificacion, t.Nombre)
					//calificaciones[4].Lista.ImprimirLista()
				}
			}
			Departamentos[depa].Calificaciones = calificaciones
		}
		Matriz[ind].Departamentos = Departamentos
	}
	matricita = Matriz
	println()
	//println("Matricitia")
	//for _,indi := range matricita{
	//	for _,depa := range indi.Departamentos{
	//		for _,cal := range depa.Calificaciones{
	//			cal.Lista.ImprimirLista()
	//		}
	//	}
	//}
}

func ConvertToArray(Matriz []Indice){
	arreglo := IniciarArreglo(Matriz)
	println()
	println("Arreglo")
	for  i,r := range Matriz{
		for  j,s := range r.Departamentos {
			for  k,cali := range s.Calificaciones {
				if cali.Lista.IsEmpty() {
					arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].IndiceNombre = r.Nombre
					arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].Departamento = s.Nombre
					arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].Calificacion = cali.Calificacion
					arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].Lista = cali.Lista
					//arreglo[i + (len(Matriz)-1) * (j + (len(r.Departamentos)-1) * k) ].Lista.ImprimirLista()
				}else{
					for _,tiendaenarreglo := range ArregloBurbuja(cali.Lista.ReturnNodes()) {
						arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].IndiceNombre = r.Nombre
						arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].Departamento = s.Nombre
						arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].Calificacion = cali.Calificacion
						arreglo[i+(len(Matriz))*(j+(len(r.Departamentos))*k)].Lista.Add(tiendaenarreglo)
					}
				}
			}
		}
	}

	ArregloCali = arreglo

	for num,i:= range ArregloCali{
		println("++++++++++++++++++++++++++++++++++++++++++++")
		print( num, i.IndiceNombre, " ",i.Departamento, " ",i.Calificacion, " ")
		i.Lista.ImprimirLista()
		println("++++++++++++++++++++++++++++++++++++++++++++")
	}
}

//Inicia el arreglo para poder ingresar las listas con la formula de column major
func IniciarArreglo(Matriz []Indice)  []ArregloLinea{
	var arreglo[]ArregloLinea
	for  _,r := range Matriz{
		for  _,s := range r.Departamentos {
			for  _,t := range s.Calificaciones {
				var nuevodatoarreglo = ArregloLinea{
					IndiceNombre: r.Nombre,
					Departamento: s.Nombre,
					Calificacion: t.Calificacion,
					Lista:        Estructuras.NewListaDoble(),
				}
				arreglo = append(arreglo, nuevodatoarreglo)
			}
		}
	}
	return arreglo
}

func ArregloBurbuja(ListaAOrdenar []Estructuras.Tienda) []Estructuras.Tienda{
	var auxiliar Estructuras.Tienda
	for i := 0; i < len(ListaAOrdenar); i++ {
		for j := 0; j < len(ListaAOrdenar)-1; j++ {
			if ListaAOrdenar[j].Nombre > ListaAOrdenar[j+1].Nombre {
				auxiliar = ListaAOrdenar[j]
				ListaAOrdenar[j] = ListaAOrdenar[j+1]
				ListaAOrdenar[j+1] = auxiliar
			}
		}
	}
	return ListaAOrdenar
}



//BUSCAR TIENDA ESPECIFICA-----------------------------------------------------------------------------------------
func TiendaEspecifica(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Estructuras.BusquedaEspecifica
	json.Unmarshal(body, &re)
	println("Buscando coincidencias con los siguientes parametros: ","Nombre: " ,re.Nombre, "Departamento: ",re.Departamento,"Calificacion: " ,re.Calificacion)
//EJECUTAR METODO PARA ENCONTRAR LA TIENDA
	var TiendaEncontrada Estructuras.Tienda
	imprimir := true
	for _, cali := range ArregloCali{
		fmt.Println(cali.Calificacion )
			if (cali.Calificacion+1 == re.Calificacion) && (cali.Departamento == re.Departamento) {

				if cali.Lista.Search(re.Nombre) != -1 {
					TiendaEncontrada = cali.Lista.SearchNReturn(re.Nombre)
					fmt.Fprint(w,"Se encontro coincidencia ", TiendaEncontrada.Calificacion, TiendaEncontrada.Nombre, TiendaEncontrada.Descripcion, TiendaEncontrada.Contacto)
					println("Se encontro coincidencia ", TiendaEncontrada.Calificacion, TiendaEncontrada.Nombre, TiendaEncontrada.Descripcion, TiendaEncontrada.Contacto)
					imprimir = false
					GenerarJsonCoincidencia(TiendaEncontrada)
					json.NewEncoder(w).Encode(TiendaEncontrada)
				}
			}
	}
	if imprimir {
		println("No se encontro ninguna coincidencia")
		fmt.Fprint(w, "No se encontro ninguna coincidencia")
	}
}


func GenerarJsonCoincidencia(Tienda Estructuras.Tienda) {

	var path = "D:/Escritorio/USAC/EDD/Practica 1/Coincidencia.Json"

	crear_json, _ := json.Marshal(Tienda)
	almacenarEnArchivo(path, crear_json)
}

func almacenarEnArchivo(path string, text []byte){
	var _, erro = os.Stat(path)

	//Crea el archivo si no existe
	if os.IsNotExist(erro) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()

	// Escribe algo de texto linea por linea
	_, err = file.WriteString(string(text))
	if existeError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
}


func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}


//ELIMINAR UNA TIENDA--------------------------------------------------------------------------------------------------
func Eliminar(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Estructuras.EliminarTienda
	json.Unmarshal(body, &re)
	println("Buscando coincidencias con los siguientes parametros: ","Nombre: " ,re.Nombre, "Categoria: ",re.Categoria,"Calificacion: " ,re.Calificacion)
	//EJECUTAR METODO PARA ELIMINAR UNA TIENDA

	imprimir := true
	for _, cali := range ArregloCali{
		if (cali.Calificacion == re.Calificacion) && (cali.Departamento == re.Categoria) {
			if cali.Lista.Search(re.Nombre) != -1 {
				println("Se encontro coincidencia ", re.Calificacion, re.Nombre, re.Categoria)
				cali.Lista.DeleteByName(re.Nombre)
				imprimir = false

			}
		}
	}
	if imprimir {
		println("No se encontro ninguna coincidencia")
		fmt.Fprint(w, "No se encontro ninguna coincidencia")
	}
}



//GUARDAR EN JSON
func Guardar(w http.ResponseWriter, r *http.Request) {
	var path = "D:/Escritorio/USAC/EDD/Practica 1/Copia.Json"

	Matriz := matricita
	var sbr Estructuras.Sobre


	for  i:=0; i<len(Matriz);i++ {
		dato := Estructuras.Dato{
			Indice:      Matriz[i].Nombre,
		}
		sbr.Datos = append(sbr.Datos, dato)


		for j :=0; j <len(Matriz[i].Departamentos); j++ {
			depa := Estructuras.Departamento{
				Nombre:   Matriz[i].Departamentos[j].Nombre,

			}
			sbr.Datos[i].Departamentos = append(sbr.Datos[i].Departamentos, depa)
			//var tiendas []Json.Tienda
			for k :=0; k <len(Matriz[i].Departamentos[j].Calificaciones); k++ {
				Matriz[i].Departamentos[j].Calificaciones[k].Lista = ArregloCali[i+(len(Matriz))*(j+(len(Matriz[i].Departamentos))*k)].Lista

				if Matriz[i].Departamentos[j].Calificaciones[k].Lista.IsEmpty() {

				}else{
					for _, store :=  range Matriz[i].Departamentos[j].Calificaciones[k].Lista.ReturnNodes() {
						sbr.Datos[i].Departamentos[j].Tiendas = append(sbr.Datos[i].Departamentos[j].Tiendas, store)
						//println(tiendas[l].Nombre)
					}
				}
			}
		}
	}





	crear_json, _ := json.Marshal(sbr)
	almacenarEnArchivo(path, crear_json)
}

//MOSTRAR RESULTADO POR INDICE
func ShowByID(w http.ResponseWriter, r *http.Request){
	parametro:= mux.Vars(r)
	i:= parametro["indice"]
	println(i)
	for x,correr := range ArregloCali {
		if  strconv.Itoa(x+1) == i{
			path := "D:/Escritorio/USAC/EDD/Practica 1/CoincidenciaPorIndice.Json"
			tienditas := correr.Lista.ReturnNodes()
			crear_json, _ := json.Marshal(tienditas)
			json.NewEncoder(w).Encode(tienditas)
			almacenarEnArchivo( path,crear_json)
		}
	}

}
//Cargar Inventarios --------------------------------------------------------------------------------------------------------------------

func CargarInventario(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Estructuras.SobreInventario
	json.Unmarshal(body, &re)
	//fmt.Println(body)
	fmt.Println(re)
	//fmt.Println(re.Invetarios)
	agregarAArbol(re)

}
func agregarAArbol(re Estructuras.SobreInventario){
	for _,inventario := range re.Invetarios{
		//fmt.Println(inventario)
		if buscarTienda(inventario){//Se busca si la tienda existe
			//TiendaA := RegresarTienda(inventario)
			for _,producto := range inventario.Productos {
				fmt.Println(producto.Nombre, producto.Descripcion, producto.Cantidad)
				if buscarProducto( inventario,producto){
					ActualizarInventarioSumando(inventario,producto)
				}else{
					agregarAArreglo(inventario,producto)
				}
			}
		}
	}
}
func ActualizarInventarioSumando(inventario Estructuras.Invetarios ,producto Estructuras.Producto){
	fmt.Println("sumando al producto repertido ",producto.Nombre," La cantidad: " ,producto.Cantidad)
	for   _, cali := range ArregloCali{
		if (cali.Calificacion+1 == inventario.Calificacion) && (cali.Departamento == inventario.Departamento) {
			if cali.Lista.Search(inventario.Tienda) != -1 {
				cali.Lista.SearchIndex(cali.Lista.Search(inventario.Tienda)).GetTienda().Productos.SumarInventario(producto)
			}
		}
	}
}

func agregarAArreglo(inventario Estructuras.Invetarios ,producto Estructuras.Producto){
	for   _, cali := range ArregloCali{
		if (cali.Calificacion+1 == inventario.Calificacion) && (cali.Departamento == inventario.Departamento) {
			//fmt.Println(producto)
			if cali.Lista.Search(inventario.Tienda) != -1 {
				cali.Lista.SearchIndex(cali.Lista.Search(inventario.Tienda)).GetTienda().Productos.Insertar(producto)
			}
		}
	}
}

func buscarProducto(inventario Estructuras.Invetarios ,producto Estructuras.Producto) bool{
	for   _, cali := range ArregloCali{
		if (cali.Calificacion+1 == inventario.Calificacion) && (cali.Departamento == inventario.Departamento) {
			if cali.Lista.Search(inventario.Tienda) != -1 {
				if cali.Lista.SearchIndex(cali.Lista.Search(inventario.Tienda)).GetTienda().Productos.Raiz == nil{
					return false
				}else if cali.Lista.SearchIndex(cali.Lista.Search(inventario.Tienda)).GetTienda().Productos.SearchProductExists(producto){
					return true
				} else {
					return false
				}
			}
		}
	}

	return false
}
func buscarTienda(inventario Estructuras.Invetarios) bool{
	//EJECUTAR METODO PARA ENCONTRAR LA TIENDA
	//var TiendaEncontrada Estructuras.Tienda
		//fmt.Println("Se busca la coincidencia: ", inventario.Tienda)
	for _, cali := range ArregloCali{
		if (cali.Calificacion+1 == inventario.Calificacion) && (cali.Departamento == inventario.Departamento) {
			//fmt.Println(inventario.Tienda)
			//fmt.Println(cali.Lista.Search(inventario.Tienda))
			if cali.Lista.Search(inventario.Tienda) != -1 {
				return true
			}
		}
	}
	return false
}
func RegresarTienda(inventario Estructuras.Invetarios) *Estructuras.Tienda{
	for _, cali := range ArregloCali{
		if (cali.Calificacion+1 == inventario.Calificacion) && (cali.Departamento == inventario.Departamento) {
			//fmt.Println(inventario.Tienda)
			//fmt.Println(cali.Lista.Search(inventario.Tienda))
			if cali.Lista.Search(inventario.Tienda) != -1 {
				return cali.Lista.SearchIndex(cali.Lista.Search(inventario.Tienda)).GetTienda()
			}
		}
	}
	return nil
}

//Cargar Pedidos --------------------------------------------------------------------------------------------------------------------

var ArbolPedidos Estructuras.AvlPedidos
func CargarPedido(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Estructuras.SobrePedidos
	json.Unmarshal(body, &re)
	//fmt.Println(re.Pedidos)
	Agregarpedidos(re)
}

func Agregarpedidos(res Estructuras.SobrePedidos){
  for _,Recorrer:= range res.Pedidos {
		fecha := strings.Split(Recorrer.Fecha, "-")
		anio,_ := strconv.Atoi(fecha[2])
		mes,_ := strconv.Atoi(fecha[1])
		dia,_ := strconv.Atoi(fecha[0])
		//fmt.Println("fecha: ", dia,mes,anio)
		inven:= Estructuras.Invetarios{
			Tienda:       Recorrer.Tienda,
			Departamento: Recorrer.Departamento,
			Calificacion: Recorrer.Calificacion,
		}
		//fmt.Println(inven)
		if buscarTienda(inven){
			if !ArbolPedidos.SearchAnioExists(anio){
				ArbolPedidos.Insertar(anio)
			}
			for _,Rec:= range Recorrer.Productos {
				pro := Estructuras.Producto{
					Codigo: Rec.Codigo,
				}
				if buscarProducto(inven,pro){
					if ArbolPedidos.ReturnAnioNode(anio).Mes.IsEmpty(){
						ArbolPedidos.ReturnAnioNode(anio).Mes.Add(mes)
						fmt.Println(mes)
					}else{
						if	ArbolPedidos.ReturnAnioNode(anio).Mes.Search(mes) != -1{
							ArbolPedidos.ReturnAnioNode(anio).Mes.Add(mes)
							fmt.Println(mes)
						}
						ArbolPedidos.ReturnAnioNode(anio).Mes.SearchIndex(ArbolPedidos.ReturnAnioNode(anio).Mes.Search(mes)).GetMatriz().Insert(dia,inven.Departamento)
						ArbolPedidos.ReturnAnioNode(anio).Mes.SearchIndex(ArbolPedidos.ReturnAnioNode(anio).Mes.Search(mes)).GetMatriz().SearchNreturn(dia, inven.Departamento).Productos.Add(Rec)
						//ArbolPedidos.ReturnAnioNode(anio).Mes.SearchNReturnM(mes).GetMatriz().Insert(dia,inven.Departamento)
						//ArbolPedidos.ReturnAnioNode(anio).Mes.SearchNReturnM(mes).GetMatriz().SearchNreturn(dia, inven.Departamento).Productos.Add(Rec)
					}
				}
			}

		}
	}
}

func DevolverDotPedidos(w http.ResponseWriter, r *http.Request){

}

//Devolver la lista de las tiendas al front end

func DevolverlistaTiendas(w http.ResponseWriter, r *http.Request){
	var tiendas []Estructuras.TiendaAngular
	for _, depas := range ArregloCali {
		if !depas.Lista.IsEmpty(){
			for _,tienda := range depas.Lista.ReturnNodes(){

				tiendas = append(tiendas, Estructuras.TiendaAngular{
					Nombre:       tienda.Nombre,
					Descripcion:  tienda.Descripcion,
					Contacto:     tienda.Contacto,
					Calificacion: tienda.Calificacion,
					Logo:         tienda.Logo,
					Departamento: depas.Departamento,
					IndiceNombre: depas.IndiceNombre,
				})
			}
		}
	}
	json.NewEncoder(w).Encode(tiendas)
}


func DevolverlistaProductos(w http.ResponseWriter, r *http.Request){
	for _, depas := range ArregloCali {
		if depas.Calificacion+1 == tiendaseleccionada.Calificacion && depas.Departamento == tiendaseleccionada.Departamento{

			if depas.Lista.Search(tiendaseleccionada.Nombre) != -1{
				fmt.Println("Devolviendo productos de  ", tiendaseleccionada.Nombre)
				fmt.Println(*depas.Lista.SearchNReturnM(tiendaseleccionada.Nombre).Productos.ReturnProductsOfTree())
				json.NewEncoder(w).Encode(*depas.Lista.SearchNReturnM(tiendaseleccionada.Nombre).Productos.ReturnProductsOfTree())
			}
		}
	}
}



//Establecer la tienda para buscar los productos
var tiendaseleccionada Estructuras.TiendaAngular // Variable para almacenar un bufer de la tienda en la que se esta actualmente en el front end

func TiendaActual(w http.ResponseWriter, r *http.Request)  {
	var tiendasec Estructuras.TiendaAngular
	tiendaseleccionada = tiendasec
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &tiendaseleccionada)
	//print("se entrro")
	fmt.Println(tiendaseleccionada)
}

// -------------------Carrito de compras
var ListaCarro []Estructuras.ProductoAngular // Variable para almacenar el carrito en el servidor

func ActualizarListaCarro(w http.ResponseWriter, r *http.Request){
	var ProductoAgregaro Estructuras.Producto
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println(w, "Error")
	}
	json.Unmarshal(reqBody,&ProductoAgregaro)
	nuevoProduct:=Estructuras.ProductoAngular{
		Tienda:       tiendaseleccionada.Nombre,
		Departamento: tiendaseleccionada.Departamento,
		Calificacion: tiendaseleccionada.Calificacion,
		Nombre:       ProductoAgregaro.Nombre,
		Codigo:       ProductoAgregaro.Codigo,
		Descripcion:  ProductoAgregaro.Descripcion,
		Precio:       ProductoAgregaro.Precio,
		Cantidad:     1,
		CantidadMax:  ProductoAgregaro.Cantidad,
		Imagen:       ProductoAgregaro.Imagen,
		Almacenamiento: ProductoAgregaro.Almacenamiento,
	}

	if(noDuplicarCarro(nuevoProduct)){
		fmt.Println("Se agrego sumo 1 al producto: " ,nuevoProduct)
	}else{
		ListaCarro = append(ListaCarro, nuevoProduct)
		fmt.Println("Se agrego al carrito: " ,nuevoProduct)
	}
}
func noDuplicarCarro(nuevoProduct Estructuras.ProductoAngular) bool{
	for i,producto:= range ListaCarro{
		if producto.Codigo == nuevoProduct.Codigo{
			if NoExcederCantidad(ListaCarro[i]){
			ListaCarro[i].Cantidad ++
			}
			return true
		}
	}
	return false
}
func NoExcederCantidad(producto Estructuras.ProductoAngular) bool{
	for _, depas := range ArregloCali {
		if depas.Calificacion+1 == producto.Calificacion && depas.Departamento == producto.Departamento {
			if depas.Lista.Search(producto.Tienda) != -1 {
				if depas.Lista.SearchIndex(depas.Lista.Search(producto.Tienda)).GetTienda().Productos.VerificarDisponibilidad(producto){
					return true
				}
			}
		}
	}
	return false
}


func DevolverlistaCarrito (w http.ResponseWriter, r *http.Request) {
	//fmt.Println(ListaCarro)
	//var ListaProductos []Estructuras.Producto
	//for _,productocarro := range ListaCarro{
		//NuevoPro := Estructuras.Producto{
		//	Nombre:      productocarro.Nombre,
		//	Codigo:      productocarro.Codigo,
		//	Descripcion: productocarro.Descripcion,
		//	Precio:      productocarro.Precio,
		//	Cantidad:    productocarro.Cantidad,
		//	Imagen:      productocarro.Imagen,
//		}
//		ListaProductos = append(ListaProductos, NuevoPro)
	//}
	json.NewEncoder(w).Encode(ListaCarro)
}

func EliminarProductoCarro(w http.ResponseWriter, r *http.Request){
	var tiendasec Estructuras.ProductoAngular
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &tiendasec)
 	
	for x,rec:= range ListaCarro {
		if rec.Departamento == tiendasec.Departamento && rec.Calificacion == tiendasec.Calificacion && rec.Nombre == tiendasec.Nombre{
			ListaCarro = append(ListaCarro[:x], ListaCarro[x+1:]...)
		}
	}
	json.NewEncoder(w).Encode(ListaCarro)
}

func ActualizarCarroCambio(w http.ResponseWriter, r *http.Request){
	var tiendasec []Estructuras.ProductoAngular
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &tiendasec)
	for _,rec := range tiendasec{
		validarpedidos(rec)
	}
	json.NewEncoder(w).Encode(ListaCarro)
}

func validarpedidos(tiendaselec Estructuras.ProductoAngular){
	for x,rec := range ListaCarro {
		if rec.Departamento == tiendaselec.Departamento && rec.Nombre == tiendaselec.Nombre && rec.Calificacion == tiendaselec.Calificacion{
			println("Front end", tiendaselec.Nombre, tiendaselec.Cantidad)
			ListaCarro[x].Cantidad = tiendaselec.Cantidad
			println("Servidor: " ,rec.Nombre, rec.Cantidad)
		}
	}
}


func DevolversumaCarro (w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(SumaCarro())
}

func SumaCarro() float32 {
	var suma float32 = 0
	for _, rec := range ListaCarro{
		suma = suma +  rec.Precio*float32(rec.Cantidad)
	}
	return suma
}

//Actualizar los datos del servidor
func ActualizarInventario (w http.ResponseWriter, r *http.Request) {
	// var fecha
	for _, rec :=  range ListaCarro{
		ActualizarEstadoInventario(rec)
		GenerarPedido(rec)
	}
}

func ActualizarEstadoInventario(rec Estructuras.ProductoAngular){
	for _, depas := range ArregloCali {
		if depas.Calificacion+1 == rec.Calificacion && rec.Departamento == rec.Departamento {
			if depas.Lista.Search(rec.Tienda) != -1 {
				producto := Estructuras.Producto{
					Nombre:      rec.Nombre,
					Codigo:      rec.Codigo,
					Descripcion: rec.Descripcion,
					Precio:      rec.Precio,
					Cantidad:    -rec.Cantidad,
					Imagen:      rec.Imagen,
				}
				depas.Lista.SearchIndex(depas.Lista.Search(rec.Tienda)).GetTienda().Productos.SumarInventario(producto)
			}
		}
	}
	var ReiniciarLista []Estructuras.ProductoAngular
	var NuevaListaCarro []Estructuras.ProductoRobot
	for _,i := range ListaCarro{
		nuevo:= Estructuras.ProductoRobot{
			Tienda:         i.Tienda,
			Departamento:   i.Departamento,
			Calificacion:   i.Calificacion,
			Nombre:         i.Nombre,
			Codigo:         i.Codigo,
			Descripcion:    i.Descripcion,
			Precio:         i.Precio,
			Cantidad:       i.Cantidad,
			CantidadMax:    i.CantidadMax,
			Imagen:         i.Imagen,
			Almacenamiento: i.Almacenamiento,
			Recolectado:    false,
		}
		NuevaListaCarro = append(NuevaListaCarro, nuevo)
	}
	AccionarRobot(NuevaListaCarro)
	ListaCarro = ReiniciarLista
}
func GenerarPedido(rec Estructuras.ProductoAngular){
	today := time.Now()
	dia := today.Day()
	var mes int
	anio := today.Year()
	if today.Month().String() == "March"{
		mes = 3
	}else if today.Month().String() == "April"{
		mes = 4
	}
	inven:= Estructuras.Invetarios{
		Tienda:       rec.Tienda,
		Departamento: rec.Departamento,
		Calificacion: rec.Calificacion,
	}
	if buscarTienda(inven) {
		if !ArbolPedidos.SearchAnioExists(anio) {
			ArbolPedidos.Insertar(anio)
		}

		pro := Estructuras.Producto{
			Codigo: rec.Codigo,
		}
		if buscarProducto(inven, pro) {
			if ArbolPedidos.ReturnAnioNode(anio).Mes.IsEmpty() {
				ArbolPedidos.ReturnAnioNode(anio).Mes.Add(mes)
			} else {
				if ArbolPedidos.ReturnAnioNode(anio).Mes.Search(mes) != -1 {
					ArbolPedidos.ReturnAnioNode(anio).Mes.Add(mes)
				}
				ArbolPedidos.ReturnAnioNode(anio).Mes.SearchNReturnM(mes).GetMatriz().Insert(dia, inven.Departamento)
				nuevo := Estructuras.Productos{Codigo: rec.Codigo}
				ArbolPedidos.ReturnAnioNode(anio).Mes.SearchNReturnM(mes).GetMatriz().SearchNreturn(dia, inven.Departamento).Productos.Add(nuevo)
			}
		}
	}
}

// Devolver datos para la intefaz

func ObtenerAniosPedido(w http.ResponseWriter, r *http.Request){
	fmt.Println("devolviendo años ")
	//ListaAnios := *ArbolPedidos.ReturnProductsOfTree()
	fmt.Println(*ArbolPedidos.ReturnProductsOfTree())
	json.NewEncoder(w).Encode(*ArbolPedidos.ReturnProductsOfTree())
}


var anioselectedAngular int
func ObtenerMesPedido(w http.ResponseWriter, r *http.Request){
	var anio int
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &anio)
	anioselectedAngular = anio
}
var messelectedAngular int
func EstablecerMesBack(w http.ResponseWriter, r *http.Request){
	var anio int
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &anio)
	messelectedAngular = anio
	fmt.Println(messelectedAngular)
}



func DevolverMesesPedido (w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(ArbolPedidos.ReturnAnioNode(anioselectedAngular).Mes.ReturnMeses())
}

func DevolverDias(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(ArbolPedidos.ReturnAnioNode(anioselectedAngular).Mes.SearchIndex(ArbolPedidos.ReturnAnioNode(anioselectedAngular).Mes.Search(messelectedAngular)).GetMatriz().ReturnDias())
}

var DiaActual int
var CategoriaAct string
func EstablecerDia(w http.ResponseWriter, r *http.Request){
	var anio Estructuras.ProductosPedido
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &anio)
	fmt.Println(anio)
	DiaActual = anio.Dia
	CategoriaAct = anio.Categoria
}


func DevolverListaProductPedidos(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(ArbolPedidos.ReturnAnioNode(anioselectedAngular).Mes.SearchIndex(ArbolPedidos.ReturnAnioNode(anioselectedAngular).Mes.Search(messelectedAngular)).GetMatriz().SearchNreturn(DiaActual,CategoriaAct).Productos.ReturnNodes())
}

func CrearDotAnio (w http.ResponseWriter, r *http.Request){
	//var path = "D:/Escritorio/USAC/EDD/Proyecto/Practica 1/ComponentesAngular/Proyecto/src/assets/archivosd/anios.dot"
}

//Cuentas y sus verificaciones ------------------------------------------------------------------------------------------------
var UsuarioLoged *Estructuras.Usuario

func VerificarCuenta (w http.ResponseWriter, r *http.Request) {
	var UsuarioVer Estructuras.UsuarioJson
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &UsuarioVer)
	fmt.Println(UsuarioVer)

	UsuarioN := Estructuras.Usuario{
		Dpi:      UsuarioVer.Dpi,
		Nombre:   "",
		Correo:   "",
		Password: UsuarioVer.Password,
		Cuenta:   "",
	}
	if ArbolUsuarios.Search(UsuarioN) != nil {

		selected := ArbolUsuarios.Search(UsuarioN)

		UsuarioN = Estructuras.Usuario{
			Dpi:      selected.Dpi,
			Nombre:   selected.Nombre,
			Correo:   selected.Correo,
			Password: selected.Password,
			Cuenta:   selected.Cuenta,
		}

		if UsuarioN.Password == UsuarioVer.Password{
			json.NewEncoder(w).Encode(true)
			UsuarioLoged = &UsuarioN
		}else{
			json.NewEncoder(w).Encode(false)
		}

	}else{
		json.NewEncoder(w).Encode(false)
	}
}

func VerificarSiAdmin (w http.ResponseWriter, r *http.Request) {

	if UsuarioLoged != nil {
		if ArbolUsuarios.Search(*UsuarioLoged) != nil {
			fmt.Println(*UsuarioLoged)
			if UsuarioLoged.Cuenta == "Admin"{
				json.NewEncoder(w).Encode(true)
			}else{
				json.NewEncoder(w).Encode(false)
			}
		} else {
			json.NewEncoder(w).Encode(false)
		}
	}else{
		json.NewEncoder(w).Encode(false)
	}
}

func DevolverCuenta (w http.ResponseWriter, r *http.Request) {
	var mandar Estructuras.MostrarUsuario

	Usuario := Estructuras.Usuario{
		Dpi:      0,
		Nombre:   "",
		Correo:   "",
		Password: "",
		Cuenta:   "",
	}
	if UsuarioLoged != nil {
		mandar = Estructuras.MostrarUsuario{true, *UsuarioLoged}
	}else{
		mandar = Estructuras.MostrarUsuario{false, Usuario}
	}

	json.NewEncoder(w).Encode(mandar)
}

func CargarUsuarios(w http.ResponseWriter, r *http.Request){
	var Users CargaDeUsers
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &Users)
	for _,i:= range Users.Usuarios {//recorer el arreglo de usuarios
		IngresarUsuario(i)
	}
	//fmt.Println(ArbolUsuarios.Gragicar(0))
}
type CargaDeUsers struct{
	Usuarios []Estructuras.Usuario
}
func IngresarUsuario(usuario Estructuras.Usuario){
	ArbolUsuarios.Insert(usuario)
}

func RegistrarUsuario(w http.ResponseWriter, r *http.Request){
	var Users Estructuras.Usuario
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//fmt.Println(r.Body)
	json.Unmarshal(reqBody, &Users)
	fmt.Println(Users)
	IngresarUsuario(Users)
	//json.NewEncoder(w).Encode(true)
}

func EliminarUsuario(w http.ResponseWriter, r *http.Request){
	//fmt.Println("elminando")
	var contraseña Estructuras.UsuarioJson
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	//fmt.Println(r.Body)
	json.Unmarshal(reqBody, &contraseña)



	if contraseña.Password == UsuarioLoged.Password{
		var Usuario *Estructuras.Usuario
		ArbolUsuarios.Remove(*UsuarioLoged)
		UsuarioLoged = Usuario
		ArbolUsuarios.ImprimirArbol()
		fmt.Println("Eliminado", Usuario)
		json.NewEncoder(w).Encode(true)
	}else {
		json.NewEncoder(w).Encode(false)
	}
}

func GenerarReporte(w http.ResponseWriter, r *http.Request){
	var llave string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &llave)

	if llave != ""{
		Estructuras.LLaveEncriptado = llave
	}

	ArbolUsuarios.Gragicar(0)
	ArbolUsuarios.Gragicar(1)
	ArbolUsuarios.Gragicar(2)
}


//Grafo-----------------------------------------------------------------------------------------------------------------
var GrafoEntragas Estructuras.Grafo
var PosicionIniciarRobot string
var PosicionActualRobot string
var EntregaRobot string
func CargarEntregas(w http.ResponseWriter, r *http.Request)  {

	var Paquete Estructuras.PaqueteGrafos
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos Inválidos")
	}
	json.Unmarshal(reqBody, &Paquete)


	for _,i:= range Paquete.Nodos {//Insertar uno por uno los vertices
		if GrafoEntragas.GetVertice(i.Nombre) == nil {
			GrafoEntragas.InsertVertice(i.Nombre)
		}
	}
	for _,i:= range Paquete.Nodos {//Insertar las aristas, o transiciones
		for _,j:= range i.Enlaces {
			//fmt.Println("Nombre i: ",i.Nombre," Nombre j: " ,j.Nombre," Distancia: " ,j.Distancia)
			GrafoEntragas.InsertArista(i.Nombre,j.Nombre,j.Distancia)
			GrafoEntragas.InsertArista(j.Nombre,i.Nombre,j.Distancia)//Hace que sea bi direccional cada arista
		}
	}
	fmt.Println("----------------------------------------")
	//GrafoEntragas.RutaMasCorta("Despacho","Textiles",true)
	PosicionIniciarRobot = Paquete.PosicionInicialRobot
	PosicionActualRobot =Paquete.PosicionInicialRobot
	EntregaRobot = Paquete.Entrega
	GrafoEntragas.GraficarGrafo()
}
var ContarMovRobot int
func AccionarRobot(ListaCarro []Estructuras.ProductoRobot){
	var Comparador int
	contadorRecolectados := 0
	for _,i := range ListaCarro{
		if !i.Recolectado {
			encontrado,actual := GrafoEntragas.RutaMasCorta(PosicionActualRobot,i.Almacenamiento,false)
			if encontrado{
				Comparador = actual
				break
			}
		}
	}
	posicionDeMenor := -1
	for x,i := range ListaCarro{
		if !i.Recolectado {
			encontrado,distanciaAct := GrafoEntragas.RutaMasCorta(PosicionActualRobot,i.Almacenamiento,false)
			if encontrado {
				if Comparador >= distanciaAct {
					Comparador = distanciaAct
					posicionDeMenor = x
				}
			}
		}else{
			contadorRecolectados ++
		}
	}
	if posicionDeMenor != -1 {
		GrafoEntragas.RutaMasCorta(PosicionActualRobot, ListaCarro[posicionDeMenor].Almacenamiento, true)
		PosicionActualRobot = ListaCarro[posicionDeMenor].Almacenamiento
		ListaCarro[posicionDeMenor].Recolectado = true
		ContarMovRobot++
		AccionarRobot(ListaCarro)
	}
	if contadorRecolectados == len(ListaCarro)-1{
		GrafoEntragas.RutaMasCorta(PosicionActualRobot, EntregaRobot, true)
		ContarMovRobot++
		GrafoEntragas.RutaMasCorta(EntregaRobot, PosicionIniciarRobot, true)
		ContarMovRobot++
		PosicionActualRobot = PosicionIniciarRobot
	}

}
func MostrarMovimientosRobot(w http.ResponseWriter, r *http.Request){
	var arreglo []int
	for i := 0; i<ContarMovRobot;i++{
		arreglo = append(arreglo, i)
	}
	json.NewEncoder(w).Encode(arreglo)
}
//CREAR UNA MATRIZ CON DIMENSIONES ESTATICAS
//actualmente no se usa



func crearMatriz3D( indice, departamento, calificacion int) [][][]Estructuras.ListaDoble {
	result := make([][][]Estructuras.ListaDoble,indice)
	for i := 0 ; i < indice ; i++ {
		result[i] = make([][]Estructuras.ListaDoble,departamento);
		for j := 0; j < departamento; j++ {
			result[i][j] = make([]Estructuras.ListaDoble,calificacion);
			for k := 0 ; k < calificacion; k++ {
				//result[i][j][k] = Estructuras.NewListaDoble()
			}
		}
	}
	return result
}



