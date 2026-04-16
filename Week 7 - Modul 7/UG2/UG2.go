package main

import "fmt"

type angka int
type kata string
type Buku struct {
	judul, penulis, penerbit   kata
	tahunTerbit, jumlahHalaman angka
}

func main() {
	var b Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukan Judul Buku : ")
	fmt.Scan(&b.judul)
	fmt.Print("Masukan Nama Penulis : ")
	fmt.Scan(&b.penulis)
	fmt.Print("Masukan Penerbit : ")
	fmt.Scan(&b.penerbit)
	fmt.Print("Masukan Tahun Terbit : ")
	fmt.Scan(&b.tahunTerbit)
	fmt.Print("Masukan Jumlah Halaman : ")
	fmt.Scan(&b.jumlahHalaman)

	fmt.Println(" ")

	fmt.Println("=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", b.judul)
	fmt.Println("Penulis :", b.penulis)
	fmt.Println("Penerbit :", b.penerbit)
	fmt.Println("Tahun Terbit :", b.tahunTerbit)
	fmt.Println("Jumlah Halaman :", b.jumlahHalaman)
}
