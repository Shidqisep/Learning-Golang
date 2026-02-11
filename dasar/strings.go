package main

import (
	"fmt"
	"strings"
)

func stringss() {
	// menggabungkan string
	str1 := "Shidqi," + "Andra" 
	str2 := str1 + " " +  str1
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(len(str2))
	fmt.Println(str1[2])
	fmt.Println(strings.HasPrefix(str1, "Sh"))
	fmt.Println(strings.ToUpper(str2))
	fmt.Println(strings.Contains(str1, "Shidqi"))
	fmt.Println(strings.Split(str2, ","))
	fmt.Println(strings.ReplaceAll(str1, "Shidqi", "Reza"))
}