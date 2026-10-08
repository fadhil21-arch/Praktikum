package main

import "fmt"

func main() {
	var mk string = "algortima dan pemograman"
	var kode, sks int

	fmt.Print("tuliskan kode mk dan sks: ")
	fmt.Scan(&kode, &sks)
	fmt.Println("Kredit MK", kode, "-", mk, "1 adalah", sks, "sks")
}
