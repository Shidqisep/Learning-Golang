package main

import "fmt"

type Cities struct{
	city, province, country string
}

func main() {
	cities1 := Cities{"Depok", "Jawa Barat", "Indonesia"}
	//ini artinya cities2 mengopy semua data cities 1
	cities2 := cities1
	cities2.city = "Bandung"

	//ini dia cities 1 tidak berubah
	fmt.Println(cities1)
	fmt.Println(cities2)

	cities3 := &cities1 //pointer
	cities3.city = "Tangerang"

	//ini dia cities 1 ikut berubah
	fmt.Println(cities1)
	fmt.Println(*cities3)
}