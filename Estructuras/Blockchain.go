package Estructuras

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"../EncriptacionFernet"
	"time"
)

type BlockChain struct {

}

type Bloque struct {
	Indice int
	Fecha string
	Data string
	Nonce int
	ProviousHash string
	Hash string
}

func NewBloque(indice, nonce int,  data, previous, hash string) *Bloque {

	today := time.Now()
	dia := today.Day()
	var mes int
	anio := today.Year()
	var hora int
	hora = today.Hour()
	var minuto int
	minuto =  today.Minute()
	var segundo int
	segundo = today.Second()
	if today.Month().String() == "March"{
		mes = 3
	}else if today.Month().String() == "April"{
		mes = 4
	}else if today.Month().String() == "May"{
		mes = 5
	}

	return &Bloque{
		Indice:       indice,
		Fecha:        strconv.Itoa(dia) + "-"+strconv.Itoa(mes)+"-"+strconv.Itoa(anio)+"::"+strconv.Itoa(hora)+":"+strconv.Itoa(minuto)+":"+strconv.Itoa(segundo),
		Data:         data,
		Nonce:        nonce,
		ProviousHash: previous,
		Hash:         hash,
	}
}

func(b *Bloque) crearHash(){
	var texto string
	var hash string
	texto = strconv.Itoa(b.Indice) + b.Fecha + b.ProviousHash + b.Data + strconv.Itoa(b.Nonce)

	hash =  EncriptacionFernet.Encriptar256(texto)
	b.Nonce++
	fmt.Print(hash)

	b.Hash = hash
}


func(b *Bloque) guardarBloque() {
	f, err := os.Create("ArchivosBlock\\"+strconv.Itoa(b.Indice)+".txt")
	today := time.Now()
	dia := today.Day()
	var mes int
	anio := today.Year()
	var hora int
	hora = today.Hour()
	var minuto int
	minuto =  today.Minute()
	var segundo int
	segundo = today.Second()
	if today.Month().String() == "March"{
		mes = 3
	}else if today.Month().String() == "April"{
		mes = 4
	}else if today.Month().String() == "May"{
		mes = 5
	}

	text := strconv.Itoa(b.Indice) + "\n" + strconv.Itoa(dia) + "-"+strconv.Itoa(mes)+"-"+strconv.Itoa(anio)+"::"+strconv.Itoa(hora)+":"+strconv.Itoa(minuto)+":"+strconv.Itoa(segundo) + "\n" + b.Data + "\n" +strconv.Itoa(b.Nonce) + "\n"+  b.ProviousHash + "\n" + b.Hash
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(text)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("$$$ Bloque guardado.")
}