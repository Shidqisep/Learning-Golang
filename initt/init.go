package initt

import "fmt"

var kehidupan string

func init() {
	kehidupan = "Bernegara"
	fmt.Println("hello from inits")
}

func Test() {
	fmt.Println("hello from test")
}	