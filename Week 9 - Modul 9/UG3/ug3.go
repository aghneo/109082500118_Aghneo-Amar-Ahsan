package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang [1000]string
	var jumlahPertandingan int

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pertandinganKe := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandinganKe)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[jumlahPertandingan] = klubA
		} else if skorB > skorA {
			pemenang[jumlahPertandingan] = klubB
		} else {
			pemenang[jumlahPertandingan] = "Draw"
		}

		jumlahPertandingan++
		pertandinganKe++
	}

	fmt.Println()
	for i := 0; i < jumlahPertandingan; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}
