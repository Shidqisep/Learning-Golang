package main

import (
	"fmt"
	"math/rand"
	_ "strconv"
	"strings"
)

func main() {
	// _nama := "Diandra"
	// sayHai("Shidqi")
	// fmt.Println(sayHello("Thierry", "Yudha"))
	manyParam("Shidqi", "Anton", "Tono", 10, 500)
	// sayHai(strconv.Itoa(returnInt()))
	// fmt.Println(returnBanyakData())

	//yang bool tidak diambil
	// hasil1, hasil2, hasil3, _ := returnBanyakData()
	// fmt.Println(hasil1, hasil2, hasil3,)
	// fmt.Println(getCompleteName())
	// defaultValue("", "", 0)
	fmt.Println(sumAll(5,4,5,1))
	//misalkan juga ada slice 
	nomor := []int{10,10,10,10}
	//harus pake titik titik kalo mau slice as a argument
	fmt.Println(sumAll(nomor...))

	//function juga bisa digunakan jadi variabel
	// notes: jangan pake kurung buka kurung tutup karena akan memanggil function
	sayHey := sayHello
	fmt.Println(sayHey("Shidqi", "Tono"))

	//function as a parameter
	sayHelloWithFilter("Anjing Lu", filtering)
	
}

func sayHai(name string){
	fmt.Println("Hello", name)
}

func sayHello(firstName string, lastName string) string {
	hello := "hello " + firstName + " " + lastName
	return (hello)	
}

//ini namanya variadic function, parameter a ini di treat/dianggap semacam slice
func manyParam(a ...any){
	fmt.Println(a[1], a[2])
	fmt.Printf("%T\n", a)
}

//bisa juga seperti ini
func sumAll(numbers ...int)int{
	total := 0
	for _, val := range numbers{
		total += val
	}
	return total
}

func returnInt()int{
	return (rand.Intn(10))
}

func returnBanyakData()(string, string, int, bool){
	return "Hello", "Fungsi ini mengembalikan banyak tipe", 100, true
}

func defaultValue(firstName, lastName string, umur int)() {
	if lastName == "" || firstName == "" {
		panic("WOI NAMANYA KOSONG")
	}else if umur == 0 {
		panic("WOI UMURNYA KOSONG")
	}
	fmt.Println("Sukses")
}

func getCompleteName()(firstName, lastName string, umur int){
	firstName = "Joni"
	lastName = "Saputra"
	umur = 100
	return firstName, lastName, umur
}

func filtering(name string)string{
	return strings.ReplaceAll(name, "Anjing", "***")
}

func sayHelloWithFilter(words string, filter func(string)string)  {
	filteredName := filter(words)
	fmt.Println("Hello " + filteredName)
}

