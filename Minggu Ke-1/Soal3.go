// Soal No.3

/*
Buatlah sebuah program yang digunakan untuk mengecek kelipatan
suatu bilangan x.
Masukan berupa sebuah bilangan asli x.
Keluaran berupa string “Fizz” apabila x adalah kelipatan 3, atau/dan
string “Bazz” apabila kelipatan 5.

x = input()
if int(x) % 3 == 0:
print("Fizz")
if int(x) % 5 == 0:
print("Bazz")


No 		Masukan 	Keluaran
1 		5 			Bazz
2 		9 			Fizz
3 		314000 		Bazz
4 		30 			Fizz
5 		8 			Bazz
*/

package main

import "fmt"

func main() {
	var (
		x int
	)
	fmt.Scan(&x)

	if x%3 == 0 {
		fmt.Println("Fizz")
	}
	if x%5 == 0 {
		fmt.Println("Bazz")
	}
}
