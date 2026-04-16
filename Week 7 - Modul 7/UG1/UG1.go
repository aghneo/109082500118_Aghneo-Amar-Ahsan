package main

import "fmt"

type suhu float64

func reamur(c suhu) suhu {
	return c * (4.0 / 5.0)
}

func fahrenheit(c suhu) suhu {
	return (c * (9.0 / 5.0)) + 32
}

func kelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukan Suhu Dalam Celcius : ")
	fmt.Scan(&c)

	fmt.Println(c, " Celcius : ", reamur(c), "Reamur")
	fmt.Println(c, " Celcius : ", fahrenheit(c), "Fahrenheit")
	fmt.Println(c, " Celcius : ", kelvin(c), "Kelvin")
}
