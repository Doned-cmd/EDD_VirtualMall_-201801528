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
	IndiceNOmbre string
	Departamento string
	Calificacion int

	Lista *ListaDobleEnlazada.ListaDoble
}


var matricita []Indice

func main() {

  	if "hola"<"i"{
  		println("b es mayor")
	}

	//fmt.Print()
	//request()

}




func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Servidor en Go")
}

func getArreglo(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "[1,2,3,4]")
}

func metodopost(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Json.Sobre
	json.Unmarshal(body, &re)
	//fmt.Println(re)
	fmt.Println(re.Datos)
	ConvertToMatrix(re)
}

func request(){
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", homePage)
	myrouter.HandleFunc("/getArreglo",getArreglo).Methods("GET")
	myrouter.HandleFunc("/metodopost", metodopost).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", myrouter))
}





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



func crearMatriz3D( indice, departamento, calificacion int) [][][]*ListaDobleEnlazada.ListaDoble {
	result := make([][][]*ListaDobleEnlazada.ListaDoble,indice)
	for i := 0 ; i < indice ; i++ {
		result[i] = make([][]*ListaDobleEnlazada.ListaDoble,departamento);
		for j := 0; j < departamento; j++ {
			result[i][j] = make([]*ListaDobleEnlazada.ListaDoble,calificacion);
			for k := 0 ; k < calificacion; k++ {
				//result[i][j][k] = ListaDobleEnlazada.NewListaDoble()
			}
		}
	}
	return result
}



