// Soal No.1

/*
Tulis ulang beberapa kode python berikut menjadi sebuah program Go.
Perhatikan format umum penulisan dari Go berikut ini!

a. Buatlah sebuah program dengan ketentuan sebagai berikut :
Masukan berupa sebuah baris yang berisi sebuah string S tanpa spasi,
dan 2 buah bilangan bulat A dan B yang masing-masing inputnya
dipisahkan dengan spasi.
Keluaran terdiri dua baris. Baris pertama adalah S, sedangkan baris
kedua adalah hasil operasi penjumlahan A dengan B. Tampilan harus
sesuai dengan format contoh yang diberikan berikut:

Kata = S
Jumlah = hasil_penjumlahan
S, A, B = input().split()
hasil_penjumlahan = int(A) + int(B)
print("Kata =", S)
print("Jumlah =", hasil_penjumlahan)

No 		Masukan 		Keluaran
1  		mbi 16 62 		Kata = mbi
						Jumlah = 78

2 		mzjpl -18 27 	Kata = mzjpl
						Jumlah = 9

3 		qejey 3 -79 	Kata = qejey
						Jumlah = -76

4 		zirwz -19 -18 	Kata = zirwz
						Jumlah = -37

5 		jdxcv 10 60 	Kata = jdxcv
						Jumlah = 70
*/
package main

import "fmt"

func main() {
	/*
		// Cara Penulisan Var Ke-1
		var s string
		var a, b int
		var hasilPenjumlahan int
	*/
	// Cara Penulisan Var Ke-2
	var (
		s                string
		a, b             int
		hasilPenjumlahan int
	)

	fmt.Scan(&s, &a, &b)
	hasilPenjumlahan = a + b

	fmt.Println("Kata =", s)
	fmt.Println("Jumlah =", hasilPenjumlahan)
}
