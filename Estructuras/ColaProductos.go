package Estructuras

type NodoCola struct {
	producto Productos
	next, previous *NodoCola
}


type ColaPedidos struct {
	head,last *NodoCola
	size int
}

func NewColaPedidos() *ColaPedidos {
	return &ColaPedidos{head: nil, last: nil, size: -1}
}
//saber si esta vacio
func (d *ColaPedidos) IsEmpty() bool{
	if d.size == -1 {
		return true
	}
	return false
}

func (d *ColaPedidos) Add(dato Productos)  {
	nuevoNodo := &NodoCola{dato,nil,nil}
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

func (d *ColaPedidos) DeleteByName(codigoPro int) *NodoCola {
	if d.IsEmpty(){
		println("No se puede eliminar nada, la lista esta vacia")
	}else{
		aux := d.head
		final := d.last
		if aux.producto.Codigo == codigoPro{
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
		}else if final.producto.Codigo == codigoPro{
			final.previous.next = nil
			d.last = final.previous
			final.previous = nil// atencion aqui <---------------------------------
			d.size--
			return final
		}else {
			for aux.next != nil {
				if aux.producto.Codigo == codigoPro {
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

func (d *ColaPedidos) DeleteByIndex(index int) *NodoCola {
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

func (d *ColaPedidos) Search(mes int) int{
	contador:=0
	if !d.IsEmpty(){
		aux := d.head
		for aux!=nil{
			if mes == aux.producto.Codigo{
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

func (d *ColaPedidos) SearchNReturn(mes int) Productos{
	var coincidencia Productos
	aux := d.head
	for d.head != nil {
		if  d.head.producto.Codigo == mes {
			coincidencia = d.head.producto
		}
		d.head = d.head.next
	}
	d.head = aux
	return coincidencia
}
func (d *ColaPedidos) SearchNReturnM(productocode int) *NodoCola {
	aux := d.head
	for d.head != nil {
		if  d.head.producto.Codigo == productocode {
			return d.head
		}
		d.head = d.head.next
	}
	d.head = aux
	return nil
}

func (d *ColaPedidos) SearchIndex(index int) *NodoCola {
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

func (d *ColaPedidos) ImprimirLista() {
	if d.IsEmpty(){
		println("-----------------------------------------------------------------------------------------")
		println("No se puede imprimir nada, la lista esta vacia")
		println("-----------------------------------------------------------------------------------------")
	}else{
		aux := d.head
		println("-----------------------------------------------------------------------------------------")
		for aux != nil{

			//println(aux.dato.Nombre, " ", aux.dato.Calificacion, " ",aux.dato.Descripcion ," ",aux.dato.Contacto)

			aux = aux.next
		}
		println("-----------------------------------------------------------------------------------------")
	}
}

func (d *ColaPedidos) ReturnNodes() []Productos {
	var ArregloTiendas []Productos
	aux := d.head
	for aux != nil{
		ArregloTiendas = append(ArregloTiendas, aux.producto)
		aux = aux.next
	}
	return ArregloTiendas
}