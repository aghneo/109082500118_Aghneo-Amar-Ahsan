package main

import "fmt"
func hitungBiaya(jenis string, masuk, keluar int) int {
	biaya := 0
	for jam := masuk; jam < keluar; jam++ {
		if jenis == "motor" {
			if jam < 17 {
				biaya += 4000 
			} else {
				biaya += 5000 
			}
		} else if jenis == "mobil" {
			if jam < 17 {
				biaya += 6000 
			} else {
				biaya += 7000 
			}
		}
	}
	return biaya
}

func main() {
	var jenis string
	var masuk, keluar int
	var totalPendapatan int
	kendaraanKe := 1

	for {
		fmt.Printf("Kendaraan Ke-%d\n", kendaraanKe)
		fmt.Print("Jenis Kendaraan (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&masuk)
		fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&keluar)

		biaya := hitungBiaya(jenis, masuk, keluar)
		fmt.Printf("Biaya parkir %s %d: %d\n", jenis, kendaraanKe, biaya)
		fmt.Println("========================================")

		totalPendapatan += biaya
		kendaraanKe++
	}
	fmt.Println(" ")
	fmt.Printf("Total Pendapatan Parkir Hari Ini : %d", totalPendapatan)
}