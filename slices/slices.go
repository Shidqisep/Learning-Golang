package main

import "fmt"

func mains() {
	//slice hanya jendela view ke array
	merekMobil := [5]string{"Bmw", "Mercedes", "Toyota", "Honda", "Suzuki"}
	//ambil semua indeks dari array merekMobil
	sliceMobil := merekMobil[:]
	fmt.Println(sliceMobil)
	//ambil dari awal hingga indeks tertentu
	// note tidak termasuk indeks 3
	sliceMobil2 := merekMobil[:3]
	fmt.Println(sliceMobil2)
	//ambil dari indeks tertentu hingga akhir
	sliceMobil3 := merekMobil[2:]
	fmt.Println(sliceMobil3)
	//ambil dari indeks tertentu ke indeks tertentu
	// note tidak termasuk indeks 4
	sliceMobil4 := merekMobil[2:4]
	fmt.Println(sliceMobil4)
}