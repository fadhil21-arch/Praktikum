package main
import "fmt"

func main(){
	var alas, tinggi float64
	var luas float64

	fmt.Print("alas :")
	fmt.Scan(&alas)
	fmt.Print("tinggi :")
	fmt.Scan(&tinggi)

	luas = alas * tinggi / 2
	fmt.Println("alas : ", alas)
	fmt.Println("tinggi :", tinggi)
	fmt.Println("luas :", luas)
}