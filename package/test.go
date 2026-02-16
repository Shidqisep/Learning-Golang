package main

import (
	"Dasar_Golang/helper"
	"Dasar_Golang/database"
	//underscore hanya untuk menjalankan init saja
	_"Dasar_Golang/internal"
	"fmt"
)

func main()  {
	//kalian bisa gunakan package yang lain ya. ini ngambil dari package helper
	fmt.Println(helper.SayHello())
	fmt.Println(helper.Application)
	helper.SayMwhahaha()
	fmt.Print(database.GetDatabase())
	
}