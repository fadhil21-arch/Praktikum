package main
import "fmt"

func main(){
	const panjang = 4
	const lebar = 2
	var volume, berat, tinggi int

	fmt.Print("masukan tinggi :")
	fmt.Scan(&tinggi)
	volume = panjang * lebar * tinggi 
	berat = 1000 * volume
	fmt.Println("volume :", volume)
	fmt.Print("berat :", berat)
}