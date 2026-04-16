package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai N : ") 
	fmt.Scan(&n)
	barisanBilangan(n)
	fmt.Println()
}

func barisanBilangan(n int) {
	if n == 1 {
		fmt.Print(n, " ")
	} else {
		fmt.Print(n, " ")       
		barisanBilangan(n - 1)  
		fmt.Print(n, " ")       
	}
}