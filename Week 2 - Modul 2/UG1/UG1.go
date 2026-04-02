package  main

import "fmt"

func main(){
	var thn int

	fmt.Print("Masukan Tahun: ")
	fmt.Scan(&thn) 
	
	if (thn % 4 == 0 && thn % 100 != 0) || (thn % 400 == 0) {
		fmt.Printf("Tahun %d Adalah Tahun Kabisat\n", thn)
	} else {
		fmt.Printf("Tahun %d Bukan Tahun Kabisat\n", thn)
	} 
}