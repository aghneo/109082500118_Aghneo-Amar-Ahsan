package main

import "fmt"

func SequnsialSearch(arr []string, dataCari string) int {
	var idx_found int = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == dataCari {
			idx_found = i
			break
		}
	}
	return idx_found
}

func main() {
	var array [5]string
	var dataCari string

	for i := 0; i < len(array); i++ {
		fmt.Print("Masukan Data : ")
		fmt.Scan(&array[i])
	}

	fmt.Print("Masukan Data Yang Mau Dicari : ")
	fmt.Scan(&dataCari)
	idx_found := SequnsialSearch(array[:], dataCari)

	if idx_found > -1 {
		fmt.Println("Data Ditemukan Pada Index Ke-", idx_found)
	} else {
		fmt.Println("Data Tidak Ditemukan!!!")
	}
}
