package main

import "fmt"

func main() {
	var vote int
	var totalSuara, suaraSah int
	var counts [21]int 

	for {
		fmt.Scan(&vote)
		if vote == 0 {
			break
		}
		
		totalSuara++
		
		if vote >= 1 && vote <= 20 {
			suaraSah++
			counts[vote]++
		}
	}

	var ketua, wakil int

	for i := 1; i <= 20; i++ {
		if counts[i] > counts[ketua] {
			wakil = ketua
			ketua = i
		} else if counts[i] > counts[wakil] {
			wakil = i
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	
	if ketua != 0 {
		fmt.Printf("Ketua RT: %d\n", ketua)
	}
	if wakil != 0 {
		fmt.Printf("Wakil ketua: %d\n", wakil)
	}
}