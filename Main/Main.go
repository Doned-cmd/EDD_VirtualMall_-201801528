package main

import (
	"../Json"
	"../ListaDobleEnlazada"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	Lista *ListaDobleEnlazada.ListaDoble
}


//Arreglo Linealizado

type ArregloLinea struct {
	IndiceNombre string
	Departamento string
	Calificacion int

	Lista *ListaDobleEnlazada.ListaDoble
}

//Arreglos de datos
var matricita []Indice
var ArregloCali []ArregloLinea


func main() {

	//fmt.Print()
	request()
	//print(Regresar(3))
}

func Regresar(b int) int {
	 contador := -1
	a := [5]int{0,1,2,3,4}
	for _, i:=  range a {
		contador++
		if i == b {
			return contador
		}
	}
	println("se realiza")
	return -1
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
	log.Fatal(http.ListenAndServe(":3000", myrouter))
}




func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Servidor en Go")
}



//REGRESAR GRAFICO DE COMO ESTA LA MEMORIA
func getArreglo(w http.ResponseWriter, r *http.Request)  {
	var path = "D:/Escritorio/USAC/EDD/Practica 1/Grafico.dot"
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

	// Escribe algo de texto linea por linea
	_, err = file.WriteString(text)
	if existeError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
}

//CARGAR JSON CON TIENDAS-----------------------------------------------------------------------------------------
func cargartienda(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Json.Sobre
	json.Unmarshal(body, &re)
	fmt.Println(re.Datos)
	ConvertToMatrix(re)
	ConvertToArray(matricita)
}
//convertir a matriz
func  ConvertToMatrix (sbr Json.Sobre) {

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
					Lista:      ListaDobleEnlazada.NewListaDoble(),
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
	println("Matricitia")
	for _,indi := range matricita{
		for _,depa := range indi.Departamentos{
			for _,cal := range depa.Calificaciones{
				cal.Lista.ImprimirLista()
			}
		}
	}
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

	for num,i:= range arreglo{
		print( num, i.IndiceNombre, " ",i.Departamento, " ",i.Calificacion, " ")
		i.Lista.ImprimirLista()
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
					Lista:        ListaDobleEnlazada.NewListaDoble(),
				}
				arreglo = append(arreglo, nuevodatoarreglo)
			}
		}
	}
	return arreglo
}

func ArregloBurbuja(ListaAOrdenar []Json.Tienda) []Json.Tienda{
	var auxiliar Json.Tienda
	for i := 0; i < len(ListaAOrdenar); i++ {
		for j := i+1; j < len(ListaAOrdenar); j++ {
			if ListaAOrdenar[i].Nombre > ListaAOrdenar[j].Nombre {
				auxiliar = ListaAOrdenar[i]
				ListaAOrdenar[i] = ListaAOrdenar[j]
				ListaAOrdenar[j] = auxiliar
			}
		}
	}
	return ListaAOrdenar
}



//BUSCAR TIENDA ESPECIFICA-----------------------------------------------------------------------------------------
func TiendaEspecifica(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Json.BusquedaEspecifica
	json.Unmarshal(body, &re)
	println("Buscando coincidencias con los siguientes parametros: ","Nombre: " ,re.Nombre, "Departamento: ",re.Departamento,"Calificacion: " ,re.Calificacion)
//EJECUTAR METODO PARA ENCONTRAR LA TIENDA
	var TiendaEncontrada Json.Tienda
	imprimir := true
	for _, cali := range ArregloCali{
			if (cali.Calificacion == re.Calificacion) && (cali.Departamento == re.Departamento) {
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


func GenerarJsonCoincidencia(Tienda Json.Tienda) {

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
	var re Json.EliminarTienda
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

	var Matriz = matricita

	for  i:=0; i<len(Matriz);i++ {
		for j :=0; j <len(Matriz[i].Departamentos); j++ {
			for k :=0; k <len(Matriz[i].Departamentos[j].Calificaciones); k++ {
				
				Matriz[i].Departamentos[j].Calificaciones[k].Lista = ArregloCali[k+(len(Matriz))*(k+(len(Matriz[i].Departamentos))*k)].Lista
			}
		}
	}

	crear_json, _ := json.Marshal(Matriz)
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



//CREAR UNA MATRIZ CON DIMENSIONES ESTATICAS
//actualmente no se usa

func crearMatriz3D( indice, departamento, calificacion int) [][][]ListaDobleEnlazada.ListaDoble {
	result := make([][][]ListaDobleEnlazada.ListaDoble,indice)
	for i := 0 ; i < indice ; i++ {
		result[i] = make([][]ListaDobleEnlazada.ListaDoble,departamento);
		for j := 0; j < departamento; j++ {
			result[i][j] = make([]ListaDobleEnlazada.ListaDoble,calificacion);
			for k := 0 ; k < calificacion; k++ {
				//result[i][j][k] = ListaDobleEnlazada.NewListaDoble()
			}
		}
	}
	return result
}



