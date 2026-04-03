package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan Batas : ")
	fmt.Scan(&n)
	fmt.Println("________________________")
	bintang(n)
}

func bintang(n int)  {
	if n > 0 {
		bintang(n - 1)
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
