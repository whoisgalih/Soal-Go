// Soal No.9

/*
Sebuah tim yang terdiri dari 3 anggota mengikuti sebuah perlombaan.
Dalam perlombaan tersebut, sebuah tim diharuskan menyelesaikan
soal-soal yang diberikan. Tim tersebut akan memutuskan untuk
menjawab suatu soal jika lebih dari satu anggota merasa bahwa
jawaban tersebut benar.
Perlombaan tersebut memberikan n soal. Bantu tim tersebut untuk
mengetahui berapa soal yang bisa mereka jawab dengan cara
membuat program yang dapat mengetahui apakah soal tersebut
dapat dijawab atau tidak.
Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan
bulat n yang menyatakan jumlah soal yang diberikan. Selanjutnya
untuk n baris berikutnya, masing-masing terdapat t1, t2, dan, t3 yang
menunjukan keputusan setiap anggota. Jika suatu anggota bernilai 1,
maka ia yakin bisa menjawab soal tersebut, dan jika suatu anggota
bernilai 0, maka ia tidak yakin bisa menjawab soal tersebut.
Keluaran adalah jumlah soal yang dapat dijawab oleh tim mereka
(tanpa memperdulikan jawaban benar atau salah).

n = int(input())
jumlah = 0
for i in range(n):
	t1,t2,t3 = input().split()
	if int(t1) + int(t2) + int(t3) >= 2 :
		jumlah = jumlah + 1
print(jumlah)

No 		Masukan 	Keluaran
1 		7 			4
		1 0 1
		1 1 1
		1 1 0
		0 1 0
		0 1 0
		1 0 0
		1 1 0

2 		4 			2
		0 1 0
		0 1 1
		0 1 1
		0 1 0

3 		6 			3
		1 0 0
		0 0 0
		1 1 0
		1 0 0
		1 1 1
		1 1 1

4 		10			2
		0 1 1
		0 0 0
		1 0 0
		1 0 1
		1 0 0
		0 0 0
		0 0 0
		0 0 1
		0 1 0
		0 0 0
		0 0 0
*/

package main

import "fmt"

func main() {
	var (
		banyakSoal          int
		jumlahSoalBisaJawab int
		t1, t2, t3          int
	)

	fmt.Scan(&banyakSoal)

	for i := 0; i < banyakSoal; i++ {
		fmt.Scan(&t1, &t2, &t3)
		if t1+t2+t3 >= 2 {
			jumlahSoalBisaJawab++
		}
	}
	fmt.Println(jumlahSoalBisaJawab)
}
