package Estructuras

import "fmt"



type NodoArbolPedidos struct {
	Anio int
	Mes  ListaDobleMes

	altura int
	Hizq* NodoArbolPedidos
	Hder* NodoArbolPedidos
}

func newNodoArbolPedidos(anio int) *NodoArbolPedidos {
	return &NodoArbolPedidos{anio, *NewListaDoblePedido(),0,nil, nil}
}



type AvlPedidos struct {
	Raiz *NodoArbolPedidos
}

func newAVLPedidos() *AvlPedidos{
	return &AvlPedidos{nil}
}

func maxp(val1 int, val2 int) int {
	if val1>val2{
		return val1
	}
	return val2
}
func alturap(temp *NodoArbolPedidos) int{
	if temp != nil{
		return temp.altura
	}
	return -1
}
func rotacionIzquierdap(temp **NodoArbolPedidos){
	aux := (*temp).Hizq
	(*temp).Hizq = aux.Hder
	aux.Hder = *temp
	(*temp).altura = maxp(alturap((*temp).Hder), alturap((*temp).Hizq)) + 1
	aux.altura = maxp(alturap((*temp).Hizq), (*temp).altura) + 1
	*temp = aux
}
func rotacionDerechap(temp **NodoArbolPedidos) {
	aux := (*temp).Hder
	(*temp).Hder = aux.Hizq
	aux.Hizq = *temp
	(*temp).altura = maxp(alturap((*temp).Hder), alturap((*temp).Hizq)) + 1
	aux.altura = maxp(alturap((*temp).Hder), (*temp).altura) + 1
	*temp = aux
}
func rotacionDobleIzquierdap(temp **NodoArbolPedidos) {
	rotacionDerechap(&(*temp).Hizq)
	rotacionIzquierdap(temp)
}

func rotacionDobleDerechap(temp **NodoArbolPedidos) {
	rotacionIzquierdap(&(*temp).Hder)
	rotacionDerechap(temp)
}

func insertp(anio int, root **NodoArbolPedidos) {
	if (*root) == nil {
		*root = newNodoArbolPedidos(anio)
		return
	}
	if anio < (*root).Anio{
		insertp(anio, &(*root).Hizq)
		if (alturap((*root).Hizq) - alturap((*root).Hder)) == -2 {
			if anio < (*root).Hizq.Anio{
				rotacionIzquierdap(root)
			}else{
				rotacionDobleIzquierdap(root)
			}
		}
	}else if anio > (*root).Anio{
		insertp(anio, &(*root).Hder)
		if (alturap((*root).Hder) - alturap((*root).Hizq)) == 2{
			if anio > (*root).Hder.Anio {
				rotacionDerechap(root)
			}else{
				rotacionDobleDerechap(root)
			}
		}
	}else{
		fmt.Println("Ya se inserto el indice")
	}

	(*root).altura = maxp(alturap((*root).Hizq), alturap((*root).Hder))+1
}

func (avl *AvlPedidos) Insertar(anio int) {
	insertp(anio, &avl.Raiz)
}

func (avl *AvlPedidos) SearchAnioExists(anio int) bool {
	return searchMesExists(anio, avl.Raiz)
}

func searchMesExists(anio int, temp *NodoArbolPedidos) bool {
	for temp!=nil {
		if anio==temp.Anio{
			return true
		}else if anio>temp.Anio{
			temp = temp.Hder
		}else if anio<temp.Anio{
			temp = temp.Hizq
		}
	}
	return false
}

func (avl *AvlPedidos) ReturnAnioNode(anio int) *NodoArbolPedidos {
	return returnMesNode(anio, avl.Raiz)
}
func returnMesNode(anio int, temp *NodoArbolPedidos) *NodoArbolPedidos{
	for temp!=nil {
		if anio==temp.Anio{
			return temp
		}else if anio>temp.Anio{
			temp = temp.Hder
		}else if anio<temp.Anio{
			temp = temp.Hizq
		}
	}
	return nil
}
func (alv *AvlPedidos) ReturnProductsOfTree() *[]int {
	var ArregloProductos []int
	recorrerArbolParaListaP(&(ArregloProductos),alv.Raiz)
	return &ArregloProductos
}

func recorrerArbolParaListaP(ArregloProductos *[]int, root *NodoArbolPedidos){
	if(root != nil) {
		*ArregloProductos = append(*ArregloProductos, root.Anio)
		recorrerArbolParaListaP(ArregloProductos, root.Hizq)
		recorrerArbolParaListaP(ArregloProductos, root.Hder)
	}
}