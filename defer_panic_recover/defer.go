package main

import "fmt"

func logging() {
	fmt.Println("Logging....")
}

func runApp()  {
	defer logging() // defer ini akan memanggil function logging di akhirnya
	fmt.Println("Running APP")
}

func mainDefer()  {
	runApp()
}