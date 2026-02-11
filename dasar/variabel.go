package main 

import "fmt"

func main(){

	var bilangan16 uint32
	var nama string
	const Nama string = "Shidqi"
	var umur int  = 20
	

	//ini tipenya golang yang nentuin
	var angka = 123123
	var kata = "Hello"
	//var itu cara lama, ini yang dibawah cara baru
	pesan := "Shidqi Ganteng"
	const Pi = 3.14
	pesanPanjang := `Shidqi Sangat NPD LOL 
	KOCAK`
	kata  = "Kehidupan"
	fmt.Println(angka)
	fmt.Println(kata)
	//ini dibawah semacam type of golang
	fmt.Printf("%T\n", bilangan16)
	fmt.Println(pesan)
	fmt.Println(nama)
	fmt.Println(Nama)
	fmt.Println(umur)
	fmt.Println(pesanPanjang)
}