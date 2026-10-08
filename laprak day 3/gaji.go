package main
import "fmt"

func main() {
	var gb, jam, gp, bl, pt float64 	
	fmt.Print("Masukan Gaji Pokok :")
	fmt.Scan(&gp)
	fmt.Print("Masukan total Jam Lembur :")
	fmt.Scan(&jam)
	bl = (jam / 1) * 45000
	pt = (0.02 + 0.035) * gp
	gb = gp + bl - pt
	fmt.Printf("Gaji Bersih : Rp %.0f\n", gb)

}
