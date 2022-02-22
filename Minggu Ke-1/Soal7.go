// Soal No.7

/*
Buatlah program yang digunakan untuk menampilkan matrik tertentu.
Masukan terdiri dari dua baris dan tiga kolom (yang menyatakan
matrik berukuran 2x3). Terakhir sebuah bilangan x yang menyatakan
besarnya rotasi matrik yang dilakukan (x adalah 0, 90, 180, atau 270, 360).
Keluaran terdiri dari tiga baris dan dua kolom, yang menyatakan hasil
rotasi searah jarum jam sebesar x derajat.

a,b,c = input().split()
d,e,f = input().split()
x = int(input())

if x == 0 or x == 360:
	print(a,b,c)
	print(d,e,f)
elif x == 90 :
	print(d,a)
	print(e,b)
	print(f,c)
elif x == 180 :
	print(f,e,d)
	print(c,b,a)
elif x == 270 :
	print(c,f)
	print(b,e)
	print(a,d)

No 		Masukan 		Keluaran
1 		10 -2 73 		-4 10
		-4 25 6			25 -2
		90 				6 73

2 		10 -2 73 		10 -2 73
		-4 25 6 		-4 25 6
		0

3 		-71 82 -93 		-71 82 -93
		34 -25 16 34 	-25 16
		360

4 		-71 82 -93 		-93 16
		34 -25 16 		82 -25
		270 			-71 34

5 		-10 -2 73 		6 25 -4
		-4 25 6 		73 -2 -10
		180
*/

package main

import "fmt"

func main() {
	var (
		a, b, c int
		d, e, f int
		x       int
	)

	fmt.Scan(&a, &b, &c)
	fmt.Scan(&d, &e, &f)
	fmt.Scan(&x)

	if x == 0 || x == 360 {
		fmt.Println(a, b, c)
		fmt.Println(d, e, f)
	} else if x == 90 {
		fmt.Println(d, a)
		fmt.Println(e, b)
		fmt.Println(f, c)
	} else if x == 180 {
		fmt.Println(f, e, d)
		fmt.Println(c, b, a)
	} else if x == 270 {
		fmt.Println(c, f)
		fmt.Println(b, e)
		fmt.Println(a, d)
	}
}
