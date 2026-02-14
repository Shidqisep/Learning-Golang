package main

import "fmt"

type Cities struct{
	city, province, country string
}

func main() {
	cities1 := Cities{"Depok", "Jawa Barat", "Indonesia"}

	cities2 := &cities1 //pointer
	cities2.city = "Tangerang"

	//ini dia cities 1 ikut berubah
	fmt.Println(cities1)
	fmt.Println(*cities2)

	/*
	cities2 = &Cities{"Jakarta Pusat", "Jakarta", "Indonesia"}
	//cities 1 tidak ikut berubah menjadi jakarta pusat
	fmt.Println(cities1)
	fmt.Println(*cities2)
	*/

	//ini berarti mengubah value dari yang cities 2 tunjuk yang dimana ini merupakan cities 1
	*cities2 = Cities{"Jakarta Pusat", "Jakarta", "Indonesia"}
	fmt.Println(cities1)
	fmt.Println(*cities2)


}