package main

import "fmt"

func random(a any) interface{} {
	return a
}


func main() {
	var result any = random(nil)
	/*
	var resultString = result.(string)
	fmt.Println(resultString)

	var resultInt = result.(int) //panic
	fmt.Println(resultInt)
	*/

	//sebaiknya menggunakan switch untuk type assertion
	switch result.(type) {
		case string:
			fmt.Println("String", result)
		case int:
			fmt.Println("Int", result)
		case bool:
			fmt.Println("Bool", result)
		default:
			fmt.Println("Tipe data tidak dikenal")
		}
}