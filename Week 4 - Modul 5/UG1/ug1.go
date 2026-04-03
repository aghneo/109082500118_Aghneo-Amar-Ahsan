package main

import "fmt"

func bintang(n int) string {
	i := 1
	if i < n {
		fmt.Println("")
	} else {
		return "*"
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(bintang(n))
}