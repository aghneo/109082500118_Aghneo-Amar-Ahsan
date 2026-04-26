package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var c rune
	*n = 0
	for {
		_, err := fmt.Scanf("%c", &c)
		if err != nil || c == '.' || *n >= NMAX {
			break
		}
		if c != '\n' && c != '\r' {
			(*t)[*n] = c
			(*n)++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
}

func balikanArray(t *tabel, n int) {
	kiri := 0
	kanan := n - 1
	for kiri < kanan {
		temp := (*t)[kiri]
		(*t)[kiri] = (*t)[kanan]
		(*t)[kanan] = temp
		kiri++
		kanan--
	}
}

func palindrom(t tabel, n int) bool {
	tCopy := t
	
	balikanArray(&tCopy, n)

	for i := 0; i < n; i++ {
		if t[i] != tCopy[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Println("Masukkan Teks (Diakhiri dengan TITIK '.'):")
	isiArray(&tab, &m)

	fmt.Print("\nIsi array awal: ")
	cetakArray(tab, m)
	fmt.Println()

	isPalindrom := palindrom(tab, m)
	fmt.Printf("Apakah Palindrom? %v\n", isPalindrom)

	balikanArray(&tab, m)

	fmt.Print("Array dibalik: ")
	cetakArray(tab, m)
	fmt.Println()
}