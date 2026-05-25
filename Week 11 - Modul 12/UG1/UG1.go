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

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	for i := 1; i <= 20; i++ {
		if counts[i] > 0 {
			fmt.Printf("%d: %d\n", i, counts[i])
		}
	}
}
