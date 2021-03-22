package Estructuras
import "fmt"

type Product struct {
	nombre string
	codigo int
	descripcion string
}

type NodoMatriz struct {
	// Estos atributos son especificos para la matriz
	x,y int //saber en que cabecera se esta

	producto *Product                             // objeto a guardar
	izquierdo, derecho, arriba, abajo *NodoMatriz // nodos con los que se desplaza dentro de la matriz
	//Estos atributos son especificos para la lista
	header int                      // tipo interno de caecera
	siguiente, anterior *NodoMatriz // nodos con los que se va a desplazar dentro de las listas
}

type lista struct {
	first, last *NodoMatriz
}

type matriz struct {
	lst_h, lst_y *lista
}

func newLista() *lista{
	return &lista{nil,nil}
}

func NewMatriz()*matriz{
	return &matriz{nil,nil}
}

func nodoMatriz(x int, y int, producto *Product) *NodoMatriz {
	return &NodoMatriz{x, y,producto, nil,nil,nil,nil,0,nil,nil}

}
func nodoLista(header int) *NodoMatriz {

	return &NodoMatriz{0,0,nil,nil,nil,nil,nil,header,nil,nil}
}
func (n *NodoMatriz) headerx() int  { return n.x }
func  (n *NodoMatriz) headery() int { return n.y }
func (n *NodoMatriz) toString() string { return "Nombre: " + n.producto.nombre + "\nDespcripcion: "+n.producto.descripcion}

func (l *lista) ordenar(nuevo *NodoMatriz) {
	aux:= l.first
	for(aux != nil){
		if nuevo.header > aux.header{
			aux = aux.siguiente
		}else{

		}
	}
}

func (l *lista) insert(header int)  {
	nuevo := nodoLista(header)
	if l.first == nil{
		l.first =nuevo
		l.last = nuevo
	}else {
		l.ordenar(nuevo)
	}
}

func (l *lista) search(header int) *NodoMatriz {
	temp := l.first
	for temp != nil{
		if temp.header == header{
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}
func (l *lista ) print() {
	temp := l.first
	for temp != nil {
		fmt.Print("Cabecera", temp.header)
		temp = temp.siguiente
	}
}

func(m *matriz) insert(proucto *Product, x int, y int){
	h := m.lst_h.search(x)
	v := m.lst_y.search(y)

	if h == nil && v == nil{
		m.NoExiste(proucto, x,y)
	} else if h == nil {

	} else if v == nil{

	}
}

func (m *matriz) NoExiste(producto *Product, x int, y int) {
	m.lst_h.insert(x)//se inserta en la lista que emula la cabecera horizontal
	m.lst_y.insert(y)// se inserta en la lista que emula la cabecera vertical


	//se busca el NodoMatriz que se acaba de insertar, para poder enlazarlos
	h:= m.lst_h.search(x)
	v := m.lst_y.search(y)

	nuevo := nodoMatriz(x,y,producto)//se crea un nuevo NodoMatriz
	h.abajo = nuevo// se enlaza el NodoMatriz horizontal hacia abajo
	nuevo.arriba = h// se enlaza el nuevo NodoMatriz hacia arriba
	v.derecho = nuevo// se enlaza el NodoMatriz vertical hacia la derecha
	nuevo.izquierdo = v// se enlaza el nuevo NodoMatriz hacia la izquierda



}