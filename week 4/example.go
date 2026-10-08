package main
import "fmt"

func main() {
	var jam, menit, detik, sisa int

	fmt.Scan(&detik)

	jam = detik / 3600
	sisa = detik % 3600
	menit = sisa / 60
	detik = sisa % 60
	fmt.Printf("%d jam %d menit %d detik\n", jam, menit, detik)

}