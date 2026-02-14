package main

import "fmt"

type Man struct{
	name string
}

func (m *Man) Married() {
	m.name = "Mr. " + m.name
}

func main() {
	man1 := Man{}
	man1.name = "Tono"
	man1.Married()
	fmt.Println(man1)
}