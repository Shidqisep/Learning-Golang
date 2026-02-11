package main

import "fmt"

func datatypes() {
	/*
	tipe data di golang ada banyak
	int untuk integer dasar
	int8 untuk -128 sampai 128
	int16 untuk -32768 sampai 32767
	int32 untuk  -2,147,483,648 sampai 2,147,483,647.
	int64	untuk yang gede banget

	ada juga unsigned integer yaitu uint
	valuenya sama akan tetapi hanya positif saja
	*/

	//contoh
	var aInt int = 10
	var aInt8 int8 = -128
	var aInt16 int16 = 32000
	var aInt32 int32 = 2000000000
	var aInt64 int64 = 9000000000000000000
	var aUint uint = 10
	var aUint8 uint8 = 255
	var aUint16 uint16 = 65535
	var aUint32 uint32 = 4000000000
	var aUint64 uint64 = 18000000000000000000
	var aFloat32 float32 = 3.14
	var aFloat64 float64 = 3.141592653589793
	var aComplex64 complex64 = 1 + 2i
	var aComplex128 complex128 = 3.14 + 2.71i
	var aBool bool = true
	var aString string = "Halo Golang"
	var aByte byte = 65        // alias uint8
	var aRune rune = '中'      // alias int32 (Unicode)
	
	// Integer types
	fmt.Println("aInt:", aInt)
	fmt.Println("aInt8:", aInt8)
	fmt.Println("aInt16:", aInt16)
	fmt.Println("aInt32:", aInt32)
	fmt.Println("aInt64:", aInt64)

	// Unsigned Integer types
	fmt.Println("aUint:", aUint)
	fmt.Println("aUint8:", aUint8)
	fmt.Println("aUint16:", aUint16)
	fmt.Println("aUint32:", aUint32)
	fmt.Println("aUint64:", aUint64)

	// Float types
	fmt.Println("aFloat32:", aFloat32)
	fmt.Println("aFloat64:", aFloat64)

	// Complex types
	fmt.Println("aComplex64:", aComplex64)
	fmt.Println("aComplex128:", aComplex128)

	// Other types
	fmt.Println("aBool:", aBool)
	fmt.Println("aString:", aString)
	
	// Special Note: By default, Println prints the numeric value of bytes and runes
	fmt.Println("aByte (Numeric):", aByte) 
	fmt.Println("aRune (Numeric):", aRune)
	fmt.Println("Hello")

	//di golang, string merupakan array of character atau array dari char
}