package main

import (
	"fmt"
	"strings"
)

//function bisa jadi suatu type di golang
type Filter func(string) string

type Blacklist func(string)bool

func sayHelloWithFilter(name string, filter Filter) string {
	filtered := filter(name)
	return "Hello" + filtered
}

func spamFilter(name string) string {
	return strings.ReplaceAll(name , "Anjing", "***")
}

func registerUser(name string, blacklist Blacklist) string {
	if blacklist(name){
		return "Anda di blacklist!"
	}else{
		return "Registrasi Berhasil"
	}
}

func factorialLoop(num int)int{
	result := 0
	for i := num; i  >= 0; i-- {
		result *= i 
	}
	return result
}

func main() {
	fmt.Println(sayHelloWithFilter("Shidqi Anjing Banget", spamFilter))
	fmt.Println(factorialLoop(5))

	//anonymous function. jadi functionnya langsung kita bikin saat registerUser

	blacklist := func(name string)bool{
		if name == "Reza" {
			return true
		}
		return false
	}

	// fmt.Println(blacklist("Reza"))

		fmt.Println(registerUser("Shidqi", blacklist))

	result := registerUser("Shidqi", func(s string) bool {
		if s == "Shidqi" {
			return true
		} else {
			return false
		}
	})
	fmt.Println(result)
}

