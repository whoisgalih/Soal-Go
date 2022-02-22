// Soal No.8

/*
Buatlah sebuah fungsi hitungVolume yang akan mengembalikan nilai
volume dengan r dan t sebagai parameternya
Masukan berupa bilangan bulat r dan t,
Keluaran berupa volume silinder haril perhitungan dengan
menggunakan fungsi hitungVolume.
Catatan: gunakan pi = 3.14

def hitungVolume(r, t):
	pi = 3.14
	return r*r*pi*t

r,t = input().split()
print(hitungVolume(int(r), int(t)))


No 		Masukan 	Keluaran
1 		1 9 		28.26
2 		4 6 		301.44
3 		3 7 		197.82000000000002
4 		5 5 		392.5
5 		2 8 		100.48
*/

package main

import "fmt"

func hitungVolume(r, t float64) float64 {
	var (
		pi    float64
		hasil float64
	)
	pi = 3.14
	hasil = r * r * pi * t
	return hasil
}

func main() {
	var (
		r, t   float64
		volume float64
	)

	fmt.Scan(&r, &t)
	volume = hitungVolume(r, t)

	fmt.Println(volume)

}
