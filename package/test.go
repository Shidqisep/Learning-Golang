package main

import (
	"Dasar_Golang/helper"
	"fmt"
)

func main()  {
	//kalian bisa gunakan package yang lain ya. ini ngambil dari package helper
	fmt.Println(helper.SayHello())
	fmt.Println(helper.Application)
	helper.SayMwhahaha()
}