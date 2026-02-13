package main

import "fmt"

func main() {
	hitungan := 0
	for hitungan < 10 {
		fmt.Println("Woi ke ", hitungan)
		hitungan++
	}
	fmt.Println(hitungan)

	for i := 0; i < 5; i++ {
		fmt.Println("Counter ke ", i)
	}

	names := []string{"Shidqi", "Anton", "Karen", "Doni"}	
	//menggunakan for range looping slice, array atau map
	for index, value := range names{
		fmt.Println(index, "isinya", value)
	}

	//ini kalo ga butuh indexnya
	for _, value := range names{
		fmt.Println("isinya", value)
	}

	mapAku := map[string]string{
		"User1":"Anton",
		"user2":"Tono",
		"User3":"Prapto",
		"User4":"Raisa",
	}

	//bisa pake break dan continue

	for key, value := range mapAku {
		fmt.Println(key, value)
		if mapAku[key] == "Prapto" {
			continue
		}else {
			mapAku[key] = mapAku[key] + "1"
		}
	}
	fmt.Println(mapAku)

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Angka Pasti Ganjil", i)
	}



}