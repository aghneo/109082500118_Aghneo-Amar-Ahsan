package main

import "fmt"

func main() {
	var ikan [1000]float64
	var x, y int

	fmt.Print("Masukan Banyak Ikan Yang Akan Dijual (x) : ")
	fmt.Scan(&x)
	
	fmt.Print("Masukan Kapasitas Ikan Per Wadah (y) : ")
	fmt.Scan(&y)

	var totalBeratKeseluruhan float64 = 0
	
	for i := 0; i < x; i++ {
		fmt.Print("Masukan Berat Ikan Ke-", i+1, " : ")
		fmt.Scan(&ikan[i])
		totalBeratKeseluruhan += ikan[i]
	}

	var beratWadahSaatIni float64 = 0
	var jumlahWadah float64 = 0

	fmt.Println("\n--- Hasil Keluaran ---")
	fmt.Print("Total Berat Ikan di Setiap Wadah : ")
	
	for i := 0; i < x; i++ {
		beratWadahSaatIni += ikan[i]

		if (i+1)%y == 0 || i == x-1 {
			fmt.Print(beratWadahSaatIni, " ")
			jumlahWadah++
			beratWadahSaatIni = 0
		}
	}

	fmt.Println()

	rataRata := totalBeratKeseluruhan / jumlahWadah
	fmt.Printf("Berat Rata-Rata Ikan di Setiap Wadah : %.2f", rataRata)
}