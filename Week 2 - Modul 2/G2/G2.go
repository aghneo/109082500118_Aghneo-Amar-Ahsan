package main

import "fmt"

func main (){
	var angka int
	fmt.Print("Masukan Inputan : ")
	fmt.Scan(&angka)

	for angka < 10 {
		fmt.Println("Angka Saat Ini : ", angka)
		angka++
	}
}