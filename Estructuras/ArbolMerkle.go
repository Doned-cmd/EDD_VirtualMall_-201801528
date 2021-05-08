package Estructuras

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"../EncriptacionFernet"
	"os"
	"os/exec"
	"strconv"
)

type NodeMerkle struct{
	HASH,HASHDER,HASHIZQ string
	indice int

	altura int
	Hizq* NodeMerkle
	Hder* NodeMerkle

	datos string
}

func newMerkleNode(indice int) *NodeMerkle {
	return &NodeMerkle{"", "","",indice,0,nil, nil,""}
}

type Merkle struct {
	Raiz *NodeMerkle
}

func NewMerkle(transacciones []string) *Merkle{
	nuevo := &Merkle{nil}
	//fmt.Println(calcularlongitudparahash(transacciones), "Este es el numero de nodos ")
	for i := 0; i < calcularlongitudparahash(transacciones); i++ {
		//print(i)
		nuevo.InsertarMerk(i)
	}

	insertarDatosUltimoNivel(&transacciones, nuevo.Raiz)
	return nuevo
}

func calcularlongitudparahash(transacciones []string) int{
	cantidadniveles := math.Trunc(math.Log2(float64(len(transacciones))))
	//fmt.Println(cantidadniveles+1)
	devolver := 0
	for i := 0; i <= int(cantidadniveles)+1; i++ {
		p := float64(i)
		devolver += int(math.Pow(2,p))
	}
	return devolver
}

func insertarDatosUltimoNivel(transacciones *[]string, temp *NodeMerkle){
	if temp != nil{
		insertarDatosUltimoNivel(transacciones,temp.Hizq)
		insertarDatosUltimoNivel(transacciones,temp.Hder)

		if temp.Hizq == nil && temp.Hder == nil{
			if len(*transacciones) != 0 {
				temp.datos = (*transacciones)[0]
				temp.HASH = EncriptacionFernet.Encriptar256((*transacciones)[0])
				*transacciones = append((*transacciones)[:0], (*transacciones)[0+1:]...) //cola
				//fmt.Println("tamano de coal", len(*transacciones))
			}else{
				temp.datos = ""
				temp.HASH = ""
			}
		}else{
			temp.HASHIZQ = temp.Hizq.HASH
			temp.HASHDER = temp.Hder.HASH
			temp.HASH = EncriptacionFernet.Encriptar256(temp.HASHIZQ+temp.HASHDER)
		}
	}

}


func (n NodeMerkle) graficar(i int) string {
	var texto string

	var informacion string
	informacion += n.HASH + "\\n"
	if n.Hizq == nil && n.Hder == nil{

		informacion += n.datos + "\\n"
	} else {
		informacion += n.HASHDER
		informacion += n.HASHIZQ + "\\n"
	}

	texto+= "nodo" + strconv.Itoa(i) + "[label=\""+informacion+"\"];\n"

	return  texto
}

func (a *Merkle) Graficar() string {

	texto := "digraph g{\n    \n\tnode [style=\"filled\" shape=\"rectangle\" fillcolor=\"#ff00005f\"];\n"
	texto += a.graficar(a.Raiz)

	texto+= "\n}"

	return  texto
}

func (a Merkle) graficar(actual *NodeMerkle) string {
	var texto string

	if actual != nil{

		texto += actual.graficar(actual.indice)
		texto += a.graficar(actual.Hizq)
		texto += a.graficar(actual.Hder)

		if actual.Hizq != nil{
			texto+= "nodo" + strconv.Itoa(actual.indice) + " -> nodo" + strconv.Itoa(actual.Hizq.indice) + ";\n"
		}
		if actual.Hder != nil{
			texto+= "nodo" + strconv.Itoa(actual.indice) + " -> nodo" + strconv.Itoa(actual.Hder.indice) + ";\n"
		}
	}

	return  texto
}

func escribirDOT1(text string, nombreArchivo string) {
	f, err := os.Create("D:\\Escritorio\\USAC\\EDD\\Proyecto\\Reportes" + "/"+nombreArchivo+".dot")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(text)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("$$$ Archivo dot escrito.")
	ejecutarComand1(nombreArchivo)
}

func ejecutarComand1(nombreArchivo string) {
	path,_ := exec.LookPath("dot")
	cmd,_ := exec.Command(path, "-Tsvg", "D:\\Escritorio\\USAC\\EDD\\Proyecto\\Reportes" + "/"+nombreArchivo+".dot").Output()
	mode := int(0777)
	ioutil.WriteFile("D:\\Escritorio\\USAC\\EDD\\Proyecto\\Reportes" + "/"+nombreArchivo+".svg", cmd, os.FileMode(mode))

	fmt.Println("$$$ Reporte Arbol Usuarios png completado")
}




func maxMerk(val1 int, val2 int) int {
	if val1>val2{
		return val1
	}
	return val2
}
func alturaMerk(temp *NodeMerkle) int{
	if temp != nil{
		return temp.altura
	}
	return -1
}
func rotacionIzquierdaMerk(temp **NodeMerkle){
	aux := (*temp).Hizq
	(*temp).Hizq = aux.Hder
	aux.Hder = *temp
	(*temp).altura = maxMerk(alturaMerk((*temp).Hder), alturaMerk((*temp).Hizq)) + 1
	aux.altura = maxMerk(alturaMerk((*temp).Hizq), (*temp).altura) + 1
	*temp = aux
}
func rotacionDerechaMerk(temp **NodeMerkle) {
	aux := (*temp).Hder
	(*temp).Hder = aux.Hizq
	aux.Hizq = *temp
	(*temp).altura = maxMerk(alturaMerk((*temp).Hder), alturaMerk((*temp).Hizq)) + 1
	aux.altura = maxMerk(alturaMerk((*temp).Hder), (*temp).altura) + 1
	*temp = aux
}
func rotacionDobleIzquierdaMerk(temp **NodeMerkle) {
	rotacionDerechaMerk(&(*temp).Hizq)
	rotacionIzquierdaMerk(temp)
}

func rotacionDobleDerechaMerk(temp **NodeMerkle) {
	rotacionIzquierdaMerk(&(*temp).Hder)
	rotacionDerechaMerk(temp)
}
func insertMerk( indice int,root **NodeMerkle) {
	if (*root) == nil {
		*root = newMerkleNode(indice)
		return
	}
	if indice < (*root).indice {
		insertMerk( indice,&(*root).Hizq)
		if (alturaMerk((*root).Hizq) - alturaMerk((*root).Hder)) == -2 {
			if indice < (*root).Hizq.indice{
				rotacionIzquierdaMerk(root)
			}else{
				rotacionDobleIzquierdaMerk(root)
			}
		}
	}else if indice > (*root).indice{
		insertMerk( indice,&(*root).Hder)
		if (alturaMerk((*root).Hder) - alturaMerk((*root).Hizq)) == 2{
			if indice > (*root).Hder.indice {
				rotacionDerechaMerk(root)
			}else{
				rotacionDobleDerechaMerk(root)
			}
		}
	}else{
		fmt.Println("Ya se inserto el indice")
	}

	(*root).altura = maxMerk(alturaMerk((*root).Hizq), alturaMerk((*root).Hder))+1
}

func (avl *Merkle) InsertarMerk(indice int) {
	insertMerk(indice, &avl.Raiz)
}

func (avl *Avl) SearchProductExistsMerk(producto Producto) bool {
	return searchProductExists(producto, avl.Raiz)
}







