package main

import "fmt"

func main() {
	x := 20
	if v := x*2; v > 5 {
		fmt.Println(v, "Lebih besar dari 5")
	}

	switch x {
	case 10:
		fmt.Println(x, "adalah 10")
	case 15:
		fmt.Println(x, "adalah 15")
	default:
		fmt.Println("Gaada yang cocok kocak")
	}

	hari := "Kamis"

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari Kerja")
	case "Sabtu", "minggu":
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari tidak dikenal")
	}

	switch panjang := len(hari); {
	case panjang <= 5:
		fmt.Println("Dibawah 5")
	default:
		fmt.Println("Diatas 5 Panjangnya")
	}

	
}