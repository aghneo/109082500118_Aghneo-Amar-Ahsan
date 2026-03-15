package main

import "fmt"

func faktorial(n int) int{
	if n == 0 || n ==1 {
		return 1
	}
	hasil := 1
	for i := 2 ; i<= n ; i++ {
		hasil *= i
	}
	return hasil
}

func kombinasi(n, r int) int{
	return faktorial(n) / (faktorial(r) * faktorial(n - r))
}

func permutasi(n, r int) int{
	return faktorial(n) / (faktorial(n - r))
}

func main(){
	var a, b, c, d int

	fmt.Print("Masukan Angka A : ")
	fmt.Scan(&a)
	fmt.Print("Masukan Angka B : ")
	fmt.Scan(&b)
	fmt.Print("Masukan Angka C : ")
	fmt.Scan(&c)
	fmt.Print("Masukan Angka D : ")
	fmt.Scan(&d)

	if a >= c && b >= d {
		fmt.Println("Hasil : ")
		fmt.Printf("%d %d\n", permutasi(a, c), kombinasi(a, c))
		fmt.Printf("%d %d\n", permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println("Angka Tidak Valid")
	}
}