package main
import "fmt"

func main() {
	var kecepatan, menit, m1, jam, totaljarak int

	fmt.Scan(&kecepatan)
	totaljarak = 100+60+170
	menit = (totaljarak * 60) / kecepatan
	jam = menit / 60
	m1 = menit % 60

	fmt.Print(jam, m1)
}