package main

import "fmt"

type Cities struct {
	city, province, country string
}

func changeCountryToIndonesia(cities *Cities) {
	cities.country = "Indonesia"
}

func main() {
	cities := Cities{"Jakarta Timur", "Jakarta", "Malaysia"}
	changeCountryToIndonesia(&cities) // ini artinya kita passing memory addressnya
	fmt.Println(cities) // berubah
}