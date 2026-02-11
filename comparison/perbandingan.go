package main

import "fmt"

func main() {
	var angka int
	// sama dengan ==
	fmt.Println(5 == 5)
	fmt.Println(5 != 5)
	fmt.Println(5 >= 5)
	fmt.Println(5 <= 5)
	//harus make pointer
	fmt.Scanln(&angka)
	fmt.Println(angka)
}