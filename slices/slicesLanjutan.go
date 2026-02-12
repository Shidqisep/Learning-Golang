package main

import (
	"fmt"
)

func main() {
	//slice hanya jendela view ke array
	merekMobil := [5]string{"Bmw", "Mercedes", "Toyota", "Honda", "Suzuki"}
	sliceMobil := merekMobil[:]

	slice2 := make([]string,3)
	fmt.Println(slice2)
	// function copy ini masukiin sliceMobil ke slice2 sesuai kapasitas di slice 2
	jumlahYangTerCopy := copy(slice2,sliceMobil)
	fmt.Println(slice2)
	fmt.Println(jumlahYangTerCopy)

	mySlice := append(slice2, "Mitsubishi")
	fmt.Println(slice2)
	fmt.Println(mySlice)
}