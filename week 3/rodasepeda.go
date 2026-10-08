package main

import "fmt"

func main() {
	var keliling, jarak, r float64

	fmt.Print("masukan jari jari :")
	fmt.Scan(&r)
	keliling = 2.0 * 22 / 7 * r
	jarak = 10 * keliling
	fmt.Println("keliling :", keliling)
	fmt.Print("jarak: ", jarak)

}
