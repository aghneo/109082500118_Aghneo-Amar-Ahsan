package main 

import "fmt"

type arrInt [2023]int

func terkecil_1(tabInt arrInt, n int) int {
	var min int = tabInt[0]
	var x int = 1

	for x < n{
		if min > tabInt[x]{
			min = tabInt[x]
		}
		x++
	}
	return min
}

func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0
	var x int = 1

	for x < n{
		if tabInt[idx] > tabInt[x]{
			idx = x
		}
		x++
	}
	return idx
}

func main(){
	var tabInt arrInt 
	var N int

	fmt.Print("Masukan Jumlah Data : ")
	fmt.Scan(&N)

	fmt.Print("Masukan Elemen Array : ")

	for i := 0; i < N; i ++ {
		fmt.Scan(&tabInt[i])
	}

	min := terkecil_1(tabInt, N)
	idx := terkecil_2(tabInt, N)

	fmt.Println("Nilai Minimal Adalah : ", min)
	fmt.Println("Nilai Minimal Ada Di Index : ", idx)
}