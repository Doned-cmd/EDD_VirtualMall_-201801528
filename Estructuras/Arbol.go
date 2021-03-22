package Estructuras

import "fmt"

type NodoArbol struct{
	producto Producto

	altura int
	Hizq* NodoArbol
	Hder* NodoArbol
}

func newNodoArbol(producto Producto) *NodoArbol {
	return &NodoArbol{producto, 0,nil, nil}
}

type Avl struct {
	Raiz *NodoArbol
}
func newAVL() *Avl{
	return &Avl{nil}
}

func max(val1 int, val2 int) int {
	if val1>val2{
		return val1
	}
	return val2
}
func altura(temp *NodoArbol) int{
	if temp != nil{
		return temp.altura
	}
	return -1
}
func rotacionIzquierda(temp **NodoArbol){
	aux := (*temp).Hizq
	(*temp).Hizq = aux.Hder
	aux.Hder = *temp
	(*temp).altura = max(altura((*temp).Hder), altura((*temp).Hizq)) + 1
	aux.altura = max(altura((*temp).Hizq), (*temp).altura) + 1
	*temp = aux
}
func rotacionDerecha(temp **NodoArbol) {
	aux := (*temp).Hder
	(*temp).Hder = aux.Hizq
	aux.Hizq = *temp
	(*temp).altura = max(altura((*temp).Hder), altura((*temp).Hizq)) + 1
	aux.altura = max(altura((*temp).Hder), (*temp).altura) + 1
	*temp = aux
}
func rotacionDobleIzquierda(temp **NodoArbol) {
	rotacionDerecha(&(*temp).Hizq)
	rotacionIzquierda(temp)
}

func rotacionDobleDerecha(temp **NodoArbol) {
	rotacionIzquierda(&(*temp).Hder)
	rotacionDerecha(temp)
}
func insert(producto Producto, root **NodoArbol) {
	if (*root) == nil {
		*root = newNodoArbol(producto)
		return
	}
	if producto.Codigo < (*root).producto.Codigo {
		insert(producto, &(*root).Hizq)
		if (altura((*root).Hizq) - altura((*root).Hder)) == -2 {
			if producto.Codigo < (*root).Hizq.producto.Codigo{
				rotacionIzquierda(root)
			}else{
				rotacionDobleIzquierda(root)
			}
		}
	}else if producto.Codigo > (*root).producto.Codigo{
		insert(producto, &(*root).Hder)
		if (altura((*root).Hder) - altura((*root).Hizq)) == 2{
			if producto.Codigo > (*root).Hder.producto.Codigo {
				rotacionDerecha(root)
			}else{
				rotacionDobleDerecha(root)
			}
		}
	}else{
		fmt.Println("Ya se inserto el indice")
	}

	(*root).altura = max(altura((*root).Hizq), altura((*root).Hder))+1
}

func (avl *Avl) Insertar(producto Producto) {
	insert(producto, &avl.Raiz)
}

func (avl *Avl) SearchProductExists(producto Producto) bool {
	return searchProductExists(producto, avl.Raiz)
}

func searchProductExists(producto Producto, temp *NodoArbol) bool {
	for temp!=nil {
		if producto.Codigo==temp.producto.Codigo{
			return true
		}else if producto.Codigo>temp.producto.Codigo{
			temp = temp.Hder
		}else if producto.Codigo<temp.producto.Codigo{
			temp = temp.Hizq
		}
	}
	return false
}

func(avl *Avl) SumarInventario( producto Producto){
	recorrersumando(producto, &avl.Raiz)
}

func recorrersumando(producto Producto, temp **NodoArbol){
	for temp!=nil {
		if producto.Codigo==(*temp).producto.Codigo{
			(*temp).producto.Cantidad = (*temp).producto.Cantidad+producto.Cantidad
		}else if producto.Codigo>(*temp).producto.Codigo{
			*temp = (*temp).Hder
		}else if producto.Codigo<(*temp).producto.Codigo{
			*temp = (*temp).Hizq
		}
	}
}
func (alv *Avl) ReturnProductsOfTree() *[]Producto {
	var ArregloProductos []Producto
	recorrerArbolParaLista(&(ArregloProductos),alv.Raiz)
	return &ArregloProductos
}

func recorrerArbolParaLista(ArregloProductos *[]Producto, root *NodoArbol){
	if(root != nil) {
		*ArregloProductos = append(*ArregloProductos, root.producto)
		recorrerArbolParaLista(ArregloProductos, root.Hizq)
		recorrerArbolParaLista(ArregloProductos, root.Hder)
	}
}

