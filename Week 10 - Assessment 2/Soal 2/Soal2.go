package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func nilaiPertama(arr arrayMahasiswa, n int, targetNIM string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			return arr[i].nilai
		}
	}
	return -1
}

func nilaiTerbesar(arr arrayMahasiswa, n int, targetNIM string) int {
	max := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM == targetNIM {
			if arr[i].nilai > max {
				max = arr[i].nilai
			}
		}
	}
	return max
}

func main() {
	var n int
	var arrMhs arrayMahasiswa
	var NIM string

	fmt.Print("Masukkan Jumlah Data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan Data Ke-", i+1, " : ")
		fmt.Scan(&arrMhs[i].NIM, &arrMhs[i].nama, &arrMhs[i].nilai)
	}

	fmt.Print("Masukkan NIM Mahasiswa Yang Ingin Dicari Nilai Pertama dan Nilai Terbesarnya : ")
	fmt.Scan(&NIM)

	Pertama := nilaiPertama(arrMhs, n, NIM)
	Terbesar := nilaiTerbesar(arrMhs, n, NIM)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", NIM, Pertama)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", NIM, Terbesar)
}