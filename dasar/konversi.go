package main

import (
	"fmt"
	"strconv"
)

func konversis() {
	var a int = 10
	// variabel a bisa diubah tipe datanya ke float dengan cara berikut
	b := float64(a)
	//mengubah int ke string akan menghasilkan string kosong jika konversinya seperti dibawah ini
	c := string(a)
	//cara yang benar adalah menggunakan library bawaan si golang
	var text string = strconv.Itoa(a)
	// cara konversi dari string ke number
	str := "100"
	fmt.Printf("tipe data variabel b: %T\n", b)
	fmt.Printf("tipe data variabel a: %T\n", a)
	fmt.Printf("tipe data variabel c: %T\n", c)
	fmt.Println(text)
	fmt.Println(text + c)

	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error mengubah ke integer: ", err.Error())
	} else {
		fmt.Println("Berhasil mengubah ke integer", number)
	}

	// boolean ke string
	truth := true
	boolStr := strconv.FormatBool(truth)
	fmt.Println(boolStr)
	// string ke boolean
	falsy := "false"
	strBool, _ := strconv.ParseBool(falsy)
	fmt.Println(strBool)

}