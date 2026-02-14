package main

import "fmt"

type Cities struct {
	city, province, country string
}

func main() {
	cities1 := new(Cities)
	cities2 := cities1
	cities2.province = "Jakarta Raya"

	fmt.Println(cities1) //cities1 juga ikut berubah
	fmt.Println(cities2)
}