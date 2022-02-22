// Soal No.2

/*
Buatlah sebuah program yang digunakan untuk menghitung luas
lingkaran dengan jari-jari r.
Masukan berupa sebuah bilangan asli r.
Keluaran adalah sebuah string yang berisi luas lingkaran, dengan
format seperti berikut ini:

Luas lingkaran dengan jari-jari = r adalah luas_lingkaran
r = input()
luas_lingkaran = 22 / 7 * int(r)
print("Luas lingkaran dengan jari-jari =",r,"adalah",luas_lingkaran)
Catatan: Gunakan fungsi float64(x) untuk casting data dari integer ke
float / real


No 		Masukan 	Keluaran
1 		5 			Luas lingkaran dengan jari-jari = 5 adalah 15.71428...
2 		14 			Luas lingkaran dengan jari-jari = 14 adalah 44
3 		314000 		Luas lingkaran dengan jari-jari = 314000 adalah 986857.142...
4 		1 			Luas lingkaran dengan jari-jari = 1 adalah 3.142857...
*/

package main

import "fmt"

func main() {
	var (
		r              int
		luas_lingkaran float64
	)
	fmt.Scan(&r)
	luas_lingkaran = float64(22) / float64(7) * float64(r)

	fmt.Println("Luas lingkaran dengan jari-jari=", r, "adalah", luas_lingkaran)
}
