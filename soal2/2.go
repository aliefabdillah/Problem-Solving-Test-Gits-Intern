package main

import (
	"fmt"
	"strings"
)

func main(){
	//input panjang array skor peserta
	var n int
	fmt.Scan(&n)

	//deklrasi variabel array skor peserta
	listSkor := make([]int, n)
	//iterasi skor peserta kedalam
	for i := 0; i < n; i++ {
		fmt.Scan(&listSkor[i])
	}

	//input panjang array skor gits
	var m int
	fmt.Scan(&m)

	//deklrasi variabel array skor gits
	listSkorGits := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&listSkorGits[i])
	}

	//merubah list skor peserta menjadi unique menggunakan function makeUniqueList
	listSkorUnique := makeUniqueList(listSkor)

	//membuat array kosong  untuk hasil 
	result := make([]int, 0)

	for i := 0; i < len(listSkorGits); i++ {
		valGits := listSkorGits[i]

		//pengecekan jika skor gits lebih besar dari skor peserta paling tinggi
		if valGits > listSkorUnique[0] {
			//menambahkan result dengan nilai (posisi) 1
			result = append(result, 1)
			//memasukan nilai skor gits kedalam array skor peserta
			listSkorUnique = append([]int{valGits}, listSkorUnique...)
		}else{
			//variable found untuk penanda apakah skor gits sama dengan skor peserta yang ada atau tidak
			found := false
			for j := 0; j < len(listSkorUnique) ; j++ {
				valSkor := listSkorUnique[j]
				
				//jika skor gits sama dengan skor peserta yang sudah ada
				if valGits == valSkor {
					//found berubah menjadi true
					found = true
					//nilai result yaitu index dari skor peserta + 1
					result = append(result, j+1)
					break
				}
			}

			//jika skor gits tidak ditemukan didalam list skor peserta
			if !found {
				var index int
				smallest := true		//varibel smallest untuk menentukan apakah varibel gits merupakan nilai terkecil atau bukan
				for k := 0; k < len(listSkorUnique) ; k++{
					// if skor gits lebih dari skor peserta
					if valGits > listSkorUnique[k] {
						// ambil nilai indeks dari skor peserta
						index = k
						//set variable smallest menjadi false
						smallest = false
						break;
					}
				}

				//jika skor gits merupakan nilai terkecil				
				if smallest {
					// tambahkan skor gits ke akhir array skor peserta
					listSkorUnique = append(listSkorUnique, valGits)
					//ambil panjang array skor peserta menjadi result (posisi)
					result = append(result, len(listSkorUnique))
				}else{
					//ambil index + 1 dari skor peserta menjadi result (posisi)
					result = append(result, index+1)
					//tambahkan skor gits ke dalam skor peserta sesuai index
					listSkorUnique = append(listSkorUnique[:index], append([]int{valGits}, listSkorUnique[index:]...)...)
				}
			}
		}
	}
	
	// print hasil posisi skor gits
	arrayString := strings.Trim(fmt.Sprint(result), "[]")
	fmt.Println("Output :", arrayString)
}

// fungsi membuat skor menjadi unique
func makeUniqueList(listSkor []int) []int {
	uniqueMapSkor := make(map[int]bool)		//varibale untuk menampung unique bertipe map
	uniqueArrSkor := []int{}				//variable untuk menampung unique bertipe array integer

	for _, val := range listSkor {
		if !uniqueMapSkor[val] {
			uniqueMapSkor[val] = true
			uniqueArrSkor = append(uniqueArrSkor, val)
		}
	}

	return uniqueArrSkor
}