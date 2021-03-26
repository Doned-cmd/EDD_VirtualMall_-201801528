package Estructuras

type NodoMes struct {
	mes int
	matriz matriz
	next, previous *NodoMes
}
func (n *NodoMes) GetMatriz() *matriz {
	return &n.matriz
}


type ListaDobleMes struct {
	head,last *NodoMes
	size int
}

func NewListaDoblePedido() *ListaDobleMes {
	return &ListaDobleMes{head: nil, last: nil, size: -1}
}
//saber si esta vacio
func (d *ListaDobleMes) IsEmpty() bool{
	if d.size == -1 {
		return true
	}
	return false
}

func (d *ListaDobleMes) Add(dato int)  {
	nuevoNodo := &NodoMes{dato, *NewMatriz(),nil,nil}
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

func (d *ListaDobleMes) DeleteByName(mes int) *NodoMes {
	if d.IsEmpty(){
		println("No se puede eliminar nada, la lista esta vacia")
	}else{
		aux := d.head
		final := d.last
		if aux.mes == mes{
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
		}else if final.mes == mes{
			final.previous.next = nil
			d.last = final.previous
			final.previous = nil// atencion aqui <---------------------------------
			d.size--
			return final
		}else {
			for aux.next != nil {
				if aux.mes == mes {
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

func (d *ListaDobleMes) DeleteByIndex(index int) *NodoMes {
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

func (d *ListaDobleMes) Search(mes int) int{
	contador:=0
	if !d.IsEmpty(){
		aux := d.head
		for aux!=nil{
			if mes == aux.mes{
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

func (d *ListaDobleMes) SearchNReturn(mes int) matriz{
	var coincidencia matriz
	aux := d.head
	for d.head != nil {
		if  d.head.mes == mes {
			coincidencia = d.head.matriz
		}
		d.head = d.head.next
	}
	d.head = aux
	return coincidencia
}
func (d *ListaDobleMes) SearchNReturnM(mes int) *NodoMes {
	aux := d.head
	for d.head != nil {
		if  d.head.mes == mes {
			return d.head
		}
		d.head = d.head.next
	}
	d.head = aux
	return nil
}

func (d *ListaDobleMes) SearchIndex(index int) *NodoMes {
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

func (d *ListaDobleMes) ImprimirLista() {
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

func (d *ListaDobleMes) ReturnNodes() []matriz {
	var ArregloTiendas []matriz
	aux := d.head
	for aux != nil{
		ArregloTiendas = append(ArregloTiendas, aux.matriz)
		aux = aux.next
	}
	return ArregloTiendas
}

func (d *ListaDobleMes) ReturnMeses() []int {
	var ArregloTiendas []int
	aux := d.head
	for aux != nil{
		ArregloTiendas = append(ArregloTiendas, aux.mes)
		aux = aux.next
	}
	return ArregloTiendas
}