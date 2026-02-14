package helper

import "fmt"

var version = "1.0.0.0"
var Application = "Golang-Ganteng"

// kalo nama functionnya dari huruf besar artinya bisa diakses di package lain
// jika namanya berawalan dari huruf kecil maka dia tidak bisa di akses di package lain
func SayHello() string {
	return "Hello World"
}

func SayMwhahaha() {
	fmt.Println("mwhahaha")
}