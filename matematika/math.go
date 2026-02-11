package main

import "fmt"

func main() {
	penjumlahan := 2 + 3
	pengurangan := 2 - 3
	perkalian := 2 * 3
	
	var pembagian float64 = 2 / 3
	// kalo mau hasil yang diharapkan berarti harus pake float juga di pembagiannya
	var pembagianAsil = 2.0/3.0
	modulus := 6 % 3

	fmt.Println(penjumlahan, pengurangan, perkalian, pembagian, pembagianAsil, modulus)

	//assingment operator
	a := 2
	fmt.Println(a)
	a+= 2
	fmt.Println(a)
	a-= 2
	fmt.Println(a)
	a/= 2
	fmt.Println(a)
	a*= 2
	fmt.Println(a)
}