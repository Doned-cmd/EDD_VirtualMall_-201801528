package Estructuras

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type Arista struct {
	next *Arista
	adjacent *Vertice
	peso int
}

func NewArista(next *Arista, adjecent *Vertice, peso int) *Arista {
	return &Arista{
		next:     next,
		adjacent: adjecent,
		peso:     peso,
	}
}

type Vertice struct {
	next *Vertice
	adjecent *Arista
	nombre string
}
func NewVertice(next *Vertice, adjecent *Arista, nombre string) *Vertice {
	return &Vertice{
		next:     next,
		adjecent: adjecent,
		nombre:   nombre,
	}
}
func (A *Vertice) CompareTo(o *Vertice ) int {
	if o.nombre == A.nombre{
		return 0
	}else if o.nombre > A.nombre{
		return 1
	}else{
		return -1
	}
}

type Grafo struct {
	root *Vertice
}

func (g *Grafo) isEmpety() bool{
	return g.root == nil
}

func (g *Grafo) GetVertice(nombre string) *Vertice{
	aux := g.root
	for aux != nil{
		if aux.nombre == nombre{
			return aux
		}
		aux = aux.next
	}
	return nil
}

func (g *Grafo) NumeroVertice (nombre string) int {
	aux := g.root
	contador := 0
	for aux != nil{
		if aux.nombre == nombre{
			return contador
		}
		aux = aux.next
		contador ++
	}
	return 0
}

func (g *Grafo) InsertArista(Origen string, Destino string, peso int) {
	nuevo := NewArista(nil,nil,peso)
	//fmt.Println("Desde: ", Origen," a :", Destino)
	origen := g.GetVertice(Origen)
	fmt.Println(origen)
	destino := g.GetVertice(Destino)
	fmt.Println(destino)
	aux := origen.adjecent
	if aux == nil {
		origen.adjecent = nuevo
		nuevo.adjacent = destino
		//fmt.Println("Se inserto con exito")
	}else{
		for aux.next != nil{
			aux = aux.next
		}
		aux.next = nuevo
		nuevo.adjacent = destino
		//fmt.Println("Se inserto con exito")
	}
}



func (g *Grafo) InsertVertice(nombre string){
	nuevo := NewVertice(nil,nil,nombre)
	if g.isEmpety(){
		g.root = nuevo
	}else{
		aux := g.root
		for aux.next != nil{
			aux = aux.next
		}
		aux.next = nuevo
	}
	fmt.Println(nuevo)
}

func (g *Grafo) ListADyecencia() string {
	auxVertice :=g.root
	var auxArista *Arista
	cadena := "digraph grafo {"
	cadena += "node [shape=box]\n"
	back := ""
	next := ""
	alinear := "{rank=same;"
	for auxVertice != nil{
		auxArista = auxVertice.adjecent
		cadena += g.generarId(auxVertice,auxArista) + g.generarLabel(auxVertice) + "\n"
		back = g.generarId2(auxVertice)
		alinear += back + ";"
		for auxArista != nil{
			cadena += g.generarId(auxVertice, auxArista) + g.generarLabel(auxArista.adjacent) + "\n"
			next = g.generarId(auxVertice, auxArista)
			cadena += back + "->" + next + "\n"
			back = g.generarId(auxVertice, auxArista)
			alinear += back + ";"
			auxArista = auxArista.next
		}
		alinear += "}\n"
		if auxVertice.next != nil{
			cadena += g.generarId2(auxVertice) + "->" + g.generarId2(auxVertice.next) + "\n"
		}
		cadena += alinear
		alinear = "{rank=same;"
		auxVertice = auxVertice.next
	}
	cadena += "}"
	Graficar(cadena,"Adyecencia",1000)
	return cadena
}

func (g *Grafo) GraficarGrafo(){
	auxVertice :=g.root
	var auxArista *Arista
	cadena := "digraph grafo {"
	cadena += "Node [shape=box]\n"

	//alinear := "{rang=same;"
	for auxVertice != nil{
		cadena += "Node" + strconv.Itoa(g.NumeroVertice(auxVertice.nombre)) + g.generarLabel(auxVertice) + "\n"
		//if auxVertice
		auxArista = auxVertice.adjecent
		for auxArista!= nil{
			cadena += "Node" + strconv.Itoa(g.NumeroVertice(auxVertice.nombre)) + "->" + "Node" + strconv.Itoa(g.NumeroVertice(auxArista.adjacent.nombre)) +" [ label=\"" +strconv.Itoa(auxArista.peso)+"\""+"]" + "\n"
			auxArista = auxArista.next
		}
		auxVertice = auxVertice.next
	}

	cadena += "}"
	fmt.Println("Dot generado")
	Graficar(cadena,"Grafo",0)
}


func (g *Grafo) Show(){
	auxVertice := g.root
	var auxArista *Arista
	for auxVertice != nil{
		auxArista = auxVertice.adjecent
		fmt.Println( auxVertice.nombre)
		for auxArista != nil{
			fmt.Println("->", auxArista.adjacent.nombre)
			auxArista = auxArista.next
		}
		auxVertice = auxVertice.next
		fmt.Println("")
	}
}

func (g *Grafo) GraficarRutas() string{
	cadena := "digraph grafo {"
	cadena += "graph [ dpi = 300 ];\n"
	cadena +=  "rankdir=LR size=\"8,5\"\n" +
		       "    node [shape=circle]\n"
	auxVertice := g.root
	var auxArista *Arista
	for auxVertice != nil{
		cadena += "\"" + auxVertice.nombre + "\"\n"
		auxArista = auxVertice.adjecent
		for auxArista != nil{
			cadena += "\"" + auxVertice.nombre + "\"->\"" + auxArista.adjacent.nombre + "\"[label=\"" + strconv.Itoa(auxArista.peso) + "\"]\n" // string(auxArista.peso)
			auxArista = auxArista.next
		}
		auxVertice = auxVertice.next
	}
	cadena += "}"
	Graficar(cadena,"RutasSolamente",0)
	return cadena
}


func (g *Grafo) generarLabel(vertice *Vertice) string{
	cadena := ""
	cadena += "[label=\"" + vertice.nombre + "\"]"
	return cadena
}


func (g *Grafo) generarId(vertice *Vertice, arista *Arista) string{
	cadena := ""
	cadena += "\"" + vertice.nombre + arista.adjacent.nombre+ "\""
	return cadena
}

func (g *Grafo) generarId2(vertice *Vertice) string{
	cadena := ""
	cadena += "\"" + vertice.nombre + "\""
	return cadena
}

func (g *Grafo) RecorridoAnchura(origen *Vertice){
	var band int
	var band2 int
	var actual *Vertice

	var cola []*Vertice
	var lista []*Vertice

	cola = append(cola, origen)
	for !(len(cola)==0){
		band = 0
		actual = cola[0]
		cola = append(cola[:0], cola[0+1:]...)
		for _,i := range lista {
			if i == actual{
				band = 1
			}
		}
		if band == 0{
			fmt.Println(actual.nombre+ ", ")
			lista = append(lista, actual)
			var aux *Arista
			aux = actual.adjecent
			for aux != nil{
				band2 = 0
				for _,i := range lista {
					if aux.adjacent == i {
						band2 = 1
					}
				}
				if band2 == 0{
					cola = append(cola, aux.adjacent)
				}
				aux = aux.next
			}
		}
	}
}

func (g *Grafo) RecorridoProfundidad(origen *Vertice)  {
	var band, band2 int
	var actual *Vertice
	var pila []*Vertice
	var lista []*Vertice
	pila = append(pila, origen)

	for !(len(pila)==0) {
		band = 0
		actual = pila[len(pila)-1]
		pila = append(pila[:len(pila)-1], pila[len(pila)-1+1:]...)

		for _, i := range lista {
			if i == actual {
				band = 1
			}
		}

		if band == 0 {
			fmt.Println(actual.nombre + ", ")
			lista = append(lista, actual)

			aux := actual.adjecent
			for aux != nil {
				band2 = 0
				for _, i := range lista {
					if aux.adjacent == i {
						band2 = 1
					}
				}
				if band2 == 0{
					pila = append(pila, aux.adjacent)
				}
				aux = aux.next
			}
		}
	}
}
var NumeroGrafo int

func (g *Grafo) RutaMasCorta(Origen string, Destino string, graficar bool) (bool,int) {
	origen:= g.GetVertice(Origen)
	destino:= g.GetVertice(Destino)
	var ruta []*Ruta
	CostoActual := 0
	var VerticeActual, destinoActual  *Vertice
	var aux *Arista
	band := 0
	band2 := 0
	var listaCosto []*VerticeCosto
	var listaOrdenada []*VerticeCosto
	var pila []*VerticeToVertice
	//var Camino []*VerticeToVertice
	listaCosto = append(listaCosto, NewVerticeCosto(origen,0))
	listaOrdenada = append(listaOrdenada, NewVerticeCosto(origen,0))
	for !(len(listaOrdenada)==0){
		VerticeActual = listaOrdenada[0].vertice
		CostoActual = listaOrdenada[0].costo
		listaOrdenada = append(listaOrdenada[:0], listaOrdenada[0+1:]...)
		if VerticeActual.nombre == destino.nombre{
			fmt.Println("Costo " , CostoActual)
			band2 = 1
			destinoActual = destino
			for !(len(pila)==0){

				//Agregar ruta------ a alguna lista aqui
				ruta = append(ruta, NewRuta(destinoActual.nombre,CostoActual))
				//Camino = append(Camino, NewVerticeToVertice(destinoActual,VerticeActual))
				for (len(pila)!=0) && (pila[len(pila)-1].destino.nombre != destinoActual.nombre){
					pila =  append(pila[:len(pila)-1], pila[len(pila)-1+1:]...)
				}
				if len(pila)!=0{
					destinoActual = pila[len(pila)-1].origen
				}
			}
			fmt.Println("Costo Real:" + strconv.Itoa(g.EncontrarDistancia(ruta)))
			break
		}
		aux = VerticeActual.adjecent
		for aux !=nil{
			band = 0
			CostoActual = CostoActual + aux.peso
			for x,i := range listaCosto {
				if aux.adjacent.nombre == listaCosto[x].vertice.nombre {
					band = 1
					if CostoActual < i.costo {
						listaCosto[x].costo = CostoActual
						for y,_ := range listaOrdenada{
							if listaOrdenada[y].vertice.nombre == aux.adjacent.nombre{
								listaOrdenada[y].costo = CostoActual
							}
						}
						listaOrdenada = ordenar(listaOrdenada)
						pila = append(pila, NewVerticeToVertice(VerticeActual,aux.adjacent))
						CostoActual = CostoActual - aux.peso
					}
				}
			}
			if band == 0{
				listaCosto = append(listaCosto, NewVerticeCosto(aux.adjacent,CostoActual))
				listaOrdenada = append(listaOrdenada, NewVerticeCosto(aux.adjacent, CostoActual))
				listaOrdenada = ordenar(listaOrdenada)
				pila = append(pila, NewVerticeToVertice(VerticeActual,aux.adjacent))
				CostoActual = CostoActual -aux.peso
			}
			aux = aux.next
		}
	}
	fmt.Println("")
	if band2 == 0{
		fmt.Println("No se encontro ruta")
		return false,0
	}else{
		fmt.Println("Se encontro ruta")
		if graficar {
			//fmt.Println(listaOrdenada)
			//temporal := ruta[0]
			cadena := "digraph grafo {"
			//cadena +=   "graph [ dpi = 300 ];\n"
			cadena += "rankdir=LR size=\"8,5\"\n" +
				"	node [shape=record]\n"
			cadena += label(ruta)
			for i := len(ruta) - 1; i >= 0; i-- {
				if i > 0 {
					cadena += ruta[i].lugar + "->"
				} else {
					cadena += ruta[i].lugar
				}
			}
			cadena += "\n}"
			//Metodo para graficar
			Graficar(cadena, "CaminoCorto", NumeroGrafo)
			fmt.Println("Imagen generada")
			NumeroGrafo++
		}
		return true, g.EncontrarDistancia(ruta)
	}
}

func (g *Grafo) EncontrarDistancia(ruta []*Ruta) int {
	auxVertice := g.root
	contador := 0
	for x:=0 ; x<len(ruta)-1;x++{
		for auxVertice != nil{
			if ruta[x].lugar == auxVertice.nombre{
				auxArista := auxVertice.adjecent
				for  auxArista!= nil{
					if auxArista.adjacent.nombre == ruta[x+1].lugar{
						contador = contador + auxArista.peso
					}
					auxArista = auxArista.next
				}
			}
			auxVertice = auxVertice.next
		}
		auxVertice = g.root
	}
	return contador
}

func label(ruta []*Ruta) string {
	cadena := ""
	for i := len(ruta)-1;i>=0;i--{
		//di := strconv.Itoa(ruta[i].distancia)
		cadena += ruta[i].lugar  + "[label=\"" + "Lugar: " + ruta[i].lugar  + "\"]\n"
		// + di  + "\\nDistancia: "
	}
	return cadena
}

func Graficar(texto string, nombre string ,num int) string{
	direccion :="D:/Escritorio/USAC/EDD/Proyecto/Reportes/"+nombre+strconv.Itoa(num)+".dot"
	direccionImg := "D:/Escritorio/USAC/EDD/Proyecto/Practica 1/ComponentesAngular/Proyecto/src/assets/archivosd/"+nombre+strconv.Itoa(num) +".svg"

	f, err := os.Create(direccion)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(texto)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Archivo dot escrito.")


	path,_ := exec.LookPath("dot")
	cmd,_ := exec.Command(path, "-Tsvg", direccion).Output()
	mode := int(0777)
	ioutil.WriteFile(direccionImg, cmd, os.FileMode(mode))
	fmt.Println("Imagen generada")
	return texto
}

func DevolverTope(pila []*VerticeToVertice) *VerticeToVertice{
	if  len(pila) !=0 {
		return pila[0]
	}
	return nil
}

func ordenar(ListaAOrdenar []*VerticeCosto) []*VerticeCosto{
	var auxiliar *VerticeCosto
	for i := 0; i < len(ListaAOrdenar); i++ {
		for j := 0; j < len(ListaAOrdenar); j++ {
			if ListaAOrdenar[i].compareTo(ListaAOrdenar[j]) > 0{
				auxiliar = ListaAOrdenar[i]
				ListaAOrdenar[i] = ListaAOrdenar[j]
				ListaAOrdenar[j] = auxiliar
			}
		}
	}
	//fmt.Println("------------------")
	//for _,x := range ListaAOrdenar {
	//	fmt.Println("Lista Ordenada", x.costo)
	//}
	//fmt.Println("------------------")
	return ListaAOrdenar
}


type VerticeCosto struct {
	vertice *Vertice
	costo int
}

func NewVerticeCosto(vertice *Vertice, costo int) *VerticeCosto {
	return &VerticeCosto{vertice: vertice, costo: costo}
}

func (v *VerticeCosto) compareTo(o *VerticeCosto) int {
	if o.costo == v.costo{
		return 0
	}else if o.costo > v.costo{
		return 1
	}else{
		return -1
	}
}

type VerticeToVertice struct {
	origen *Vertice
	destino *Vertice
}

func NewVerticeToVertice(origen *Vertice, destino *Vertice) *VerticeToVertice {
	return &VerticeToVertice{origen: origen, destino: destino}
}

func (v *VerticeToVertice) compareTo(o *VerticeToVertice) int {
	if (o.origen.CompareTo(v.origen) == 0)&&(o.destino.CompareTo(v.destino)) == 0{
		return 0
	}
		return -1
}

type Ruta struct {
	lugar string
	distancia int
}

func NewRuta(lugar string, distancia int) *Ruta {
	return &Ruta{lugar: lugar, distancia: distancia}
}