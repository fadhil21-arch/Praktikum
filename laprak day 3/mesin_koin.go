package main
import "fmt"

func main() {
	var koin, koin_emas, koin_perak, koin_tembaga int
	fmt.Print("Masukan koin :")
	fmt.Scan(&koin)
	koin_emas = koin / 9
	koin_perak = (koin % 9) / 3
	koin_tembaga = (koin % 9) / 1
	fmt.Println("koin_emas:", koin_emas)
	fmt.Println("koin_perak:", koin_perak)
	fmt.Println("koin_tembaga:", koin_tembaga)
	
}
