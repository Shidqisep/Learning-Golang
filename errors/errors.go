package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int,error) {
	if pembagi == 0 {
		return 0, errors.New("Undefined divide by 0")
	}else{
		return  nilai/pembagi, nil
	}
}

func main() {
	// ini kalo dia balikin 2 value
	result, err := pembagian(10, 0)

	/*
		ini yang biasanya dilakukan oleh programmer golang, cek error secara manual
	*/
	if err == nil {
		fmt.Println(result, "dan errornya adalah", err)
	}else{
		fmt.Println("ini error:", err)
	}


}