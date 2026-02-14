package main

import (
	"fmt"
	_"log"
)

func endApp() {
	fmt.Println("End App")
}

func runApplication(errors bool)  {
	defer endApp()
	if errors {
		// log.Fatal("ERROR") //log fatal ini beneran memberhentikan aplikasi secara menyeluruh
		panic("ERROR")
	}
	fmt.Println("Running App")
}

func panics()  {
	runApplication(true)
}



