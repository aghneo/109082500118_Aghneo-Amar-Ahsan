package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
	return 0
}

func biayaPerHari(jumlahMhs int) int {
	return jumlahMhs * (70000 + 250000 + 300000)
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hari := tanggunganHari(lamaPerjalanan, tujuan)
	biayaHarianDomestik := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hari * biayaHarianDomestik)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hari*biayaHarianDomestik) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukan lama hari study tour: ")
	fmt.Scan(&lama)
	fmt.Print("Masukan tujuan study (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U: Rp.%.f\n", biaya)
}