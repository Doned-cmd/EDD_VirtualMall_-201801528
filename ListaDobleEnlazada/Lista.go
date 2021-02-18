package ListaDobleEnlazada

import (
	"../Json"
)

type Nodo struct {
	 dato Json.Tienda
	 next, previous *Nodo
}

type ListaDoble struct {
	head,last *Nodo
	size int
}

func NewListaDoble() ListaDoble {
	return ListaDoble{head: nil, last: nil, size: -1}
}
//saber si esta vacio
func (d *ListaDoble) IsEmpty() bool{
	if d.size == -1 {
		return true
	}
	return false
}

func (d *ListaDoble) Add(dato Json.Tienda)  {
	nuevoNodo := &Nodo{dato,nil,nil}
	if d.IsEmpty(){
		d.head = nuevoNodo
		d.last = nuevoNodo
		d.size = 0
	}else {
		d.last.next = nuevoNodo
		nuevoNodo.previous = d.last
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
	}else{
		aux := d.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
	}
	return &Nodo{}
}

func (d *ListaDoble) Search(nombre string) int{
	var contador int
	if d.IsEmpty(){
	}else{
		aux := d.head
		for d.head.next != nil {
			if  d.head.dato.Nombre == nombre {
				break
			}
			d.head = d.head.next
			contador++
		}
		return contador
		d.head = aux
	}
	return -1
}

func (d *ListaDoble) SearchIndex(index int) *Nodo {
	if d.IsEmpty(){
		println("No se puede encontrar nada, la lista esta vacia")
	}else{
		aux := d.head
		for i := 0; i < index; i++ {
			aux = aux.next
		}
		return aux
	}
	return nil
}

func (d *ListaDoble) ImprimirLista() {
	if d.IsEmpty(){
		println("No se puede imprimir nada, la lista esta vacia")
	}else{
		aux := d.head
		for aux != nil{
			println(aux.dato.Nombre, " ", aux.dato.Calificacion, " ",aux.dato.Descripcion ," ",aux.dato.Contacto)
			aux = aux.next
		}
	}

}