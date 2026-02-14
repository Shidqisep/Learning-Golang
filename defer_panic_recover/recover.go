package main

import "fmt"

func endApps()  {
	fmt.Println("End Apps")
	message := recover()
	fmt.Println("terjadi error:", message)
}



// dibawah ini adalah cara salah menggunakan recover
/*
func runApps(err bool) {
	defer endApps()
	message := recover()
	fmt.Println("terjadi error", message)
	if err {
		panic("ERROR")
	}
}

*/

//cara yang benar adalah recovernya berada di defer function
func runApps(err bool) {
	defer endApps()
	if err {
		panic("ERROR")
	}
}

func main()  {
	runApps(true)
	fmt.Println("Programnya Lanjut...")
}