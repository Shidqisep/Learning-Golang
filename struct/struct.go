package main

import "fmt"

type Person struct {
	name, jobs, address string
	age        int
	anyType interface{}
}

type Place struct{
	name, location string
	isGood bool
	age int
}

//ini berarti func sayHellonya punya si Person
func (p Person) sayHello() {
	fmt.Println("Halo dari", p.name)
}

func (p Place)welcoming(name string)  {
	fmt.Println("Halo", name, "Selamat Datang di", p.name)
}

func (p Place)placeRating(){
	if p.isGood {
		fmt.Println(p.name,", Tempat ini bagus")
	}else{
		fmt.Println(p.name,", Tempat ini jelek")
	}
}

func main() {
	var per1 Person
	per1.name = "Shidqi"
	per1.age = 20

	fmt.Println(per1)
	per1.sayHello()

	per2 := Person{"Anton", "Mahasiswa", "Alamat", 21, 100}
	fmt.Println(per2)
	per2.sayHello()

	per3 :=Person{
		name: "Tono",
		address: "HBTB BLOK C",
	}
	fmt.Println(per3)
	per3.sayHello()

	pla := Place{
		name: "Rumah Tono",
		location: "Sukatani",
		isGood: false,
	}

	pla.placeRating()

	pla2 := Place{
		name: "Rumah Aku",
		location: "Tapos",
		isGood: true,
	}
	pla2.placeRating()
	pla2.welcoming(per1.name)
}