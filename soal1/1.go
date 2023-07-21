package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//deklarasi variable dan output
	var input int
	var outputString []string

	//input menggunakan scan
	fmt.Print("Masukkan angka: ")
	_, err := fmt.Scan(&input)

	//cek error input
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input:", err)
		os.Exit(1)
	}

	//perulangan selama i < input
	for i := 0; i < input; i++ {
		//penghitungan rumus A000124 of Sloane’s OEIS
		output := (i*i + i + 2)/2
		//merubah output int menjadi string
		outputString = append(outputString, strconv.Itoa(output))
	}

	//menambahkan separator "-""
	result := strings.Join(outputString, "-")
	fmt.Println("output:", result)
}