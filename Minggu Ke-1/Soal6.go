// Soal No.6

/*
Buatlah sebuah program yang menentukan rata-rata nilai seorang
mahasiswa.
Masukan berupa beberapa nilai (bilangan bulat) mahasiswa yang
dipisahkan oleh spasi. Nilai akhir adalah -1 untuk mengakhiri input.
Keluaran berupa bilangan desimal yang menyatakan rata-rata nilai
mahasiswa.

nilai = int(input())
jumlah = 0
n = 0

while nilai != -1:
	n = n + 1
	jumlah = jumlah + nilai
	nilai = int(input())
if n == 0:
	rata2 = 0.0
else:
	rata2 = jumlah / n
	print(rata2)


No 		Masukan 		Keluaran
1 		10 				46.42857142857143
		97
		53
		5
		33
		65
		62
		-1

2 		-1 				0.0
*/

package main

import "fmt"

func main() {
	var (
		nilai, jumlahNilai int
		jumlahMahasiswa    int
		rerata             float64
	)
	fmt.Scan(&nilai)
	jumlahNilai = 0
	jumlahMahasiswa = 0
	rerata = 0

	for nilai != -1 {
		jumlahMahasiswa = jumlahMahasiswa + 1
		jumlahNilai = jumlahNilai + nilai
		fmt.Scan(&nilai)
	}

	if jumlahMahasiswa == 0 {
		rerata = 0
	} else {
		rerata = float64(jumlahNilai) / float64(jumlahMahasiswa)
	}

	fmt.Println(rerata)
}
