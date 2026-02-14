package main

import "fmt"

func main() {
	counter := 0
	//function ini bisa mengubah variabel yang berada di luar scopenya
	// fitur ini berbahaya karena akan membuat bingung
	inc := func() {
		counter++
	}
	inc()
	inc()
	inc()
	fmt.Println(counter)
}