package main

import "fmt"

func main(){
	var berat, kg, g int
	var b1, b2, ttl int

	fmt.Print("Masukan Total Berat (gram): ")
	fmt.Scan(&berat)

	fmt.Println("_____RINCIAN PEMBAYARAN_____")
	fmt.Println(" ")
	kg = berat / 1000
	g = berat % 1000

	if kg > 10 {
		b1 = kg * 10000
		b2 = g * 0
		ttl = b1 + b2
		fmt.Printf("Detail Berat: %d kg + %d gram\n", kg, g)
		fmt.Printf("Detail Pembayaran: Rp. %d + Rp. %d\n", b1, b2)
		fmt.Printf("Total Pembayaran: Rp. %d\n", ttl)
	} else if g >= 500 {
		b1 = kg * 10000
		b2 = g * 5
		ttl = b1 + b2
		fmt.Printf("Detail Berat: %d kg + %d gram\n", kg, g)
		fmt.Printf("Detail Pembayaran: Rp. %d + Rp. %d\n", b1, b2)
		fmt.Printf("Total Pembayaran: Rp. %d\n", ttl)
	} else if g < 500 {
		b1 = kg * 10000
		b2 = g * 15
		ttl = b1 + b2
		fmt.Printf("Detail Berat: %d kg + %d gram\n", kg, g)
		fmt.Printf("Detail Pembayaran: Rp. %d + Rp. %d\n", b1, b2)
		fmt.Printf("Total Pembayaran: Rp. %d\n", ttl)
	}
	fmt.Println(" ")
	fmt.Println("_____THANK YOU_____")
}