package main

import (
	"fmt"
	"math"
)

func main() {

	var r float32
	fmt.Print("Masukan panjang jari-jari lingkaran: ")
	fmt.Scanln(&r)
	luas := math.Pi * r * r

	fmt.Printf("luas lingkaran  : %.2f cm2\n", luas)
}
