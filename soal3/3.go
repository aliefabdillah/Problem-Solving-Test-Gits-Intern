package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// deklarasi input string
	var inputString string
	fmt.Print("input: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')

	//pengecekan input apakah sudah sesuai atau tidak
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input:", err)
		os.Exit(1)
	}

	// Menghapus whitespace (spasi) dari inputString
	inputString = strings.ReplaceAll(inputString, " ", "")
	
	//pemanggilan fungsi untuk mengecek bracket dari input user
	if isBracketValid(inputString) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/* Fungsi cek pola bracket dari input user */
func isBracketValid(str string) bool {
	//deklrasi variable stack dengan tipe array int32 untuk menampung setiap char dari input string
	stack := []int32{}

	// variable bracket berupa map int32 berupa pasangan bracket yang memenuhi syarat
	brackets := map[int32]int32{
		')': '(',
		'}': '{',
		']': '[',
	}

	// perulangan pengecekan selama panjang input string
	for _, char := range str {
		switch char {
		//jika index char berupa kurung buka
		case '(', '{', '[':					
			stack = append(stack, char)		//masukan ke dalam variable stack
		//jika index char berupa kurung tutup
		case ')', '}', ']':
			//lakukan pengecekan apabila stack sudah kosong atau stack paling akhir tidak sama dengan key pada variabel brackets atau nilainya bukan tutup kurung
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				//return false sebagai kembalian fungsi
				return false
			}

			//hapus index stack paling terakhir
			stack = stack[:len(stack)-1]
		}
	}

	//return kondisi true / jika stuck sudah kosong 
	return len(stack) == 0
}

/* 
Kompleksitas waktu = O(n)

pada fungsi isBracketValid yang berisi algoritma dari Balanced bracket terdiri dari syntax sederhana seperti inisialisasi variabel,
pengecekan switch case dan if, pengurangan nilai pada array, dan return, yang dimana memiliki kompleksitas O(1), lalu terdapat satu
perulangan for yang dimana kompleksitasnya tergantung dari panjang perulangan atau n sehingga bernilai O(n) lalu didalamnya terdapat
operasi append yang bernilai O(m), didapat m karena nilainya bergantung pada karakter kurung buka yang telah ditentukan sebelumnya, 
maka dapat dijumlahkan kompleksitas pada algoritma difungsi tersebut yaitu:

O(1) + O(1) + O(1) + O(1) + O(1) + O(1) + O(n) + O(m) = O(n + m)

dalam skala big O sendiri O(n + m) tergolong kategori kompleksitas yang "Fair"
*/
