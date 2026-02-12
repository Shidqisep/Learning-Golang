package main

import "fmt"

func main() {
	// var angka int
	// sama dengan ==
	fmt.Println(5 == 5)
	fmt.Println(5 != 5)
	fmt.Println(5 >= 5)
	fmt.Println(5 <= 5)
	//harus make pointer
	// fmt.Scanln(&angka)
	// fmt.Println(angka)

	//operator logika 
	//and &&
	if (5 > 3 && 8 < 9) {
		fmt.Println("pernyataan diatas benar")
	}else{
		fmt.Println("Pernyataan diatas salah")
	}

	//or ||
		if (5 > 6 || 8 < 6) {
		fmt.Println("salah satu pernyataan diatas benar")
	}else{
		fmt.Println("Pernyataan diatas salah semua")
	}
}