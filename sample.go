package main

import (
	"Dasar_Golang/anuan"
	"Dasar_Golang/initt"
	"fmt"
)

type Married struct {
	Nama string
	Benar bool
}


func (s Married) Ganti() {
	s.Nama = "Jakarta Selatan " + s.Nama
}

func main() {

	eko := Married{"Samuelers", true}
	eko.Ganti()

	fmt.Println(eko)

	x := anuan.Hailo("Tono")
	fmt.Println(x)

	initt.Test();

	
}
