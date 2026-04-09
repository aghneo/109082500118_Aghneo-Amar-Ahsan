package main 

import "fmt"

const phi float64 = 3.14

func vol(r, t float64) (vol float64){
	vol = phi * r * r * t 
	return vol
}

func massa(r, t, p float64) (massa float64){
	massa = vol(r, t) * p
	return massa
}

func display(m1, m2 float64){
	if m1 == m2 {
		fmt.Println("BALANCE WELLLLL")
	} else {
		selisih := m1 - m2
		if selisih < 0 {
			selisih = selisih * (-1)
		} 
		fmt.Print("Selisih : ", selisih)
	}
}

func main(){
	var r float64
	var t1, p1 float64
	var t2, p2 float64

	fmt.Print("Masukan Jari - Jari : ")
	fmt.Scan(&r)
	fmt.Println("______________________________")
	fmt.Print("Masukan Tinggi Tabung Kiri : ")
	fmt.Scan(&t1)
	fmt.Print("Masukan Massa Jenis Zat Cair Kiri : ")
	fmt.Scan(&p1)
	fmt.Println("______________________________")
	fmt.Print("Masukan Tinggi Tabung Kanan : ")
	fmt.Scan(&t2)
	fmt.Print("Masukan Massa Jenis Zat Cair Kanan : ")
	fmt.Scan(&p2)
	fmt.Println("______________________________")

	m1 := massa(r, t1, p1)
	m2 := massa(r, t2, p2)

	display(m1, m2)
}