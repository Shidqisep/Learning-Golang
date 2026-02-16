package internal

import "fmt"


//disini dia tidak ada method, kalo impor biasa dia akan error karena kan tidak digunakan
func init() {
	fmt.Println("This is internal only")
}