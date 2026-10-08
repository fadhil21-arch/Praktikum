package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	// Perhitungan
	kecepatan := 60
	jarak := kecepatan * t
	bahanBakar := jarak / 12

	fmt.Println(jarak)
	fmt.Println(bahanBakar)
}