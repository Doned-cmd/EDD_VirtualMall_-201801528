package EncriptacionFernet

import (
	"crypto/sha256"
	"fmt"
	"github.com/fernet/fernet-go"

	//"github.com/fernet/fernet-go"
)

func EncriptarString(mensaje string, key string) (string){

	//llave2:=fernet.Key{}
	KeyMaster,_ := fernet.DecodeKey(key)


	bytes := []byte(mensaje)
	tok,_ := fernet.EncryptAndSign(bytes,KeyMaster)

	//decript := fernet.VerifyAndDecrypt(encript,100,)
	str := string(tok)
	//fmt.Println(KeyMaster.Encode())
	//fmt.Println(str)
	return str
}

func Desencriptar(token string, key string){
	llave,_ := fernet.DecodeKeys(key)
	b := []byte(token)
	textoB := fernet.VerifyAndDecrypt(b, 100000000000, llave)
	texto := string(textoB)
	fmt.Println(texto)
}

func EncriptarInt(mensaje string){

}

func Encriptar256(texto string) string{
	sum := fmt.Sprintf("%x", sha256.Sum256([]byte(texto)))
	return sum
}