package Estructuras

type Nodo struct {
	 dato Tienda
	 next, previous *Nodo
}

func (n *Nodo) GetTienda() *Tienda {
	return &n.dato
}

type ListaDoble struct {
	head,last *Nodo
	size int
}

func NewListaDoble() *ListaDoble {
	return &ListaDoble{head: nil, last: nil, size: -1}
}
//saber si esta vacio
func (d *ListaDoble) IsEmpty() bool{
	if d.size == -1 {
		return true
	}
	return false
}

func (d *ListaDoble) Add(dato Tienda)  {
	nuevoNodo := &Nodo{dato,nil,nil}
	if d.IsEmpty(){
		d.head = nuevoNodo
		d.last = nuevoNodo
		d.size = 0
	}else {
		d.last.next = nuevoNodo
		nuevoNodo.previous = d.last
		d.last = nuevoNodo
		d.size ++
	}
}

func (d *ListaDoble) DeleteByName(nombre string) *Nodo{
	if d.IsEmpty(){
		println("No se puede eliminar nada, la lista esta vacia")
	}else{
		aux := d.head
		final := d.last
		if aux.dato.Nombre == nombre{
			if d.size == 0{
				d.head = aux.next
				d.last = aux.next
				d.size--
			}else {
				aux.next.previous = nil
				d.head = aux.next
				aux.next = nil// atencion aqui <---------------------------------
				d.size--
			}
			return aux
		}else if final.dato.Nombre == nombre{
			final.previous.next = nil
			d.last = final.previous
			final.previous = nil// atencion aqui <---------------------------------
			d.size--
			return final
		}else {
			for aux.next != nil {
				if aux.dato.Nombre == nombre {
					aux.previous.next = aux.next
					aux.next.previous = aux.previous
					d.size--
					return aux
				}
				aux = aux.next
			}
		}
	}
	println("No se encontro")
	return nil
}

func (d *ListaDoble) DeleteByIndex(index int) *Nodo{
	if d.IsEmpty(){
		println("No se puede eliminar nada, la lista esta vacia")
		return nil
	}else{
		aux := d.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux
	}
}

func (d *ListaDoble) Search(nombre string) int{
	contador:=0
	if !d.IsEmpty(){
		aux := d.head
		for aux!=nil{
			if nombre == aux.dato.Nombre{
				return contador
			}
			//fmt.Print(nombre)
			//fmt.Println(aux.dato.Nombre)
			//fmt.Println(aux.dato.Nombre, contador)
			aux = aux.next
			contador++
		}
	}
	return -1
}

func (d *ListaDoble) SearchNReturn(nombre string) Tienda{
	var coincidencia Tienda
	aux := d.head
	for d.head != nil {
		if  d.head.dato.Nombre == nombre {
			coincidencia = d.head.dato
		}
		d.head = d.head.next
	}
	d.head = aux
	return coincidencia
}
func (d *ListaDoble) SearchNReturnM(nombre string) *Tienda{
	var coincidencia *Tienda
	aux := d.head
	for d.head != nil {
		if  d.head.dato.Nombre == nombre {
			coincidencia = (&d.head.dato)
		}
		d.head = d.head.next
	}
	d.head = aux
	return coincidencia
}

func (d *ListaDoble) SearchIndex(index int) *Nodo {
	if d.IsEmpty(){
		println("No se puede encontrar nada, la lista esta vacia")
		return nil
	}else{
		aux := d.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux
	}
}

func (d *ListaDoble) ImprimirLista() {
	if d.IsEmpty(){
		println("-----------------------------------------------------------------------------------------")
		println("No se puede imprimir nada, la lista esta vacia")
		println("-----------------------------------------------------------------------------------------")
	}else{
		aux := d.head
		println("-----------------------------------------------------------------------------------------")
		for aux != nil{

			println(aux.dato.Nombre, " ", aux.dato.Calificacion, " ",aux.dato.Descripcion ," ",aux.dato.Contacto)

			aux = aux.next
		}
		println("-----------------------------------------------------------------------------------------")
	}

}

func (d *ListaDoble) ReturnNodes() []Tienda {
	var ArregloTiendas []Tienda
	aux := d.head
	for aux != nil{
		ArregloTiendas = append(ArregloTiendas, aux.dato)
		aux = aux.next
	}
	return ArregloTiendas
}
