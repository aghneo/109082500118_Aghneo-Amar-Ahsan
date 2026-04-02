package main

import "fmt"

func main(){
	var p1, p2, p3, p4 string
	var x bool = true

	for i := 1 ; i <= 5 ; i++ {
		fmt.Printf("Percobaan Ke-%d: ", i)
		fmt.Scan(&p1, &p2, &p3, &p4)

		if p1 != "merah" || p2 != "kuning" ||  p3 != "hijau" ||  p4 != "ungu"{
			x = false
		} 
	}
	fmt.Println("PERCOBAAN BERHASIL: ", x)
}