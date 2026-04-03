package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x :")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&y) 

	
	hasil := hitungPangkat(x, y)
	fmt.Println(hasil)
}

func hitungPangkat(x, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * hitungPangkat(x, y-1)
	}
}