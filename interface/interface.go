package main

import "fmt"

type hasName interface {
	getName() string
}

//parameter say hello ini tipenya interface
func sayHello(has hasName) {
	fmt.Println("Halo",has.getName())
}

type Person struct{
	name string
}

type Animal struct{
	name string
}

//Person ini secara umum adalah implementasi dari interface hasName
func (person Person)getName()string  {
	return person.name
}

func (animal Animal)getName()string  {
	return animal.name
}

//any dan interface itu bisa return tipe data apa saja
func ups(param any)any{
	return param
}



func main()  {
	person := Person{"Anton"}
	sayHello(person)
	anim := Animal{"Kucing"}
	sayHello(anim)
}