// Soal No.4

/*
Buatlah sebuah program untuk menghitung rekor.
Masukan berupa 4 buah bilangan bulat a, b, c, dan d.
Keluaran berupa string "Ada rekor baru" apabila nilai d lebih besar dari
a, b, dan c, atau "Ada rekor baru" apabila sebaliknya.
a, b, c, d = input().split()

if int(d) > int(a) and int(d) > int(b) and int(d) > int(c):
	print("Ada rekor baru")
else:
	print("Tidak ada rekor baru")


No 		Masukan 		Keluaran
1 		27 28 11 8 		Tidak ada rekor baru
2 		13 2 19 30 		Ada rekor baru
3 		30 27 5 280 	Ada rekor baru
4 		0 0 0 0 		Tidak ada rekor baru
5 		-19 -25 29 25 	Tidak ada rekor baru
*/

package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	if d > a && d > b && d > c {
		fmt.Println("Ada rekor baru")
	} else {
		fmt.Println("Tidak ada rekor baru")
	}
}
