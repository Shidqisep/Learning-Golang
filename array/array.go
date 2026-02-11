package main

import "fmt"

func main() {
	// cara buat array di golang. ini dia make kurung kurawal
	// array di golang itu length nya bersifat immutable
	kumpulanAngka := [3]int{1, 2, 3}
	nama_nama := []string{"Tono", "Joni", "Anton", "MohanPedo"}
	fmt.Println(kumpulanAngka)
	fmt.Println(len(nama_nama))
	// fmt.Printf("%T",nama_nama)
	// yang dibawah ini katanya mirip dengan var_dump di php
	fmt.Printf("%#v\n", nama_nama)

	// ubah array 
	kumpulanAngka[1]  = 3
	fmt.Println("ini adalah value index ke 1 dari array kumpulanAngka",kumpulanAngka[1])
	
	// value array bisa diubah
	for i := 0; i < len(kumpulanAngka); i++ {
		kumpulanAngka[i] = kumpulanAngka[i] + 2
	}
	fmt.Println(kumpulanAngka)

	for index , value := range kumpulanAngka{
		fmt.Println("index ke", index, "=", value)
	}

	//array bisa dibandingkan di golang

	arr1 := [3]string{"Shidqi", "Suci", "Shifan"}
	arr2 := [3]string{"Shidqi", "Suci"}


	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 != arr2)
}