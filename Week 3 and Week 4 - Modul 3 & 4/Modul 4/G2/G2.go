package main

import "fmt"

func hitungTotal(harga, jumlah int) {
	var total int
	total = harga * jumlah
	fmt.Println("Total harga = ", total)
	hitungDiskon(total)
}

func hitungDiskon(total int) {
	var diskon, hargaAkhir int
	diskon = total * 10 / 100
	hargaAkhir = total * diskon
	fmt.Println("Diskon 10%:", diskon)
	fmt.Println("Harga setelah diskon : ", hargaAkhir)
}

func main() {
	var price, qty int

	fmt.Print("Masukan Harga Barang : ")
	fmt.Scan(&price)
	fmt.Print("Masukan Jumlah Barang : ")
	fmt.Scan(&qty)

	hitungTotal(price, qty)
}
