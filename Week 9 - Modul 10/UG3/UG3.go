package main

import "fmt"

type arrayBalita [1000]float64

func MinMax(N int, arrayB arrayBalita) (min, max float64) {
	min = arrayB[0]
	for i := 1; i < N; i++ {
		if min > arrayB[i] {
			min = arrayB[i]
		}
	}

	max = arrayB[0]
	for i := 1; i < N; i++ {
		if max < arrayB[i] {
			max = arrayB[i]
		}
	}
	return min, max
}

func rata2(N int, arrayB arrayBalita) float64 {
	var total float64 = 0
	for i := 0; i < N; i++ {
		total += arrayB[i]
	}
	return total / float64(N)
}

func main() {
	var arrayB arrayBalita
	var N int

	fmt.Print("Masukan Banyak Data Berat Balita : ")
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Print("Masukan Berat Balita Ke-", i+1, " : ")
		fmt.Scan(&arrayB[i])
	}

	min, max := MinMax(N, arrayB)
	rata2 := rata2(N, arrayB)

	fmt.Println("Berat Balita Minimum : ", min)
	fmt.Println("Berat Balita Maksimum : ", max)
	fmt.Printf("Rerata Berat Balita : %.2f", rata2)
}
