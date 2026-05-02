package main

import "fmt"

type array [1000]float64

func Terkecil(N int, arrayKelinci array) float64 {
	var min float64 = arrayKelinci[0]
	var i int = 1
	for i < N {
		if min > arrayKelinci[i] {
			min = arrayKelinci[i]
		}
		i++
	}
	return min
}

func Terbesar(N int, arrayKelinci array) float64 {
	var  max float64 = arrayKelinci[0]
	var i int = 1
	for i < N {
		if max < arrayKelinci[i] {
			max = arrayKelinci[i]
		}
		i++
	}
	return max
}

func main(){
	var N int
	var arrayKelinci array

	fmt.Print("Masukan Banyaknya Kelinci Yang Akan Ditimbang : ")
	fmt.Scan(&N)

	for i := 0 ; i < N ; i++ {
		fmt.Print("Masukan Berat Kelinci Ke-", i+1, " : ")
		fmt.Scan(&arrayKelinci[i])
	}

	min := Terkecil(N, arrayKelinci)
	max := Terbesar(N, arrayKelinci)

	fmt.Println("Berat Kelinci Terkecil Adalah : ", min)
	fmt.Println("Berat Kelinci Terbesar Adalah : ", max)
}