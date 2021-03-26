package Estructuras
import "fmt"



type NodoMatriz struct {
	// Estos atributos son especificos para la matriz
	Dia int
	Categoria string //saber en que cabecera se esta
	Productos *ColaPedidos // objeto a guardar
	izquierdo, derecho, arriba, abajo *NodoMatriz // nodos con los que se desplaza dentro de la matriz
	//Estos atributos son especificos para la lista
	headerd int // tipo interno de caecera
	headercat string
	siguiente, anterior *NodoMatriz // nodos con los que se va a desplazar dentro de las listas
}

type lista struct {
	first, last *NodoMatriz
}

type matriz struct {
	lst_h, lst_y *lista //y = categorias, h = dias
}

func newLista() *lista{
	return &lista{nil,nil}
}

func NewMatriz()*matriz{
	return &matriz{newLista(),newLista()}
}

func nodoMatriz(dia int, categoria string) *NodoMatriz {
	return &NodoMatriz{dia, categoria, NewColaPedidos(),nil,nil,nil,nil,0,"",nil, nil,}

}
func nodoListaDia(header int) *NodoMatriz {
	return &NodoMatriz{0,"",nil,nil,nil,nil,nil,header,"",nil,nil}
}
func nodoListaCat(header string) *NodoMatriz {
	return &NodoMatriz{0,"",nil,nil,nil,nil,nil,0,header,nil,nil}
}
func (n *NodoMatriz) headerx() int  { return n.Dia }
func  (n *NodoMatriz) headery() string { return n.Categoria }
//func (n *NodoMatriz) toString() string { return "Nombre: " + n.producto.Nombre + "\nDespcripcion: "+n.producto.Descripcion}

func (l *lista) ordenardia(nuevo *NodoMatriz) {
	aux:= l.first
	for aux != nil {
		if nuevo.headerd > aux.headerd{
			aux = aux.siguiente
		}else{
			if aux == l.first{
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			}else{
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}
func (l *lista) ordenarcat(nuevo *NodoMatriz) {
	aux:= l.first
	for aux != nil {
		if nuevo.headercat > aux.headercat{
			aux = aux.siguiente
		}else{
			if aux == l.first{
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			}else{
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}

func (l *lista) insertdia(header int)  {
	nuevo := nodoListaDia(header)
	if l.first == nil{
		l.first =nuevo
		l.last = nuevo
	}else {
		l.ordenardia(nuevo)
	}
}
func (l *lista) insertcat(header string)  {
	nuevo := nodoListaCat(header)
	if l.first == nil{
		l.first =nuevo
		l.last = nuevo
	}else {
		l.ordenarcat(nuevo)
	}
}




func(m *matriz) Insert( dia int, categoria string){
	h := m.lst_h.searchDia(dia)
	v := m.lst_y.searchCat(categoria)

	if h == nil && v == nil{
		m.noExiste(dia, categoria)
	} else if h == nil {
		m.existeVertical(dia, categoria)
	} else if v == nil{
		m.existeHorizontal( dia, categoria)
	}else{
		m.existe(dia, categoria)
	}
}


func (l *lista) searchDia(header int) *NodoMatriz {
	temp := l.first
	for temp != nil{
		if temp.headerd == header{
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}
func (l *lista) searchCat(header string) *NodoMatriz {
	temp := l.first
	for temp != nil{
		if temp.headercat == header{
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}


func (l *lista ) print() {
	temp := l.first
	for temp != nil {
		fmt.Print("Cabecera", temp.headerd)
		temp = temp.siguiente
	}
}

func (m *matriz) noExiste( x int, y string) {
	m.lst_h.insertdia(x)//se inserta en la lista que emula la cabecera horizontal
	m.lst_y.insertcat(y)// se inserta en la lista que emula la cabecera vertical


	//se busca el NodoMatriz que se acaba de insertar, para poder enlazarlos
	h:= m.lst_h.searchDia(x)
	v := m.lst_y.searchCat(y)

	nuevo := nodoMatriz(x,y)//se crea un nuevo NodoMatriz

	h.abajo = nuevo// se enlaza el NodoMatriz horizontal hacia abajo
	nuevo.arriba = h// se enlaza el nuevo NodoMatriz hacia arriba

	v.derecho = nuevo// se enlaza el NodoMatriz vertical hacia la derecha
	nuevo.izquierdo = v// se enlaza el nuevo NodoMatriz hacia la izquierda
}

func (m *matriz) existeVertical( x int, y string) {
	m.lst_h.insertdia(x)
	h := m.lst_h.searchDia(x)
	v := m.lst_y.searchCat(y)

	nuevo := nodoMatriz(x,y)
	agregado := false

	aux := v.derecho
	var cabecera int

	for aux != nil{
		cabecera = aux.headerx()
		if cabecera < x {
			aux = aux.derecho
		}else{
			nuevo.derecho = aux
			nuevo.izquierdo = aux.izquierdo
			aux.izquierdo.derecho = nuevo
			aux.izquierdo = nuevo
			agregado = true
			break
		}
	}
	if agregado == false{
		aux = v.derecho
		for aux.derecho != nil{
			aux = aux.derecho
		}
		nuevo.izquierdo = aux
		aux.derecho = nuevo
	}
	nuevo.arriba = h
	h.abajo = nuevo
}

func (m *matriz) existeHorizontal( x int, y string) {
	m.lst_y.insertcat(y)
	h := m.lst_h.searchDia(x)
	v := m.lst_y.searchCat(y)

	nuevo := nodoMatriz(x,y)
	agregado := false

	aux := h.abajo
	var cabecera string

	for aux != nil{
		cabecera = aux.headery()
		if cabecera < y {
			aux = aux.abajo
		}else{
			nuevo.abajo = aux
			nuevo.arriba = aux.arriba
			aux.arriba.abajo = nuevo
			aux.arriba = nuevo
			agregado = true
			break
		}
	}
	if agregado == false{
		aux = h.abajo
		for aux.abajo != nil{
			aux = aux.abajo
		}
		nuevo.arriba = aux
		aux.abajo = nuevo
	}
	nuevo.izquierdo = v
	v.derecho = nuevo
}

func (m *matriz) existe( x int, y string) {
	h := m.lst_h.searchDia(x)
	v := m.lst_y.searchCat(y)

	nuevo := nodoMatriz(x,y)
	agregado := false

	aux := v.derecho

	var cabecera int

	for aux != nil{
		cabecera = aux.headerx()
		if cabecera < x {
			aux = aux.derecho
		}else{
			nuevo.derecho = aux
			nuevo.izquierdo = aux.izquierdo
			aux.izquierdo.derecho = nuevo
			aux.izquierdo = nuevo
			agregado = true
			break
		}
	}
	if agregado == false{
		aux = v.derecho
		for aux.derecho != nil{
			aux = aux.derecho
		}
		nuevo.izquierdo = aux
		aux.derecho = nuevo
	}


	agregado = false

	aux = h.abajo
	var cabecera2 string

	for aux != nil{
		cabecera2 = aux.headery()
		if cabecera2 < y {
			aux = aux.abajo
		}else{
			nuevo.abajo = aux
			nuevo.arriba = aux.arriba
			aux.arriba.abajo = nuevo
			aux.arriba = nuevo
			agregado = true
			break
		}
	}
	if agregado == false{
		aux = h.abajo
		for aux.abajo != nil{
			aux = aux.abajo
		}
		nuevo.arriba = aux
		aux.abajo = nuevo
	}
}




func (m *matriz) printVertical(){
	cabecera := m.lst_y.first
	for cabecera != nil{
		aux := cabecera.derecho
		for aux != nil{
			aux = aux.derecho
		}
		cabecera = cabecera.siguiente
	}
}

func (m *matriz) printHorizontal(){
	cabecera := m.lst_h.first
	for cabecera != nil{
		aux := cabecera.abajo
		for aux != nil{
			aux = aux.abajo
		}
		cabecera = cabecera.siguiente
	}
}

func (m *matriz) SearchNreturn (dia int, categoria string) *NodoMatriz{
	cabecera := m.lst_h.first
	for cabecera != nil{
		if cabecera.headerd == dia{
			aux := cabecera.abajo
			for aux != nil {
				if aux.Dia == dia && aux.Categoria == categoria{
					return aux
				}
				aux = aux.abajo
			}
		}
		cabecera = cabecera.siguiente
	}
	return nil
}


func (m *matriz) ReturnDias () []ProductosPedido{
	var Lista []ProductosPedido
	cabecera := m.lst_h.first
	for cabecera != nil{
		aux := cabecera.abajo
		for aux != nil {
			Nuevo := ProductosPedido{
				Dia:       aux.Dia,
				Categoria: aux.Categoria,
			}
			Lista = append(Lista,Nuevo)
			aux = aux.abajo
		}
		cabecera = cabecera.siguiente
	}
	return Lista
}
