package Estructuras

import "fmt"

type NodoBTree struct {
	keys []*Usuario
	n int
	children []*NodoBTree
	t int
	leaf bool
}

type BTree struct {
	root *NodoBTree
	t int
}
//metodos de la clase Btree

func NewBTree() *BTree{
	return &BTree{
		root: nil,
		t:    5,
	}
}
func (t *BTree) Traverse(){
	if t.root != nil {
		t.root.traverse()
	}
	fmt.Println()
}

func (t *BTree) IsEmpty() bool{
	return t.root == nil
}

func (t *BTree) GetRoot() *NodoBTree {
	return t.root
}

func (t *BTree) Show() {
	if t.root != nil{
		t.root.show()
	}
}

func (t *BTree) searchTree(data Usuario) *NodoBTree {
	if t.root == nil {
		return nil
	} else {
		return t.root.search(data)
	}
}

func (t *BTree) Search(data Usuario) *Usuario{
	aux := t.searchTree(data)
	if aux != nil {
		for i := 0; i < aux.n; i++{
			if aux.keys[i].Dpi == data.Dpi {
				return aux.keys[i]
			}
		}
	}
	return nil
}

func (t *BTree) searchTreeCorreo(data Usuario) *NodoBTree {
	if t.root == nil {
		return nil
	} else {
		return t.root.searchCorreoNPassword(data)
	}
}

func (t *BTree) SearchCorreoNPass(data Usuario) *Usuario{
	aux := t.searchTreeCorreo(data)
	if aux != nil {
		for i := 0; i < aux.n; i++{
			if aux.keys[i].Correo == data.Correo && aux.keys[i].Password == data.Password {
				return aux.keys[i]
			}
		}
	}
	return nil
}


func (t *BTree) Insert(data Usuario) {
	if t.root == nil {
		t.root = newNodoBTree(t.t, true)
		t.root.keys[0] = &data
		t.root.n = 1
	} else {
		if t.root.n == 2 * t.t - 1 {
			aux := newNodoBTree(t.t, false)
			aux.children[0] = t.root
			aux.splitChild(0, t.root)

			i := 0
			if aux.keys[0].Dpi < data.Dpi {
				i++
			}
			aux.children[i].insertNonNull(data)
			t.root = aux
		} else {
			t.root.insertNonNull(data)
		}
	}
}

func (T *BTree) ImprimirArbol(){
	if T.root != nil {
		imprimirArbol(*T.root)
	}
}

func  imprimirArbol( Temp NodoBTree){
	if Temp.leaf{
		for i:=0; i< Temp.n;i++{
			fmt.Println(Temp.keys[i].Dpi)
		}

	}else{
		for i := 0;i < Temp.n;i++ {
			if !Temp.leaf{
				imprimirArbol(*Temp.children[i])
			}
		}
	}
}

func (T *BTree) SearchCorreo(usuario Usuario) *Usuario{
	if T.root != nil {
		return searchCorreo(*T.root, usuario)
	} else{
		return nil
	}
}

func  searchCorreo( Temp NodoBTree, usuario Usuario) *Usuario{
	if Temp.leaf{
		for i:=0; i< Temp.n;i++{
			if Temp.keys[i].Correo == usuario.Correo && Temp.keys[i].Password == usuario.Password {
				fmt.Println(Temp.keys[i].Dpi)
				return Temp.keys[i]
			}
		}

	}else{
		for i := 0;i < Temp.n;i++ {
			if !Temp.leaf{
				//imprimirArbol(*Temp.children[i])
				return searchCorreo(*Temp.children[i], usuario)
			}
		}
	}
	return nil
}


func (t *BTree) Remove(data Usuario) {
	if t.root == nil {
		fmt.Println("El arbol esta vacio")
	} else {
		t.root.remove(data)
		if t.root.n == 0 {
			//aux := t.root
			if t.root.leaf {
				t.root = nil
			} else {
				t.root = t.root.children[0]
			}
		}
	}
}

//Metodos de los nodos
func newNodoBTree(t int, leaf bool) *NodoBTree {
	var newkeys []*Usuario
	for i := 0; i< 2*t-1;i++{
		//var usuario Usuario
		newkeys = append(newkeys, nil)
	}
	var newChildren []*NodoBTree
	for i := 0; i< 2*t ;i++{
		newChildren = append(newChildren, nil)
	}
	return &NodoBTree{
		keys:     newkeys,
		n:        0,
		children: newChildren,
		t:        t,
		leaf:     leaf,
	}
}

func (t *NodoBTree) traverse(){
	i := 0
	for i = 0; i < t.n; i++ {
		if !t.leaf{
			t.children[i].traverse()
		}
		fmt.Println(t.keys[i])
	}
	if !t.leaf{
		t.children[i].traverse()
	}
}

func (t *NodoBTree)  show( ) {
	 i := 0
	for i = 0; i < t.n; i++ {

		if t.leaf == false {
			t.children[i].show()
		}
		//modelo.addRow(keys[i].toString().split("-"));
	}

	if t.leaf == false{
		t.children[i].show()
	}
}

func (t *NodoBTree) search( data Usuario) *NodoBTree {

	i := 0
	for i < t.n && data.Dpi > t.keys[i].Dpi {
		i++
	}

	if t.keys[i].Dpi ==  data.Dpi  {
		return t
	}

	if t.leaf == true {
		return nil
	}

	return t.children[i].search(data)

}

func (t *NodoBTree) searchCorreoNPassword( data Usuario) *NodoBTree {

	i := 0
	for i < t.n && data.Correo > t.keys[i].Correo {
		i++
	}

	if t.keys[i].Correo ==  data.Correo && t.keys[i].Password ==  data.Password  {
		return t
	}

	if t.leaf == true {
		return nil
	}

	return t.children[i].searchCorreoNPassword(data)

}

func (t *NodoBTree) insertNonNull(data Usuario) {
	i := t.n - 1
	if t.leaf == true {
		for i >= 0 && data.Dpi < t.keys[i].Dpi {
			t.keys[i + 1] = t.keys[i]
			i--
		}

		t.keys[i + 1] = &data
		t.n = t.n + 1
	} else {
		for i >= 0 && t.keys[i].Dpi > data.Dpi {
			i--
		}
		if t.children[i + 1].n == (2 * t.t - 1) {
			t.splitChild(i + 1, t.children[i + 1])
			if t.keys[i + 1].Dpi < data.Dpi {
				i++
			}
		}
		t.children[i + 1].insertNonNull(data)
	}
}


func (t *NodoBTree) splitChild( index int , node *NodoBTree) {
	aux := newNodoBTree(node.t, node.leaf)
	aux.n = t.t - 1

	for  j := 0; j < t.t - 1; j++{
		aux.keys[j] = node.keys[j + t.t]
	}
	if node.leaf == false{
		for j := 0; j < t.t; j++ {
			aux.children[j] = node.children[j + t.t]
		}
	}
	node.n = t.t - 1

	for j:= t.n; j >= (index + 1); j-- {
		t.children[j + 1] = t.children[j]
	}

	t.children[index + 1] = aux

	for j := (t.n - 1); j >= index; j-- {
		t.keys[j + 1] = t.keys[j]
	}

	t.keys[index] = node.keys[t.t - 1]

	t.n = t.n + 1
}

func (t *NodoBTree) findKey( data Usuario) int {
	index := 0
	for index < t.n && t.keys[index].Dpi < data.Dpi {
		index++
	}
	return index
}

func (t *NodoBTree) remove(data Usuario) {
	index := t.findKey(data);
	if index < t.n && t.keys[index].Dpi == data.Dpi {
		if (t.leaf) {
			t.removeFromLeaf(index);
		} else {
			t.removeFromNonLeaf(index);
		}
	} else {
		if (t.leaf) {
			fmt.Println("el dato " , data , "\nNo existe");
			return;
		}
		flag := index == t.n

		if t.children[index].n < t.t {
			t.fill(index);
		}
		if flag && index > t.n {
			t.children[index - 1].remove(data)
		} else {
			t.children[index].remove(data)
		}
	}
}

func (t *NodoBTree) removeFromLeaf(index int) {
	for i := index + 1; i < t.n; i++ {
		t.keys[i - 1] = t.keys[i]
	}
	t.n--
}

func (t *NodoBTree) removeFromNonLeaf(index int) {
	key := t.keys[index]

	if t.children[index].n >= t.t {
		pred := t.getPred(index)
		t.keys[index] = pred
		t.children[index].remove(*pred)
	} else if t.children[index + 1].n >= t.t {
		succ := t.getSucc(index)
		t.keys[index] = succ
		t.children[index + 1].remove(*succ)
	} else {
		t.merge(index)
		t.children[index].remove(*key)
	}
}

func (t *NodoBTree) getPred(index int) *Usuario{
	aux := t.children[index]
	for !aux.leaf {
		aux = aux.children[aux.n]
	}
	return aux.keys[aux.n - 1]
}

func (t *NodoBTree) getSucc(index int) *Usuario{
	aux := t.children[index + 1]
	for !aux.leaf {
		aux = aux.children[0]
	}
	return aux.keys[0]
}

func (t *NodoBTree) merge( index int) {
	child := t.children[index]
	sibling := t.children[index + 1]

	child.keys[t.t - 1] = t.keys[index]

	for i:= 0; i < sibling.n; i++ {
		child.keys[i + t.t] = sibling.keys[i]
	}

	if !child.leaf {
		for i:= 0; i <= sibling.n; i++ {
			child.children[i + t.t] = sibling.children[i]
		}
	}

	for i:= index + 1; i < t.n; i++ {
		t.keys[i - 1] = t.keys[i]
	}

	for i:= index + 2; i <= t.n; i++ {
		t.children[i - 1] = t.children[i]
	}
	child.n += sibling.n + 1
	t.n--
}

func (t *NodoBTree) fill(index int) {
	if index != 0 && t.children[index - 1].n >= t.t {
		t.borrowFromPrev(index)
	} else if index != t.n && t.children[index + 1].n >= t.t {
		t.borrowFromNext(index)
	} else {
		if index != t.n {
			t.merge(index)
		} else {
			t.merge(index - 1)
		}
	}
}

func (t *NodoBTree) borrowFromPrev(index int) {
	child := t.children[index]
	sibling := t.children[index - 1]

	for i:= child.n - 1; i >= 0; i-- {
		child.keys[i + 1] = child.keys[i]
	}

	if !child.leaf {
		for  i:= child.n; i >= 0; i-- {
			child.children[i + 1] = child.children[i]
		}
	}

	child.keys[0] = t.keys[index - 1]

	if !child.leaf{
		child.children[0] = sibling.children[sibling.n]
	}
	t.keys[index - 1] = sibling.keys[sibling.n - 1]
	child.n += 1
	sibling.n -= 1
}

func (t *NodoBTree) borrowFromNext(index int) {
	child := t.children[index]
	sibling := t.children[index + 1]
	child.keys[(child.n)] = t.keys[index]

	if !child.leaf {
		child.children[(child.n) + 1] = sibling.children[0]
	}

	t.keys[index] = sibling.keys[0]

	for  i:= 1; i < sibling.n; i++ {
		sibling.keys[i - 1] = sibling.keys[i]
	}

	if !sibling.leaf {
		for i:= 1; i <= sibling.n; i++ {
			sibling.children[i - 1] = sibling.children[i]
		}
	}

	child.n += 1
	sibling.n -= 1
}
