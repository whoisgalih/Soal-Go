// Soal No.5

/*
Buatlah sebuah program yang akan terus menerima masukan berupa
string secara terus menerus dan menampilkannya hingga masukan
berupa string “selesai”.

string = input()
while string != "selesai":
print(string)
string = input()

No 		Masukan 		Keluaran
1 		nk 				nk
		ugrp 			ugrp
		qib 			qib
		racxmw 			racxmw
		zvuatp 			zvuatp
		selesai

2 		yeebc 			yeebc
		rvmwq 			rvmwq
		qz 				qz
		g 				g
		selesai

3 		selesai
*/

package main

import "fmt"

func main() {
	var kata string

	fmt.Scan(&kata)

	for kata != "selesai" {
		fmt.Println(kata)
		fmt.Scan(&kata)
	}
}
