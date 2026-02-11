
package main

import "fmt"

func main(){
	mahasiswa := map[string]string{
		"nama":"Riyan Gaming",
		"umur":"20 Tahun",
		"Pekerjaan": "Mahasiswa",
	}

	fmt.Println("Nama:", mahasiswa["nama"], "\nUmur:", mahasiswa["umur"])

	nilaiMahasiswa := map[string]int{
		"Tono":20,
		"Joko":80,
	}

	nilaiMahasiswa["Anton"] = 50
	//ambil value map
	nilaiJoko := nilaiMahasiswa["Joko"]
	fmt.Println(nilaiJoko)

	//cek apakah key ada
	value, ok := nilaiMahasiswa["Shidqi"]
	if !ok {
		fmt.Println("key tidak ada",value , ok)
	}else{
		fmt.Println("key ada, nilainya:", value, ok)
	}

	//multi dimensional map
	multiMap := make(map[string]map[string]int)
	multiMap["Shidqi"] = make(map[string]int)
	multiMap["Anton"] = make(map[string]int)
	multiMap["Shidqi"]["ganteng"] = 10
	multiMap["Shidqi"]["jelek"] = 20
	multiMap["Anton"]["jelek"] = 20
	fmt.Println(multiMap["Shidqi"])
	fmt.Println(multiMap)

}