# Problem-Solving-Test-Gits-Intern
Repository untuk jawaban soal Problem Solving Test Gits Intern

## Screenshot Hasil Program
### 1. A000124 of Sloane’s OEIS

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/1_1.png)

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/1_2.png)

### 2. Dense Ranking

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/2_1.png)

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/2_2.png)

### 3. Balanced Bracket

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/3_1.png)

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/3_2.png)

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/3_3.png)

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/3_4.png)

#### Kompleksitas Algoritma

![alt text](https://github.com/aliefabdillah/Problem-Solving-Test-Gits-Intern/blob/main/screenshot/isBracketValid.png)

- Pada fungsi isBracketValid diatas yang berisi algoritma dari Balanced bracket terdiri dari syntax sederhana seperti inisialisasi variabel, pengecekan switch case dan if, pengurangan nilai pada array, dan return, yang dimana memiliki kompleksitas O(1),
- lalu terdapat satu perulangan for yang dimana kompleksitasnya tergantung dari panjang perulangan atau kita deskripsikan sebagai n sehingga bernilai O(n)
- lalu didalamnya terdapat operasi append, append sendiri nilainya adalah O(1) karena operasinya hanya menambahkan nilai ke baris paling belakang, tetapi append ini dalam case terburuk dapat bernilai O(n) jika memang array nya sudah penuh maka append akan membuat sebuah array baru lagi dan mengcopy seluruh data pada array lama ke array baru, akan tetapi karena disini array di set panjangnya dinamis sehingga membuat kompleksitasnya bernilai O(1)

maka dapat dijumlahkan kompleksitas pada algoritma difungsi tersebut yaitu:

O(1) + O(1) + O(1) + O(1) + O(1) + O(1) + O(n) + O(1) = O(n)

dalam skala big O sendiri O(n) / Linear tergolong kategori kompleksitas yang "Fair"
